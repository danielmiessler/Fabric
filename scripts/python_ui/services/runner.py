from __future__ import annotations
import subprocess, time
from typing import List, Dict, Optional
from utils.logging import logger
from utils.errors import ExecutionError
from utils.typing import RunResult, ChainStep
from utils.security import validate_pattern_name, sanitize_input
from services.monitoring import track_execution, update_execution_progress, complete_execution

FABRIC_BIN = "fabric"
DEFAULT_TIMEOUT = 90  # seconds
MAX_OUTPUT_BYTES = 1_000_000  # 1MB cap

def _run_cmd(args: List[str], input_text: str, timeout_s: int) -> RunResult:
    t0 = time.time()
    try:
        proc = subprocess.run(
            args,
            input=input_text.encode("utf-8"),
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            check=False,
            timeout=timeout_s,
        )
        out_b = proc.stdout or b""
        err_b = proc.stderr or b""

        # enforce output size cap
        truncated = False
        if len(out_b) > MAX_OUTPUT_BYTES:
            out_b = out_b[:MAX_OUTPUT_BYTES]
            truncated = True

        ok = proc.returncode == 0
        result = RunResult(
            success=ok,
            output=out_b.decode("utf-8", errors="replace"),
            error=None if ok else err_b.decode("utf-8", errors="replace"),
            duration_ms=int((time.time() - t0) * 1000),
            exit_code=proc.returncode,
            meta={"truncated": str(truncated)},
        )
        if not ok:
            logger.warning("runner: nonzero exit (%s): %s", proc.returncode, args)
        return result
    except subprocess.TimeoutExpired:
        dur = int((time.time() - t0) * 1000)
        msg = f"Execution timed out after {timeout_s}s"
        logger.error("runner timeout: %s", args)
        return RunResult(False, "", msg, dur, exit_code=-1, meta={"timeout": "true"})
    except Exception as e:
        dur = int((time.time() - t0) * 1000)
        logger.exception("runner exception for %s", args)
        return RunResult(False, "", f"Runner exception: {e}", dur, exit_code=-1, meta={})

def run_fabric(
    pattern: str,
    input_text: str,
    provider: str | None = None,
    model: str | None = None,
    timeout_s: int = DEFAULT_TIMEOUT,
    execution_id: Optional[str] = None,
) -> RunResult:
    """
    Execute a single Fabric pattern securely with real-time monitoring.
    - Validates pattern name, sanitizes input, enforces timeout & output caps.
    - Never uses shell=True.
    - Tracks execution progress for real-time monitoring.
    """
    validate_pattern_name(pattern)
    safe_input = sanitize_input(input_text, max_length=50_000)

    # Create or use existing execution tracking
    if execution_id is None:
        execution_id = track_execution(
            pattern=pattern,
            provider=provider,
            model=model,
            input_size=len(safe_input)
        )
    
    # Update progress to indicate execution started
    update_execution_progress(execution_id, 0.1)

    args: List[str] = [FABRIC_BIN, "-p", pattern]
    
    # Handle vendor/model specification safely
    norm_provider = (provider or "").strip() or None
    norm_model = (model or "").strip() or None
    if norm_model:
        if norm_provider and norm_provider != "All Providers":
            model_arg = f"{norm_provider}/{norm_model}"
        else:
            model_arg = norm_model
        args += ["-m", model_arg]

    logger.info("runner: %s (execution_id: %s)", " ".join(args), execution_id)
    
    # Update progress before execution
    update_execution_progress(execution_id, 0.2)
    
    # Execute command
    res = _run_cmd(args, safe_input, timeout_s=timeout_s)
    
    # Complete execution tracking
    complete_execution(execution_id, res)
    
    if not res.success:
        logger.debug("runner failure meta=%s", res.meta)
    
    return res

def run_chain(
    patterns: List[str],
    seed_input: str,
    provider: str | None = None,
    model: str | None = None,
    timeout_s: int = DEFAULT_TIMEOUT,
    continue_on_error: bool = False,
) -> List[ChainStep]:
    """
    Execute a sequence of patterns; each step consumes prior output as input.
    Returns a list of ChainStep with outputs/errors per stage.
    """
    steps: List[ChainStep] = []
    current_input = sanitize_input(seed_input, max_length=50_000)

    for name in patterns:
        validate_pattern_name(name)
        res = run_fabric(name, current_input, provider=provider, model=model, timeout_s=timeout_s)
        step = ChainStep(pattern=name, input=current_input, output=None, error=None, duration_ms=res.duration_ms)
        if res.success:
            step.output = res.output
            current_input = res.output
        else:
            step.error = res.error or "Unknown runner error"
            if not continue_on_error:
                steps.append(step)
                break
        steps.append(step)
    return steps