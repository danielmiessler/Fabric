#!/usr/bin/env python3
"""Deterministic parser for note-enhancement artifact blocks."""

from __future__ import annotations

from dataclasses import dataclass


BEGIN_PREFIX = "<<<BEGIN_ARTIFACT:"
END_MARKER = "<<<END_ARTIFACT>>>"


@dataclass
class ArtifactBlock:
    path: str
    content: str


def parse_artifact_blocks(text: str) -> dict[str, ArtifactBlock]:
    lines = text.splitlines()
    idx = 0
    blocks: dict[str, ArtifactBlock] = {}

    while idx < len(lines):
        line = lines[idx].strip()
        if not line.startswith(BEGIN_PREFIX) or not line.endswith(">>>"):
            idx += 1
            continue

        path = line[len(BEGIN_PREFIX) : -3].strip()
        idx += 1
        content_lines: list[str] = []

        while idx < len(lines) and lines[idx].strip() != END_MARKER:
            content_lines.append(lines[idx])
            idx += 1

        if idx >= len(lines):
            raise ValueError(f"unterminated artifact block for {path}")

        content = "\n".join(content_lines).rstrip("\n") + "\n"
        blocks[path] = ArtifactBlock(path=path, content=content)
        idx += 1

    return blocks


def require_blocks(text: str, required_paths: list[str]) -> dict[str, ArtifactBlock]:
    blocks = parse_artifact_blocks(text)
    missing = [path for path in required_paths if path not in blocks]
    if missing:
        raise ValueError(f"missing artifact blocks: {', '.join(missing)}")
    return blocks
