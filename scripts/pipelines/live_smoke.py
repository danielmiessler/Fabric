#!/usr/bin/env python3
"""Live provider-backed smoke suite for model-backed built-in pipelines.

This script executes a bounded pipeline slice for each built-in pipeline that
contains model-backed `fabric_pattern` stages. It stops at the first model
stage (`--to-stage <stage-id>`) to keep latency and token usage predictable.
"""

from __future__ import annotations

import argparse
import json
import os
import subprocess
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Any


@dataclass(frozen=True)
class SmokeCase:
    pipeline: str
    to_stage: str
    stdin_text: str


GENERAL_SOURCE_TEXT = """\
Today we explored retrieval systems and prompting strategy.
The speaker compared sparse retrieval versus dense retrieval, then explained
hybrid retrieval and reranking.
The final section focused on how to validate claims against source material.
"""

NOTE_ENHANCEMENT_TEXT = """\
# rough notes
- talked about retrieval strategies
- sparse search is cheap and transparent
- dense search helps semantic matching
- use reranking before final answer
"""

THERAPY_SOURCE_TEXT = """\
Client shared difficulty with evening anxiety and rumination.
Client noticed episodes increase after social conflict and poor sleep.
We reflected on grounding patterns that have helped before and identified
small, realistic next actions for the coming week.
"""

ZOOM_CAPTION_TEXT = """\
[Instructor] 00:00:01
Welcome everyone. Today we cover deterministic pipelines.
[Instructor] 00:00:08
We begin with source ingestion and then build stable artifacts.
[Student] 00:00:17
How do we validate output quality before publishing?
[Instructor] 00:00:23
Use explicit validation stages and fail fast on contract violations.
"""


DEFAULT_CASES: tuple[SmokeCase, ...] = (
    SmokeCase(
        pipeline="technical-study-guide",
        to_stage="semantic_map_generate",
        stdin_text=GENERAL_SOURCE_TEXT,
    ),
    SmokeCase(
        pipeline="nontechnical-study-guide",
        to_stage="semantic_map_generate",
        stdin_text=GENERAL_SOURCE_TEXT,
    ),
    SmokeCase(
        pipeline="note-enhancement",
        to_stage="enhance_generate",
        stdin_text=NOTE_ENHANCEMENT_TEXT,
    ),
    SmokeCase(
        pipeline="therapy-conversation-notes",
        to_stage="analyze_generate",
        stdin_text=THERAPY_SOURCE_TEXT,
    ),
    SmokeCase(
        pipeline="zoom-tech-note",
        to_stage="stage1_refine_generate",
        stdin_text=ZOOM_CAPTION_TEXT,
    ),
    SmokeCase(
        pipeline="zoom-tech-note-deep-pass",
        to_stage="stage1_refine_generate",
        stdin_text=ZOOM_CAPTION_TEXT,
    ),
)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run live model-backed smoke checks for built-in pipelines.")
    parser.add_argument("--fabric-bin", default="./fabric", help="Path to the fabric binary")
    parser.add_argument("--vendor", default=os.environ.get("SMOKE_VENDOR", "OpenAI"), help="Vendor name")
    parser.add_argument("--model", default=os.environ.get("SMOKE_MODEL", "gpt-4.1-mini"), help="Model id")
    parser.add_argument(
        "--pipelines",
        default=os.environ.get("SMOKE_PIPELINES", "").strip(),
        help="Optional comma-separated pipeline subset",
    )
    parser.add_argument("--cwd", default=".", help="Working directory for fabric execution")
    parser.add_argument(
        "--timeout-seconds",
        type=int,
        default=int(os.environ.get("SMOKE_TIMEOUT_SECONDS", "240")),
        help="Per-pipeline timeout in seconds",
    )
    parser.add_argument("--list-cases", action="store_true", help="Print available smoke cases and exit")
    return parser.parse_args()


def parse_events(stderr_text: str) -> list[dict[str, Any]]:
    events: list[dict[str, Any]] = []
    for raw_line in stderr_text.splitlines():
        line = raw_line.strip()
        if not line or not line.startswith("{"):
            continue
        try:
            decoded = json.loads(line)
        except json.JSONDecodeError:
            continue
        if isinstance(decoded, dict) and isinstance(decoded.get("type"), str):
            events.append(decoded)
    return events


def event_exists(events: list[dict[str, Any]], event_type: str, stage_id: str | None = None) -> bool:
    for event in events:
        if event.get("type") != event_type:
            continue
        if stage_id is not None and event.get("stage_id") != stage_id:
            continue
        return True
    return False


def summarize_tail(value: str, limit: int = 1400) -> str:
    if len(value) <= limit:
        return value
    return value[-limit:]


def run_case(case: SmokeCase, fabric_bin: Path, cwd: Path, vendor: str, model: str, timeout_seconds: int) -> None:
    cmd = [
        str(fabric_bin),
        "--pipeline",
        case.pipeline,
        "--to-stage",
        case.to_stage,
        "--pipeline-events-json",
        "--vendor",
        vendor,
        "--model",
        model,
    ]

    completed = subprocess.run(
        cmd,
        cwd=str(cwd),
        input=case.stdin_text,
        text=True,
        capture_output=True,
        timeout=timeout_seconds,
        env=os.environ.copy(),
    )

    if completed.returncode != 0:
        raise RuntimeError(
            f"{case.pipeline}: non-zero exit ({completed.returncode})\n"
            f"stderr tail:\n{summarize_tail(completed.stderr)}\n"
            f"stdout tail:\n{summarize_tail(completed.stdout)}"
        )

    events = parse_events(completed.stderr)
    if not events:
        raise RuntimeError(f"{case.pipeline}: no JSON events observed on stderr")

    if not event_exists(events, "stage_started", stage_id=case.to_stage):
        raise RuntimeError(f"{case.pipeline}: missing stage_started event for {case.to_stage}")
    if not event_exists(events, "stage_passed", stage_id=case.to_stage):
        raise RuntimeError(f"{case.pipeline}: missing stage_passed event for {case.to_stage}")

    run_summaries = [event for event in events if event.get("type") == "run_summary"]
    if not run_summaries:
        raise RuntimeError(f"{case.pipeline}: missing run_summary event")
    if run_summaries[-1].get("status") != "passed":
        raise RuntimeError(f"{case.pipeline}: run_summary status is not passed")

    print(f"PASS {case.pipeline} -> {case.to_stage}")


def select_cases(pipeline_filter: str) -> list[SmokeCase]:
    if not pipeline_filter:
        return list(DEFAULT_CASES)

    requested = {item.strip() for item in pipeline_filter.split(",") if item.strip()}
    unknown = requested.difference({case.pipeline for case in DEFAULT_CASES})
    if unknown:
        unknown_list = ", ".join(sorted(unknown))
        raise ValueError(f"unknown pipeline(s) in --pipelines: {unknown_list}")

    return [case for case in DEFAULT_CASES if case.pipeline in requested]


def main() -> int:
    args = parse_args()

    if args.list_cases:
        for case in DEFAULT_CASES:
            print(f"{case.pipeline}:{case.to_stage}")
        return 0

    fabric_bin = Path(args.fabric_bin).expanduser().resolve()
    if not fabric_bin.exists():
        raise FileNotFoundError(f"fabric binary not found: {fabric_bin}")

    cwd = Path(args.cwd).expanduser().resolve()
    if not cwd.exists() or not cwd.is_dir():
        raise NotADirectoryError(f"invalid working directory: {cwd}")

    cases = select_cases(args.pipelines)
    print(
        f"Running {len(cases)} live smoke case(s) with vendor={args.vendor!r} model={args.model!r} cwd={cwd}"
    )
    for case in cases:
        run_case(
            case=case,
            fabric_bin=fabric_bin,
            cwd=cwd,
            vendor=args.vendor,
            model=args.model,
            timeout_seconds=args.timeout_seconds,
        )

    print("All live pipeline smoke checks passed.")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except KeyboardInterrupt:
        print("Interrupted.", file=sys.stderr)
        raise SystemExit(130)
    except Exception as exc:  # pragma: no cover - script-level error boundary
        print(f"ERROR: {exc}", file=sys.stderr)
        raise SystemExit(1)
