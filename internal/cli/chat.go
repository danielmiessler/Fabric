package cli

import (
	"context"
	"errors"
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

func handleChatProcessing(currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (err error) {
	if messageTools != "" {
		currentFlags.AppendMessage(messageTools)
	}
	// FABRIC_MODEL_<PATTERN> sets a "vendor|model" pair or a model name for one pattern.
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
		currentFlags.Vendor, currentFlags.Stream, currentFlags.DryRun); err != nil {
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

	// Check for an existing output file before the TTS call. A failed check must not cost a model call.
	if isTTSModel && isAudioOutput {
		outputFile := currentFlags.Output
		// CreateAudioOutputFile also adds .wav, so check the same name.
		if filepath.Ext(outputFile) == "" {
			outputFile += ".wav"
		}
		if _, err = os.Stat(outputFile); err == nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("file_already_exists_choose_different"), outputFile))
			return
		}
	}

	chatOptions.AudioOutput = isAudioOutput
	if isAudioOutput {
		chatOptions.AudioFormat = "wav"
	}

	if session, err = chatter.Send(context.Background(), chatReq, chatOptions); err != nil {
		return
	}

	result := session.GetLastMessage().Content

	// Keep only the requested fenced code block. The session keeps the full response.
	if (currentFlags.Extract || currentFlags.ExtractLast) && !(isTTSModel && isAudioOutput) {
		if code, found := domain.ExtractFencedCodeBlock(result, currentFlags.ExtractLast); found {
			result = code
		}
	}

	// Print the result unless streaming already printed it.
	if !currentFlags.Stream || currentFlags.SuppressThink || chatOptions.BufferStream {
		// Do not print raw audio data.
		if isTTSModel && isAudioOutput && strings.HasPrefix(result, "FABRIC_AUDIO_DATA:") {
			fmt.Printf(i18n.T("tts_audio_generated_successfully"), currentFlags.Output)
		} else {
			fmt.Println(result)
		}
	}

	if currentFlags.Copy {
		if err = CopyToClipboard(result); err != nil {
			return
		}
	}

	if currentFlags.Output != "" {
		if currentFlags.OutputSession {
			sessionAsString := session.String()
			err = CreateOutputFile(sessionAsString, currentFlags.Output)
		} else {
			if isTTSModel && isAudioOutput {
				if strings.HasPrefix(result, "FABRIC_AUDIO_DATA:") {
					audioData := result[len("FABRIC_AUDIO_DATA:"):]
					err = CreateAudioOutputFile([]byte(audioData), currentFlags.Output)
				} else {
					// A TTS response without the prefix is an error message. Write it as text.
					err = CreateOutputFile(result, currentFlags.Output)
				}
			} else {
				err = CreateOutputFile(result, currentFlags.Output)
			}
		}
	}

	if chatOptions.Notification {
		if err = sendNotification(chatOptions, chatReq.PatternName, result); err != nil {
			// A notification failure does not fail the command.
			debuglog.Log("Failed to send notification: %v\n", err)
		}
	}

	return
}

// sendNotification runs the custom notification command, or the built-in notifier when none is set.
func sendNotification(options *domain.ChatOptions, patternName, result string) error {
	title := i18n.T("fabric_command_complete")
	if patternName != "" {
		title = fmt.Sprintf(i18n.T("fabric_command_complete_with_pattern"), patternName)
	}

	// Truncate at 100 code points, not grapheme clusters. A multi-code-point emoji can be cut in the middle.
	message := i18n.T("command_completed_successfully")
	if result != "" {
		maxLength := 100
		runes := []rune(result)
		if len(runes) > maxLength {
			message = fmt.Sprintf(i18n.T("output_truncated"), string(runes[:maxLength]))
		} else {
			message = fmt.Sprintf(i18n.T("output_full"), result)
		}
		// Notifications show one line.
		message = strings.ReplaceAll(message, "\n", " ")
	}

	if options.NotificationCommand != "" {
		// SECURITY: pass title and message as the positional arguments $1 and $2, never inside the command string.
		// docs/Desktop-Notifications.md documents this interface.
		cmd := exec.Command("sh", "-c", options.NotificationCommand+" \"$1\" \"$2\"", "--", title, message)

		// Show the command output on the terminal.
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}

	notificationManager := notifications.NewNotificationManager()
	if !notificationManager.IsAvailable() {
		return errors.New(i18n.T("no_notification_system_available"))
	}

	return notificationManager.Send(title, message)
}

// isTTSModel matches model names that contain "tts" or "text-to-speech".
func isTTSModel(modelName string) bool {
	lowerModel := strings.ToLower(modelName)
	return strings.Contains(lowerModel, "tts") ||
		strings.Contains(lowerModel, "preview-tts") ||
		strings.Contains(lowerModel, "text-to-speech")
}
