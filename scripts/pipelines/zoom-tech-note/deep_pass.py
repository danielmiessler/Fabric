#!/usr/bin/env python3
"""Strict deep-pass quality gate for the Fabric-native zoom-tech-note pipeline."""

from __future__ import annotations

import argparse
import json
import re
from dataclasses import dataclass
from pathlib import Path

from common import load_session_context


@dataclass
class DeepPassCheck:
    check_id: str
    title: str
    passed: bool
    details: str


def find_section(text: str, keywords: list[str]) -> str:
    flags = re.IGNORECASE | re.MULTILINE
    heading_pattern = re.compile(r"^##\s+.*$", flags)
    matches = list(heading_pattern.finditer(text))
    if not matches:
        return ""

    kw_pattern = re.compile("|".join(re.escape(keyword) for keyword in keywords), re.IGNORECASE)
    for idx, match in enumerate(matches):
        if not kw_pattern.search(match.group(0)):
            continue
        start = match.start()
        end = matches[idx + 1].start() if idx + 1 < len(matches) else len(text)
        return text[start:end]
    return ""


def count_list_items(section: str) -> int:
    count = 0
    for line in section.splitlines():
        if re.match(r"^\s*[-*]\s+\S+", line):
            count += 1
        elif re.match(r"^\s*\d+\.\s+\S+", line):
            count += 1
    return count


def run_deep_pass_checks(final_notes_text: str) -> list[DeepPassCheck]:
    text = final_notes_text
    lower = text.lower()

    prereq_section = find_section(text, ["prereq", "prerequisite"])
    hots_section = find_section(text, ["hots", "high-order", "high order"])
    faq_section = find_section(text, ["faq"])
    practice_section = find_section(text, ["practice roadmap", "practice plan", "practice"])

    intuition_hits = len(re.findall(r"\bintuition\b", lower))
    mermaid_hits = len(re.findall(r"```mermaid", lower))
    hots_items = count_list_items(hots_section)
    faq_q_hits = len(re.findall(r"(?im)^\s*(?:###\s+)?q(?:\d+)?(?:\s*[:.\-]|\b)", faq_section))
    practice_items = count_list_items(practice_section)

    return [
        DeepPassCheck(
            check_id="prereq_rescue",
            title="Prerequisite rescue section",
            passed=bool(prereq_section.strip()),
            details="Requires a dedicated prerequisites/prereq section.",
        ),
        DeepPassCheck(
            check_id="intuition_depth",
            title="Intuition-first depth",
            passed=intuition_hits >= 3,
            details=f"Found 'intuition' mentions: {intuition_hits} (min: 3).",
        ),
        DeepPassCheck(
            check_id="mermaid_diagrams",
            title="Mermaid diagrams present",
            passed=mermaid_hits >= 1,
            details=f"Found mermaid blocks: {mermaid_hits} (min: 1).",
        ),
        DeepPassCheck(
            check_id="hots_section",
            title="HOTS section with actionable questions",
            passed=bool(hots_section.strip()) and hots_items >= 2,
            details=f"HOTS list items: {hots_items} (min: 2).",
        ),
        DeepPassCheck(
            check_id="faq_section",
            title="FAQ section with multiple Q/A entries",
            passed=bool(faq_section.strip()) and faq_q_hits >= 2,
            details=f"FAQ question markers: {faq_q_hits} (min: 2).",
        ),
        DeepPassCheck(
            check_id="practice_plan",
            title="Practice plan/roadmap section",
            passed=bool(practice_section.strip()) and practice_items >= 2,
            details=f"Practice items: {practice_items} (min: 2).",
        ),
    ]


def main() -> int:
    parser = argparse.ArgumentParser(description="Run deep-pass quality checks for zoom-tech-note.")
    parser.add_argument("--session-context", required=True, help="Path to session_context.json")
    args = parser.parse_args()

    context = load_session_context(Path(args.session_context).expanduser().resolve())
    session_dir = Path(context.session_dir).expanduser().resolve()
    pipeline_dir = session_dir / ".pipeline"
    pipeline_dir.mkdir(parents=True, exist_ok=True)

    final_notes_path = session_dir / "final_notes.md"
    report_path = pipeline_dir / "deep_pass_report.md"
    exceptions_path = pipeline_dir / "deep_pass_exceptions.json"

    if not final_notes_path.exists():
        report_path.write_text(
            "# Deep Pass Report\n\n- **status:** FAIL\n- **reason:** final_notes.md missing\n",
            encoding="utf-8",
        )
        exceptions_path.write_text(
            json.dumps(
                {
                    "status": "FAIL",
                    "reason": "missing_final_notes",
                    "missing": ["final_notes.md"],
                },
                ensure_ascii=True,
                indent=2,
            )
            + "\n",
            encoding="utf-8",
        )
        return 1

    text = final_notes_path.read_text(encoding="utf-8")
    checks = run_deep_pass_checks(text)
    failed = [check for check in checks if not check.passed]
    passed = not failed

    lines = ["# Deep Pass Report", "", f"- **status:** {'PASS' if passed else 'FAIL'}", ""]
    lines.append("## Checks")
    for check in checks:
        lines.append(f"- **{check.title}:** {'PASS' if check.passed else 'FAIL'} - {check.details}")
    lines.append("")
    if failed:
        lines.append("## Missing Requirements")
        for check in failed:
            lines.append(f"- `{check.check_id}`: {check.title}")
        lines.append("")

    report_path.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
    exceptions_path.write_text(
        json.dumps(
            {
                "status": "PASS" if passed else "FAIL",
                "checks": [
                    {
                        "id": check.check_id,
                        "title": check.title,
                        "passed": check.passed,
                        "details": check.details,
                    }
                    for check in checks
                ],
                "failed_ids": [check.check_id for check in failed],
            },
            ensure_ascii=True,
            indent=2,
        )
        + "\n",
        encoding="utf-8",
    )
    return 0 if passed else 1


if __name__ == "__main__":
    raise SystemExit(main())
