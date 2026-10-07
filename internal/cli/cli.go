package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	"github.com/danielmiessler/fabric/internal/tools/converter"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
)

// Cli Controls the cli. It takes in the flags and runs the appropriate functions
func Cli(version string) (err error) {
	var currentFlags *Flags
	if currentFlags, err = Init(); err != nil {
		return
	}

	if _, err = i18n.Init(currentFlags.Language); err != nil {
		return
	}

	if currentFlags.Setup {
		if err = ensureEnvFile(); err != nil {
			return
		}
	}

	if currentFlags.Version {
		fmt.Println(version)
		return
	}

	var registry, err2 = initializeFabric()
	if err2 != nil {
		if !currentFlags.Setup {
			debuglog.Log("%s\n", err2.Error())
			currentFlags.Setup = true
		}
		// The handlers below dereference registry.
		if registry == nil {
			return err2
		}
	}

	if registry != nil {
		configureOpenAIResponsesAPI(registry, currentFlags.DisableResponsesAPI)
	}

	var handled bool
	if handled, err = handleSetupAndServerCommands(currentFlags, registry, version); err != nil || handled {
		return
	}

	if handled, err = handleConfigurationCommands(currentFlags, registry); err != nil || handled {
		return
	}

	if handled, err = handleListingCommands(currentFlags, registry.Db, registry); err != nil || handled {
		return
	}

	if handled, err = handleManagementCommands(currentFlags, registry.Db); err != nil || handled {
		return
	}

	if handled, err = handleExtensionCommands(currentFlags, registry); err != nil || handled {
		return
	}

	// Validate prompt-export-only combinations before any expensive preprocessing.
	if currentFlags.PrintPrompt {
		if err = validatePromptExportFlags(currentFlags); err != nil {
			return
		}
	}

	if currentFlags.TranscribeFile != "" {
		var transcriptionMessage string
		if transcriptionMessage, err = handleTranscription(currentFlags, registry); err != nil {
			return
		}
		currentFlags.Message = AppendMessage(currentFlags.Message, transcriptionMessage)
	}

	if currentFlags.HtmlReadability {
		if msg, cleanErr := converter.HtmlReadability(currentFlags.Message); cleanErr != nil {
			fmt.Println(i18n.T("html_readability_error"), cleanErr)
		} else {
			currentFlags.Message = msg
		}
	}

	var messageTools string
	if messageTools, err = handleToolProcessing(currentFlags, registry); err != nil {
		return
	}

	// handleToolProcessing already printed the tool output.
	if messageTools != "" && !currentFlags.IsChatRequest() {
		return nil
	}

	if handled, err = handlePromptExport(currentFlags, registry, messageTools); err != nil || handled {
		return
	}

	if currentFlags.Workflow != "" {
		err = handleWorkflowProcessing(currentFlags, registry, messageTools)
		return
	}

	if err = handleChatProcessing(currentFlags, registry, messageTools); err != nil && currentFlags.patternFromBinary {
		err = fmt.Errorf("%w\n"+i18n.T("pattern_from_binary_name_hint"), err, currentFlags.Pattern)
	}
	return
}

func processYoutubeVideo(
	flags *Flags, registry *core.PluginRegistry, videoId string) (message string, err error) {

	if (!flags.YouTubeComments && !flags.YouTubeMetadata && !flags.YouTubeVisual) || flags.YouTubeTranscript || flags.YouTubeTranscriptWithTimestamps {
		var transcript string
		var language = "en"
		if flags.Language != "" || registry.Language.DefaultLanguage.Value != "" {
			if flags.Language != "" {
				language = flags.Language
			} else {
				language = registry.Language.DefaultLanguage.Value
			}
		}
		if flags.YouTubeTranscriptWithTimestamps {
			if transcript, err = registry.YouTube.GrabTranscriptWithTimestampsWithArgs(videoId, language, flags.YtDlpArgs); err != nil {
				return
			}
		} else {
			if transcript, err = registry.YouTube.GrabTranscriptWithArgs(videoId, language, flags.YtDlpArgs); err != nil {
				return
			}
		}
		message = AppendMessage(message, transcript)
	}

	if flags.YouTubeVisual {
		var visualText string
		var language = "en"
		if flags.Language != "" {
			language = flags.Language
		} else if registry.Language.DefaultLanguage.Value != "" {
			language = registry.Language.DefaultLanguage.Value
		}
		if visualText, err = registry.YouTube.GrabVisual(videoId, language, flags.YtDlpArgs, flags.YouTubeVisualSensitivity, flags.YouTubeVisualFps); err != nil {
			return
		}
		message = AppendMessage(message, visualText)
	}

	if flags.YouTubeComments {
		var comments []string
		if comments, err = registry.YouTube.GrabComments(videoId); err != nil {
			return
		}

		commentsString := strings.Join(comments, "\n")

		message = AppendMessage(message, commentsString)
	}

	if flags.YouTubeMetadata {
		var metadata *youtube.VideoMetadata
		if metadata, err = registry.YouTube.GrabMetadata(videoId); err != nil {
			return
		}
		metadataJson, _ := json.MarshalIndent(metadata, "", "  ")
		message = AppendMessage(message, string(metadataJson))
	}

	return
}

func WriteOutput(message string, outputFile string) (err error) {
	fmt.Println(message)
	if outputFile != "" {
		err = CreateOutputFile(message, outputFile)
	}
	return
}

func configureOpenAIResponsesAPI(registry *core.PluginRegistry, disableResponsesAPI bool) {
	if registry != nil && registry.VendorsAll != nil {
		for _, vendor := range registry.VendorsAll.Vendors {
			if vendor.GetName() == "OpenAI" {
				if openaiClient, ok := vendor.(*openai.Client); ok {
					enableResponsesAPI := !disableResponsesAPI
					openaiClient.SetResponsesAPIEnabled(enableResponsesAPI)
				}
				break
			}
		}
	}
}
