#!/usr/bin/env python3
"""Materialize Stage 1 zoom-tech-note artifacts from Fabric pattern output."""

from __future__ import annotations

import argparse
import sys
from pathlib import Path

from artifact_blocks import require_blocks
from common import load_session_context, run_dir, stage2_input_text


REQUIRED_PATHS = [
    ".pipeline/refined_transcript.md",
    ".pipeline/topic_inventory.json",
    ".pipeline/corrections_log.csv",
    ".pipeline/uncertainty_report.json",
]


def write_session_artifact(session_dir: Path, relative_path: str, content: str) -> Path:
    target = session_dir / relative_path
    target.parent.mkdir(parents=True, exist_ok=True)
    target.write_text(content, encoding="utf-8")
    return target


def main() -> int:
    parser = argparse.ArgumentParser(description="Materialize Stage 1 outputs for zoom-tech-note.")
    parser.add_argument("--session-context", required=True, help="Path to session_context.json")
    args = parser.parse_args()

    raw_output = sys.stdin.read()
    blocks = require_blocks(raw_output, REQUIRED_PATHS)
    context = load_session_context(Path(args.session_context).expanduser().resolve())
    session_dir = Path(context.session_dir).expanduser().resolve()

    for path in REQUIRED_PATHS:
        write_session_artifact(session_dir, path, blocks[path].content)

    next_input_path = run_dir() / "stage2_input.md"
    next_input_path.write_text(stage2_input_text(session_dir), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
