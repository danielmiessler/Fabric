"""LLM-powered Markdown → Schema extractor.

Optimized for Anthropic Claude (claude-sonnet-4-20250514).
Falls back to rule-based parsing when no API key is available.
"""

from __future__ import annotations

import json
from pathlib import Path
from typing import TYPE_CHECKING, Any

from parser import parse_markdown_prompt
from schema import PROMPT_SCHEMA, FabricPromptSchema

if TYPE_CHECKING:
    from anthropic import Anthropic
    from openai import OpenAI

__all__ = ["PromptExtractor", "extract_file"]

_SYSTEM_PROMPT = """Extract the Markdown prompt into the JSON schema provided.
Be precise. Do not add information not present in the source.
Output valid JSON only, no markdown fences."""


def _build_user_prompt(markdown: str) -> str:
    """Build extraction prompt with schema and content."""
    schema_json = json.dumps(PROMPT_SCHEMA, indent=2)
    return f"""SCHEMA:
{schema_json}

MARKDOWN:
{markdown}

JSON:"""


class PromptExtractor:
    """Extract structured data from Markdown prompts.

    Supports Anthropic and OpenAI backends, with rule-based fallback.
    """

    MODELS = {
        "anthropic": "claude-sonnet-4-20250514",
        "openai": "gpt-4o-mini",
    }

    def __init__(self, provider: str = "anthropic", model: str | None = None):
        if provider not in self.MODELS:
            raise ValueError(f"Unsupported provider: {provider}")
        self.provider = provider
        self.model = model or self.MODELS[provider]
        self._client: Anthropic | OpenAI | None = None

    @property
    def client(self) -> Anthropic | OpenAI:
        """Lazy-load the API client."""
        if self._client is None:
            if self.provider == "openai":
                from openai import OpenAI
                self._client = OpenAI()
            else:
                from anthropic import Anthropic
                self._client = Anthropic()
        return self._client

    def extract(self, markdown: str) -> dict:
        """Extract structured data using LLM."""
        prompt = _build_user_prompt(markdown)

        if self.provider == "openai":
            raw = self._call_openai(prompt)
        else:
            raw = self._call_anthropic(prompt)

        return json.loads(_strip_code_fences(raw))

    def extract_rule_based(self, markdown: str) -> dict:
        """Extract using rule-based parser (no API required)."""
        return parse_markdown_prompt(markdown)

    def _call_openai(self, prompt: str) -> str:
        response = self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": _SYSTEM_PROMPT},
                {"role": "user", "content": prompt},
            ],
            response_format={"type": "json_object"},
            temperature=0,
        )
        return response.choices[0].message.content or ""

    def _call_anthropic(self, prompt: str) -> str:
        response = self.client.messages.create(
            model=self.model,
            max_tokens=8192,
            system=_SYSTEM_PROMPT,
            messages=[{"role": "user", "content": prompt}],
            temperature=0,
        )
        return response.content[0].text


def _strip_code_fences(text: str) -> str:
    """Remove markdown code fences from JSON response."""
    if "```json" in text:
        return text.split("```json")[1].split("```")[0]
    if "```" in text:
        return text.split("```")[1].split("```")[0]
    return text


def extract_file(
    path: Path,
    provider: str = "anthropic",
    model: str | None = None,
    rule_based: bool = False,
) -> dict:
    """Extract structured data from a Markdown prompt file.

    Args:
        path: Path to the Markdown file.
        provider: LLM provider ('anthropic' or 'openai').
        model: Model name override.
        rule_based: Use rule-based parser instead of LLM.

    Returns:
        Dict matching FabricPromptSchema structure.
    """
    content = path.read_text()
    extractor = PromptExtractor(provider=provider, model=model)
    if rule_based:
        return extractor.extract_rule_based(content)
    return extractor.extract(content)
