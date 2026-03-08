#!/usr/bin/env python3
"""Deterministic hard-gate validator for transcript pipeline artifacts.

No API usage. Pure local deterministic checks.
"""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path


SOURCE_TAG_RE = re.compile(r"\[source:\s*([A-Za-z0-9._:-]+)\]")


def read_jsonl(path: Path) -> list[dict]:
    rows: list[dict] = []
    with path.open("r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            rows.append(json.loads(line))
    return rows


def parse_source_ids_from_text(text: str) -> set[str]:
    return set(SOURCE_TAG_RE.findall(text))


def parse_coverage_matrix(path: Path) -> set[str]:
    if not path.exists():
        return set()
    obj = json.loads(path.read_text(encoding="utf-8"))
    ids: set[str] = set()
    if isinstance(obj, dict):
        for key, value in obj.items():
            if isinstance(key, str):
                ids.add(key)
            if isinstance(value, list):
                for item in value:
                    if isinstance(item, dict) and isinstance(item.get("segment_id"), str):
                        ids.add(item["segment_id"])
                    elif isinstance(item, str):
                        ids.add(item)
            elif isinstance(value, dict) and isinstance(value.get("segment_id"), str):
                ids.add(value["segment_id"])
    elif isinstance(obj, list):
        for item in obj:
            if isinstance(item, dict) and isinstance(item.get("segment_id"), str):
                ids.add(item["segment_id"])
            elif isinstance(item, str):
                ids.add(item)
    return ids


def parse_unresolved_uncertainty_ids(path: Path) -> set[str]:
    if not path.exists():
        return set()
    data = json.loads(path.read_text(encoding="utf-8"))
    rows = data if isinstance(data, list) else [data]
    unresolved: set[str] = set()
    for row in rows:
        if not isinstance(row, dict):
            continue
        seg_id = row.get("segment_id")
        if not isinstance(seg_id, str):
            continue
        status = str(row.get("status", "")).lower()
        tier = str(row.get("confidence_tier", "")).upper()
        resolved = row.get("resolved")
        if status in {"open", "unresolved"} or tier == "LOW" or resolved is False:
            unresolved.add(seg_id)
    return unresolved


def write_markdown_report(
    output_path: Path,
    passed: bool,
    summary: dict,
    missing_content: list[str],
    orphan_sources: list[str],
    missing_uncertain: list[str],
) -> None:
    lines = []
    lines.append("# Validation Report")
    lines.append("")
    lines.append(f"- **status:** {'PASS' if passed else 'FAIL'}")
    lines.append(f"- **segments_total:** {summary['segments_total']}")
    lines.append(f"- **segments_content:** {summary['segments_content']}")
    lines.append(f"- **segments_noise:** {summary['segments_noise']}")
    lines.append(f"- **coverage_source_ids:** {summary['coverage_source_ids']}")
    lines.append(f"- **coverage_matrix_ids:** {summary['coverage_matrix_ids']}")
    lines.append(f"- **missing_content_count:** {len(missing_content)}")
    lines.append(f"- **orphan_source_count:** {len(orphan_sources)}")
    lines.append(f"- **missing_uncertainty_retention_count:** {len(missing_uncertain)}")
    lines.append("")

    if missing_content:
        lines.append("## Missing Content Segment IDs")
        lines.extend(f"- `{sid}`" for sid in missing_content[:200])
        lines.append("")
    if orphan_sources:
        lines.append("## Orphan Source IDs In Notes/Coverage")
        lines.extend(f"- `{sid}`" for sid in orphan_sources[:200])
        lines.append("")
    if missing_uncertain:
        lines.append("## Missing Uncertainty Retention IDs")
        lines.extend(f"- `{sid}`" for sid in missing_uncertain[:200])
        lines.append("")

    output_path.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser(description="Validate segment coverage deterministically.")
    parser.add_argument("--ledger", required=True, help="Path to .pipeline/segment_ledger.jsonl")
    parser.add_argument("--final-notes", required=True, help="Path to final_notes.md")
    parser.add_argument("--coverage-matrix", required=False, help="Path to .pipeline/coverage_matrix.json")
    parser.add_argument("--uncertainty-report", required=False, help="Path to .pipeline/uncertainty_report.json")
    parser.add_argument("--human-review-queue", required=False, help="Path to .pipeline/human_review_queue.md")
    parser.add_argument("--report-out", required=True, help="Path to .pipeline/validation_report.md")
    parser.add_argument("--exceptions-out", required=True, help="Path to .pipeline/exceptions.json")
    args = parser.parse_args()

    ledger_path = Path(args.ledger).expanduser().resolve()
    final_notes_path = Path(args.final_notes).expanduser().resolve()
    coverage_matrix_path = Path(args.coverage_matrix).expanduser().resolve() if args.coverage_matrix else None
    uncertainty_path = Path(args.uncertainty_report).expanduser().resolve() if args.uncertainty_report else None
    review_queue_path = Path(args.human_review_queue).expanduser().resolve() if args.human_review_queue else None
    report_out = Path(args.report_out).expanduser().resolve()
    exceptions_out = Path(args.exceptions_out).expanduser().resolve()

    ledger_rows = read_jsonl(ledger_path)
    notes_text = final_notes_path.read_text(encoding="utf-8")
    source_ids_from_notes = parse_source_ids_from_text(notes_text)
    source_ids_from_matrix = parse_coverage_matrix(coverage_matrix_path) if coverage_matrix_path else set()
    all_source_ids = source_ids_from_notes | source_ids_from_matrix

    ledger_ids = {row["segment_id"] for row in ledger_rows}
    content_ids = {row["segment_id"] for row in ledger_rows if row.get("type") != "noise"}
    noise_ids = {row["segment_id"] for row in ledger_rows if row.get("type") == "noise"}

    missing_content = sorted(sid for sid in content_ids if sid not in all_source_ids)
    orphan_sources = sorted(sid for sid in all_source_ids if sid not in ledger_ids)

    unresolved_ids = parse_unresolved_uncertainty_ids(uncertainty_path) if uncertainty_path else set()
    retained_uncertain_ids = set()
    if unresolved_ids:
        retained_uncertain_ids |= {sid for sid in unresolved_ids if sid in all_source_ids}
        if review_queue_path and review_queue_path.exists():
            rq_text = review_queue_path.read_text(encoding="utf-8")
            retained_uncertain_ids |= {sid for sid in unresolved_ids if sid in rq_text}
    missing_uncertain = sorted(unresolved_ids - retained_uncertain_ids)

    passed = not missing_content and not orphan_sources and not missing_uncertain

    summary = {
        "segments_total": len(ledger_rows),
        "segments_content": len(content_ids),
        "segments_noise": len(noise_ids),
        "coverage_source_ids": len(source_ids_from_notes),
        "coverage_matrix_ids": len(source_ids_from_matrix),
    }

    write_markdown_report(
        output_path=report_out,
        passed=passed,
        summary=summary,
        missing_content=missing_content,
        orphan_sources=orphan_sources,
        missing_uncertain=missing_uncertain,
    )

    exceptions = {
        "status": "PASS" if passed else "FAIL",
        "summary": summary,
        "missing_content_segment_ids": missing_content,
        "orphan_source_ids": orphan_sources,
        "missing_uncertainty_retention_ids": missing_uncertain,
    }
    exceptions_out.write_text(json.dumps(exceptions, ensure_ascii=True, indent=2) + "\n", encoding="utf-8")

    print(f"status: {'PASS' if passed else 'FAIL'}")
    print(f"wrote: {report_out}")
    print(f"wrote: {exceptions_out}")
    if missing_content:
        print(f"missing_content_count: {len(missing_content)}")
    if orphan_sources:
        print(f"orphan_source_count: {len(orphan_sources)}")
    if missing_uncertain:
        print(f"missing_uncertainty_retention_count: {len(missing_uncertain)}")
    return 0 if passed else 1


if __name__ == "__main__":
    raise SystemExit(main())

