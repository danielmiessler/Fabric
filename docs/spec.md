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
- Persist pipeline artifacts in a deterministic structure for audit and recovery.
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

The system must execute and report a stage model with, at minimum:

1. intake
2. normalization
3. semantic mapping
4. rendering
5. validation
6. publish

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

Full pipeline runs must persist artifacts such as:

- `final_notes.md`
- `.pipeline/run_manifest.json`
- `.pipeline/source_manifest.json`
- `.pipeline/semantic_map.json`
- `.pipeline/validation_report.md`
- optional publish artifacts

#### FR9. Profile support

The system must support multiple note-generation profiles, at minimum:

- `technical-study-guide`
- `nontechnical-study-guide`

#### FR10. Output control

The operator must be able to choose:

- stdout only
- output file
- artifact folder output

#### FR11. Custom pipeline support

The system must support user-defined pipelines in addition to built-in pipelines.

At minimum:

- built-in pipeline definitions must live in the Fabric repository
- user-defined pipeline definitions must live in a Fabric-managed config directory
- user-defined pipelines must be loadable without recompiling Fabric
- user-defined pipelines may override built-in pipelines by name

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
   - Parses metadata, profile bindings, stage graph, and output contract.

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
   - Owns `.pipeline/` directories, output paths, and stage file references.

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

- schema version
- pipeline name
- profile name
- accepted input types
- stages in order
- executor for each stage
- stage inputs and outputs
- artifact policy
- default output behavior

### 7.2 Logical example

```yaml
schema_version: 1
name: zoom-tech-note
profile: technical-study-guide
accepts:
  - stdin
  - file
  - directory
stages:
  - id: intake
    executor: builtin
  - id: normalize
    executor: command
    command:
      program: python3
      args:
        - /path/to/existing/zoom/intake.py
        - "{{source}}"
      cwd: /path/to/existing/zoom
      timeout: 10m
  - id: semantic_map
    executor: fabric_pattern
    pattern: tech_semantic_map
    stream: true
  - id: render
    executor: fabric_pattern
    pattern: tech_note
    stream: true
  - id: validate
    executor: builtin
  - id: publish
    executor: builtin
```

This is illustrative. The exact syntax may change, but the contract must preserve the same semantics.

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
- artifact folder mode

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

Required behaviors:

- stage start marker
- stage completion marker
- stage failure marker
- files written summary
- final run summary

## 10. Artifact Contract

For full pipeline mode, the system must persist a `.pipeline/` folder adjacent to the output or source session directory.

Required V1 artifacts:

- `.pipeline/run_manifest.json`
- `.pipeline/source_manifest.json`
- `.pipeline/semantic_map.json` if semantic mapping is part of the selected pipeline
- `.pipeline/validation_report.md` or `.json`
- `final_notes.md`

Recommended optional artifacts:

- `.pipeline/stage_logs/`
- `.pipeline/publish_manifest.json`
- `.pipeline/exceptions.json`

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

## 12. External Pipeline Reuse

The current Zoom technical pipeline must remain external and unchanged.

The Fabric-side implementation should wrap it through a `command` executor or compatible adapter.

This allows Fabric to become the operator-facing interface while Zoom remains the source of truth for that pipeline’s current implementation.

This is required to satisfy the “do not reinvent the wheel” constraint.

The same `command` executor model must support any host-language runtime that can be invoked as a process, including Python, Bash, Java, Node.js, and compiled binaries.

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

### Phase 2

- integrate Zoom technical pipeline adapter
- add validation and publish builtins
- align quick note patterns with pipeline mode

### Phase 3

- add richer status output
- add JSON event stream mode
- add more profiles and pipeline definitions
- add partial stage execution controls

## 16. Decision

Fabric should be extended with a pipeline runner and stage executor model.

The system should:

- reuse Fabric for prompt execution
- use pipeline definitions for orchestration
- use stage executors to mix Fabric patterns, builtins, and external commands
- load built-in and user-defined pipeline definitions
- support language-agnostic command execution
- preserve existing external pipelines instead of rewriting them

That is the cleanest way to satisfy the product requirement while keeping Fabric central and minimizing reinvention.
