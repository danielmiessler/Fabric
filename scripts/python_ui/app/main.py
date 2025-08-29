"""
Fabric Pattern Studio - Clean Architecture Entry Point (PR-1)
"""
import os, sys
import streamlit as st

# Ensure package imports resolve when running via 'streamlit run app/main.py'
BASE_DIR = os.path.dirname(os.path.dirname(__file__))  # scripts/python_ui
if BASE_DIR not in sys.path:
    sys.path.insert(0, BASE_DIR)

# Early dependency check before importing application modules
try:
    from utils.safe_imports import ensure_streamlit_components
    ensure_streamlit_components()
except ImportError as e:
    st.error(f"❌ Critical dependencies missing: {e}")
    st.code("pip install streamlit>=1.27.0 pandas>=1.5.0", language="bash")
    st.stop()
except Exception as e:
    # Don't let dependency check completely break startup
    st.warning(f"⚠️ Dependency check failed: {e}")

from utils import errors, logging as app_logging  # noqa: E402
from app import routing, state  # noqa: E402
from components import header, sidebar  # noqa: E402
from views import execution, management, dashboard  # noqa: E402

def configure_page() -> None:
    st.set_page_config(
        page_title="Fabric Pattern Studio",
        page_icon="🎭",
        layout="wide",
        initial_sidebar_state="expanded",
    )

@errors.ui_error_boundary
def main() -> None:
    app_logging.init()
    
    # Check and install missing dependencies on startup
    try:
        from services.dependencies import ensure_dependencies, get_dependency_manager
        
        if not ensure_dependencies():
            st.error("❌ Critical dependencies missing. Please install requirements manually.")
            st.code("pip install -r requirements.txt", language="bash")
            st.stop()
        
        # Show any fallback messages for optional dependencies
        manager = get_dependency_manager()
        fallback_messages = manager.get_fallback_messages()
        if fallback_messages:
            with st.expander("⚠️ Optional Features Disabled", expanded=False):
                for message in fallback_messages:
                    st.warning(message)
                st.info("💡 Run `pip install -r requirements.txt` to enable all features")
    
    except Exception as e:
        # Don't let dependency checking break the app
        app_logging.logger.warning(f"Dependency check failed: {e}")
    
    configure_page()
    state.initialize()      # defaults + future persistence hooks
    header.render()
    sidebar.render()

    view = routing.get_current_view()
    if view == "Run Patterns":
        execution.render()
    elif view == "Pattern Management":
        management.render()
    elif view == "Analysis Dashboard":
        dashboard.render()
    else:
        # Default to dashboard if view is unknown
        dashboard.render()

if __name__ == "__main__":
    main()