"""md2toon - Markdown → TOON transpiler for fabric prompts.

Converts Markdown system prompts to structured JSON, then to TOON format.
Achieves ~80% token savings on typical fabric prompts.

Example:
    >>> from md2toon import convert
    >>> result = convert(Path("system.md"), rule_based=True, benchmark=True)
    >>> print(result["toon"])
"""

from .agent import PromptExtractor, extract_file
from .converter import json_to_toon
from .md2toon import convert, convert_batch
from .parser import parse_markdown_prompt
from .schema import PROMPT_SCHEMA, FabricPromptSchema

__all__ = [
    "convert",
    "convert_batch",
    "extract_file",
    "json_to_toon",
    "parse_markdown_prompt",
    "FabricPromptSchema",
    "PROMPT_SCHEMA",
    "PromptExtractor",
]
