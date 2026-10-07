package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chatfmt"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
)

// handlePromptExport prints the rendered prompt. Cli validates the flags before it calls this function.
func handlePromptExport(
	currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (handled bool, err error) {
	if !currentFlags.PrintPrompt {
		return false, nil
	}

	var prompt string
	if prompt, err = renderPromptExport(currentFlags, registry, strings.Join(os.Args[1:], " "), messageTools); err != nil {
		return true, err
	}

	if currentFlags.Output == "" {
		fmt.Print(prompt)
	} else if err = CreateOutputFile(prompt, currentFlags.Output); err != nil {
		return true, err
	}
	if currentFlags.Copy {
		err = CopyToClipboard(prompt)
	}
	return true, err
}

func validatePromptExportFlags(currentFlags *Flags) error {
	if currentFlags.DryRun {
		return errors.New(i18n.T("print_prompt_error_dry_run"))
	}
	if currentFlags.OutputSession {
		return errors.New(i18n.T("print_prompt_error_output_session"))
	}
	if currentFlags.Workflow != "" {
		return errors.New(i18n.T("print_prompt_error_workflow"))
	}
	if currentFlags.Output != "" && IsAudioFormat(currentFlags.Output) {
		return fmt.Errorf(i18n.T("print_prompt_error_audio_output"), currentFlags.Output)
	}
	return nil
}

func renderPromptExport(
	currentFlags *Flags, registry *core.PluginRegistry, meta string, messageTools string) (string, error) {
	flagsCopy := *currentFlags
	if messageTools != "" {
		flagsCopy.Message = AppendMessage(flagsCopy.Message, messageTools)
	}

	chatReq, err := flagsCopy.BuildChatRequest(meta)
	if err != nil {
		return "", err
	}

	if chatReq.Language == "" {
		chatReq.Language = registry.Language.DefaultLanguage.Value
	}

	// Resolve the vendor only for a named model, as Send does. Without a model, the export
	// makes no network call and shows the non-raw structure unless --raw is set.
	chatter := core.NewChatter(registry.Db)
	if applyPatternModel(&flagsCopy); flagsCopy.Model != "" {
		if chatter, err = registry.GetChatter(flagsCopy.Model, flagsCopy.ModelContextLength,
			flagsCopy.Vendor, false, false); err != nil {
			return "", err
		}
	}

	session, err := chatter.BuildSession(chatReq, flagsCopy.Raw || chatter.NeedsRawMode(), false)
	if err != nil {
		return "", err
	}

	return chatfmt.FormatMessages(session.GetVendorMessages()), nil
}
