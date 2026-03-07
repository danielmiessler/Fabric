# Fabric Pipeline Vertical Slice

## Goal
Build the first end-to-end Fabric pipeline slice: YAML-backed pipeline loading, preflight validation, pipeline discovery, validation commands, and a minimal `--pipeline` runner skeleton.

## Tasks
- [ ] Add `internal/pipeline` types and loader for `.yaml` pipeline definitions from `data/pipelines/` and `~/.config/fabric/pipelines/` -> Verify: a unit test can load a valid pipeline and reject a malformed one.
- [ ] Add preflight validation for required top-level fields, stage IDs, executor types, final output rules, and executor-specific config -> Verify: validation tests cover valid pipeline, duplicate stage IDs, missing `name`, missing final output, and invalid executor configs.
- [ ] Add CLI flags and handlers for `--listpipelines`, `--validate-pipeline <file>`, and `--pipeline <name> --validate-only` -> Verify: commands print expected results and exit cleanly for valid/invalid inputs.
- [ ] Add a minimal pipeline runner skeleton for `--pipeline <name>` with source selection, `.pipeline/<run-id>/` creation, run manifest writing, and sequential stage loop placeholders -> Verify: a simple built-in sample pipeline creates a run directory, manifest, and stage progress output.
- [ ] Add one built-in sample pipeline under `data/pipelines/` for verification and smoke-test the new command surface -> Verify: `fabric --listpipelines` shows it and `fabric --pipeline <sample> --validate-only` passes.
- [ ] Run focused tests for the new package and CLI behavior -> Verify: targeted `go test` commands pass for touched packages.

## Done When
- [ ] Fabric can discover pipelines, validate them, and start a named pipeline run through the new CLI surface without touching the existing chat flow.
