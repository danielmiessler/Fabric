#!/usr/bin/env python3
"""Materialize note-enhancement model outputs."""

from __future__ import annotations

import argparse
from pathlib import Path

from artifact_blocks import require_blocks
from common import load_session_context, run_dir


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    session_context = load_session_context(Path(args.session_context))
    raw_output = Path("/dev/stdin").read_text(encoding="utf-8")
    required = require_blocks(raw_output, ["enhanced_notes.md", "edit_log.md"])

    enhanced_notes = required["enhanced_notes.md"].content.rstrip("\n") + "\n"
    edit_log_text = required["edit_log.md"].content.rstrip("\n") + "\n"

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    stable_pipeline_dir = Path(session_context.stable_pipeline_dir)
    stable_pipeline_dir.mkdir(parents=True, exist_ok=True)

    run_enhanced_path = current_run_dir / "enhanced_notes.md"
    stable_output_path = Path(session_context.final_output_path)

    run_enhanced_path.write_text(enhanced_notes, encoding="utf-8")
    stable_output_path.parent.mkdir(parents=True, exist_ok=True)
    stable_output_path.write_text(enhanced_notes, encoding="utf-8")

    run_edit_log_path = current_run_dir / "edit_log.md"
    stable_edit_log_path = stable_pipeline_dir / "edit_log.md"
    run_edit_log_path.write_text(edit_log_text, encoding="utf-8")
    stable_edit_log_path.write_text(edit_log_text, encoding="utf-8")

    print(enhanced_notes, end="")


if __name__ == "__main__":
    main()
