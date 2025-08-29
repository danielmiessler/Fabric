"""
Safe import utilities with automatic dependency management.
Provides decorators and utilities for handling optional dependencies gracefully.
"""
import functools
from typing import Any, Callable, Optional, TypeVar, Union
import streamlit as st

from utils.logging import logger

T = TypeVar('T')


def requires_dependency(
    module_name: str, 
    package_name: str = None, 
    fallback_message: str = None,
    auto_install: bool = True
):
    """
    Decorator that ensures a dependency is available before executing a function.
    
    Args:
        module_name: Python module name to check
        package_name: Pip package name (defaults to module_name)
        fallback_message: Message to show if dependency is missing
        auto_install: Whether to attempt automatic installation
    """
    def decorator(func: Callable[..., T]) -> Callable[..., Optional[T]]:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            from services.dependencies import check_and_install_if_missing
            
            pkg_name = package_name or module_name
            
            if check_and_install_if_missing(module_name, pkg_name, auto_install):
                return func(*args, **kwargs)
            else:
                message = fallback_message or f"Function requires {pkg_name} - please install it"
                logger.warning(f"Function {func.__name__} skipped: {message}")
                
                # Show user-friendly message in Streamlit
                try:
                    st.warning(f"⚠️ {message}")
                    if st.button(f"📦 Install {pkg_name}", key=f"install_{pkg_name}_{func.__name__}"):
                        if check_and_install_if_missing(module_name, pkg_name, True):
                            st.success(f"✅ {pkg_name} installed! Please refresh the page.")
                            st.rerun()
                        else:
                            st.error(f"❌ Failed to install {pkg_name}")
                except:
                    # Not in Streamlit context, just log
                    pass
                
                return None
        
        return wrapper
    return decorator


def optional_import_with_fallback(
    module_name: str,
    fallback_function: Callable = None,
    error_message: str = None
):
    """
    Decorator for functions that use optional imports with fallback behavior.
    
    Args:
        module_name: Module to import
        fallback_function: Function to call if module not available
        error_message: Custom error message
    """
    def decorator(func: Callable[..., T]) -> Callable[..., T]:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            from services.dependencies import check_and_install_if_missing
            
            if check_and_install_if_missing(module_name, module_name):
                return func(*args, **kwargs)
            else:
                if fallback_function:
                    logger.info(f"Using fallback for {func.__name__} due to missing {module_name}")
                    return fallback_function(*args, **kwargs)
                else:
                    message = error_message or f"Function {func.__name__} requires {module_name}"
                    logger.warning(message)
                    raise ImportError(message)
        
        return wrapper
    return decorator


class LazyImport:
    """Lazy import class that attempts import on first access."""
    
    def __init__(self, module_name: str, package_name: str = None, auto_install: bool = True):
        self.module_name = module_name
        self.package_name = package_name or module_name
        self.auto_install = auto_install
        self._module = None
        self._attempted = False
    
    def __getattr__(self, name: str) -> Any:
        if not self._attempted:
            self._attempt_import()
        
        if self._module is None:
            raise AttributeError(f"Module {self.module_name} not available")
        
        return getattr(self._module, name)
    
    def _attempt_import(self):
        """Attempt to import the module."""
        self._attempted = True
        
        from services.dependencies import check_and_install_if_missing
        
        if check_and_install_if_missing(self.module_name, self.package_name, self.auto_install):
            try:
                import importlib
                self._module = importlib.import_module(self.module_name)
                logger.debug(f"Lazy import successful: {self.module_name}")
            except ImportError as e:
                logger.warning(f"Lazy import failed: {self.module_name} - {e}")
        else:
            logger.warning(f"Lazy import dependency not available: {self.module_name}")
    
    @property
    def available(self) -> bool:
        """Check if the module is available."""
        if not self._attempted:
            self._attempt_import()
        return self._module is not None


# Pre-configured lazy imports for common optional dependencies
plotly = LazyImport("plotly", "plotly>=5.0.0")
sentence_transformers = LazyImport("sentence_transformers", "sentence-transformers") 
cachetools = LazyImport("cachetools", "cachetools>=5.0.0")
pyperclip = LazyImport("pyperclip", "pyperclip>=1.8.0")


def with_dependency_check(dependencies: Union[str, List[str]]):
    """
    Class decorator that checks dependencies before class instantiation.
    
    Args:
        dependencies: Single dependency or list of dependencies to check
    """
    if isinstance(dependencies, str):
        dependencies = [dependencies]
    
    def class_decorator(cls):
        original_init = cls.__init__
        
        @functools.wraps(original_init)
        def new_init(self, *args, **kwargs):
            # Check all dependencies before initializing
            missing_deps = []
            
            for dep in dependencies:
                from services.dependencies import check_and_install_if_missing
                if not check_and_install_if_missing(dep, dep):
                    missing_deps.append(dep)
            
            if missing_deps:
                logger.error(f"Cannot initialize {cls.__name__}: missing {missing_deps}")
                raise ImportError(f"Missing required dependencies: {missing_deps}")
            
            original_init(self, *args, **kwargs)
        
        cls.__init__ = new_init
        return cls
    
    return class_decorator


def ensure_streamlit_components():
    """Ensure Streamlit and its components are available."""
    from services.dependencies import check_and_install_if_missing
    
    required = [
        ("streamlit", "streamlit>=1.27.0"),
        ("pandas", "pandas>=1.5.0") 
    ]
    
    missing = []
    for module_name, package_spec in required:
        if not check_and_install_if_missing(module_name, package_spec):
            missing.append(package_spec)
    
    if missing:
        error_msg = f"Critical Streamlit dependencies missing: {missing}"
        logger.critical(error_msg)
        raise ImportError(error_msg)
    
    return True
