## Zoom Tech Note Design

Status: draft for implementation

### Goal

Ship a Fabric-native `zoom-tech-note` pipeline inside this repository that preserves the current Zoom transcript pipeline behavior as closely as possible without depending on a private local checkout at runtime.

The reference behavior is the existing Zoom pipeline at:

- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/scripts/ingest_zoom_captions.py`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/scripts/run_chat_pipeline.py`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/scripts/validate_coverage.py`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/scripts/publish_tutorial_notes.py`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/docs/prompts/stages/stage1-refine.md`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/docs/prompts/stages/stage2-synthesize.md`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/docs/prompts/stages/stage3-enhance.md`
- `/Users/praxlannister/Documents/Zoom/transcript-pipeline-kit/docs/prompts/references/tutorial-tech-bar-raiser.md`

The parity target is:

1. same stage intent,
2. same artifact names,
3. same validation and pass/fail semantics,
4. same learner-facing outputs,
5. same deep-pass quality gate semantics,
6. no runtime dependency on the external Zoom repo.

### Approaches Considered

#### Option 1: External adapter to the existing Zoom repo

Fabric would call the current Zoom pipeline scripts through absolute or configurable paths.

Pros:

- fastest proof of reuse,
- nearly zero translation work,
- lowest parity risk in the short term.

Cons:

- not shippable,
- impossible to merge as a real built-in pipeline,
- breaks for users who do not have the private Zoom repo,
- violates the requirement that the feature should live in Fabric.

Decision: rejected.

#### Option 2: Fabric-owned parity port using copied scripts and prompts

Copy the deterministic Zoom scripts and stage prompts into Fabric-owned pipeline assets, then orchestrate them through the Fabric pipeline runner.

Pros:

- shippable,
- closest possible parity,
- preserves the original artifact contract,
- avoids rewriting the already-working deterministic scripts.

Cons:

- needs a small compatibility layer for AI-stage artifact writing,
- still requires careful prompt and script packaging.

Decision: recommended.

#### Option 3: Full semantic rebuild in native Fabric abstractions

Re-express the Zoom pipeline from scratch using new builtins, new patterns, and refactored artifacts.

Pros:

- cleaner long term,
- fewer copied assets.

Cons:

- highest behavioral drift risk,
- duplicates design work already solved by the Zoom reference,
- violates the user's instruction to preserve the real pipeline behavior.

Decision: rejected for the first shipped version.

### Recommended Architecture

The Fabric-native `zoom-tech-note` should be implemented as a Fabric-owned parity port.

That means:

1. deterministic scripts are copied into Fabric with minimal path and packaging changes only,
2. the three AI-stage prompts are copied into Fabric-owned patterns without semantic rewriting,
3. the Fabric pipeline YAML reproduces the same stage order and output contract,
4. any unavoidable behavior change is documented explicitly before implementation.

The design boundary is:

- preserve reference behavior by default,
- only change what is required to make the pipeline shippable inside Fabric,
- treat improvements as explicit follow-ups, not silent edits.

### Reference Stage Model

The Zoom reference pipeline is effectively:

1. deterministic ingest,
2. Stage 1 refine,
3. Stage 2 synthesize,
4. Stage 3 enhance,
5. optional deep pass,
6. deterministic validation,
7. deterministic publish normalization.

Reference artifacts by stage:

#### Stage 0: ingest

Outputs:

- `.pipeline/segment_ledger.jsonl`
- `.pipeline/segment_manifest.jsonl`

#### Stage 1: refine

Outputs:

- `.pipeline/refined_transcript.md`
- `.pipeline/topic_inventory.json`
- `.pipeline/corrections_log.csv`
- `.pipeline/uncertainty_report.json`

#### Stage 2: synthesize

Outputs:

- `.pipeline/structured_notes.md`
- `.pipeline/coverage_matrix.json`

#### Stage 3: enhance

Outputs:

- `.pipeline/enhanced_notes.md`
- `final_notes.md`
- `bootcamp_index.md`

#### Optional deep pass

Outputs:

- `.pipeline/deep_pass_report.md`
- `.pipeline/deep_pass_exceptions.json`

#### Deterministic validation

Outputs:

- `.pipeline/validation_report.md`
- `.pipeline/exceptions.json`

#### Publish normalization

Effects:

- rewrite `final_notes.md` title/frontmatter/H1,
- create published tutorial filename,
- rewrite `bootcamp_index.md`.

### Fabric-Native Stage Graph

The Fabric-native port should use the following logical stages:

1. `capture_source`
2. `ingest_segments`
3. `stage1_refine_generate`
4. `stage1_refine_materialize`
5. `stage2_synthesize_generate`
6. `stage2_synthesize_materialize`
7. `stage3_enhance_generate`
8. `stage3_enhance_materialize`
9. optional `deep_pass`
10. `validate_outputs`
11. optional `publish_notes`

The reason for splitting each AI stage into `generate` + `materialize` is the main unavoidable Fabric compatibility adaptation:

- the old Zoom pipeline supports a tool-enabled mode where the model writes files directly,
- Fabric `fabric_pattern` stages return content to stdout but do not write multiple files directly,
- therefore the Fabric port should use the old tool-restricted output style and add deterministic materializer scripts that parse the AI output and write the expected artifacts.

This is the closest shippable equivalent to the original pipeline and avoids changing the deterministic file outputs.

### Asset Layout Inside Fabric

Recommended new assets:

#### Pipeline definition

- `data/pipelines/zoom-tech-note.yaml`

#### Fabric patterns

- `data/patterns/zoom_stage1_refine/`
- `data/patterns/zoom_stage2_synthesize/`
- `data/patterns/zoom_stage3_enhance/`

These patterns should preserve the existing prompt semantics. The stage prompts from the Zoom repo should be transplanted into Fabric pattern files with only the minimum adaptation needed for Fabric pattern formatting.

#### Deterministic scripts

- `scripts/pipelines/zoom-tech-note/ingest_zoom_captions.py`
- `scripts/pipelines/zoom-tech-note/validate_coverage.py`
- `scripts/pipelines/zoom-tech-note/publish_tutorial_notes.py`
- `scripts/pipelines/zoom-tech-note/deep_pass.py`
- `scripts/pipelines/zoom-tech-note/materialize_stage1.py`
- `scripts/pipelines/zoom-tech-note/materialize_stage2.py`
- `scripts/pipelines/zoom-tech-note/materialize_stage3.py`

The copied deterministic scripts should remain semantically equivalent to the Zoom originals. Any changes should be constrained to:

- relative-path resolution,
- working-directory handling,
- runner integration,
- packaging in this repository.

### Artifact Contract

The Fabric pipeline should preserve the same artifact names as the Zoom reference.

Within `.pipeline/<run-id>/` we should keep Fabric runner manifests:

- `run_manifest.json`
- `source_manifest.json`
- `run.json`

Within the session working directory `.pipeline/` we should preserve the Zoom tutorial artifacts:

- `segment_ledger.jsonl`
- `segment_manifest.jsonl`
- `refined_transcript.md`
- `topic_inventory.json`
- `corrections_log.csv`
- `uncertainty_report.json`
- `structured_notes.md`
- `coverage_matrix.json`
- `enhanced_notes.md`
- `deep_pass_report.md`
- `deep_pass_exceptions.json`
- `validation_report.md`
- `exceptions.json`

Learner-facing outputs remain:

- `final_notes.md`
- published tutorial filename
- `bootcamp_index.md`

Note: the Fabric runner's ephemeral `.pipeline/<run-id>/` is separate from the Zoom session artifact folder. The implementation should root session outputs in the invocation directory exactly as the old pipeline expects, while the runner keeps its own transient run-state metadata isolated.

### Pass / Fail Semantics

The Fabric-native port should preserve these outcomes:

- ingest failure: run fails,
- any AI stage missing required materialized outputs: run fails,
- deep pass enabled and failing: run fails before validation,
- deterministic validation failure: run fails,
- publish failure after a valid `final_notes.md`: run returns failure state for publish but should still preserve the validated final note.

We should preserve the Zoom-style manifest status vocabulary as closely as practical:

- `pass`
- `fail`
- `pass_with_publish_warning`
- deep-pass failure status

### Required Compatibility Adaptations

These are required to ship the pipeline inside Fabric and are not optional product changes:

1. AI stages must run in Fabric through patterns, so multi-file output must be materialized by deterministic helper scripts.
2. prompt files must be expressed as Fabric patterns rather than free markdown files.
3. copied deterministic scripts must be vendored into Fabric because the feature cannot depend on a private local checkout.

These are not intended behavior changes. They are packaging and orchestration adaptations needed to keep the same observable outputs.

### Improvement Candidates Requiring User Awareness

These should not be applied silently:

1. tightening any of the Zoom prompts beyond formatting needed for Fabric patterns,
2. changing any artifact names,
3. changing the deep-pass thresholds,
4. changing class naming logic,
5. changing validation rules,
6. replacing copied deterministic Python logic with new Go builtins.

If we do any of these, we should call them out explicitly first.

### Open Design Fork

The one material decision still open is how to preserve the old optional deep-pass behavior.

The current Fabric pipeline runner does not yet support pipeline-specific conditional stages or user-provided pipeline execution flags beyond the global CLI flags.

There are two clean options:

1. ship two built-in pipelines:
   - `zoom-tech-note`
   - `zoom-tech-note-deep-pass`
2. extend Fabric core to support pipeline parameters or conditional stages before shipping this pipeline.

Recommendation:

- ship two built-in pipelines first,
- keep the runner core unchanged,
- preserve the user-facing choice without blocking delivery on new orchestration features.

### Implementation Order

1. Copy deterministic Zoom scripts into Fabric-owned pipeline scripts.
2. Create Fabric patterns that preserve Stage 1, Stage 2, Stage 3 prompt behavior.
3. Add materializer scripts for Stage 1, Stage 2, Stage 3.
4. Add `zoom-tech-note` pipeline YAML.
5. Add `zoom-tech-note-deep-pass` pipeline YAML or equivalent deep-pass variant.
6. Add regression fixtures and tests for artifact creation, validation, and publish normalization.
7. Only after parity works, consider approved improvements.
