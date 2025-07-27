package gemini

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/plugins"

	"github.com/danielmiessler/fabric/internal/domain"
	"google.golang.org/genai"
)

// WAV audio constants
const (
	DefaultChannels      = 1
	DefaultSampleRate    = 24000
	DefaultBitsPerSample = 16
	WAVHeaderSize        = 44
	RIFFHeaderSize       = 36
	MaxAudioDataSize     = 100 * 1024 * 1024 // 100MB limit for security
	MinAudioDataSize     = 44                // Minimum viable audio data
	AudioDataPrefix      = "FABRIC_AUDIO_DATA:"
)

func NewClient() (ret *Client) {
	vendorName := "Gemini"
	ret = &Client{}

	ret.PluginBase = &plugins.PluginBase{
		Name:          vendorName,
		EnvNamePrefix: plugins.BuildEnvVariablePrefix(vendorName),
	}

	ret.ApiKey = ret.PluginBase.AddSetupQuestion("API key", true)

	return
}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func (o *Client) ListModels() (ret []string, err error) {
	ctx := context.Background()
	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	// List available models using the correct API
	resp, err := client.Models.List(ctx, &genai.ListModelsConfig{})
	if err != nil {
		return nil, err
	}

	for _, model := range resp.Items {
		// Strip the "models/" prefix for user convenience
		modelName := strings.TrimPrefix(model.Name, "models/")
		ret = append(ret, modelName)
	}
	return
}

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	// Check if this is a TTS model request
	if o.isTTSModel(opts.Model) {
		if !opts.AudioOutput {
			err = fmt.Errorf("TTS model '%s' requires audio output. Please specify an audio output file with -o flag ending in .wav", opts.Model)
			return
		}

		// Handle TTS generation
		return o.generateTTSAudio(ctx, msgs, opts)
	}

	// Regular text generation
	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	// Convert messages to new SDK format
	contents := o.convertMessages(msgs)

	// Generate content
	temperature := float32(opts.Temperature)
	topP := float32(opts.TopP)
	response, err := client.Models.GenerateContent(ctx, o.buildModelNameFull(opts.Model), contents, &genai.GenerateContentConfig{
		Temperature:     &temperature,
		TopP:            &topP,
		MaxOutputTokens: int32(opts.ModelContextLength),
	})
	if err != nil {
		return "", err
	}

	// Extract text from response
	ret = o.extractTextFromResponse(response)
	return
}

func (o *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string) (err error) {
	ctx := context.Background()
	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	// Convert messages to new SDK format
	contents := o.convertMessages(msgs)

	// Generate streaming content
	temperature := float32(opts.Temperature)
	topP := float32(opts.TopP)
	stream := client.Models.GenerateContentStream(ctx, o.buildModelNameFull(opts.Model), contents, &genai.GenerateContentConfig{
		Temperature:     &temperature,
		TopP:            &topP,
		MaxOutputTokens: int32(opts.ModelContextLength),
	})

	for response, err := range stream {
		if err != nil {
			channel <- fmt.Sprintf("Error: %v\n", err)
			close(channel)
			break
		}

		text := o.extractTextFromResponse(response)
		if text != "" {
			channel <- text
		}
	}
	close(channel)

	return
}

func (o *Client) NeedsRawMode(modelName string) bool {
	return false
}

// buildModelNameFull adds the "models/" prefix for API calls
func (o *Client) buildModelNameFull(modelName string) string {
	if strings.HasPrefix(modelName, "models/") {
		return modelName
	}
	return "models/" + modelName
}

// isTTSModel checks if the model is a text-to-speech model
func (o *Client) isTTSModel(modelName string) bool {
	lowerModel := strings.ToLower(modelName)
	return strings.Contains(lowerModel, "tts") ||
		strings.Contains(lowerModel, "preview-tts") ||
		strings.Contains(lowerModel, "text-to-speech")
}

// extractTextForTTS extracts text content from chat messages for TTS generation
func (o *Client) extractTextForTTS(msgs []*chat.ChatCompletionMessage) (string, error) {
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role == chat.ChatMessageRoleUser && msgs[i].Content != "" {
			return msgs[i].Content, nil
		}
	}
	return "", fmt.Errorf("no text content found for TTS generation")
}

// createGenaiClient creates a new GenAI client for TTS operations
func (o *Client) createGenaiClient(ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	})
}

// generateTTSAudio handles TTS audio generation using the new SDK
func (o *Client) generateTTSAudio(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	textToSpeak, err := o.extractTextForTTS(msgs)
	if err != nil {
		return "", err
	}

	// Validate voice name before making API call
	if opts.Voice != "" && !IsValidGeminiVoice(opts.Voice) {
		validVoices := GetGeminiVoiceNames()
		return "", fmt.Errorf("invalid voice '%s'. Valid voices are: %v", opts.Voice, validVoices)
	}

	client, err := o.createGenaiClient(ctx)
	if err != nil {
		return "", err
	}

	return o.performTTSGeneration(ctx, client, textToSpeak, opts)
}

// performTTSGeneration performs the actual TTS generation and audio processing
func (o *Client) performTTSGeneration(ctx context.Context, client *genai.Client, textToSpeak string, opts *domain.ChatOptions) (string, error) {

	// Create content for TTS
	contents := []*genai.Content{{
		Parts: []*genai.Part{{Text: textToSpeak}},
	}}

	// Configure for TTS generation
	voiceName := opts.Voice
	if voiceName == "" {
		voiceName = "Kore" // Default voice if none specified
	}

	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"AUDIO"},
		SpeechConfig: &genai.SpeechConfig{
			VoiceConfig: &genai.VoiceConfig{
				PrebuiltVoiceConfig: &genai.PrebuiltVoiceConfig{
					VoiceName: voiceName,
				},
			},
		},
	}

	// Generate TTS content
	response, err := client.Models.GenerateContent(ctx, o.buildModelNameFull(opts.Model), contents, config)
	if err != nil {
		return "", fmt.Errorf("TTS generation failed: %w", err)
	}

	// Extract and process audio data
	if len(response.Candidates) > 0 && response.Candidates[0].Content != nil && len(response.Candidates[0].Content.Parts) > 0 {
		part := response.Candidates[0].Content.Parts[0]
		if part.InlineData != nil && len(part.InlineData.Data) > 0 {
			// Validate audio data format and size
			if part.InlineData.MIMEType != "" && !strings.HasPrefix(part.InlineData.MIMEType, "audio/") {
				return "", fmt.Errorf("unexpected data type: %s, expected audio data", part.InlineData.MIMEType)
			}

			pcmData := part.InlineData.Data
			if len(pcmData) < MinAudioDataSize {
				return "", fmt.Errorf("audio data too small: %d bytes, minimum required: %d", len(pcmData), MinAudioDataSize)
			}

			// Generate WAV file with proper headers and return the binary data
			wavData, err := o.generateWAVFile(pcmData)
			if err != nil {
				return "", fmt.Errorf("failed to generate WAV file: %w", err)
			}

			// Validate generated WAV data
			if len(wavData) < WAVHeaderSize {
				return "", fmt.Errorf("generated WAV data is invalid: %d bytes, minimum required: %d", len(wavData), WAVHeaderSize)
			}

			// Store the binary audio data in a special format that the CLI can detect
			// Use more efficient string concatenation
			return AudioDataPrefix + string(wavData), nil
		}
	}

	return "", fmt.Errorf("no audio data received from TTS model")
}

// generateWAVFile creates WAV data from PCM data with proper headers
func (o *Client) generateWAVFile(pcmData []byte) ([]byte, error) {
	// Validate input size to prevent potential security issues
	if len(pcmData) == 0 {
		return nil, fmt.Errorf("empty PCM data provided")
	}
	if len(pcmData) > MaxAudioDataSize {
		return nil, fmt.Errorf("PCM data too large: %d bytes, maximum allowed: %d", len(pcmData), MaxAudioDataSize)
	}

	// WAV file parameters (Gemini TTS default specs)
	channels := DefaultChannels
	sampleRate := DefaultSampleRate
	bitsPerSample := DefaultBitsPerSample

	// Calculate required values
	byteRate := sampleRate * channels * bitsPerSample / 8
	blockAlign := channels * bitsPerSample / 8
	dataLen := uint32(len(pcmData))
	riffSize := RIFFHeaderSize + dataLen

	// Pre-allocate buffer with known size for better performance
	totalSize := int(riffSize + 8) // +8 for RIFF header
	buf := bytes.NewBuffer(make([]byte, 0, totalSize))

	// RIFF header
	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, riffSize)
	buf.WriteString("WAVE")

	// fmt chunk
	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16))            // subchunk1Size
	binary.Write(buf, binary.LittleEndian, uint16(1))             // audioFormat = PCM
	binary.Write(buf, binary.LittleEndian, uint16(channels))      // numChannels
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))    // sampleRate
	binary.Write(buf, binary.LittleEndian, uint32(byteRate))      // byteRate
	binary.Write(buf, binary.LittleEndian, uint16(blockAlign))    // blockAlign
	binary.Write(buf, binary.LittleEndian, uint16(bitsPerSample)) // bitsPerSample

	// data chunk
	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, dataLen)

	// Write PCM data to buffer
	buf.Write(pcmData)

	// Validate generated WAV data
	result := buf.Bytes()
	if len(result) < WAVHeaderSize {
		return nil, fmt.Errorf("generated WAV data is invalid: %d bytes, minimum required: %d", len(result), WAVHeaderSize)
	}

	return result, nil
}

// convertMessages converts fabric chat messages to genai Content format
func (o *Client) convertMessages(msgs []*chat.ChatCompletionMessage) []*genai.Content {
	var contents []*genai.Content

	for _, msg := range msgs {
		content := &genai.Content{Parts: []*genai.Part{}}

		if msg.Content != "" {
			content.Parts = append(content.Parts, &genai.Part{Text: msg.Content})
		}

		// Handle multi-content messages (images, etc.)
		for _, part := range msg.MultiContent {
			switch part.Type {
			case chat.ChatMessagePartTypeText:
				content.Parts = append(content.Parts, &genai.Part{Text: part.Text})
			case chat.ChatMessagePartTypeImageURL:
				// TODO: Handle image URLs if needed
				// This would require downloading and converting to inline data
			}
		}

		contents = append(contents, content)
	}

	return contents
}

// extractTextFromResponse extracts text content from the response
func (o *Client) extractTextFromResponse(response *genai.GenerateContentResponse) string {
	var result strings.Builder

	for _, candidate := range response.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				if part.Text != "" {
					result.WriteString(part.Text)
				}
			}
		}
	}

	return result.String()
}
