"""
Real-time monitoring components for pattern execution dashboard.
Provides live execution progress, system health, and performance metrics.
"""
import streamlit as st
import time
from datetime import datetime, timedelta
from typing import Dict, Any, List

# Optional plotly imports for enhanced charts with dependency fallback
from services.dependencies import OptionalImport, check_and_install_if_missing

PLOTLY_AVAILABLE = check_and_install_if_missing("plotly", "plotly>=5.0.0")

if PLOTLY_AVAILABLE:
    try:
        import plotly.graph_objects as go
        import plotly.express as px
        from plotly.subplots import make_subplots
    except ImportError:
        PLOTLY_AVAILABLE = False
        go = px = None

from services.monitoring import get_execution_monitor, ExecutionStatus, ExecutionMetrics
from utils.logging import logger
from utils.errors import ui_error_boundary


@ui_error_boundary
def render_active_executions() -> None:
    """Render real-time active executions with progress bars."""
    monitor = get_execution_monitor()
    active_executions = monitor.get_active_executions()
    
    if not active_executions:
        st.info("🟢 No active pattern executions")
        return
    
    st.subheader("🚀 Active Executions")
    
    for execution in active_executions:
        with st.container():
            col1, col2, col3 = st.columns([3, 1, 1])
            
            with col1:
                # Progress bar with pattern name
                progress_text = f"**{execution.pattern}** ({execution.status.value})"
                if execution.estimated_completion:
                    eta = execution.estimated_completion.strftime("%H:%M:%S")
                    progress_text += f" - ETA: {eta}"
                
                st.write(progress_text)
                st.progress(execution.progress)
            
            with col2:
                # Provider/Model info
                if execution.provider and execution.model:
                    st.write(f"📡 {execution.provider}")
                    st.write(f"🤖 {execution.model}")
                else:
                    st.write("📡 Default")
            
            with col3:
                # Duration and controls
                if execution.start_time:
                    duration = datetime.now() - execution.start_time
                    st.write(f"⏱️ {duration.total_seconds():.1f}s")
                
                if st.button("❌", key=f"cancel_{execution.execution_id}", 
                           help="Cancel execution"):
                    monitor.cancel_execution(execution.execution_id)
                    st.rerun()
            
            st.divider()


@ui_error_boundary
def render_execution_stats() -> None:
    """Render real-time execution statistics."""
    monitor = get_execution_monitor()
    stats = monitor.get_execution_stats()
    
    st.subheader("📊 Execution Statistics")
    
    col1, col2, col3, col4 = st.columns(4)
    
    with col1:
        st.metric(
            label="Total Executions",
            value=stats["total_executions"],
            help="Total number of pattern executions"
        )
    
    with col2:
        success_rate = stats["success_rate"] * 100
        st.metric(
            label="Success Rate",
            value=f"{success_rate:.1f}%",
            delta=f"+{success_rate:.1f}%" if success_rate > 90 else None,
            help="Percentage of successful executions"
        )
    
    with col3:
        avg_duration = stats["average_duration"] / 1000  # Convert to seconds
        st.metric(
            label="Avg Duration",
            value=f"{avg_duration:.1f}s",
            help="Average execution time in seconds"
        )
    
    with col4:
        st.metric(
            label="Active Now",
            value=stats["active_count"],
            delta=f"+{stats['active_count']}" if stats["active_count"] > 0 else "0",
            help="Currently running executions"
        )


@ui_error_boundary
def render_system_health() -> None:
    """Render system health status panel."""
    monitor = get_execution_monitor()
    health = monitor.get_system_health()
    
    st.subheader("🏥 System Health")
    
    col1, col2 = st.columns(2)
    
    with col1:
        st.write("**Fabric CLI**")
        if health.fabric_cli_available:
            st.success("✅ Available")
            if health.fabric_version:
                st.write(f"Version: {health.fabric_version}")
        else:
            st.error("❌ Unavailable")
            st.write("Check fabric installation")
    
    with col2:
        st.write("**System Resources**")
        if health.system_load is not None:
            st.write(f"Load: {health.system_load:.2f}")
        if health.memory_usage_percent is not None:
            st.write(f"Memory: {health.memory_usage_percent:.1f}%")
    
    # Provider status
    if health.provider_status:
        st.write("**Provider Status**")
        provider_cols = st.columns(len(health.provider_status))
        
        for i, (provider, status) in enumerate(health.provider_status.items()):
            with provider_cols[i]:
                icon = "✅" if status else "❌"
                st.write(f"{icon} {provider.title()}")
    
    # Last health check
    if health.last_check:
        st.caption(f"Last checked: {health.last_check.strftime('%H:%M:%S')}")


@ui_error_boundary
def render_execution_timeline() -> None:
    """Render execution timeline chart."""
    monitor = get_execution_monitor()
    recent_executions = monitor.get_recent_executions(limit=20)
    
    if not recent_executions:
        st.info("No recent executions to display")
        return
    
    st.subheader("📈 Execution Timeline")
    
    # Prepare data for timeline chart
    timeline_data = []
    for execution in recent_executions:
        if execution.start_time and execution.end_time:
            timeline_data.append({
                'Pattern': execution.pattern,
                'Start': execution.start_time.strftime('%H:%M:%S'),
                'End': execution.end_time.strftime('%H:%M:%S'),
                'Duration': f"{(execution.end_time - execution.start_time).total_seconds():.1f}s",
                'Status': execution.status.value,
                'Provider': execution.provider or 'default'
            })
    
    if timeline_data:
        if PLOTLY_AVAILABLE:
            # Create Gantt chart for execution timeline
            fig = px.timeline(
                timeline_data,
                x_start="Start",
                x_end="End",
                y="Pattern",
                color="Status",
                hover_data=["Duration", "Provider"],
                title="Pattern Execution Timeline"
            )
            
            fig.update_layout(
                height=400,
                xaxis_title="Time",
                yaxis_title="Pattern",
                showlegend=True
            )
            
            st.plotly_chart(fig, use_container_width=True)
        else:
            # Fallback to table view
            st.dataframe(timeline_data, use_container_width=True)


@ui_error_boundary
def render_performance_trends() -> None:
    """Render pattern performance trends."""
    monitor = get_execution_monitor()
    recent_executions = monitor.get_recent_executions(limit=50)
    
    if len(recent_executions) < 2:
        st.info("Need more execution data for trend analysis")
        return
    
    st.subheader("📊 Performance Trends")
    
    # Group executions by pattern
    pattern_metrics = {}
    for execution in recent_executions:
        if execution.pattern not in pattern_metrics:
            pattern_metrics[execution.pattern] = {
                'durations': [],
                'success_count': 0,
                'total_count': 0,
                'timestamps': []
            }
        
        metrics = pattern_metrics[execution.pattern]
        metrics['total_count'] += 1
        
        if execution.status == ExecutionStatus.COMPLETED:
            metrics['success_count'] += 1
            if execution.duration_ms:
                metrics['durations'].append(execution.duration_ms / 1000)  # Convert to seconds
                metrics['timestamps'].append(execution.start_time)
    
    if pattern_metrics:
        # Create performance summary table
        performance_data = []
        for pattern, metrics in pattern_metrics.items():
            avg_duration = sum(metrics['durations']) / len(metrics['durations']) if metrics['durations'] else 0
            success_rate = metrics['success_count'] / metrics['total_count'] if metrics['total_count'] > 0 else 0
            
            performance_data.append({
                'Pattern': pattern,
                'Avg Duration (s)': f"{avg_duration:.2f}",
                'Success Rate': f"{success_rate * 100:.1f}%",
                'Total Runs': metrics['total_count']
            })
        
        st.dataframe(performance_data, use_container_width=True)
        
        # Duration trend chart for top patterns
        top_patterns = sorted(pattern_metrics.items(), 
                            key=lambda x: x[1]['total_count'], 
                            reverse=True)[:5]
        
        if top_patterns and PLOTLY_AVAILABLE:
            fig = go.Figure()
            
            for pattern, metrics in top_patterns:
                if metrics['durations'] and metrics['timestamps']:
                    fig.add_trace(go.Scatter(
                        x=metrics['timestamps'],
                        y=metrics['durations'],
                        mode='lines+markers',
                        name=pattern,
                        line=dict(width=2)
                    ))
            
            fig.update_layout(
                title="Pattern Duration Trends",
                xaxis_title="Time",
                yaxis_title="Duration (seconds)",
                height=400
            )
            
            st.plotly_chart(fig, use_container_width=True)
        elif top_patterns:
            st.info("📊 Duration trends chart available with plotly installation")


@ui_error_boundary
def render_monitoring_dashboard() -> None:
    """Render complete real-time monitoring dashboard."""
    st.header("🔍 Real-time Execution Monitor")
    
    # Auto-refresh controls
    col1, col2 = st.columns([3, 1])
    with col1:
        auto_refresh = st.checkbox("🔄 Auto-refresh", value=True)
    with col2:
        if st.button("🔄 Refresh Now"):
            st.rerun()
    
    # Auto-refresh logic (non-blocking)
    if auto_refresh:
        st.autorefresh(interval=2000, key="monitor_auto_refresh")
    
    # Dashboard sections
    render_active_executions()
    
    col1, col2 = st.columns(2)
    with col1:
        render_execution_stats()
    with col2:
        render_system_health()
    
    render_execution_timeline()
    render_performance_trends()


def start_monitoring_session() -> None:
    """Initialize monitoring session."""
    if 'monitoring_session' not in st.session_state:
        st.session_state.monitoring_session = True
        monitor = get_execution_monitor()
        logger.info("Real-time monitoring session started")


def get_monitoring_stats() -> Dict[str, Any]:
    """Get current monitoring statistics for sidebar display."""
    monitor = get_execution_monitor()
    stats = monitor.get_execution_stats()
    return {
        "active_executions": stats["active_count"],
        "total_executions": stats["total_executions"],
        "success_rate": stats["success_rate"]
    }