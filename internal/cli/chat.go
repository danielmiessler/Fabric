package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools/notifications"
)

// handleChatProcessing handles the main chat processing logic
func handleChatProcessing(currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (err error) {
	if messageTools != "" {
		currentFlags.AppendMessage(messageTools)
	}
	// Check for pattern-specific model via environment variable
	if currentFlags.Pattern != "" && currentFlags.Model == "" {
		envVar := "FABRIC_MODEL_" + strings.ToUpper(strings.ReplaceAll(currentFlags.Pattern, "-", "_"))
		if modelSpec := os.Getenv(envVar); modelSpec != "" {
			parts := strings.SplitN(modelSpec, "|", 2)
			if len(parts) == 2 {
				currentFlags.Vendor = parts[0]
				currentFlags.Model = parts[1]
			} else {
				currentFlags.Model = modelSpec
			}
		}
	}

	var chatter *core.Chatter
	if chatter, err = registry.GetChatter(currentFlags.Model, currentFlags.ModelContextLength,
		currentFlags.Vendor, currentFlags.Strategy, currentFlags.Stream, currentFlags.DryRun); err != nil {
		return
	}

	var session *fsdb.Session
	var chatReq *domain.ChatRequest
	if chatReq, err = currentFlags.BuildChatRequest(strings.Join(os.Args[1:], " ")); err != nil {
		return
	}

	if chatReq.Language == "" {
		chatReq.Language = registry.Language.DefaultLanguage.Value
	}
	var chatOptions *domain.ChatOptions
	if chatOptions, err = currentFlags.BuildChatOptions(); err != nil {
		return
	}

	// Check if user is requesting audio output or using a TTS model
	isAudioOutput := currentFlags.Output != "" && IsAudioFormat(currentFlags.Output)
	isTTSModel := isTTSModel(currentFlags.Model)

	if isTTSModel && !isAudioOutput {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("tts_model_requires_audio_output"), currentFlags.Model))
		return
	}

	if isAudioOutput && !isTTSModel {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("audio_output_file_specified_but_not_tts_model"), currentFlags.Output, currentFlags.Model))
		return
	}

	// For TTS models, check if output file already exists BEFORE processing
	if isTTSModel && isAudioOutput {
		outputFile := currentFlags.Output
		// Add .wav extension if not provided
		if filepath.Ext(outputFile) == "" {
			outputFile += ".wav"
		}
		if _, err = os.Stat(outputFile); err == nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("file_already_exists_choose_different"), outputFile))
			return
		}
	}

	// Set audio options in chat config
	chatOptions.AudioOutput = isAudioOutput
	if isAudioOutput {
		chatOptions.AudioFormat = "wav" // Default to WAV format
	}

	if session, err = chatter.Send(chatReq, chatOptions); err != nil {
		return
	}

	result := session.GetLastMessage().Content

	if !currentFlags.Stream || currentFlags.SuppressThink {
		// For TTS models with audio output, show a user-friendly message instead of raw data
		if isTTSModel && isAudioOutput && strings.HasPrefix(result, "FABRIC_AUDIO_DATA:") {
			fmt.Printf(i18n.T("tts_audio_generated_successfully"), currentFlags.Output)
		} else {
			// print the result if it was not streamed already or suppress-think disabled streaming output
			fmt.Println(result)
		}
	}

	// if the copy flag is set, copy the message to the clipboard
	if currentFlags.Copy {
		if err = CopyToClipboard(result); err != nil {
			return
		}
	}

	// if the output flag is set, create an output file
	if currentFlags.Output != "" {
		if currentFlags.OutputSession {
			sessionAsString := session.String()
			err = CreateOutputFile(sessionAsString, currentFlags.Output)
		} else {
			// For TTS models, we need to handle audio output differently
			if isTTSModel && isAudioOutput {
				// Check if result contains actual audio data
				if strings.HasPrefix(result, "FABRIC_AUDIO_DATA:") {
					// Extract the binary audio data
					audioData := result[len("FABRIC_AUDIO_DATA:"):]
					err = CreateAudioOutputFile([]byte(audioData), currentFlags.Output)
				} else {
					// Fallback for any error messages or unexpected responses
					err = CreateOutputFile(result, currentFlags.Output)
				}
			} else {
				err = CreateOutputFile(result, currentFlags.Output)
			}
		}
	}

	// Send notification if requested
	if chatOptions.Notification {
		if err = sendNotification(chatOptions, chatReq.PatternName, result); err != nil {
			// Log notification error but don't fail the main command
			debuglog.Log("Failed to send notification: %v\n", err)
		}
	}

	return
}

// sendNotification sends a desktop notification about command completion.
//
// When truncating the result for notification display, this function counts Unicode code points,
// not grapheme clusters. As a result, complex emoji or accented characters with multiple combining
// characters may be truncated improperly. This is a limitation of the current implementation.
func sendNotification(options *domain.ChatOptions, patternName, result string) error {
	title := i18n.T("fabric_command_complete")
	if patternName != "" {
		title = fmt.Sprintf(i18n.T("fabric_command_complete_with_pattern"), patternName)
	}

	// Limit message length for notification display (counts Unicode code points)
	message := i18n.T("command_completed_successfully")
	if result != "" {
		maxLength := 100
		runes := []rune(result)
		if len(runes) > maxLength {
			message = fmt.Sprintf(i18n.T("output_truncated"), string(runes[:maxLength]))
		} else {
			message = fmt.Sprintf(i18n.T("output_full"), result)
		}
		// Clean up newlines for notification display
		message = strings.ReplaceAll(message, "\n", " ")
	}

	// Use custom notification command if provided
	if options.NotificationCommand != "" {
		// SECURITY: Pass title and message as proper shell positional arguments $1 and $2
		// This matches the documented interface where custom commands receive title and message as shell variables
		cmd := exec.Command("sh", "-c", options.NotificationCommand+" \"$1\" \"$2\"", "--", title, message)

		// For debugging: capture and display output from custom commands
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}

	// Use built-in notification system
	notificationManager := notifications.NewNotificationManager()
	if !notificationManager.IsAvailable() {
		return fmt.Errorf("%s", i18n.T("no_notification_system_available"))
	}

	return notificationManager.Send(title, message)
}

// isTTSModel checks if the model is a text-to-speech model
func isTTSModel(modelName string) bool {
	lowerModel := strings.ToLower(modelName)
	return strings.Contains(lowerModel, "tts") ||
		strings.Contains(lowerModel, "preview-tts") ||
		strings.Contains(lowerModel, "text-to-speech")
}
