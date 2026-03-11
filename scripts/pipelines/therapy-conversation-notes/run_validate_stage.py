#!/usr/bin/env python3
"""Validate therapy-conversation-notes outputs."""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path

from common import load_session_context, run_dir


REQUIRED_HEADINGS = [
    "# 🧠 Therapy Conversation Notes",
    "## Conversation Summary",
    "## Emotional Signals",
    "## Thought Patterns",
    "## Actionable Reflection",
    "## Safety and Boundaries",
]

REQUIRED_SNAPSHOT_KEYS = [
    "summary_points",
    "emotions_observed",
    "patterns_observed",
    "actions",
    "safety_flags",
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
    snapshot_path = stable_pipeline_dir / "analysis_snapshot.json"
    if not snapshot_path.exists():
        issues.append(f"missing analysis snapshot: {snapshot_path}")
    else:
        try:
            snapshot = json.loads(snapshot_path.read_text(encoding="utf-8"))
        except json.JSONDecodeError as err:
            issues.append(f"analysis_snapshot.json is not valid JSON: {err}")
        else:
            if not isinstance(snapshot, dict):
                issues.append("analysis_snapshot.json must be a JSON object")
            else:
                for key in REQUIRED_SNAPSHOT_KEYS:
                    if key not in snapshot:
                        issues.append(f"analysis_snapshot.json missing key: {key}")
                        continue
                    value = snapshot[key]
                    if not isinstance(value, list) or not all(isinstance(item, str) for item in value):
                        issues.append(f"analysis_snapshot.json key '{key}' must be an array of strings")

    final_output_path = Path(session_context.final_output_path)
    if not final_output_path.exists():
        issues.append(f"missing therapy notes: {final_output_path}")
        return issues

    final_notes = final_output_path.read_text(encoding="utf-8")
    for heading in REQUIRED_HEADINGS:
        if heading not in final_notes:
            issues.append(f"missing heading: {heading}")

    reflection_body = section_body(final_notes, "## Actionable Reflection")
    if count_bullets(reflection_body) < 3:
        issues.append("Actionable Reflection must include at least 3 bullet items")

    boundaries_body = section_body(final_notes, "## Safety and Boundaries").lower()
    if "not a substitute for professional care" not in boundaries_body:
        issues.append("Safety and Boundaries must include professional-care disclaimer language")

    if "<<<BEGIN_ARTIFACT:" in final_notes or "<<<END_ARTIFACT>>>" in final_notes:
        issues.append("final notes contain raw artifact markers")

    return issues


def write_report(report_path: Path, issues: list[str]) -> None:
    lines = ["# Validation Report", ""]
    if issues:
        lines.extend(["status: fail", ""])
        lines.extend(f"- {issue}" for issue in issues)
    else:
        lines.extend(["status: pass", "", "- Therapy conversation notes outputs are present and valid."])
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
