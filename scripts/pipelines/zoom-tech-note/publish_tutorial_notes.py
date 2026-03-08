#!/usr/bin/env python3
"""Normalize learner-facing tutorial notes across sessions.

This script intentionally does NOT modify raw transcript files.
It updates tutorial presentation artifacts only:
- final_notes.md (title/H1 normalization + source-tag sanitization)
- creates a published renamed tutorial file per session
- rewrites bootcamp_index.md to point to the published tutorial file
"""

from __future__ import annotations

import argparse
import re
from dataclasses import dataclass
from datetime import datetime
from pathlib import Path


@dataclass
class SessionMeta:
    session_dir: Path
    domain: str
    topic: str
    date_display: str  # DD/MM/YYYY
    class_number: int


DOMAIN_FILE_LABEL = {
    "AI/ML": "AI-ML",
    "WebDev": "WebDev",
    "Web3": "Web3",
}


KNOWN_OVERRIDES: list[tuple[str, str, str, str | None]] = [
    ("Fast-tracking the AI course", "AI/ML", "Fast-tracking the Course of AI", "18/01/2026"),
    ("Neural Networks from Scratch", "AI/ML", "Neural Networks from Scratch", "24/01/2026"),
    ("Transformers_ Part 1", "AI/ML", "Transformers Part 1", "31/01/2026"),
    ("Transformers_ Part 2", "AI/ML", "Transformers Part 2", "06/02/2026"),
    (
        "Promises, Callbacks, CPU vs IO Tasks",
        "WebDev",
        "Promises, Callbacks, CPU vs IO Tasks",
        None,
    ),
    ("Web Development _ Promises _", "WebDev", "Promises", None),
    (
        "Writing promises and async await",
        "WebDev",
        "Writing Promises and Async/Await",
        None,
    ),
    ("Introduction to Blockchains", "Web3", "Introduction to Blockchains", None),
    ("Wallets and Private Keys", "Web3", "Wallets and Private Keys", None),
    ("Token Program", "Web3", "Token Program", None),
]


def parse_folder_date(folder_name: str) -> str:
    m = re.match(r"^(\d{4})-(\d{2})-(\d{2})\s+\d{2}\.\d{2}\.\d{2}", folder_name)
    if not m:
        return datetime.now().strftime("%d/%m/%Y")
    yyyy, mm, dd = m.groups()
    return f"{dd}/{mm}/{yyyy}"


def infer_domain_topic_date(folder_name: str) -> tuple[str, str, str]:
    for needle, domain, topic, date_override in KNOWN_OVERRIDES:
        if needle in folder_name:
            date_display = date_override or parse_folder_date(folder_name)
            return domain, topic, date_display

    date_display = parse_folder_date(folder_name)
    raw = re.sub(r"^\d{4}-\d{2}-\d{2}\s+\d{2}\.\d{2}\.\d{2}\s*", "", folder_name)
    raw = raw.replace("_", " ")
    raw = re.sub(r"\s+", " ", raw).strip()
    raw = re.sub(r"\s*Bootcamp\s*1\.0\s*$", "", raw, flags=re.I)
    raw = re.sub(r"^Week\s*\d+\s*[-:]?\s*", "", raw, flags=re.I)

    if re.search(r"AI\s*&?\s*ML|AI and ML", raw, flags=re.I):
        domain = "AI/ML"
        topic = re.sub(r"AI\s*&?\s*ML", "", raw, flags=re.I).strip(" -|_") or "AI/ML Session"
    elif "Web Development" in raw:
        domain = "WebDev"
        topic = raw.replace("Web Development", "").strip(" -|_") or "Web Development Session"
    elif "Web3" in raw:
        domain = "Web3"
        topic = raw.replace("Web3", "").strip(" -|_") or "Web3 Session"
    else:
        domain = "WebDev"
        topic = raw or "Session"

    topic = re.sub(r"\s+", " ", topic)
    return domain, topic, date_display


def strip_source_tags(text: str) -> str:
    return re.sub(r"\s*\[source:\s*[^\]]+\]", "", text)


def normalize_frontmatter_and_h1(content: str, title: str) -> str:
    content = strip_source_tags(content)
    if content.startswith("---\n"):
        end = content.find("\n---\n", 4)
        if end != -1:
            fm = content[4:end]
            body = content[end + 5 :]
            if re.search(r"^title:\s*\".*\"\s*$", fm, flags=re.M):
                fm = re.sub(r"^title:\s*\".*\"\s*$", f'title: "{title}"', fm, flags=re.M)
            elif re.search(r"^title:\s*.*$", fm, flags=re.M):
                fm = re.sub(r"^title:\s*.*$", f'title: "{title}"', fm, flags=re.M)
            else:
                fm = f'title: "{title}"\n' + fm
            content = f"---\n{fm}\n---\n{body}"
    else:
        content = (
            "---\n"
            f'title: "{title}"\n'
            "type: \"tutorial-note\"\n"
            "status: \"ready\"\n"
            "---\n\n"
            + content
        )

    if re.search(r"^#\s+", content, flags=re.M):
        content = re.sub(r"^#\s+.*$", f"# 🎓 {title}", content, count=1, flags=re.M)
    else:
        content = f"# 🎓 {title}\n\n" + content

    return content


def safe_filename(value: str) -> str:
    value = value.replace("/", "-")
    value = value.replace(":", "-")
    value = re.sub(r"\s+", " ", value).strip()
    return value


def collect_sessions(root: Path, session_filter: Path | None) -> list[SessionMeta]:
    sessions: list[tuple[Path, str, str, str]] = []
    direct_final_note = root / "final_notes.md"
    if direct_final_note.exists():
        domain, topic, date_display = infer_domain_topic_date(root.name)
        sessions.append((root, domain, topic, date_display))
    for final_note in sorted(root.glob("*/final_notes.md")):
        session_dir = final_note.parent
        domain, topic, date_display = infer_domain_topic_date(session_dir.name)
        sessions.append((session_dir, domain, topic, date_display))

    by_domain: dict[str, list[tuple[Path, str, str]]] = {}
    for sdir, domain, topic, date_display in sessions:
        by_domain.setdefault(domain, []).append((sdir, topic, date_display))

    out: list[SessionMeta] = []
    for domain, items in by_domain.items():
        items.sort(key=lambda x: x[0].name)
        for idx, (sdir, topic, date_display) in enumerate(items, start=1):
            out.append(SessionMeta(sdir, domain, topic, date_display, idx))

    out.sort(key=lambda m: m.session_dir.name)
    if session_filter:
        out = [m for m in out if m.session_dir == session_filter]
    return out


def write_bootcamp_index(session_dir: Path, title: str, published_filename: str, domain: str) -> None:
    note_name = Path(published_filename).stem
    content = (
        "# Bootcamp Index\n\n"
        f"- [[{note_name}|{title}]]\n"
        f"- Domain: {domain}\n"
    )
    (session_dir / "bootcamp_index.md").write_text(content, encoding="utf-8")


def process_session(meta: SessionMeta, dry_run: bool = False) -> tuple[Path, Path, str]:
    title = f"{meta.domain} Class {meta.class_number:02d} [{meta.date_display}] - {meta.topic}"
    final_note = meta.session_dir / "final_notes.md"
    content = final_note.read_text(encoding="utf-8")
    updated = normalize_frontmatter_and_h1(content, title)

    file_domain = DOMAIN_FILE_LABEL.get(meta.domain, meta.domain.replace("/", "-"))
    date_for_file = meta.date_display.replace("/", "-")
    published_name = safe_filename(f"{file_domain} Class {meta.class_number:02d} [{date_for_file}] - {meta.topic}.md")
    published_path = meta.session_dir / published_name

    if not dry_run:
        final_note.write_text(updated, encoding="utf-8")
        published_path.write_text(updated, encoding="utf-8")
        write_bootcamp_index(meta.session_dir, title, published_name, meta.domain)

    return final_note, published_path, title


def main() -> int:
    parser = argparse.ArgumentParser(description="Publish normalized tutorial note names and headings.")
    parser.add_argument("--root", default=".", help="Root directory containing session folders")
    parser.add_argument("--session-dir", default=None, help="Optional single session directory")
    parser.add_argument("--dry-run", action="store_true", help="Show planned updates without writing")
    args = parser.parse_args()

    root = Path(args.root).expanduser().resolve()
    session_filter = Path(args.session_dir).expanduser().resolve() if args.session_dir else None

    metas = collect_sessions(root, session_filter)
    if not metas:
        print("No session final_notes.md files found.")
        return 2

    changed: list[tuple[Path, Path, str]] = []
    for meta in metas:
        changed.append(process_session(meta, dry_run=args.dry_run))

    print(f"processed_sessions: {len(changed)}")
    for final_note, published_path, title in changed:
        print(f"- final: {final_note}")
        print(f"  published: {published_path}")
        print(f"  title: {title}")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
