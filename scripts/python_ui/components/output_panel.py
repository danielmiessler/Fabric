"""
Output panel component for displaying pattern execution results with feedback.
"""
import streamlit as st
from typing import List, Dict, Any, Optional
from datetime import datetime
from utils.errors import ui_error_boundary
from utils.logging import logger
from services import storage

@ui_error_boundary
def render_output_panel(outputs: List[str]) -> None:
    """
    Render the output panel with results from pattern execution.
    
    Args:
        outputs: List of output strings from pattern execution
    """
    if not outputs:
        return
    
    st.markdown("---")
    
    # Output header with actions
    col1, col2, col3, col4 = st.columns([2, 1, 1, 1])
    with col1:
        st.header("📤 Pattern Outputs")
    with col2:
        if st.button("📋 Copy All", use_container_width=True):
            all_outputs = "\n\n".join(outputs)
            success, error = _set_clipboard_content(all_outputs)
            if success:
                st.toast("All outputs copied!", icon="📋")
            else:
                st.error(error)
    with col3:
        if st.button("💾 Save All", use_container_width=True):
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            all_content = "\n\n".join(outputs)
            _save_output_log("Batch Execution", "Multiple patterns", all_content, timestamp)
            st.toast("Outputs saved!", icon="💾")
    with col4:
        if st.button("🧹 Clear", use_container_width=True):
            if "chat_output" in st.session_state:
                st.session_state.chat_output = []
            st.toast("Outputs cleared!", icon="🧹")
            st.rerun()

    # Display outputs with enhanced formatting
    for i, message in enumerate(outputs):
        render_single_output(message, i)


@ui_error_boundary
def render_single_output(output: str, index: int) -> None:
    """
    Render a single output with metadata and actions.
    
    Args:
        output: The output string to display
        index: Index of the output for unique keys
    """
    with st.container():
        # Extract pattern name for better display
        pattern_name = "Output"
        if output.startswith("### 🎯"):
            pattern_name = output.split("\n")[0].replace("### 🎯 ", "")
        elif output.startswith("###"):
            pattern_name = output.split("\n")[0].replace("### ", "")

        # Output header with metadata
        col1, col2, col3 = st.columns([2, 1, 1])
        with col1:
            st.markdown(f"**Output {index+1}:** {pattern_name}")
        with col2:
            word_count = len(output.split())
            st.caption(f"📊 {word_count} words")
        with col3:
            if st.button("📋", key=f"copy_output_{index}", help="Copy this output"):
                success, error = _set_clipboard_content(output)
                if success:
                    st.toast("Copied!", icon="📋")
                else:
                    st.error(error)

        # Display the actual output
        st.markdown(output)

        # Add feedback and actions
        render_output_feedback(pattern_name, output, index)
        
        st.markdown("---")


@ui_error_boundary
def render_output_feedback(pattern_name: str, output: str, index: int) -> None:
    """
    Render feedback controls for an output.
    
    Args:
        pattern_name: Name of the pattern that generated this output
        output: The output content
        index: Index for unique keys
    """
    col1, col2, col3 = st.columns([1, 1, 2])
    
    with col1:
        try:
            feedback = st.feedback(
                "thumbs",
                key=f"output_feedback_{index}"
            )
            if feedback is not None:
                # Store feedback in session state
                if "pattern_feedback" not in st.session_state:
                    st.session_state.pattern_feedback = {}
                st.session_state.pattern_feedback[pattern_name] = feedback
                
                if feedback == 1:
                    st.toast("Thanks for the positive feedback!", icon="👍")
                else:
                    st.toast("Thanks for the feedback!", icon="👎")
        except Exception:
            # Fallback to simple buttons if feedback widget is not available
            col_like, col_dislike = st.columns(2)
            with col_like:
                if st.button("👍", key=f"like_{index}"):
                    if "pattern_feedback" not in st.session_state:
                        st.session_state.pattern_feedback = {}
                    st.session_state.pattern_feedback[pattern_name] = 1
                    st.success("Thanks!")
            with col_dislike:
                if st.button("👎", key=f"dislike_{index}"):
                    if "pattern_feedback" not in st.session_state:
                        st.session_state.pattern_feedback = {}
                    st.session_state.pattern_feedback[pattern_name] = 0
                    st.info("Thanks for feedback!")

    with col2:
        if st.button("⭐ Star", key=f"star_output_{index}"):
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            render_star_dialog(pattern_name, output, timestamp)


@ui_error_boundary
def render_star_dialog(pattern_name: str, output: str, timestamp: str) -> None:
    """
    Render a dialog for starring an output.
    
    Args:
        pattern_name: Name of the pattern
        output: The output content
        timestamp: Timestamp of the output
    """
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
                
                # Initialize starred outputs if not exists
                if "starred_outputs" not in st.session_state:
                    st.session_state.starred_outputs = []
                
                st.session_state.starred_outputs.append(log_entry)
                _save_outputs()
                st.success("Output starred successfully!")
                st.rerun()

        with col2:
            if st.button("Cancel", use_container_width=True):
                st.rerun()

    star_dialog()


@ui_error_boundary
def render_analysis_tab(outputs: List[str]) -> None:
    """
    Render the analysis tab with output metrics and search.
    
    Args:
        outputs: List of output strings to analyze
    """
    st.header("📊 Output Analysis")

    if outputs:
        # Enhanced analysis with metrics
        col1, col2, col3 = st.columns(3)
        with col1:
            st.metric("Total Outputs", len(outputs))
        with col2:
            total_chars = sum(len(output) for output in outputs)
            st.metric("Total Characters", f"{total_chars:,}")
        with col3:
            avg_length = total_chars / len(outputs)
            st.metric("Avg Length", f"{avg_length:.0f}")

        # Filter and search
        search_analysis = st.text_input("🔍 Search in outputs", placeholder="Search content...")

        # Display outputs with enhanced formatting
        filtered_outputs = [
            (i, output) for i, output in enumerate(reversed(outputs), 1)
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
                        success, error = _set_clipboard_content(output)
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


def _set_clipboard_content(content: str) -> tuple[bool, str]:
    """Set clipboard content with automatic dependency installation."""
    from services.dependencies import check_and_install_if_missing
    
    # Try to ensure pyperclip is available
    if not check_and_install_if_missing("pyperclip", "pyperclip>=1.8.0"):
        return False, "Pyperclip not available. Please install: pip install pyperclip"
    
    try:
        import pyperclip
        pyperclip.copy(content)
        return True, ""
    except Exception as e:
        return False, f"Failed to copy to clipboard: {e}"


def _save_output_log(pattern_name: str, input_content: str, output_content: str, timestamp: str) -> None:
    """Save output log. This should use the storage service."""
    try:
        log_entry = {
            "timestamp": timestamp,
            "pattern_name": pattern_name,
            "input": input_content,
            "output": output_content,
        }
        
        # Initialize output logs if not exists
        if "output_logs" not in st.session_state:
            st.session_state.output_logs = []
        
        st.session_state.output_logs.append(log_entry)
        _save_outputs()
        
    except Exception as e:
        logger.error(f"Failed to save output log: {e}")


def _save_outputs() -> None:
    """Save outputs to storage. This should use the storage service."""
    try:
        # This is a placeholder - should use storage service
import json
import os
import tempfile
        

        outputs_dir = os.path.expanduser("~/.config/fabric/outputs")
        os.makedirs(outputs_dir, exist_ok=True)

        def _atomic_write(path: str, data: Any) -> None:
            dir_name = os.path.dirname(path)
            fd, tmp_path = tempfile.mkstemp(dir=dir_name, prefix=".tmp_", suffix=".json")
            try:
                with os.fdopen(fd, "w") as f:
                    json.dump(data, f, indent=2)
                os.replace(tmp_path, path)
            finally:
                if os.path.exists(tmp_path):
                    try:
                        os.remove(tmp_path)
                    except Exception:
                        pass
        
        if "output_logs" in st.session_state:
            output_logs_file = os.path.join(outputs_dir, "output_logs.json")
            _atomic_write(output_logs_file, st.session_state.output_logs)
        
        if "starred_outputs" in st.session_state:
            starred_outputs_file = os.path.join(outputs_dir, "starred_outputs.json")
            _atomic_write(starred_outputs_file, st.session_state.starred_outputs)
    except Exception as e:
        logger.error(f"Failed to save outputs: {e}")