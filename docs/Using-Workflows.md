# Using workflows

A workflow runs a list of patterns in sequence. Fabric sends the output of each step to the next step as its input. You write the list one time in a YAML or JSON file. Then you run all the steps with one command.

Without a workflow, you connect the patterns with shell pipes:

```bash
cat transcript.txt | fabric -p summarize_meeting | fabric -p create_formal_email
```

With a workflow, you run the same patterns with one command:

```bash
cat transcript.txt | fabric --workflow meeting-followup.yaml
```

A workflow file also lets you set a different model, vendor, or variables for each step.

## Write a workflow file

A workflow file has an optional `name` and a list of `steps`. Each step must have a `pattern`.

```yaml
name: meeting-followup
steps:
  - pattern: summarize_meeting
  - pattern: create_formal_email
```

Fabric also reads JSON files, because JSON is also YAML:

```json
{"steps": [{"pattern": "summarize_meeting"}, {"pattern": "create_formal_email"}]}
```

### Step fields

| Field       | Necessary | What it does                                                                                       |
| ----------- | --------- | -------------------------------------------------------------------------------------------------- |
| `pattern`   | Yes       | The pattern name, or a path to a pattern file (a path starts with `/`, `~`, `.`, or `\`).          |
| `model`     | No        | The model for this step. If you do not set it, the step uses `-m` or your default model.           |
| `vendor`    | No        | The vendor for this step. If you do not set it, the step uses `-V` or your default vendor.         |
| `variables` | No        | Pattern variables for this step. They replace `-v` values that have the same name.                 |
| `input`     | No        | Your own input text for this step. It replaces the output of the previous step for this step only. |

## Run a workflow

1. Send the input to Fabric on stdin, or give it as a message argument.
2. Add `--workflow` and the path to the workflow file.

```bash
cat notes.txt | fabric --workflow my-workflow.yaml
fabric --workflow my-workflow.yaml "Text to process"
fabric -y "https://youtu.be/<id>" | fabric --workflow my-workflow.yaml
```

Fabric writes progress lines to stderr, for example `[step 1/2 summarize_meeting] running...`. Only the output of the last step goes to stdout. Thus you can send the result to a file or to a different command.

### Flags that apply to all steps

These flags have the same effect as in a run with one pattern:

- `-m` and `-V` set the model and vendor for each step that does not set its own.
- `-v` sets pattern variables. A step `variables` value with the same name replaces the `-v` value.
- `-C` (context), `--strategy`, and `-g` (language) apply to each step.
- `-c` copies the last output to the clipboard. `-o` writes it to a file.
- `--dry-run` shows the prompts, but does not send them to a model.

`--stream` applies to the last step only. Fabric does not show the output of the other steps. It gives that output only to the next step.

Fabric ignores `--session` when you use `--workflow`.

## Checks before the run

Before Fabric runs the first step, it examines the full file. It stops with an error if:

- The file has no steps.
- A step has no `pattern`.
- A step uses the same pattern as the step before it.
- A pattern name is not in your patterns directory.

Fabric does not examine pattern file paths before the run. It loads them when that step starts.

If a step has an error, Fabric stops. The error message shows the step, for example `[step 2/3 create_formal_email] failed: ...`.

## Examples

The [`examples`](./examples/) directory has three workflow files. Each file uses patterns that Fabric installs.

### Meeting transcript to recap email

[`meeting-followup.yaml`](./examples/meeting-followup.yaml) makes a summary of a meeting transcript, then writes a recap email from the summary.

```bash
cat transcript.txt | fabric --workflow docs/examples/meeting-followup.yaml --copy
```

`--copy` puts the email on the clipboard.

### Lecture to flash cards

[`study-kit.yaml`](./examples/study-kit.yaml) makes notes from a lecture or talk, then makes flash cards from the notes.

```bash
fabric -y "https://youtu.be/<id>" | fabric --workflow docs/examples/study-kit.yaml -o cards.md
```

To make a quiz, change `create_flash_cards` to `create_quiz`.

### Article to fact-checked brief

[`fact-checked-brief.yaml`](./examples/fact-checked-brief.yaml) examines each claim in an article and gives the evidence for and against it. Then it writes a five-sentence summary of the analysis.

```bash
fabric -u "https://example.com/article" | fabric --workflow docs/examples/fact-checked-brief.yaml
```

This example sets a different model for each step. A strong model does the analysis of the claims. A small, fast model writes the summary, because this step does not need a strong model. Change the `model` and `vendor` values to models that are available to you. To see the list, run `fabric -L`.

## Give one step its own input

Usually each step reads the output of the step before it. To give a step different text, set `input`:

```yaml
steps:
  - pattern: summarize
  - pattern: create_tags
    input: "Tags for a blog post about home network security"
```

The step after a step with `input` reads the output of that step, as usual. If `input` contains only spaces or blank lines, Fabric ignores it.
