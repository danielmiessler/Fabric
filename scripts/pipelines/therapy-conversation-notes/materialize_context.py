#!/usr/bin/env python3
"""Build optional context artifacts for therapy-conversation-notes."""

from __future__ import annotations

import argparse
import json
from dataclasses import asdict, dataclass
from pathlib import Path

from common import discover_context_files, load_session_context, run_dir


@dataclass
class ContextEntry:
    path: str
    kind: str
    status: str
    excerpt: str
    error: str = ""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


def extract_pdf_text(path: Path) -> tuple[str, str, str]:
    try:
        from pypdf import PdfReader  # type: ignore
    except Exception:
        return "", "metadata_only", "pypdf not installed; only file metadata is available"

    try:
        reader = PdfReader(str(path))
        chunks: list[str] = []
        for page in reader.pages[:12]:
            chunks.append((page.extract_text() or "").strip())
        text = "\n\n".join(chunk for chunk in chunks if chunk).strip()
        if not text:
            return "", "metadata_only", "PDF contained no extractable text"
        return text, "extracted", ""
    except Exception as err:
        return "", "metadata_only", f"PDF extraction failed: {err}"


def extract_context_entry(path: Path) -> ContextEntry:
    suffix = path.suffix.lower()
    if suffix in {".md", ".txt"}:
        text = path.read_text(encoding="utf-8", errors="ignore").strip()
        excerpt = text[:4000]
        status = "extracted" if excerpt else "metadata_only"
        return ContextEntry(
            path=str(path),
            kind="markdown" if suffix == ".md" else "text",
            status=status,
            excerpt=excerpt,
            error="" if excerpt else "file was empty",
        )
    if suffix == ".pdf":
        text, status, error = extract_pdf_text(path)
        return ContextEntry(
            path=str(path),
            kind="pdf",
            status=status,
            excerpt=text[:4000],
            error=error,
        )
    return ContextEntry(path=str(path), kind=suffix.lstrip("."), status="ignored", excerpt="", error="unsupported type")


def render_context_pack(entries: list[ContextEntry]) -> str:
    lines = [
        "# Optional Therapeutic Context",
        "",
        "This context is optional and must be treated as supporting reference only.",
        "",
    ]
    if not entries:
        lines.extend(["- No optional context files were discovered.", ""])
        return "\n".join(lines).rstrip() + "\n"

    for idx, entry in enumerate(entries, start=1):
        lines.extend(
            [
                f"## Context Source {idx}",
                f"- path: `{entry.path}`",
                f"- kind: `{entry.kind}`",
                f"- status: `{entry.status}`",
            ]
        )
        if entry.error:
            lines.append(f"- note: {entry.error}")
        lines.append("")
        if entry.excerpt:
            lines.extend(
                [
                    "```text",
                    entry.excerpt.rstrip(),
                    "```",
                    "",
                ]
            )
    return "\n".join(lines).rstrip() + "\n"


def main() -> None:
    args = parse_args()
    session_context = load_session_context(Path(args.session_context))

    context_files = discover_context_files(
        Path(session_context.context_root),
        Path(session_context.input_path),
    )
    entries = [extract_context_entry(path) for path in context_files]

    context_pack = render_context_pack(entries)
    manifest_json = json.dumps({"entries": [asdict(entry) for entry in entries]}, ensure_ascii=True, indent=2) + "\n"

    current_run_dir = run_dir()
    current_run_dir.mkdir(parents=True, exist_ok=True)
    stable_pipeline_dir = Path(session_context.stable_pipeline_dir)
    stable_pipeline_dir.mkdir(parents=True, exist_ok=True)

    run_context_pack_path = current_run_dir / "context_pack.md"
    run_context_manifest_path = current_run_dir / "context_manifest.json"
    stable_context_pack_path = stable_pipeline_dir / "context_pack.md"
    stable_context_manifest_path = stable_pipeline_dir / "context_manifest.json"

    run_context_pack_path.write_text(context_pack, encoding="utf-8")
    run_context_manifest_path.write_text(manifest_json, encoding="utf-8")
    stable_context_pack_path.write_text(context_pack, encoding="utf-8")
    stable_context_manifest_path.write_text(manifest_json, encoding="utf-8")


if __name__ == "__main__":
    main()
