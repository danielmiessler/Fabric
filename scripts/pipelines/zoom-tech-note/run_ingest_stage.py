#!/usr/bin/env python3
"""Fabric wrapper for the copied Zoom deterministic ingest stage."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

from common import load_session_context


def main() -> int:
    parser = argparse.ArgumentParser(description="Run deterministic ingest for zoom-tech-note.")
    parser.add_argument("--session-context", required=True, help="Path to session_context.json")
    args = parser.parse_args()

    input_path = sys.stdin.read().strip()
    if not input_path:
        raise SystemExit("missing input path on stdin")

    context = load_session_context(Path(args.session_context).expanduser().resolve())
    session_dir = Path(context.session_dir).expanduser().resolve()

    cmd = [
        sys.executable,
        str(Path(__file__).resolve().parent / "ingest_zoom_captions.py"),
        input_path,
        "--output-dir",
        str(session_dir / ".pipeline"),
        "--session-id",
        context.session_id,
    ]
    completed = subprocess.run(cmd)
    return completed.returncode


if __name__ == "__main__":
    raise SystemExit(main())
