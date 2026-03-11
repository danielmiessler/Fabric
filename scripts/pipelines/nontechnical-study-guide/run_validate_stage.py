#!/usr/bin/env python3
"""Validate nontechnical-study-guide outputs."""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path

from common import load_session_context, run_dir


REQUIRED_HEADINGS = [
    "## Session Focus",
    "## Context You Need Before This",
    "## Big Ideas",
    "## What Was Actually Said",
    "## Why It Matters",
    "## Intuition and Real-Life Examples",
    "## Argument or Theme Map",
    "## Difficult Terms Explained Simply",
    "## Misunderstandings to Avoid",
    "## Reflection Questions",
    "## Revision Summary",
    "## Connections to Other Topics",
]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


def section_body(final_notes: str, heading: str) -> str:
    marker = "\n" + heading + "\n"
    start = final_notes.find(marker)
    if start < 0:
        if final_notes.startswith(heading + "\n"):
            start = 0
            offset = len(heading) + 1
        else:
            return ""
    else:
        start += 1
        offset = len(heading) + 1

    body_start = start + offset
    next_heading = final_notes.find("\n## ", body_start)
    if next_heading < 0:
        return final_notes[body_start:]
    return final_notes[body_start:next_heading]


def count_bullets(markdown: str) -> int:
    return len(re.findall(r"(?m)^\s*[-*]\s+", markdown))


def validate_outputs(session_context) -> list[str]:
    issues: list[str] = []

    semantic_map_path = Path(session_context.stable_pipeline_dir) / "semantic_map.json"
    if not semantic_map_path.exists():
        issues.append(f"missing semantic map: {semantic_map_path}")
    else:
        try:
            data = json.loads(semantic_map_path.read_text(encoding="utf-8"))
        except json.JSONDecodeError as err:
            issues.append(f"semantic_map.json is not valid JSON: {err}")
            data = None
        if data is not None and not isinstance(data, dict):
            issues.append("semantic_map.json must be a JSON object")

    final_output_path = Path(session_context.final_output_path)
    if not final_output_path.exists():
        issues.append(f"missing final notes: {final_output_path}")
        return issues

    final_notes = final_output_path.read_text(encoding="utf-8")
    if not final_notes.lstrip().startswith("# 📝 "):
        issues.append("final notes must start with a top-level '# 📝 ' heading")
    for heading in REQUIRED_HEADINGS:
        if heading not in final_notes:
            issues.append(f"missing heading: {heading}")

    if final_notes.count("```mermaid") < 1:
        issues.append("expected at least 1 Mermaid diagram")
    if "<<<BEGIN_ARTIFACT:" in final_notes or "<<<END_ARTIFACT>>>" in final_notes:
        issues.append("final notes contain raw artifact markers")

    examples_body = section_body(final_notes, "## Intuition and Real-Life Examples")
    if count_bullets(examples_body) < 3:
        issues.append("Intuition and Real-Life Examples must include at least 3 examples")

    reflections_body = section_body(final_notes, "## Reflection Questions")
    if count_bullets(reflections_body) < 4:
        issues.append("Reflection Questions must include at least 4 questions")

    terms_body = section_body(final_notes, "## Difficult Terms Explained Simply")
    if count_bullets(terms_body) < 3:
        issues.append("Difficult Terms Explained Simply must include at least 3 terms")

    return issues


def write_report(report_path: Path, issues: list[str]) -> None:
    lines = ["# Validation Report", ""]
    if issues:
        lines.extend(["status: fail", ""])
        lines.extend(f"- {issue}" for issue in issues)
    else:
        lines.extend(["status: pass", "", "- All required nontechnical study-guide outputs are present and valid."])
    report_path.parent.mkdir(parents=True, exist_ok=True)
    report_path.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")


def main() -> None:
    args = parse_args()
    session_context = load_session_context(Path(args.session_context))
    issues = validate_outputs(session_context)

    current_run_dir = run_dir()
    run_report_path = current_run_dir / "validation_report.md"
    stable_report_path = Path(session_context.stable_pipeline_dir) / "validation_report.md"
    write_report(run_report_path, issues)
    write_report(stable_report_path, issues)

    if issues:
        raise SystemExit("\n".join(issues))


if __name__ == "__main__":
    main()
