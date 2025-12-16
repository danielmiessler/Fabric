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
		return err
	}

	// initialize internationalization using requested language
	if _, err = i18n.Init(currentFlags.Language); err != nil {
		return err
	}

	if currentFlags.Setup {
		if err = ensureEnvFile(); err != nil {
			return err
		}
	}

	if currentFlags.Version {
		fmt.Println(version)
		return err
	}

	// Initialize database and registry
	registry, err2 := initializeFabric()
	if err2 != nil {
		if !currentFlags.Setup {
			debuglog.Log("%s\n", err2.Error())
			currentFlags.Setup = true
		}
		// Return early if registry is nil to prevent panics in subsequent handlers
		if registry == nil {
			return err2
		}
	}

	// Configure OpenAI Responses API setting based on CLI flag
	if registry != nil {
		configureOpenAIResponsesAPI(registry, currentFlags.DisableResponsesAPI)
	}

	// Handle setup and server commands
	var handled bool
	if handled, err = handleSetupAndServerCommands(currentFlags, registry, version); err != nil || handled {
		return err
	}

	// Handle configuration commands
	if handled, err = handleConfigurationCommands(currentFlags, registry); err != nil || handled {
		return err
	}

	// Handle listing commands
	if handled, err = handleListingCommands(currentFlags, registry.Db, registry); err != nil || handled {
		return err
	}

	// Handle management commands
	if handled, err = handleManagementCommands(currentFlags, registry.Db); err != nil || handled {
		return err
	}

	// Handle extension commands
	if handled, err = handleExtensionCommands(currentFlags, registry); err != nil || handled {
		return err
	}

	// Handle transcription if specified
	if currentFlags.TranscribeFile != "" {
		var transcriptionMessage string
		if transcriptionMessage, err = handleTranscription(currentFlags, registry); err != nil {
			return err
		}
		currentFlags.Message = AppendMessage(currentFlags.Message, transcriptionMessage)
	}

	// Process HTML readability if needed
	if currentFlags.HtmlReadability {
		if msg, cleanErr := converter.HtmlReadability(currentFlags.Message); cleanErr != nil {
			fmt.Println(i18n.T("html_readability_error"), cleanErr)
		} else {
			currentFlags.Message = msg
		}
	}

	// Handle tool-based message processing
	var messageTools string
	if messageTools, err = handleToolProcessing(currentFlags, registry); err != nil {
		return err
	}

	// Return early for non-chat tool operations
	if messageTools != "" && !currentFlags.IsChatRequest() {
		return nil
	}

	// Handle chat processing
	err = handleChatProcessing(currentFlags, registry, messageTools)
	return err
}

func processYoutubeVideo(
	flags *Flags, registry *core.PluginRegistry, videoId string,
) (message string, err error) {
	if (!flags.YouTubeComments && !flags.YouTubeMetadata) || flags.YouTubeTranscript || flags.YouTubeTranscriptWithTimestamps {
		var transcript string
		language := "en"
		if flags.Language != "" || registry.Language.DefaultLanguage.Value != "" {
			if flags.Language != "" {
				language = flags.Language
			} else {
				language = registry.Language.DefaultLanguage.Value
			}
		}
		if flags.YouTubeTranscriptWithTimestamps {
			if transcript, err = registry.YouTube.GrabTranscriptWithTimestampsWithArgs(videoId, language, flags.YtDlpArgs); err != nil {
				return message, err
			}
		} else {
			if transcript, err = registry.YouTube.GrabTranscriptWithArgs(videoId, language, flags.YtDlpArgs); err != nil {
				return message, err
			}
		}
		message = AppendMessage(message, transcript)
	}

	if flags.YouTubeComments {
		var comments []string
		if comments, err = registry.YouTube.GrabComments(videoId); err != nil {
			return message, err
		}

		commentsString := strings.Join(comments, "\n")

		message = AppendMessage(message, commentsString)
	}

	if flags.YouTubeMetadata {
		var metadata *youtube.VideoMetadata
		if metadata, err = registry.YouTube.GrabMetadata(videoId); err != nil {
			return message, err
		}
		metadataJson, _ := json.MarshalIndent(metadata, "", "  ")
		message = AppendMessage(message, string(metadataJson))
	}

	return message, err
}

func WriteOutput(message, outputFile string) (err error) {
	fmt.Println(message)
	if outputFile != "" {
		err = CreateOutputFile(message, outputFile)
	}
	return err
}

// configureOpenAIResponsesAPI configures the OpenAI client's Responses API setting based on the CLI flag
func configureOpenAIResponsesAPI(registry *core.PluginRegistry, disableResponsesAPI bool) {
	// Find the OpenAI vendor in the registry
	if registry != nil && registry.VendorsAll != nil {
		for _, vendor := range registry.VendorsAll.Vendors {
			if vendor.GetName() == "OpenAI" {
				// Type assertion to access the OpenAI-specific method
				if openaiClient, ok := vendor.(*openai.Client); ok {
					// Invert the disable flag to get the enable flag
					enableResponsesAPI := !disableResponsesAPI
					openaiClient.SetResponsesAPIEnabled(enableResponsesAPI)
				}
				break
			}
		}
	}
}
