import streamlit as st
from utils.errors import ui_error_boundary
from services import patterns
from utils.typing import PatternSpec
from components import pattern_editor, pattern_list

@ui_error_boundary
def render() -> None:
    st.header("⚙️ Pattern Management")
    
    # Create tabs for different management operations
    create_tab, edit_tab, delete_tab, descriptions_tab = st.tabs(["Create", "Edit", "Delete", "📋 Descriptions"])
    
    with create_tab:
        _render_create_tab()
    
    with edit_tab:
        _render_edit_tab()
    
    with delete_tab:
        _render_delete_tab()
    
    with descriptions_tab:
        # Pattern descriptions management (matching original functionality)
                st.info("Pattern descriptions management not available")

@ui_error_boundary
def _render_create_tab() -> None:
    """Render the pattern creation interface."""
    st.header("Create New Pattern")
    
    pattern_name = st.text_input(
        "Pattern Name",
        placeholder="e.g., my_new_pattern",
        help="Name should contain only letters, numbers, dots, underscores, and hyphens"
    )
    
    if not pattern_name:
        st.info("Enter a pattern name to create a new pattern")
        return
    
    # Check if pattern already exists
    try:
        existing_patterns = patterns.list_patterns()
        existing_names = {p.name for p in existing_patterns}
        
        if pattern_name in existing_names:
            st.error(f"Pattern '{pattern_name}' already exists. Choose a different name.")
            return
    except Exception as e:
        st.error(f"Error checking existing patterns: {e}")
        return
    
    creation_mode = st.radio(
        "Creation Mode",
        ["Simple Editor", "Advanced (Wizard)"],
        key="creation_mode_main",
        horizontal=True,
    )
    
    if creation_mode == "Simple Editor":
        _render_simple_creation(pattern_name)
    else:
        _render_wizard_creation(pattern_name)

def _render_simple_creation(pattern_name: str) -> None:
    """Render the simple pattern creation interface."""
    default_content = """# IDENTITY and PURPOSE
You are an AI assistant designed to {purpose}.

# STEPS
- Step 1: Analyze the input
- Step 2: Process the information
- Step 3: Generate the output

# OUTPUT INSTRUCTIONS
- Provide clear and concise output
- Follow the specified format
- Be helpful and accurate
"""
    
    new_content = st.text_area(
        "Enter Pattern Content",
        value=default_content,
        height=400,
        help="Write your pattern content in Markdown format"
    )
    
    if st.button("Create Pattern", type="primary"):
        try:
            spec = PatternSpec(
                name=pattern_name,
                path=None,  # Will be set by save_pattern
                content=new_content
            )
            patterns.save_pattern(spec)
            st.success(f"Pattern '{pattern_name}' created successfully!")
            st.rerun()
        except Exception as e:
            st.error(f"Error creating pattern: {e}")

def _render_wizard_creation(pattern_name: str) -> None:
    """Render the wizard-based pattern creation interface."""
    st.info("Use the wizard to create a structured pattern step by step.")
    
    sections = ["IDENTITY", "GOAL", "STEPS", "OUTPUT INSTRUCTIONS"]
    current_section = st.radio(
        "Edit Section", 
        sections, 
        key="pattern_creation_section_select",
        horizontal=True
    )
    
    # Initialize session state for wizard sections
    for section in sections:
        key = f"new_pattern_{section.lower().replace(' ', '_')}"
        if key not in st.session_state:
            st.session_state[key] = ""
    
    # Section-specific editors
    if current_section == "IDENTITY":
        identity = st.text_area(
            "Define the IDENTITY and PURPOSE", 
            value=st.session_state.new_pattern_identity,
            height=200,
            help="Describe what this AI assistant is and its main purpose"
        )
        st.session_state.new_pattern_identity = identity
        
    elif current_section == "GOAL":
        goal = st.text_area(
            "Define the GOAL", 
            value=st.session_state.new_pattern_goal,
            height=200,
            help="Specify the specific goal or objective"
        )
        st.session_state.new_pattern_goal = goal
        
    elif current_section == "STEPS":
        steps = st.text_area(
            "Define the STEPS", 
            value=st.session_state.new_pattern_steps,
            height=200,
            help="List the steps the AI should follow"
        )
        st.session_state.new_pattern_steps = steps
        
    elif current_section == "OUTPUT INSTRUCTIONS":
        instructions = st.text_area(
            "Define the OUTPUT INSTRUCTIONS", 
            value=st.session_state.new_pattern_output_instructions,
            height=200,
            help="Specify how the output should be formatted"
        )
        st.session_state.new_pattern_output_instructions = instructions
    
    # Preview the complete pattern
    with st.expander("Preview Complete Pattern"):
        pattern_content = f"""# IDENTITY and PURPOSE
{st.session_state.get('new_pattern_identity', '')}

# GOAL
{st.session_state.get('new_pattern_goal', '')}

# STEPS
{st.session_state.get('new_pattern_steps', '')}

# OUTPUT INSTRUCTIONS
{st.session_state.get('new_pattern_output_instructions', '')}"""
        st.code(pattern_content, language="markdown")
    
    if st.button("Create Pattern", type="primary"):
        try:
            pattern_content = f"""# IDENTITY and PURPOSE
{st.session_state.get('new_pattern_identity', '')}

# GOAL
{st.session_state.get('new_pattern_goal', '')}

# STEPS
{st.session_state.get('new_pattern_steps', '')}

# OUTPUT INSTRUCTIONS
{st.session_state.get('new_pattern_output_instructions', '')}"""
            
            spec = PatternSpec(
                name=pattern_name,
                path=None,  # Will be set by save_pattern
                content=pattern_content
            )
            patterns.save_pattern(spec)
            st.success(f"Pattern '{pattern_name}' created successfully!")
            
            # Clear wizard state
            for section in sections:
                key = f"new_pattern_{section.lower().replace(' ', '_')}"
                if key in st.session_state:
                    del st.session_state[key]
            
            st.rerun()
        except Exception as e:
            st.error(f"Error creating pattern: {e}")

@ui_error_boundary
def _render_edit_tab() -> None:
    """Render the pattern editing interface."""
    st.header("Edit Patterns")
    
    try:
        available_patterns = patterns.list_patterns()
        if not available_patterns:
            st.warning("No patterns available. Create a pattern first.")
            return
        
        pattern_names = [p.name for p in available_patterns]
        selected_pattern = st.selectbox(
            "Select Pattern to Edit", 
            [""] + pattern_names,
            key="edit_pattern_selector"
        )
        
        if selected_pattern:
            pattern_editor.render(selected_pattern)
            
    except Exception as e:
        st.error(f"Error loading patterns: {e}")

@ui_error_boundary
def _render_delete_tab() -> None:
    """Render the pattern deletion interface."""
    st.header("Delete Patterns")
    
    try:
        available_patterns = patterns.list_patterns()
        if not available_patterns:
            st.warning("No patterns available.")
            return
        
        pattern_names = [p.name for p in available_patterns]
        patterns_to_delete = st.multiselect(
            "Select Patterns to Delete",
            pattern_names,
            key="delete_patterns_selector",
        )
        
        if patterns_to_delete:
            st.warning(f"You are about to delete {len(patterns_to_delete)} pattern(s):")
            for pattern in patterns_to_delete:
                st.markdown(f"- {pattern}")
            
            confirm_delete = st.checkbox(
                "I understand that this action cannot be undone",
                key="confirm_delete_checkbox"
            )
            
            if st.button(
                "🗑️ Delete Selected Patterns",
                type="primary",
                disabled=not confirm_delete,
            ):
                if confirm_delete:
                    try:
                        deleted_count = patterns.delete_patterns(patterns_to_delete)
                        if deleted_count > 0:
                            st.success(f"✓ Successfully deleted {deleted_count} pattern(s)")
                        else:
                            st.warning("No patterns were deleted")
                        st.rerun()
                    except Exception as e:
                        st.error(f"Error deleting patterns: {e}")
                else:
                    st.error("Please confirm deletion by checking the box above.")
        else:
            st.info("Select one or more patterns to delete.")
            
    except Exception as e:
        st.error(f"Error loading patterns: {e}")