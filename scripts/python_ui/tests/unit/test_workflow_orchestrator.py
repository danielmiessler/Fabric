"""
Unit tests for workflow orchestrator service.
"""
import pytest
from unittest.mock import Mock, patch
from datetime import datetime, timedelta

from services.workflow_orchestrator import (
    WorkflowOrchestrator, WorkflowStep, Workflow, ExecutionConfig, WorkflowStatus,
    get_workflow_orchestrator, create_simple_workflow
)
from utils.typing import RunResult


@pytest.fixture
def orchestrator():
    """Create a fresh WorkflowOrchestrator for testing."""
    return WorkflowOrchestrator()


@pytest.fixture
def sample_config():
    """Sample execution configuration."""
    return ExecutionConfig(
        provider="test_provider",
        model="test_model",
        timeout_s=60,
        max_retries=2
    )


@pytest.fixture
def sample_steps():
    """Sample workflow steps."""
    return [
        WorkflowStep(id="step_1", pattern="pattern_a"),
        WorkflowStep(id="step_2", pattern="pattern_b", depends_on=["step_1"]),
        WorkflowStep(id="step_3", pattern="pattern_c", depends_on=["step_1"])
    ]


def test_create_workflow(orchestrator, sample_config):
    """Test basic workflow creation."""
    patterns = ["pattern_a", "pattern_b", "pattern_c"]
    workflow = orchestrator.create_workflow("test_workflow", patterns, sample_config)
    
    assert workflow.name == "test_workflow"
    assert len(workflow.steps) == 3
    assert workflow.config == sample_config
    assert workflow.status == WorkflowStatus.PENDING
    
    # Check sequential dependencies
    assert workflow.steps[0].depends_on == []
    assert workflow.steps[1].depends_on == ["step_0"]
    assert workflow.steps[2].depends_on == ["step_1"]


def test_create_parallel_workflow(orchestrator):
    """Test parallel workflow creation."""
    parallel_groups = {
        "group_1": ["pattern_a", "pattern_b"],
        "group_2": ["pattern_c"]
    }
    sequential_patterns = ["pattern_d"]
    
    workflow = orchestrator.create_parallel_workflow(
        "parallel_test",
        parallel_groups,
        sequential_patterns
    )
    
    assert len(workflow.steps) == 4  # 3 parallel + 1 sequential
    
    # Check parallel groups
    parallel_steps = [step for step in workflow.steps if step.parallel_group]
    assert len(parallel_steps) == 3
    
    # Check sequential step depends on all parallel steps
    sequential_step = [step for step in workflow.steps if not step.parallel_group][0]
    assert len(sequential_step.depends_on) == 3


def test_create_custom_workflow(orchestrator, sample_steps):
    """Test custom workflow creation."""
    workflow = orchestrator.create_custom_workflow("custom_test", sample_steps)
    
    assert workflow.name == "custom_test"
    assert len(workflow.steps) == 3
    assert workflow.steps == sample_steps


def test_create_custom_workflow_invalid_dependencies(orchestrator):
    """Test custom workflow creation with invalid dependencies."""
    invalid_steps = [
        WorkflowStep(id="step_1", pattern="pattern_a", depends_on=["nonexistent"])
    ]
    
    with pytest.raises(ValueError, match="depends on non-existent step"):
        orchestrator.create_custom_workflow("invalid_test", invalid_steps)


def test_workflow_step_methods():
    """Test workflow step helper methods."""
    steps = [
        WorkflowStep(id="step_1", pattern="pattern_a", status=WorkflowStatus.COMPLETED),
        WorkflowStep(id="step_2", pattern="pattern_b", status=WorkflowStatus.RUNNING),
        WorkflowStep(id="step_3", pattern="pattern_c", status=WorkflowStatus.PENDING)
    ]
    
    workflow = Workflow(id="test", name="test", steps=steps)
    
    # Test get_step
    step = workflow.get_step("step_1")
    assert step is not None
    assert step.pattern == "pattern_a"
    
    # Test get_steps_by_status
    pending_steps = workflow.get_steps_by_status(WorkflowStatus.PENDING)
    assert len(pending_steps) == 1
    assert pending_steps[0].id == "step_3"
    
    # Test get_ready_steps (no dependencies)
    ready_steps = workflow.get_ready_steps()
    assert len(ready_steps) == 1  # Only step_3 is pending and has no deps


def test_plan_execution_groups(orchestrator):
    """Test execution group planning."""
    steps = [
        WorkflowStep(id="step_1", pattern="pattern_a"),  # No deps
        WorkflowStep(id="step_2", pattern="pattern_b"),  # No deps  
        WorkflowStep(id="step_3", pattern="pattern_c", depends_on=["step_1", "step_2"])  # Depends on 1,2
    ]
    
    workflow = Workflow(id="test", name="test", steps=steps)
    
    groups = orchestrator._plan_execution_groups(workflow)
    
    # Should have 2 groups: [step_1, step_2] then [step_3]
    assert len(groups) == 2
    assert len(groups[0]) == 2  # step_1 and step_2 can run in parallel
    assert len(groups[1]) == 1  # step_3 runs after


def test_resolve_step_input(orchestrator):
    """Test step input resolution."""
    step_outputs = {
        "_initial": "initial input",
        "step_1": "output from step 1",
        "step_2": "output from step 2"
    }
    
    # No dependencies - should use initial
    step_no_deps = WorkflowStep(id="step_a", pattern="pattern_a")
    input_text = orchestrator._resolve_step_input(step_no_deps, step_outputs)
    assert input_text == "initial input"
    
    # Single dependency
    step_single_dep = WorkflowStep(id="step_b", pattern="pattern_b", depends_on=["step_1"])
    input_text = orchestrator._resolve_step_input(step_single_dep, step_outputs)
    assert input_text == "output from step 1"
    
    # Multiple dependencies - should concatenate
    step_multi_dep = WorkflowStep(id="step_c", pattern="pattern_c", depends_on=["step_1", "step_2"])
    input_text = orchestrator._resolve_step_input(step_multi_dep, step_outputs)
    assert "output from step 1" in input_text
    assert "output from step 2" in input_text
    assert "---" in input_text  # Separator


@patch('services.runner.run_fabric')
def test_execute_workflow_success(mock_run_fabric, orchestrator):
    """Test successful workflow execution."""
    # Mock successful execution
    mock_run_fabric.return_value = RunResult(
        success=True,
        output="test output",
        error=None,
        duration_ms=1000,
        exit_code=0
    )
    
    workflow = orchestrator.create_workflow("test", ["pattern_a"])
    result = orchestrator.execute_workflow(workflow, "test input")
    
    assert result["success"] is True
    assert result["final_output"] == "test output"
    assert workflow.status == WorkflowStatus.COMPLETED


@patch('services.runner.run_fabric')
def test_execute_workflow_failure(mock_run_fabric, orchestrator):
    """Test workflow execution with failure."""
    # Mock failed execution
    mock_run_fabric.return_value = RunResult(
        success=False,
        output="",
        error="execution failed",
        duration_ms=500,
        exit_code=1
    )
    
    workflow = orchestrator.create_workflow("test", ["pattern_a"])
    result = orchestrator.execute_workflow(workflow, "test input")
    
    assert result["success"] is False
    assert "execution failed" in str(result.get("error", ""))
    assert workflow.status == WorkflowStatus.FAILED


def test_validate_workflow_valid(orchestrator):
    """Test workflow validation with valid workflow."""
    steps = [
        WorkflowStep(id="step_1", pattern="pattern_a"),
        WorkflowStep(id="step_2", pattern="pattern_b", depends_on=["step_1"])
    ]
    workflow = Workflow(id="test", name="test", steps=steps)
    
    # Mock pattern availability
    with patch('services.patterns.list_patterns') as mock_patterns:
        mock_patterns.return_value = [
            Mock(name="pattern_a"),
            Mock(name="pattern_b")
        ]
        
        errors = orchestrator.validate_workflow(workflow)
        assert errors == []


def test_validate_workflow_invalid(orchestrator):
    """Test workflow validation with invalid workflow."""
    # Empty workflow
    empty_workflow = Workflow(id="test", name="test", steps=[])
    errors = orchestrator.validate_workflow(empty_workflow)
    assert "no steps" in errors[0].lower()
    
    # Duplicate step IDs
    duplicate_steps = [
        WorkflowStep(id="step_1", pattern="pattern_a"),
        WorkflowStep(id="step_1", pattern="pattern_b")  # Duplicate ID
    ]
    dup_workflow = Workflow(id="test", name="test", steps=duplicate_steps)
    errors = orchestrator.validate_workflow(dup_workflow)
    assert any("duplicate" in error.lower() for error in errors)
    
    # Invalid dependencies
    invalid_deps = [
        WorkflowStep(id="step_1", pattern="pattern_a", depends_on=["nonexistent"])
    ]
    invalid_workflow = Workflow(id="test", name="test", steps=invalid_deps)
    errors = orchestrator.validate_workflow(invalid_workflow)
    assert any("non-existent" in error.lower() for error in errors)


def test_has_circular_dependencies(orchestrator):
    """Test circular dependency detection."""
    # Valid workflow (no cycles)
    valid_steps = [
        WorkflowStep(id="step_1", pattern="pattern_a"),
        WorkflowStep(id="step_2", pattern="pattern_b", depends_on=["step_1"])
    ]
    valid_workflow = Workflow(id="test", name="test", steps=valid_steps)
    assert not orchestrator._has_circular_dependencies(valid_workflow)
    
    # Circular dependency
    circular_steps = [
        WorkflowStep(id="step_1", pattern="pattern_a", depends_on=["step_2"]),
        WorkflowStep(id="step_2", pattern="pattern_b", depends_on=["step_1"])
    ]
    circular_workflow = Workflow(id="test", name="test", steps=circular_steps)
    assert orchestrator._has_circular_dependencies(circular_workflow)


def test_get_workflow_progress(orchestrator):
    """Test workflow progress tracking."""
    workflow = orchestrator.create_workflow("test", ["pattern_a", "pattern_b"])
    
    # Initially all pending
    progress = orchestrator.get_workflow_progress(workflow.id)
    assert progress["progress"] == 0.0
    assert progress["total_steps"] == 2
    assert progress["completed_steps"] == 0
    
    # Complete first step
    workflow.steps[0].status = WorkflowStatus.COMPLETED
    progress = orchestrator.get_workflow_progress(workflow.id)
    assert progress["progress"] == 0.5
    assert progress["completed_steps"] == 1


def test_cancel_workflow(orchestrator):
    """Test workflow cancellation."""
    workflow = orchestrator.create_workflow("test", ["pattern_a"])
    workflow.status = WorkflowStatus.RUNNING
    
    success = orchestrator.cancel_workflow(workflow.id)
    assert success is True
    assert workflow.status == WorkflowStatus.CANCELLED


def test_cleanup_completed_workflows(orchestrator):
    """Test cleanup of old workflows."""
    # Create old completed workflow
    old_workflow = orchestrator.create_workflow("old", ["pattern_a"])
    old_workflow.status = WorkflowStatus.COMPLETED
    old_workflow.end_time = datetime.now() - timedelta(hours=25)  # 25 hours ago
    
    # Create recent workflow
    recent_workflow = orchestrator.create_workflow("recent", ["pattern_b"])
    recent_workflow.status = WorkflowStatus.COMPLETED
    recent_workflow.end_time = datetime.now() - timedelta(hours=1)  # 1 hour ago
    
    cleaned_count = orchestrator.cleanup_completed_workflows(max_age_hours=24)
    
    assert cleaned_count == 1
    assert old_workflow.id not in orchestrator._active_workflows
    assert recent_workflow.id in orchestrator._active_workflows


def test_singleton_pattern():
    """Test that get_workflow_orchestrator returns the same instance."""
    orch1 = get_workflow_orchestrator()
    orch2 = get_workflow_orchestrator()
    assert orch1 is orch2


def test_create_simple_workflow_convenience():
    """Test convenience function for creating simple workflows."""
    patterns = ["pattern_a", "pattern_b"]
    workflow = create_simple_workflow("test", patterns)
    
    assert workflow.name == "test"
    assert len(workflow.steps) == 2
    assert workflow.steps[1].depends_on == ["step_0"]  # Sequential dependency
