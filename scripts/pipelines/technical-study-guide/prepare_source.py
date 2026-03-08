#!/usr/bin/env python3
"""Prepare technical-study-guide source artifacts."""

from __future__ import annotations

from pathlib import Path

from common import read_stdin, resolve_session_context, run_dir, save_session_context, semantic_map_input_text


def main() -> None:
    stdin_text = read_stdin()
    context, raw_text = resolve_session_context(stdin_text)

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)

    session_context_path = current_run_dir / "session_context.json"
    semantic_map_input_path = current_run_dir / "semantic_map_input.md"

    save_session_context(session_context_path, context)
    semantic_map_input_path.write_text(semantic_map_input_text(context, raw_text), encoding="utf-8")

    print(str(Path(context.input_path)))


if __name__ == "__main__":
    main()
