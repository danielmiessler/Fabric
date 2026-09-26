package cli

import (
	"context"
	"fmt"
	"maps"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/util"
	"gopkg.in/yaml.v3"
)

// Workflow describes a simple sequential pattern composition.
// Each step's output becomes the next step's input.
type Workflow struct {
	Name  string         `yaml:"name,omitempty"`
	Steps []WorkflowStep `yaml:"steps"`
}

type WorkflowStep struct {
	Pattern   string            `yaml:"pattern"`
	Input     string            `yaml:"input,omitempty"`
	Variables map[string]string `yaml:"variables,omitempty"`
	Model     string            `yaml:"model,omitempty"`
	Vendor    string            `yaml:"vendor,omitempty"`
}

// patternResolver is the minimal contract required to verify that a pattern
// name points at a real system prompt before the workflow starts running.
type patternResolver interface {
	GetRaw(name string) (*fsdb.Pattern, error)
}

// stepLabel produces the canonical "[step N/TOTAL pattern]" prefix used by
// every workflow log line and error so that users grep for one shape only.
func stepLabel(idx, total int, pattern string) string {
	return fmt.Sprintf("[step %d/%d %s]", idx+1, total, pattern)
}

// LoadWorkflow parses a workflow definition from a YAML or JSON file (JSON is
// valid YAML). It does not validate the contents – call Validate before running.
func LoadWorkflow(path string) (wf *Workflow, err error) {
	absPath, err := util.GetAbsolutePath(path)
	if err != nil {
		return nil, fmt.Errorf("invalid workflow path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("error reading workflow file: %w", err)
	}

	wf = &Workflow{}
	if err = yaml.Unmarshal(data, wf); err != nil {
		return nil, fmt.Errorf("error parsing workflow file: %w", err)
	}
	return
}

// Validate performs all pre-flight checks on a workflow:
//   - at least one step is defined
//   - every step has a non-empty pattern name
//   - no pattern appears twice in a row (accidental duplication)
//   - every named pattern resolves against the patterns store (unless the
//     name looks like a file path, which is resolved lazily at run time)
//
// Passing nil for patterns skips the existence check – useful for unit tests
// that don't have a configured fsdb.
func (wf *Workflow) Validate(patterns patternResolver) error {
	if len(wf.Steps) == 0 {
		return fmt.Errorf("workflow has no steps")
	}

	var prev string
	total := len(wf.Steps)
	for i, step := range wf.Steps {
		name := strings.TrimSpace(step.Pattern)
		label := stepLabel(i, total, name)

		if name == "" {
			return fmt.Errorf("%s missing required field 'pattern'", stepLabel(i, total, "?"))
		}
		if name == prev {
			return fmt.Errorf("%s repeats the previous step", label)
		}
		prev = name

		if patterns != nil && !fsdb.LooksLikePatternFilePath(name) {
			if _, err := patterns.GetRaw(name); err != nil {
				return fmt.Errorf("%s pattern not found: %w", label, err)
			}
		}
	}
	return nil
}

// runWorkflow runs the validated steps in order and pipes each output into
// the next step. It returns the final step output; the caller prints it.
func runWorkflow(
	registry *core.PluginRegistry,
	wf *Workflow,
	input string,
	flags *Flags,
	chatOptions *domain.ChatOptions,
) (result string, err error) {
	language := flags.Language
	if language == "" {
		language = registry.Language.DefaultLanguage.Value
	}
	meta := strings.Join(os.Args[1:], " ")

	total := len(wf.Steps)
	for i, step := range wf.Steps {
		isLast := i == total-1
		label := stepLabel(i, total, step.Pattern)
		stepErr := func(e error) error {
			return fmt.Errorf("%s failed: %w", label, e)
		}

		// Per-step input override: trim first so a YAML value that is only
		// whitespace/newlines does NOT count as an override. Only a non-empty
		// trimmed string replaces the carried input.
		stepInput, usedOverride := resolveStepInputWithOverride(step, input)
		if usedOverride {
			fmt.Fprintf(os.Stderr, "%s using custom input (%d chars)\n", label, len(stepInput))
		}

		// Lightweight progress indicator on stderr so stdout stays pipe-clean.
		fmt.Fprintf(os.Stderr, "%s running...\n", label)

		model := flags.Model
		if step.Model != "" {
			model = step.Model
		}
		vendor := flags.Vendor
		if step.Vendor != "" {
			vendor = step.Vendor
		}

		// Only stream the final step; intermediate output is captured whole.
		stream := flags.Stream && isLast

		var chatter *core.Chatter
		if chatter, err = registry.GetChatter(model, flags.ModelContextLength,
			vendor, stream, flags.DryRun); err != nil {
			return "", stepErr(err)
		}

		req := &domain.ChatRequest{
			ContextName:           flags.Context,
			PatternName:           step.Pattern,
			PatternVariables:      mergeVars(flags.PatternVariables, step.Variables),
			InputHasVars:          flags.InputHasVars,
			NoVariableReplacement: flags.NoVariableReplacement,
			StrategyName:          flags.Strategy,
			Language:              language,
			Meta:                  meta,
		}
		if stepInput != "" {
			req.Message = &chat.ChatCompletionMessage{
				Role:    chat.ChatMessageRoleUser,
				Content: stepInput,
			}
		}

		// Shallow copy so per-step mutations don't leak between iterations.
		opts := *chatOptions
		opts.Model = model
		opts.Quiet = !isLast

		debuglog.Log("%s model=%s vendor=%s inputLen=%d override=%t\n",
			label, model, vendor, len(stepInput), usedOverride)

		var session *fsdb.Session
		if session, err = chatter.Send(context.Background(), req, &opts); err != nil {
			return "", stepErr(err)
		}

		out, extractErr := extractStepOutput(session)
		if extractErr != nil {
			return "", stepErr(extractErr)
		}

		fmt.Fprintf(os.Stderr, "%s done (%d chars)\n", label, len(out))

		result = out
		input = out // pipe into next step
	}
	return result, nil
}

// extractStepOutput pulls the assistant response from a completed session and
// guards against the two realistic edge cases: no message appended at all, or
// a message whose content is blank/whitespace-only.
func extractStepOutput(session *fsdb.Session) (string, error) {
	if session == nil {
		return "", fmt.Errorf("model returned no session")
	}
	last := session.GetLastMessage()
	if last == nil {
		return "", fmt.Errorf("model returned no message")
	}
	out := strings.TrimSpace(last.Content)
	if out == "" {
		return "", fmt.Errorf("model returned empty output")
	}
	return out, nil
}

// handleWorkflowProcessing wires the CLI flags into the load → validate → run
// pipeline and handles terminal output (print / copy / file).
func handleWorkflowProcessing(currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (err error) {
	wf, err := LoadWorkflow(currentFlags.Workflow)
	if err != nil {
		return
	}

	if err = wf.Validate(registry.Db.Patterns); err != nil {
		return fmt.Errorf("workflow %s is invalid: %w", currentFlags.Workflow, err)
	}

	if messageTools != "" {
		currentFlags.AppendMessage(messageTools)
	}

	var chatOptions *domain.ChatOptions
	if chatOptions, err = currentFlags.BuildChatOptions(); err != nil {
		return
	}

	var result string
	if result, err = runWorkflow(registry, wf, currentFlags.Message, currentFlags, chatOptions); err != nil {
		return
	}

	// Print final result unless it was already streamed live.
	if !currentFlags.Stream || chatOptions.SuppressThink {
		fmt.Println(result)
	}

	if currentFlags.Copy {
		if err = CopyToClipboard(result); err != nil {
			return
		}
	}

	if currentFlags.Output != "" {
		err = CreateOutputFile(result, currentFlags.Output)
	}
	return
}

// resolveStepInputWithOverride picks the user input for a workflow step.
// A non-empty (after TrimSpace) step.Input overrides the carried value;
// whitespace-only or missing input falls back to carriedInput.
func resolveStepInputWithOverride(step WorkflowStep, carriedInput string) (chosenInput string, usedOverride bool) {
	if override := strings.TrimSpace(step.Input); override != "" {
		return override, true
	}
	return carriedInput, false
}

func mergeVars(base, override map[string]string) map[string]string {
	if len(base) == 0 && len(override) == 0 {
		return nil
	}
	out := maps.Clone(base)
	if out == nil {
		out = make(map[string]string, len(override))
	}
	maps.Copy(out, override)
	return out
}
