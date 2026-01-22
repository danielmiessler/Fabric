"""Rule-based Markdown → Schema parser.

Extracts structured data from fabric prompts using regex patterns.
No LLM required. Handles standard fabric prompt section headers:
IDENTITY, PURPOSE, STEPS, OUTPUT INSTRUCTIONS, RESTRICTIONS.

This is a *lossy* extraction — prose is normalized into structured fields.
The tradeoff enables ~80% token savings via TOON compression.
"""

from __future__ import annotations

import re
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from collections.abc import Callable

__all__ = ["parse_markdown_prompt"]


# Section header aliases for flexible parsing
_IDENTITY_HEADERS = frozenset({"IDENTITY AND PURPOSE", "IDENTITY", "PURPOSE", "IDENTITY & PURPOSE"})
_STEP_HEADERS = frozenset({"STEPS", "ACTIONS", "TASK", "PROCESS"})
_OUTPUT_HEADERS = frozenset({"OUTPUT INSTRUCTIONS", "OUTPUT", "FORMAT"})
_RESTRICTION_HEADERS = frozenset({"RESTRICTIONS", "CONSTRAINTS", "RULES", "LIMITATIONS"})


def parse_markdown_prompt(content: str) -> dict:
    """Parse a fabric Markdown prompt into a structured dict.

    Args:
        content: Raw Markdown content of a fabric system prompt.

    Returns:
        Dict matching FabricPromptSchema structure, ready for TOON conversion.
    """
    sections = _split_sections(content)

    result: dict = {
        "role": "",
        "expertise": [],
        "purpose": "",
        "steps": [],
        "output_format": "markdown",
        "output_sections": [],
        "output_instructions": [],
        "restrictions": [],
    }

    # Parse identity
    for key in _IDENTITY_HEADERS:
        if key in sections:
            text = sections[key]
            result["role"], result["expertise"], result["purpose"] = _parse_identity(text)
            if _contains_thinking_instruction(text):
                result["thinking_instruction"] = "Think step by step"
            break

    # Parse steps
    for key in _STEP_HEADERS:
        if key in sections:
            result["steps"] = _parse_steps(sections[key])
            break

    # Parse output
    for key in _OUTPUT_HEADERS:
        if key in sections:
            result["output_sections"], result["output_instructions"] = _parse_output(sections[key])
            break

    # Parse restrictions
    for key in _RESTRICTION_HEADERS:
        if key in sections:
            result["restrictions"] = _parse_restrictions(sections[key])
            break

    return result


def _split_sections(content: str) -> dict[str, str]:
    """Split markdown into sections by # headers."""
    sections: dict[str, str] = {}
    current_header: str | None = None
    lines: list[str] = []

    for line in content.split("\n"):
        match = re.match(r"^#+\s*(.+)$", line)
        if match:
            if current_header:
                sections[current_header] = "\n".join(lines).strip()
            current_header = match.group(1).strip().upper()
            lines = []
        else:
            lines.append(line)

    if current_header:
        sections[current_header] = "\n".join(lines).strip()

    return sections


def _contains_thinking_instruction(text: str) -> bool:
    """Check if text contains 'step by step' meta-instruction."""
    lower = text.lower()
    return "step by step" in lower or "think step" in lower


def _parse_identity(text: str) -> tuple[str, list[str], str]:
    """Extract role, expertise, and purpose from identity section."""
    sentences = [s.strip() for s in re.split(r"(?<=[.!?])\s+", text.strip()) if s.strip()]
    if not sentences:
        return "", [], ""

    role = ""
    expertise: list[str] = []
    purpose = ""

    # First sentence is typically the role definition
    first = sentences[0]
    if "you are" in first.lower() or first.startswith("You"):
        role = first
        expertise = _extract_expertise(first)

    # Find purpose and additional expertise
    purpose_keywords = {"purpose", "goal", "aim", "objective", "task"}
    for sentence in sentences[1:]:
        lower = sentence.lower()
        if any(kw in lower for kw in purpose_keywords):
            purpose = sentence
        elif "specialize" in lower:
            expertise.extend(_extract_expertise(sentence))

    # Fallback: use last sentence as purpose
    if not purpose and len(sentences) > 1:
        purpose = sentences[-1]

    return role, expertise, purpose


def _extract_expertise(text: str) -> list[str]:
    """Extract expertise mentions from text."""
    pattern = r"(?:expert in|specialize[sd]? in|skilled in|proficient in)\s+([^.]+)"
    matches = re.findall(pattern, text, re.IGNORECASE)
    result: list[str] = []
    for match in matches:
        items = re.split(r",\s*(?:and\s+)?|\s+and\s+", match)
        result.extend(item.strip() for item in items if item.strip())
    return result


def _parse_steps(text: str) -> list[dict]:
    """Extract action steps from steps section."""
    steps: list[dict] = []

    for line in text.split("\n"):
        line = line.strip()
        if not line:
            continue

        # Bullet points: -, *, •
        if match := re.match(r"^[-*•]\s+(.+)$", line):
            action = _strip_markdown_bold(match.group(1))
            steps.append({"action": action})
        # Numbered items: 1. or 1)
        elif match := re.match(r"^\d+[.)]\s+(.+)$", line):
            steps.append({"action": match.group(1).strip()})

    # Fallback: sentence splitting
    if not steps:
        sentences = [s.strip() for s in text.split(".") if len(s.strip()) > 10]
        steps = [{"action": s} for s in sentences[:5]]

    return steps


def _strip_markdown_bold(text: str) -> str:
    """Remove **bold** markers from text."""
    return re.sub(r"\*\*([^*]+)\*\*", r"\1", text).strip()


# Patterns for detecting formatting instructions
_OUTPUT_PATTERNS: list[tuple[str, str | Callable[[re.Match], str]]] = [
    (r"output.*markdown", "Output in Markdown format"),
    (r"output.*json", "Output in JSON format"),
    (r"do not use.*bold", "Do not use bold formatting"),
    (r"do not use.*italic", "Do not use italic formatting"),
    (r"use bulleted lists", "Use bulleted lists"),
    (r"(\d+)\s*words?\s*(?:or\s*)?(?:less|max)", lambda m: f"Maximum {m.group(1)} words"),
    (r"(\d+)\s*bullets?", lambda m: f"Use {m.group(1)} bullet points"),
]


def _parse_output(text: str) -> tuple[list[dict], list[dict]]:
    """Extract output sections and formatting instructions."""
    sections: list[dict] = []
    instructions: list[dict] = []

    # Named sections: "in a section called X"
    pattern = r"(?:in a section called|under the heading|in a subsection called)\s*[\"\']?([A-Z][A-Z\s_-]+)[\"\']?"
    for name in re.findall(pattern, text, re.IGNORECASE):
        clean_name = name.strip().rstrip(":")
        sections.append({"name": clean_name, "description": f"Output section: {clean_name}"})

    # Formatting instructions from patterns
    for pattern, result in _OUTPUT_PATTERNS:
        if match := re.search(pattern, text, re.IGNORECASE):
            instruction = result(match) if callable(result) else result
            instructions.append({"instruction": instruction})

    # Bullet point instructions
    for bullet in re.findall(r"^\s*[-*]\s+(.+)$", text, re.MULTILINE):
        if len(bullet) < 100:
            instructions.append({"instruction": bullet.strip()})

    return sections, instructions


def _parse_restrictions(text: str) -> list[dict]:
    """Extract restrictions and constraints."""
    restrictions: list[dict] = []

    # Bullets: - rule text
    for bullet in re.findall(r"^\s*[-*•]\s+(.+)$", text, re.MULTILINE):
        restrictions.append({"rule": _strip_markdown_bold(bullet)})

    return restrictions
