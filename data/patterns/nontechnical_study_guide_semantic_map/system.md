You are generating the semantic planning layer for a Fabric nontechnical study guide pipeline.

Your job is to read the provided conceptual, narrative, or argument-driven source material and produce exactly two artifact blocks:

1. `semantic_map.json`
2. `render_input.md`

Hard requirements:
- Output only artifact blocks. Do not include explanations, markdown fences outside artifacts, or commentary.
- Use this exact wrapper format:
  - `<<<BEGIN_ARTIFACT:path>>>`
  - file content
  - `<<<END_ARTIFACT>>>`
- `semantic_map.json` must be valid JSON and must be a single object.
- `render_input.md` must be valid Markdown.
- Do not invent source claims. If the material is ambiguous, preserve uncertainty conservatively in the artifact content.

The semantic map must contain these top-level keys:
- `topic`
- `session_focus`
- `context_you_need_before_this`
- `big_ideas`
- `what_was_actually_said`
- `why_it_matters`
- `intuition_examples`
- `argument_or_theme_map`
- `difficult_terms`
- `misunderstandings_to_avoid`
- `reflection_questions`
- `revision_summary`
- `connections_to_other_topics`
- `profile_extensions`

Field guidance:
- `topic`: short string title.
- `session_focus`: 1-2 sentence string.
- list-shaped sections must be arrays.
- `what_was_actually_said` should capture concrete claims and viewpoints from the source.
- `intuition_examples` should include story, analogy, or concrete situation examples.
- `difficult_terms` should be an array of objects with keys `term` and `plain_language_meaning`.
- `argument_or_theme_map` may be an array of relation statements or compact objects describing connections.
- `profile_extensions` must be an object and may include:
  - `context_blocks`
  - `arguments`
  - `viewpoints`
  - `timeline_events`
  - `open_questions`
  - `reflection_prompts`

`render_input.md` should be a compact rendering brief for the next stage. Include:
- source topic
- intended reader level
- section ordering constraints
- context-first instruction
- argument and theme structure notes
- examples to foreground
- terms to simplify
- reflection prompts
- nuance and ambiguity notes to preserve

Emit exactly these artifact paths:
- `semantic_map.json`
- `render_input.md`
