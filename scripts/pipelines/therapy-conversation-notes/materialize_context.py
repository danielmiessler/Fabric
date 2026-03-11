#!/usr/bin/env python3
"""Build optional context artifacts for therapy-conversation-notes."""

from __future__ import annotations

import argparse
import json
from dataclasses import asdict, dataclass
from pathlib import Path

from common import discover_context_files, load_session_context, run_dir


MAX_CONTEXT_FILES = 8
MAX_CONTEXT_CHARS_PER_FILE = 4000
MAX_TOTAL_CONTEXT_CHARS = 20000
MAX_PDF_PAGES = 12


@dataclass
class ContextEntry:
    path: str
    kind: str
    status: str
    excerpt: str
    source_chars: int
    included_chars: int
    truncated: bool
    error: str = ""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


def append_error(existing: str, new_note: str) -> str:
    if not existing:
        return new_note
    return f"{existing}; {new_note}"


def read_text_excerpt(path: Path, limit: int) -> tuple[str, bool]:
    with path.open("r", encoding="utf-8", errors="ignore") as file:
        content = file.read(limit + 1)
    if len(content) <= limit:
        return content, False
    return content[:limit], True


def extract_pdf_text(path: Path, per_file_limit: int) -> tuple[str, str, str, int, bool]:
    try:
        from pypdf import PdfReader  # type: ignore
    except Exception:
        return "", "metadata_only", "pypdf not installed; only file metadata is available", 0, False

    try:
        reader = PdfReader(str(path))
        chunks: list[str] = []
        source_chars = 0
        truncated = False
        for page in reader.pages[:MAX_PDF_PAGES]:
            page_text = (page.extract_text() or "").strip()
            if not page_text:
                continue

            delimiter = "\n\n" if chunks else ""
            source_chars += len(delimiter) + len(page_text)
            if source_chars > per_file_limit:
                truncated = True

            remaining = per_file_limit - sum(len(chunk) for chunk in chunks)
            if remaining > 0:
                chunks.append((delimiter + page_text)[:remaining])
            if truncated:
                break

        text = "".join(chunks).strip()
        if not text:
            return "", "metadata_only", "PDF contained no extractable text", 0, False

        error = ""
        if truncated:
            error = f"excerpt truncated to {per_file_limit} chars"
        return text, "extracted", error, source_chars if source_chars > 0 else len(text), truncated
    except Exception as err:
        return "", "metadata_only", f"PDF extraction failed: {err}", 0, False


def extract_context_entry(path: Path) -> ContextEntry:
    suffix = path.suffix.lower()
    if suffix in {".md", ".txt"}:
        text, truncated = read_text_excerpt(path, MAX_CONTEXT_CHARS_PER_FILE)
        excerpt = text.strip()
        status = "extracted" if excerpt else "metadata_only"
        source_chars = len(excerpt) + (1 if truncated else 0)
        error = ""
        if not excerpt:
            error = "file was empty"
        elif truncated:
            error = f"excerpt truncated to {MAX_CONTEXT_CHARS_PER_FILE} chars"
        return ContextEntry(
            path=str(path),
            kind="markdown" if suffix == ".md" else "text",
            status=status,
            excerpt=excerpt,
            source_chars=source_chars,
            included_chars=len(excerpt),
            truncated=truncated and bool(excerpt),
            error=error,
        )
    if suffix == ".pdf":
        text, status, error, source_chars, truncated = extract_pdf_text(path, MAX_CONTEXT_CHARS_PER_FILE)
        return ContextEntry(
            path=str(path),
            kind="pdf",
            status=status,
            excerpt=text,
            source_chars=source_chars,
            included_chars=len(text),
            truncated=truncated and bool(text),
            error=error,
        )
    return ContextEntry(
        path=str(path),
        kind=suffix.lstrip("."),
        status="ignored",
        excerpt="",
        source_chars=0,
        included_chars=0,
        truncated=False,
        error="unsupported type",
    )


def apply_total_char_limit(entries: list[ContextEntry], total_char_limit: int) -> int:
    included_total = 0
    for entry in entries:
        entry.included_chars = len(entry.excerpt)
        if not entry.excerpt:
            continue

        remaining = total_char_limit - included_total
        if remaining <= 0:
            entry.excerpt = ""
            entry.included_chars = 0
            entry.truncated = True
            if entry.status == "extracted":
                entry.status = "metadata_only"
            entry.error = append_error(
                entry.error,
                f"excerpt omitted because total context cap ({total_char_limit} chars) was reached",
            )
            continue

        if len(entry.excerpt) > remaining:
            entry.excerpt = entry.excerpt[:remaining]
            entry.included_chars = len(entry.excerpt)
            entry.truncated = True
            entry.error = append_error(
                entry.error,
                f"excerpt truncated to fit total context cap ({total_char_limit} chars)",
            )

        included_total += entry.included_chars
    return included_total


def render_context_pack(entries: list[ContextEntry], warnings: list[str]) -> str:
    lines = [
        "# Optional Therapeutic Context",
        "",
        "This context is optional and must be treated as supporting reference only.",
        "",
    ]
    if warnings:
        lines.extend(["## Context Warnings", ""])
        lines.extend(f"- {warning}" for warning in warnings)
        lines.append("")

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
                f"- included_chars: `{entry.included_chars}`",
            ]
        )
        if entry.truncated:
            lines.append("- truncated: `true`")
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

    discovered_context_files = discover_context_files(
        Path(session_context.context_root),
        Path(session_context.input_path),
    )
    selected_context_files = discovered_context_files[:MAX_CONTEXT_FILES]
    skipped_by_file_cap = len(discovered_context_files) - len(selected_context_files)

    entries = [extract_context_entry(path) for path in selected_context_files]
    included_context_chars = apply_total_char_limit(entries, MAX_TOTAL_CONTEXT_CHARS)

    truncated_entries = sum(1 for entry in entries if entry.truncated)
    total_cap_events = sum(1 for entry in entries if "total context cap" in entry.error)

    warnings: list[str] = []
    if skipped_by_file_cap > 0:
        warnings.append(
            f"processed first {MAX_CONTEXT_FILES} context files in deterministic path order; "
            f"skipped {skipped_by_file_cap} additional files"
        )
    if total_cap_events > 0:
        warnings.append(
            f"total context excerpt size capped at {MAX_TOTAL_CONTEXT_CHARS} chars; "
            f"{total_cap_events} file(s) were truncated or omitted"
        )
    if truncated_entries > total_cap_events:
        warnings.append(
            f"{truncated_entries - total_cap_events} file(s) exceeded per-file limit "
            f"of {MAX_CONTEXT_CHARS_PER_FILE} chars"
        )

    context_pack = render_context_pack(entries, warnings)
    manifest_payload = {
        "limits": {
            "max_files": MAX_CONTEXT_FILES,
            "max_chars_per_file": MAX_CONTEXT_CHARS_PER_FILE,
            "max_total_chars": MAX_TOTAL_CONTEXT_CHARS,
        },
        "summary": {
            "discovered_files": len(discovered_context_files),
            "processed_files": len(selected_context_files),
            "skipped_by_file_cap": skipped_by_file_cap,
            "included_chars": included_context_chars,
            "truncated_entries": truncated_entries,
        },
        "warnings": warnings,
        "entries": [asdict(entry) for entry in entries],
    }
    manifest_json = json.dumps(manifest_payload, ensure_ascii=True, indent=2) + "\n"

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
