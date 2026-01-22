#!/usr/bin/env python3
"""md2toon - Markdown → TOON transpiler for fabric prompts.

Converts Markdown system prompts to structured JSON, then to TOON format.
Achieves ~80% token savings on typical fabric prompts.

Usage:
    md2toon.py <input.md> [-o output.toon] [-b] [-r]
    md2toon.py --batch <dir> [--output-dir <dir>] [-r]
"""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path
from typing import TYPE_CHECKING

from agent import PromptExtractor
from converter import json_to_toon
from parser import parse_markdown_prompt

if TYPE_CHECKING:
    import tiktoken as _tiktoken

__all__ = ["convert", "convert_batch", "main"]

# Lazy-load tiktoken for token counting
_tiktoken_enc: _tiktoken.Encoding | None = None


def _count_tokens(text: str) -> int:
    """Count tokens using tiktoken, with fallback estimation."""
    global _tiktoken_enc
    try:
        if _tiktoken_enc is None:
            import tiktoken
            _tiktoken_enc = tiktoken.get_encoding("cl100k_base")
        return len(_tiktoken_enc.encode(text))
    except ImportError:
        return len(text) // 4  # ~4 chars/token estimate


def convert(
    path: Path,
    *,
    output: Path | None = None,
    as_json: bool = False,
    benchmark: bool = False,
    provider: str = "anthropic",
    model: str | None = None,
    rule_based: bool = False,
) -> dict:
    """Convert a single Markdown prompt to TOON.

    Args:
        path: Input Markdown file.
        output: Output file path (optional).
        as_json: Write JSON instead of TOON.
        benchmark: Include token count comparison.
        provider: LLM provider for extraction.
        model: Model name override.
        rule_based: Use rule-based parser instead of LLM.

    Returns:
        Dict with structured_json, toon, and optional benchmark.
    """
    content = path.read_text()

    # Extract structure
    if rule_based:
        structured = parse_markdown_prompt(content)
    else:
        extractor = PromptExtractor(provider=provider, model=model)
        structured = extractor.extract(content)

    # Convert to TOON
    toon = json_to_toon(structured)
    json_str = json.dumps(structured, indent=2)

    result: dict = {
        "input": str(path),
        "structured_json": structured,
        "toon": toon,
    }

    if benchmark:
        md_tokens = _count_tokens(content)
        toon_tokens = _count_tokens(toon)
        result["benchmark"] = {
            "original_markdown": {"chars": len(content), "tokens": md_tokens},
            "structured_json": {"chars": len(json_str), "tokens": _count_tokens(json_str)},
            "toon": {"chars": len(toon), "tokens": toon_tokens},
            "savings": {
                "toon_vs_markdown": f"{(md_tokens - toon_tokens) / md_tokens * 100:.1f}%" if md_tokens else "N/A",
            },
        }

    if output:
        output.write_text(json_str if as_json else toon)

    return result


def convert_batch(
    input_dir: Path,
    *,
    output_dir: Path | None = None,
    provider: str = "anthropic",
    model: str | None = None,
    rule_based: bool = False,
) -> dict:
    """Convert all system.md files in a directory.

    Args:
        input_dir: Directory containing pattern subdirectories.
        output_dir: Directory for .toon output files.
        provider: LLM provider.
        model: Model name override.
        rule_based: Use rule-based parser.

    Returns:
        Aggregate statistics and per-file results.
    """
    files = sorted(input_dir.rglob("system.md"))
    if not files:
        raise ValueError(f"No system.md files in {input_dir}")

    if output_dir:
        output_dir.mkdir(parents=True, exist_ok=True)

    results: list[dict] = []
    total_md, total_toon = 0, 0

    for i, path in enumerate(files, 1):
        name = path.parent.name
        print(f"[{i}/{len(files)}] {name}", file=sys.stderr)

        try:
            out = output_dir / f"{name}.toon" if output_dir else None
            res = convert(
                path,
                output=out,
                benchmark=True,
                provider=provider,
                model=model,
                rule_based=rule_based,
            )
            bm = res.get("benchmark", {})
            total_md += bm.get("original_markdown", {}).get("tokens", 0)
            total_toon += bm.get("toon", {}).get("tokens", 0)
            results.append({"file": str(path), "status": "success", "benchmark": bm})
        except Exception as e:
            results.append({"file": str(path), "status": "error", "error": str(e)})

    savings = f"{(total_md - total_toon) / total_md * 100:.1f}%" if total_md else "N/A"
    return {
        "processed": len(results),
        "successful": sum(1 for r in results if r["status"] == "success"),
        "failed": sum(1 for r in results if r["status"] == "error"),
        "aggregate": {"total_md_tokens": total_md, "total_toon_tokens": total_toon, "savings": savings},
        "files": results,
    }


def _build_parser() -> argparse.ArgumentParser:
    """Build CLI argument parser."""
    p = argparse.ArgumentParser(
        prog="md2toon",
        description="Convert Markdown prompts to TOON format",
    )
    p.add_argument("input", type=Path, nargs="?", help="Input Markdown file")
    p.add_argument("-o", "--output", type=Path, help="Output file")
    p.add_argument("-b", "--benchmark", action="store_true", help="Show token savings")
    p.add_argument("-r", "--rule-based", action="store_true", help="Use rule-based parser (no API)")
    p.add_argument("--json", action="store_true", help="Output JSON instead of TOON")
    p.add_argument("--batch", type=Path, metavar="DIR", help="Batch convert directory")
    p.add_argument("--output-dir", type=Path, help="Output directory for batch")
    p.add_argument("--provider", choices=["anthropic", "openai"], default="anthropic")
    p.add_argument("--model", help="Model name override")
    p.add_argument("--show-toon", action="store_true", help="Print TOON to stdout")
    p.add_argument("--show-json", action="store_true", help="Print JSON to stdout")
    return p


def main() -> None:
    """CLI entry point."""
    args = _build_parser().parse_args()

    if args.batch:
        result = convert_batch(
            args.batch,
            output_dir=args.output_dir,
            provider=args.provider,
            model=args.model,
            rule_based=args.rule_based,
        )
        print(json.dumps(result, indent=2))
        return

    if not args.input:
        _build_parser().print_help()
        sys.exit(1)

    result = convert(
        args.input,
        output=args.output,
        as_json=args.json,
        benchmark=args.benchmark,
        provider=args.provider,
        model=args.model,
        rule_based=args.rule_based,
    )

    if args.show_json:
        print(json.dumps(result["structured_json"], indent=2))
    elif args.show_toon:
        print(result["toon"])

    if args.benchmark and "benchmark" in result:
        bm = result["benchmark"]
        print(f"\nTokens: {bm['original_markdown']['tokens']} MD → {bm['toon']['tokens']} TOON ({bm['savings']['toon_vs_markdown']} savings)", file=sys.stderr)


if __name__ == "__main__":
    main()
