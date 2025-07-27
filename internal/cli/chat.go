package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// handleChatProcessing handles the main chat processing logic
func handleChatProcessing(currentFlags *Flags, registry *core.PluginRegistry, messageTools string) (err error) {
	if messageTools != "" {
		currentFlags.AppendMessage(messageTools)
	}

	var chatter *core.Chatter
	if chatter, err = registry.GetChatter(currentFlags.Model, currentFlags.ModelContextLength,
		currentFlags.Strategy, currentFlags.Stream, currentFlags.DryRun); err != nil {
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
		err = fmt.Errorf("TTS model '%s' requires audio output. Please specify an audio output file with -o flag (e.g., -o output.wav)", currentFlags.Model)
		return
	}

	if isAudioOutput && !isTTSModel {
		err = fmt.Errorf("audio output file '%s' specified but model '%s' is not a TTS model. Please use a TTS model like gemini-2.5-flash-preview-tts", currentFlags.Output, currentFlags.Model)
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
			err = fmt.Errorf("file %s already exists. Please choose a different filename or remove the existing file", outputFile)
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
			fmt.Printf("TTS audio generated successfully and saved to: %s\n", currentFlags.Output)
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
	return
}

// isTTSModel checks if the model is a text-to-speech model
func isTTSModel(modelName string) bool {
	lowerModel := strings.ToLower(modelName)
	return strings.Contains(lowerModel, "tts") ||
		strings.Contains(lowerModel, "preview-tts") ||
		strings.Contains(lowerModel, "text-to-speech")
}
