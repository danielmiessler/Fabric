package cli

import (
	"fmt"
	"os"
	"strconv"

	openai "github.com/openai/openai-go"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/gemini"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// handleListingCommands handles listing-related commands
// Returns (handled, error) where handled indicates if a command was processed and should exit
func handleListingCommands(currentFlags *Flags, fabricDb *fsdb.Db, registry *core.PluginRegistry) (handled bool, err error) {
	if currentFlags.LatestPatterns != "0" {
		var parsedToInt int
		if parsedToInt, err = strconv.Atoi(currentFlags.LatestPatterns); err != nil {
			return true, err
		}

		if err = fabricDb.Patterns.PrintLatestPatterns(parsedToInt); err != nil {
			return true, err
		}
		return true, nil
	}

	if currentFlags.ListPatterns {
		err = fabricDb.Patterns.ListNames(currentFlags.ShellCompleteOutput)
		return true, err
	}

	if currentFlags.ListAllModels {
		var models *ai.VendorsModels
		if models, err = registry.VendorManager.GetModels(); err != nil {
			return true, err
		}
		if currentFlags.ShellCompleteOutput {
			models.Print(true)
		} else {
			models.PrintWithVendor(false, registry.Defaults.Vendor.Value, registry.Defaults.Model.Value)
		}
		return true, nil
	}

	if currentFlags.ListAllContexts {
		err = fabricDb.Contexts.ListNames(currentFlags.ShellCompleteOutput)
		return true, err
	}

	if currentFlags.ListAllSessions {
		err = fabricDb.Sessions.ListNames(currentFlags.ShellCompleteOutput)
		return true, err
	}

	if currentFlags.ListStrategies {
		err = registry.Strategies.ListStrategies(currentFlags.ShellCompleteOutput)
		return true, err
	}

	if currentFlags.ListVendors {
		err = registry.ListVendors(os.Stdout)
		return true, err
	}

	if currentFlags.ListGeminiVoices {
		voicesList := gemini.ListGeminiVoices(currentFlags.ShellCompleteOutput)
		fmt.Print(voicesList)
		return true, nil
	}

	if currentFlags.ListTranscriptionModels {
		listTranscriptionModels(currentFlags.ShellCompleteOutput)
		return true, nil
	}

	return false, nil
}

// listTranscriptionModels lists all available transcription models
func listTranscriptionModels(shellComplete bool) {
	models := []string{
		string(openai.AudioModelWhisper1),
		string(openai.AudioModelGPT4oMiniTranscribe),
		string(openai.AudioModelGPT4oTranscribe),
	}

	if shellComplete {
		for _, model := range models {
			fmt.Println(model)
		}
	} else {
		fmt.Println(i18n.T("available_transcription_models"))
		for _, model := range models {
			fmt.Printf("  %s\n", model)
		}
	}
}
