#!/usr/bin/env python3
"""Shared helpers for the Fabric-native zoom-tech-note pipeline."""

from __future__ import annotations

import json
import os
import re
from dataclasses import asdict, dataclass
from pathlib import Path


FABRIC_RUN_DIR_ENV = "FABRIC_PIPELINE_RUN_DIR"
FABRIC_STAGE_ID_ENV = "FABRIC_PIPELINE_STAGE_ID"
FABRIC_SOURCE_MODE_ENV = "FABRIC_PIPELINE_SOURCE_MODE"
FABRIC_SOURCE_REFERENCE_ENV = "FABRIC_PIPELINE_SOURCE_REFERENCE"
FABRIC_INVOCATION_DIR_ENV = "FABRIC_PIPELINE_INVOCATION_DIR"

HEADER_RE = re.compile(r"^\[(?P<speaker>[^\]]+)\]\s+(?P<timestamp>\d{2}:\d{2}:\d{2})\s*$")


@dataclass
class SessionContext:
    source_mode: str
    source_reference: str
    invocation_dir: str
    session_dir: str
    input_path: str
    session_name: str
    session_id: str


def slugify(value: str) -> str:
    slug = re.sub(r"[^a-zA-Z0-9]+", "-", value.strip().lower()).strip("-")
    return slug or "session"


def env_path(name: str) -> Path:
    value = os.environ.get(name, "")
    if not value:
        raise RuntimeError(f"missing required environment variable: {name}")
    return Path(value).expanduser().resolve()


def run_dir() -> Path:
    return env_path(FABRIC_RUN_DIR_ENV)


def invocation_dir() -> Path:
    return env_path(FABRIC_INVOCATION_DIR_ENV)


def source_mode() -> str:
    return os.environ.get(FABRIC_SOURCE_MODE_ENV, "")


def source_reference() -> str:
    return os.environ.get(FABRIC_SOURCE_REFERENCE_ENV, "").strip()


def resolve_zoom_input(path: Path) -> Path:
    path = path.expanduser().resolve()
    if path.is_file():
        return path
    if path.is_dir():
        exact = path / "meeting_saved_closed_caption.txt"
        if exact.exists():
            return exact
        txts = sorted(path.glob("*.txt"))
        if len(txts) == 1:
            return txts[0]
        raise FileNotFoundError(f"Could not resolve unique transcript .txt in directory: {path}")
    raise FileNotFoundError(f"Input path not found: {path}")


def read_stdin() -> str:
    return os.fdopen(0, "r", encoding="utf-8").read()


def resolve_session_context(stdin_text: str) -> tuple[SessionContext, str]:
    current_invocation_dir = invocation_dir()
    current_source_mode = source_mode()
    current_source_reference = source_reference()

    if current_source_mode == "source" and current_source_reference:
        resolved_input = resolve_zoom_input(Path(current_source_reference))
        session_dir = resolved_input.parent
        raw_text = resolved_input.read_text(encoding="utf-8")
    elif current_source_mode == "stdin":
        session_dir = current_invocation_dir
        resolved_input = run_dir() / "meeting_saved_closed_caption.txt"
        raw_text = stdin_text
        resolved_input.parent.mkdir(parents=True, exist_ok=True)
        resolved_input.write_text(raw_text, encoding="utf-8")
    else:
        raise RuntimeError(
            "zoom-tech-note supports source files/directories or stdin only; "
            f"got source mode {current_source_mode!r}"
        )

    context = SessionContext(
        source_mode=current_source_mode,
        source_reference=current_source_reference,
        invocation_dir=str(current_invocation_dir),
        session_dir=str(session_dir),
        input_path=str(resolved_input),
        session_name=session_dir.name,
        session_id=slugify(session_dir.name),
    )
    return context, raw_text


def save_session_context(path: Path, context: SessionContext) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(asdict(context), ensure_ascii=True, indent=2) + "\n", encoding="utf-8")


def load_session_context(path: Path) -> SessionContext:
    data = json.loads(path.read_text(encoding="utf-8"))
    return SessionContext(**data)


def stage1_input_text(context: SessionContext, raw_text: str) -> str:
    return (
        "# Stage 1 Input\n\n"
        f"- session_id: `{context.session_id}`\n"
        f"- session_name: `{context.session_name}`\n"
        "- run_mode: `tool-restricted`\n\n"
        "## Raw Transcript\n\n"
        "```text\n"
        f"{raw_text.rstrip()}\n"
        "```\n"
    )


def stage2_input_text(session_dir: Path) -> str:
    pipeline_dir = session_dir / ".pipeline"
    refined = (pipeline_dir / "refined_transcript.md").read_text(encoding="utf-8").rstrip()
    inventory = (pipeline_dir / "topic_inventory.json").read_text(encoding="utf-8").rstrip()
    manifest = (pipeline_dir / "segment_manifest.jsonl").read_text(encoding="utf-8").rstrip()
    return (
        "# Stage 2 Input\n\n"
        "- run_mode: `tool-restricted`\n\n"
        "## .pipeline/refined_transcript.md\n\n"
        "```markdown\n"
        f"{refined}\n"
        "```\n\n"
        "## .pipeline/topic_inventory.json\n\n"
        "```json\n"
        f"{inventory}\n"
        "```\n\n"
        "## .pipeline/segment_manifest.jsonl\n\n"
        "```jsonl\n"
        f"{manifest}\n"
        "```\n"
    )


def stage3_input_text(session_dir: Path) -> str:
    pipeline_dir = session_dir / ".pipeline"
    structured = (pipeline_dir / "structured_notes.md").read_text(encoding="utf-8").rstrip()
    coverage = (pipeline_dir / "coverage_matrix.json").read_text(encoding="utf-8").rstrip()
    inventory = (pipeline_dir / "topic_inventory.json").read_text(encoding="utf-8").rstrip()
    return (
        "# Stage 3 Input\n\n"
        "- run_mode: `tool-restricted`\n\n"
        "## .pipeline/structured_notes.md\n\n"
        "```markdown\n"
        f"{structured}\n"
        "```\n\n"
        "## .pipeline/coverage_matrix.json\n\n"
        "```json\n"
        f"{coverage}\n"
        "```\n\n"
        "## .pipeline/topic_inventory.json\n\n"
        "```json\n"
        f"{inventory}\n"
        "```\n"
    )
