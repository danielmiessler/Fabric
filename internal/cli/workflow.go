package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"gopkg.in/yaml.v3"
)

// loadWorkflow reads and validates a workflow YAML file.
func loadWorkflow(path string) (*domain.WorkflowDefinition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("workflow_error_read_file"), err)
	}

	var wf domain.WorkflowDefinition
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, fmt.Errorf(i18n.T("workflow_error_parse_file"), err)
	}

	if len(wf.Steps) == 0 {
		return nil, fmt.Errorf("%s", i18n.T("workflow_error_no_steps"))
	}

	for idx, step := range wf.Steps {
		if step.Pattern == "" {
			return nil, fmt.Errorf(i18n.T("workflow_error_step_no_pattern"), idx+1)
		}
	}

	return &wf, nil
}

// validateWorkflowPatterns checks that every step references a pattern that exists in the registry.
func validateWorkflowPatterns(wf *domain.WorkflowDefinition, registry *core.PluginRegistry) error {
	names, err := registry.Db.Patterns.GetNames()
	if err != nil {
		return err
	}

	available := make(map[string]bool, len(names))
	for _, name := range names {
		available[name] = true
	}

	for idx, step := range wf.Steps {
		if !available[step.Pattern] {
			return fmt.Errorf(i18n.T("workflow_error_pattern_not_found"), idx+1, step.Pattern)
		}
	}

	return nil
}

// runWorkflow executes workflow steps sequentially, piping each step's output as input to the next.
// Intermediate steps run non-streaming with Quiet=true; the last step uses the caller's stream setting.
func runWorkflow(registry *core.PluginRegistry, wf *domain.WorkflowDefinition, input string, currentFlags *Flags, chatOptions *domain.ChatOptions) (string, error) {
	totalSteps := len(wf.Steps)
	debuglog.Debug(debuglog.Basic, "workflow: starting %d steps\n", totalSteps)

	for idx, step := range wf.Steps {
		isLast := idx == totalSteps-1
		stepNum := idx + 1

		debuglog.Debug(debuglog.Detailed, "workflow step %d/%d: pattern=%s input_len=%d\n", stepNum, totalSteps, step.Pattern, len(input))

		chatter, err := registry.GetChatter(
			currentFlags.Model, currentFlags.ModelContextLength,
			currentFlags.Vendor, isLast && currentFlags.Stream, currentFlags.DryRun,
		)
		if err != nil {
			return "", fmt.Errorf(i18n.T("workflow_error_step_failed"), stepNum, totalSteps, step.Pattern, err)
		}

		chatReq := &domain.ChatRequest{
			ContextName:           currentFlags.Context,
			PatternName:           step.Pattern,
			PatternVariables:      mergeVariables(currentFlags.PatternVariables, step.Variables),
			InputHasVars:          currentFlags.InputHasVars,
			NoVariableReplacement: currentFlags.NoVariableReplacement,
			Meta:                  strings.Join(os.Args[1:], " "),
		}

		if chatReq.Language == "" {
			chatReq.Language = registry.Language.DefaultLanguage.Value
		}

		if input != "" {
			chatReq.Message = &chat.ChatCompletionMessage{
				Role:    chat.ChatMessageRoleUser,
				Content: strings.TrimSpace(input),
			}
		}

		stepOpts := *chatOptions
		if !isLast {
			stepOpts.Quiet = true
			fmt.Fprintf(os.Stderr, i18n.T("workflow_step_progress"), stepNum, totalSteps, step.Pattern)
		}

		session, err := chatter.Send(chatReq, &stepOpts)
		if err != nil {
			return "", fmt.Errorf(i18n.T("workflow_error_step_failed"), stepNum, totalSteps, step.Pattern, err)
		}

		input = session.GetLastMessage().Content
		debuglog.Debug(debuglog.Detailed, "workflow step %d/%d: complete output_len=%d\n", stepNum, totalSteps, len(input))
	}

	return input, nil
}

// handleWorkflowProcessing loads a workflow file, validates patterns, runs the workflow,
// and handles final output (print, copy, file).
func handleWorkflowProcessing(currentFlags *Flags, registry *core.PluginRegistry, messageTools string) error {
	wf, err := loadWorkflow(currentFlags.Workflow)
	if err != nil {
		return err
	}

	if err := validateWorkflowPatterns(wf, registry); err != nil {
		return err
	}

	input := currentFlags.Message
	if messageTools != "" {
		input = AppendMessage(input, messageTools)
	}

	if strings.TrimSpace(input) == "" {
		fmt.Fprint(os.Stderr, i18n.T("workflow_warning_empty_input"))
	}

	chatOptions, err := currentFlags.BuildChatOptions()
	if err != nil {
		return err
	}

	result, err := runWorkflow(registry, wf, input, currentFlags, chatOptions)
	if err != nil {
		return err
	}

	if !currentFlags.Stream || currentFlags.SuppressThink {
		fmt.Println(result)
	}

	if currentFlags.Copy {
		if err := CopyToClipboard(result); err != nil {
			return err
		}
	}

	if currentFlags.Output != "" {
		if err := CreateOutputFile(result, currentFlags.Output); err != nil {
			return err
		}
	}

	fmt.Fprintf(os.Stderr, i18n.T("workflow_complete"), len(wf.Steps))

	return nil
}

// mergeVariables returns a merged map where step-level variables override CLI-level variables.
func mergeVariables(cliVars, stepVars map[string]string) map[string]string {
	merged := make(map[string]string)
	for k, v := range cliVars {
		merged[k] = v
	}
	for k, v := range stepVars {
		merged[k] = v
	}
	return merged
}
