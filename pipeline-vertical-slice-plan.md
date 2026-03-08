# Fabric Pipeline Vertical Slice

## Goal
Build the first end-to-end Fabric pipeline slice: YAML-backed pipeline loading, preflight validation, pipeline discovery, validation commands, and a minimal `--pipeline` runner skeleton.

## Tasks
- [x] Add `internal/pipeline` types and loader for `.yaml` pipeline definitions from `data/pipelines/` and `~/.config/fabric/pipelines/` -> Verified by loader/unit coverage and current discovery behavior.
- [x] Add preflight validation for required top-level fields, stage IDs, executor types, final output rules, and executor-specific config -> Verified by structural and preflight tests, including env/executable/pattern checks.
- [x] Add CLI flags and handlers for `--listpipelines`, `--validate-pipeline <file>`, and `--pipeline <name> --validate-only` -> Verified by current CLI behavior and regression coverage.
- [x] Add a minimal pipeline runner skeleton for `--pipeline <name>` with source selection, `.pipeline/<run-id>/` creation, run manifest writing, and sequential stage loop placeholders -> Completed and expanded beyond the skeleton into real executor-backed execution.
- [x] Add one built-in sample pipeline under `data/pipelines/` for verification and smoke-test the new command surface -> `data/pipelines/passthrough.yaml` exists and is runnable.
- [x] Run focused tests for the new package and CLI behavior -> Verified by focused package tests and full `go test ./...`.

## Done When
- [x] Fabric can discover pipelines, validate them, and start a named pipeline run through the new CLI surface without touching the existing chat flow.

## Checkpoint

Status as of 2026-03-07:

- The original vertical slice is complete.
- The implementation moved beyond the original slice in three commits:
  - `38da8c0b` initial pipeline runner slice
  - `96bed3fb` stage roles and builtins
  - `f886b511` runner contract tightening and regression fixes

Implemented beyond the original slice:

- `command`, `fabric_pattern`, and `builtin` executors
- stage-role semantics for `validate` and `publish`
- stdout/stderr chaining behavior
- `-o` behavior that still works when a later publish stage fails
- preflight validation beyond schema shape:
  - env expansion
  - command executable resolution
  - pattern/context/strategy existence checks
- temporary `.pipeline/<run-id>/` lifecycle with delayed cleanup and startup janitor
- source handling for:
  - stdin
  - `--source` file
  - `--source` directory
  - `--scrape_url`
- explicit primary-output semantics across stdout and artifact-backed stages
- publish-manifest finalization from the finalized run snapshot
- CLI regression coverage for publish-failure output-file behavior

Still left after this slice:

- built-in note pipelines beyond quick mode and Zoom parity
- remaining promised pipeline profiles:
  - `nontechnical-study-guide`
- richer operator-facing examples and authoring guidance
- pipeline dry-run / introspection mode
- partial-stage execution controls
- JSON event stream mode

Known issues at this checkpoint:

- No confirmed correctness bug remains in the implemented core slice after the current audit and `go test ./...`.
- The remaining gaps are feature-completeness gaps, not known correctness regressions in the shipped pipeline kernel.

## Zoom Tech Note Extension Checkpoint

Status as of 2026-03-07 after the next major slice:

- Fabric now includes two real built-in parity pipelines:
  - `zoom-tech-note`
  - `zoom-tech-note-deep-pass`
- The parity implementation is Fabric-owned and shippable:
  - deterministic Zoom scripts were copied into `scripts/pipelines/zoom-tech-note/`
  - Stage 1/2/3 prompts were ported into Fabric patterns under `data/patterns/`
  - Fabric-only materializer scripts bridge multi-artifact AI stages back into the original Zoom artifact contract
- The implementation preserves the original Zoom stage flow and stable artifacts while still using the Fabric pipeline runner for orchestration.

Implemented in this extension slice:

- source preparation for file, directory, and stdin-backed session inputs
- copied ingest, validation, and publish scripts
- copied and isolated deep-pass logic
- Stage 1/2/3 materialization that writes the expected stable artifacts into the real session directory
- pipeline-owned pattern file resolution through relative paths
- parity-focused deterministic tests for:
  - prepare
  - ingest
  - Stage 1 materialization
  - Stage 2 materialization
  - Stage 3 materialization
  - deep pass
  - validation
  - publish
- CLI preflight validation for both built-in Zoom pipelines
- full repo verification with `go test ./...`
- built-in quick-mode note patterns aligned with the note product surface:
  - `data/patterns/techNote/system.md`
  - `data/patterns/nontechNote/system.md`
- deterministic built-in pattern tests for:
  - `techNote`
  - `nontechNote`

What remains after the Zoom parity slice:

- first-class built-in note pipelines beyond quick mode and Zoom parity remain incomplete as a category, but the first generic profile is now implemented:
  - implemented:
    - `technical-study-guide`
  - remaining:
    - `nontechnical-study-guide`
- richer operator-facing authoring documentation and examples
- dry-run / introspection mode
- JSON event stream mode
- partial stage execution controls

Known status after the Zoom parity slice:

- The current verification set covers the implemented runner contracts and the shipped Zoom parity slice after:
  - `python3 -m py_compile scripts/pipelines/zoom-tech-note/*.py`
  - `go run ./cmd/fabric --validate-pipeline data/pipelines/zoom-tech-note.yaml`
  - `go run ./cmd/fabric --validate-pipeline data/pipelines/zoom-tech-note-deep-pass.yaml`
  - `go test ./internal/plugins/db/fsdb -run 'TestBuiltinQuickNotePatterns' -count=1`
  - `go test ./internal/pipeline -run 'TestZoomTechNote' -count=1`
  - `go test ./...`
- Phase 2 is now complete on the current branch.
- The main remaining work is product-surface expansion plus ongoing parity hardening, not a new runner architecture phase.
