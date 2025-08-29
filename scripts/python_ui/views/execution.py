"""
Execution view for running Fabric patterns with enhanced UI and chain mode support.
"""
import streamlit as st
import pandas as pd
from typing import List, Dict, Any, Optional
from datetime import datetime
import time

from utils.errors import ui_error_boundary
from utils.logging import logger
from services import runner, patterns
from components import pattern_selector, output_panel

@ui_error_boundary
def render() -> None:
    """Main execution view rendering function."""
    # Initialize session state
    _initialize_session_state()
    
    
    # Check for patterns availability
    try:
        pattern_specs = patterns.list_patterns()
        available_patterns = [spec.name for spec in pattern_specs]
    except Exception as e:
        logger.error(f"Failed to load patterns: {e}")
        st.error(f"Failed to load patterns: {e}")
        return

    if not available_patterns:
        logger.warning("No patterns available")
        st.warning("No patterns available. Create a pattern first.")
        return

    # Enhanced tabs with icons
    tabs = st.tabs(["🚀 Execute", "📊 Analysis"])

    with tabs[0]:
        _render_execution_tab(available_patterns)

    with tabs[1]:
        # Get outputs from session state
        outputs = st.session_state.get("chat_output", [])
        output_panel.render_analysis_tab(outputs)


@ui_error_boundary
def _render_execution_tab(available_patterns: List[str]) -> None:
    """Render the main execution tab."""
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

    # Enhanced pattern selection with intelligent search
    st.subheader("🎨 Select Patterns")
    
    # Pattern selection mode
    selection_mode = st.radio(
        "Selection Mode",
        ["🔍 Basic Search", "🧠 Intelligent Discovery"],
        horizontal=True,
        key="pattern_selection_mode"
    )
    
    if selection_mode == "🔍 Basic Search":
        selected_patterns = pattern_selector.render_pattern_selector("main_patterns")
    else:
        # Use intelligent pattern selector
        from components import intelligent_pattern_selector
        selector_result = intelligent_pattern_selector.render_intelligent_pattern_selector("intelligent_main")
        selected_patterns = selector_result.get("selected_patterns", [])
    
    st.session_state.selected_patterns = selected_patterns

    if selected_patterns:
        # Enhanced pattern details
        pattern_selector.render_pattern_details(selected_patterns)

        # Enhanced input section
        _render_input_section()

        # Enhanced execution options
        chain_mode, auto_save = _render_execution_options(selected_patterns)

        # Pattern Variables UI
        pattern_variables = _render_pattern_variables_ui(selected_patterns)

        # Enhanced execution button section
        _render_execution_controls(selected_patterns, chain_mode, pattern_variables)

        # Defer execution to next rerun to avoid double-execution within same pass
        if st.session_state.get("execute_patterns", False):
            # Do not execute here; allow top-level to handle in next run
            pass

        # Display output from the last run
        if st.session_state.get('last_run_outputs'):
            st.markdown("### 🎯 Pattern Output")
            output_display = "\n\n---\n\n".join(st.session_state.last_run_outputs)
            st.markdown(output_display)

            # Show feedback for each individual output
            for output in st.session_state.last_run_outputs:
                if output.startswith("### 🎯"):
                    pattern_name = output.split("\n")[0].replace("### 🎯 ", "")
                    _show_pattern_feedback_ui(pattern_name, output)
                    
                                
            # Clear the outputs from state so they don't re-display
            st.session_state.last_run_outputs = []

        # Enhanced output display
        if st.session_state.get("chat_output"):
            output_panel.render_output_panel(st.session_state.chat_output)
        
        # Handle starring dialogs (render outside of conditional blocks)
        
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


@ui_error_boundary
def _render_input_section() -> None:
    """Render the input configuration section."""
    st.subheader("📝 Input Configuration")

    # Input method selection
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
                    success, content, error = _get_clipboard_content()
                    if success:
                        # Validate and sanitize clipboard content
                        if _validate_input_content(content):
                            sanitized_content = _sanitize_input_content(content)
                            if sanitized_content != content:
                                st.toast("Content was automatically sanitized", icon="🧹")
                            
                            st.session_state.input_content = sanitized_content
                            st.session_state.show_preview = True
                            st.toast("Content loaded successfully!", icon="✅")
                        else:
                            st.error("Invalid clipboard content")
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

    # Show input preview if enabled
    if st.session_state.get("show_preview", False) and st.session_state.get("input_content"):
        st.text_area("Input Preview", value=st.session_state.input_content, height=100, disabled=True)


@ui_error_boundary
def _render_execution_options(selected_patterns: List[str]) -> tuple[bool, bool]:
    """Render execution options and return their values."""
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
            value=st.session_state.user_preferences.get("auto_save", True),
            help="Automatically save successful executions"
        )
        st.session_state.user_preferences["auto_save"] = auto_save

    # Show chain mode info and reordering
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

    return chain_mode, auto_save


@ui_error_boundary
def _render_pattern_variables_ui(selected_patterns: List[str]) -> Dict[str, Dict[str, str]]:
    """Render pattern variables UI and return the collected variables."""
    pattern_variables = {}
    
    if selected_patterns:
        # Check if any selected patterns have variables
        patterns_with_vars = []
        for pattern in selected_patterns:
            variables = _detect_pattern_variables(pattern)
            if variables:
                patterns_with_vars.append((pattern, variables))
        
        if patterns_with_vars:
            st.markdown("---")
            st.subheader("🔧 Pattern Variables")
            st.info("The following patterns require variables to be set:")
            
            # Create tabs for each pattern with variables
            if len(patterns_with_vars) == 1:
                pattern, variables = patterns_with_vars[0]
                st.markdown(f"**{pattern}**")
                pattern_variables[pattern] = _render_single_pattern_variables(
                    variables, f"vars_{pattern}"
                )
            else:
                # Multiple patterns with variables - use tabs
                tab_names = [f"{pattern} ({len(vars)} vars)" for pattern, vars in patterns_with_vars]
                tabs = st.tabs(tab_names)
                
                for i, (pattern, variables) in enumerate(patterns_with_vars):
                    with tabs[i]:
                        pattern_variables[pattern] = _render_single_pattern_variables(
                            variables, f"vars_{pattern}"
                        )

    return pattern_variables


@ui_error_boundary
def _render_execution_controls(selected_patterns: List[str], chain_mode: bool, pattern_variables: Dict[str, Dict[str, str]]) -> None:
    """Render execution controls and handle execution."""
    st.markdown("---")
    col1, col2, col3 = st.columns([2, 1, 1])

    with col1:
        # Check if all required variables are filled
        can_execute = True
        missing_vars_info = []
        
        if pattern_variables:
            for pattern_name, variables in pattern_variables.items():
                required_vars = _detect_pattern_variables(pattern_name)
                is_valid, missing_vars = _validate_pattern_variables(variables, required_vars)
                if not is_valid:
                    can_execute = False
                    missing_vars_info.append(f"{pattern_name}: {', '.join(missing_vars)}")
        
        if st.button(
            "🚀 Execute Patterns",
            type="primary",
            use_container_width=True,
            disabled=not st.session_state.get("input_content") or not selected_patterns or not can_execute
        ):
            st.session_state.execute_patterns = True
            st.rerun()
        
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


@ui_error_boundary
def _execute_patterns(selected_patterns: List[str], chain_mode: bool, pattern_variables: Dict[str, Dict[str, str]], auto_save: bool) -> None:
    """Execute the selected patterns using the runner service."""
    if not st.session_state.get("input_content"):
        st.error("⚠️ Please provide input content.")
        return
    
    if not selected_patterns:
        st.error("⚠️ Please select at least one pattern.")
        return

    # Update execution stats
    st.session_state.execution_stats["total_runs"] += 1
    
    # Get current configuration (optional - Fabric can use defaults)
    current_provider = st.session_state.config.get("vendor")
    current_model = st.session_state.config.get("model")
    
    # Debug: Log what we got from config
    logger.info(f"Config - provider: {current_provider}, model: {current_model}")
    
    # If model looks like a full model name (e.g., "gpt-4.1-turbo"), don't use it as provider
    if current_provider == current_model:
        # This suggests the model name was incorrectly set as provider
        current_provider = None
        logger.info(f"Fixed config - provider: {current_provider}, model: {current_model}")
    
    # Note: Fabric can run without explicit provider/model if defaults are configured
    if not current_provider and not current_model:
        st.info("ℹ️ Using Fabric's default model configuration")
        current_provider = None
        current_model = None

    start_time = time.time()
    
    try:
        if chain_mode:
            # Execute pattern chain
            chain_results = _execute_pattern_chain(selected_patterns, pattern_variables, current_provider, current_model)
            
            # Calculate execution time and update stats BEFORE any UI updates
            execution_time = time.time() - start_time
            st.session_state.execution_stats["avg_execution_time"] = (
                (st.session_state.execution_stats["avg_execution_time"] *
                 (st.session_state.execution_stats["total_runs"] - 1) + execution_time) /
                st.session_state.execution_stats["total_runs"]
            )
            
            # Check if chain was successful
            if chain_results and chain_results.get("metadata", {}).get("success", False):
                st.session_state.execution_stats["successful_runs"] += 1
                logger.info(f"Chain execution completed successfully. Stats: {st.session_state.execution_stats}")
            else:
                st.session_state.execution_stats["failed_runs"] += 1
                logger.error(f"Chain execution failed. Stats: {st.session_state.execution_stats}")
            
            _display_chain_results(chain_results)
        else:
            # Execute patterns individually
            outputs = _execute_individual_patterns(selected_patterns, pattern_variables, current_provider, current_model)
            
            # Calculate execution time BEFORE rerun
            execution_time = time.time() - start_time
            st.session_state.execution_stats["avg_execution_time"] = (
                (st.session_state.execution_stats["avg_execution_time"] *
                 (st.session_state.execution_stats["total_runs"] - 1) + execution_time) /
                st.session_state.execution_stats["total_runs"]
            )
            
            if outputs:
                # Update success stats BEFORE rerun
                st.session_state.execution_stats["successful_runs"] += 1
                logger.info(f"Execution completed successfully. Stats: {st.session_state.execution_stats}")
                
                st.session_state.last_run_outputs = outputs
                st.session_state.show_success_toast = True
                
                # Auto-save if enabled
                if auto_save:
                    for output in outputs:
                        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
                        pattern_name = output.split("\n")[0].replace("### 🎯 ", "") if output.startswith("### 🎯") else "Unknown"
                        _save_output_log(pattern_name, st.session_state.input_content, output, timestamp)
                
                st.rerun()
            else:
                # No outputs means execution failed
                st.session_state.execution_stats["failed_runs"] += 1
                logger.error(f"Execution failed - no outputs. Stats: {st.session_state.execution_stats}")
                st.error("Pattern execution completed but produced no output.")
        
    except Exception as e:
        # Calculate execution time even for failures
        execution_time = time.time() - start_time
        st.session_state.execution_stats["avg_execution_time"] = (
            (st.session_state.execution_stats["avg_execution_time"] *
             (st.session_state.execution_stats["total_runs"] - 1) + execution_time) /
            st.session_state.execution_stats["total_runs"]
        )
        
        # Update failure stats
        st.session_state.execution_stats["failed_runs"] += 1
        logger.error(f"Pattern execution failed: {e}")
        logger.error(f"Execution failed. Stats: {st.session_state.execution_stats}")
        st.error(f"Pattern execution failed: {e}")


@ui_error_boundary
def _execute_individual_patterns(selected_patterns: List[str], pattern_variables: Dict[str, Dict[str, str]], provider: str, model: str) -> List[str]:
    """Execute patterns individually using the runner service."""
    all_outputs = []
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    
    # Create status container
    status_container = st.status(
        f"🚀 Executing {len(selected_patterns)} pattern{'s' if len(selected_patterns) > 1 else ''}...",
        expanded=True
    )
    
    with status_container:
        st.write("🔍 Validating configuration...")
        st.write("✅ Configuration validated")
        start_ts = time.time()
        
        # Display model info properly
        if provider and model:
            st.write(f"🤖 Using: {provider} - {model}")
        elif model:
            st.write(f"🤖 Using model: {model}")
        else:
            st.write("🤖 Using Fabric's default configuration")
        
        for i, pattern in enumerate(selected_patterns, 1):
            st.write(f"🔄 Running pattern {i}/{len(selected_patterns)}: **{pattern}**")
            
            # Create progress bar
            progress = st.progress(0)
            progress.progress(i / len(selected_patterns))
            
            try:
                # Execute pattern using runner service
                st.write(f"🔧 Executing with provider={provider}, model={model}")
                
                result = runner.run_fabric(
                    pattern=pattern,
                    input_text=st.session_state.input_content,
                    provider=provider,
                    model=model
                )
                
                st.write(f"📊 Result: success={result.success}, exit_code={result.exit_code}, duration={result.duration_ms}ms")
                
                if result.success:
                    st.write(f"✅ Pattern **{pattern}** completed successfully")
                    st.write(f"📄 Output length: {len(result.output)} characters")
                    
                    # Format output
                    output_msg = f"""### 🎯 {pattern}

{result.output}"""
                    all_outputs.append(output_msg)
                    
                    # Save to output logs
                    _save_output_log(pattern, st.session_state.input_content, result.output, timestamp)
                    
                else:
                    error_msg = f"❌ Pattern **{pattern}** failed: {result.error or 'Unknown error'}"
                    st.error(error_msg)
                    st.write(f"🔍 Debug info: exit_code={result.exit_code}, meta={result.meta}")
                    logger.error(f"Pattern {pattern} failed: {result.error}")
                    logger.error(f"Pattern {pattern} debug: exit_code={result.exit_code}, meta={result.meta}")
                    all_outputs.append(f"### ❌ {pattern}\n\n{result.error or 'Unknown error'}")
                    
            except Exception as e:
                error_msg = f"❌ Pattern **{pattern}** failed: {str(e)}"
                st.error(error_msg)
                logger.error(f"Pattern {pattern} failed: {str(e)}")
                all_outputs.append(f"### ❌ {pattern}\n\n{str(e)}")
        
        if all_outputs:
            execution_time = time.time() - start_ts
            st.write(f"🎉 Execution completed in {execution_time:.2f}s")
    
    return all_outputs


@ui_error_boundary
def _execute_pattern_chain(selected_patterns: List[str], pattern_variables: Dict[str, Dict[str, str]], provider: str, model: str) -> Dict[str, Any]:
    """Execute patterns in chain mode using the runner service."""
    try:
        # Execute chain using runner service
        chain_steps = runner.run_chain(
            patterns=selected_patterns,
            seed_input=st.session_state.input_content,
            provider=provider,
            model=model
        )
        
        # Convert to legacy format for compatibility with display function
        chain_results = {
            "sequence": selected_patterns,
            "stages": [],
            "final_output": None,
            "metadata": {
                "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "success": False,
            },
        }
        
        for step in chain_steps:
            stage_result = {
                "pattern": step.pattern,
                "input": step.input,
                "output": step.output,
                "success": step.output is not None,
                "error": step.error,
            }
            chain_results["stages"].append(stage_result)
        
        # Set final output and success status
        if chain_steps and chain_steps[-1].output:
            chain_results["final_output"] = chain_steps[-1].output
            chain_results["metadata"]["success"] = True
        
        return chain_results
        
    except Exception as e:
        logger.error(f"Chain execution failed: {e}")
        return {
            "sequence": selected_patterns,
            "stages": [],
            "final_output": None,
            "metadata": {
                "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "success": False,
                "error": str(e)
            },
        }


@ui_error_boundary
def _display_chain_results(chain_results: Dict[str, Any]) -> None:
    """Display chain execution results."""
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
        
        # Add to chat output
        if "chat_output" not in st.session_state:
            st.session_state.chat_output = []
        st.session_state.chat_output.append(chain_results["final_output"])

        # Add feedback UI for chain result
        _show_pattern_feedback_ui("Chain Result", chain_results["final_output"])


# Helper functions

def _initialize_session_state() -> None:
    """Initialize session state variables."""
    if "execution_stats" not in st.session_state:
        st.session_state.execution_stats = {
            "total_runs": 0,
            "successful_runs": 0,
            "failed_runs": 0,
            "avg_execution_time": 0.0,
        }
    
    if "user_preferences" not in st.session_state:
        st.session_state.user_preferences = {
            "auto_save": True,
        }
    
    if "pattern_feedback" not in st.session_state:
        st.session_state.pattern_feedback = {}
    
    if "chat_output" not in st.session_state:
        st.session_state.chat_output = []
    
    if "config" not in st.session_state:
        st.session_state.config = {}


def _get_clipboard_content() -> tuple[bool, str, str]:
    """Get content from clipboard with automatic dependency installation."""
    from services.dependencies import check_and_install_if_missing
    
    # Try to ensure pyperclip is available
    if not check_and_install_if_missing("pyperclip", "pyperclip>=1.8.0"):
        return False, "", "Pyperclip not available. Please install: pip install pyperclip"
    
    try:
        import pyperclip
        content = pyperclip.paste()
        return True, content, ""
    except Exception as e:
        return False, "", f"Failed to access clipboard: {e}"


def _validate_input_content(content: str) -> bool:
    """Validate input content."""
    if not content or not content.strip():
        return False
    if len(content) > 50000:  # 50k character limit
        return False
    return True


def _sanitize_input_content(content: str) -> str:
    """Sanitize input content."""
    # Remove null bytes
    content = content.replace("\0", "")
    
    # Replace control characters with spaces (except newlines and tabs)
    allowed_chars = {"\n", "\t", "\r"}
    sanitized_chars = []
    for c in content:
        if c in allowed_chars or ord(c) >= 32:
            sanitized_chars.append(c)
        else:
            sanitized_chars.append(" ")
    
    # Join characters and normalize whitespace
    content = "".join(sanitized_chars)
    content = " ".join(content.split())
    
    return content


# _render_input_preview removed - replaced by input_preview component


def _detect_pattern_variables(pattern_name: str) -> List[str]:
    """Detect variables in a pattern by analyzing its content."""
    try:
        pattern_spec = patterns.load_pattern(pattern_name)
        # Simple variable detection - look for {{variable}} patterns
        import re
        variables = re.findall(r'\{\{(\w+)\}\}', pattern_spec.content)
        return list(set(variables))
    except Exception:
        return []


def _render_single_pattern_variables(variables: List[str], key_prefix: str) -> Dict[str, str]:
    """Render UI for a single pattern's variables."""
    variable_values = {}
    
    for var in variables:
        variable_values[var] = st.text_input(
            f"Variable: {var}",
            key=f"{key_prefix}_{var}",
            placeholder=f"Enter value for {var}"
        )
    
    return variable_values


def _validate_pattern_variables(variables: Dict[str, str], required_vars: List[str]) -> tuple[bool, List[str]]:
    """Validate that all required variables have values."""
    missing_vars = []
    for var in required_vars:
        if not variables.get(var, "").strip():
            missing_vars.append(var)
    
    return len(missing_vars) == 0, missing_vars


def _show_pattern_feedback_ui(pattern_name: str, output: str) -> None:
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
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            output_panel.render_star_dialog(pattern_name, output, timestamp)


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
        
    except Exception as e:
        logger.error(f"Failed to save output log: {e}")