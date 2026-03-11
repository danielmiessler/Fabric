#!/usr/bin/env python3
"""Build analysis input for therapy-conversation-notes."""

from __future__ import annotations

import argparse
from pathlib import Path

from common import load_session_context, run_dir


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    parser.add_argument("--conversation-input", required=True)
    parser.add_argument("--context-pack", required=True)
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    session_context = load_session_context(Path(args.session_context))
    conversation_input = Path(args.conversation_input).read_text(encoding="utf-8")
    context_pack = Path(args.context_pack).read_text(encoding="utf-8")

    analysis_input = (
        "# Therapy Conversation Analysis Brief\n\n"
        f"- session_id: `{session_context.session_id}`\n"
        f"- session_name: `{session_context.session_name}`\n"
        f"- source_mode: `{session_context.source_mode}`\n\n"
        "## Analysis Contract\n\n"
        "- Produce reflection-oriented notes from the conversation.\n"
        "- Use optional context only as supporting lens, never as a diagnosis source.\n"
        "- Keep tone grounded, empathetic, and bounded.\n"
        "- Include explicit safety boundary language.\n\n"
        "## Primary Conversation Input\n\n"
        f"{conversation_input.rstrip()}\n\n"
        "## Optional Context Pack\n\n"
        f"{context_pack.rstrip()}\n"
    )

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    analysis_input_path = current_run_dir / "analysis_input.md"
    analysis_input_path.write_text(analysis_input.rstrip() + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
