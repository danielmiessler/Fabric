#!/usr/bin/env python3
"""Validate technical-study-guide outputs."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from common import load_session_context, run_dir


REQUIRED_HEADINGS = [
    "# 🎓 ",
    "## 🧠 Session Focus",
    "## 🎯 Prerequisites",
    "## ✅ Learning Outcomes",
    "## 🧭 Topic Index",
    "## 🗺️ Conceptual Roadmap",
    "## 🏗️ Systems Visualization",
    "## 🌆 Skyline Intuition Diagram",
    "## 📚 Core Concepts (Intuition First)",
    "## ➗ Mathematical Intuition",
    "## 💻 Coding Walkthroughs",
    "## 🚀 Advanced Real-World Scenario",
    "## 🧩 HOTS (High-Order Thinking)",
    "## ❓ FAQ",
    "## 🛠️ Practice Roadmap",
    "## 🔭 Next Improvements",
    "## 🔗 Related Notes",
]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument("--session-context", required=True)
    return parser.parse_args()


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
        else:
            if not isinstance(data, dict):
                issues.append("semantic_map.json must be a JSON object")

    final_output_path = Path(session_context.final_output_path)
    if not final_output_path.exists():
        issues.append(f"missing final notes: {final_output_path}")
        return issues

    final_notes = final_output_path.read_text(encoding="utf-8")
    for heading in REQUIRED_HEADINGS:
        if heading not in final_notes:
            issues.append(f"missing heading: {heading}")

    if final_notes.count("```mermaid") < 2:
        issues.append("expected at least 2 Mermaid diagrams")
    if "<<<BEGIN_ARTIFACT:" in final_notes or "<<<END_ARTIFACT>>>" in final_notes:
        issues.append("final notes contain raw artifact markers")

    return issues


def write_report(report_path: Path, issues: list[str]) -> None:
    lines = ["# Validation Report", ""]
    if issues:
        lines.extend(["status: fail", ""])
        lines.extend(f"- {issue}" for issue in issues)
    else:
        lines.extend(["status: pass", "", "- All required study-guide outputs are present and valid."])
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
