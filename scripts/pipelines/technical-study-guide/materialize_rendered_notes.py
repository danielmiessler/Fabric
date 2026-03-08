#!/usr/bin/env python3
"""Materialize final notes for technical-study-guide."""

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
    blocks = require_blocks(Path("/dev/stdin").read_text(encoding="utf-8"), ["final_notes.md"])
    final_notes = blocks["final_notes.md"].content.rstrip("\n") + "\n"

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    run_final_output_path = current_run_dir / "final_notes.md"
    run_final_output_path.write_text(final_notes, encoding="utf-8")

    stable_final_output_path = Path(session_context.final_output_path)
    stable_final_output_path.parent.mkdir(parents=True, exist_ok=True)
    stable_final_output_path.write_text(final_notes, encoding="utf-8")

    print(final_notes, end="")


if __name__ == "__main__":
    main()
