#!/usr/bin/env python3
"""Materialize semantic map artifacts for technical-study-guide."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from artifact_blocks import require_blocks
from common import load_session_context, render_input_text, run_dir


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    session_context = load_session_context(Path(args.session_context))
    blocks = require_blocks(
        Path("/dev/stdin").read_text(encoding="utf-8"),
        ["semantic_map.json", "render_input.md"],
    )

    semantic_map_content = blocks["semantic_map.json"].content
    semantic_map = json.loads(semantic_map_content)
    if not isinstance(semantic_map, dict):
        raise ValueError("semantic_map.json must contain a JSON object")

    render_input_content = blocks["render_input.md"].content

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    stable_pipeline_dir = Path(session_context.stable_pipeline_dir)
    stable_pipeline_dir.mkdir(parents=True, exist_ok=True)

    run_semantic_map_path = current_run_dir / "semantic_map.json"
    stable_semantic_map_path = stable_pipeline_dir / "semantic_map.json"
    run_render_input_path = current_run_dir / "render_input.md"

    normalized_semantic_map = json.dumps(semantic_map, ensure_ascii=True, indent=2) + "\n"
    run_semantic_map_path.write_text(normalized_semantic_map, encoding="utf-8")
    stable_semantic_map_path.write_text(normalized_semantic_map, encoding="utf-8")

    if not render_input_content.strip():
        source_text = Path(session_context.input_path).read_text(encoding="utf-8")
        render_input_content = render_input_text(session_context, semantic_map, source_text)
    run_render_input_path.write_text(render_input_content.rstrip("\n") + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
