"""
Workflow orchestrator service for managing complex pattern execution workflows.
Provides workflow building, dependency management, and execution coordination.
"""
from __future__ import annotations
import asyncio
import time
import uuid
from dataclasses import dataclass, field
from datetime import datetime, timedelta
from enum import Enum
from typing import List, Dict, Any, Optional, Callable, Set
from concurrent.futures import ThreadPoolExecutor, as_completed

from utils.logging import logger
from utils.errors import ExecutionError
from utils.typing import RunResult, ChainStep
from services import runner
from services.monitoring import get_execution_monitor, ExecutionStatus


class WorkflowStatus(Enum):
    """Workflow execution status."""
    PENDING = "pending"
    RUNNING = "running" 
    COMPLETED = "completed"
    FAILED = "failed"
    CANCELLED = "cancelled"


@dataclass
class ExecutionConfig:
    """Configuration for pattern execution."""
    provider: Optional[str] = None
    model: Optional[str] = None
    timeout_s: int = 90
    max_retries: int = 1
    retry_delay_s: float = 0.0
    parallel_limit: int = 3  # Maximum concurrent executions
    continue_on_error: bool = False


@dataclass
class WorkflowStep:
    """Individual step in a workflow."""
    id: str
    pattern: str
    condition: Optional[str] = None  # Python expression for conditional execution
    parallel_group: Optional[str] = None  # Group name for parallel execution
    timeout: int = 90
    depends_on: List[str] = field(default_factory=list)  # Step IDs this step depends on
    variables: Dict[str, str] = field(default_factory=dict)
    
    # Execution state
    status: WorkflowStatus = WorkflowStatus.PENDING
    start_time: Optional[datetime] = None
    end_time: Optional[datetime] = None
    result: Optional[RunResult] = None
    error: Optional[str] = None


@dataclass 
class Workflow:
    """Complete workflow definition with steps and execution plan."""
    id: str
    name: str
    steps: List[WorkflowStep]
    config: ExecutionConfig = field(default_factory=ExecutionConfig)
    
    # Execution state
    status: WorkflowStatus = WorkflowStatus.PENDING
    start_time: Optional[datetime] = None
    end_time: Optional[datetime] = None
    current_step: Optional[str] = None
    results: Dict[str, Any] = field(default_factory=dict)
    
    def get_step(self, step_id: str) -> Optional[WorkflowStep]:
        """Get workflow step by ID."""
        return next((step for step in self.steps if step.id == step_id), None)
    
    def get_steps_by_status(self, status: WorkflowStatus) -> List[WorkflowStep]:
        """Get all steps with specified status."""
        return [step for step in self.steps if step.status == status]
    
    def get_ready_steps(self) -> List[WorkflowStep]:
        """Get steps that are ready to execute (dependencies satisfied)."""
        ready_steps = []
        
        for step in self.steps:
            if step.status != WorkflowStatus.PENDING:
                continue
                
            # Check if all dependencies are completed
            dependencies_met = all(
                self.get_step(dep_id) and self.get_step(dep_id).status == WorkflowStatus.COMPLETED
                for dep_id in step.depends_on
            )
            
            if dependencies_met:
                ready_steps.append(step)
        
        return ready_steps


class WorkflowOrchestrator:
    """Orchestrates workflow execution with dependency management and parallel processing."""
    
    def __init__(self):
        self._active_workflows: Dict[str, Workflow] = {}
        self._execution_monitor = get_execution_monitor()
        self._max_workers = 4  # Maximum concurrent executions
        
    def create_workflow(
        self, 
        name: str, 
        patterns: List[str], 
        config: Optional[ExecutionConfig] = None
    ) -> Workflow:
        """
        Create a simple sequential workflow from pattern list.
        
        Args:
            name: Workflow name
            patterns: List of pattern names
            config: Optional execution configuration
            
        Returns:
            Created workflow
        """
        workflow_id = str(uuid.uuid4())
        
        if config is None:
            config = ExecutionConfig()
        
        # Create sequential steps
        steps = []
        for i, pattern in enumerate(patterns):
            step_id = f"step_{i}"
            depends_on = [f"step_{i-1}"] if i > 0 else []
            
            step = WorkflowStep(
                id=step_id,
                pattern=pattern,
                depends_on=depends_on,
                timeout=config.timeout_s
            )
            steps.append(step)
        
        workflow = Workflow(
            id=workflow_id,
            name=name,
            steps=steps,
            config=config
        )
        
        self._active_workflows[workflow_id] = workflow
        logger.info(f"Created workflow {name} with {len(patterns)} steps")
        
        return workflow
    
    def create_parallel_workflow(
        self, 
        name: str, 
        parallel_groups: Dict[str, List[str]], 
        sequential_patterns: List[str] = None,
        config: Optional[ExecutionConfig] = None
    ) -> Workflow:
        """
        Create a workflow with parallel execution groups.
        
        Args:
            name: Workflow name
            parallel_groups: Dict of group_name -> patterns that can run in parallel
            sequential_patterns: Patterns that must run sequentially
            config: Optional execution configuration
            
        Returns:
            Created workflow
        """
        workflow_id = str(uuid.uuid4())
        
        if config is None:
            config = ExecutionConfig()
        if sequential_patterns is None:
            sequential_patterns = []
        
        steps = []
        step_counter = 0
        
        # Add parallel groups
        for group_name, group_patterns in parallel_groups.items():
            for pattern in group_patterns:
                step = WorkflowStep(
                    id=f"step_{step_counter}",
                    pattern=pattern,
                    parallel_group=group_name,
                    timeout=config.timeout_s
                )
                steps.append(step)
                step_counter += 1
        
        # Add sequential patterns (depend on all parallel groups)
        parallel_step_ids = [step.id for step in steps]
        
        for pattern in sequential_patterns:
            step = WorkflowStep(
                id=f"step_{step_counter}",
                pattern=pattern,
                depends_on=parallel_step_ids.copy(),
                timeout=config.timeout_s
            )
            steps.append(step)
            step_counter += 1
        
        workflow = Workflow(
            id=workflow_id,
            name=name,
            steps=steps,
            config=config
        )
        
        self._active_workflows[workflow_id] = workflow
        logger.info(f"Created parallel workflow {name} with {len(steps)} steps")
        
        return workflow
    
    def create_custom_workflow(
        self, 
        name: str, 
        steps: List[WorkflowStep], 
        config: Optional[ExecutionConfig] = None
    ) -> Workflow:
        """
        Create a custom workflow with user-defined steps.
        
        Args:
            name: Workflow name
            steps: List of workflow steps
            config: Optional execution configuration
            
        Returns:
            Created workflow
        """
        workflow_id = str(uuid.uuid4())
        
        if config is None:
            config = ExecutionConfig()
        
        # Validate dependencies
        step_ids = {step.id for step in steps}
        for step in steps:
            for dep_id in step.depends_on:
                if dep_id not in step_ids:
                    raise ValueError(f"Step {step.id} depends on non-existent step {dep_id}")
        
        workflow = Workflow(
            id=workflow_id,
            name=name,
            steps=steps,
            config=config
        )
        
        self._active_workflows[workflow_id] = workflow
        logger.info(f"Created custom workflow {name} with {len(steps)} steps")
        
        return workflow
    
    def execute_workflow(
        self, 
        workflow: Workflow, 
        input_text: str,
        progress_callback: Optional[Callable[[str, float], None]] = None
    ) -> Dict[str, Any]:
        """
        Execute a workflow with dependency management and parallel processing.
        
        Args:
            workflow: Workflow to execute
            input_text: Initial input text
            progress_callback: Optional callback for progress updates
            
        Returns:
            Execution results dictionary
        """
        workflow.status = WorkflowStatus.RUNNING
        workflow.start_time = datetime.now()
        
        try:
            # Group steps for parallel execution
            execution_groups = self._plan_execution_groups(workflow)
            
            # Execute groups sequentially, steps within groups in parallel
            step_outputs: Dict[str, str] = {}
            step_outputs["_initial"] = input_text  # Initial input
            
            total_groups = len(execution_groups)
            
            for group_index, group_steps in enumerate(execution_groups):
                logger.info(f"Executing group {group_index + 1}/{total_groups} with {len(group_steps)} steps")
                
                # Execute steps in this group (parallel if possible)
                group_results = self._execute_step_group(
                    group_steps, step_outputs, workflow.config
                )
                
                # Update step outputs and workflow state
                for step_id, result in group_results.items():
                    step = workflow.get_step(step_id)
                    if step:
                        step.result = result
                        step.end_time = datetime.now()
                        
                        if result.success:
                            step.status = WorkflowStatus.COMPLETED
                            step_outputs[step_id] = result.output
                        else:
                            step.status = WorkflowStatus.FAILED
                            step.error = result.error
                            
                            # Stop workflow on error if not configured to continue
                            if not workflow.config.continue_on_error:
                                raise ExecutionError(f"Step {step_id} failed: {result.error}")
                
                # Update progress
                if progress_callback:
                    progress = (group_index + 1) / total_groups
                    progress_callback(workflow.id, progress)
            
            # Determine final output
            final_output = None
            if workflow.steps:
                # Get output from last successful step
                for step in reversed(workflow.steps):
                    if step.status == WorkflowStatus.COMPLETED and step.result:
                        final_output = step.result.output
                        break
            
            workflow.status = WorkflowStatus.COMPLETED
            workflow.end_time = datetime.now()
            
            return {
                "workflow_id": workflow.id,
                "success": True,
                "final_output": final_output,
                "step_results": {step.id: step.result for step in workflow.steps},
                "execution_time": (workflow.end_time - workflow.start_time).total_seconds(),
                "metadata": {
                    "total_steps": len(workflow.steps),
                    "successful_steps": len(workflow.get_steps_by_status(WorkflowStatus.COMPLETED)),
                    "failed_steps": len(workflow.get_steps_by_status(WorkflowStatus.FAILED))
                }
            }
            
        except Exception as e:
            workflow.status = WorkflowStatus.FAILED
            workflow.end_time = datetime.now()
            
            logger.error(f"Workflow {workflow.id} execution failed: {e}")
            return {
                "workflow_id": workflow.id,
                "success": False,
                "final_output": None,
                "error": str(e),
                "step_results": {step.id: step.result for step in workflow.steps},
                "execution_time": (workflow.end_time - workflow.start_time).total_seconds() if workflow.start_time else 0,
                "metadata": {
                    "total_steps": len(workflow.steps),
                    "successful_steps": len(workflow.get_steps_by_status(WorkflowStatus.COMPLETED)),
                    "failed_steps": len(workflow.get_steps_by_status(WorkflowStatus.FAILED))
                }
            }
    
    def _plan_execution_groups(self, workflow: Workflow) -> List[List[WorkflowStep]]:
        """
        Plan execution groups based on dependencies and parallel groups.
        
        Returns:
            List of step groups to execute sequentially
        """
        groups = []
        remaining_steps = workflow.steps.copy()
        
        while remaining_steps:
            # Find steps that can execute now
            ready_steps = []
            
            for step in remaining_steps:
                # For initial planning, treat steps with no dependencies as ready
                if not step.depends_on:
                    ready_steps.append(step)
                else:
                    # Check if dependencies are satisfied
                    dependencies_satisfied = all(
                        workflow.get_step(dep_id) and workflow.get_step(dep_id).status == WorkflowStatus.COMPLETED
                        for dep_id in step.depends_on
                    )
                    
                    if dependencies_satisfied:
                        ready_steps.append(step)
            
            if not ready_steps:
                # Check for circular dependencies
                remaining_step_ids = [step.id for step in remaining_steps]
                logger.error(f"No ready steps found. Remaining: {remaining_step_ids}")
                raise ExecutionError("Circular dependency detected in workflow")
            
            # Group ready steps by parallel group
            if any(step.parallel_group for step in ready_steps):
                # Group by parallel_group
                parallel_groups = {}
                for step in ready_steps:
                    group_key = step.parallel_group or f"_sequential_{step.id}"
                    if group_key not in parallel_groups:
                        parallel_groups[group_key] = []
                    parallel_groups[group_key].append(step)
                
                # Add each parallel group as a separate execution group
                for group_steps in parallel_groups.values():
                    groups.append(group_steps)
            else:
                # No parallel groups, add all ready steps as one group
                groups.append(ready_steps)
            
            # Remove ready steps from remaining and mark as completed for dependency resolution
            for step in ready_steps:
                remaining_steps.remove(step)
                # Temporarily mark as completed for dependency resolution in planning
                step.status = WorkflowStatus.COMPLETED
        
        # Reset all steps back to pending after planning
        for step in workflow.steps:
            step.status = WorkflowStatus.PENDING
        
        return groups
    
    def _execute_step_group(
        self, 
        steps: List[WorkflowStep], 
        step_outputs: Dict[str, str],
        config: ExecutionConfig
    ) -> Dict[str, RunResult]:
        """
        Execute a group of steps (potentially in parallel).
        
        Args:
            steps: Steps to execute
            step_outputs: Available outputs from previous steps
            config: Execution configuration
            
        Returns:
            Dictionary of step_id -> RunResult
        """
        if len(steps) == 1:
            # Single step execution
            step = steps[0]
            input_text = self._resolve_step_input(step, step_outputs)
            
            step.status = WorkflowStatus.RUNNING
            step.start_time = datetime.now()
            
            result = runner.run_fabric(
                pattern=step.pattern,
                input_text=input_text,
                provider=config.provider,
                model=config.model,
                timeout_s=step.timeout
            )
            
            return {step.id: result}
        
        # Multiple steps - execute in parallel if possible
        results = {}
        
        # Check if steps can truly run in parallel (same parallel_group)
        can_parallel = (
            len(set(step.parallel_group for step in steps)) == 1 and 
            steps[0].parallel_group is not None
        )
        
        if can_parallel and len(steps) <= config.parallel_limit:
            # Execute in parallel using ThreadPoolExecutor
            with ThreadPoolExecutor(max_workers=min(len(steps), config.parallel_limit)) as executor:
                # Submit all steps
                future_to_step = {}
                
                for step in steps:
                    input_text = self._resolve_step_input(step, step_outputs)
                    step.status = WorkflowStatus.RUNNING
                    step.start_time = datetime.now()
                    
                    future = executor.submit(
                        runner.run_fabric,
                        pattern=step.pattern,
                        input_text=input_text,
                        provider=config.provider,
                        model=config.model,
                        timeout_s=step.timeout
                    )
                    future_to_step[future] = step
                
                # Collect results as they complete
                for future in as_completed(future_to_step):
                    step = future_to_step[future]
                    try:
                        result = future.result()
                        results[step.id] = result
                    except Exception as e:
                        logger.error(f"Step {step.id} execution failed: {e}")
                        results[step.id] = RunResult(
                            success=False,
                            output="",
                            error=str(e),
                            duration_ms=0,
                            exit_code=-1
                        )
        else:
            # Execute sequentially
            for step in steps:
                input_text = self._resolve_step_input(step, step_outputs)
                
                step.status = WorkflowStatus.RUNNING
                step.start_time = datetime.now()
                
                result = runner.run_fabric(
                    pattern=step.pattern,
                    input_text=input_text,
                    provider=config.provider,
                    model=config.model,
                    timeout_s=step.timeout
                )
                
                results[step.id] = result
                
                # Update step_outputs for next step
                if result.success:
                    step_outputs[step.id] = result.output
        
        return results
    
    def _resolve_step_input(self, step: WorkflowStep, step_outputs: Dict[str, str]) -> str:
        """
        Resolve input for a workflow step based on dependencies.
        
        Args:
            step: Step to resolve input for
            step_outputs: Available outputs from previous steps
            
        Returns:
            Input text for the step
        """
        if not step.depends_on:
            # No dependencies, use initial input
            return step_outputs.get("_initial", "")
        
        if len(step.depends_on) == 1:
            # Single dependency, use its output
            dep_id = step.depends_on[0]
            return step_outputs.get(dep_id, step_outputs.get("_initial", ""))
        
        # Multiple dependencies, concatenate outputs
        inputs = []
        for dep_id in step.depends_on:
            if dep_id in step_outputs:
                inputs.append(step_outputs[dep_id])
        
        if not inputs:
            return step_outputs.get("_initial", "")
        
        # Join multiple inputs with clear separators
        return "\n\n---\n\n".join(inputs)
    
    def get_workflow(self, workflow_id: str) -> Optional[Workflow]:
        """Get workflow by ID."""
        return self._active_workflows.get(workflow_id)
    
    def list_active_workflows(self) -> List[Workflow]:
        """Get list of all active workflows."""
        return [
            workflow for workflow in self._active_workflows.values()
            if workflow.status in [WorkflowStatus.PENDING, WorkflowStatus.RUNNING]
        ]
    
    def cancel_workflow(self, workflow_id: str) -> bool:
        """Cancel a running workflow."""
        workflow = self._active_workflows.get(workflow_id)
        if not workflow:
            return False
        
        workflow.status = WorkflowStatus.CANCELLED
        workflow.end_time = datetime.now()
        
        # Cancel any running steps
        for step in workflow.steps:
            if step.status == WorkflowStatus.RUNNING:
                step.status = WorkflowStatus.CANCELLED
        
        logger.info(f"Cancelled workflow {workflow_id}")
        return True
    
    def get_workflow_progress(self, workflow_id: str) -> Dict[str, Any]:
        """
        Get workflow execution progress.
        
        Args:
            workflow_id: Workflow ID
            
        Returns:
            Progress information dictionary
        """
        workflow = self._active_workflows.get(workflow_id)
        if not workflow:
            return {"error": "Workflow not found"}
        
        total_steps = len(workflow.steps)
        completed_steps = len(workflow.get_steps_by_status(WorkflowStatus.COMPLETED))
        failed_steps = len(workflow.get_steps_by_status(WorkflowStatus.FAILED))
        running_steps = len(workflow.get_steps_by_status(WorkflowStatus.RUNNING))
        
        progress = completed_steps / total_steps if total_steps > 0 else 0.0
        
        return {
            "workflow_id": workflow_id,
            "name": workflow.name,
            "status": workflow.status.value,
            "progress": progress,
            "total_steps": total_steps,
            "completed_steps": completed_steps,
            "failed_steps": failed_steps,
            "running_steps": running_steps,
            "current_step": workflow.current_step,
            "start_time": workflow.start_time.isoformat() if workflow.start_time else None,
            "estimated_completion": self._estimate_completion_time(workflow)
        }
    
    def _estimate_completion_time(self, workflow: Workflow) -> Optional[str]:
        """Estimate workflow completion time based on remaining steps."""
        if workflow.status != WorkflowStatus.RUNNING:
            return None
        
        try:
            remaining_steps = workflow.get_steps_by_status(WorkflowStatus.PENDING)
            if not remaining_steps:
                return None
            
            # Estimate time based on pattern complexity and remaining work
            total_estimated_seconds = 0
            for step in remaining_steps:
                # Use pattern-based time estimation
                estimated_time = 30.0  # Default 30 seconds per step
                total_estimated_seconds += estimated_time
            
            # Adjust for parallel execution potential
            parallel_groups = set(step.parallel_group for step in remaining_steps if step.parallel_group)
            if parallel_groups:
                # Assume some parallelization benefit
                total_estimated_seconds *= 0.7
            
            completion_time = datetime.now() + timedelta(seconds=total_estimated_seconds)
            return completion_time.isoformat()
            
        except Exception as e:
            logger.warning(f"Completion time estimation failed: {e}")
            return None
    
    def validate_workflow(self, workflow: Workflow) -> List[str]:
        """
        Validate workflow for common issues.
        
        Args:
            workflow: Workflow to validate
            
        Returns:
            List of validation errors (empty if valid)
        """
        errors = []
        
        try:
            # Check for empty workflow
            if not workflow.steps:
                errors.append("Workflow has no steps")
                return errors
            
            # Check step IDs are unique
            step_ids = [step.id for step in workflow.steps]
            if len(set(step_ids)) != len(step_ids):
                errors.append("Duplicate step IDs found")
            
            # Check dependencies exist
            step_id_set = set(step_ids)
            for step in workflow.steps:
                for dep_id in step.depends_on:
                    if dep_id not in step_id_set:
                        errors.append(f"Step {step.id} depends on non-existent step {dep_id}")
            
            # Check for circular dependencies
            if self._has_circular_dependencies(workflow):
                errors.append("Circular dependencies detected")
            
            # Check pattern names are valid
            try:
                available_patterns = patterns.list_patterns()
                available_names = {spec.name for spec in available_patterns}
                
                for step in workflow.steps:
                    if step.pattern not in available_names:
                        errors.append(f"Step {step.id} references unknown pattern: {step.pattern}")
            except Exception as e:
                logger.warning(f"Could not validate pattern names: {e}")
                
        except Exception as e:
            logger.error(f"Workflow validation failed: {e}")
            errors.append(f"Validation error: {e}")
        
        return errors
    
    def _has_circular_dependencies(self, workflow: Workflow) -> bool:
        """Check for circular dependencies in workflow."""
        try:
            # Use topological sort to detect cycles
            in_degree = {step.id: len(step.depends_on) for step in workflow.steps}
            queue = [step.id for step in workflow.steps if len(step.depends_on) == 0]
            processed = 0
            
            while queue:
                current = queue.pop(0)
                processed += 1
                
                # Find steps that depend on current
                for step in workflow.steps:
                    if current in step.depends_on:
                        in_degree[step.id] -= 1
                        if in_degree[step.id] == 0:
                            queue.append(step.id)
            
            # If we couldn't process all steps, there's a cycle
            return processed != len(workflow.steps)
            
        except Exception as e:
            logger.error(f"Circular dependency check failed: {e}")
            return True  # Assume circular dependency on error
    
    def cleanup_completed_workflows(self, max_age_hours: int = 24) -> int:
        """
        Clean up old completed workflows to prevent memory leaks.
        
        Args:
            max_age_hours: Maximum age in hours for keeping workflows
            
        Returns:
            Number of workflows cleaned up
        """
        cutoff_time = datetime.now() - timedelta(hours=max_age_hours)
        cleaned_count = 0
        
        workflows_to_remove = []
        for workflow_id, workflow in self._active_workflows.items():
            if (workflow.status in [WorkflowStatus.COMPLETED, WorkflowStatus.FAILED, WorkflowStatus.CANCELLED] and
                workflow.end_time and workflow.end_time < cutoff_time):
                workflows_to_remove.append(workflow_id)
        
        for workflow_id in workflows_to_remove:
            del self._active_workflows[workflow_id]
            cleaned_count += 1
        
        if cleaned_count > 0:
            logger.info(f"Cleaned up {cleaned_count} old workflows")
        
        return cleaned_count


# Singleton instance
_workflow_orchestrator = None


def get_workflow_orchestrator() -> WorkflowOrchestrator:
    """Get singleton WorkflowOrchestrator instance."""
    global _workflow_orchestrator
    if _workflow_orchestrator is None:
        _workflow_orchestrator = WorkflowOrchestrator()
    return _workflow_orchestrator


# Convenience functions for easy access
def create_simple_workflow(name: str, patterns: List[str], config: ExecutionConfig = None) -> Workflow:
    """Create a simple sequential workflow."""
    orchestrator = get_workflow_orchestrator()
    return orchestrator.create_workflow(name, patterns, config)


def execute_workflow_simple(workflow: Workflow, input_text: str) -> Dict[str, Any]:
    """Execute a workflow with default settings."""
    orchestrator = get_workflow_orchestrator()
    return orchestrator.execute_workflow(workflow, input_text)
