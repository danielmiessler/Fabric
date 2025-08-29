"""
Dependency management service for automatic installation of missing packages.
Provides fallback mechanisms and graceful degradation when dependencies are unavailable.
"""
import subprocess
import sys
import importlib
from typing import Dict, List, Optional, Tuple, Any
from dataclasses import dataclass
from pathlib import Path

from utils.logging import logger


@dataclass
class DependencySpec:
    """Specification for a dependency with fallback options."""
    name: str
    package: str  # pip package name
    version: Optional[str] = None
    optional: bool = False
    fallback_message: Optional[str] = None
    install_command: Optional[str] = None


class DependencyManager:
    """Manages dependency checking, installation, and fallbacks."""
    
    # Core dependencies that must be present
    CORE_DEPENDENCIES = {
        "streamlit": DependencySpec("streamlit", "streamlit", ">=1.27.0"),
        "pandas": DependencySpec("pandas", "pandas", ">=1.5.0"),
        "numpy": DependencySpec("numpy", "numpy", ">=1.23.0"),
    }
    
    # Optional dependencies with fallbacks
    OPTIONAL_DEPENDENCIES = {
        "plotly": DependencySpec(
            "plotly", 
            "plotly", 
            ">=5.0.0",
            optional=True,
            fallback_message="Advanced charts disabled - install plotly for enhanced visualizations"
        ),
        "pyperclip": DependencySpec(
            "pyperclip", 
            "pyperclip", 
            ">=1.8.0",
            optional=True,
            fallback_message="Clipboard functionality disabled - install pyperclip for copy/paste features"
        ),
        "python-dotenv": DependencySpec(
            "dotenv", 
            "python-dotenv", 
            ">=1.0.0",
            optional=True,
            fallback_message="Environment file loading disabled - install python-dotenv for .env support"
        ),
        "matplotlib": DependencySpec(
            "matplotlib", 
            "matplotlib", 
            ">=3.5.0",
            optional=True,
            fallback_message="Basic plotting disabled - install matplotlib for charts"
        ),
        "seaborn": DependencySpec(
            "seaborn", 
            "seaborn", 
            ">=0.12.0",
            optional=True,
            fallback_message="Advanced plotting disabled - install seaborn for statistical charts"
        ),
        "cachetools": DependencySpec(
            "cachetools",
            "cachetools",
            ">=5.0.0", 
            optional=True,
            fallback_message="Advanced caching disabled - install cachetools for better performance"
        ),
        "sentence-transformers": DependencySpec(
            "sentence_transformers",
            "sentence-transformers",
            optional=True,
            fallback_message="Semantic search disabled - install sentence-transformers for AI-powered search"
        )
    }
    
    def __init__(self, auto_install: bool = True, user_confirm: bool = True):
        """
        Initialize dependency manager.
        
        Args:
            auto_install: Whether to automatically install missing dependencies
            user_confirm: Whether to ask user confirmation before installing
        """
        self._auto_install = auto_install
        self._user_confirm = user_confirm
        self._availability_cache: Dict[str, bool] = {}
        self._fallback_messages: List[str] = []
        
    def check_all_dependencies(self) -> Dict[str, Any]:
        """
        Check all dependencies and return status report.
        
        Returns:
            Dictionary with dependency status information
        """
        report = {
            "core_status": {},
            "optional_status": {},
            "missing_core": [],
            "missing_optional": [],
            "fallback_messages": [],
            "installation_commands": []
        }
        
        # Check core dependencies
        for name, spec in self.CORE_DEPENDENCIES.items():
            available = self.is_available(spec.name)
            report["core_status"][name] = available
            
            if not available:
                report["missing_core"].append(spec)
                report["installation_commands"].append(f"pip install {spec.package}")
        
        # Check optional dependencies
        for name, spec in self.OPTIONAL_DEPENDENCIES.items():
            available = self.is_available(spec.name)
            report["optional_status"][name] = available
            
            if not available:
                report["missing_optional"].append(spec)
                if spec.fallback_message:
                    report["fallback_messages"].append(spec.fallback_message)
        
        return report
    
    def is_available(self, module_name: str) -> bool:
        """
        Check if a module is available for import.
        
        Args:
            module_name: Name of the module to check
            
        Returns:
            True if module can be imported, False otherwise
        """
        if module_name in self._availability_cache:
            return self._availability_cache[module_name]
        
        try:
            importlib.import_module(module_name)
            self._availability_cache[module_name] = True
            return True
        except ImportError:
            self._availability_cache[module_name] = False
            return False
    
    def install_missing_dependencies(self, dependencies: List[DependencySpec]) -> Dict[str, bool]:
        """
        Install missing dependencies.
        
        Args:
            dependencies: List of dependency specifications to install
            
        Returns:
            Dictionary mapping package names to installation success status
        """
        results = {}
        
        for spec in dependencies:
            if self.is_available(spec.name):
                results[spec.package] = True
                continue
            
            logger.info(f"Installing missing dependency: {spec.package}")
            
            try:
                success = self._install_package(spec)
                results[spec.package] = success
                
                if success:
                    # Clear cache and recheck
                    self._availability_cache.pop(spec.name, None)
                    available = self.is_available(spec.name)
                    if available:
                        logger.info(f"Successfully installed and verified: {spec.package}")
                    else:
                        logger.warning(f"Installation reported success but module still not importable: {spec.name}")
                        results[spec.package] = False
                else:
                    logger.error(f"Failed to install: {spec.package}")
                    
            except Exception as e:
                logger.error(f"Exception installing {spec.package}: {e}")
                results[spec.package] = False
        
        return results
    
    def _install_package(self, spec: DependencySpec) -> bool:
        """
        Install a single package using pip.
        
        Args:
            spec: Dependency specification
            
        Returns:
            True if installation succeeded, False otherwise
        """
        try:
            # Build pip install command
            cmd = [sys.executable, "-m", "pip", "install"]
            
            if spec.version:
                package_spec = f"{spec.package}{spec.version}"
            else:
                package_spec = spec.package
            
            cmd.append(package_spec)
            
            # Add common pip flags for better user experience
            cmd.extend(["--user", "--no-warn-script-location"])
            
            logger.info(f"Running: {' '.join(cmd)}")
            
            # Execute installation
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                timeout=300  # 5 minute timeout
            )
            
            if result.returncode == 0:
                logger.info(f"Successfully installed {spec.package}")
                return True
            else:
                logger.error(f"Installation failed for {spec.package}: {result.stderr}")
                return False
                
        except subprocess.TimeoutExpired:
            logger.error(f"Installation timeout for {spec.package}")
            return False
        except Exception as e:
            logger.error(f"Installation exception for {spec.package}: {e}")
            return False
    
    def ensure_core_dependencies(self) -> bool:
        """
        Ensure all core dependencies are available, installing if necessary.
        
        Returns:
            True if all core dependencies are available, False otherwise
        """
        missing_core = []
        
        for spec in self.CORE_DEPENDENCIES.values():
            if not self.is_available(spec.name):
                missing_core.append(spec)
        
        if not missing_core:
            return True
        
        logger.warning(f"Missing core dependencies: {[spec.name for spec in missing_core]}")
        
        if self._auto_install:
            logger.info("Attempting to install missing core dependencies...")
            results = self.install_missing_dependencies(missing_core)
            
            # Check if all installations succeeded
            return all(results.values())
        else:
            logger.error("Auto-install disabled, core dependencies missing")
            return False
    
    def get_fallback_messages(self) -> List[str]:
        """Get list of fallback messages for missing optional dependencies."""
        messages = []
        
        for spec in self.OPTIONAL_DEPENDENCIES.values():
            if not self.is_available(spec.name) and spec.fallback_message:
                messages.append(spec.fallback_message)
        
        return messages
    
    def install_requirements_txt(self, requirements_file: Optional[Path] = None) -> bool:
        """
        Install dependencies from requirements.txt file.
        
        Args:
            requirements_file: Path to requirements.txt (defaults to project requirements.txt)
            
        Returns:
            True if installation succeeded, False otherwise
        """
        try:
            if requirements_file is None:
                # Default to project requirements.txt
                script_dir = Path(__file__).parent.parent
                requirements_file = script_dir / "requirements.txt"
            
            if not requirements_file.exists():
                logger.warning(f"Requirements file not found: {requirements_file}")
                return False
            
            logger.info(f"Installing from requirements.txt: {requirements_file}")
            
            cmd = [
                sys.executable, "-m", "pip", "install",
                "-r", str(requirements_file),
                "--user", "--no-warn-script-location"
            ]
            
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                timeout=600  # 10 minute timeout for full requirements
            )
            
            if result.returncode == 0:
                logger.info("Successfully installed requirements.txt dependencies")
                # Clear cache to recheck availability
                self._availability_cache.clear()
                return True
            else:
                logger.error(f"Requirements installation failed: {result.stderr}")
                return False
                
        except Exception as e:
            logger.error(f"Exception installing requirements.txt: {e}")
            return False
    
    def create_installation_script(self) -> str:
        """
        Create a shell script for manual dependency installation.
        
        Returns:
            Shell script content for installing dependencies
        """
        script_lines = [
            "#!/bin/bash",
            "# Fabric Pattern Studio Dependency Installation Script",
            "# Generated automatically by the dependency manager",
            "",
            "set -e  # Exit on any error",
            "",
            "echo '🎭 Installing Fabric Pattern Studio dependencies...'",
            ""
        ]
        
        # Add core dependencies
        script_lines.append("echo '📦 Installing core dependencies...'")
        for spec in self.CORE_DEPENDENCIES.values():
            package_spec = f"{spec.package}{spec.version}" if spec.version else spec.package
            script_lines.append(f'python3 -m pip install --user "{package_spec}"')
        
        script_lines.append("")
        
        # Add optional dependencies
        script_lines.append("echo '✨ Installing optional dependencies...'")
        for spec in self.OPTIONAL_DEPENDENCIES.values():
            package_spec = f"{spec.package}{spec.version}" if spec.version else spec.package
            script_lines.append(f'python3 -m pip install --user "{package_spec}" || echo "⚠️ Optional: {spec.package} installation failed"')
        
        script_lines.extend([
            "",
            "echo '✅ Dependency installation completed!'",
            "echo 'You can now run: ./run.sh'"
        ])
        
        return "\n".join(script_lines)


# Context managers for graceful fallbacks
class OptionalImport:
    """Context manager for optional imports with fallbacks."""
    
    def __init__(self, module_name: str, fallback_message: str = None):
        self.module_name = module_name
        self.fallback_message = fallback_message or f"{module_name} not available"
        self.module = None
        self.available = False
    
    def __enter__(self):
        try:
            self.module = importlib.import_module(self.module_name)
            self.available = True
            return self.module
        except ImportError:
            self.available = False
            logger.debug(self.fallback_message)
            return None
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        pass


def safe_import(module_name: str, fallback_value: Any = None) -> Any:
    """
    Safely import a module with fallback value.
    
    Args:
        module_name: Name of module to import
        fallback_value: Value to return if import fails
        
    Returns:
        Imported module or fallback value
    """
    try:
        return importlib.import_module(module_name)
    except ImportError:
        logger.debug(f"Optional module {module_name} not available, using fallback")
        return fallback_value


def check_and_install_if_missing(
    module_name: str, 
    package_name: str, 
    auto_install: bool = True
) -> bool:
    """
    Check if module exists and optionally install if missing.
    
    Args:
        module_name: Python module name for importing
        package_name: Pip package name for installation
        auto_install: Whether to attempt automatic installation
        
    Returns:
        True if module is available after check/install
    """
    try:
        importlib.import_module(module_name)
        return True
    except ImportError:
        if auto_install:
            logger.info(f"Installing missing dependency: {package_name}")
            try:
                result = subprocess.run(
                    [sys.executable, "-m", "pip", "install", "--user", package_name],
                    capture_output=True,
                    text=True,
                    timeout=120
                )
                
                if result.returncode == 0:
                    logger.info(f"Successfully installed {package_name}")
                    # Try import again
                    try:
                        importlib.import_module(module_name)
                        return True
                    except ImportError:
                        logger.error(f"Module {module_name} still not available after installation")
                        return False
                else:
                    logger.error(f"Installation failed for {package_name}: {result.stderr}")
                    return False
            except Exception as e:
                logger.error(f"Exception during installation of {package_name}: {e}")
                return False
        else:
            logger.warning(f"Module {module_name} not available and auto-install disabled")
            return False


# Singleton instance
_dependency_manager = None


def get_dependency_manager() -> DependencyManager:
    """Get singleton DependencyManager instance."""
    global _dependency_manager
    if _dependency_manager is None:
        _dependency_manager = DependencyManager()
    return _dependency_manager


def ensure_dependencies() -> bool:
    """
    Ensure all required dependencies are available.
    
    Returns:
        True if all core dependencies are available
    """
    manager = get_dependency_manager()
    return manager.ensure_core_dependencies()
