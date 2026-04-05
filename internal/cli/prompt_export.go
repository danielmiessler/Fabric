package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/chatfmt"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
)

func handlePromptExport(
	currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (handled bool, err error) {
	if !currentFlags.PrintPrompt {
		return false, nil
	}

	if err = validatePromptExportFlags(currentFlags); err != nil {
		return true, err
	}

	var prompt string
	if prompt, err = renderPromptExport(currentFlags, registry, strings.Join(os.Args[1:], " "), messageTools); err != nil {
		return true, err
	}

	if err = outputPromptExport(prompt, currentFlags.Output, currentFlags.Copy, os.Stdout, CopyToClipboard); err != nil {
		return true, err
	}

	return true, nil
}

func validatePromptExportFlags(currentFlags *Flags) error {
	if currentFlags.DryRun {
		return fmt.Errorf("%s", i18n.T("print_prompt_error_dry_run"))
	}
	if currentFlags.OutputSession {
		return fmt.Errorf("%s", i18n.T("print_prompt_error_output_session"))
	}
	return nil
}

func renderPromptExport(
	currentFlags *Flags, registry *core.PluginRegistry, meta string, messageTools string) (string, error) {
	if registry == nil || registry.Db == nil {
		return "", fmt.Errorf("registry database not initialized")
	}

	flagsCopy := *currentFlags
	if messageTools != "" {
		flagsCopy.Message = AppendMessage(flagsCopy.Message, messageTools)
	}

	chatReq, err := flagsCopy.BuildChatRequest(meta)
	if err != nil {
		return "", err
	}

	if chatReq.Language == "" &&
		registry.Language != nil &&
		registry.Language.DefaultLanguage != nil {
		chatReq.Language = registry.Language.DefaultLanguage.Value
	}

	chatter := core.NewChatter(registry.Db)
	session, err := chatter.BuildSessionQuiet(chatReq, currentFlags.Raw)
	if err != nil {
		return "", err
	}

	return chatfmt.FormatMessages(session.GetVendorMessages()), nil
}

func outputPromptExport(
	prompt string,
	outputPath string,
	copy bool,
	stdout io.Writer,
	copyToClipboard func(string) error,
) error {
	if outputPath == "" {
		outputText := prompt
		if !strings.HasSuffix(outputText, "\n") {
			outputText += "\n"
		}
		if _, err := io.WriteString(stdout, outputText); err != nil {
			return err
		}
	} else {
		if err := CreateOutputFile(prompt, outputPath); err != nil {
			return err
		}
	}

	if copy {
		return copyToClipboard(prompt)
	}

	return nil
}
