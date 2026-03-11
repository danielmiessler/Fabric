#!/usr/bin/env python3
"""Materialize model output artifacts for therapy-conversation-notes."""

from __future__ import annotations

import argparse
import json
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
    blocks = require_blocks(Path("/dev/stdin").read_text(encoding="utf-8"), ["final_notes.md", "analysis_snapshot.json"])

    final_notes = blocks["final_notes.md"].content.rstrip("\n") + "\n"
    snapshot = json.loads(blocks["analysis_snapshot.json"].content)
    if not isinstance(snapshot, dict):
        raise ValueError("analysis_snapshot.json must contain a JSON object")
    snapshot_json = json.dumps(snapshot, ensure_ascii=True, indent=2) + "\n"

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    stable_pipeline_dir = Path(session_context.stable_pipeline_dir)
    stable_pipeline_dir.mkdir(parents=True, exist_ok=True)

    run_final_notes_path = current_run_dir / "final_notes.md"
    run_snapshot_path = current_run_dir / "analysis_snapshot.json"
    stable_snapshot_path = stable_pipeline_dir / "analysis_snapshot.json"
    stable_final_output_path = Path(session_context.final_output_path)

    run_final_notes_path.write_text(final_notes, encoding="utf-8")
    run_snapshot_path.write_text(snapshot_json, encoding="utf-8")
    stable_snapshot_path.write_text(snapshot_json, encoding="utf-8")

    stable_final_output_path.parent.mkdir(parents=True, exist_ok=True)
    stable_final_output_path.write_text(final_notes, encoding="utf-8")

    print(final_notes, end="")


if __name__ == "__main__":
    main()
