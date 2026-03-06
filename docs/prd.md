# Product Requirements Document: Fabric Pipeline UX for Note Generation

Status: Draft v1

Owner: TBD

## 1. Summary

This product adds a pipeline-oriented note-generation experience on top of Fabric. The feature allows users to run technical and non-technical note workflows from stdin, files, folders, or URLs while seeing the output of each stage in the terminal UI.

The key product decision is that Fabric remains the engine. We are not creating a new AI runtime. Instead, we are adding a Fabric-native way to execute note pipelines, including support for existing external pipelines such as the Zoom technical pipeline.

## 2. Background

Fabric already offers a strong prompt-centric workflow:

- pattern-based prompting
- stdin-driven usage
- URL scraping
- streaming output
- custom patterns

That makes it excellent for one-step transformations, but the target user experience here is broader. The desired experience is:

- one command to start a note pipeline
- visible stage-by-stage execution
- recoverable artifacts
- support for both quick note generation and richer multi-stage workflows

There is also an explicit constraint:

- do not reinvent the wheel
- do not replace existing pipelines unnecessarily
- make changes in the Fabric repository only

This means the product should be designed so that Fabric becomes the operator-facing surface while existing working pipelines can be reused through adapters.

## 3. Problem

Today, note-generation workflows are fragmented.

Users may have:

- a strong standalone pipeline in one repository
- prompt templates in another place
- no consistent pipe-first terminal UX
- no stage-by-stage UI
- no unified command surface

This creates three user problems:

1. the user experience is inconsistent
2. existing working pipelines are difficult to access ergonomically
3. quick single-shot commands and full deterministic runs are not unified

## 4. Objective

Build a Fabric-native pipeline experience for note generation that:

- feels as easy as a Unix pipe
- preserves stage visibility
- supports technical and non-technical workflows
- reuses Fabric features and external pipelines instead of replacing them

### Success criteria

Primary success means a user can do both of the following without extra glue code:

```bash
cat input.txt | fabric --pattern techNote --stream
fabric --pipeline zoom-tech-note --source ./session.txt
```

And in both cases the user can understand what happened.

## 5. Target Users

### Primary user

An operator who works from the terminal and wants a fast path from source material to useful notes.

Typical sources:

- lecture transcripts
- meeting notes
- therapy or journal exports
- articles
- web pages

### Secondary user

A builder or maintainer who already has an external pipeline and wants to surface it through Fabric without rewriting it.

## 6. User Needs

Users need:

- one easy command surface
- support for stdin, file, folder, and URL input
- both quick mode and full pipeline mode
- stage-by-stage visibility
- output they can trust
- a way to keep existing pipelines intact
- a way to add custom pipelines without editing Fabric source
- a way to run external stage logic regardless of implementation language

## 7. Product Principles

1. Fabric first
   - Reuse Fabric before introducing new runtime concepts.

2. Preserve what already works
   - Existing external pipelines should be wrapped if they are already valuable.

3. Keep the first command simple
   - The default path should require minimal configuration.

4. Make progress visible
   - Long-running jobs should never feel opaque.

5. Separate prompting from orchestration
   - Patterns solve prompting.
   - Pipelines solve sequencing, validation, and artifacts.

## 8. User Stories

### Quick technical note generation

As a terminal-first user, I want to pipe a technical transcript into a command and get a note immediately, so I can process source material without setting up a workflow manually.

### Quick non-technical note generation

As a user with conceptual or reflective text, I want a non-technical note command that does not assume code or formulas, so I get an appropriate output style.

### Full pipeline execution

As a power user, I want to run a named multi-stage pipeline and see the result of every stage, so I can trust and debug the workflow.

### External pipeline reuse

As a maintainer of an existing pipeline, I want Fabric to call my current pipeline as-is, so I avoid rewriting proven logic.

### Artifact inspection

As an operator, I want manifests and validation reports, so I can inspect failures and recover runs confidently.

## 9. Scope

### In scope for V1

- a Fabric-native pipeline execution model
- technical and non-technical note workflows
- stdin, file, folder, and URL inputs
- stage-by-stage terminal visibility
- support for Fabric pattern stages
- support for external command stages
- support for user-defined custom pipeline definitions
- support for language-agnostic script and executable stages
- deterministic artifact output for full runs
- output to stdout and files

### Out of scope for V1

- graphical UI
- remote orchestration service
- distributed workers
- rewriting external pipelines into Fabric patterns
- generalized workflow automation unrelated to note generation

## 10. Requirements

### Functional requirements

1. The product must support a pipe-first CLI workflow.
2. The product must support URL-based source ingestion.
3. The product must support file and directory sources.
4. The product must provide quick note commands.
5. The product must provide a full pipeline run command.
6. The product must display stage-by-stage progress in the terminal.
7. The product must persist artifacts for full pipeline runs.
8. The product must support multiple profiles.
9. The product must support wrapping existing external pipelines.
10. The product must support output to stdout or files.
11. The product must support loading custom pipeline definitions from a user-controlled path.
12. The product must support command stages that run arbitrary host-language executables or scripts.

### Non-functional requirements

1. Fabric remains the execution engine.
2. The architecture should minimize invasive changes to Fabric.
3. The design must be upgrade-safe relative to upstream Fabric.
4. The user experience must remain low-friction.
5. Sensitive data must remain local by default.
6. Failures must be stage-specific and inspectable.
7. The system must be extensible to new profiles and new pipeline definitions.
8. The pipeline definition format must be versioned for future compatibility.
9. User-defined pipelines must not be overwritten by Fabric upgrades.

## 11. Solution Overview

The product introduces a Fabric-native pipeline layer.

Instead of forcing every workflow into one pattern, the system defines:

- pipeline definitions
- a pipeline runner
- stage executors

Each stage can be one of:

- a Fabric pattern stage
- a built-in deterministic stage
- an external command stage

This gives us the best of both worlds:

- Fabric remains central
- existing pipelines can be reused
- stage UI and artifacts become possible

## 12. Why This Is Better Than the Alternatives

### Why not patterns only?

Patterns are good for one-shot transformations. They are not enough for stage orchestration, manifests, and validation.

### Why not executable aliases only?

Aliases improve command ergonomics, but they do not create a real pipeline product.

### Why not rewrite the Zoom pipeline into Fabric?

That violates the “do not reinvent the wheel” constraint and creates avoidable migration risk.

## 13. Proposed UX

### Quick mode

```bash
cat input.txt | fabric --pattern techNote --stream
cat input.txt | fabric --pattern nontechNote -o notes.md
fabric --scrape_url https://example.com --pattern techNote --stream
```

### Full pipeline mode

```bash
fabric --pipeline zoom-tech-note --source ./session.txt
fabric --pipeline nontech-note --source ./session.txt
```

### Stage output model

The operator sees stage status as the run progresses:

```text
[1/6] intake ........ PASS
[2/6] normalize ..... PASS
[3/6] semantic-map .. RUNNING
<streamed output>
[3/6] semantic-map .. PASS
[4/6] render ........ PASS
[5/6] validate ...... PASS
[6/6] publish ....... PASS
```

## 14. Success Metrics

### Adoption metrics

- number of pipeline runs per week
- percentage of runs using quick mode vs full run mode
- number of users creating custom pipeline definitions

### Quality metrics

- successful run rate
- validation pass rate
- percentage of failures with actionable stage-specific error messages

### UX metrics

- time from source input to first useful output
- number of operator commands required for a typical run

## 15. Risks

### Risk 1: Overbuilding

If the system becomes a generic workflow engine, it will become too broad and hard to ship.

Mitigation:

- constrain V1 to note-generation pipelines only

### Risk 2: Too much Fabric core churn

If the solution rewrites major Fabric internals, upgrades become painful.

Mitigation:

- keep the pipeline layer modular
- reuse existing Fabric primitives

### Risk 3: External pipeline integration becomes brittle

Existing pipelines may vary in behavior and artifact shape.

Mitigation:

- define adapter boundaries clearly
- use explicit command executor contracts

### Risk 4: Users confuse strategies with pipelines

Fabric already has strategies, but strategies are prompt modifiers, not multi-stage workflows.

Mitigation:

- document the distinction clearly
- keep the pipeline UX separate from strategy selection

## 16. Release Plan

### Release 1

- pipeline definition format
- built-in and custom pipeline definition discovery
- pipeline runner
- stage UI renderer
- Fabric pattern stage executor
- external command stage executor

### Release 2

- Zoom technical pipeline adapter
- validation and publish builtins
- quick technical and non-technical note commands

### Release 3

- richer JSON run events
- more profiles
- better output destinations and integration hooks

## 17. Resolved Product Decisions and Remaining Question

Resolved product decisions:

1. The operator-facing command remains `fabric`.
2. Multi-stage execution will use `fabric --pipeline ...`.
3. Quick single-step usage will continue to use `fabric --pattern ...`.
4. Built-in pipeline definitions will live in `data/pipelines/`.
5. User-defined pipeline definitions will live in `~/.config/fabric/pipelines/`.

Remaining question:

1. How much of the Zoom pipeline contract should be normalized vs simply wrapped?

## 18. Decision

The product should be built as a Fabric-native pipeline experience, not as a new external runtime.

Fabric should remain the core engine.

The pipeline layer should add:

- stage orchestration
- stage visibility
- artifact persistence
- custom pipeline definition loading
- language-agnostic command execution
- reusable adapters for existing external pipelines

That is the most direct way to meet the user requirement while preserving Fabric and avoiding reinvention.
