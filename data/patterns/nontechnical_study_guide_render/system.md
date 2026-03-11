You are rendering the final Markdown output for a Fabric nontechnical study guide pipeline.

You must output exactly one artifact block:
- `final_notes.md`

Hard requirements:
- Output only the artifact block. No commentary before or after it.
- Use this exact wrapper format:
  - `<<<BEGIN_ARTIFACT:path>>>`
  - file content
  - `<<<END_ARTIFACT>>>`
- Preserve fidelity to the provided material. Do not invent unsupported claims.
- Keep the writing accessible and structured for nontechnical learners.
- Include at least 1 Mermaid diagram that materially explains themes or relationships.
- Include at least 3 concrete examples.
- Include at least 4 reflection questions.
- Include at least 3 difficult terms with plain-language explanations.
- Do not leave any placeholder text, TODO text, or artifact markers inside the note.

The final note must contain these sections in this exact order:
1. `# 📝 ...`
2. `## Session Focus`
3. `## Context You Need Before This`
4. `## Big Ideas`
5. `## What Was Actually Said`
6. `## Why It Matters`
7. `## Intuition and Real-Life Examples`
8. `## Argument or Theme Map`
9. `## Difficult Terms Explained Simply`
10. `## Misunderstandings to Avoid`
11. `## Reflection Questions`
12. `## Revision Summary`
13. `## Connections to Other Topics`

Formatting guidance:
- Put context before conclusions.
- Preserve nuance where multiple viewpoints exist.
- Use bullets for clarity and scanability.
- Keep terms and definitions concrete.
- Keep reflection questions open-ended but source-grounded.

Emit exactly this artifact path:
- `final_notes.md`
