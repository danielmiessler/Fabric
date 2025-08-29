"""
Unit tests for dependency management service.
"""
import pytest
from unittest.mock import Mock, patch, MagicMock
import subprocess

from services.dependencies import (
    DependencyManager, DependencySpec, OptionalImport,
    safe_import, check_and_install_if_missing, get_dependency_manager
)


@pytest.fixture
def dependency_manager():
    """Create a fresh DependencyManager for testing."""
    return DependencyManager(auto_install=False, user_confirm=False)


def test_dependency_spec():
    """Test DependencySpec dataclass."""
    spec = DependencySpec(
        name="test_module",
        package="test-package", 
        version=">=1.0.0",
        optional=True,
        fallback_message="Test fallback"
    )
    
    assert spec.name == "test_module"
    assert spec.package == "test-package"
    assert spec.version == ">=1.0.0"
    assert spec.optional is True
    assert spec.fallback_message == "Test fallback"


def test_is_available_existing_module(dependency_manager):
    """Test checking for an existing module."""
    # Test with a module that should always exist
    assert dependency_manager.is_available("os") is True
    
    # Should cache the result
    assert "os" in dependency_manager._availability_cache
    assert dependency_manager._availability_cache["os"] is True


def test_is_available_missing_module(dependency_manager):
    """Test checking for a non-existent module."""
    assert dependency_manager.is_available("nonexistent_module_xyz") is False
    
    # Should cache the result
    assert "nonexistent_module_xyz" in dependency_manager._availability_cache
    assert dependency_manager._availability_cache["nonexistent_module_xyz"] is False


@patch('subprocess.run')
def test_install_package_success(mock_run, dependency_manager):
    """Test successful package installation."""
    # Mock successful pip install
    mock_run.return_value = Mock(returncode=0, stderr="", stdout="Successfully installed test-package")
    
    spec = DependencySpec("test_module", "test-package")
    success = dependency_manager._install_package(spec)
    
    assert success is True
    mock_run.assert_called_once()
    
    # Check that pip install command was constructed correctly
    call_args = mock_run.call_args[0][0]
    assert "pip" in call_args
    assert "install" in call_args
    assert "test-package" in call_args


@patch('subprocess.run')
def test_install_package_failure(mock_run, dependency_manager):
    """Test failed package installation."""
    # Mock failed pip install
    mock_run.return_value = Mock(returncode=1, stderr="Installation failed", stdout="")
    
    spec = DependencySpec("test_module", "test-package")
    success = dependency_manager._install_package(spec)
    
    assert success is False


@patch('subprocess.run')
def test_install_package_timeout(mock_run, dependency_manager):
    """Test package installation timeout."""
    # Mock timeout
    mock_run.side_effect = subprocess.TimeoutExpired("pip", 300)
    
    spec = DependencySpec("test_module", "test-package")
    success = dependency_manager._install_package(spec)
    
    assert success is False


def test_check_all_dependencies(dependency_manager):
    """Test comprehensive dependency checking."""
    with patch.object(dependency_manager, 'is_available') as mock_available:
        # Mock some dependencies as available, others not
        mock_available.side_effect = lambda name: name in ["streamlit", "pandas"]
        
        report = dependency_manager.check_all_dependencies()
        
        assert "core_status" in report
        assert "optional_status" in report
        assert "missing_core" in report
        assert "missing_optional" in report
        assert "fallback_messages" in report


def test_optional_import_context_manager():
    """Test OptionalImport context manager."""
    # Test with existing module
    with OptionalImport("os", "OS not available") as os_module:
        assert os_module is not None
        assert hasattr(os_module, 'path')  # os module should have path
    
    # Test with non-existent module
    with OptionalImport("nonexistent_module_xyz", "Module not available") as missing_module:
        assert missing_module is None


def test_safe_import_existing():
    """Test safe_import with existing module."""
    os_module = safe_import("os")
    assert os_module is not None
    assert hasattr(os_module, 'path')


def test_safe_import_missing_with_fallback():
    """Test safe_import with missing module and fallback."""
    fallback_value = "fallback"
    result = safe_import("nonexistent_module_xyz", fallback_value)
    assert result == fallback_value


def test_check_and_install_if_missing_auto_install():
    """Test automatic installation of missing dependency."""
    with patch('services.dependencies.subprocess.run') as mock_run, \
         patch('services.dependencies.importlib.import_module') as mock_import:
        
        # First import fails (module missing), second succeeds (after install)
        mock_import.side_effect = [ImportError(), Mock()]
        mock_run.return_value = Mock(returncode=0, stderr="", stdout="Successfully installed test-package")
        
        result = check_and_install_if_missing("test_module", "test-package", auto_install=True)
        
        assert result is True
        mock_run.assert_called_once()
        assert mock_import.call_count == 2  # First check + verification after install


@patch('importlib.import_module')
def test_check_and_install_if_missing_no_auto_install(mock_import):
    """Test check without auto-install when module is missing."""
    mock_import.side_effect = ImportError()
    
    result = check_and_install_if_missing("test_module", "test-package", auto_install=False)
    
    assert result is False
    assert mock_import.call_count == 1  # Only initial check


def test_singleton_pattern():
    """Test that get_dependency_manager returns the same instance."""
    manager1 = get_dependency_manager()
    manager2 = get_dependency_manager()
    assert manager1 is manager2


def test_create_installation_script(dependency_manager):
    """Test installation script generation."""
    script_content = dependency_manager.create_installation_script()
    
    assert "#!/bin/bash" in script_content
    assert "pip install" in script_content
    assert "streamlit" in script_content  # Should include core dependencies
    assert "plotly" in script_content  # Should include optional dependencies


@patch('subprocess.run')
def test_install_requirements_txt_success(mock_run, dependency_manager, tmp_path):
    """Test installing from requirements.txt file."""
    # Create temporary requirements.txt
    requirements_file = tmp_path / "requirements.txt"
    requirements_file.write_text("streamlit>=1.27.0\npandas>=1.5.0\n")
    
    # Mock successful installation
    mock_run.return_value = Mock(returncode=0, stderr="", stdout="")
    
    success = dependency_manager.install_requirements_txt(requirements_file)
    
    assert success is True
    mock_run.assert_called_once()
    
    # Check command structure
    call_args = mock_run.call_args[0][0]
    assert "-r" in call_args
    assert str(requirements_file) in call_args


@patch('subprocess.run')
def test_install_requirements_txt_missing_file(mock_run, dependency_manager, tmp_path):
    """Test handling of missing requirements.txt file."""
    missing_file = tmp_path / "missing_requirements.txt"
    
    success = dependency_manager.install_requirements_txt(missing_file)
    
    assert success is False
    mock_run.assert_not_called()
