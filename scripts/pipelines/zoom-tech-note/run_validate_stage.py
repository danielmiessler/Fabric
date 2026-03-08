#!/usr/bin/env python3
"""Fabric wrapper for the copied Zoom validation stage."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

from common import load_session_context


def main() -> int:
    parser = argparse.ArgumentParser(description="Run deterministic validation for zoom-tech-note.")
    parser.add_argument("--session-context", required=True, help="Path to session_context.json")
    args = parser.parse_args()

    context = load_session_context(Path(args.session_context).expanduser().resolve())
    session_dir = Path(context.session_dir).expanduser().resolve()
    pipeline_dir = session_dir / ".pipeline"

    cmd = [
        sys.executable,
        str(Path(__file__).resolve().parent / "validate_coverage.py"),
        "--ledger",
        str(pipeline_dir / "segment_ledger.jsonl"),
        "--final-notes",
        str(session_dir / "final_notes.md"),
        "--coverage-matrix",
        str(pipeline_dir / "coverage_matrix.json"),
        "--uncertainty-report",
        str(pipeline_dir / "uncertainty_report.json"),
        "--human-review-queue",
        str(pipeline_dir / "human_review_queue.md"),
        "--report-out",
        str(pipeline_dir / "validation_report.md"),
        "--exceptions-out",
        str(pipeline_dir / "exceptions.json"),
    ]
    completed = subprocess.run(cmd)
    return completed.returncode


if __name__ == "__main__":
    raise SystemExit(main())
