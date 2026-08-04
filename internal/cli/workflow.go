package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
	Name  string         `yaml:"name,omitempty"  json:"name,omitempty"`
	Steps []WorkflowStep `yaml:"steps"           json:"steps"`
}

type WorkflowStep struct {
	Pattern   string            `yaml:"pattern"             json:"pattern"`
	Input     string            `yaml:"input,omitempty"     json:"input,omitempty"`
	Variables map[string]string `yaml:"variables,omitempty" json:"variables,omitempty"`
	Model     string            `yaml:"model,omitempty"     json:"model,omitempty"`
	Vendor    string            `yaml:"vendor,omitempty"    json:"vendor,omitempty"`
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

// LoadWorkflow parses a workflow definition from a YAML or JSON file. It does
// not validate the contents – call Validate before running.
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
	ext := strings.ToLower(filepath.Ext(absPath))
	if ext == ".json" {
		if err = json.Unmarshal(data, wf); err != nil {
			return nil, fmt.Errorf("error parsing workflow JSON: %w", err)
		}
	} else {
		if err = yaml.Unmarshal(data, wf); err != nil {
			return nil, fmt.Errorf("error parsing workflow YAML: %w", err)
		}
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
	if wf == nil {
		return fmt.Errorf("workflow is nil")
	}
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

		if patterns != nil && !isFilePathPattern(name) {
			if _, err := patterns.GetRaw(name); err != nil {
				return fmt.Errorf("%s pattern not found: %w", label, err)
			}
		}
	}
	return nil
}

// runWorkflow executes a validated Workflow sequentially.
//
// Contract
//   - Input:  wf must have passed wf.Validate(); chatOptions comes from flags.
//   - Piping: each step's assistant output becomes the carried input for the
//     next step unless that step declares its own `input:` override.
//   - Reuse:  every step goes through the standard registry.GetChatter → Send
//     path, so --stream (last step only), --dry-run, -m/-V, -v vars, context,
//     strategy and language all behave exactly as in a single-pattern run.
//   - Errors: the run stops at the first failing step; the error is wrapped
//     with the canonical stepLabel prefix so it reads
//     "[step N/TOTAL pattern] failed: <cause>".
//   - Output: returns the trimmed content of the final step's assistant
//     message; printing / copy / file-output is left to the caller.
func runWorkflow(
	registry *core.PluginRegistry,
	wf *Workflow,
	input string,
	flags *Flags,
	chatOptions *domain.ChatOptions,
) (result string, err error) {
	globalStream := flags.Stream
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
		stream := globalStream && isLast

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
		if session, err = chatter.Send(req, &opts); err != nil {
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

// isFilePathPattern mirrors the heuristic used by PatternsEntity.loadPattern:
// patterns that look like file paths are loaded directly from disk rather
// than from the patterns database, so we can't validate them up front.
func isFilePathPattern(p string) bool {
	return strings.HasPrefix(p, "/") ||
		strings.HasPrefix(p, "~") ||
		strings.HasPrefix(p, ".") ||
		strings.HasPrefix(p, "\\")
}

func mergeVars(base, override map[string]string) map[string]string {
	if len(base) == 0 && len(override) == 0 {
		return nil
	}
	out := make(map[string]string, len(base)+len(override))
	for k, v := range base {
		out[k] = v
	}
	for k, v := range override {
		out[k] = v
	}
	return out
}
