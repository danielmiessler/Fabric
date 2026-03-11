#!/usr/bin/env python3
"""Validate note-enhancement outputs."""

from __future__ import annotations

import argparse
import re
from pathlib import Path

from common import load_session_context, run_dir


REQUIRED_HEADINGS = [
    "# ✨ Enhanced Notes",
    "## Improved Summary",
    "## Key Takeaways",
    "## Clarifications Added",
    "## Suggested Next Questions",
]

REQUIRED_EDIT_LOG_HEADINGS = [
    "# Edit Log",
    "## Structural Changes",
    "## Language Tightening",
    "## Assumptions Avoided",
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
    stable_pipeline_dir = Path(session_context.stable_pipeline_dir)

    final_output_path = Path(session_context.final_output_path)
    if not final_output_path.exists():
        issues.append(f"missing enhanced notes: {final_output_path}")
        return issues

    final_notes = final_output_path.read_text(encoding="utf-8")
    for heading in REQUIRED_HEADINGS:
        if heading not in final_notes:
            issues.append(f"missing heading: {heading}")

    key_takeaways_body = section_body(final_notes, "## Key Takeaways")
    if count_bullets(key_takeaways_body) < 3:
        issues.append("Key Takeaways must include at least 3 bullet items")

    if "<<<BEGIN_ARTIFACT:" in final_notes or "<<<END_ARTIFACT>>>" in final_notes:
        issues.append("enhanced notes contain raw artifact markers")

    edit_log_path = stable_pipeline_dir / "edit_log.md"
    if not edit_log_path.exists():
        issues.append(f"missing edit log: {edit_log_path}")
        return issues

    edit_log = edit_log_path.read_text(encoding="utf-8")
    for heading in REQUIRED_EDIT_LOG_HEADINGS:
        if heading not in edit_log:
            issues.append(f"missing edit log heading: {heading}")
    if "<<<BEGIN_ARTIFACT:" in edit_log or "<<<END_ARTIFACT>>>" in edit_log:
        issues.append("edit log contains raw artifact markers")

    return issues


def write_report(report_path: Path, issues: list[str]) -> None:
    lines = ["# Validation Report", ""]
    if issues:
        lines.extend(["status: fail", ""])
        lines.extend(f"- {issue}" for issue in issues)
    else:
        lines.extend(["status: pass", "", "- Note enhancement outputs are present and valid."])
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
