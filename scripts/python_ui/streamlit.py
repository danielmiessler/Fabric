import shutil
import json
import os
import streamlit as st
from subprocess import run, CalledProcessError
from dotenv import load_dotenv
import re
import time
import logging
from typing import Dict, List, Optional, Tuple
from datetime import datetime
import sys
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import numpy as np
from io import StringIO
import uuid

# Create formatters
console_formatter = logging.Formatter(
    "\033[92m%(asctime)s\033[0m - "  # Green timestamp
    "\033[94m%(levelname)s\033[0m - "  # Blue level
    "\033[95m[%(funcName)s]\033[0m "  # Purple function name
    "%(message)s"  # Regular message
)
file_formatter = logging.Formatter(
    "%(asctime)s - %(levelname)s - [%(funcName)s] %(message)s"
)

# Configure root logger
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

# Clear any existing handlers
logger.handlers = []

# Console Handler
console_handler = logging.StreamHandler(sys.stdout)
console_handler.setFormatter(console_formatter)
console_handler.setLevel(logging.INFO)
logger.addHandler(console_handler)

# File Handler
log_dir = os.path.expanduser("~/.config/fabric/logs")
os.makedirs(log_dir, exist_ok=True)
log_file = os.path.join(log_dir, f"fabric_ui_{datetime.now().strftime('%Y%m%d')}.log")
file_handler = logging.FileHandler(log_file)
file_handler.setFormatter(file_formatter)
file_handler.setLevel(logging.DEBUG)  # More detailed logging in file
logger.addHandler(file_handler)

# Detect operating system
PLATFORM = sys.platform

# Log startup message
logger.info("🚀 Fabric UI Starting Up")
logger.info(f"💾 Log file: {log_file}")
logger.info(f"🖥️ Platform detected: {PLATFORM}")

# Global variables
pattern_dir = os.path.expanduser("~/.config/fabric/patterns")
MAX_RETRIES = 3
RETRY_DELAY = 1  # seconds


def initialize_session_state():
    """Initialize necessary session state attributes.

    Error handling:
    - Ensures all required session state variables are initialized
    - Loads saved outputs from persistent storage
    - Handles missing or corrupted saved output files
    """
    logger.info("Initializing session state")
    default_configs = {
        # Configuration state
        "config_loaded": False,
        "vendors": {},
        "available_models": [],
        "selected_vendor": None,
        "selected_model": None,
        # Pattern execution state
        "input_content": "",
        "selected_patterns": [],
        "chat_output": [],
        "current_view": "run",
        # Pattern creation state
        "wizard_step": "Basic Info",
        "session_name": "",
        "context_name": "",
        # Model configuration
        "config": {"vendor": "", "model": "", "context_length": "2048"},
        # Model caching
        "cached_models": None,
        "last_model_fetch": 0,
        # UI state
        "active_tab": 0,
        # Output management
        "output_logs": [],
        "starred_outputs": [],
        "starring_output": None,
        "temp_star_name": "",
        # Enhanced UI state
        "show_advanced_options": False,
        "execution_mode": "standard",
        "current_session_id": str(uuid.uuid4()),
        "user_preferences": {
            "auto_save": True,
            "show_timestamps": True,
            "compact_view": False,
            "enable_notifications": True
        },
        # New interaction states
        "pattern_feedback": {},
        "execution_stats": {
            "total_runs": 0,
            "successful_runs": 0,
            "failed_runs": 0,
            "avg_execution_time": 0
        }
    }

    for key, value in default_configs.items():
        if key not in st.session_state:
            st.session_state[key] = value

    # Load saved outputs if they exist
    load_saved_outputs()


def parse_models_output(output: str) -> Dict[str, List[str]]:
    """Parse the output of fabric --listmodels command."""
    logger.debug("Parsing models output")
    providers = {}
    current_provider = None

    lines = output.split("\n")
    for line in lines:
        line = line.strip()
        if not line:
            continue

        if line == "Available models:":
            continue

        if not line.startswith("\t") and not line.startswith("["):
            current_provider = line.strip()
            providers[current_provider] = []
        elif current_provider and (line.startswith("\t") or line.startswith("[")):
            model = line.strip()
            if "[" in model and "]" in model:
                model = model.split("]", 1)[1].strip()
            providers[current_provider].append(model)

    logger.debug(f"Found providers: {list(providers.keys())}")
    return providers


def safe_run_command(command: List[str], retry: bool = True) -> Tuple[bool, str, str]:
    """Safely run a command with retries."""
    cmd_str = " ".join(command)
    logger.info(f"Executing command: {cmd_str}")

    for attempt in range(MAX_RETRIES if retry else 1):
        try:
            logger.debug(f"Attempt {attempt + 1}/{MAX_RETRIES if retry else 1}")
            result = run(command, capture_output=True, text=True)
            if result.returncode == 0:
                logger.debug("Command executed successfully")
                return True, result.stdout, ""
            if attempt == MAX_RETRIES - 1 or not retry:
                logger.error(
                    f"Command failed with return code {result.returncode}: {result.stderr}"
                )
                return False, "", result.stderr
        except Exception as e:
            if attempt == MAX_RETRIES - 1 or not retry:
                logger.error(f"Command execution failed: {str(e)}")
                return False, "", str(e)
        logger.debug(f"Retrying in {RETRY_DELAY} seconds...")
        time.sleep(RETRY_DELAY)
    logger.error("Max retries exceeded")
    return False, "", "Max retries exceeded"


def fetch_models_once() -> Dict[str, List[str]]:
    """Fetch models once and cache the results."""
    logger.info("Fetching models")
    current_time = time.time()
    cache_timeout = 300  # 5 minutes

    if (
        st.session_state.cached_models is not None
        and current_time - st.session_state.last_model_fetch < cache_timeout
    ):
        logger.debug("Using cached models")
        return st.session_state.cached_models

    logger.debug("Cache expired or not available, fetching new models")
    success, stdout, stderr = safe_run_command(["fabric", "--listmodels"])
    if not success:
        logger.error(f"Failed to fetch models: {stderr}")
        st.error(f"Failed to fetch models: {stderr}")
        return {}

    providers = parse_models_output(stdout)
    logger.info(f"Found {len(providers)} providers")
    st.session_state.cached_models = providers
    st.session_state.last_model_fetch = current_time
    return providers


def get_configured_providers() -> Dict[str, List[str]]:
    """Get list of configured providers using fabric --listmodels."""
    return fetch_models_once()


def update_provider_selection(new_provider: str) -> None:
    """Update provider and reset related states."""
    logger.info(f"Updating provider selection to: {new_provider}")
    if new_provider != st.session_state.config["vendor"]:
        logger.debug("Provider changed, resetting model selection")
        st.session_state.config["vendor"] = new_provider
        st.session_state.selected_vendor = new_provider
        st.session_state.config["model"] = None
        st.session_state.selected_model = None
        st.session_state.available_models = []
        if "model_select" in st.session_state:
            del st.session_state.model_select
        logger.debug("Model state reset completed")


def load_configuration() -> bool:
    """Load environment variables and initialize configuration."""
    logger.info("Loading configuration")
    try:
        env_path = os.path.expanduser("~/.config/fabric/.env")
        logger.debug(f"Looking for .env file at: {env_path}")

        if not os.path.exists(env_path):
            logger.error(f"Configuration file not found at {env_path}")
            st.error(f"Configuration file not found at {env_path}")
            return False

        load_dotenv(dotenv_path=env_path)
        logger.debug("Environment variables loaded")

        with st.spinner("Loading providers and models..."):
            providers = get_configured_providers()

        if not providers:
            logger.error("No providers configured")
            st.error("No providers configured. Please run 'fabric --setup' first.")
            return False

        default_vendor = os.getenv("DEFAULT_VENDOR")
        default_model = os.getenv("DEFAULT_MODEL")
        context_length = os.getenv("DEFAULT_MODEL_CONTEXT_LENGTH", "2048")

        logger.debug(
            f"Default configuration - Vendor: {default_vendor}, Model: {default_model}"
        )

        if not default_vendor or default_vendor not in providers:
            default_vendor = next(iter(providers))
            default_model = (
                providers[default_vendor][0] if providers[default_vendor] else None
            )
            logger.info(
                f"Using fallback configuration - Vendor: {default_vendor}, Model: {default_model}"
            )

        st.session_state.config = {
            "vendor": default_vendor,
            "model": default_model,
            "context_length": context_length,
        }
        st.session_state.vendors = providers
        st.session_state.config_loaded = True

        logger.info("Configuration loaded successfully")
        return True

    except Exception as e:
        logger.error(f"Configuration error: {str(e)}", exc_info=True)
        st.error(f"Configuration error: {str(e)}")
        return False


def load_models_and_providers() -> None:
    """Load models and providers from fabric configuration with enhanced UI."""
    try:
        # Enhanced sidebar header with status indicator
        col1, col2 = st.sidebar.columns([3, 1])
        with col1:
            st.header("🤖 AI Configuration")
        with col2:
            if st.session_state.config.get("model"):
                st.success("✓")
            else:
                st.error("⚠️")

        providers: Dict[str, List[str]] = fetch_models_once()

        if not providers:
            st.sidebar.error("No providers configured")
            return

        current_vendor = st.session_state.config.get("vendor", "")
        available_providers = list(providers.keys())

        try:
            provider_index = (
                available_providers.index(current_vendor)
                if current_vendor in available_providers
                else 0
            )
        except ValueError:
            provider_index = 0
            logger.warning(
                f"Current vendor {current_vendor} not found in available providers"
            )

        # Use segmented control for provider selection if available, otherwise fallback to selectbox
        try:
            if len(available_providers) <= 4:
                selected_provider = st.sidebar.segmented_control(
                    "Provider",
                    available_providers,
                    default=available_providers[provider_index] if available_providers else None,
                    key="provider_segmented"
                )
            else:
                raise AttributeError("Too many providers for segmented control")
        except (AttributeError, TypeError):
            # Fallback to selectbox if segmented_control is not available
            selected_provider = st.sidebar.selectbox(
                "Provider",
                available_providers,
                index=provider_index,
                key="provider_select",
                on_change=lambda: update_provider_selection(
                    st.session_state.provider_select
                ),
            )

        if selected_provider and selected_provider != st.session_state.config.get("vendor"):
            update_provider_selection(selected_provider)
            st.sidebar.toast(f"Switched to {selected_provider}", icon="🔄")

        available_models = providers.get(selected_provider, [])
        if not available_models:
            st.sidebar.warning(f"No models available for {selected_provider}")
            return

        current_model = st.session_state.config.get("model")
        try:
            model_index = (
                available_models.index(current_model)
                if current_model in available_models
                else 0
            )
        except ValueError:
            model_index = 0
            logger.warning(
                f"Current model {current_model} not found in available models for {selected_provider}"
            )

        model_key = f"model_select_{selected_provider}"
        selected_model = st.sidebar.selectbox(
            "Model",
            available_models,
            index=model_index,
            key=model_key,
            help=f"Available models for {selected_provider}"
        )

        if selected_model != st.session_state.config.get("model"):
            logger.debug(f"Updating model selection to: {selected_model}")
            st.session_state.config["model"] = selected_model
            st.session_state.selected_model = selected_model
            st.sidebar.toast(f"Model updated to {selected_model}", icon="⚙️")

        # Add model info expander
        with st.sidebar.expander("ℹ️ Model Info", expanded=False):
            st.write(f"**Provider:** {selected_provider}")
            st.write(f"**Model:** {selected_model}")
            st.write(f"**Context Length:** {st.session_state.config.get('context_length', 'Unknown')}")

    except Exception as e:
        logger.error(f"Error loading models and providers: {str(e)}", exc_info=True)
        st.sidebar.error(f"Error loading models and providers: {str(e)}")
        st.session_state.selected_model = None
        st.session_state.config["model"] = None


def load_pattern_descriptions():
    """Load pattern descriptions and tags from JSON file with fallback."""
    try:
        # Try to load from the pattern_descriptions directory
        descriptions_path = os.path.join(
            os.path.dirname(os.path.dirname(__file__)),
            "pattern_descriptions",
            "pattern_descriptions.json"
        )

        if os.path.exists(descriptions_path):
            with open(descriptions_path, "r", encoding="utf-8") as f:
                data = json.load(f)
                logger.info(f"Loaded {len(data.get('patterns', []))} pattern descriptions")
                return data.get("patterns", [])
        else:
            logger.warning(f"Pattern descriptions file not found at {descriptions_path}")
            return []
    except Exception as e:
        logger.error(f"Error loading pattern descriptions: {str(e)}")
        return []


def get_pattern_metadata(pattern_name):
    """Get pattern metadata from system.md."""
    pattern_path = os.path.join(pattern_dir, pattern_name, "system.md")
    if os.path.exists(pattern_path):
        with open(pattern_path, "r") as f:
            return f.read()
    return None


def detect_pattern_variables(pattern_name: str) -> List[str]:
    """Detect variables like {{variable_name}} in pattern content.
    
    Args:
        pattern_name: Name of the pattern to analyze
        
    Returns:
        List of variable names found in the pattern
    """
    logger.debug(f"Detecting variables in pattern: {pattern_name}")
    
    pattern_content = get_pattern_metadata(pattern_name)
    if not pattern_content:
        return []
    
    # Find all variables in {{variable_name}} format
    variable_pattern = r'\{\{([a-zA-Z_][a-zA-Z0-9_]*)\}\}'
    matches = re.findall(variable_pattern, pattern_content)
    
    # Remove duplicates and sort
    variables = sorted(list(set(matches)))
    
    logger.debug(f"Found variables in {pattern_name}: {variables}")
    return variables


def render_pattern_variables_ui(variables: List[str], key_prefix: str = "vars") -> Dict[str, str]:
    """Render input fields for pattern variables.
    
    Args:
        variables: List of variable names to render inputs for
        key_prefix: Prefix for streamlit keys to avoid conflicts
        
    Returns:
        Dictionary mapping variable names to their values
    """
    if not variables:
        return {}
    
    st.subheader("📝 Pattern Variables")
    st.info("This pattern requires the following variables to be set:")
    
    variable_values = {}
    
    for var in variables:
        # Create a more user-friendly label
        label = var.replace('_', ' ').title()
        
        # Provide helpful placeholders based on common variable names
        placeholder = ""
        if "name" in var.lower():
            placeholder = "Enter a name..."
        elif "author" in var.lower():
            placeholder = "e.g., Paul Graham, Maya Angelou, etc."
        elif "topic" in var.lower():
            placeholder = "Enter the topic..."
        elif "style" in var.lower():
            placeholder = "Enter the style..."
        else:
            placeholder = f"Enter value for {var}..."
        
        value = st.text_input(
            f"**{label}** (`{var}`)",
            placeholder=placeholder,
            key=f"{key_prefix}_{var}",
            help=f"This value will replace {{{{{{var}}}}}} in the pattern"
        )
        
        variable_values[var] = value
    
    return variable_values


def substitute_pattern_variables(content: str, variables: Dict[str, str]) -> str:
    """Substitute variables in pattern content.
    
    Args:
        content: The pattern content with variables
        variables: Dictionary mapping variable names to values
        
    Returns:
        Content with variables substituted
    """
    if not variables:
        return content
    
    result = content
    for var_name, var_value in variables.items():
        if var_value:  # Only substitute if value is provided
            pattern = f"{{{{{var_name}}}}}"
            result = result.replace(pattern, var_value)
            logger.debug(f"Substituted {pattern} with '{var_value}'")
    
    return result


def validate_pattern_variables(variables: Dict[str, str], required_vars: List[str]) -> Tuple[bool, List[str]]:
    """Validate that all required variables have values.
    
    Args:
        variables: Dictionary of variable values
        required_vars: List of required variable names
        
    Returns:
        Tuple of (is_valid, list_of_missing_variables)
    """
    missing_vars = []
    
    for var in required_vars:
        if var not in variables or not variables[var].strip():
            missing_vars.append(var)
    
    return len(missing_vars) == 0, missing_vars


def get_pattern_description_and_tags(pattern_name, descriptions_data=None):
    """Get pattern description and tags from the descriptions JSON."""
    if descriptions_data is None:
        descriptions_data = load_pattern_descriptions()

    for pattern in descriptions_data:
        if pattern.get("patternName") == pattern_name:
            return {
                "description": pattern.get("description", "No description available"),
                "tags": pattern.get("tags", [])
            }

    # Fallback: try to extract from system.md
    metadata = get_pattern_metadata(pattern_name)
    if metadata:
        # Extract first line or paragraph as description
        lines = metadata.split('\n')
        for line in lines:
            line = line.strip()
            if line and not line.startswith('#') and len(line) > 20:
                return {
                    "description": line[:200] + "..." if len(line) > 200 else line,
                    "tags": ["UNCATEGORIZED"]
                }

    return {
        "description": "No description available",
        "tags": []
    }


def get_all_tags(descriptions_data=None):
    """Get all unique tags from pattern descriptions."""
    if descriptions_data is None:
        descriptions_data = load_pattern_descriptions()

    all_tags = set()
    for pattern in descriptions_data:
        tags = pattern.get("tags", [])
        all_tags.update(tags)

    return sorted(list(all_tags))


def filter_patterns_by_tags(patterns, selected_tags, descriptions_data=None):
    """Filter patterns by selected tags."""
    if not selected_tags:
        return patterns

    if descriptions_data is None:
        descriptions_data = load_pattern_descriptions()

    # Create a mapping of pattern names to their tags
    pattern_tags = {}
    for pattern_data in descriptions_data:
        pattern_name = pattern_data.get("patternName")
        tags = pattern_data.get("tags", [])
        pattern_tags[pattern_name] = tags

    filtered_patterns = []
    for pattern in patterns:
        pattern_tags_list = pattern_tags.get(pattern, [])
        # Check if any of the selected tags match the pattern's tags
        if any(tag in pattern_tags_list for tag in selected_tags):
            filtered_patterns.append(pattern)

    return filtered_patterns


def search_patterns_by_description(patterns, search_term, descriptions_data=None):
    """Search patterns by description content."""
    if not search_term:
        return patterns

    if descriptions_data is None:
        descriptions_data = load_pattern_descriptions()

    search_term = search_term.lower()

    # Create a mapping of pattern names to their descriptions
    pattern_descriptions = {}
    for pattern_data in descriptions_data:
        pattern_name = pattern_data.get("patternName")
        description = pattern_data.get("description", "")
        pattern_descriptions[pattern_name] = description.lower()

    filtered_patterns = []
    for pattern in patterns:
        # Search in pattern name
        if search_term in pattern.lower():
            filtered_patterns.append(pattern)
            continue

        # Search in description
        description = pattern_descriptions.get(pattern, "")
        if search_term in description:
            filtered_patterns.append(pattern)
            continue

        # Fallback: search in system.md content
        metadata = get_pattern_metadata(pattern)
        if metadata and search_term in metadata.lower():
            filtered_patterns.append(pattern)

    return filtered_patterns


def run_pattern_extraction():
    """Run the pattern extraction script to update descriptions."""
    try:
        script_path = os.path.join(
            os.path.dirname(os.path.dirname(__file__)),
            "pattern_descriptions",
            "extract_patterns.py"
        )

        if os.path.exists(script_path):
            import subprocess
            result = subprocess.run(
                [sys.executable, script_path],
                capture_output=True,
                text=True,
                cwd=os.path.dirname(script_path)
            )

            if result.returncode == 0:
                return True, result.stdout
            else:
                return False, result.stderr
        else:
            return False, f"Extract patterns script not found at {script_path}"
    except Exception as e:
        return False, str(e)


def show_pattern_management_ui():
    """Show pattern descriptions and tags management interface."""
    st.header("🏷️ Pattern Descriptions & Tags Management")

    # Check if pattern descriptions system is available
    descriptions_data = load_pattern_descriptions()

    if not descriptions_data:
        st.warning("⚠️ Pattern descriptions system not found or empty.")
        st.info("This feature requires the pattern descriptions JSON files. You can:")

        col1, col2 = st.columns(2)
        with col1:
            if st.button("🔄 Run Pattern Extraction", type="primary"):
                with st.spinner("Running pattern extraction..."):
                    success, output = run_pattern_extraction()
                    if success:
                        st.success("✅ Pattern extraction completed!")
                        st.code(output)
                        st.rerun()
                    else:
                        st.error(f"❌ Pattern extraction failed: {output}")

        with col2:
            st.markdown("""
            **Manual Setup:**
            1. Navigate to `scripts/pattern_descriptions/`
            2. Run `python extract_patterns.py`
            3. Refresh this page
            """)
        return

    # Show statistics
    st.subheader("📊 Pattern Statistics")
    col1, col2, col3, col4 = st.columns(4)

    total_patterns = len(descriptions_data)
    pending_descriptions = len([p for p in descriptions_data if p.get("description") == "[Description pending]"])
    tagged_patterns = len([p for p in descriptions_data if p.get("tags")])
    all_tags = get_all_tags(descriptions_data)

    with col1:
        st.metric("Total Patterns", total_patterns)
    with col2:
        st.metric("Pending Descriptions", pending_descriptions)
    with col3:
        st.metric("Tagged Patterns", tagged_patterns)
    with col4:
        st.metric("Unique Tags", len(all_tags))

    # Management tabs
    tab1, tab2, tab3 = st.tabs(["📝 Edit Descriptions", "🏷️ Manage Tags", "🔄 Sync Patterns"])

    with tab1:
        st.subheader("Edit Pattern Descriptions")

        # Filter for pending descriptions
        if st.checkbox("Show only pending descriptions"):
            patterns_to_show = [p for p in descriptions_data if p.get("description") == "[Description pending]"]
        else:
            patterns_to_show = descriptions_data

        if patterns_to_show:
            for i, pattern_data in enumerate(patterns_to_show[:10]):  # Show first 10
                pattern_name = pattern_data.get("patternName", "")
                current_description = pattern_data.get("description", "")

                with st.expander(f"📝 {pattern_name}", expanded=current_description == "[Description pending]"):
                    # Show current info
                    col1, col2 = st.columns([2, 1])
                    with col1:
                        st.markdown(f"**Pattern:** {pattern_name}")
                        st.markdown(f"**Current Description:** {current_description}")
                    with col2:
                        current_tags = pattern_data.get("tags", [])
                        if current_tags:
                            tag_display = " ".join([f"`{tag}`" for tag in current_tags])
                            st.markdown(f"**Tags:** {tag_display}")

                    # Show pattern preview
                    metadata = get_pattern_metadata(pattern_name)
                    if metadata:
                        with st.expander("📖 Pattern Content Preview", expanded=False):
                            preview = metadata[:1000] + "..." if len(metadata) > 1000 else metadata
                            st.code(preview, language="markdown")

                    st.info("💡 Use this preview to write a concise description of what this pattern does.")
        else:
            st.success("🎉 All patterns have descriptions!")

    with tab2:
        st.subheader("Tag Management")

        if all_tags:
            st.markdown("**Current Tags:**")
            tag_display = " ".join([f"`{tag}`" for tag in all_tags])
            st.markdown(tag_display)

            # Tag usage statistics
            tag_usage = {}
            for pattern_data in descriptions_data:
                for tag in pattern_data.get("tags", []):
                    tag_usage[tag] = tag_usage.get(tag, 0) + 1

            if tag_usage:
                st.markdown("**Tag Usage:**")
                for tag, count in sorted(tag_usage.items(), key=lambda x: x[1], reverse=True):
                    st.write(f"• `{tag}`: {count} patterns")
        else:
            st.info("No tags found. Add tags to patterns to see them here.")

    with tab3:
        st.subheader("Sync Pattern Data")
        st.info("Use this to update pattern descriptions when new patterns are added.")

        if st.button("🔄 Run Pattern Extraction", type="primary"):
            with st.spinner("Extracting pattern information..."):
                success, output = run_pattern_extraction()
                if success:
                    st.success("✅ Pattern extraction completed!")
                    with st.expander("📋 Extraction Output", expanded=False):
                        st.code(output)
                    st.rerun()
                else:
                    st.error(f"❌ Pattern extraction failed:")
                    st.code(output)


def get_patterns():
    """Get the list of available patterns from the specified directory."""
    if not os.path.exists(pattern_dir):
        st.error(f"Pattern directory not found: {pattern_dir}")
        return []
    try:
        patterns = [
            item
            for item in os.listdir(pattern_dir)
            if os.path.isdir(os.path.join(pattern_dir, item))
        ]
        return patterns
    except PermissionError:
        st.error(f"Permission error accessing pattern directory: {pattern_dir}")
        return []
    except Exception as e:
        st.error(f"An unexpected error occurred: {e}")
        return []


def enhanced_pattern_selector(patterns: List[str], key: str = "pattern_selector") -> List[str]:
    """Enhanced pattern selector with descriptions, tags, and search functionality."""
    if not patterns:
        st.warning("No patterns available. Create a pattern first.")
        return []

    # Load pattern descriptions and tags
    descriptions_data = load_pattern_descriptions()
    all_tags = get_all_tags(descriptions_data)

    selected_tags = []
    with st.expander("Filter by Category", expanded=False):
        if all_tags:
            try:
                selected_tags = st.pills(
                    "Select tags to filter patterns",
                    all_tags,
                    selection_mode="multi",
                    key=f"{key}_tag_pills",
                    label_visibility="collapsed"
                )
            except (AttributeError, TypeError):
                selected_tags = st.multiselect(
                    "Select tags to filter patterns",
                    all_tags,
                    key=f"{key}_tag_multiselect",
                    label_visibility="collapsed"
                )

    # Define search_term as empty since it's removed
    search_term = ""

    # Apply filters
    filtered_patterns = patterns

    # Filter by tags
    if selected_tags:
        filtered_patterns = filter_patterns_by_tags(
            filtered_patterns, selected_tags, descriptions_data
        )

    if not filtered_patterns:
        st.info("No patterns match your filters. Try adjusting your search or tags.")
        return []

    # Show filter results summary
    if selected_tags:
        st.caption(f"📊 Showing {len(filtered_patterns)} patterns matching tags: {', '.join(selected_tags)}")

    # Enhanced pattern display with variable indicators
    pattern_display_options = []
    for pattern in filtered_patterns:
        variables = detect_pattern_variables(pattern)
        if variables:
            # Add variable indicator to pattern name
            display_name = f"{pattern} 🔧 ({len(variables)} vars)"
        else:
            display_name = pattern
        pattern_display_options.append(display_name)

    # Use pills for pattern selection if available and not too many patterns
    try:
        if len(filtered_patterns) <= 10:
            selected_display_patterns = st.pills(
                "Select Patterns",
                pattern_display_options,
                selection_mode="multi",
                key=f"{key}_pills"
            )
        else:
            raise AttributeError("Too many patterns for pills")
    except (AttributeError, TypeError):
        # Fall back to multiselect if pills is not available
        selected_display_patterns = st.multiselect(
            "Select Patterns",
            pattern_display_options,
            key=f"{key}_multiselect"
        )

    # Convert display names back to actual pattern names
    selected_patterns = []
    if selected_display_patterns:
        for display_pattern in selected_display_patterns:
            # Extract the actual pattern name (before any indicators)
            actual_pattern = display_pattern.split(' 🔧')[0]
            selected_patterns.append(actual_pattern)

    return selected_patterns if selected_patterns else []


def create_execution_status_container():
    """Create an enhanced status container for pattern execution."""
    return st.status("🚀 Preparing execution...", expanded=True)


def show_pattern_feedback_ui(pattern_name: str, output: str):
    """Show feedback UI for pattern execution results."""
    col1, col2, col3 = st.columns([2, 1, 1])

    with col1:
        st.markdown(f"**{pattern_name}** execution completed")

    with col2:
        try:
            feedback = st.feedback(
                "thumbs",
                key=f"feedback_{pattern_name}_{hash(output)}"
            )
            if feedback is not None:
                st.session_state.pattern_feedback[pattern_name] = feedback
                if feedback == 1:
                    try:
                        st.toast("Thanks for the positive feedback! 👍", icon="✅")
                    except (AttributeError, TypeError):
                        st.success("Thanks for the positive feedback! 👍")
                else:
                    try:
                        st.toast("Thanks for the feedback. We'll improve! 👎", icon="📝")
                    except (AttributeError, TypeError):
                        st.info("Thanks for the feedback. We'll improve! 👎")
        except (AttributeError, TypeError):
            # Fallback to simple buttons if feedback widget is not available
            col_like, col_dislike = st.columns(2)
            with col_like:
                if st.button("👍", key=f"like_{pattern_name}_{hash(output)}"):
                    st.session_state.pattern_feedback[pattern_name] = 1
                    st.success("Thanks!")
            with col_dislike:
                if st.button("👎", key=f"dislike_{pattern_name}_{hash(output)}"):
                    st.session_state.pattern_feedback[pattern_name] = 0
                    st.info("Thanks for feedback!")

    with col3:
        if st.button("⭐ Star Output", key=f"star_{pattern_name}_{hash(output)}"):
            # Add to starred outputs
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            star_output_enhanced(pattern_name, output, timestamp)


def star_output_enhanced(pattern_name: str, output: str, timestamp: str):
    """Enhanced starring functionality with dialog."""
    @st.dialog("⭐ Star this output")
    def star_dialog():
        st.write(f"**Pattern:** {pattern_name}")
        st.write(f"**Timestamp:** {timestamp}")

        custom_name = st.text_input(
            "Give this output a name (optional):",
            placeholder=f"Great {pattern_name} result"
        )

        col1, col2 = st.columns(2)
        with col1:
            if st.button("⭐ Star it!", type="primary", use_container_width=True):
                log_entry = {
                    "timestamp": timestamp,
                    "pattern_name": pattern_name,
                    "input": st.session_state.get("input_content", ""),
                    "output": output,
                    "is_starred": True,
                    "custom_name": custom_name or f"Starred {pattern_name} output",
                }
                st.session_state.starred_outputs.append(log_entry)
                save_outputs()
                st.success("Output starred successfully!")
                st.rerun()

        with col2:
            if st.button("Cancel", use_container_width=True):
                st.rerun()

    star_dialog()


def create_pattern(
    pattern_name: str, content: Optional[str] = None
) -> Tuple[bool, str]:
    """Create a new pattern with necessary files and structure."""
    new_pattern_path = None
    try:
        # Validate pattern name
        if not pattern_name:
            logger.error("Pattern name cannot be empty")
            return False, "Pattern name cannot be empty."

        # Check if pattern already exists
        new_pattern_path = os.path.join(pattern_dir, pattern_name)
        if os.path.exists(new_pattern_path):
            logger.error(f"Pattern {pattern_name} already exists")
            return False, "Pattern already exists."

        # Create pattern directory
        os.makedirs(new_pattern_path)
        logger.info(f"Created pattern directory: {new_pattern_path}")

        # If content is provided, use fabric create_pattern to structure it
        if content:
            logger.info(
                f"Structuring content for pattern '{pattern_name}' using Fabric"
            )
            try:
                # Get current model and provider configuration
                current_provider = st.session_state.config.get("vendor")
                current_model = st.session_state.config.get("model")

                if not current_provider or not current_model:
                    raise ValueError("Please select a provider and model first.")

                # Execute fabric create_pattern with input content
                cmd = ["fabric", "--pattern", "create_pattern"]
                if current_provider and current_model:
                    cmd.extend(["--vendor", current_provider, "--model", current_model])

                logger.debug(f"Running command: {' '.join(cmd)}")
                logger.debug(f"Input content:\n{content}")

                # Execute pattern
                result = run(
                    cmd, input=content, capture_output=True, text=True, check=True
                )
                structured_content = result.stdout.strip()

                if not structured_content:
                    raise ValueError("No output received from create_pattern")

                # Save the structured content to system.md
                system_file = os.path.join(new_pattern_path, "system.md")
                with open(system_file, "w") as f:
                    f.write(structured_content)

                # Validate the created pattern
                is_valid, validation_message = validate_pattern(pattern_name)
                if not is_valid:
                    raise ValueError(f"Pattern validation failed: {validation_message}")

                logger.info(
                    f"Successfully created pattern '{pattern_name}' with structured content"
                )

            except CalledProcessError as e:
                error_msg = f"Error running create_pattern: {e.stderr}"
                logger.error(error_msg)
                if os.path.exists(new_pattern_path):
                    shutil.rmtree(new_pattern_path)
                return False, error_msg

            except Exception as e:
                error_msg = f"Unexpected error during content structuring: {str(e)}"
                logger.error(error_msg)
                if os.path.exists(new_pattern_path):
                    shutil.rmtree(new_pattern_path)
                return False, error_msg
        else:
            # Create minimal template for manual editing
            logger.info(f"Creating minimal template for pattern '{pattern_name}'")
            system_file = os.path.join(new_pattern_path, "system.md")
            with open(system_file, "w") as f:
                f.write("# IDENTITY and PURPOSE\n\n# STEPS\n\n# OUTPUT INSTRUCTIONS\n")

            # Validate the created pattern
            is_valid, validation_message = validate_pattern(pattern_name)
            if not is_valid:
                logger.warning(
                    f"Pattern created but validation failed: {validation_message}"
                )

        return True, f"Pattern '{pattern_name}' created successfully."

    except Exception as e:
        error_msg = f"Error creating pattern: {str(e)}"
        logger.error(error_msg)
        # Clean up on any error
        if new_pattern_path and os.path.exists(new_pattern_path):
            shutil.rmtree(new_pattern_path)
        return False, error_msg


def delete_pattern(pattern_name):
    """Delete an existing pattern."""
    try:
        if not pattern_name:
            return False, "Pattern name cannot be empty."

        pattern_path = os.path.join(pattern_dir, pattern_name)
        if not os.path.exists(pattern_path):
            return False, "Pattern does not exist."

        shutil.rmtree(pattern_path)
        return True, f"Pattern '{pattern_name}' deleted successfully."
    except Exception as e:
        return False, f"Error deleting pattern: {str(e)}"


def pattern_creation_wizard():
    """Multi-step wizard for creating a new pattern."""
    st.header("Create New Pattern")

    pattern_name = st.text_input("Pattern Name")
    if pattern_name:
        edit_mode = st.radio(
            "Edit Mode",
            ["Simple Editor", "Advanced (Wizard)"],
            key="pattern_creation_edit_mode",
            horizontal=True,
        )

        if edit_mode == "Simple Editor":
            new_content = st.text_area("Enter Pattern Content", height=400)

            if st.button("Create Pattern", type="primary"):
                success, message = create_pattern(pattern_name, new_content)
                if success:
                    st.success(message)
                    st.experimental_rerun()
                else:
                    st.error(message)

        else:
            sections = ["IDENTITY", "GOAL", "OUTPUT", "OUTPUT INSTRUCTIONS"]
            current_section = st.radio(
                "Edit Section", sections, key="pattern_creation_section_select"
            )

            if current_section == "IDENTITY":
                identity = st.text_area("Define the IDENTITY", height=200)
                st.session_state.new_pattern_identity = identity

            elif current_section == "GOAL":
                goal = st.text_area("Define the GOAL", height=200)
                st.session_state.new_pattern_goal = goal

            elif current_section == "OUTPUT":
                output = st.text_area("Define the OUTPUT", height=200)
                st.session_state.new_pattern_output = output

            elif current_section == "OUTPUT INSTRUCTIONS":
                instructions = st.text_area(
                    "Define the OUTPUT INSTRUCTIONS", height=200
                )
                st.session_state.new_pattern_instructions = instructions

            pattern_content = f"""# IDENTITY
{st.session_state.get('new_pattern_identity', '')}

# GOAL
{st.session_state.get('new_pattern_goal', '')}

# OUTPUT
{st.session_state.get('new_pattern_output', '')}

# OUTPUT INSTRUCTIONS
{st.session_state.get('new_pattern_instructions', '')}"""

            if st.button("Create Pattern", type="primary"):
                success, message = create_pattern(pattern_name, pattern_content)
                if success:
                    st.success(message)
                    for key in [
                        "new_pattern_identity",
                        "new_pattern_goal",
                        "new_pattern_output",
                        "new_pattern_instructions",
                    ]:
                        if key in st.session_state:
                            del st.session_state[key]
                    st.experimental_rerun()
                else:
                    st.error(message)
    else:
        st.info("Enter a pattern name to create a new pattern")


def bulk_edit_patterns(patterns_to_edit, field_to_update, new_value):
    """Perform bulk edits on multiple patterns."""
    results = []
    for pattern in patterns_to_edit:
        try:
            pattern_path = os.path.join(pattern_dir, pattern)
            system_file = os.path.join(pattern_path, "system.md")

            if not os.path.exists(system_file):
                results.append((pattern, False, "system.md not found"))
                continue

            with open(system_file, "r") as f:
                content = f.read()

            if field_to_update == "purpose":
                sections = content.split("#")
                updated_sections = []
                for section in sections:
                    if section.strip().startswith("IDENTITY and PURPOSE"):
                        lines = section.split("\n")
                        for i, line in enumerate(lines):
                            if "You are an AI assistant designed to" in line:
                                lines[i] = (
                                    f"You are an AI assistant designed to {new_value}."
                                )
                        updated_sections.append("\n".join(lines))
                    else:
                        updated_sections.append(section)

                new_content = "#".join(updated_sections)
                with open(system_file, "w") as f:
                    f.write(new_content)
                results.append((pattern, True, "Updated successfully"))
            else:
                results.append(
                    (
                        pattern,
                        False,
                        f"Field {field_to_update} not supported for bulk edit",
                    )
                )

        except Exception as e:
            results.append((pattern, False, str(e)))

    return results


def pattern_creation_ui():
    """UI component for creating patterns with simple and wizard modes."""
    pattern_name = st.text_input("Pattern Name")
    if not pattern_name:
        st.info("Enter a pattern name to create a new pattern")
        return

    system_content = """# IDENTITY and PURPOSE

You are an AI assistant designed to {purpose}.

# STEPS

- Step 1
- Step 2
- Step 3

# OUTPUT INSTRUCTIONS

- Output format instructions here
"""
    new_content = st.text_area("Edit Pattern Content", system_content, height=400)

    if st.button("Create Pattern", type="primary"):
        if not pattern_name:
            st.error("Pattern name cannot be empty.")
        else:
            success, message = create_pattern(pattern_name)
            if success:
                system_file = os.path.join(pattern_dir, pattern_name, "system.md")
                with open(system_file, "w") as f:
                    f.write(new_content)
                st.success(f"Pattern '{pattern_name}' created successfully!")
                st.experimental_rerun()
            else:
                st.error(message)


def pattern_management_ui():
    """UI component for pattern management."""
    st.sidebar.title("Pattern Management")


def save_output_log(
    pattern_name: str, input_content: str, output_content: str, timestamp: str
):
    """Save pattern execution log."""
    log_entry = {
        "timestamp": timestamp,
        "pattern_name": pattern_name,
        "input": input_content,
        "output": output_content,
        "is_starred": False,
        "custom_name": "",
    }
    st.session_state.output_logs.append(log_entry)
    # Save outputs after each new log entry
    save_outputs()


def star_output(log_index: int, custom_name: str = "") -> bool:
    """Star/favorite an output log.

    Args:
        log_index: Index of the output log to star
        custom_name: Optional custom name for the starred output

    Returns:
        bool: True if output was starred successfully, False otherwise
    """
    try:
        if 0 <= log_index < len(st.session_state.output_logs):
            log_entry = st.session_state.output_logs[log_index].copy()
            log_entry["is_starred"] = True
            log_entry["custom_name"] = (
                custom_name
                or f"Starred Output #{len(st.session_state.starred_outputs) + 1}"
            )

            # Check if this output is already starred (by timestamp)
            if not any(
                s["timestamp"] == log_entry["timestamp"]
                for s in st.session_state.starred_outputs
            ):
                st.session_state.starred_outputs.append(log_entry)
                save_outputs()  # Save after starring
                return True

        return False
    except Exception as e:
        logger.error(f"Error starring output: {str(e)}")
        return False


def unstar_output(log_index: int):
    """Remove an output from starred/favorites."""
    if 0 <= log_index < len(st.session_state.starred_outputs):
        st.session_state.starred_outputs.pop(log_index)
        # Save outputs after unstarring
        save_outputs()


def validate_input_content(input_text: str) -> Tuple[bool, str]:
    """Validate input content for potentially problematic characters or patterns.

    Args:
        input_text: The input text to validate

    Returns:
        Tuple[bool, str]: (is_valid, error_message)
    """
    if not input_text or input_text.isspace():
        return False, "Input content cannot be empty or only whitespace."

    # Check for minimum length
    if len(input_text.strip()) < 2:
        return False, "Input content must be at least 2 characters long."

    # Check for maximum length (e.g., 100KB)
    if len(input_text.encode("utf-8")) > 100 * 1024:
        return False, "Input content exceeds maximum size of 100KB."

    # Check for high concentration of special characters
    special_chars = set("!@#$%^&*()_+[]{}|\\;:'\",.<>?`~")
    special_char_count = sum(1 for c in input_text if c in special_chars)
    special_char_ratio = special_char_count / len(input_text)

    if special_char_ratio > 0.3:  # More than 30% special characters
        return (
            False,
            "Input contains too many special characters. Please check your input.",
        )

    # Check for control characters
    control_chars = set(
        chr(i) for i in range(32) if i not in [9, 10, 13]
    )  # Allow tab, newline, carriage return
    if any(c in control_chars for c in input_text):
        return False, "Input contains invalid control characters."

    # Check for proper UTF-8 encoding
    try:
        input_text.encode("utf-8").decode("utf-8")
    except UnicodeError:
        return False, "Input contains invalid Unicode characters."

    return True, ""


def sanitize_input_content(input_text: str) -> str:
    """Sanitize input content by removing or replacing problematic characters.

    Args:
        input_text: The input text to sanitize

    Returns:
        str: Sanitized input text
    """
    # Remove null bytes
    text = input_text.replace("\0", "")

    # Replace control characters with spaces (except newlines and tabs)
    allowed_chars = {"\n", "\t", "\r"}
    sanitized_chars = []
    for c in text:
        if c in allowed_chars or ord(c) >= 32:
            sanitized_chars.append(c)
        else:
            sanitized_chars.append(" ")

    # Join characters and normalize whitespace
    text = "".join(sanitized_chars)
    text = " ".join(text.split())

    return text


def execute_patterns_enhanced(
    patterns_to_run: List[str],
    chain_mode: bool = False,
    initial_input: Optional[str] = None,
    pattern_variables: Optional[Dict[str, Dict[str, str]]] = None,
) -> List[str]:
    """Execute patterns with enhanced UI feedback and streaming.
    
    Args:
        patterns_to_run: List of pattern names to execute
        chain_mode: Whether to chain patterns (output of one becomes input of next)
        initial_input: Initial input text
        pattern_variables: Dict mapping pattern names to their variable values
    """
    logger.info(f"Executing {len(patterns_to_run)} patterns")

    # Update execution stats
    st.session_state.execution_stats["total_runs"] += 1

    # Create enhanced status container
    status_container = st.status(
        f"🚀 Executing {len(patterns_to_run)} pattern{'s' if len(patterns_to_run) > 1 else ''}...",
        expanded=True
    )

    with status_container:
        st.write("🔍 Validating configuration...")

        # Validate configuration
        current_provider = st.session_state.config.get("vendor")
        current_model = st.session_state.config.get("model")

        if not current_provider or not current_model:
            error_msg = "Please select a provider and model first."
            logger.error(error_msg)
            st.error(error_msg)
            st.session_state.execution_stats["failed_runs"] += 1
            return []

        st.write("✅ Configuration validated")
        st.write("🔍 Validating input content...")

        current_input = initial_input or st.session_state.input_content

        # Validate input content
        is_valid, error_message = validate_input_content(current_input)
        if not is_valid:
            logger.error(f"Input validation failed: {error_message}")
            st.error(f"Input validation failed: {error_message}")
            st.session_state.execution_stats["failed_runs"] += 1
            return []

        st.write("✅ Input validated")

        # Validate pattern variables
        if pattern_variables:
            st.write("🔍 Validating pattern variables...")
            for pattern_name, variables in pattern_variables.items():
                required_vars = detect_pattern_variables(pattern_name)
                is_valid, missing_vars = validate_pattern_variables(variables, required_vars)
                if not is_valid:
                    error_msg = f"Pattern '{pattern_name}' is missing required variables: {', '.join(missing_vars)}"
                    logger.error(error_msg)
                    st.error(error_msg)
                    st.session_state.execution_stats["failed_runs"] += 1
                    return []
            st.write("✅ Pattern variables validated")

        # Sanitize input content
        try:
            sanitized_input = sanitize_input_content(current_input)
            if sanitized_input != current_input:
                logger.info("Input content was sanitized")
                st.warning("Input content was automatically sanitized")
            current_input = sanitized_input
        except Exception as e:
            logger.error(f"Error sanitizing input: {str(e)}")
            st.error(f"Error processing input: {str(e)}")
            st.session_state.execution_stats["failed_runs"] += 1
            return []

        st.write(f"🤖 Using: {current_provider} - {current_model}")

        all_outputs = []
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        start_time = time.time()

        try:
            for i, pattern in enumerate(patterns_to_run, 1):
                st.write(f"🔄 Running pattern {i}/{len(patterns_to_run)}: **{pattern}**")

                # Create progress bar
                progress = st.progress(0)
                progress.progress(i / len(patterns_to_run))

                logger.info(f"Running pattern: {pattern}")
                try:
                    # Handle pattern variables
                    cmd = ["fabric", "--pattern", pattern]
                    
                    # Add variables if present for this pattern
                    if pattern_variables and pattern in pattern_variables:
                        variables = pattern_variables[pattern]
                        for var_name, var_value in variables.items():
                            if var_value:  # Only add non-empty variables
                                cmd.extend(["--variable", f"{var_name}={var_value}"])
                                logger.debug(f"Added variable {var_name}={var_value} for pattern {pattern}")
                    
                    message = current_input if chain_mode else st.session_state.input_content
                    input_data = str(message)

                    # Run the command
                    result = run(
                        cmd, input=input_data, capture_output=True, text=True, check=True
                    )

                    pattern_output = result.stdout.strip()

                    if pattern_output:
                        st.write(f"✅ Pattern **{pattern}** completed successfully")

                        # Save to output logs
                        save_output_log(pattern, message, pattern_output, timestamp)

                        # Add to outputs with enhanced formatting
                        output_msg = f"""### 🎯 {pattern}

{pattern_output}"""
                        all_outputs.append(output_msg)

                        if chain_mode:
                            current_input = pattern_output
                    else:
                        st.warning(f"Pattern **{pattern}** generated no output")
                        all_outputs.append(f"### ⚠️ {pattern}\n\nNo output generated.")

                except Exception as e:
                    error_msg = f"❌ Pattern **{pattern}** failed: {str(e)}"
                    st.error(error_msg)
                    logger.error(f"Pattern {pattern} failed: {str(e)}")
                    all_outputs.append(f"### ❌ {pattern}\n\n{str(e)}")
                    if chain_mode:
                        break

            # Calculate execution time
            execution_time = time.time() - start_time
            st.session_state.execution_stats["avg_execution_time"] = (
                (st.session_state.execution_stats["avg_execution_time"] *
                 (st.session_state.execution_stats["total_runs"] - 1) + execution_time) /
                st.session_state.execution_stats["total_runs"]
            )

            if all_outputs:
                st.session_state.execution_stats["successful_runs"] += 1
                st.write(f"🎉 Execution completed in {execution_time:.2f}s")
            else:
                st.session_state.execution_stats["failed_runs"] += 1

        except Exception as e:
            error_msg = f"❌ Error in pattern execution: {str(e)}"
            logger.error(error_msg, exc_info=True)
            st.error(error_msg)
            st.session_state.execution_stats["failed_runs"] += 1

    logger.info("Pattern execution completed")
    return all_outputs


def execute_patterns(
    patterns_to_run: List[str],
    chain_mode: bool = False,
    initial_input: Optional[str] = None,
    pattern_variables: Optional[Dict[str, Dict[str, str]]] = None,
) -> List[str]:
    """Legacy wrapper for backward compatibility."""
    return execute_patterns_enhanced(patterns_to_run, chain_mode, initial_input, pattern_variables)


def validate_pattern(pattern_name):
    """Validate a pattern's structure and content."""
    try:
        pattern_path = os.path.join(pattern_dir, pattern_name)

        if not os.path.exists(os.path.join(pattern_path, "system.md")):
            return False, f"Missing required file: system.md."

        with open(os.path.join(pattern_path, "system.md")) as f:
            content = f.read()
            required_sections = ["# IDENTITY", "# STEPS", "# OUTPUT"]
            missing_sections = []
            for section in required_sections:
                if section.lower() not in content.lower():
                    missing_sections.append(section)

            if missing_sections:
                return (
                    True,
                    f"Warning: Missing sections in system.md: {', '.join(missing_sections)}",
                )

        return True, "Pattern is valid."
    except Exception as e:
        return False, f"Error validating pattern: {str(e)}"


def pattern_editor(pattern_name):
    """Edit pattern content with simple and advanced editing options."""
    if not pattern_name:
        return

    pattern_path = os.path.join(pattern_dir, pattern_name)
    system_file = os.path.join(pattern_path, "system.md")
    user_file = os.path.join(pattern_path, "user.md")

    st.markdown(f"### Editing Pattern: {pattern_name}")
    is_valid, message = validate_pattern(pattern_name)
    if not is_valid:
        st.error(message)
    elif message != "Pattern is valid.":
        st.warning(message)
    else:
        st.success("Pattern structure is valid")

    edit_mode = st.radio(
        "Edit Mode",
        ["Simple Editor", "Advanced (Wizard)"],
        key=f"edit_mode_{pattern_name}",
        horizontal=True,
    )

    if edit_mode == "Simple Editor":
        if os.path.exists(system_file):
            with open(system_file) as f:
                content = f.read()
            new_content = st.text_area("Edit system.md", content, height=600)
            if st.button("Save system.md"):
                with open(system_file, "w") as f:
                    f.write(new_content)
                st.success("Saved successfully!")
        else:
            st.error("system.md file not found")

        if os.path.exists(user_file):
            with open(user_file) as f:
                content = f.read()
            new_content = st.text_area("Edit user.md", content, height=300)
            if st.button("Save user.md"):
                with open(user_file, "w") as f:
                    f.write(new_content)
                st.success("Saved successfully!")

    else:
        if os.path.exists(system_file):
            with open(system_file) as f:
                content = f.read()

            sections = content.split("#")
            edited_sections = []

            for section in sections:
                if not section.strip():
                    continue

                lines = section.strip().split("\n", 1)
                if len(lines) > 1:
                    title, content = lines
                else:
                    title, content = lines[0], ""

                st.markdown(f"#### {title}")
                new_content = st.text_area(
                    f"Edit {title} section",
                    value=content.strip(),
                    height=200,
                    key=f"section_{title}",
                )
                edited_sections.append(f"# {title}\n\n{new_content}")

            if st.button("Save Changes"):
                new_content = "\n\n".join(edited_sections)
                with open(system_file, "w") as f:
                    f.write(new_content)
                st.success("Changes saved successfully!")

                is_valid, message = validate_pattern(pattern_name)
                if not is_valid:
                    st.error(message)
                elif message != "Pattern is valid.":
                    st.warning(message)
        else:
            st.error("system.md file not found")


def get_outputs_dir() -> str:
    """Get the directory for storing outputs."""
    outputs_dir = os.path.expanduser("~/.config/fabric/outputs")
    os.makedirs(outputs_dir, exist_ok=True)
    return outputs_dir


def save_outputs():
    """Save pattern outputs and starred outputs to files.

    Error handling:
    - Creates output directory if it doesn't exist
    - Handles file write permissions
    - Handles JSON serialization errors
    - Logs all errors for debugging
    """
    logger.info("Saving outputs to persistent storage")
    outputs_dir = get_outputs_dir()

    output_logs_file = os.path.join(outputs_dir, "output_logs.json")
    starred_outputs_file = os.path.join(outputs_dir, "starred_outputs.json")

    try:
        # Save output logs
        with open(output_logs_file, "w") as f:
            json.dump(st.session_state.output_logs, f, indent=2)
        logger.debug(f"Saved output logs to {output_logs_file}")

        # Save starred outputs
        with open(starred_outputs_file, "w") as f:
            json.dump(st.session_state.starred_outputs, f, indent=2)
        logger.debug(f"Saved starred outputs to {starred_outputs_file}")

    except PermissionError as e:
        error_msg = f"Permission denied when saving outputs: {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)
    except json.JSONEncodeError as e:
        error_msg = f"Error encoding outputs to JSON: {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)
    except Exception as e:
        error_msg = f"Unexpected error saving outputs: {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)


def load_saved_outputs():
    """Load saved pattern outputs from files.

    Error handling:
    - Handles missing output files
    - Handles corrupted JSON files
    - Handles file read permissions
    - Initializes empty state if files don't exist
    """
    logger.info("Loading saved outputs")
    outputs_dir = get_outputs_dir()
    output_logs_file = os.path.join(outputs_dir, "output_logs.json")
    starred_outputs_file = os.path.join(outputs_dir, "starred_outputs.json")

    try:
        # Load output logs
        if os.path.exists(output_logs_file):
            with open(output_logs_file, "r") as f:
                st.session_state.output_logs = json.load(f)
            logger.debug(f"Loaded output logs from {output_logs_file}")

        # Load starred outputs
        if os.path.exists(starred_outputs_file):
            with open(starred_outputs_file, "r") as f:
                st.session_state.starred_outputs = json.load(f)
            logger.debug(f"Loaded starred outputs from {starred_outputs_file}")

    except json.JSONDecodeError as e:
        error_msg = f"Error decoding saved outputs (corrupted files): {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)
        # Initialize empty state
        st.session_state.output_logs = []
        st.session_state.starred_outputs = []
    except PermissionError as e:
        error_msg = f"Permission denied when loading outputs: {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)
    except Exception as e:
        error_msg = f"Unexpected error loading saved outputs: {str(e)}"
        logger.error(error_msg)
        st.error(error_msg)
        # Initialize empty state
        st.session_state.output_logs = []
        st.session_state.starred_outputs = []


def handle_star_name_input(log_index: int, name: str):
    """Handle the starring process when a name is input.

    Args:
        log_index: Index of the output to star
        name: Name to give the starred output
    """
    try:
        if star_output(log_index, name):
            st.success("Output starred successfully!")
        else:
            st.error("Failed to star output. Please try again.")
    except Exception as e:
        logger.error(f"Error handling star name input: {str(e)}")
        st.error(f"Error starring output: {str(e)}")


def execute_pattern_chain(patterns_sequence: List[str], initial_input: str, pattern_variables: Optional[Dict[str, Dict[str, str]]] = None) -> Dict:
    """Execute a sequence of patterns in a chain, passing output from each to the next.

    Args:
        patterns_sequence: List of pattern names to execute in sequence
        initial_input: Initial input text to start the chain
        pattern_variables: Dict mapping pattern names to their variable values

    Returns:
        Dict containing results from each stage of the chain
    """
    logger.info(
        f"Starting pattern chain execution with {len(patterns_sequence)} patterns"
    )
    chain_results = {
        "sequence": patterns_sequence,
        "stages": [],
        "final_output": None,
        "metadata": {
            "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            "success": False,
        },
    }

    current_input = initial_input

    try:
        for i, pattern in enumerate(patterns_sequence, 1):
            logger.info(f"Chain Stage {i}: Executing pattern '{pattern}'")
            stage_result = {
                "pattern": pattern,
                "input": current_input,
                "output": None,
                "success": False,
                "error": None,
            }

            try:
                cmd = ["fabric", "--pattern", pattern]
                
                # Add variables if present for this pattern
                if pattern_variables and pattern in pattern_variables:
                    variables = pattern_variables[pattern]
                    for var_name, var_value in variables.items():
                        if var_value:  # Only add non-empty variables
                            cmd.extend(["--variable", f"{var_name}={var_value}"])
                            logger.debug(f"Added variable {var_name}={var_value} for pattern {pattern} in chain")
                
                result = run(
                    cmd, input=current_input, capture_output=True, text=True, check=True
                )
                output = result.stdout.strip()

                if output:
                    stage_result["output"] = output
                    stage_result["success"] = True
                    current_input = output  # Use this output as input for next pattern
                    logger.debug(f"Stage {i} completed successfully")
                else:
                    stage_result["error"] = "Pattern generated no output"
                    logger.warning(f"Pattern {pattern} generated no output")

            except CalledProcessError as e:
                error_msg = f"Error executing pattern: {e.stderr.strip()}"
                stage_result["error"] = error_msg
                logger.error(error_msg)
                break

            except Exception as e:
                error_msg = f"Unexpected error: {str(e)}"
                stage_result["error"] = error_msg
                logger.error(error_msg)
                break

            chain_results["stages"].append(stage_result)

            # Save stage output to logs
            save_output_log(
                pattern,
                stage_result["input"],
                stage_result["output"] or stage_result["error"],
                chain_results["metadata"]["timestamp"],
            )

        # Set final output and success status
        successful_stages = [s for s in chain_results["stages"] if s["success"]]
        if successful_stages:
            chain_results["final_output"] = successful_stages[-1]["output"]
            chain_results["metadata"]["success"] = True

    except Exception as e:
        logger.error(f"Chain execution failed: {str(e)}", exc_info=True)
        chain_results["metadata"]["error"] = str(e)

    return chain_results


def enhance_input_preview():
    """Display a preview of the input content with basic statistics.

    Shows:
    - Input text preview
    - Character count
    - Word count
    """
    if "input_content" in st.session_state and st.session_state.input_content:
        with st.expander("Input Preview", expanded=True):
            st.markdown("### Current Input")
            st.code(st.session_state.input_content, language="text")

            # Basic statistics
            char_count = len(st.session_state.input_content)
            word_count = len(st.session_state.input_content.split())

            col1, col2 = st.columns(2)
            with col1:
                st.metric("Characters", char_count)
            with col2:
                st.metric("Words", word_count)


def get_clipboard_content() -> Tuple[bool, str, str]:
    """Get content from clipboard with proper error handling.

    Returns:
        Tuple[bool, str, str]: (success, content, error_message)
    """
    try:
        # macOS - use pbpaste
        if PLATFORM == "darwin":
            result = run(["pbpaste"], capture_output=True, text=True, check=True)
        # Windows - fallback to pyperclip if available
        elif PLATFORM == "win32":
            try:
                import pyperclip

                content = pyperclip.paste()
                return True, content, ""
            except ImportError:
                return (
                    False,
                    "",
                    "The pyperclip module is required for clipboard operations on Windows.\nPlease install it with: pip install pyperclip",
                )
            except Exception as e:
                return False, "", f"Windows clipboard error: {str(e)}"
        # Linux - use xclip
        else:
            result = run(
                ["xclip", "-selection", "clipboard", "-o"],
                capture_output=True,
                text=True,
                check=True,
            )

        content = result.stdout
        # Validate the content is proper UTF-8
        try:
            content.encode("utf-8").decode("utf-8")
            return True, content, ""
        except UnicodeError:
            return False, "", "Clipboard contains invalid Unicode characters"
    except FileNotFoundError:
        if PLATFORM == "darwin":
            return (
                False,
                "",
                "Could not access clipboard. Please ensure you have the proper permissions.",
            )
        elif PLATFORM == "win32":
            return (
                False,
                "",
                "Windows clipboard access failed. Try installing pyperclip with: pip install pyperclip",
            )
        else:
            return (
                False,
                "",
                "xclip is not installed. Please install it with: sudo apt-get install xclip",
            )
    except CalledProcessError as e:
        return False, "", f"Failed to read clipboard: {e.stderr}"
    except Exception as e:
        return False, "", f"Unexpected error reading clipboard: {str(e)}"


def set_clipboard_content(content: str) -> Tuple[bool, str]:
    """Set content to clipboard with proper error handling.

    Args:
        content: The content to copy to clipboard

    Returns:
        Tuple[bool, str]: (success, error_message)
    """
    try:
        # Validate content is proper UTF-8 before attempting to copy
        try:
            input_bytes = content.encode("utf-8")
        except UnicodeError:
            return False, "Content contains invalid Unicode characters"

        # macOS - use pbcopy
        if PLATFORM == "darwin":
            run(["pbcopy"], input=input_bytes, check=True)
        # Windows - fallback to pyperclip if available
        elif PLATFORM == "win32":
            try:
                import pyperclip

                pyperclip.copy(content)
            except ImportError:
                return (
                    False,
                    "The pyperclip module is required for clipboard operations on Windows.\nPlease install it with: pip install pyperclip",
                )
            except Exception as e:
                return False, f"Windows clipboard error: {str(e)}"
        # Linux - use xclip
        else:
            run(["xclip", "-selection", "clipboard"], input=input_bytes, check=True)
        return True, ""
    except FileNotFoundError:
        if PLATFORM == "darwin":
            return (
                False,
                "Could not access clipboard. Please ensure you have the proper permissions.",
            )
        elif PLATFORM == "win32":
            return (
                False,
                "Windows clipboard access failed. Try installing pyperclip with: pip install pyperclip",
            )
        else:
            return (
                False,
                "xclip is not installed. Please install it with: sudo apt-get install xclip",
            )
    except CalledProcessError as e:
        return False, f"Failed to copy to clipboard: {e.stderr}"
    except Exception as e:
        return False, f"Unexpected error copying to clipboard: {str(e)}"


def show_welcome_screen():
    """Display an enhanced welcome screen for new users."""
    if st.session_state.execution_stats["total_runs"] == 0 and not st.session_state.output_logs:
        st.markdown("""
        <div style='text-align: center; padding: 2rem; background: linear-gradient(135deg, rgba(155, 108, 255, 0.1) 0%, rgba(76, 181, 255, 0.1) 100%); border-radius: 10px; margin: 1rem 0;'>
            <h2>🎉 Welcome to Fabric Pattern Studio!</h2>
            <p style='font-size: 1.1em; color: #666;'>Your AI-powered pattern execution environment</p>
        </div>
        """, unsafe_allow_html=True)

        col1, col2, col3 = st.columns(3)

        with col1:
            st.markdown("""
            ### 🚀 Quick Start
            1. Select patterns from the sidebar
            2. Enter your input text
            3. Click "Execute Patterns"
            4. View and manage results
            """)

        with col2:
            st.markdown("""
            ### 🎯 Features
            - **Smart Discovery**: Search by tags & descriptions
            - **Chain Mode**: Connect patterns sequentially
            - **Enhanced Analytics**: Performance tracking
            - **Real-time Feedback**: Rate pattern outputs
            """)

        with col3:
            st.markdown("""
            ### 💡 Pro Tips
            - ⭐ Star your best outputs
            - 📊 Monitor execution stats
            - 🔗 Use chain mode for workflows
            - 📋 Use clipboard for quick input
            """)

        # Quick action buttons
        st.markdown("### 🎬 Quick Actions")
        col1, col2, col3, col4 = st.columns(4)

        with col1:
            if st.button("🎨 Browse Patterns", use_container_width=True):
                patterns = get_patterns()
                if patterns:
                    st.success(f"Found {len(patterns)} patterns!")
                    with st.expander("Available Patterns", expanded=True):
                        for pattern in patterns[:10]:  # Show first 10
                            st.write(f"• {pattern}")
                        if len(patterns) > 10:
                            st.write(f"... and {len(patterns) - 10} more")
                else:
                    st.warning("No patterns found. Create some first!")

        with col2:
            if st.button("📊 View Analytics", use_container_width=True):
                st.info("Analytics will appear after you run some patterns!")

        with col3:
            if st.button("🔗 Try Chain Mode", use_container_width=True):
                st.session_state.view_selector = "🚀 Run"
                st.info("Use Chain Mode to connect multiple patterns!")

        with col4:
            if st.button("⚙️ Manage Patterns", use_container_width=True):
                st.session_state.view_selector = "⚙️ Manage"
                st.info("Pattern management tools are in the sidebar!")


def apply_custom_styling():
    """Apply enhanced custom styling to the app."""
    st.markdown("""
    <style>
        /* Enhanced gradient background */
        .stApp {
            background: linear-gradient(180deg,
                rgba(25, 25, 35, 0.95) 0%,
                rgba(35, 35, 45, 0.95) 50%,
                rgba(25, 25, 35, 0.95) 100%);
        }

        /* Enhanced header styling */
        .fabric-header {
            padding: 1.5rem;
            margin-bottom: 1.5rem;
            background: linear-gradient(135deg,
                rgba(155, 108, 255, 0.15) 0%,
                rgba(76, 181, 255, 0.15) 100%);
            border-radius: 12px;
            border: 1px solid rgba(155, 108, 255, 0.2);
            backdrop-filter: blur(10px);
        }

        .fabric-title {
            font-size: 2.8em;
            margin: 0;
            background: linear-gradient(135deg, #9B6CFF 0%, #4CB5FF 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            font-weight: 700;
            text-align: center;
            text-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }

        /* Enhanced metrics */
        .metric-container {
            background: rgba(255, 255, 255, 0.05);
            padding: 1rem;
            border-radius: 8px;
            border: 1px solid rgba(255, 255, 255, 0.1);
        }

        /* Enhanced buttons */
        .stButton > button {
            border-radius: 8px;
            border: 1px solid rgba(155, 108, 255, 0.3);
            background: linear-gradient(135deg,
                rgba(155, 108, 255, 0.1) 0%,
                rgba(76, 181, 255, 0.1) 100%);
            transition: all 0.3s ease;
        }

        .stButton > button:hover {
            border-color: rgba(155, 108, 255, 0.6);
            background: linear-gradient(135deg,
                rgba(155, 108, 255, 0.2) 0%,
                rgba(76, 181, 255, 0.2) 100%);
            transform: translateY(-1px);
            box-shadow: 0 4px 12px rgba(155, 108, 255, 0.3);
        }

        /* Enhanced expanders */
        .streamlit-expanderHeader {
            background: rgba(255, 255, 255, 0.05);
            border-radius: 8px;
            border: 1px solid rgba(255, 255, 255, 0.1);
        }

        /* Enhanced chat messages */
        .stChatMessage {
            background: rgba(255, 255, 255, 0.05);
            border-radius: 12px;
            border: 1px solid rgba(255, 255, 255, 0.1);
            margin: 0.5rem 0;
        }

        /* Enhanced status containers */
        .stStatus {
            background: rgba(255, 255, 255, 0.05);
            border-radius: 8px;
            border: 1px solid rgba(155, 108, 255, 0.2);
        }

        /* Enhanced pills */
        .stPills {
            background: rgba(255, 255, 255, 0.05);
            border-radius: 8px;
            padding: 0.5rem;
        }

        /* Enhanced segmented control */
        .stSegmentedControl {
            background: rgba(255, 255, 255, 0.05);
            border-radius: 8px;
            border: 1px solid rgba(255, 255, 255, 0.1);
        }

        /* Enhanced sidebar */
        .css-1d391kg {
            background: linear-gradient(180deg,
                rgba(20, 20, 30, 0.95) 0%,
                rgba(30, 30, 40, 0.95) 100%);
        }

        /* Enhanced assistant avatar */
        .assistant-container {
            position: fixed;
            right: 20px;
            bottom: 40px;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 8px;
            z-index: 1000;
        }

        .assistant-avatar {
            width: 48px;
            height: 48px;
            background: linear-gradient(135deg,
                rgba(155, 108, 255, 0.2) 0%,
                rgba(76, 181, 255, 0.2) 100%);
            border: 2px solid rgba(155, 108, 255, 0.4);
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            cursor: pointer;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            backdrop-filter: blur(12px);
            box-shadow: 0 4px 16px rgba(155, 108, 255, 0.3);
        }

        .assistant-avatar:hover {
            background: linear-gradient(135deg,
                rgba(155, 108, 255, 0.3) 0%,
                rgba(76, 181, 255, 0.3) 100%);
            border-color: rgba(155, 108, 255, 0.6);
            transform: translateY(-3px) scale(1.05);
            box-shadow: 0 8px 24px rgba(155, 108, 255, 0.4);
        }

        .assistant-avatar::before {
            content: "🤖";
            font-size: 22px;
            opacity: 0.8;
            transition: all 0.3s ease;
        }

        .assistant-avatar:hover::before {
            opacity: 1;
            transform: scale(1.1);
        }

        /* Enhanced signature */
        .signature {
            position: fixed;
            right: 10px;
            bottom: 10px;
            font-size: 0.7em;
            color: rgba(155, 108, 255, 0.4);
            z-index: 999;
            text-decoration: none;
            font-family: 'JetBrains Mono', monospace;
            font-weight: 500;
            transition: all 0.3s ease;
        }

        .signature:hover {
            color: rgba(155, 108, 255, 0.8);
            transform: translateY(-1px);
        }

        /* Toast notifications */
        .stToast {
            background: rgba(155, 108, 255, 0.1);
            border: 1px solid rgba(155, 108, 255, 0.3);
            border-radius: 8px;
            backdrop-filter: blur(8px);
        }
    </style>
    """, unsafe_allow_html=True)


def main():
    """Main function to run the Streamlit app."""
    logger.info("Starting Fabric Pattern Studio")
    # Set page config
    st.set_page_config(
        page_title="Fabric Pattern Studio",
        page_icon="🧬",
        layout="wide",
        initial_sidebar_state="expanded",
    )

    try:

        # Apply enhanced custom styling
        apply_custom_styling()

        # Enhanced header with modern styling
        st.markdown(
            """
            <div class="fabric-header">
                <h1 class="fabric-title">✨ Fabric Pattern Studio</h1>
                <p style='text-align: center; color: rgba(155, 108, 255, 0.8); font-size: 1.1em; margin: 0.5rem 0 0 0;'>
                    AI-Powered Pattern Execution Environment
                </p>
            </div>
            <div class="assistant-container">
                <div class="assistant-avatar" onclick="window.open('https://github.com/danielmiessler/fabric', '_blank')" title="Visit Fabric on GitHub"></div>
            </div>
            <a href="https://github.com/sosacrazy126" target="_blank" class="signature" title="Made by zo6">made by zo6 ✨</a>
        """,
            unsafe_allow_html=True,
        )

        initialize_session_state()

        if not st.session_state.config_loaded:
            logger.info("Loading initial configuration")
            success = load_configuration()
            if not success:
                logger.error("Failed to load configuration")
                st.error("Failed to load configuration. Please check your .env file.")
                st.stop()

        with st.sidebar:
            # Enhanced GitHub link with stats
            st.markdown(
                """
                <div style='text-align: center; margin-bottom: 1rem;'>
                    <a href="https://github.com/danielmiessler/fabric" target="_blank">
                        <img src="https://img.shields.io/github/stars/danielmiessler/fabric?style=social" alt="GitHub Repo">
                    </a>
                    <br>
                    <small style='color: #666;'>⭐ Star the project!</small>
                </div>
                """,
                unsafe_allow_html=True,
            )

            # Enhanced configuration section
            load_models_and_providers()

            st.markdown("---")

            # Enhanced navigation with icons
            st.header("🧭 Navigation")
            try:
                view = st.segmented_control(
                    "Select View",
                    ["🚀 Run", "⚙️ Manage", "📊 Analytics"],
                    key="view_selector",
                )

                # Map segmented control values to full names
                view_mapping = {
                    "🚀 Run": "Run Patterns",
                    "⚙️ Manage": "Pattern Management",
                    "📊 Analytics": "Analysis Dashboard"
                }
                view = view_mapping.get(view, "Run Patterns")
            except (AttributeError, TypeError):
                # Fallback to radio buttons if segmented_control is not available
                view = st.radio(
                    "Select View",
                    ["Run Patterns", "Pattern Management", "Analysis Dashboard"],
                    key="view_selector_radio",
                )
            logger.debug(f"Selected view: {view}")

            st.markdown("---")

            # User preferences section
            with st.expander("⚙️ Preferences", expanded=False):
                st.markdown("**🎨 Interface**")

                # Compact view toggle with fallback
                try:
                    compact_view = st.toggle(
                        "📱 Compact View",
                        value=st.session_state.user_preferences["compact_view"],
                        help="Reduce spacing and use smaller components"
                    )
                except (AttributeError, TypeError):
                    compact_view = st.checkbox(
                        "📱 Compact View",
                        value=st.session_state.user_preferences["compact_view"],
                        help="Reduce spacing and use smaller components"
                    )
                st.session_state.user_preferences["compact_view"] = compact_view

                # Notifications toggle with fallback
                try:
                    enable_notifications = st.toggle(
                        "🔔 Notifications",
                        value=st.session_state.user_preferences["enable_notifications"],
                        help="Show toast notifications for actions"
                    )
                except (AttributeError, TypeError):
                    enable_notifications = st.checkbox(
                        "🔔 Notifications",
                        value=st.session_state.user_preferences["enable_notifications"],
                        help="Show toast notifications for actions"
                    )
                st.session_state.user_preferences["enable_notifications"] = enable_notifications

                # Timestamps toggle with fallback
                try:
                    show_timestamps = st.toggle(
                        "🕐 Show Timestamps",
                        value=st.session_state.user_preferences["show_timestamps"],
                        help="Display timestamps in outputs"
                    )
                except (AttributeError, TypeError):
                    show_timestamps = st.checkbox(
                        "🕐 Show Timestamps",
                        value=st.session_state.user_preferences["show_timestamps"],
                        help="Display timestamps in outputs"
                    )
                st.session_state.user_preferences["show_timestamps"] = show_timestamps

                st.markdown("**💾 Data**")

                # Auto-save toggle with fallback
                try:
                    auto_save = st.toggle(
                        "💾 Auto-save Results",
                        value=st.session_state.user_preferences["auto_save"],
                        help="Automatically save successful executions"
                    )
                except (AttributeError, TypeError):
                    auto_save = st.checkbox(
                        "💾 Auto-save Results",
                        value=st.session_state.user_preferences["auto_save"],
                        help="Automatically save successful executions"
                    )
                st.session_state.user_preferences["auto_save"] = auto_save

                # Clear data button
                if st.button("🗑️ Clear All Data", type="secondary", use_container_width=True):
                    if st.checkbox("⚠️ Confirm deletion"):
                        st.session_state.output_logs = []
                        st.session_state.starred_outputs = []
                        st.session_state.execution_stats = {
                            "total_runs": 0,
                            "successful_runs": 0,
                            "failed_runs": 0,
                            "avg_execution_time": 0
                        }
                        save_outputs()
                        st.toast("All data cleared!", icon="🗑️")
                        st.rerun()

            # Quick stats in sidebar
            with st.expander("📊 Quick Stats", expanded=False):
                stats = st.session_state.execution_stats
                st.metric("Total Executions", stats["total_runs"])
                if stats["total_runs"] > 0:
                    success_rate = (stats["successful_runs"] / stats["total_runs"]) * 100
                    st.metric("Success Rate", f"{success_rate:.1f}%")
                    st.metric("Avg Time", f"{stats['avg_execution_time']:.1f}s")

                st.metric("Saved Outputs", len(st.session_state.output_logs))
                st.metric("Starred Items", len(st.session_state.starred_outputs))

                # Pattern descriptions stats
                descriptions_data = load_pattern_descriptions()
                if descriptions_data:
                    total_patterns = len(descriptions_data)
                    pending_descriptions = len([p for p in descriptions_data if p.get("description") == "[Description pending]"])
                    st.metric("Pattern Descriptions", f"{total_patterns - pending_descriptions}/{total_patterns}")

                    all_tags = get_all_tags(descriptions_data)
                    st.metric("Available Tags", len(all_tags))

            # Session info
            with st.expander("ℹ️ Session Info", expanded=False):
                st.caption(f"**Session ID:** {st.session_state.current_session_id[:8]}...")
                st.caption(f"**Platform:** {PLATFORM}")
                st.caption(f"**Patterns Dir:** {pattern_dir}")

                # Show current configuration
                if st.session_state.config.get("vendor") and st.session_state.config.get("model"):
                    st.success(f"✅ {st.session_state.config['vendor']} - {st.session_state.config['model']}")
                else:
                    st.warning("⚠️ No model configured")

        if view != st.session_state.get("current_view"):
            st.session_state["current_view"] = view

        # Show welcome screen for new users
        show_welcome_screen()

        if view == "Run Patterns":
            patterns = get_patterns()
            logger.debug(f"Available patterns: {patterns}")

            if not patterns:
                logger.warning("No patterns available")
                st.warning("No patterns available. Create a pattern first.")
                return

            # Enhanced tabs with icons
            tabs = st.tabs(["🚀 Execute", "📊 Analysis"])

            with tabs[0]:
                # Enhanced header with stats
                col1, col2, col3, col4 = st.columns([2, 1, 1, 1])
                with col1:
                    st.header("🎯 Pattern Execution")
                with col2:
                    st.metric("Total Runs", st.session_state.execution_stats["total_runs"])
                with col3:
                    success_rate = (
                        st.session_state.execution_stats["successful_runs"] /
                        max(st.session_state.execution_stats["total_runs"], 1) * 100
                    )
                    st.metric("Success Rate", f"{success_rate:.1f}%")
                with col4:
                    avg_time = st.session_state.execution_stats["avg_execution_time"]
                    st.metric("Avg Time", f"{avg_time:.1f}s")

                # Enhanced pattern selection
                st.subheader("🎨 Select Patterns")
                selected_patterns = enhanced_pattern_selector(patterns, "main_patterns")
                st.session_state.selected_patterns = selected_patterns

                if selected_patterns:
                    # Enhanced pattern details with descriptions and tags
                    with st.expander("📋 Pattern Details", expanded=False):
                        descriptions_data = load_pattern_descriptions()

                        for pattern in selected_patterns:
                            # Get pattern info
                            pattern_info = get_pattern_description_and_tags(pattern, descriptions_data)

                            # Create pattern card
                            with st.container():
                                col1, col2 = st.columns([3, 1])

                                with col1:
                                    st.markdown(f"### 🎯 {pattern}")

                                    # Show description
                                    description = pattern_info["description"]
                                    if description and description != "No description available":
                                        st.markdown(f"**Description:** {description}")
                                    else:
                                        # Fallback to system.md preview
                                        metadata = get_pattern_metadata(pattern)
                                        if metadata:
                                            preview = metadata[:200] + "..." if len(metadata) > 200 else metadata
                                            st.markdown(f"**Preview:** {preview}")
                                        else:
                                            st.info("No description available")

                                    # Show tags only in a subtle way (small text, muted)
                                    tags = pattern_info["tags"]
                                    if tags and len(tags) > 0:
                                        # Only show if user wants to see categories
                                        if st.session_state.get("show_pattern_categories", False):
                                            tag_display = " • ".join(tags)
                                            st.caption(f"Categories: {tag_display}")
                                        else:
                                            # Just show a small indicator that categories exist
                                            st.caption(f"📂 {len(tags)} categories available")

                                with col2:
                                    # Pattern feedback display
                                    if pattern in st.session_state.pattern_feedback:
                                        feedback = st.session_state.pattern_feedback[pattern]
                                        if feedback == 1:
                                            st.success("👍 Liked")
                                        else:
                                            st.error("👎 Disliked")

                                    # Quick action buttons
                                    if st.button("ℹ️ Full Details", key=f"details_{pattern}"):
                                        with st.expander(f"Full details for {pattern}", expanded=True):
                                            full_metadata = get_pattern_metadata(pattern)
                                            if full_metadata:
                                                st.code(full_metadata, language="markdown")
                                            else:
                                                st.warning("No system.md file found")

                                st.markdown("---")

                    # Enhanced input section
                    st.subheader("📝 Input Configuration")

                    # Input method selection with segmented control
                    try:
                        input_method = st.segmented_control(
                            "Input Method",
                            ["📋 Clipboard", "✏️ Manual"],
                            default="✏️ Manual"
                        )
                    except (AttributeError, TypeError):
                        input_method = st.radio(
                            "Input Method",
                            ["📋 Clipboard", "✏️ Manual"],
                            horizontal=True
                        )

                    if input_method == "📋 Clipboard":
                        col_load, col_preview = st.columns([2, 1])
                        with col_load:
                            if st.button(
                                "📋 Load from Clipboard",
                                use_container_width=True,
                                type="secondary"
                            ):
                                with st.spinner("Loading from clipboard..."):
                                    success, content, error = get_clipboard_content()
                                    if success:
                                        # Validate clipboard content
                                        is_valid, error_message = validate_input_content(content)
                                        if not is_valid:
                                            st.error(f"Invalid clipboard content: {error_message}")
                                        else:
                                            # Sanitize clipboard content
                                            sanitized_content = sanitize_input_content(content)
                                            if sanitized_content != content:
                                                st.toast(
                                                    "Content was automatically sanitized",
                                                    icon="🧹"
                                                )

                                            st.session_state.input_content = sanitized_content
                                            st.session_state.show_preview = True
                                            st.toast("Content loaded successfully!", icon="✅")
                                    else:
                                        st.error(error)

                        with col_preview:
                            preview_enabled = st.toggle(
                                "👁 Preview",
                                value=st.session_state.get("show_preview", False),
                                key="preview_toggle"
                            )
                            st.session_state.show_preview = preview_enabled

                    else:  # Manual input
                        st.session_state.input_content = st.text_area(
                            "✏️ Enter your input text",
                            value=st.session_state.get("input_content", ""),
                            height=200,
                            placeholder="Type or paste your content here...",
                            help="This text will be processed by the selected patterns"
                        )

                    if (
                        st.session_state.get("show_preview", False)
                        or input_method == "Manual Input"
                    ):
                        if st.session_state.get("input_content"):
                            enhance_input_preview()

                    # Enhanced execution options
                    st.subheader("⚙️ Execution Options")

                    col1, col2 = st.columns(2)
                    with col1:
                        chain_mode = st.toggle(
                            "🔗 Chain Mode",
                            help="Execute patterns in sequence, passing output of each pattern as input to the next",
                        )
                    with col2:
                        auto_save = st.toggle(
                            "💾 Auto-save results",
                            value=st.session_state.user_preferences["auto_save"],
                            help="Automatically save successful executions"
                        )
                        st.session_state.user_preferences["auto_save"] = auto_save

                    if chain_mode and len(selected_patterns) > 1:
                        st.info("🔗 Patterns will be executed in sequence")
                        st.markdown("##### 📋 Pattern Execution Order:")

                        # Enhanced pattern reordering with data editor
                        patterns_df = pd.DataFrame({
                            "Order": range(1, len(selected_patterns) + 1),
                            "Pattern": selected_patterns,
                            "Status": ["⏳ Pending"] * len(selected_patterns)
                        })

                        edited_df = st.data_editor(
                            patterns_df,
                            use_container_width=True,
                            key="pattern_reorder",
                            hide_index=True,
                            column_config={
                                "Order": st.column_config.NumberColumn("Order", width="small"),
                                "Pattern": st.column_config.TextColumn("Pattern Name"),
                                "Status": st.column_config.TextColumn("Status", width="small")
                            },
                        )

                        # Update selected patterns if order changed
                        new_patterns = edited_df["Pattern"].tolist()
                        if new_patterns != selected_patterns:
                            st.session_state.selected_patterns = new_patterns

                    # Pattern Variables UI
                    pattern_variables = {}
                    if selected_patterns:
                        # Check if any selected patterns have variables
                        patterns_with_vars = []
                        for pattern in selected_patterns:
                            variables = detect_pattern_variables(pattern)
                            if variables:
                                patterns_with_vars.append((pattern, variables))
                        
                        if patterns_with_vars:
                            st.markdown("---")
                            st.subheader("🔧 Pattern Variables")
                            st.info(f"The following patterns require variables to be set:")
                            
                            # Create tabs for each pattern with variables
                            if len(patterns_with_vars) == 1:
                                pattern, variables = patterns_with_vars[0]
                                st.markdown(f"**{pattern}**")
                                pattern_variables[pattern] = render_pattern_variables_ui(
                                    variables, f"vars_{pattern}"
                                )
                            else:
                                # Multiple patterns with variables - use tabs
                                tab_names = [f"{pattern} ({len(vars)} vars)" for pattern, vars in patterns_with_vars]
                                tabs = st.tabs(tab_names)
                                
                                for i, (pattern, variables) in enumerate(patterns_with_vars):
                                    with tabs[i]:
                                        pattern_variables[pattern] = render_pattern_variables_ui(
                                            variables, f"vars_{pattern}"
                                        )

                    # Enhanced execution button section
                    st.markdown("---")
                    col1, col2, col3 = st.columns([2, 1, 1])

                    with col1:
                        # Check if all required variables are filled
                        can_execute = True
                        missing_vars_info = []
                        
                        if pattern_variables:
                            for pattern_name, variables in pattern_variables.items():
                                required_vars = detect_pattern_variables(pattern_name)
                                is_valid, missing_vars = validate_pattern_variables(variables, required_vars)
                                if not is_valid:
                                    can_execute = False
                                    missing_vars_info.append(f"{pattern_name}: {', '.join(missing_vars)}")
                        
                        run_button = st.button(
                            "🚀 Execute Patterns",
                            type="primary",
                            use_container_width=True,
                            disabled=not st.session_state.input_content or not selected_patterns or not can_execute
                        )
                        
                        # Show missing variables warning
                        if missing_vars_info:
                            st.warning(f"⚠️ Missing required variables:\n" + "\n".join([f"• {info}" for info in missing_vars_info]))

                    with col2:
                        if st.button("🧹 Clear Output", use_container_width=True):
                            st.session_state.chat_output = []
                            st.toast("Output cleared!", icon="🧹")
                            st.rerun()

                    with col3:
                        if st.button("📊 View Stats", use_container_width=True):
                            st.session_state.show_stats = not st.session_state.get("show_stats", False)

                    # Show execution stats if requested
                    if st.session_state.get("show_stats", False):
                        with st.expander("📊 Execution Statistics", expanded=True):
                            stats = st.session_state.execution_stats
                            col1, col2, col3, col4 = st.columns(4)
                            with col1:
                                st.metric("Total Runs", stats["total_runs"])
                            with col2:
                                st.metric("Successful", stats["successful_runs"])
                            with col3:
                                st.metric("Failed", stats["failed_runs"])
                            with col4:
                                st.metric("Avg Time", f"{stats['avg_execution_time']:.2f}s")

                    # Execute patterns
                    if run_button:
                        if not st.session_state.input_content:
                            st.error("⚠️ Please provide input content.")
                        elif not selected_patterns:
                            st.error("⚠️ Please select at least one pattern.")
                        else:
                            if chain_mode:
                                # Execute pattern chain with enhanced UI
                                chain_results = execute_pattern_chain(
                                    selected_patterns,
                                    st.session_state.input_content,
                                    pattern_variables if pattern_variables else None
                                )

                                # Enhanced chain results display
                                st.markdown("## 🔗 Chain Execution Results")

                                # Success indicator
                                if chain_results["metadata"]["success"]:
                                    st.success("✅ Chain execution completed successfully!")
                                else:
                                    st.error("❌ Chain execution failed")

                                # Show sequence with visual flow
                                st.markdown("### 📋 Pattern Sequence")
                                sequence_text = " ➡️ ".join(chain_results["sequence"])
                                st.code(sequence_text)

                                # Enhanced stage display
                                st.markdown("### 🔄 Execution Stages")
                                for i, stage in enumerate(chain_results["stages"], 1):
                                    status_icon = "✅" if stage["success"] else "❌"
                                    with st.expander(
                                        f"{status_icon} Stage {i}: {stage['pattern']}",
                                        expanded=not stage["success"],
                                    ):
                                        col1, col2 = st.columns(2)
                                        with col1:
                                            st.markdown("#### 📥 Input")
                                            st.code(stage["input"][:500] + "..." if len(stage["input"]) > 500 else stage["input"])
                                        with col2:
                                            st.markdown("#### 📤 Output")
                                            if stage["success"]:
                                                st.markdown(stage["output"])
                                            else:
                                                st.error(stage["error"])

                                # Show final output with enhanced formatting
                                if chain_results["metadata"]["success"]:
                                    st.markdown("### 🎯 Final Output")
                                    st.markdown(chain_results["final_output"])
                                    st.session_state.chat_output.append(chain_results["final_output"])

                                    # Add feedback UI for chain result
                                    show_pattern_feedback_ui("Chain Result", chain_results["final_output"])
                            else:
                                # Enhanced normal pattern execution
                                outputs = execute_patterns_enhanced(
                                    selected_patterns,
                                    pattern_variables=pattern_variables if pattern_variables else None
                                )
                                if outputs:
                                    st.session_state.last_run_outputs = outputs
                                    st.session_state.show_success_toast = True
                                    st.rerun()

                    # Display output from the last run and then clear it
                    if st.session_state.get('last_run_outputs'):
                        st.markdown("### 🎯 Pattern Output")
                        # Join and display the outputs
                        output_display = "\n\n---\n\n".join(st.session_state.last_run_outputs)
                        st.markdown(output_display)

                        # Show feedback for each individual output
                        for output in st.session_state.last_run_outputs:
                            if output.startswith("### 🎯"):
                                pattern_name = output.split("\n")[0].replace("### 🎯 ", "")
                                show_pattern_feedback_ui(pattern_name, output)
                        
                        # Clear the outputs from state so they don't re-display
                        st.session_state.last_run_outputs = []

                    # Enhanced output display
                    if st.session_state.chat_output:
                        st.markdown("---")

                        # Output header with actions
                        col1, col2, col3, col4 = st.columns([2, 1, 1, 1])
                        with col1:
                            st.header("📤 Pattern Outputs")
                        with col2:
                            if st.button("📋 Copy All", use_container_width=True):
                                all_outputs = "\n\n".join(st.session_state.chat_output)
                                success, error = set_clipboard_content(all_outputs)
                                if success:
                                    st.toast("All outputs copied!", icon="📋")
                                else:
                                    st.error(error)
                        with col3:
                            if st.button("💾 Save All", use_container_width=True):
                                timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
                                all_content = "\n\n".join(st.session_state.chat_output)
                                save_output_log("Batch Execution", "Multiple patterns", all_content, timestamp)
                                st.toast("Outputs saved!", icon="💾")
                        with col4:
                            if st.button("🧹 Clear", use_container_width=True):
                                st.session_state.chat_output = []
                                st.toast("Outputs cleared!", icon="🧹")
                                st.rerun()

                        # Display outputs with enhanced formatting
                        for i, message in enumerate(st.session_state.chat_output):
                            # Create a container for each output
                            with st.container():
                                # Extract pattern name for better display
                                pattern_name = "Output"
                                if message.startswith("### 🎯"):
                                    pattern_name = message.split("\n")[0].replace("### 🎯 ", "")
                                elif message.startswith("###"):
                                    pattern_name = message.split("\n")[0].replace("### ", "")

                                # Output header with metadata
                                col1, col2, col3 = st.columns([2, 1, 1])
                                with col1:
                                    st.markdown(f"**Output {i+1}:** {pattern_name}")
                                with col2:
                                    word_count = len(message.split())
                                    st.caption(f"📊 {word_count} words")
                                with col3:
                                    if st.button("📋", key=f"copy_output_{i}", help="Copy this output"):
                                        success, error = set_clipboard_content(message)
                                        if success:
                                            st.toast("Copied!", icon="📋")
                                        else:
                                            st.error(error)

                                # Display the actual output
                                st.markdown(message)

                                # Add feedback and actions
                                col1, col2, col3 = st.columns([1, 1, 2])
                                with col1:
                                    feedback = st.feedback(
                                        "thumbs",
                                        key=f"output_feedback_{i}"
                                    )
                                    if feedback is not None:
                                        if feedback == 1:
                                            st.toast("Thanks for the positive feedback!", icon="👍")
                                        else:
                                            st.toast("Thanks for the feedback!", icon="👎")

                                with col2:
                                    if st.button("⭐ Star", key=f"star_output_{i}"):
                                        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
                                        star_output_enhanced(pattern_name, message, timestamp)

                                st.markdown("---")

                else:
                    # Enhanced empty state
                    st.info("🎯 Select one or more patterns to run and see the magic happen!")

                    # Show some helpful tips
                    with st.expander("💡 Quick Tips", expanded=False):
                        st.markdown("""
                        **Getting Started:**
                        1. 🎨 Select patterns using search and tag filters
                        2. ✏️ Enter your input text or load from clipboard
                        3. 🚀 Click "Execute Patterns"

                        **Pro Tips:**
                        - 🏷️ Use tag filters to find relevant patterns
                        - 🔍 Search by pattern name or description
                        - 🔗 Use Chain Mode to connect patterns
                        - ⭐ Star your favorite outputs
                        - 📊 Check the Analysis tab for insights
                        """)

            with tabs[1]:
                st.header("📊 Output Analysis")

                if st.session_state.chat_output:
                    # Enhanced analysis with metrics
                    col1, col2, col3 = st.columns(3)
                    with col1:
                        st.metric("Total Outputs", len(st.session_state.chat_output))
                    with col2:
                        total_chars = sum(len(output) for output in st.session_state.chat_output)
                        st.metric("Total Characters", f"{total_chars:,}")
                    with col3:
                        avg_length = total_chars / len(st.session_state.chat_output)
                        st.metric("Avg Length", f"{avg_length:.0f}")

                    # Filter and search
                    search_analysis = st.text_input("🔍 Search in outputs", placeholder="Search content...")

                    # Display outputs with enhanced formatting
                    filtered_outputs = [
                        (i, output) for i, output in enumerate(reversed(st.session_state.chat_output), 1)
                        if not search_analysis or search_analysis.lower() in output.lower()
                    ]

                    for i, output in filtered_outputs:
                        # Extract pattern name if available
                        pattern_name = "Unknown"
                        if output.startswith("### 🎯"):
                            pattern_name = output.split("\n")[0].replace("### 🎯 ", "")
                        elif output.startswith("###"):
                            pattern_name = output.split("\n")[0].replace("### ", "")

                        with st.expander(f"📄 Output #{i} - {pattern_name}", expanded=False):
                            # Add copy button and word count
                            col1, col2, col3 = st.columns([2, 1, 1])
                            with col1:
                                st.markdown(f"**Pattern:** {pattern_name}")
                            with col2:
                                word_count = len(output.split())
                                st.caption(f"Words: {word_count}")
                            with col3:
                                if st.button("📋 Copy", key=f"copy_analysis_{i}"):
                                    success, error = set_clipboard_content(output)
                                    if success:
                                        st.toast("Copied to clipboard!", icon="📋")
                                    else:
                                        st.error(error)

                            st.markdown(output)

                            # Add analysis metrics for this output
                            with st.expander("📈 Output Metrics", expanded=False):
                                lines = output.count('\n') + 1
                                chars = len(output)
                                words = len(output.split())

                                col1, col2, col3 = st.columns(3)
                                with col1:
                                    st.metric("Lines", lines)
                                with col2:
                                    st.metric("Words", words)
                                with col3:
                                    st.metric("Characters", chars)
                else:
                    st.info("🎯 Run some patterns to see output analysis here.")

        elif view == "Pattern Management":
            create_tab, edit_tab, delete_tab, descriptions_tab = st.tabs(["Create", "Edit", "Delete", "📋 Descriptions"])

            with create_tab:
                st.header("Create New Pattern")
                creation_mode = st.radio(
                    "Creation Mode",
                    ["Simple Editor", "Advanced (Wizard)"],
                    key="creation_mode_main",
                    horizontal=True,
                )

                if creation_mode == "Simple Editor":
                    pattern_creation_ui()
                else:
                    pattern_creation_wizard()

            with edit_tab:
                st.header("Edit Patterns")
                patterns = get_patterns()
                if not patterns:
                    st.warning("No patterns available. Create a pattern first.")
                else:
                    selected_pattern = st.selectbox(
                        "Select Pattern to Edit", [""] + patterns
                    )
                    if selected_pattern:
                        pattern_editor(selected_pattern)

            with delete_tab:
                st.header("Delete Patterns")
                patterns = get_patterns()
                if not patterns:
                    st.warning("No patterns available.")
                else:
                    patterns_to_delete = st.multiselect(
                        "Select Patterns to Delete",
                        patterns,
                        key="delete_patterns_selector",
                    )

                    if patterns_to_delete:
                        st.warning(
                            f"You are about to delete {len(patterns_to_delete)} pattern(s):"
                        )
                        for pattern in patterns_to_delete:
                            st.markdown(f"- {pattern}")

                        confirm_delete = st.checkbox(
                            "I understand that this action cannot be undone"
                        )

                        if st.button(
                            "🗑️ Delete Selected Patterns",
                            type="primary",
                            disabled=not confirm_delete,
                        ):
                            if confirm_delete:
                                for pattern in patterns_to_delete:
                                    success, message = delete_pattern(pattern)
                                    if success:
                                        st.success(f"✓ {pattern}: {message}")
                                    else:
                                        st.error(f"✗ {pattern}: {message}")
                                st.experimental_rerun()
                            else:
                                st.error(
                                    "Please confirm deletion by checking the box above."
                                )
                    else:
                        st.info("Select one or more patterns to delete.")

            with descriptions_tab:
                show_pattern_management_ui()

        else:
            st.header("Pattern Output History")

            # Create tabs for All Outputs and Starred Outputs
            all_tab, starred_tab = st.tabs(["All Outputs", "⭐ Starred"])

            with all_tab:
                if not st.session_state.output_logs:
                    st.info(
                        "No pattern outputs recorded yet. Run some patterns to see their logs here."
                    )
                else:
                    for i, log in enumerate(reversed(st.session_state.output_logs)):
                        with st.expander(
                            f"Output #{len(st.session_state.output_logs)-i} - {log['pattern_name']} ({log['timestamp']})",
                            expanded=False,
                        ):
                            st.markdown("### Input")
                            st.code(log["input"], language="text")
                            st.markdown("### Output")
                            st.markdown(log["output"])

                            # Check if this output is already starred
                            is_starred = any(
                                s["timestamp"] == log["timestamp"]
                                for s in st.session_state.starred_outputs
                            )

                            col1, col2 = st.columns([1, 4])
                            with col1:
                                if not is_starred:
                                    if st.button(
                                        "⭐ Star",
                                        key=f"star_{i}",
                                        use_container_width=True,
                                    ):
                                        st.session_state.starring_output = (
                                            len(st.session_state.output_logs) - i - 1
                                        )
                                        st.session_state.temp_star_name = ""
                                else:
                                    st.write("⭐ Starred")

                            with col2:
                                if st.button("📋 Copy Output", key=f"copy_{i}"):
                                    success, error = set_clipboard_content(
                                        log["output"]
                                    )
                                    if success:
                                        st.success("Output copied to clipboard!")
                                    else:
                                        st.error(error)

                            # Show starring form inside the expander if this is the output being starred
                            if (
                                st.session_state.starring_output
                                == len(st.session_state.output_logs) - i - 1
                            ):
                                st.markdown("---")
                                with st.form(key=f"star_name_form_{i}"):
                                    name_input = st.text_input(
                                        "Enter a name for this output (optional):",
                                        key=f"star_name_input_{i}",
                                    )
                                    col1, col2 = st.columns(2)
                                    with col1:
                                        submit = st.form_submit_button(
                                            "Save", use_container_width=True
                                        )
                                    with col2:
                                        cancel = st.form_submit_button(
                                            "Cancel", use_container_width=True
                                        )

                                    if submit:
                                        handle_star_name_input(
                                            st.session_state.starring_output, name_input
                                        )
                                        # Reset starring state after handling
                                        st.session_state.starring_output = None
                                        st.experimental_rerun()
                                    elif cancel:
                                        # Reset starring state
                                        st.session_state.starring_output = None
                                        st.experimental_rerun()

                # Remove the old starring form from the bottom
                st.markdown("---")

            with starred_tab:
                if not st.session_state.starred_outputs:
                    st.info(
                        "No starred outputs yet. Star some outputs to see them here!"
                    )
                else:
                    for i, starred in enumerate(st.session_state.starred_outputs):
                        with st.expander(
                            f"⭐ {starred.get('custom_name', f'Starred Output #{i+1}')} ({starred['timestamp']})",
                            expanded=False,
                        ):
                            col1, col2 = st.columns([3, 1])
                            with col1:
                                st.markdown(
                                    f"### {starred.get('custom_name', f'Starred Output #{i+1}')}"
                                )
                            with col2:
                                if st.button("✏️ Edit Name", key=f"edit_name_{i}"):
                                    st.session_state[f"editing_name_{i}"] = True

                            if st.session_state.get(f"editing_name_{i}", False):
                                new_name = st.text_input(
                                    "Enter new name:",
                                    value=starred.get("custom_name", ""),
                                    key=f"new_name_{i}",
                                )
                                col1, col2 = st.columns([1, 1])
                                with col1:
                                    if st.button("Save", key=f"save_name_{i}"):
                                        st.session_state.starred_outputs[i][
                                            "custom_name"
                                        ] = new_name
                                        del st.session_state[f"editing_name_{i}"]
                                        st.success("Name updated!")
                                        st.experimental_rerun()
                                with col2:
                                    if st.button("Cancel", key=f"cancel_name_{i}"):
                                        del st.session_state[f"editing_name_{i}"]
                                        st.experimental_rerun()

                            st.markdown("### Pattern")
                            st.code(starred["pattern_name"], language="text")
                            st.markdown("### Input")
                            st.code(
                                starred["input"], language="text"
                            )  # Display input as code block
                            st.markdown("### Output")
                            st.markdown(starred["output"])  # Display output as markdown

                            col1, col2 = st.columns([1, 4])
                            with col1:
                                if st.button("❌ Remove Star", key=f"unstar_{i}"):
                                    unstar_output(i)
                                    st.success("Output unstarred!")
                                    st.experimental_rerun()

                            with col2:
                                if st.button("📋 Copy Output", key=f"copy_starred_{i}"):
                                    try:
                                        run(
                                            ["xclip", "-selection", "clipboard"],
                                            input=starred["output"].encode(),
                                            check=True,
                                        )
                                        st.success("Output copied to clipboard!")
                                    except Exception as e:
                                        st.error(f"Error copying to clipboard: {e}")

                    if st.button("Clear All Starred"):
                        if st.checkbox("Confirm clearing all starred outputs"):
                            st.session_state.starred_outputs = []
                            save_outputs()  # Save after clearing
                            st.success("All starred outputs cleared!")
                            st.experimental_rerun()

    except Exception as e:
        logger.error("Unexpected error in main function", exc_info=True)
        st.error(f"An unexpected error occurred: {str(e)}")
        st.stop()


    finally:
        logger.info("Application shutdown")

if __name__ == "__main__":
    logger.info("Application startup")
    main()
