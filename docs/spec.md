# Fabric Pipeline Execution Specification

Status: Draft v1

Purpose: Define a Fabric-native way to run multi-stage note-generation pipelines while preserving Fabric as the execution engine and allowing existing external pipelines, such as the Zoom technical pipeline, to be reused without modification.

## 1. Problem Statement

Fabric is already strong at single-shot prompt execution:

- stdin -> pattern -> output
- URL -> scrape -> pattern -> output
- optional streaming, output files, and prompt strategies

That solves the prompt layer, but it does not yet provide a first-class pipeline surface for:

- multi-stage note generation
- visible stage-by-stage progress in the UI
- deterministic artifacts such as manifests and validation reports
- reuse of existing external pipelines without rewriting them into one prompt

The system to be built must let an operator invoke note-generation pipelines through a simple command surface while reusing Fabric where Fabric is already strongest.

Important boundary:

- Fabric remains the prompt runtime and command host.
- Pipelines define how stages are sequenced and how outputs are persisted.
- Existing external pipelines may be wrapped through adapters instead of being rewritten.

## 2. Goals and Non-Goals

### 2.1 Goals

- Provide a pipe-first pipeline UX on top of Fabric.
- Support stdin, file path, directory path, and URL inputs.
- Preserve stage boundaries and show each stage’s output in the terminal UI.
- Reuse existing Fabric concepts wherever possible:
  - patterns
  - strategies
  - streaming
  - output files
  - custom patterns
- Reuse external pipelines, such as the Zoom technical pipeline, without modifying them.
- Support both quick single-step note generation and full artifact-producing pipeline runs.
- Persist pipeline artifacts in a deterministic per-run structure during execution and short-lived post-run inspection.
- Keep the design extensible across technical and non-technical note workflows.

### 2.2 Non-Goals

- Replacing Fabric with a separate AI runtime.
- Rewriting external pipelines into Fabric-only prompts for V1.
- Building a web UI in V1.
- Building a distributed workflow scheduler.
- Creating a fully generic automation framework unrelated to note-generation pipelines.
- Forcing a new mandatory global settings system beyond Fabric’s existing configuration model.

## 3. Design Principles

1. Do not reinvent the wheel.
   - Use Fabric as the prompt runtime.
   - Extend Fabric only where a real gap exists.

2. Preserve working systems.
   - Existing external pipelines should be wrapped, not reauthored, when possible.

3. Separate orchestration from prompting.
   - Patterns are for prompt behavior.
   - Pipelines are for sequencing, artifacts, validation, and publish flow.

4. Default to local and inspectable.
   - Operators must be able to see stage progress, failures, and written files.

5. Keep the architecture upgrade-safe.
   - Minimize invasive changes to Fabric internals.

## 4. Core Requirements

### 4.1 Functional Requirements

#### FR1. Pipe-first invocation

The system must accept piped stdin as a first-class source.

Examples:

```bash
cat lecture.txt | fabric --pattern techNote --stream
cat therapy.txt | fabric --pattern nontechNote -o notes.md
```

#### FR2. URL-driven invocation

The system must accept URL input and route it through Fabric’s scrape flow where applicable.

Examples:

```bash
fabric --scrape_url https://example.com/article --pattern techNote --stream
```

#### FR3. File and directory invocation

The system must accept direct file or directory paths for full pipeline runs.

Examples:

```bash
fabric --pipeline zoom-tech-note --source ./session.txt
fabric --pipeline nontech-note --source ./MyTherapySessions
```

#### FR4. Stage-aware execution

The system must execute and report an explicit, linear stage model.

At minimum:

- every stage must declare an explicit `id`
- every stage must declare an explicit executor type
- stage names are conventions, not reserved semantics
- exactly one stage must be designated as the final output stage in V1

#### FR5. Stage UI visibility

The terminal UI must show:

- current stage
- stage status
- stream output where applicable
- files written by the stage
- final PASS or FAIL per stage

#### FR6. Multi-mode execution

The system must support:

- quick prompt execution via `--pattern`
- full multi-stage execution via `--pipeline`
- future partial stage execution without redesigning the model

#### FR7. External pipeline compatibility

The system must support using existing external pipelines as-is through stage adapters.

At minimum, the design must support wrapping the current Zoom technical pipeline without changing the Zoom pipeline repository.

#### FR8. Artifact preservation

Full pipeline runs must write Fabric-managed artifacts under:

- `.pipeline/<run-id>/`

in the current working directory where `fabric` was invoked.

At minimum, the system must support run-scoped artifacts such as:

- `run_manifest.json`
- `source_manifest.json`
- declared named stage artifacts
- validation artifacts when validation exists
- optional publish artifacts

These artifacts are temporary by default and must follow the agreed hybrid cleanup model.

#### FR9. Profile support

The system must support multiple note-generation profiles, at minimum:

- `technical-study-guide`
- `nontechnical-study-guide`

#### FR10. Output control

The operator must be able to choose:

- stdout final output
- `-o/--output` for the final rendered file

Pipeline-managed artifacts remain under `.pipeline/<run-id>/` in the current working directory and are not user-routed through a separate artifact-root flag in V1.

#### FR11. Custom pipeline support

The system must support user-defined pipelines in addition to built-in pipelines.

At minimum:

- built-in pipeline definitions must live in the Fabric repository
- user-defined pipeline definitions must live in `~/.config/fabric/pipelines/`
- user-defined pipelines must be loadable without recompiling Fabric
- user-defined pipelines may override built-in pipelines by name
- pipeline discovery must be available through `fabric --listpipelines`
- validation must be available through both:
  - `fabric --validate-pipeline <file>`
  - `fabric --pipeline <name> --validate-only`

#### FR12. Polyglot command execution

The system must support command stages that can run executables or scripts written in any language available on the host system.

At minimum, the command executor must support:

- explicit executable plus argument lists
- stdin passthrough where required
- working directory selection
- environment variable injection
- timeout control
- stdout and stderr capture
- exit-code based failure handling

Examples include:

- `python3 script.py`
- `bash script.sh`
- `java -jar tool.jar`
- `node script.mjs`

#### FR13. Schema versioning

Pipeline definitions must declare a schema version so the format can evolve without breaking existing custom pipelines.

#### FR14. Pipeline definition language

Pipeline definitions must:

- be authored in YAML
- use the `.yaml` extension in V1
- declare a mandatory top-level `name`
- match that `name` exactly to the filename stem
- be validated against a versioned structured schema plus semantic preflight checks

### 4.2 Non-Functional Requirements

#### NFR1. Fabric-first architecture

Fabric must remain the core prompt runtime.

#### NFR2. Minimal invasive change

V1 should avoid deep changes to Fabric’s core prompt execution path unless required by a validated gap.

#### NFR3. Upgrade safety

The solution should be resilient to upstream Fabric updates.

#### NFR4. Observability

Runs must be inspectable without requiring a debugger.

#### NFR5. Reproducibility

Deterministic stages must produce deterministic artifacts from the same inputs.

#### NFR6. Privacy

Sensitive local content, such as therapy sessions, must remain local by default.

#### NFR7. Extensibility

New profiles and stage executors must be addable without redesigning the system.

#### NFR8. Low operator overhead

The default UX must remain easy to remember and practical for daily use.

#### NFR9. Explicit failure semantics

The system must fail with stage-specific errors instead of a generic command failure.

#### NFR10. Scalability

The architecture must scale by adding new pipeline definitions, profiles, and executor types without requiring changes to the core execution path for every new workflow.

#### NFR11. Upgrade-safe customization

User-added pipelines must remain separate from built-in pipelines so that upstream Fabric updates do not overwrite local pipeline definitions.

## 5. Architectural Options

### 5.1 Option A: Fabric patterns only

Everything is implemented as custom Fabric patterns and invoked directly through `fabric --pattern`.

Pros:

- lowest effort
- maximum reuse of current Fabric behavior
- simple to prototype

Cons:

- patterns are not true orchestration units
- hard to express stage UI and stage artifacts cleanly
- poor fit for wrapping existing external pipelines

Assessment:

- useful for simple note commands
- insufficient for the full requirement

### 5.2 Option B: Executable patterns only

Pattern aliases or generated shell commands behave like executables.

Pros:

- better command ergonomics
- very low engineering effort

Cons:

- still command sugar only
- does not solve orchestration
- does not solve artifacts, validation, or publish logic

Assessment:

- good as convenience
- not enough as the primary architecture

### 5.3 Option C: Fabric-native pipeline runner with stage executors

Fabric gains a pipeline subsystem that runs a declarative pipeline definition. Each stage is executed by one of several executor types.

Pros:

- keeps Fabric as the host runtime
- supports stage UI and artifacts
- supports external pipeline reuse
- scales to technical and non-technical profiles

Cons:

- more engineering work than aliases or patterns alone
- introduces a new orchestration layer inside the repo

Assessment:

- best fit for the requirement
- recommended

### 5.4 Option D: Fabric fork with custom subcommand redesign

Large CLI redesign where Fabric becomes a general workflow host.

Pros:

- deepest integration

Cons:

- highest maintenance risk
- large surface area
- premature for V1

Assessment:

- not recommended

## 6. Recommended Architecture

Adopt Option C: Fabric-native pipeline runner with stage executors.

The central idea is that Fabric should add a declarative pipeline model rather than treating multi-stage work as one prompt or one shell alias.

### 6.1 Core components

1. `Pipeline Definition Loader`
   - Loads pipeline definitions from Fabric-managed paths.
   - Parses top-level metadata, accepted input modes, stage graph, and output contract.

2. `Pipeline Runner`
   - Owns stage sequencing.
   - Handles stage start, completion, failure, cancellation, and summary.
   - Writes run manifests.

3. `Stage Executor Registry`
   - Maps each stage to an executor implementation.

4. `Fabric Pattern Executor`
   - Runs a Fabric pattern as a stage.
   - Supports stdin input, stream mode, variables, and strategies.

5. `Builtin Stage Executor`
   - Runs deterministic internal logic such as:
     - manifest creation
     - file writing
     - section validation
     - publish bookkeeping

6. `External Command Executor`
   - Runs an external command stage.
   - Allows wrapping existing pipelines or scripts without changing them.
   - Must be language-agnostic and process-oriented rather than Python-specific.

7. `Artifact Manager`
   - Owns `.pipeline/<run-id>/`, run-scoped manifests, declared artifact references, and cleanup lifecycle.

8. `Terminal Status Renderer`
   - Prints live stage status, stream output, and run summaries.

### 6.2 Stage executor model

Each stage must declare an executor type:

- `fabric_pattern`
- `builtin`
- `command`

This is the key design decision that allows:

- direct reuse of Fabric prompt execution
- direct reuse of existing external pipelines
- deterministic in-process helpers for artifacts and validation

## 7. Pipeline Definition Model

### 7.1 Pipeline definition purpose

A pipeline definition is the contract that tells Fabric how to run a named multi-stage workflow.

It must declare:

- version
- pipeline name
- accepted input types when restricted
- stages in order
- optional stage role when post-final-output semantics matter
- executor for each stage
- stage inputs and outputs
- artifact policy
- final output behavior

Stage role semantics in V1:

- `role: validate` marks a stage whose success gates final stdout emission
- `role: publish` marks a downstream operational stage that may fail after validated output exists
- built-in stages may imply these roles even when `role` is omitted:
  - `validate_declared_outputs` behaves as `validate`
  - `write_publish_manifest` behaves as `publish`
- any validate or publish stage must appear after the single `final_output` stage
- publish stages must appear after the last validate stage

### 7.2 Logical example

```yaml
version: 1
name: zoom-tech-note
description: Technical lecture-note pipeline
accepts:
  - stdin
  - source
  - scrape_url
stages:
  - id: intake
    executor: builtin
    builtin:
      name: source_capture
  - id: normalize
    executor: command
    input:
      from: source
    command:
      program: python3
      args:
        - /path/to/existing/zoom/intake.py
        - "{{source}}"
      cwd: /path/to/existing/zoom
      timeout: 600
    artifacts:
      - name: normalized_source
        path: normalized_source.md
  - id: semantic_map
    executor: fabric_pattern
    pattern: tech_semantic_map
    stream: true
    primary_output:
      from: stdout
  - id: render
    executor: fabric_pattern
    pattern: tech_note
    stream: true
    final_output: true
    primary_output:
      from: stdout
  - id: validate_contract
    role: validate
    executor: builtin
    builtin:
      name: validate_declared_outputs
  - id: publish
    role: publish
    executor: builtin
    builtin:
      name: write_publish_manifest
```

This is the intended V1 direction. Implementation may add small field-level details, but the semantics here are now the source-of-truth model.

### 7.3 Definition discovery and override rules

Fabric must support two pipeline definition locations:

- built-in definitions: `data/pipelines/`
- user definitions: `~/.config/fabric/pipelines/`

Rules:

- user definitions override built-in definitions with the same name
- both locations must use the same schema
- invalid user definitions must fail with actionable validation errors

## 8. CLI Surface

### 8.1 Recommended command structure

The operator-facing command must remain `fabric`.

V1 will use:

- `fabric --pattern ...` for quick prompt execution
- `fabric --pipeline ...` for multi-stage pipeline execution
- `fabric --listpipelines` for discovery
- `fabric --validate-pipeline ...` and `fabric --pipeline ... --validate-only` for preflight validation

### 8.2 Required command families

#### Quick prompt mode

```bash
fabric --pattern techNote
fabric --pattern nontechNote
```

These remain the fast-entry mode.

#### Full pipeline mode

```bash
fabric --pipeline zoom-tech-note --source ./session.txt
fabric --pipeline nontech-note --source ./session.txt
```

This is the orchestrated mode with artifacts and stage visibility.

#### Advanced future mode

Future versions may add partial execution controls such as `--from-stage` and `--to-stage`, but these are not required for V1.

### 8.3 Input modes

Required:

- stdin
- `--source`
- `--scrape_url`

### 8.4 Output modes

Required:

- stdout
- `-o/--output`
- automatic `.pipeline/<run-id>/` artifact management in the current working directory

## 9. Stage UI Requirements

The operator must see the progression of a run in real time.

Minimum terminal rendering:

```text
[1/6] intake ........ PASS
[2/6] normalize ..... PASS
[3/6] semantic-map .. RUNNING
<streamed output>
[3/6] semantic-map .. PASS
[4/6] render ........ RUNNING
<streamed output>
[4/6] render ........ PASS
[5/6] validate ...... PASS
[6/6] publish ....... PASS
```

These labels are illustrative. Stage names are not reserved semantics; Fabric renders the declared stage IDs.

Required behaviors:

- stage start marker
- stage completion marker
- stage failure marker
- files written summary
- final run summary

## 10. Artifact Contract

For full pipeline mode, the system must create `.pipeline/<run-id>/` in the current working directory where `fabric` was invoked.

Required V1 artifacts:

- `run_manifest.json`
- `source_manifest.json`
- declared named artifacts referenced by the pipeline
- validation outputs when a validation stage or builtin validation exists

Recommended optional artifacts:

- `stage_logs/`
- `publish_manifest.json`
- `exceptions.json`

Lifecycle rule:

- `.pipeline/<run-id>/` is temporary by default
- cleanup uses the agreed hybrid cleanup model
- startup janitor cleanup and delayed per-run cleanup are both part of the design

## 11. Fabric Features to Reuse Directly

The implementation should reuse current Fabric behavior for:

- custom patterns
- prompt strategies
- URL scraping
- stdin handling
- stream output
- output files

Important note:

- Fabric strategies are prompt modifiers, not orchestration definitions.
- Strategies may improve stage quality but do not replace the pipeline model.

## 12. External Pipeline Reuse and Shippable Built-Ins

Fabric must support both:

- user-authored external pipelines executed through `command` stages
- Fabric-owned built-in pipelines that are fully shippable from this repository alone

For local experimentation, existing external pipelines may be reused directly through `command` stages.

For built-in user-facing pipelines that Fabric ships, local absolute-path dependencies are not acceptable. Those pipelines must live entirely inside the Fabric repository.

This is the rule that now governs the Zoom technical pipeline:

- the original Zoom pipeline remains the reference behavior
- the shippable Fabric implementation is a parity port owned by this repo
- deterministic scripts may be copied into Fabric unchanged where possible
- AI stages may be adapted into `generate -> materialize` pairs when Fabric execution semantics require it

This preserves behavior without forcing Fabric to depend on a user’s private local Zoom checkout.

The same `command` executor model must still support any host-language runtime that can be invoked as a process, including Python, Bash, Java, Node.js, and compiled binaries.

## 13. Error Handling

The system must report failures in a stage-specific manner.

Examples:

- pattern not found
- invalid pipeline definition
- external command failed
- output artifact missing
- validation failed

Failure output must include:

- stage ID
- executor type
- error summary
- relevant file path or command
- exit code if applicable

## 14. Security and Privacy

The system must default to local execution and local file output.

Sensitive content must not be auto-published.

External command execution must be explicit in the pipeline definition and visible in run logs.

## 15. Rollout Plan

### Phase 1

- define a versioned pipeline definition format
- implement built-in and custom pipeline loading
- implement direct Fabric pattern stages
- implement external command stages
- implement terminal stage renderer
- implement basic `fabric --pipeline` flow

Current status:

- implemented on the feature branch
- includes stronger behavior than the original phase description:
  - preflight validation
  - cleanup lifecycle
  - output-file behavior
  - stage-role-aware runtime semantics

### Phase 2

- ship a Fabric-owned Zoom technical parity pipeline
- add validation and publish builtins
- align quick note patterns with pipeline mode

Current status:

- fully implemented
- done:
  - `zoom-tech-note`
  - `zoom-tech-note-deep-pass`
  - validation builtin
  - publish builtin
  - built-in quick note patterns aligned with pipeline mode:
    - `techNote`
    - `nontechNote`

### Phase 3

- add richer status output
- add JSON event stream mode
- add more profiles and pipeline definitions
- add partial stage execution controls

Current status:

- largely deferred
- the current implementation keeps the terminal renderer human-readable and stderr-based, but does not yet expose:
  - JSON event stream mode
  - partial execution flags
  - richer real profile inventory

## 16. Implementation Checkpoint

Checkpoint date: 2026-03-07

Implemented and verified in the current branch:

- YAML-backed pipeline definitions under:
  - `data/pipelines/`
  - `~/.config/fabric/pipelines/`
- pipeline discovery with `fabric --listpipelines`
- shell-completion-friendly listing with `--shell-complete-list`
- preflight validation via:
  - `fabric --validate-pipeline <file>`
  - `fabric --pipeline <name> --validate-only`
- exactly-one-source enforcement for pipeline mode:
  - stdin
  - `--source`
  - `--scrape_url`
- pipeline runner with:
  - `builtin`
  - `command`
  - `fabric_pattern`
- stdout/stderr split suitable for chaining
- temporary `.pipeline/<run-id>/` artifacts in the invocation directory
- delayed cleanup plus startup janitor cleanup
- explicit primary-output handling across stdout and artifact-backed stages
- stage-role-aware behavior for validation gating and publish-after-output semantics
- output-file behavior that remains correct when a later publish stage fails
- built-in sample pipeline: `passthrough`
- built-in Zoom parity pipelines:
  - `zoom-tech-note`
  - `zoom-tech-note-deep-pass`
- built-in quick-mode note patterns aligned with the real note surface:
  - `techNote`
  - `nontechNote`
- Fabric-owned Zoom pipeline assets:
  - copied deterministic scripts under `scripts/pipelines/zoom-tech-note/`
  - Fabric patterns under `data/patterns/zoom_stage*_*/system.md`
  - materializer scripts for Stage 1/2/3 multi-artifact output
- parity-focused deterministic tests for the Zoom pipeline wrappers and artifact flow
- deterministic built-in pattern tests for:
  - `techNote`
  - `nontechNote`
- relative pipeline-path resolution for stage-owned pattern files

Implemented but still minimal:

- built-in pipeline inventory
- publish integration behavior
- operator-facing examples
- built-in study-guide profile inventory:
  - `technical-study-guide`

Not yet implemented from the broader spec:

- remaining built-in note-generation pipelines such as:
  - `nontechnical-study-guide`
- richer profile inventory
- pipeline dry-run / introspection mode
- JSON event stream mode
- partial stage execution controls

Known status at this checkpoint:

- the implemented runner/kernel contracts are covered by the current audit, focused Zoom parity tests, and `go test ./...`
- the implemented quick-note alignment is covered by focused built-in pattern tests and remains part of the shipped built-in product surface
- the platform is usable, but built-in pipeline parity should still be treated as under active audit while the shipped pipeline inventory expands
- the remaining work is primarily product-surface completion plus parity hardening, not a redesign of the runner model

## 17. Decision

Fabric should be extended with a pipeline runner and stage executor model.

The system should:

- reuse Fabric for prompt execution
- use pipeline definitions for orchestration
- use stage executors to mix Fabric patterns, builtins, and external commands
- load built-in and user-defined pipeline definitions
- support language-agnostic command execution
- preserve existing external pipelines instead of rewriting them

That is the cleanest way to satisfy the product requirement while keeping Fabric central and minimizing reinvention.
