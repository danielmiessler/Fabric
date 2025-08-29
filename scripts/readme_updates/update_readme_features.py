#!/usr/bin/env python3
"""
Generate the '### Recent Major Features' markdown section for README from the changelog SQLite DB.

- Connects to cmd/generate_changelog/changelog.db
- Extracts version, date, and AI summaries from the 'versions' table
- Heuristically filters for feature/improvement items (excludes CI/CD/docs/bug fixes)
- Formats output to match README style:
  - [vX.Y.Z](https://github.com/danielmiessler/fabric/releases/tag/vX.Y.Z) (Aug 14, 2025) — **Feature Name**: Short description

Usage:
  python scripts/readme_updates/update_readme_features.py --limit 20
"""

import argparse
import sqlite3
from pathlib import Path
from datetime import datetime
import re
import sys
from typing import List, Optional, Tuple

# Heuristics for filtering feature-related lines
EXCLUDE_RE = re.compile(
    r"(?i)\b(fix|bug|hotfix|ci|cd|pipeline|chore|docs|doc|readme|refactor|style|typo|"
    "test|tests|bump|deps|dependency|merge|revert|format|lint|build|release\b|prepare|"
    "codeowners|coverage|security)\b"
)
INCLUDE_RE = re.compile(
    r"(?i)\b(new|feature|feat|add|added|introduce|enable|support|improve|enhance|"
    "performance|speed|option|flag|argument|parameter|integration|provider|search|tts|"
    "audio|model|cli|ui|web|oauth|sync|database|notifications|desktop|reasoning|thinking)\b"
)


def parse_args():
    """Parse command-line arguments."""
    p = argparse.ArgumentParser(
        description="Generate README 'Recent Major Features' markdown from changelog DB."
    )
    p.add_argument(
        "--limit", type=int, default=20, help="Maximum number of releases to include."
    )
    p.add_argument(
        "--db",
        type=str,
        default=None,
        help="Optional path to changelog.db (defaults to repo cmd/generate_changelog/changelog.db)",
    )
    return p.parse_args()


def repo_root() -> Path:
    """Get the repository root directory."""
    # scripts/readme_updates/update_readme_features.py -> repo root is parent.parent.parent
    return Path(__file__).resolve().parent.parent.parent


def db_path(args) -> Path:
    """Determine the database path."""
    if args.db:
        return Path(args.db).expanduser().resolve()
    return repo_root() / "cmd" / "generate_changelog" / "changelog.db"


def connect(dbfile: Path):
    """Connect to the SQLite database."""
    if not dbfile.exists():
        print(f"ERROR: changelog database not found: {dbfile}", file=sys.stderr)
        sys.exit(1)
    return sqlite3.connect(str(dbfile))


def normalize_version(name: str) -> str:
    """Ensure version string starts with 'v'."""
    n = str(name).strip()
    return n if n.startswith("v") else f"v{n}"


def parse_date(value) -> str:
    """Parse various date formats and return formatted string."""
    if value is None:
        return "(Unknown date)"

    # Handle the ISO format with timezone from the database
    s = str(value).strip()

    # Try to parse the ISO format with timezone
    if "+" in s or "T" in s:
        # Remove timezone info and microseconds for simpler parsing
        s_clean = s.split("+")[0].split(".")[0]
        try:
            dt = datetime.strptime(s_clean, "%Y-%m-%d %H:%M:%S")
            return dt.strftime("%b %d, %Y").replace(" 0", " ")
        except ValueError:
            pass

    # Fallback formats
    fmts = [
        "%Y-%m-%d",
        "%Y-%m-%d %H:%M:%S",
        "%Y-%m-%dT%H:%M:%S",
        "%Y/%m/%d",
        "%m/%d/%Y",
    ]

    for fmt in fmts:
        try:
            dt = datetime.strptime(s, fmt)
            return dt.strftime("%b %d, %Y").replace(" 0", " ")
        except ValueError:
            continue

    # Return original if we can't parse it
    return f"({s})"


def split_summary(text: str) -> List[str]:
    """Split AI summary into individual lines/bullets."""
    if not text:
        return []

    lines = []
    # Split by newlines first
    for line in text.split("\n"):
        line = line.strip()
        if not line:
            continue
        # Remove markdown headers
        line = re.sub(r"^#+\s+", "", line)
        # Remove PR links and author info
        line = re.sub(
            r"^PR\s+\[#\d+\]\([^)]+\)\s+by\s+\[[^\]]+\]\([^)]+\):\s*", "", line
        )
        # Remove bullet points
        line = re.sub(r"^[-*•]\s+", "", line)
        if line:
            lines.append(line)

    return lines


def is_feature_line(line: str) -> bool:
    """Check if a line describes a feature/improvement (not a bug fix or CI/CD)."""
    line_lower = line.lower()

    # Strong exclusions first
    if any(
        word in line_lower
        for word in ["chore:", "fix:", "docs:", "test:", "ci:", "build:", "refactor:"]
    ):
        return False

    if EXCLUDE_RE.search(line):
        return False

    return bool(INCLUDE_RE.search(line))


def extract_title_desc(line: str) -> Tuple[str, str]:
    """Extract title and description from a feature line."""
    # Remove any markdown formatting
    line = re.sub(r"\*\*([^*]+)\*\*", r"\1", line)

    # Look for colon separator first
    if ":" in line:
        parts = line.split(":", 1)
        if len(parts) == 2:
            title = parts[0].strip()
            desc = parts[1].strip()

            # Clean up the title
            title = (
                title.replace("Introduce ", "")
                .replace("Enable ", "")
                .replace("Add ", "")
            )
            title = title.replace("Implement ", "").replace("Support ", "")

            # Make title more concise
            if len(title) > 30:
                # Try to extract key words
                key_words = []
                for word in title.split():
                    if word[0].isupper() or "-" in word or "_" in word:
                        key_words.append(word)
                if key_words:
                    title = " ".join(key_words[:3])

            return (title, desc)

    # Fallback: use first sentence as description
    sentences = re.split(r"[.!?]\s+", line)
    if sentences:
        desc = sentences[0].strip()
        # Extract a title from the description
        if "thinking" in desc.lower():
            return ("AI Reasoning", desc)
        elif "token" in desc.lower() and "context" in desc.lower():
            return ("Extended Context", desc)
        elif "curl" in desc.lower() or "install" in desc.lower():
            return ("Easy Setup", desc)
        elif "vendor" in desc.lower() or "model" in desc.lower():
            return ("Model Management", desc)
        elif "notification" in desc.lower():
            return ("Desktop Notifications", desc)
        elif "tts" in desc.lower() or "speech" in desc.lower():
            return ("Text-to-Speech", desc)
        elif "oauth" in desc.lower() or "auth" in desc.lower():
            return ("OAuth Auto-Auth", desc)
        elif "search" in desc.lower() and "web" in desc.lower():
            return ("Web Search", desc)
        else:
            # Generic title from first significant words
            words = desc.split()[:2]
            title = " ".join(words)
            return (title, desc)

    return ("Feature", line)


def pick_feature(ai_summary: str) -> Optional[Tuple[str, str]]:
    """Pick the best feature line from the AI summary."""
    lines = split_summary(ai_summary)

    # Look for the first feature line
    for line in lines:
        if is_feature_line(line):
            title, desc = extract_title_desc(line)
            # Clean up description - remove redundant info
            desc = desc[:200] if len(desc) > 200 else desc  # Limit length
            return (title, desc)

    return None


def build_item(
    version: str, date_str: str, feature_title: str, feature_desc: str
) -> str:
    """Build a markdown list item for a release."""
    url = f"https://github.com/danielmiessler/fabric/releases/tag/{version}"
    return f"- [{version}]({url}) ({date_str}) — **{feature_title}**: {feature_desc}"


def main():
    """Main function."""
    args = parse_args()
    dbfile = db_path(args)
    conn = connect(dbfile)
    cur = conn.cursor()

    # Query the database
    cur.execute("SELECT name, date, ai_summary FROM versions ORDER BY date DESC")
    rows = cur.fetchall()

    items = []
    for name, date, summary in rows:
        version = normalize_version(name)
        date_fmt = parse_date(date)
        feat = pick_feature(summary or "")

        if not feat:
            continue

        title, desc = feat
        items.append(build_item(version, date_fmt, title, desc))

        if len(items) >= args.limit:
            break

    conn.close()

    # Output the markdown
    print("### Recent Major Features")
    print()
    for item in items:
        print(item)


if __name__ == "__main__":
    main()
