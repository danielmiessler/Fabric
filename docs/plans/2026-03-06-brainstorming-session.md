# Brainstorming Session: Fabric Executable Pipelines

Date: 2026-03-06
Status: Validated design notes
Scope: Record the clarified design decisions for Fabric-native executable pipelines before implementation.

## Context

The goal is to add pipeline functionality inside Fabric itself rather than introducing a second binary or a separate runtime. Fabric remains the single operator-facing command.

This brainstorming session refined the execution model after reviewing:

- existing Fabric stdin and pattern behavior
- the current `docs/spec.md`
- the current `docs/prd.md`
- the need to support existing external pipelines such as the Zoom technical pipeline without rewriting them

## Confirmed Product Direction

- The operator-facing command remains `fabric`.
- Quick single-step runs continue to use `fabric --pattern ...`.
- Multi-stage runs will use `fabric --pipeline ...`.
- A second binary such as `aipipe` is not part of the design.
- Existing external pipelines should be wrapped through stage adapters rather than rewritten into patterns.

## Confirmed Pipeline Model

The executable pipeline model inside Fabric is:

- declarative pipeline definitions
- a pipeline runner
- stage executors
- deterministic artifacts
- visible stage-by-stage execution

The executor types remain:

- `fabric_pattern`
- `builtin`
- `command`

The `command` executor is explicitly language-agnostic and must support host-language runtimes such as Python, Bash, Java, Node.js, and compiled executables.

## Settled Clarification 1: Input Source Policy for `--pipeline`

Decision:

- `fabric --pipeline ...` must require exactly one input source.

Allowed sources:

- stdin
- `--source`
- `--scrape_url`

This is the strict `C` policy discussed during brainstorming.

### Valid examples

```bash
cat a.txt | fabric --pipeline tech-note
fabric --pipeline tech-note --source ./session.txt
fabric --scrape_url https://example.com/article --pipeline tech-note
```

### Invalid examples

```bash
cat a.txt | fabric --pipeline tech-note --source ./session.txt
cat a.txt | fabric --scrape_url https://example.com/article --pipeline tech-note
fabric --pipeline tech-note --source ./session.txt --scrape_url https://example.com/article
```

### Why this was chosen

- pipeline mode is orchestration, not just prompt execution
- manifests and artifacts need one clear source of truth
- silent precedence rules create confusing behavior
- strictness reduces accidental runs and improves debuggability

### Important distinction from current Fabric behavior

Current `fabric --pattern ...` mode is permissive and message-oriented. It reads stdin into the message and injects it into `{{input}}`.

New `fabric --pipeline ...` mode should be stricter because it owns:

- stage sequencing
- artifact layout
- validation
- publish flow

## Settled Clarification 2: Chaining Fabric Through Pipes

Decision:

- chaining Fabric commands must be allowed
- pipeline mode must preserve Unix pipe behavior

### Target examples

```bash
cat a.txt | fabric --pipeline tech-note | fabric --pattern summarize
curl -L https://example.com/article | fabric --pipeline nontech-note | fabric --pattern create_quiz
```

### Required contract

- `stdout` contains only the final pipeline payload
- `stderr` contains stage UI, progress, warnings, and run logs
- `.pipeline/` artifacts are written to disk and do not pollute stdout
- failure is communicated through non-zero exit codes

### Why this contract matters

If stage progress is printed to stdout, downstream Fabric commands receive status noise mixed with content and chaining breaks.

The default behavior for pipeline runs must therefore be:

- final rendered content -> stdout
- progress and stage events -> stderr
- manifests and artifacts -> filesystem

## Settled Clarification 3: Artifact Location and Automatic Cleanup

Decision:

- pipeline artifacts will be written under a `.pipeline/` folder
- `.pipeline/` will be created only in the current working directory where `fabric` is invoked
- pipeline artifacts must be automatically cleaned up after a short delay
- the agreed cleanup target is 5 seconds after run completion
- this cleanup policy applies to both successful and failed runs

### Why `.pipeline/` remains correct

The final design retains `.pipeline/` instead of a flat set of files in the current directory because:

- it groups run artifacts in one place
- it reduces visible clutter in the working directory
- it gives a stable container for manifests, semantic maps, validation reports, and other stage outputs
- it aligns with the earlier pipeline artifact model already described in the spec and PRD

### Current working directory rule

The `.pipeline/` folder must be created only in the directory from which the operator runs the Fabric pipeline command.

This means:

- no global artifact directory
- no hidden state elsewhere in the filesystem
- no node-style workspace sprawl
- no dependency on extra setup from the user

### Per-run isolation rule

Artifacts must not be written into one shared undifferentiated `.pipeline/` root.

Instead, each pipeline run must get its own run-specific artifact directory:

```text
.pipeline/<run-id>/
```

Examples:

```text
.pipeline/20260307T103015Z-a1b2c3/
.pipeline/20260307T103105Z-d4e5f6/
```

This is required to avoid:

- one run deleting another run’s artifacts
- mixed artifacts from concurrent runs
- ambiguity during debugging and cleanup

### What the user should experience

From the operator’s point of view:

1. The user runs a pipeline command.
2. The user sees stage-by-stage progress in the same terminal.
3. The final pipeline output is printed to stdout.
4. Temporary artifacts are created under `.pipeline/<run-id>/`.
5. Those artifacts are removed automatically after completion.
6. The user does not need to manually clean anything.

The user should not need to know the cleanup mechanics for the system to remain clean.

### Agreed terminal behavior

The output contract remains:

- `stdout` contains the final pipeline payload only
- `stderr` contains progress, stage markers, and diagnostics
- `.pipeline/<run-id>/` contains temporary artifacts

This ensures:

- chainability through Unix pipes
- inspectability during the run
- automatic cleanup after the run

### Agreed cleanup target

The agreed retention window is:

- delete artifacts 5 seconds after the run reaches a terminal state

Terminal states include:

- success
- failure

The retention rule is intentionally short because:

- the user is expected to inspect progress in the terminal
- the goal is to avoid directory bloat
- the system should be largely self-cleaning

### Problem with a single cleanup hook

A simple in-process delayed delete is not sufficient.

Reasons:

- the user may close the terminal immediately after completion
- the Fabric process may crash before cleanup fires
- the machine may sleep or reboot
- only some artifacts may have been written when the run terminated

This means a single delayed delete mechanism is not reliable enough on its own.

### Agreed cleanup architecture: hybrid model

The agreed solution is a hybrid cleanup strategy.

It has two parts:

1. per-run delayed cleanup
2. startup janitor cleanup

Both are required.

### Part 1: Per-run delayed cleanup

For every pipeline run:

1. Fabric creates `.pipeline/<run-id>/`.
2. Fabric writes a run-state file into that run folder.
3. Fabric updates run state as the pipeline progresses.
4. When the run finishes, Fabric marks the run complete.
5. A best-effort cleanup task deletes that run folder after 5 seconds.

This gives fast cleanup in the normal successful path.

### Part 2: Startup janitor cleanup

On every new invocation of `fabric --pipeline ...`, Fabric must also scan the current directory’s `.pipeline/` folder and delete stale run directories that should already have expired.

This janitor handles cases where:

- the previous Fabric process was terminated early
- the terminal was closed before cleanup completed
- the detached cleanup task itself failed
- the machine rebooted or slept before cleanup executed

This means stale artifacts do not persist forever even if the previous cleanup path was interrupted.

### Required run metadata

Each run directory must contain a small machine-readable state file, such as:

```text
.pipeline/<run-id>/run.json
```

At minimum, this state should track:

- `run_id`
- `status`
- `pid`
- `started_at`
- `updated_at`
- `completed_at`
- `expires_at`
- optional stage summary

The status should distinguish at least:

- `running`
- `completed`
- `failed`

This metadata enables:

- safe janitor decisions
- stale run detection
- per-run cleanup without guessing

### Cleanup rules

The cleanup rules agreed in this session are:

- delete only `.pipeline/<run-id>/`, never the entire `.pipeline/` tree blindly
- remove `.pipeline/` itself only if it is empty after per-run cleanup
- cleanup must apply to both successful and failed runs
- cleanup must be best-effort and must not block final user output

### Why only deleting the entire `.pipeline/` folder is unsafe

Deleting a shared `.pipeline/` root is unsafe because:

- multiple pipeline runs may happen close together
- one run may still be active while another completes
- deleting the root could remove artifacts for an active run

Therefore, cleanup must be scoped to the current run directory only.

### User-experience constraints that remain agreed

- the user must not need `sudo`
- the user must not need to pass a special cleanup flag
- the user must not need to manually delete artifacts
- the user should still receive immediate output and terminal progress
- cleanup must not add visible wait time to a successful run

### Permission model

No `sudo` should be required in normal operation.

The cleanup works under normal user permissions because Fabric only creates and removes files inside the current working directory where the user already has write access.

If the current directory is not writable, artifact creation would already fail, so cleanup introduces no new privilege requirement.

### Operational safety requirements

The cleanup system must be:

- per-run isolated
- best-effort
- restart-safe
- resilient to terminal closure
- resilient to main-process failure
- resilient to cleanup-helper failure through janitor fallback

### Implementation implications carried forward

The implementation must now include:

- `.pipeline/<run-id>/` creation
- run metadata file creation and updates
- run completion timestamping
- delayed cleanup scheduling for the current run
- startup janitor scan on each pipeline invocation
- safe empty-root removal for `.pipeline/`

### Explicitly rejected alternatives

The following alternatives were discussed and are not the selected design:

- only in-process delayed cleanup with no janitor
- writing all artifacts directly into one flat `.pipeline/` with no run separation
- keeping artifacts indefinitely
- requiring the user to manage cleanup manually
- requiring elevated privileges or extra setup

## Current Fabric Behavior That Informed the Design

Current pattern mode already makes chaining possible because:

- stdin is read into the message buffer
- the pattern receives that content as `{{input}}`
- the final result is emitted to stdout

This existing behavior should be preserved conceptually for the final output of `--pipeline`, while pipeline mode adds stronger source validation and artifact handling.

## Confirmed Scalability Direction

The system must scale by:

- adding new pipeline definitions without editing the runner each time
- adding new executor types without redesigning the command surface
- keeping built-in and user-defined pipelines separate
- versioning the pipeline schema

The current agreed discovery model is:

- built-in pipelines: `data/pipelines/`
- user-defined pipelines: `~/.config/fabric/pipelines/`

User-defined pipelines may override built-in ones by name.

## Confirmed Requirements to Carry Into Implementation

- Fabric remains the only operator-facing command
- `--pipeline` is the multi-stage execution switch
- `--source` is required for file or directory sources
- `--scrape_url` remains the URL source entry point
- stdin remains a valid pipeline input source
- exactly one input source is allowed for `--pipeline`
- pipeline runs must support chaining through clean stdout behavior
- command stages must support arbitrary host-language executables

## Settled Clarification 4: Meaning of `-o` in Pipeline Mode

Decision:

- in `fabric --pipeline ...`, `-o` writes the final rendered output to the specified file
- the same final rendered output must still also be emitted to stdout
- `-o` must not suppress chaining behavior

### Why this was chosen

This preserves Fabric’s pipe-first usability and keeps pipeline mode composable in Unix pipelines.

The user can both:

- save the final rendered output to a file
- continue piping that same output into another Fabric pattern or another shell command

### Example that must remain valid

```bash
cat a.txt | fabric --pipeline tech-note -o final.md | fabric --pattern summarize
```

In the agreed design:

- `final.md` receives the final rendered note
- stdout also receives the same final rendered note
- the downstream Fabric command receives the note content, not a status message

### Explicitly rejected behaviors

The following behaviors were considered and are not the selected design:

- `-o` suppresses stdout entirely
- `-o` replaces stdout with a short confirmation message
- `-o` changes pipeline mode into a file-only output mode

## Settled Clarification 5: `-o` Must Not Control Artifact Rooting

Decision:

- `-o` controls only the final rendered note file
- `-o` does not control where `.pipeline/` artifacts are written
- `.pipeline/` remains rooted in the current working directory where `fabric` is invoked

### Why this was chosen

This keeps the model simple and unsurprising:

- final user-facing output is controlled by `-o`
- temporary pipeline artifacts follow the pipeline lifecycle rules
- artifact cleanup remains consistent and local to the invocation directory

This separation avoids overloading one flag with two different meanings.

### Agreed behavior

For example:

```bash
fabric --pipeline tech-note --source ./session.txt -o final.md
```

The expected result is:

- `final.md` contains the final rendered output
- stdout also prints the final rendered output
- `.pipeline/<run-id>/` is still created in the current working directory
- artifact cleanup still follows the same 5-second hybrid cleanup model

### Explicitly rejected behavior

The following behavior is not part of the design:

- using `-o` to relocate `.pipeline/`
- using `-o` to redefine artifact-rooting semantics
- making pipeline artifact placement depend on the final note output path

## Settled Clarification 6: Exactly One Final Output Stage in V1

Decision:

- each pipeline must declare exactly one official final output stage in V1
- the output of that stage is the only content sent to stdout for normal pipeline chaining
- other stage outputs remain artifacts only unless a future feature expands this behavior

### Why this was chosen

This keeps the stdout contract simple and reliable:

- downstream commands always receive one predictable payload
- stage ordering changes do not silently change what gets piped
- validation and chaining remain easy to reason about

This is especially important because pipeline stages may produce multiple meaningful intermediate artifacts, for example:

- normalized source
- semantic map JSON
- rendered markdown
- validation report

Only one of these should be treated as the canonical pipeable output in V1.

### Agreed V1 behavior

For a pipeline like:

1. intake
2. normalize
3. semantic_map
4. render_markdown
5. validate
6. publish

the pipeline definition must identify one final-output stage, for example `render_markdown`.

Then:

```bash
cat a.txt | fabric --pipeline tech-note | fabric --pattern summarize
```

must send only the final rendered note to the downstream command.

Intermediate artifacts such as:

- `semantic_map.json`
- `validation_report.md`
- `normalized_source.md`

must remain in `.pipeline/<run-id>/` only and must not be mixed into stdout.

### Explicitly not selected for V1

The following alternatives were considered and are not part of V1:

- multiple output-producing stages where the last one implicitly wins
- multiple candidate final outputs with ambiguous stdout behavior
- requiring the user to specify an output stage on every invocation

### Future extensibility requirement

Although V1 uses exactly one final output stage, the design should not make future expansion impossible.

If users later require it, the pipeline schema and runner may be extended to support a controlled version of multi-output behavior, for example:

- multiple declared output artifacts
- an explicit output-selection mode
- structured JSON output mode for advanced consumers

However, this future capability must be additive.

It must not weaken the V1 rule that:

- one pipeline run has one canonical stdout payload

## Settled Clarification 7: Validation Must Gate Final Stdout

Decision:

- if a pipeline includes a validation stage, validation must run before final stdout emission
- final stdout must be emitted only if validation passes
- if validation fails, Fabric must not emit the final pipeline payload to stdout
- validation failure must produce a non-zero exit code and a stage-specific error on stderr

### What ambiguity this resolves

In a multi-stage pipeline, there may be a stage that renders the final note and a later stage that validates it.

The question was whether Fabric should:

- print the rendered note immediately and validate afterward, or
- treat validation as a gate before anything is released to stdout

The selected answer is:

- validation is a gate

### Why this was chosen

This preserves the safety of pipe chaining.

If invalid output is printed before validation, downstream commands may consume broken or rejected content before Fabric reports failure.

That would make pipeline mode unreliable for:

- Unix chaining
- automation
- trust in final output

### Agreed behavior

For a pipeline such as:

1. intake
2. normalize
3. semantic_map
4. render_markdown
5. validate
6. publish

the intended V1 behavior is:

1. `render_markdown` generates the candidate final output
2. `validate` checks that candidate output
3. only if `validate` passes does Fabric emit the final output to stdout
4. if `validate` fails, Fabric emits no final payload to stdout

### Example that motivated this decision

```bash
cat a.txt | fabric --pipeline tech-note | fabric --pattern summarize
```

With the selected design:

- if validation passes, `summarize` receives the final note
- if validation fails, `summarize` should not receive a misleading final note

This keeps pipeline chaining trustworthy.

### Explicitly rejected behavior

The following behavior is not part of V1:

- print final stdout first and validate afterward
- allow invalid pipeline payloads to be consumed downstream before the run is marked failed
- make validation gating optional per pipeline in a way that weakens the default stdout contract

## Settled Clarification 8: Pipelines Without a Validate Stage

Decision:

- a pipeline may omit the `validate` stage in V1
- if no `validate` stage exists, Fabric must still allow the pipeline to run
- in that case, Fabric must emit a warning on stderr before releasing the final payload to stdout
- the final payload is emitted after the declared final output stage completes

### Why this was chosen

This preserves flexibility for lightweight or draft-oriented pipelines while still making the absence of validation visible to the operator.

It avoids two bad extremes:

- forcing every pipeline to carry a validation stage even when the workflow does not need one
- silently treating unvalidated output as if it were validated

### Agreed behavior

If a pipeline contains:

1. intake
2. normalize
3. render_markdown
4. publish

and does not contain `validate`, then:

- Fabric runs the pipeline normally
- Fabric emits a warning on stderr indicating that no validation stage exists
- Fabric emits the final payload to stdout after the declared final output stage completes

### Example

```bash
cat a.txt | fabric --pipeline fast-draft-note
```

Expected behavior:

- stderr contains something like: `warning: pipeline fast-draft-note has no validate stage`
- stdout still contains the final rendered output
- the run is not rejected purely because validation is absent

### Relationship to Clarification 7

The rules now are:

- if a `validate` stage exists, it must gate final stdout
- if no `validate` stage exists, Fabric must warn on stderr and then emit the final payload normally

This keeps the stdout contract consistent while remaining flexible.

### Explicitly rejected behavior

The following is not part of V1:

- rejecting every pipeline definition that lacks a validate stage
- silently omitting any indication that output was not validated

## Settled Clarification 9: Validation-Absence Warning Format

Decision:

- if a pipeline has no `validate` stage, the warning must be emitted as a plain informational stderr line
- it must not be modeled as a full pipeline stage event
- it must not appear as a fake stage in the stage progress UI

### Why this was chosen

The absence of validation is important operator information, but it is not itself a pipeline execution stage.

Treating it as a stage would overcomplicate the stage model and make the progress UI less truthful.

### Agreed behavior

For a pipeline such as:

```bash
cat a.txt | fabric --pipeline fast-draft-note
```

where `fast-draft-note` does not define a `validate` stage:

- stderr should include a plain warning line
- stdout should still contain the final payload
- stage UI should continue to reflect only the actual defined stages

### Example warning shape

An acceptable V1 form is something like:

```text
warning: pipeline fast-draft-note has no validate stage
```

The exact wording can evolve, but it must remain:

- informational
- stderr-only
- outside the formal stage progression

### Explicitly rejected behavior

The following is not part of V1:

- inventing a pseudo-stage named `validate-warning`
- showing missing-validation as a PASS/FAIL stage row
- mixing the warning into stdout

## Settled Clarification 10: Publish Failure After Validated Final Output

Decision:

- if the final output stage succeeds and validation passes, Fabric may still emit the final payload to stdout even if a later `publish` stage fails
- if `publish` fails, the overall pipeline run must still exit non-zero
- the publish failure must be reported on stderr as a stage-specific failure

### Why this was chosen

This preserves the usability of the validated final output while still treating downstream publish/bookkeeping failures as real failures.

It prevents an unnecessary loss of already validated content while maintaining accurate failure semantics for the overall pipeline run.

### Agreed behavior

For a pipeline such as:

1. intake
2. normalize
3. render_markdown
4. validate
5. publish

if:

- `render_markdown` succeeds
- `validate` passes
- `publish` fails

then:

- stdout may still contain the validated final payload
- stderr must contain the publish-stage failure
- the process exit code must be non-zero

### Example

```bash
cat a.txt | fabric --pipeline tech-note | fabric --pattern summarize
```

Under the selected design:

- the downstream `summarize` command may still receive the validated final note
- the top-level pipeline invocation is still considered failed because publish failed

This means chaining remains useful while operational failures remain visible and machine-detectable.

### Important interpretation

This decision assumes that `publish` is not the canonical content-generation stage.

In V1:

- the pipeline’s declared final output stage remains the source of stdout
- `publish` is treated as a downstream operational stage unless a future design explicitly changes that model

### Explicitly rejected behavior

The following is not part of V1:

- suppressing an already validated final payload purely because a later publish step failed
- silently treating publish failure as non-fatal

## Settled Clarification 11: Stage Names Are Conventions, Not Reserved Semantics

Decision:

- names such as `intake`, `normalize`, `render_markdown`, `validate`, and `publish` are examples and conventions only
- they are not reserved stage IDs
- they must not carry built-in behavior purely because of their names
- Fabric must derive behavior from explicit stage metadata and executor configuration, not from stage naming

### Why this was chosen

This keeps the pipeline system flexible and scalable.

Users must be able to define pipelines for many different workflows without being forced into one canonical stage vocabulary.

Examples:

- transcript note generation
- article synthesis
- code review
- report generation
- meeting summarization
- custom domain-specific workflows

All of these may need different stage names and different stage counts.

### Agreed behavior

The following should both be valid if they obey the execution contract:

```yaml
stages:
  - id: intake
    executor: builtin
  - id: semantic_map
    executor: fabric_pattern
  - id: render_markdown
    executor: fabric_pattern
    final_output: true
```

and:

```yaml
stages:
  - id: collect_diff
    executor: command
  - id: analyze_risks
    executor: fabric_pattern
  - id: format_report
    executor: fabric_pattern
    final_output: true
```

In both cases, Fabric should look at:

- executor type
- declared final output stage
- optional validation semantics
- command configuration

and not at the literal stage names themselves.

### Core rule carried forward

Users can define:

- any number of stages
- any stage names
- any number of external code steps

The pipeline definition only needs to obey the execution contract.

### Explicitly rejected behavior

The following is not part of V1:

- hard-coding semantics based on stage names alone
- requiring every pipeline to contain the same stage IDs
- forcing users into one canonical fixed sequence

## Settled Clarification 12: `publish` Is Optional

Decision:

- `publish` is optional in V1
- a pipeline is valid without any publish stage
- stdout behavior works normally even when no publish stage exists

### Why this was chosen

Not every pipeline needs a downstream delivery or bookkeeping step.

Many pipelines only need to:

- read input
- transform input
- optionally validate output
- emit the final result

Requiring `publish` everywhere would add unnecessary boilerplate.

### Agreed behavior

A pipeline such as:

1. read_source
2. chunk_text
3. summarize
4. render_markdown
5. validate

is valid without any `publish` step.

If the declared final output stage succeeds, and validation passes when present, stdout behaves normally.

### Relationship to earlier clarifications

The rules now are:

- exactly one final output stage in V1
- validation gates stdout if a validate stage exists
- missing validate is allowed but must warn on stderr
- publish is optional
- publish failure after validated final output still fails the run overall if publish is present

### Explicitly rejected behavior

The following is not part of V1:

- requiring every pipeline to define `publish`
- forcing user-defined pipelines to add no-op publish stages just to satisfy schema rules

## Settled Clarification 13: Pipelines May Be Entirely Command-Driven

Decision:

- a valid pipeline in V1 may contain zero `fabric_pattern` stages
- pipelines may be entirely command-driven if they obey the execution contract
- Fabric patterns are supported stage executors, not mandatory stage executors

### Why this was chosen

The pipeline system must be able to host and orchestrate existing workflows, not only AI-prompt-centric workflows.

Some useful pipelines may be composed entirely of external tools or scripts, for example:

- Bash preprocessing
- Python transformation
- Node validation
- Java formatting
- compiled utility stages

Requiring at least one `fabric_pattern` stage would make the runner less generally useful and would block legitimate workflows.

### Agreed behavior

A pipeline such as:

1. read_input via Bash
2. transform via Python
3. validate_json via Node
4. format_output via Java

is valid in V1 if:

- it declares exactly one final output stage
- it obeys stdout/stderr rules
- it obeys artifact and exit-code rules
- it obeys source-input and cleanup rules

### Relationship to Fabric’s role

Even if a specific pipeline contains no `fabric_pattern` stages, it is still a Fabric pipeline because:

- Fabric loads the pipeline definition
- Fabric orchestrates stage execution
- Fabric handles source rules
- Fabric handles progress rendering
- Fabric handles artifact management
- Fabric handles cleanup
- Fabric enforces the execution contract

So Fabric remains the operator-facing runner even when a given pipeline is composed entirely of external executables.

### Explicitly rejected behavior

The following is not part of V1:

- requiring every pipeline to contain at least one LLM-driven stage
- treating command-only pipelines as invalid by definition

## Settled Clarification 14: V1 Pipelines Are Strictly Linear

Decision:

- V1 pipeline definitions are strictly linear
- Fabric does not need native branching, merging, or parallel stage graphs in V1
- stages execute in declared order, one after another

### Why this was chosen

This keeps the runner simpler and more reliable in V1.

It reduces complexity in:

- stage progress rendering
- stdout contract
- failure semantics
- artifact tracking
- cleanup behavior
- pipeline definition validation

### Agreed behavior

A V1 pipeline definition should declare one ordered list of stages.

Examples of V1-valid shape:

1. stage A
2. stage B
3. stage C

and:

1. read_input
2. transform
3. validate
4. final_output

### Branching can still happen inside a stage

The absence of native branching in the Fabric runner does not prevent users from implementing branching logic inside their own command stages.

For example, a command stage may:

- inspect the current input
- choose one of several internal code paths
- invoke different scripts or tools
- merge internal results
- emit one normalized output for the next declared Fabric stage

This means a pipeline can remain linear from Fabric’s point of view while still allowing sophisticated routing logic inside user code.

### Example interpretation

Pipeline definition:

1. `read_input`
2. `decide_path`
3. `continue_processing`
4. `render_output`

Here, `decide_path` may internally branch based on content and still return one output artifact or payload to stage 3.

So the Fabric-level rule is:

- linear orchestration

while the user-level freedom is:

- arbitrary branching inside a stage implementation

### Explicitly rejected behavior

The following is not part of V1:

- native branching stage graphs
- native fan-out/fan-in
- native parallel stage execution
- merge semantics between multiple Fabric-level branches

## Settled Clarification 15: Command Stages Are Fully Trusted and Unsandboxed

Decision:

- command stages may write or modify files outside `.pipeline/<run-id>/`
- command stages may also delete files or directories that the current user has permission to delete
- this is required for real workflows such as writing to Obsidian, updating user-managed destinations, or running fully custom automation logic
- Fabric will not add confirmation guardrails for destructive filesystem operations in V1
- the pipeline author is responsible for what their command stages do

### Why this was chosen

The pipeline system must remain useful for practical automation.

Legitimate workflows may need to:

- write notes into an Obsidian vault
- update indexes or dashboards
- export content into project directories
- sync outputs into other local destinations

Blocking all writes or deletes outside `.pipeline/<run-id>/` would make the system too weak for these real use cases.

Because user-defined pipelines are authored intentionally, the selected trust model is:

- Fabric orchestrates
- the user owns the command logic
- the user accepts responsibility for destructive behavior in their own pipeline definitions

### Agreed behavior

Allowed by default:

- creating files outside `.pipeline/<run-id>/`
- updating files outside `.pipeline/<run-id>/`
- writing outputs to user-managed destinations
- deleting files outside `.pipeline/<run-id>/`
- deleting directories outside `.pipeline/<run-id>/`
- running arbitrary host commands with the current user's permissions

### Example

Acceptable default behavior:

- a pipeline writes `W07 - Program Derived Addresses.md` into an Obsidian vault
- a pipeline updates a resource index file
- a pipeline writes a final report into the current project
- a pipeline removes or replaces files in a destination folder if the user intentionally wrote that behavior into the pipeline

### Default safety model

The selected safety model is:

- normal writes: allowed
- destructive deletes: allowed
- command stages are trusted, not sandboxed

### Explicitly rejected behavior

The following is not part of V1:

- forbidding all writes outside `.pipeline/<run-id>/`
- adding interactive confirmation prompts for command-stage deletes
- treating Fabric as a sandbox for user-authored pipeline commands

## Settled Clarification 16: Stage-to-Stage Handoff Uses a Hybrid Contract

Decision:

- stage-to-stage flow in V1 uses a hybrid contract
- each stage may produce one primary output payload for the next stage
- each stage may also produce optional artifact files in `.pipeline/<run-id>/`
- Fabric tracks the primary payload for orchestration purposes while preserving artifacts for inspection

### Why this was chosen

This combines the strengths of both pure stdout chaining and pure artifact-driven execution.

It keeps the system:

- composable
- inspectable
- suitable for command stages
- suitable for Fabric-pattern stages
- suitable for structured workflows that need files as well as next-stage inputs

### Agreed behavior

Every stage may expose:

- one primary output payload
- zero or more side artifacts

The primary output payload is what Fabric passes forward to the next stage by default.

Artifacts remain available in `.pipeline/<run-id>/` for:

- debugging
- validation
- inspection
- downstream stage file reads where explicitly configured

### Example

A pipeline may behave like this:

1. `normalize`
   - primary payload: normalized markdown text
   - artifact: `normalized_source.md`

2. `semantic_map`
   - primary payload: semantic map JSON or normalized structured content
   - artifact: `semantic_map.json`

3. `render_markdown`
   - primary payload: final markdown note
   - artifact: `final_notes.md`

Fabric uses the primary payload to move from stage to stage while keeping the artifacts on disk.

### Why pure stdout-only was not selected

Pure stdout-only handoff is too fragile for complex workflows because:

- logs and content can be mixed if stage discipline is weak
- structured multi-file workflows become awkward
- validation and debugging lose inspectable intermediate state

### Why pure file-only handoff was not selected

Pure file-only handoff is too rigid because:

- it is verbose for simple flows
- it weakens the natural chaining model
- it overemphasizes artifacts for cases where the primary payload is enough

### Core rule carried forward

The runner must distinguish between:

- the primary stage payload used for orchestration
- side artifacts used for persistence and inspection

This distinction is central to the V1 execution contract.

## Settled Clarification 17: Primary Output Source Must Be Declared Explicitly

Decision:

- when a stage can produce both stdout and artifact files, the stage definition must explicitly declare which output is the primary handoff for the next stage
- Fabric must not guess implicitly between stdout and artifact outputs in this situation

### Why this was chosen

Guessing would create ambiguity and make pipelines brittle.

Command stages and Fabric-pattern stages may produce:

- stdout content
- one or more artifact files
- both at the same time

Without an explicit declaration, the runner could incorrectly pass the wrong output to the next stage.

### Agreed behavior

Stage definitions should explicitly identify their primary output source, for example by indicating:

- primary output comes from stdout
- primary output comes from a specific artifact

The exact field syntax can be finalized during command-contract and schema work, but the behavioral rule is now fixed:

- primary output source must be declared

### Example

If a command stage:

- prints a short summary to stdout
- writes `semantic_map.json` as an artifact

then the pipeline definition must explicitly say which one the next stage should consume.

That prevents the runner from assuming incorrectly.

### Explicitly rejected behavior

The following is not part of V1:

- assuming stdout is always primary
- assuming an artifact is always primary
- letting the runner infer primary handoff source by guesswork

## Settled Clarification 18: Stage Input Defaults to Previous Primary Payload but May Read Earlier Named Artifacts

Decision:

- the default stage input in V1 is the immediately previous stage’s primary output payload
- a stage may also explicitly read named artifacts produced by earlier stages
- this is allowed only when declared explicitly in the pipeline definition

### Why this was chosen

This preserves the simplicity of a linear default flow while still allowing richer workflows when a later stage needs a specific earlier artifact.

It avoids two extremes:

- overly rigid previous-stage-only chaining
- unstructured free-for-all access without declaration

### Agreed behavior

Default behavior:

- stage N consumes stage N-1 primary output

Explicit advanced behavior:

- stage N may additionally or alternatively consume a specific named artifact from an earlier stage if the definition declares it

### Example

Default case:

1. `normalize`
2. `semantic_map`
3. `render_markdown`

Here:

- `semantic_map` consumes `normalize` primary output
- `render_markdown` consumes `semantic_map` primary output

Explicit artifact reference case:

1. `normalize`
   - artifact: `normalized_source.md`
2. `semantic_map`
   - artifact: `semantic_map.json`
3. `render_markdown`
   - primary input: previous stage primary payload
   - extra input: `normalize.normalized_source.md`

This lets a later stage use the normal pipeline flow while still referring back to earlier artifacts when needed.

### Core rule carried forward

The runner should assume previous-stage primary payload unless the pipeline definition explicitly names another artifact source.

### Explicitly rejected behavior

The following is not part of V1:

- forcing every stage to declare all inputs manually even for simple linear flows
- allowing undeclared arbitrary reads of prior artifacts by guesswork

## Settled Clarification 19: Built-In and User-Defined Pipelines Share the Same Schema

Decision:

- built-in and user-defined pipelines use the same schema in V1
- built-in and user-defined pipelines have the same capability model in V1
- Fabric must not create a separate privileged schema for built-in pipelines

### Why this was chosen

This keeps the product easier to understand, easier to document, and easier to extend.

It avoids a two-class pipeline system where:

- shipped pipelines behave one way
- user pipelines behave another way

That kind of split would make learning, debugging, and sharing pipelines harder.

### Agreed behavior

The same schema rules must apply to:

- built-in pipeline definitions stored in `data/pipelines/`
- user-defined pipeline definitions stored in `~/.config/fabric/pipelines/`

The same execution contract must apply to both.

### Benefits carried forward

- one mental model
- one validator
- one documentation surface
- easier override behavior
- easier future extensibility

### Explicitly rejected behavior

The following is not part of V1:

- hidden built-in-only pipeline capabilities
- a second privileged schema for shipped pipelines
- feature gaps between built-in and user-authored pipeline definitions

## Settled Clarification 20: User-Defined Pipelines Override Built-Ins With Informational Visibility

Decision:

- if a user-defined pipeline has the same name as a built-in pipeline, the user-defined pipeline wins
- Fabric must emit a small informational stderr line indicating that the built-in definition was overridden

### Why this was chosen

This preserves predictable override behavior while keeping the operator informed about which definition actually ran.

It avoids two bad outcomes:

- silent overrides that make debugging confusing
- noisy or blocking override mechanics for a normal use case

### Agreed behavior

If:

- `data/pipelines/tech-note.yaml` exists
- and `~/.config/fabric/pipelines/tech-note.yaml` also exists

then:

- Fabric runs the user-defined `tech-note`
- stderr includes an informational line stating that the built-in pipeline was overridden

### Example informational behavior

An acceptable V1 form is something like:

```text
info: using user-defined pipeline 'tech-note' instead of built-in definition
```

The exact wording can evolve, but the behavior should remain:

- informational
- stderr-only
- non-blocking

### Explicitly rejected behavior

The following is not part of V1:

- preferring built-ins over user-defined pipelines with the same name
- silently overriding without any visibility
- prompting the user interactively just to resolve a normal override

## Settled Clarification 21: Environment Variables Come From the Process Environment

Decision:

- pipeline definitions may reference environment variables
- Fabric expands those variables from the current process environment
- Fabric does not parse shell startup files such as `.zshrc` or `.bashrc` directly

### Why this was chosen

This matches normal Unix process behavior and keeps Fabric simple.

The shell is responsible for building the environment. Fabric simply receives that environment and expands variables from it.

### Agreed sources of environment values

Users may provide environment variables through any normal process-environment path, including:

- exported variables in the current shell
- inline one-off variable assignment before the command
- shell startup files such as `.zshrc` or `.bashrc`, once those files have been loaded by the shell
- scripts, CI jobs, or remote runners that set process environment variables

### Examples

Export in current shell:

```bash
export OBSIDIAN_VAULT="$HOME/Documents/MyVault"
fabric --pipeline tech-note --source ./session.txt
```

Inline one-off assignment:

```bash
OBSIDIAN_VAULT="$HOME/Documents/MyVault" fabric --pipeline tech-note --source ./session.txt
```

Shell startup file, loaded by the shell:

```bash
export OBSIDIAN_VAULT="$HOME/Documents/MyVault"
source ~/.zshrc
fabric --pipeline tech-note --source ./session.txt
```

### Important distinction

Fabric should not:

- read `.zshrc` directly
- read `.bashrc` directly
- attempt to interpret shell configuration files itself

Instead:

- the shell loads those files
- exported variables become part of the process environment
- Fabric expands from that process environment

### Relationship to `PATH`

`PATH` remains a normal environment variable and is used in the standard way for executable lookup.

Custom variables such as `OBSIDIAN_VAULT` follow the same general environment mechanism.

### Explicitly rejected behavior

The following is not part of V1:

- Fabric directly parsing user shell startup files
- using a separate environment-variable resolution model for pipelines only

## Settled Clarification 22: Fabric Performs Mandatory Preflight Validation and Exposes It as a User Command

Decision:

- every pipeline run in V1 must go through Fabric preflight validation before stage 1 starts
- Fabric must also expose preflight validation as a user-invokable command
- V1 will support both:
  - direct pipeline-file validation
  - installed-pipeline validation by name

### Why this was chosen

Pipelines are definition-driven execution units. They need the same kind of up-front validation discipline that users expect from build tools, schema validators, and compilers.

This keeps the system safer and more debuggable by catching definition problems before any stage actually runs.

### Validation layers now agreed

The design now distinguishes three layers:

1. schema/syntax validation
2. semantic preflight validation
3. runtime contract checks

#### 1. Schema/syntax validation

This validates the structure of the pipeline definition itself, for example:

- valid YAML or other selected file format
- required top-level fields exist
- field types are correct
- stage definitions follow the schema

#### 2. Semantic preflight validation

This validates the logical correctness of the definition before execution, for example:

- exactly one final output stage exists
- stage IDs are unique
- executor types are valid
- required environment variables are present
- declared input references are resolvable
- primary output declarations are valid
- installed source mode usage is valid

#### 3. Runtime contract checks

This happens during execution, for example:

- stage exit codes
- required primary output is actually produced
- declared artifacts actually exist
- stage failure propagation
- cleanup metadata updates

### Important distinction

Fabric’s preflight validation is not the same thing as a user-authored `validate` stage.

The rules are now:

- Fabric preflight validation is mandatory for all pipelines
- runtime contract checks are mandatory for all stages
- user-authored `validate` stages remain optional and are for workflow-specific validation, not for schema validation

### Agreed user-facing validation commands

V1 should support both of the following styles:

Direct file validation:

```bash
fabric --validate-pipeline ./my-pipeline.yaml
```

Installed pipeline validation by name:

```bash
fabric --pipeline tech-note --validate-only
```

### Why both forms are needed

Direct file validation is useful when:

- authoring a new pipeline definition
- checking a file before installing it
- validating a file in CI or scripts

Installed pipeline validation is useful when:

- verifying the resolved pipeline that Fabric will actually run
- checking behavior after override resolution
- validating built-in or user-defined pipelines by name

### Core rule carried forward

Fabric must treat pipeline validation as a first-class part of the product, not as an implementation detail hidden inside runtime failures.

## Settled Clarification 23: Preflight Validation Is Side-Effect Free and Never Executes AI Calls

Decision:

- `--validate-pipeline` validates pipeline definition correctness only
- preflight validation must be side-effect free
- preflight validation must not execute live AI calls
- preflight validation must not execute user scripts or command stages
- actual `--pipeline` execution is what runs AI stages and command stages for real
- a future `--dry-run` mode may describe execution without actually performing it

### Why this was chosen

Validation must remain:

- safe
- fast
- cheap
- deterministic
- usable in scripting and CI

If validation were allowed to execute AI stages, then it would stop being a true validation command and become a partial pipeline execution mode. That would introduce:

- token cost
- provider/network dependency
- side effects
- ambiguous operator expectations

The validation command must therefore validate the contract of the pipeline, not the truth or quality of AI-generated output.

### Agreed meaning of `--validate-pipeline`

The command:

```bash
fabric --validate-pipeline ./my-pipeline.yaml
```

checks definition correctness only.

It may perform checks such as:

- file/schema parsing
- semantic pipeline validation
- environment-variable resolution
- stage-reference validation
- final-output-stage validation
- executor-type validation
- pattern existence checks
- command executable resolution where possible

It must not:

- make live model calls
- execute stage commands
- mutate the filesystem beyond any purely internal read-only resolution work
- perform publish actions
- write to external systems

### Agreed meaning of actual `--pipeline`

The command:

```bash
fabric --pipeline tech-note --source ./session.txt
```

is the real execution path.

This is the mode that:

- runs AI stages
- runs command stages
- creates runtime artifacts
- performs stage-by-stage execution
- emits progress to stderr
- emits final validated output to stdout according to earlier decisions

### Agreed future shape of `--dry-run`

`--dry-run` is not required for V1, but the conceptual shape is now clear.

A future dry-run mode may show:

- the resolved pipeline definition
- resolved environment variables
- which Fabric patterns would run
- which commands would run
- where artifacts would be written
- what the final output stage would be

Even in that future mode, the rule should remain:

- no real AI calls
- no command execution
- no external side effects

### Important distinction carried forward

Preflight validation can prove that a pipeline is well-formed and runnable.

It cannot prove that:

- an AI stage will return valid Mermaid
- an AI stage will produce beautiful Markdown
- a model will obey every prompt constraint
- a runtime command will produce semantically correct output

Those concerns belong to:

- runtime contract checks
- optional user-authored `validate` stages
- future richer post-stage validators if needed

### Canonical model now agreed

The execution model is now:

1. `--validate-pipeline`
   - checks definition correctness only
   - no AI calls
   - no script execution
2. actual `--pipeline`
   - runs AI stages and command stages for real
3. optional future `--dry-run`
   - shows what would happen
   - still no AI calls
   - still no script execution

## Settled Clarification 24: User-Facing Pipeline Files Are YAML, Validated Through a Versioned Structured Schema

Decision:

- V1 pipeline definitions will be authored in YAML
- Fabric will parse YAML into a structured internal data model
- Fabric will validate that parsed representation against a versioned schema/rule set
- users do not need to manually convert YAML into JSON
- validation rigor should come from the schema, while authoring ergonomics should come from YAML

### Why this was chosen

The design needed to balance two things:

- a good human authoring experience
- a rigorous, evolvable validation model

YAML is the better authoring format for hand-written pipeline definitions because:

- it is easier to read than JSON
- it is easier to write and review by hand
- nested structures like stage definitions are more ergonomic
- it matches how many infrastructure and workflow tools are authored in practice

At the same time, free-form YAML without a strict schema would become fragile.

That is why the validation model must remain schema-driven and versioned.

### Agreed V1 authoring experience

Users should write pipeline definitions in YAML.

Example:

```yaml
version: 1
name: tech-note
stages:
  - id: normalize
    executor: command
    command:
      program: python3
      args: ["normalize.py"]
    primary_output:
      from: artifact
      path: normalized.md

  - id: render
    executor: fabric_pattern
    pattern: tech_note
    final_output: true
    primary_output:
      from: stdout
```

This is the file format the operator and pipeline author should see and maintain.

### Agreed internal validation model

Fabric should not treat YAML as loosely interpreted config text.

Instead, the pipeline loader should:

1. read the YAML file
2. parse it into a structured in-memory representation
3. validate that parsed structure against a versioned schema
4. then run semantic preflight checks on top of the schema pass

This is effectively the same design pattern as:

- YAML as authoring syntax
- JSON-like structured data as the internal representation
- schema-based validation over the parsed representation

### Important clarification

The user does **not** need to manually convert YAML into JSON.

Fabric is responsible for:

- parsing the YAML
- turning it into structured data
- validating that structured data

So the contract is:

- YAML for humans
- schema-validated structured representation for Fabric

### Why JSON-only was not chosen

JSON-only would be easier for purely machine-oriented validation, but it is worse for user authoring.

Pipeline definitions are expected to be hand-authored and reviewed. For that use case, JSON adds unnecessary friction.

### Why YAML-only without a schema was not chosen

YAML-only without a proper schema would weaken:

- preflight validation quality
- forward compatibility
- error clarity
- implementation safety

The system needs a formal contract, not just a parseable text format.

### Versioning rule carried forward

The schema must be versioned from the start.

That means the pipeline definition format should contain an explicit version field, for example:

```yaml
version: 1
```

This allows the pipeline language to evolve later without silently breaking older definitions.

### Canonical model now agreed

The file-format and validation model is now:

- user-authored pipeline files are YAML
- Fabric parses YAML into a structured internal representation
- Fabric validates that representation against a versioned schema/rule set
- semantic preflight validation runs after schema validation
- users never need to manually convert YAML to JSON

## Settled Clarification 25: Pipeline Discovery Uses a Dedicated `--listpipelines` Flag

Decision:

- pipeline discovery in V1 will use:

```bash
fabric --listpipelines
```

- the product will not use a generic `--list --pipelines` shape

### Why this was chosen

This decision was verified against Fabric’s existing CLI conventions before being recorded.

Fabric already exposes listing operations as dedicated flags rather than through a shared `--list` namespace.

Existing examples include:

- `--listpatterns`
- `--listmodels`
- `--listcontexts`
- `--listsessions`
- `--listextensions`
- `--liststrategies`
- `--listvendors`

So the pipeline discovery command should follow the same pattern for consistency and predictability.

### Why `--list --pipelines` was not chosen

That shape would introduce a different command style from the rest of the Fabric CLI.

Even though it can sound clean in isolation, it would create an unnecessary second listing idiom inside the same product.

### Agreed operator experience

The V1 discovery command for pipelines should therefore be:

```bash
fabric --listpipelines
```

This command should list:

- built-in pipelines
- user-defined pipelines
- whether a user-defined pipeline overrides a built-in pipeline of the same name

### Output intent

The exact formatting is still implementation detail, but the command should make it easy for an operator to answer:

- what pipelines are available
- which ones are built in
- which ones are user-defined
- which built-ins are currently overridden by user-defined versions

## Settled Clarification 26: User-Defined Pipelines Live in `~/.config/fabric/pipelines/`

Decision:

- built-in pipelines in V1 should live in the Fabric repo under a built-in path such as:

```text
data/pipelines/
```

- user-defined pipelines in V1 should live in:

```text
~/.config/fabric/pipelines/
```

- user-defined pipelines override built-in pipelines by name, following the same precedence rule already agreed earlier

### Why this was chosen

This decision was verified against Fabric’s current configuration conventions before being recorded.

Fabric already uses the `~/.config/fabric/...` model for user-defined extensibility points such as:

- patterns
- strategies
- related user-managed configuration

Examples already present in the current product and documentation include:

- `~/.config/fabric/patterns/`
- `~/.config/fabric/strategies/`

So `~/.config/fabric/pipelines/` is the most consistent and least surprising V1 location for user-authored pipeline definitions.

### Why project-local pipeline directories were not chosen for V1

A project-local-only location would initially sound convenient, but it introduces avoidable complexity early:

- ambiguous discovery rules
- more precedence questions
- weaker consistency with Fabric’s current configuration model
- harder global listing behavior

Project-local pipeline discovery may still be added later, but it is not needed for V1.

### Agreed built-in vs user-defined layout

The canonical location model now becomes:

Built-ins:

```text
data/pipelines/
```

User-defined:

```text
~/.config/fabric/pipelines/
```

### Agreed precedence rule carried forward

If a user-defined pipeline and a built-in pipeline share the same name:

- Fabric uses the user-defined version
- Fabric emits a small informational stderr line noting the override

This keeps behavior consistent with the custom override model already chosen elsewhere in the design.

### Why this path matters for operator UX

This gives users a stable and intuitive place to:

- author custom pipelines
- install their own pipeline definitions
- inspect what Fabric will load
- manage overrides cleanly without editing the Fabric repo itself

## Settled Clarification 27: V1 Uses `.yaml` as the Only Pipeline Definition File Extension

Decision:

- V1 pipeline definition files will use the `.yaml` extension only
- V1 will not support mixed `.yaml` and `.yml` discovery
- custom multi-part extensions such as `.fabric-pipeline.yaml` are also not required in V1

### Canonical examples

Built-in pipeline files:

```text
data/pipelines/tech-note.yaml
data/pipelines/zoom-tech-note.yaml
```

User-defined pipeline files:

```text
~/.config/fabric/pipelines/tech-note.yaml
~/.config/fabric/pipelines/obsidian-sync.yaml
```

### Why this was chosen

The pipeline language already has enough moving parts:

- names
- stage IDs
- override precedence
- schema versions
- built-in vs user-defined discovery

Allowing both `.yaml` and `.yml` would add avoidable file-discovery ambiguity in V1.

### Example of the ambiguity that V1 is avoiding

If both extensions were allowed, users could end up with:

```text
~/.config/fabric/pipelines/tech-note.yaml
~/.config/fabric/pipelines/tech-note.yml
```

That immediately creates extra questions:

- which file should win
- whether this is an error
- whether both should be loaded
- whether the `name:` field or the filename is authoritative

All of those questions are unnecessary for V1.

### Why `.fabric-pipeline.yaml` was not chosen

A custom extension such as:

```text
tech-note.fabric-pipeline.yaml
```

would be more self-describing, but it adds noise without solving a real V1 problem.

The parent directory already makes file purpose clear:

```text
~/.config/fabric/pipelines/
```

So the simpler extension is the better choice.

### Canonical rule now agreed

V1 file discovery for pipelines should therefore be:

- YAML authoring format
- `.yaml` extension only
- one canonical discovery rule
- no mixed-extension resolution logic required

## Settled Clarification 28: Pipeline Files Must Declare `name:`, and It Must Match the Filename Stem

Decision:

- every V1 pipeline definition must declare a top-level `name:` field
- the `.yaml` filename stem remains the lookup key used by Fabric
- the internal `name:` field must exactly match the filename stem
- a mismatch between filename and `name:` is a preflight validation error
- discovery, lookup, overrides, and validation still key off the filename stem, but the in-file name is now mandatory documentation and validation data

### Canonical examples

Built-in:

```text
data/pipelines/tech-note.yaml
data/pipelines/zoom-tech-note.yaml
```

User-defined:

```text
~/.config/fabric/pipelines/tech-note.yaml
~/.config/fabric/pipelines/obsidian-sync.yaml
```

In these cases, the pipeline identifiers are still:

- `tech-note`
- `zoom-tech-note`
- `obsidian-sync`

### Why this was chosen

The updated goal is to keep V1 explicit without making identity ambiguous.

The user asked to make `name:` mandatory in the same spirit as mandatory stage `id`s.

That is reasonable, but V1 still should not allow two competing identities.

So the design now uses both:

- filename stem as the lookup/discovery key
- `name:` as mandatory self-description and validation data

while preventing ambiguity by requiring an exact match.

### Example of the ambiguity being avoided

If Fabric allowed both values to differ, a user could create:

```text
~/.config/fabric/pipelines/my-tech.yaml
```

with:

```yaml
version: 1
name: tech-note
```

Now the system must decide:

- is the pipeline called `my-tech`
- or `tech-note`
- what should `fabric --pipeline ...` look up
- what should `--listpipelines` display

V1 still does not need that ambiguity.

So under the new rule, this definition is simply invalid.

### Canonical valid example

```text
~/.config/fabric/pipelines/tech-note.yaml
```

with:

```yaml
version: 1
name: tech-note
stages:
  - id: normalize
    executor: command
    ...
```

### What the file should still contain

The file still contains the schema-driven pipeline definition, for example:

```yaml
version: 1
name: tech-note
stages:
  - id: normalize
    executor: command
    ...
```

The file now does include a required explicit name field, but that field is constrained to match the filename stem exactly.

### Canonical V1 rule now agreed

Pipeline identity in V1 is:

- the `.yaml` filename stem
- reinforced by a mandatory matching `name:` field
- consistent across built-in and user-defined pipelines
- the key used for lookup
- the key used for override behavior
- the key shown in pipeline discovery output
- any mismatch is a validation error, not a precedence question

## Settled Clarification 29: Every Stage Must Declare an Explicit `id`

Decision:

- every stage in a V1 pipeline definition must declare an explicit `id`
- Fabric will not auto-generate stage identifiers in V1
- stage IDs are required even if no later stage explicitly references them

### Canonical example

```yaml
version: 1
stages:
  - id: normalize
    executor: command
    command:
      program: python3
      args: ["normalize.py"]

  - id: render
    executor: fabric_pattern
    pattern: tech_note
    final_output: true
    primary_output:
      from: stdout
```

### Why this was chosen

The pipeline system already depends on stable stage identity for several agreed features:

- clearer runtime diagnostics
- explicit earlier-artifact references
- stage-by-stage UI
- preflight validation messaging
- future partial execution features

If Fabric auto-generated stage IDs from order, then:

- references would become fragile after edits
- diagnostics would be less readable
- stage identity would become implicit rather than explicit

That is not a good V1 tradeoff.

### Example of ambiguity being avoided

If stage IDs were optional, a user could write:

```yaml
stages:
  - executor: command
  - executor: fabric_pattern
```

Now Fabric would need to invent internal names such as:

- `stage-1`
- `stage-2`

That creates avoidable problems:

- what should `--listpipelines` or debug output show
- how should later stages refer to earlier artifacts
- what happens if a new stage is inserted in the middle

Mandatory explicit IDs avoid that entire class of drift.

### Canonical rule now agreed

For V1:

- every stage has an explicit `id`
- stage IDs must be unique within a pipeline
- stage IDs are part of the preflight validation contract
- Fabric uses stage IDs in diagnostics, references, and runtime reporting

## Settled Clarification 30: Every Stage Declares an Explicit Top-Level `executor:` Field

Decision:

- every stage in V1 must declare an explicit top-level `executor:` field
- Fabric will not infer executor type from surrounding fields
- V1 uses the simple flat shape:
  - `executor: command`
  - `executor: fabric_pattern`
  - `executor: builtin`

### Canonical examples

Command stage:

```yaml
- id: normalize
  executor: command
  command:
    program: python3
    args: ["normalize.py"]
  primary_output:
    from: artifact
    path: normalized.md
```

Fabric-pattern stage:

```yaml
- id: render
  executor: fabric_pattern
  pattern: tech_note
  final_output: true
  primary_output:
    from: stdout
```

Builtin stage:

```yaml
- id: internal_check
  executor: builtin
  builtin: some_future_builtin_name
```

### Why this was chosen

The pipeline schema is already a real definition language with:

- mandatory pipeline `name:`
- mandatory stage `id`
- final output semantics
- primary output declarations
- artifact references

That means V1 should optimize for explicitness and strong validation, not inference.

### Why inference was rejected

If Fabric inferred executor type from keys like `pattern:` or `command:`, then definitions like the following would create unnecessary ambiguity:

```yaml
- id: something
  pattern: tech_note
  command:
    program: python3
```

Now the system must guess:

- is this a command stage
- is this a pattern stage
- should both keys be allowed
- is this simply invalid

An explicit `executor:` field turns that into a clean validation problem instead of an inference problem.

### Why nested executor objects were not chosen

A more structured shape such as:

```yaml
executor:
  type: fabric_pattern
```

is possible, but it adds verbosity without enough V1 benefit.

The simpler flat field is the better tradeoff for now.

### Canonical V1 rule now agreed

For every stage:

- `executor:` is required
- executor type is explicit
- Fabric validates executor-specific required fields based on that declared type
- Fabric does not guess executor type from other keys

## Settled Clarification 31: Command Stages Use Structured `program` and `args`, Not Raw Shell Strings

Decision:

- V1 command stages declare executables using a structured `program` plus `args` form
- V1 will not use a single raw shell command string as the canonical command shape
- V1 will not rely on implicit shell parsing as part of the core command-stage contract

### Canonical example

```yaml
- id: normalize
  executor: command
  command:
    program: python3
    args: ["normalize.py", "--mode", "fast"]
```

Another example:

```yaml
- id: compile_report
  executor: command
  command:
    program: java
    args: ["-jar", "report-builder.jar", "--input", "notes.json"]
```

### Why this was chosen

This keeps the command executor:

- language-agnostic
- explicit
- easier to validate
- safer to execute

It also avoids making shell parsing rules part of the pipeline language itself.

### Why raw shell strings were rejected for V1

A shape such as:

```yaml
command: "python3 normalize.py --mode fast"
```

looks compact, but it creates immediate ambiguity around:

- quoting
- escaping
- environment expansion behavior
- shell interpretation
- platform differences

That is not a good default contract for V1.

### Why “support both” was not chosen for V1

Allowing both:

- structured command objects
- raw shell strings

would make the schema and runtime behavior more complex immediately.

The structured form is enough for V1 and can still represent commands across Python, Bash, Java, Node, and other executable runtimes.

### Canonical V1 rule now agreed

For command stages:

- `command.program` is required
- `command.args` is the structured argument list
- Fabric executes the declared program with the declared argument vector
- Fabric does not treat the command as a shell string by default

## Settled Clarification 32: Command Stages May Optionally Declare `cwd`

Decision:

- command stages may optionally declare `cwd`
- `cwd` is not mandatory
- if `cwd` is omitted, the command runs with Fabric’s current working directory as its working directory
- relative pipeline-owned references should still be resolved from the pipeline file location when appropriate

### Why this was chosen

The ambiguity here was between two different concerns:

1. how Fabric locates files referenced by the pipeline definition
2. what directory a spawned command actually runs inside

Those are not the same thing and should not be conflated.

### Agreed default behavior

If a command stage does not declare `cwd`, then the process runs from the same directory where the user invoked `fabric`.

This keeps the default runtime behavior aligned with earlier decisions about:

- `.pipeline/` living in the current working directory
- operator-visible side effects being rooted in the current working directory
- output behavior feeling local to where the user launched the command

### Agreed resolution behavior for pipeline-local references

Pipeline definitions may still reference helper assets or scripts relative to the pipeline definition itself.

So the design now distinguishes:

- pipeline-definition-relative resolution for pipeline-owned helper references
- invocation-directory runtime behavior for default command execution

This allows a pipeline file in:

```text
~/.config/fabric/pipelines/tech-note.yaml
```

to reference helper files near the pipeline definition, while still letting the command execute from the user’s current working directory by default.

### Example where `cwd` is not needed

If Fabric can resolve a helper script from the pipeline definition area, then a stage may not need an explicit `cwd` at all.

Example:

```yaml
- id: normalize
  executor: command
  command:
    program: python3
    args: ["helpers/normalize.py"]
```

In that case:

- the helper path may be resolved relative to the pipeline definition area
- the command still executes from the operator’s current working directory by default

### Example where `cwd` is needed

Some commands are expected to run from a specific project directory because they depend on local relative files such as:

- `./templates`
- `./package.json`
- `./Makefile`
- `./config.yaml`

For those cases, the stage may explicitly declare:

```yaml
- id: build_docs
  executor: command
  command:
    program: python3
    args: ["scripts/build_docs.py"]
    cwd: /Users/me/other-project
```

### Canonical V1 rule now agreed

For command stages:

- `cwd` is optional
- default runtime working directory is the directory where `fabric` was invoked
- pipeline-local helper resolution and command runtime working directory are separate concepts
- `cwd` exists as an explicit override when a stage truly needs to run from a different location

## Settled Clarification 33: Command Stages May Declare a Stage-Level `env` Map

Decision:

- command stages may optionally declare an `env` map
- stage-level environment variables are allowed in V1
- this is in addition to inherited process environment, not a replacement for it

### Canonical example

```yaml
- id: publish
  executor: command
  command:
    program: python3
    args: ["publish.py"]
    env:
      OBSIDIAN_VAULT: "${OBSIDIAN_VAULT}"
      MODE: "append"
```

### Why this was chosen

This keeps command stages practical for real-world integrations without forcing users to create wrapper scripts for common configuration cases.

It also fits cleanly with the environment-variable rules already agreed earlier:

- Fabric reads from the process environment
- Fabric expands env references from the current process environment
- users can provide env vars via shell export, inline assignment, or shell startup files

### Agreed behavior

The command-stage environment model in V1 is now:

- the stage inherits the current process environment by default
- the stage may additionally declare `env` entries
- those entries may include variable references such as `${OBSIDIAN_VAULT}`
- Fabric resolves them during preflight where possible

### Why this matters

Without stage-level `env`, users would immediately need extra wrapper scripts just to inject small per-stage configuration differences.

That is unnecessary friction for:

- vault paths
- modes
- profile selectors
- output configuration
- environment-specific toggles

### Canonical V1 rule now agreed

For command stages:

- `env` is optional
- inherited process environment remains available
- stage-level env entries may extend or override inherited values for that stage
- env-variable expansion follows the earlier agreed process-environment model

## Settled Clarification 34: Command Stages May Declare an Optional Per-Stage `timeout`

Decision:

- command stages may optionally declare `timeout`
- `timeout` is not mandatory
- V1 will treat `timeout` as a numeric seconds value
- if a stage exceeds its timeout, that stage fails
- timeout failure is reported on stderr and the overall pipeline exits non-zero unless later behavior explicitly says otherwise for a downstream stage

### Canonical example

```yaml
- id: long_publish
  executor: command
  command:
    program: python3
    args: ["publish.py"]
    timeout: 300
```

### Why this was chosen

Pipelines that can run arbitrary external executables need a built-in way to deal with hung stages.

Without a per-stage timeout:

- one bad script can block the whole pipeline indefinitely
- operator experience degrades badly
- automation and CI usage become fragile

### Why timeout is stage-level, not just global

A single pipeline-wide timeout is too coarse for mixed workflows.

Different stages may have very different normal runtimes, for example:

- a quick normalization step
- a slower publish step
- a heavier external transformation step

Per-stage timeout keeps control local to the stage that needs it.

### Why the unit is seconds in V1

Using a numeric seconds value keeps the schema simpler for the first version.

Example:

```yaml
timeout: 300
```

This is easier to validate and document than introducing richer duration parsing in V1.

### Canonical V1 rule now agreed

For command stages:

- `timeout` is optional
- when present, it is an integer number of seconds
- timeout expiry causes stage failure
- timeout diagnostics go to stderr
- normal cleanup rules still apply after timeout-related failure

## Settled Clarification 35: Top-Level Pipeline Schema Uses a Small Required Core and a Few Optional Control Fields

Decision:

The canonical V1 top-level pipeline shape should be:

```yaml
version: 1
name: tech-note
description: Optional human-readable summary
accepts:
  - stdin
  - source
  - scrape_url
stages:
  - ...
```

### Required fields

- `version`
- `name`
- `stages`

### Optional fields

- `description`
- `accepts`

### Accepted source modes

Allowed V1 `accepts` values:

- `stdin`
- `source`
- `scrape_url`

The `source` value covers file or directory source paths.

### Default behavior when `accepts` is omitted

If `accepts` is omitted, the pipeline is treated as accepting any one supported source mode that the runner can provide.

## Settled Clarification 36: Stage Input Uses an Explicit `input` Object With a Strong Default

Decision:

Stage input in V1 should follow this model:

- first stage defaults to the pipeline source payload
- later stages default to the previous stage’s primary output payload
- a stage may explicitly override input with an `input` object

### Canonical forms

Explicit pipeline-source input:

```yaml
- id: normalize
  executor: command
  input:
    from: source
```

Explicit earlier-artifact input:

```yaml
- id: render
  executor: fabric_pattern
  input:
    from: artifact
    stage: map
    artifact: semantic_map
```

### Allowed V1 `input.from` values

- `source`
- `previous`
- `artifact`

## Settled Clarification 37: Declared Pipeline Artifacts Are Named, Relative, and Scoped to `.pipeline/<run-id>/`

Decision:

If a stage wants to expose pipeline-owned artifacts for later stages or inspection, it should declare them explicitly using an `artifacts` list.

### Canonical example

```yaml
- id: map
  executor: command
  command:
    program: python3
    args: ["build_map.py"]
  artifacts:
    - name: semantic_map
      path: semantic_map.json
      required: true
```

### Artifact rules in V1

- artifact `name` is required
- artifact `path` is required
- artifact `path` is relative to `.pipeline/<run-id>/`
- `required` is optional and defaults to `true`

### Important distinction

Command stages may still write anywhere on disk because command stages are trusted.

But only artifacts declared in the pipeline definition are treated as Fabric-managed pipeline artifacts for:

- later stage references
- manifests
- validation
- cleanup tracking

## Settled Clarification 38: `primary_output` Uses Explicit Source Selection and Artifact Names, Not Paths

Decision:

The canonical V1 `primary_output` shape should be:

For stdout-based primary output:

```yaml
primary_output:
  from: stdout
```

For artifact-based primary output:

```yaml
primary_output:
  from: artifact
  artifact: semantic_map
```

### Canonical V1 rule

- `primary_output.from` is required when a stage produces a primary payload
- if `from: artifact`, then `artifact:` must reference one declared artifact by name
- raw file paths are not used directly in `primary_output`

## Settled Clarification 39: Fabric-Pattern Stages Have a Small Explicit Schema

Decision:

The canonical V1 shape for a `fabric_pattern` stage should support:

- `pattern` as required
- optional `context`
- optional `strategy`
- optional `variables`
- optional stage-level `stream`

### Canonical example

```yaml
- id: render
  executor: fabric_pattern
  pattern: tech_note
  context: lecture
  strategy: precise
  variables:
    tone: concise
  stream: true
  final_output: true
  primary_output:
    from: stdout
```

### Canonical V1 rule

- `pattern` is required for `fabric_pattern` stages
- `context`, `strategy`, `variables`, and `stream` are optional
- other Fabric CLI flags are not automatically mirrored into stage schema in V1 unless later implementation truly needs them

## Settled Clarification 40: Builtin Stages Use an Explicit Builtin Name and Optional Config

Decision:

Builtin stages should use an explicit nested builtin descriptor.

### Canonical example

```yaml
- id: contract_check
  executor: builtin
  builtin:
    name: validate_declared_outputs
    config:
      require_primary_output: true
```

### Canonical V1 rule

- `builtin.name` is required for builtin stages
- `builtin.config` is optional
- builtin names map to Fabric-owned implementations, not user scripts

## Settled Clarification 41: Command Stage Stdout Is Captured Internally and Only Becomes Payload When Declared

Decision:

Command-stage stdout must be captured internally by Fabric.

It only becomes the stage’s primary payload when:

- `primary_output.from: stdout` is declared

### Canonical V1 rule

- command stdout is captured by the runner
- only declared stdout primary output becomes next-stage payload
- command stderr is treated as stage diagnostics and flows to Fabric stderr with stage context

## Settled Clarification 42: Validation Commands Use Normal CLI Exit Codes and Human-Readable Diagnostics

Decision:

Validation commands in V1 should behave as follows:

- valid pipeline -> exit code `0`
- invalid pipeline -> non-zero exit code
- human-readable validation result goes to stdout
- detailed diagnostics and errors go to stderr

### Canonical examples

Direct file validation:

```bash
fabric --validate-pipeline ~/.config/fabric/pipelines/tech-note.yaml
```

Installed pipeline validation:

```bash
fabric --pipeline tech-note --validate-only
```

## Settled Clarification 43: `--listpipelines` Should Integrate With `--shell-complete-list`

Decision:

The listing model should support both:

- normal human-readable listing
- raw-name output for shell completion

### Canonical behavior

Human-readable:

```bash
fabric --listpipelines
```

Raw names:

```bash
fabric --listpipelines --shell-complete-list
```

### Canonical V1 rule

- normal listing shows source and override visibility
- shell-complete mode emits only raw pipeline names suitable for completion

## Settled Clarification 44: Recommended V1 Implementation Boundary Is Now Closed Enough to Start Building

Decision:

The recommended V1 implementation boundary is now:

- one Fabric-native pipeline runner
- YAML-authored, schema-validated pipeline definitions
- built-in pipelines under `data/pipelines/`
- user-defined pipelines under `~/.config/fabric/pipelines/`
- strict preflight validation
- linear execution
- three executor types:
  - `command`
  - `fabric_pattern`
  - `builtin`
- one declared final output stage
- temporary run artifacts under `.pipeline/<run-id>/` in the current working directory with hybrid cleanup

## Operating Rule for Remaining Design Closure

Decision:

- from this point forward, low-risk design gaps should default to the recommended option without interrupting the user
- user intervention is only required when a decision would materially affect implementation, compatibility, or product semantics in a way that is still genuinely unclear

### Why this was chosen

The design has reached the point where repeatedly asking about every low-level choice creates more friction than value.

The working rule is now:

- continue using the brainstorming method internally
- keep recording chosen recommendations into this note
- only escalate true blockers or meaningful product forks back to the user

## Immediate Next Implementation Step

The next step after this brainstorming record is not more product exploration. It is implementation design closure:

- define the exact command contract for `fabric --pipeline`
- define source validation and exit-code behavior
- implement pipeline definition loading
- implement the first stage executors and terminal renderer
- implement preflight validation and listing commands

## Out of Scope for This Session

The following were intentionally not finalized here:

- exact JSON mode or structured output mode for pipelines
- exact partial-stage execution flags such as `--from-stage` and `--to-stage`

These remain implementation-design topics, not product-direction blockers.

## Implementation Checkpoint: 2026-03-07

This section records what is now implemented in Fabric after the design session and the first execution-focused implementation passes.

### Commits that established the current checkpoint

- `38da8c0b` `Add executable pipeline runner slice`
- `96bed3fb` `Add pipeline stage roles and builtins`
- `f886b511` `Tighten pipeline runner contracts`

### Confirmed implemented decisions

- Fabric remains the only operator-facing command surface.
- `--pipeline` exists and is the multi-stage execution entry point.
- `--listpipelines` exists and follows Fabric’s dedicated-list-flag style.
- `--validate-pipeline <file>` exists.
- `--pipeline <name> --validate-only` exists.
- pipeline definitions are YAML-authored and use the `.yaml` extension
- built-in definitions are loaded from `data/pipelines/`
- user-defined definitions are loaded from `~/.config/fabric/pipelines/`
- user-defined definitions override built-ins by name
- pipeline identity uses the filename stem and also requires matching `name:`
- pipeline mode requires exactly one input source:
  - stdin
  - `--source`
  - `--scrape_url`
- pipeline mode keeps chaining-safe IO semantics:
  - final payload on stdout
  - progress and diagnostics on stderr
- `-o` writes the final rendered output and still preserves stdout emission
- `.pipeline/<run-id>/` is created in the current working directory where `fabric` was invoked
- cleanup is temporary and hybrid:
  - per-run delayed cleanup
  - startup janitor cleanup
- the runner supports all three V1 executor types:
  - `command`
  - `fabric_pattern`
  - `builtin`
- command stages support:
  - `program`
  - `args`
  - optional `cwd`
  - optional `env`
  - optional `timeout`
- command stages are full-trust and may write outside `.pipeline/<run-id>/`
- stage-to-stage handoff uses the agreed hybrid model:
  - one primary payload for the next stage
  - optional artifacts on disk
- the primary output source must be declared explicitly
- exactly one final output stage is enforced in V1
- validation gates stdout when a validate stage exists
- missing validation emits an informational stderr warning and still allows final stdout
- publish-stage failures may still allow validated final stdout
- publish manifests are written from the finalized run snapshot
- role semantics are now real in code:
  - `validate`
  - `publish`
- role ordering is enforced:
  - validate/publish after the final output stage
  - publish after the last validate stage

### What is implemented but still intentionally minimal

- built-in pipeline inventory is minimal
- built-in example coverage is minimal
- operator-facing examples are still sparse
- profile inventory is still minimal

### What remains from the broader agreed direction

- first-class profile examples now include:
  - implemented:
    - `technical-study-guide`
    - `nontechnical-study-guide`
    - `note-enhancement`
    - `therapy-conversation-notes`
- richer operator-facing pipeline authoring documentation
- live provider-backed smoke coverage for model-backed stages under configured providers

### Verified implementation state at this checkpoint

- focused pipeline and CLI tests pass
- full `go test ./...` passes
- the current known gaps are product-surface and example coverage gaps, not confirmed correctness regressions in the implemented pipeline kernel

### Small doc/code alignment note

The implementation now has explicit and effective stage-role semantics that are stronger than the earliest drafts in this note.

The practical source of truth after this checkpoint is:

- this brainstorming note for design history and decisions
- `docs/spec.md` for the current technical contract
- the code for exact runtime behavior

## Implementation Checkpoint Addendum: 2026-03-07 (Zoom Tech Note Parity Port)

This addendum records the next major implementation step after the core pipeline kernel checkpoint above.

### What is now implemented in the working tree

- Fabric now contains two built-in Zoom parity pipelines:
  - `zoom-tech-note`
  - `zoom-tech-note-deep-pass`
- The implementation is Fabric-owned and shippable.
- It no longer depends on a private local Zoom repository at runtime.

### What was copied or ported into Fabric

- deterministic Zoom scripts were copied into:
  - `scripts/pipelines/zoom-tech-note/ingest_zoom_captions.py`
  - `scripts/pipelines/zoom-tech-note/validate_coverage.py`
  - `scripts/pipelines/zoom-tech-note/publish_tutorial_notes.py`
- Stage 1/2/3 prompts were ported into Fabric patterns under:
  - `data/patterns/zoom_stage1_refine/system.md`
  - `data/patterns/zoom_stage2_synthesize/system.md`
  - `data/patterns/zoom_stage3_enhance/system.md`
- Fabric-specific helper scripts were added for:
  - source/session preparation
  - ingest-stage orchestration
  - Stage 1 materialization
  - Stage 2 materialization
  - Stage 3 materialization
  - deep pass
  - validation-stage orchestration
  - publish-stage orchestration

### Important compatibility note

The old Zoom runner used a tool-enabled model flow that could write multiple files directly during an AI stage.

Fabric pattern stages emit content, not multiple file writes.

So the required compatibility adaptation is:

- `generate` stage via `fabric_pattern`
- `materialize` stage via deterministic command script

This is not a product-semantic change.
It is the minimum execution-model adaptation needed to preserve the old artifact contract inside Fabric.

### What parity now covers

- stable session-root artifacts are written into the real source session directory
- temporary Fabric run metadata still lives under `.pipeline/<run-id>/` in the invocation directory
- Stage 1/2/3 outputs preserve the original file contract
- the deep-pass variant preserves the stricter quality gate as a separate built-in pipeline

### Additional implementation detail now enforced

Pipeline stages may reference pattern files by relative path.

This is now supported in preflight and runtime resolution so a built-in pipeline can own its patterns inside the Fabric repo without relying on globally installed pattern names.

### Verification completed for this parity slice

- `python3 -m py_compile scripts/pipelines/zoom-tech-note/*.py`
- `go run ./cmd/fabric --validate-pipeline data/pipelines/zoom-tech-note.yaml`
- `go run ./cmd/fabric --validate-pipeline data/pipelines/zoom-tech-note-deep-pass.yaml`
- `go test ./internal/pipeline -run 'TestZoomTechNote' -count=1`
- `go test ./...`

### What still remains after the Zoom parity slice

- built-in quick-mode note alignment is now implemented through:
  - `techNote`
  - `nontechNote`
- first-class built-in generic note pipelines beyond quick mode and Zoom parity now include:
  - implemented:
    - `technical-study-guide`
    - `nontechnical-study-guide`
    - `note-enhancement`
    - `therapy-conversation-notes`
- richer operator-facing examples and authoring docs
- live provider-backed smoke validation for model-backed stages

## Implementation Checkpoint Addendum: 2026-03-11 (Phase 3 Runtime/CLI Controls)

This addendum records the next major runtime surface milestone after the Phase 2 profile and parity slices.

### Implemented in this checkpoint

- pipeline partial execution controls:
  - `--from-stage`
  - `--to-stage`
  - `--only-stage`
- pipeline JSON lifecycle events on stderr:
  - `--pipeline-events-json`
- pipeline dry-run/introspection output:
  - `fabric --pipeline <name> --dry-run`
- runner contract behavior for stage slicing:
  - non-selected stages are explicitly marked `skipped` in manifests
  - selected-stage progress counters are ordinal-relative to the selected slice
  - missing previous-stage payloads now fail explicitly instead of silently passing empty input

### Verification completed for this checkpoint

- focused tests:
  - `go test ./internal/pipeline ./internal/cli`
- full suite:
  - `go test ./...`
- built-in catalog sanity check:
  - `go run ./cmd/fabric --listpipelines`

### Remaining work after this checkpoint

- richer operator docs and authoring examples
- broader profile inventory and deep-pass variants where productized
- live provider-backed smoke coverage for model-backed stages

### Current confidence after this addendum

- confidence in the implemented Fabric pipeline kernel remains high
- confidence in the built-in Zoom parity port is high after the current verification set
- the remaining gaps are now mostly product-surface expansion, documentation, and continued parity hardening rather than runner-model uncertainty
