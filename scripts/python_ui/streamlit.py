import shutil
import json
import os
import streamlit as st
from subprocess import run, CalledProcessError
from dotenv import load_dotenv
import re
import secrets
import time
import logging
from typing import Dict, List, Optional, Tuple
from datetime import datetime
import sys
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import numpy as np

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
    """Load models and providers from fabric configuration."""
    try:
        st.sidebar.header("Model and Provider Selection")

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

        selected_provider = st.sidebar.selectbox(
            "Provider",
            available_providers,
            index=provider_index,
            key="provider_select",
            on_change=lambda: update_provider_selection(
                st.session_state.provider_select
            ),
        )

        if selected_provider != st.session_state.config.get("vendor"):
            update_provider_selection(selected_provider)
        st.sidebar.success(f"Using {selected_provider}")

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
            "Model", available_models, index=model_index, key=model_key
        )

        if selected_model != st.session_state.config.get("model"):
            logger.debug(f"Updating model selection to: {selected_model}")
            st.session_state.config["model"] = selected_model
            st.session_state.selected_model = selected_model

    except Exception as e:
        logger.error(f"Error loading models and providers: {str(e)}", exc_info=True)
        st.sidebar.error(f"Error loading models and providers: {str(e)}")
        st.session_state.selected_model = None
        st.session_state.config["model"] = None


def get_pattern_metadata(pattern_name):
    """Get pattern metadata from system.md."""
    pattern_path = os.path.join(pattern_dir, pattern_name, "system.md")
    if os.path.exists(pattern_path):
        with open(pattern_path, "r") as f:
            return f.read()
    return None


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


def execute_patterns(
    patterns_to_run: List[str],
    chain_mode: bool = False,
    initial_input: Optional[str] = None,
) -> List[str]:
    """Execute the selected patterns and capture their outputs."""
    logger.info(f"Executing {len(patterns_to_run)} patterns")

    st.session_state.chat_output = []
    all_outputs = []
    current_input = initial_input or st.session_state.input_content
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    # Validate configuration
    current_provider = st.session_state.config.get("vendor")
    current_model = st.session_state.config.get("model")

    if not current_provider or not current_model:
        error_msg = "Please select a provider and model first."
        logger.error(error_msg)
        st.error(error_msg)
        return all_outputs

    # Validate input content
    is_valid, error_message = validate_input_content(current_input)
    if not is_valid:
        logger.error(f"Input validation failed: {error_message}")
        st.error(f"Input validation failed: {error_message}")
        return all_outputs

    # Sanitize input content
    try:
        sanitized_input = sanitize_input_content(current_input)
        if sanitized_input != current_input:
            logger.info("Input content was sanitized")
            st.warning(
                "Input content was automatically sanitized for better compatibility."
            )

        current_input = sanitized_input
    except Exception as e:
        logger.error(f"Error sanitizing input: {str(e)}")
        st.error(f"Error processing input: {str(e)}")
        return all_outputs

    execution_info = f"**Using Model:** {current_provider} - {current_model}"
    all_outputs.append(execution_info)
    logger.info(f"Using model: {current_model} from provider: {current_provider}")

    try:
        for pattern in patterns_to_run:
            logger.info(f"Running pattern: {pattern}")
            try:
                cmd = ["fabric", "--pattern", pattern]
                logger.debug(f"Executing command: {' '.join(cmd)}")

                message = (
                    current_input if chain_mode else st.session_state.input_content
                )
                logger.debug(f"Input for pattern {pattern}:\n{message}")

                # Ensure input_data is a string
                input_data = str(message)

                # Run the command with text=True and string input
                result = run(
                    cmd, input=input_data, capture_output=True, text=True, check=True
                )

                pattern_output = result.stdout.strip()
                logger.debug(f"Raw output from pattern {pattern}:\n{pattern_output}")

                if pattern_output:
                    # Format output as markdown
                    output_msg = f"""### {pattern}

{pattern_output}"""
                    all_outputs.append(output_msg)
                    # Save to output logs with markdown formatting
                    save_output_log(pattern, message, pattern_output, timestamp)
                    if chain_mode:
                        current_input = pattern_output
                else:
                    logger.warning(f"Pattern {pattern} generated no output")
                    all_outputs.append(f"### {pattern}\n\nNo output generated.")

            except UnicodeEncodeError as e:
                error_msg = f"### {pattern}\n\n❌ Error: Input contains invalid characters: {str(e)}"
                logger.error(f"Unicode encoding error for pattern {pattern}: {str(e)}")
                all_outputs.append(error_msg)
                if chain_mode:
                    break

            except CalledProcessError as e:
                error_msg = f"### {pattern}\n\n❌ Error executing: {e.stderr.strip()}"
                logger.error(f"Pattern {pattern} failed: {e.stderr.strip()}")
                all_outputs.append(error_msg)
                if chain_mode:
                    break

            except Exception as e:
                error_msg = f"### {pattern}\n\n❌ Failed to execute: {str(e)}"
                logger.error(f"Pattern {pattern} failed: {str(e)}", exc_info=True)
                all_outputs.append(error_msg)
                if chain_mode:
                    break

    except Exception as e:
        error_msg = f"### Error\n\n❌ Error in pattern execution: {str(e)}"
        logger.error(error_msg, exc_info=True)
        st.error(error_msg)

    logger.info("Pattern execution completed")
    return all_outputs


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


def execute_pattern_chain(patterns_sequence: List[str], initial_input: str) -> Dict:
    """Execute a sequence of patterns in a chain, passing output from each to the next.

    Args:
        patterns_sequence: List of pattern names to execute in sequence
        initial_input: Initial input text to start the chain

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


def _require_access_token():
    """Ask for a shared token before the UI shows.

    If FABRIC_UI_TOKEN is set, the user must type the same token before the
    app shows. If it is not set, there is no token check, and only the
    loopback bind in .streamlit/config.toml limits access. The compare
    time does not change with the token value.
    """
    token = os.environ.get("FABRIC_UI_TOKEN", "")
    if not token:
        return
    if st.session_state.get("_authed"):
        return
    if len(token) < 16:
        logger.warning("FABRIC_UI_TOKEN is short: use a random token of 16 or more characters")
    st.title("🔒 Fabric Pattern Studio")
    entered = st.text_input("Access token", type="password")
    # Compare bytes. compare_digest refuses a str with non-ASCII characters.
    if entered and secrets.compare_digest(entered.encode(), token.encode()):
        st.session_state["_authed"] = True
        st.rerun()
    if entered:
        st.error("Invalid access token")
    st.stop()


def main():
    """Main function to run the Streamlit app."""
    logger.info("Starting Fabric Pattern Studio")
    try:
        # Set page config
        st.set_page_config(
            page_title="Fabric Pattern Studio",
            page_icon="🧬",
            layout="wide",
            initial_sidebar_state="expanded",
        )

        # Ask for the token if FABRIC_UI_TOKEN is set.
        _require_access_token()

        # Add title with gradient styling and footer signature
        st.markdown(
            """
            <style>
                [data-testid="stHeader"] {
                    background-color: rgba(0,0,0,0);
                }
                .fabric-header {
                    padding: 1rem;
                    margin-bottom: 1rem;
                    background: linear-gradient(90deg, rgba(155, 108, 255, 0.1) 0%, rgba(76, 181, 255, 0.1) 100%);
                    border-radius: 8px;
                }
                .fabric-title {
                    font-size: 2.5em;
                    margin: 0;
                    background: linear-gradient(90deg, #9B6CFF 0%, #4CB5FF 100%);
                    -webkit-background-clip: text;
                    -webkit-text-fill-color: transparent;
                    font-weight: 600;
                    text-align: center;
                }
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
                    width: 42px;
                    height: 42px;
                    background: rgba(155, 108, 255, 0.05);
                    border: 2px solid rgba(155, 108, 255, 0.1);
                    border-radius: 50%;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    cursor: pointer;
                    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                    backdrop-filter: blur(8px);
                    -webkit-backdrop-filter: blur(8px);
                }
                .assistant-avatar:hover {
                    background: rgba(155, 108, 255, 0.1);
                    border-color: rgba(155, 108, 255, 0.2);
                    transform: translateY(-2px);
                }
                .assistant-avatar::before {
                    content: "🤖";
                    font-size: 20px;
                    opacity: 0.7;
                    transition: opacity 0.3s ease;
                }
                .assistant-avatar:hover::before {
                    opacity: 0.9;
                }
                .signature {
                    position: fixed;
                    right: 10px;
                    bottom: 10px;
                    font-size: 0.7em;
                    color: rgba(155, 108, 255, 0.3);
                    z-index: 999;
                    text-decoration: none;
                    font-family: monospace;
                }
                .signature:hover {
                    color: rgba(155, 108, 255, 0.8);
                    transition: color 0.3s ease;
                }
                .stApp {
                    background: linear-gradient(180deg, rgba(25, 25, 35, 0.95) 0%, rgba(35, 35, 45, 0.95) 100%);
                }
            </style>
            <div class="fabric-header">
                <h1 class="fabric-title">Pattern Studio</h1>
            </div>
            <div class="assistant-container">
                <div class="assistant-avatar" onclick="window.open('https://github.com/danielmiessler/fabric', '_blank')"></div>
            </div>
            <a href="https://github.com/sosacrazy126" target="_blank" class="signature">made by zo6</a>
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
            # Add GitHub link
            st.markdown(
                """
                <div style='text-align: center; margin-bottom: 1rem;'>
                    <a href="https://github.com/danielmiessler/fabric" target="_blank">
                        <img src="https://img.shields.io/github/stars/danielmiessler/fabric?style=social" alt="GitHub Repo">
                    </a>
                </div>
                """,
                unsafe_allow_html=True,
            )

            st.title("Configuration")
            load_models_and_providers()

            st.markdown("---")
            st.title("Navigation")
            view = st.radio(
                "Select View",
                ["Run Patterns", "Pattern Management", "Analysis Dashboard"],
                key="view_selector",
            )
            logger.debug(f"Selected view: {view}")

        if view != st.session_state.get("current_view"):
            st.session_state["current_view"] = view

        if view == "Run Patterns":
            patterns = get_patterns()
            logger.debug(f"Available patterns: {patterns}")

            if not patterns:
                logger.warning("No patterns available")
                st.warning("No patterns available. Create a pattern first.")
                return

            tabs = st.tabs(["Run", "Analysis"])

            with tabs[0]:
                st.header("Run Patterns")
                selected_patterns = st.multiselect(
                    "Select Patterns to Run",
                    patterns,
                    default=st.session_state.selected_patterns,
                    key="selected_patterns_widget",
                )
                st.session_state.selected_patterns = selected_patterns

                if selected_patterns:
                    for pattern in selected_patterns:
                        with st.expander(f"📝 {pattern} Details", expanded=False):
                            metadata = get_pattern_metadata(pattern)
                            if metadata:
                                st.markdown(metadata)
                            else:
                                st.info("No description available")

                    st.subheader("Input")
                    input_method = st.radio(
                        "Input Method", ["Clipboard", "Manual Input"], horizontal=True
                    )

                    if input_method == "Clipboard":
                        col_load, col_preview = st.columns([2, 1])
                        with col_load:
                            if st.button(
                                "📋 Load from Clipboard", use_container_width=True
                            ):
                                success, content, error = get_clipboard_content()
                                if success:
                                    # Validate clipboard content
                                    is_valid, error_message = validate_input_content(
                                        content
                                    )
                                    if not is_valid:
                                        st.error(
                                            f"Invalid clipboard content: {error_message}"
                                        )
                                    else:
                                        # Sanitize clipboard content
                                        sanitized_content = sanitize_input_content(
                                            content
                                        )
                                        if sanitized_content != content:
                                            st.warning(
                                                "Clipboard content was automatically sanitized for better compatibility."
                                            )

                                        st.session_state.input_content = (
                                            sanitized_content
                                        )
                                        st.session_state.show_preview = True
                                        st.success("Content loaded from clipboard!")
                                else:
                                    st.error(error)

                        with col_preview:
                            if st.button("👁 Toggle Preview", use_container_width=True):
                                st.session_state.show_preview = (
                                    not st.session_state.get("show_preview", False)
                                )
                    else:
                        st.session_state.input_content = st.text_area(
                            "Enter Input Text",
                            value=st.session_state.get("input_content", ""),
                            height=200,
                        )

                    if (
                        st.session_state.get("show_preview", False)
                        or input_method == "Manual Input"
                    ):
                        if st.session_state.get("input_content"):
                            enhance_input_preview()

                    # Move chain mode checkbox before the run button
                    chain_mode = st.checkbox(
                        "Chain Mode",
                        help="Execute patterns in sequence, passing output of each pattern as input to the next",
                    )

                    if chain_mode and len(selected_patterns) > 1:
                        st.info("Patterns will be executed in the order selected above")
                        st.markdown("##### Drag to reorder patterns:")
                        # Convert patterns list to DataFrame for data editor
                        patterns_df = pd.DataFrame({"Pattern": selected_patterns})

                        edited_df = st.data_editor(
                            patterns_df,
                            use_container_width=True,
                            key="pattern_reorder",
                            hide_index=True,
                            column_config={
                                "Pattern": st.column_config.TextColumn(
                                    "Pattern", help="Drag to reorder patterns"
                                )
                            },
                        )

                        # Update selected patterns if order changed
                        new_patterns = edited_df["Pattern"].tolist()
                        if new_patterns != selected_patterns:
                            st.session_state.selected_patterns = new_patterns

                    col1, col2 = st.columns([3, 1])
                    with col1:
                        if st.button(
                            "🚀 Run Patterns", type="primary", use_container_width=True
                        ):
                            if not st.session_state.input_content:
                                st.warning("Please provide input content.")
                            else:
                                with st.spinner("Running patterns..."):
                                    if chain_mode:
                                        # Execute pattern chain
                                        chain_results = execute_pattern_chain(
                                            selected_patterns,
                                            st.session_state.input_content,
                                        )

                                        # Display chain results
                                        st.markdown("## Chain Execution Results")

                                        # Show sequence
                                        st.markdown("### Pattern Sequence")
                                        st.code(" → ".join(chain_results["sequence"]))

                                        # Show each stage
                                        st.markdown("### Execution Stages")
                                        for i, stage in enumerate(
                                            chain_results["stages"], 1
                                        ):
                                            with st.expander(
                                                f"Stage {i}: {stage['pattern']}",
                                                expanded=False,
                                            ):
                                                st.markdown("#### Input")
                                                st.code(stage["input"])
                                                st.markdown("#### Output")
                                                if stage["success"]:
                                                    st.markdown(stage["output"])
                                                else:
                                                    st.error(stage["error"])

                                        # Show final output
                                        if chain_results["metadata"]["success"]:
                                            st.markdown("### Final Output")
                                            st.markdown(chain_results["final_output"])
                                            st.session_state.chat_output.append(
                                                chain_results["final_output"]
                                            )
                                        else:
                                            st.error(
                                                "Chain execution failed. Check the stages above for details."
                                            )
                                    else:
                                        # Normal pattern execution
                                        outputs = execute_patterns(selected_patterns)
                                        st.session_state.chat_output.extend(outputs)

                    # Display outputs after execution
                    if st.session_state.chat_output:
                        st.markdown("---")
                        st.header("Pattern Outputs")
                        for message in st.session_state.chat_output:
                            st.markdown(message)
                            st.markdown("---")  # Add separator between outputs

                        # Output Actions
                        col1, col2 = st.columns(2)
                        with col1:
                            if st.button("📋 Copy All Outputs"):
                                all_outputs = "\n\n".join(st.session_state.chat_output)
                                success, error = set_clipboard_content(all_outputs)
                                if success:
                                    st.success("All outputs copied to clipboard!")
                                else:
                                    st.error(error)

                        with col2:
                            if st.button("❌ Clear Outputs"):
                                st.session_state.chat_output = []
                                st.success("Outputs cleared!")
                                st.experimental_rerun()

                    with col2:
                        st.write("")  # Empty space for layout balance

                else:
                    st.info("Select one or more patterns to run.")

            with tabs[1]:
                st.header("Output Analysis")
                if st.session_state.chat_output:
                    # Display pattern outputs in chronological order
                    for i, output in enumerate(
                        reversed(st.session_state.chat_output), 1
                    ):
                        with st.expander(f"Output #{i}", expanded=False):
                            st.markdown(output)
                else:
                    st.info("Run some patterns to see output analysis.")

        elif view == "Pattern Management":
            create_tab, edit_tab, delete_tab = st.tabs(["Create", "Edit", "Delete"])

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


if __name__ == "__main__":
    logger.info("Application startup")
    main()
