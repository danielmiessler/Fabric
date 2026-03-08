#!/usr/bin/env python3
"""Ingest Zoom caption text into deterministic segment artifacts.

No API usage. Pure local processing.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import re
import sys
from dataclasses import dataclass
from pathlib import Path


HEADER_RE = re.compile(r"^\[(?P<speaker>[^\]]+)\]\s+(?P<timestamp>\d{2}:\d{2}:\d{2})\s*$")

NOISE_EXACT = {
    "cool",
    "cool.",
    "mm-hmm",
    "mm-hmm.",
    "okay",
    "okay.",
    "ok",
    "ok.",
    "yes",
    "yes.",
    "right",
    "right.",
    "hmm",
    "hmm.",
}

NOISE_CONTAINS = (
    "can you hear me",
    "can you see me",
    "testing",
    "let me share my screen",
    "is my screen visible",
    "hello my testing",
)


@dataclass
class Segment:
    raw_speaker: str
    timestamp: str
    raw_text: str


def slugify(value: str) -> str:
    slug = re.sub(r"[^a-zA-Z0-9]+", "-", value.strip().lower()).strip("-")
    return slug or "session"


def normalize_speaker(raw: str) -> str:
    cleaned = re.sub(r"\s+", " ", raw).strip()
    return " ".join(part.capitalize() for part in cleaned.split(" "))


def classify_segment(text: str) -> tuple[str, str]:
    normalized = re.sub(r"\s+", " ", text.strip().lower())
    if not normalized:
        return "noise", "empty_text"
    if normalized in NOISE_EXACT and len(normalized.split()) <= 3:
        return "noise", "low_value_ack"
    for phrase in NOISE_CONTAINS:
        if phrase in normalized:
            return "noise", f"contains:{phrase}"
    return "content", ""


def resolve_input(path: Path) -> Path:
    if path.is_file():
        return path
    if path.is_dir():
        exact = path / "meeting_saved_closed_caption.txt"
        if exact.exists():
            return exact
        txts = sorted(path.glob("*.txt"))
        if len(txts) == 1:
            return txts[0]
        raise FileNotFoundError(
            f"Could not resolve unique transcript .txt in directory: {path}"
        )
    raise FileNotFoundError(f"Input path not found: {path}")


def parse_segments(lines: list[str]) -> list[Segment]:
    segments: list[Segment] = []
    i = 0
    while i < len(lines):
        line = lines[i].rstrip("\n")
        m = HEADER_RE.match(line.strip())
        if not m:
            if segments and line.strip():
                segments[-1].raw_text = (segments[-1].raw_text + " " + line.strip()).strip()
            i += 1
            continue

        raw_speaker = m.group("speaker").strip()
        timestamp = m.group("timestamp")
        i += 1

        text_lines: list[str] = []
        while i < len(lines):
            current = lines[i].rstrip("\n")
            if HEADER_RE.match(current.strip()):
                break
            if current.strip():
                text_lines.append(current.strip())
            i += 1

        segments.append(Segment(raw_speaker=raw_speaker, timestamp=timestamp, raw_text=" ".join(text_lines).strip()))

    return segments


def write_jsonl(path: Path, rows: list[dict]) -> None:
    with path.open("w", encoding="utf-8") as f:
        for row in rows:
            f.write(json.dumps(row, ensure_ascii=True) + "\n")


def main() -> int:
    parser = argparse.ArgumentParser(description="Ingest Zoom captions into segment artifacts.")
    parser.add_argument("input_path", help="Path to transcript file or session folder")
    parser.add_argument(
        "--output-dir",
        help="Output directory (default: <session>/.pipeline)",
        default=None,
    )
    parser.add_argument("--session-id", help="Explicit session id", default=None)
    args = parser.parse_args()

    input_path = resolve_input(Path(args.input_path).expanduser().resolve())
    session_dir = input_path.parent
    output_dir = Path(args.output_dir).expanduser().resolve() if args.output_dir else session_dir / ".pipeline"
    output_dir.mkdir(parents=True, exist_ok=True)

    with input_path.open("r", encoding="utf-8") as f:
        lines = f.readlines()

    parsed = parse_segments(lines)
    if not parsed:
        print(f"No segments parsed from {input_path}", file=sys.stderr)
        return 2

    session_id = args.session_id or slugify(session_dir.name)
    ledger_rows: list[dict] = []
    manifest_rows: list[dict] = []

    for idx, seg in enumerate(parsed, start=1):
        segment_id = f"{session_id}-seg-{idx:05d}"
        seg_type, noise_reason = classify_segment(seg.raw_text)
        norm_speaker = normalize_speaker(seg.raw_speaker)
        text_hash = hashlib.sha1(seg.raw_text.encode("utf-8")).hexdigest()[:12]

        ledger = {
            "session_id": session_id,
            "segment_id": segment_id,
            "segment_index": idx,
            "timestamp": seg.timestamp,
            "raw_speaker": seg.raw_speaker,
            "normalized_speaker": norm_speaker,
            "raw_text": seg.raw_text,
            "type": seg_type,
            "noise_reason": noise_reason,
            "word_count": len(seg.raw_text.split()),
            "text_sha1_12": text_hash,
            "source_file": str(input_path),
        }
        manifest = {
            "session_id": session_id,
            "segment_id": segment_id,
            "segment_index": idx,
            "timestamp": seg.timestamp,
            "speaker": norm_speaker,
            "type": seg_type,
        }
        ledger_rows.append(ledger)
        manifest_rows.append(manifest)

    ledger_path = output_dir / "segment_ledger.jsonl"
    manifest_path = output_dir / "segment_manifest.jsonl"
    write_jsonl(ledger_path, ledger_rows)
    write_jsonl(manifest_path, manifest_rows)

    content_count = sum(1 for row in ledger_rows if row["type"] == "content")
    noise_count = len(ledger_rows) - content_count

    print(f"input: {input_path}")
    print(f"session_id: {session_id}")
    print(f"segments_total: {len(ledger_rows)}")
    print(f"segments_content: {content_count}")
    print(f"segments_noise: {noise_count}")
    print(f"wrote: {ledger_path}")
    print(f"wrote: {manifest_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

