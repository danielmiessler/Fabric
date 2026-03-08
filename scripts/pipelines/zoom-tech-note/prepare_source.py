#!/usr/bin/env python3
"""Prepare source/session context for the Fabric-native zoom-tech-note pipeline."""

from __future__ import annotations

from pathlib import Path

from common import read_stdin, resolve_session_context, run_dir, save_session_context, stage1_input_text


def main() -> int:
    stdin_text = read_stdin()
    context, raw_text = resolve_session_context(stdin_text)

    current_run_dir = run_dir()
    session_context_path = current_run_dir / "session_context.json"
    stage1_input_path = current_run_dir / "stage1_input.md"

    save_session_context(session_context_path, context)
    stage1_input_path.write_text(stage1_input_text(context, raw_text), encoding="utf-8")

    print(context.input_path)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
