"""Pydantic schema for fabric prompt structure.

Designed to maximize TOON compression via uniform, tabular-friendly structure.
All arrays use consistent field sets to enable TOON's tabular encoding.

Note: This schema extracts *structure* from prose prompts. Some verbatim
content may be summarized or normalized. This is the tradeoff for ~80%
token savings.
"""

from __future__ import annotations

from pydantic import BaseModel, Field

__all__ = [
    "Step",
    "OutputInstruction", 
    "OutputSection",
    "Restriction",
    "FabricPromptSchema",
    "PROMPT_SCHEMA",
]


class Step(BaseModel):
    """A single step in the prompt workflow."""

    action: str = Field(description="The instruction or action to perform")


class OutputInstruction(BaseModel):
    """A formatting or output constraint."""

    instruction: str = Field(description="The formatting rule or constraint")


class OutputSection(BaseModel):
    """A named section the model should produce."""

    name: str = Field(description="Section heading name")
    description: str = Field(description="What this section should contain")


class Restriction(BaseModel):
    """A constraint or limitation on model behavior."""

    rule: str = Field(description="The restriction or constraint")


class FabricPromptSchema(BaseModel):
    """Structured representation of a fabric system prompt.

    Fields are ordered to maximize TOON tabular compression:
    - Scalar fields first (role, purpose, output_format)
    - Uniform arrays last (steps, output_instructions, restrictions)
    """

    role: str = Field(default="", description="The role/persona the model assumes")
    expertise: list[str] = Field(default_factory=list, description="Areas of expertise")
    purpose: str = Field(default="", description="The primary purpose or goal")
    thinking_instruction: str | None = Field(default=None, description="Meta-instruction like 'think step by step'")
    steps: list[Step] = Field(default_factory=list, description="Ordered steps to follow")
    output_format: str = Field(default="markdown", description="Primary output format")
    output_sections: list[OutputSection] = Field(default_factory=list, description="Named sections to produce")
    output_instructions: list[OutputInstruction] = Field(default_factory=list, description="Formatting rules")
    restrictions: list[Restriction] = Field(default_factory=list, description="Constraints on behavior")


PROMPT_SCHEMA = FabricPromptSchema.model_json_schema()
