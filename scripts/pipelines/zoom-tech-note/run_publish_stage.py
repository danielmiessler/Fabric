#!/usr/bin/env python3
"""Fabric wrapper for the copied Zoom publish stage."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

from common import load_session_context


def main() -> int:
    parser = argparse.ArgumentParser(description="Run publish normalization for zoom-tech-note.")
    parser.add_argument("--session-context", required=True, help="Path to session_context.json")
    args = parser.parse_args()

    context = load_session_context(Path(args.session_context).expanduser().resolve())
    session_dir = Path(context.session_dir).expanduser().resolve()

    cmd = [
        sys.executable,
        str(Path(__file__).resolve().parent / "publish_tutorial_notes.py"),
        "--session-dir",
        str(session_dir),
        "--root",
        str(session_dir),
    ]
    completed = subprocess.run(cmd)
    return completed.returncode


if __name__ == "__main__":
    raise SystemExit(main())
