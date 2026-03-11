# Pipeline Operations and Authoring

This guide is the operator-focused and maintainer-focused reference for Fabric pipeline mode.

It covers:

- running built-in pipelines safely
- inspecting pipeline execution plans and lifecycle events
- stage-slice controls for targeted runs
- authoring and validating custom YAML pipelines
- running provider-backed smoke checks for model-backed stages

## 1. Runtime Commands

List built-in plus user-installed pipelines:

```bash
fabric --listpipelines
```

Validate a specific pipeline definition file:

```bash
fabric --validate-pipeline data/pipelines/technical-study-guide.yaml
```

Validate a named pipeline without executing stages:

```bash
fabric --pipeline technical-study-guide --validate-only
```

Run a pipeline end-to-end (source from stdin):

```bash
cat transcript.txt | fabric --pipeline technical-study-guide
```

Run a pipeline from an explicit file or directory source:

```bash
fabric --pipeline note-enhancement --source /path/to/session/final_notes.md
```

Run with output persisted to a file:

```bash
cat transcript.txt | fabric --pipeline therapy-conversation-notes -o therapy_notes.md
```

## 2. Stage-Slice Controls

Use these flags for targeted execution:

- `--from-stage <stage-id>`: start from a specific stage and continue to the end
- `--to-stage <stage-id>`: run from the beginning through a specific stage
- `--only-stage <stage-id>`: run one stage only

Examples:

```bash
cat transcript.txt | fabric --pipeline zoom-tech-note --to-stage stage1_refine_generate
```

```bash
cat transcript.txt | fabric --pipeline technical-study-guide --from-stage semantic_map_materialize
```

```bash
cat transcript.txt | fabric --pipeline note-enhancement --only-stage enhance_generate
```

`--only-stage` cannot be combined with `--from-stage` or `--to-stage`.

## 3. JSON Events and Dry Run

Emit machine-readable lifecycle events on stderr:

```bash
cat transcript.txt | fabric --pipeline technical-study-guide --pipeline-events-json
```

Current event types:

- `run_started`
- `stage_started`
- `stage_passed`
- `stage_failed`
- `warning`
- `run_summary`

Inspect pipeline execution plan without stage execution:

```bash
fabric --pipeline zoom-tech-note --dry-run
```

`--dry-run` is the safest way to review stage ordering, stage inputs, and contract wiring before running model-backed stages.

## 4. Artifact Contract and Run Lifecycle

Pipeline runs create temporary execution state under:

```text
.pipeline/<run-id>/
```

Key runtime artifacts:

- `run_manifest.json`: stage-level status and file outputs
- `run.json`: run-level status and expiration metadata
- `source_manifest.json`: selected source mode/reference metadata

If a pipeline declares `final_output: true` on a stage, that stage payload becomes the run final output and can be written with `-o`.

## 5. Authoring Custom Pipelines

Custom pipelines are loaded from:

```text
~/.config/fabric/pipelines/
```

Minimal template:

```yaml
version: 1
name: my-pipeline
description: Example custom pipeline.
accepts:
  - stdin
  - source
stages:
  - id: prepare
    executor: builtin
    builtin:
      name: passthrough
    final_output: true
    primary_output:
      from: stdout
```

Supported executors:

- `builtin`
- `command`
- `fabric_pattern`

Optional stage roles:

- `validate`
- `publish`

Validation and publish ordering semantics are enforced by preflight checks and runtime contracts.

## 6. Provider-Backed Smoke Validation

Fabric ships a bounded smoke harness for model-backed built-in pipelines:

```text
scripts/pipelines/live_smoke.py
```

Local run:

```bash
go build -o fabric ./cmd/fabric
OPENAI_API_KEY=... \
python3 scripts/pipelines/live_smoke.py \
  --fabric-bin ./fabric \
  --vendor OpenAI \
  --model gpt-4.1-mini
```

Run a subset:

```bash
python3 scripts/pipelines/live_smoke.py --fabric-bin ./fabric --pipelines technical-study-guide,zoom-tech-note,zoom-tech-note-deep-pass
```

The smoke suite runs each pipeline only through its first model stage (`--to-stage`) and asserts:

- target stage started
- target stage passed
- run summary status is `passed`

CI workflow:

- `.github/workflows/pipeline-live-smoke.yml`
- supports scheduled and manual (`workflow_dispatch`) execution

## 7. Troubleshooting

If pipeline validation fails:

1. Run `fabric --validate-pipeline <file>` first.
2. Confirm referenced pattern paths and command executables exist.
3. Verify source mode is accepted by `accepts:`.

If smoke fails:

1. Check provider credentials for selected vendor.
2. Confirm selected model is available for that vendor account.
3. Re-run a single pipeline case locally with `--pipelines <name>`.
