#!/usr/bin/env python3
"""Shared helpers for the Fabric-native technical-study-guide pipeline."""

from __future__ import annotations

import json
import os
import re
from dataclasses import asdict, dataclass
from pathlib import Path


FABRIC_RUN_DIR_ENV = "FABRIC_PIPELINE_RUN_DIR"
FABRIC_SOURCE_MODE_ENV = "FABRIC_PIPELINE_SOURCE_MODE"
FABRIC_SOURCE_REFERENCE_ENV = "FABRIC_PIPELINE_SOURCE_REFERENCE"
FABRIC_INVOCATION_DIR_ENV = "FABRIC_PIPELINE_INVOCATION_DIR"


@dataclass
class SessionContext:
    source_mode: str
    source_reference: str
    invocation_dir: str
    session_dir: str
    input_path: str
    session_name: str
    session_id: str
    stable_pipeline_dir: str
    final_output_path: str


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
    return os.environ.get(FABRIC_SOURCE_MODE_ENV, "").strip()


def source_reference() -> str:
    return os.environ.get(FABRIC_SOURCE_REFERENCE_ENV, "").strip()


def read_stdin() -> str:
    return os.fdopen(0, "r", encoding="utf-8").read()


def resolve_input_path(path: Path) -> Path:
    resolved = path.expanduser().resolve()
    if resolved.is_file():
        if resolved.suffix.lower() not in {".md", ".txt"}:
            raise FileNotFoundError(f"Unsupported source file type: {resolved}")
        return resolved

    if resolved.is_dir():
        preferred = [
            "transcript.md",
            "transcript.txt",
            "notes.md",
            "notes.txt",
            "meeting_saved_closed_caption.txt",
        ]
        for candidate in preferred:
            candidate_path = resolved / candidate
            if candidate_path.exists():
                return candidate_path

        candidates = sorted(
            entry for entry in resolved.iterdir() if entry.is_file() and entry.suffix.lower() in {".md", ".txt"}
        )
        if len(candidates) == 1:
            return candidates[0]
        raise FileNotFoundError(f"Could not resolve unique technical-study-guide input in directory: {resolved}")

    raise FileNotFoundError(f"Input path not found: {resolved}")


def _ephemeral_session_identity(current_source_mode: str, current_source_reference: str) -> tuple[str, str]:
    if current_source_mode == "scrape_url" and current_source_reference:
        return slugify(current_source_reference), slugify(current_source_reference)
    if current_source_mode == "stdin":
        return "stdin-session", "stdin-session"
    return f"{current_source_mode or 'session'}-session", f"{current_source_mode or 'session'}-session"


def resolve_session_context(stdin_text: str) -> tuple[SessionContext, str]:
    current_invocation_dir = invocation_dir()
    current_source_mode = source_mode()
    current_source_reference = source_reference()

    if current_source_mode == "source" and current_source_reference:
        resolved_input = resolve_input_path(Path(current_source_reference))
        session_dir = resolved_input.parent
        stable_pipeline_dir = session_dir / ".pipeline"
        final_output_path = session_dir / "final_notes.md"
        session_name = session_dir.name
        session_id = slugify(session_name)
        raw_text = resolved_input.read_text(encoding="utf-8")
    elif current_source_mode in {"stdin", "scrape_url"}:
        session_dir = current_invocation_dir
        stable_pipeline_dir = run_dir()
        final_output_path = run_dir() / "final_notes.md"
        resolved_input = run_dir() / "source_input.md"
        raw_text = stdin_text
        resolved_input.parent.mkdir(parents=True, exist_ok=True)
        resolved_input.write_text(raw_text, encoding="utf-8")
        session_name, session_id = _ephemeral_session_identity(current_source_mode, current_source_reference)
    else:
        raise RuntimeError(
            "technical-study-guide supports source files/directories, stdin, or scrape_url input; "
            f"got source mode {current_source_mode!r}"
        )

    context = SessionContext(
        source_mode=current_source_mode,
        source_reference=current_source_reference,
        invocation_dir=str(current_invocation_dir),
        session_dir=str(session_dir),
        input_path=str(resolved_input),
        session_name=session_name,
        session_id=session_id,
        stable_pipeline_dir=str(stable_pipeline_dir),
        final_output_path=str(final_output_path),
    )
    return context, raw_text


def save_session_context(path: Path, context: SessionContext) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(asdict(context), ensure_ascii=True, indent=2) + "\n", encoding="utf-8")


def load_session_context(path: Path) -> SessionContext:
    data = json.loads(path.read_text(encoding="utf-8"))
    return SessionContext(**data)


def semantic_map_input_text(context: SessionContext, raw_text: str) -> str:
    return (
        "# Technical Study Guide Source\n\n"
        f"- session_id: `{context.session_id}`\n"
        f"- session_name: `{context.session_name}`\n"
        f"- source_mode: `{context.source_mode}`\n"
        f"- source_reference: `{context.source_reference or '(none)'}`\n\n"
        "## Instructions for semantic mapping\n\n"
        "- Preserve fidelity to the source.\n"
        "- Focus on technical intuition, systems understanding, math intuition, and coding progression.\n"
        "- Prepare structured material that can become a study guide.\n\n"
        "## Source Material\n\n"
        "```text\n"
        f"{raw_text.rstrip()}\n"
        "```\n"
    )


def render_input_text(context: SessionContext, semantic_map_json: dict, source_text: str) -> str:
    return (
        "# Technical Study Guide Rendering Brief\n\n"
        f"- session_id: `{context.session_id}`\n"
        f"- session_name: `{context.session_name}`\n"
        f"- source_mode: `{context.source_mode}`\n\n"
        "## Semantic Map\n\n"
        "```json\n"
        f"{json.dumps(semantic_map_json, ensure_ascii=True, indent=2)}\n"
        "```\n\n"
        "## Source Reminder\n\n"
        "```text\n"
        f"{source_text.rstrip()}\n"
        "```\n"
    )
