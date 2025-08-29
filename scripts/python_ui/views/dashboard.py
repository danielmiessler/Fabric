import streamlit as st
from utils.errors import ui_error_boundary
from components import stats_panel, history_table, real_time_monitor

@ui_error_boundary
def render() -> None:
    st.header("📊 Real-time Pattern Execution Dashboard")
    
    # Welcome message and overview
    st.markdown("""
    Welcome to the Enhanced Dashboard! Monitor real-time pattern executions, system health,
    and comprehensive analytics for your AI workflow orchestration.
    """)
    
    # Initialize monitoring session
    real_time_monitor.start_monitoring_session()
    
    # Auto-refresh controls
    col1, col2, col3 = st.columns([2, 1, 1])
    with col1:
        auto_refresh = st.checkbox("🔄 Auto-refresh (2s)", value=False, key="dashboard_auto_refresh")
    with col2:
        if st.button("🔄 Refresh Now"):
            st.rerun()
    with col3:
        view_mode = st.selectbox("View", ["Enhanced", "Classic"], index=0)
    
    # Auto-refresh logic using proper Streamlit pattern
    if auto_refresh:
        # Use a placeholder to show refresh timing
        refresh_placeholder = st.empty()
        refresh_placeholder.info("🔄 Auto-refresh enabled - Dashboard will update automatically")
        
        # Use time tracking in session state to prevent infinite loops
        current_time = st.session_state.get("last_dashboard_refresh", 0)
        import time
        now = time.time()
        
        # Only rerun if 2 seconds have passed since last refresh
        if now - current_time >= 2:
            st.session_state["last_dashboard_refresh"] = now
            st.rerun()
    
    if view_mode == "Enhanced":
        # Enhanced dashboard with real-time monitoring
        
        # Real-time monitoring section
        real_time_monitor.render_active_executions()
        
        # System overview
        col1, col2 = st.columns(2)
        with col1:
            real_time_monitor.render_execution_stats()
        with col2:
            real_time_monitor.render_system_health()
        
        st.markdown("---")
        
        # Performance analytics
        real_time_monitor.render_execution_timeline()
        real_time_monitor.render_performance_trends()
        
        st.markdown("---")
        
        # Classic components in tabs
        tab1, tab2, tab3, tab4 = st.tabs(["📊 Legacy Stats", "📋 History & Outputs", "⭐ Starred Outputs", "⚙️ Controls"])
        
        with tab1:
            render_classic_stats()
        
        with tab2:
            history_table.render_history_overview()
        
        with tab3:
                        st.info("⭐ Starred outputs feature not available")
        
        with tab4:
            render_dashboard_controls()
    
    else:
        # Classic dashboard view
        render_classic_overview()
        
        # Main dashboard content with tabs
        tab1, tab2, tab3 = st.tabs(["📊 Statistics", "📋 History & Outputs", "⭐ Starred Outputs"])
        
        with tab1:
            render_classic_stats()
        
        with tab2:
            history_table.render_history_overview()
        
        with tab3:
                        st.info("⭐ Starred outputs feature not available")
        
        render_dashboard_controls()


def render_classic_overview():
    """Render the classic dashboard overview metrics."""
    st.markdown("### 📈 Quick Overview")
    col1, col2, col3, col4 = st.columns(4)
    
    # Get basic stats for quick overview
    execution_stats = st.session_state.get("execution_stats", {})
    output_logs = st.session_state.get("output_logs", [])
    starred_outputs = st.session_state.get("starred_outputs", [])
    
    with col1:
        total_runs = execution_stats.get("total_runs", 0)
        st.metric("Total Executions", total_runs)
    
    with col2:
        success_rate = 0
        if total_runs > 0:
            successful = execution_stats.get("successful_runs", 0)
            success_rate = (successful / total_runs) * 100
        st.metric("Success Rate", f"{success_rate:.1f}%")
    
    with col3:
        st.metric("Saved Outputs", len(output_logs))
    
    with col4:
        st.metric("Starred Items", len(starred_outputs))
    
    st.markdown("---")


def render_classic_stats():
    """Render classic statistics panels."""
    stats_panel.render_storage_stats()
    st.markdown("---")
    stats_panel.render_pattern_stats()


def render_dashboard_controls():
    """Render dashboard action controls."""
    st.markdown("### ⚙️ Dashboard Actions")
    
    output_logs = st.session_state.get("output_logs", [])
    starred_outputs = st.session_state.get("starred_outputs", [])
    
    col1, col2, col3 = st.columns(3)
    
    with col1:
        if st.button("🔄 Refresh Data", use_container_width=True):
            # Force refresh by clearing any cached data
            st.cache_data.clear()
            st.success("Dashboard data refreshed!")
            st.rerun()
    
    with col2:
        if st.button("📊 Export Stats", use_container_width=True):
            # Show stats in a format that can be copied
            stats = st.session_state.get("execution_stats", {})
            stats_text = f"""
Fabric Pattern Studio Statistics
================================
Total Executions: {stats.get('total_runs', 0)}
Successful Runs: {stats.get('successful_runs', 0)}
Failed Runs: {stats.get('failed_runs', 0)}
Average Time: {stats.get('avg_execution_time', 0):.2f}s
Saved Outputs: {len(output_logs)}
Starred Items: {len(starred_outputs)}
            """.strip()
            st.code(stats_text, language="text")
            st.info("Copy the statistics above to save or share them.")
    
    with col3:
        if st.button("🗑️ Clear Session Data", use_container_width=True, type="secondary"):
            st.session_state["show_clear_warning"] = True
    
    # Clear data confirmation
    if st.session_state.get("show_clear_warning", False):
        st.warning("⚠️ This will clear all execution statistics and session data (but not saved outputs).")
        col1, col2 = st.columns(2)
        
        with col1:
            if st.button("⚠️ Confirm Clear", type="primary"):
                # Clear session statistics but preserve saved outputs
                st.session_state.execution_stats = {
                    "total_runs": 0,
                    "successful_runs": 0,
                    "failed_runs": 0,
                    "avg_execution_time": 0
                }
                st.session_state["show_clear_warning"] = False
                st.success("Session statistics cleared!")
                st.rerun()
        
        with col2:
            if st.button("❌ Cancel"):
                st.session_state["show_clear_warning"] = False
                st.rerun()