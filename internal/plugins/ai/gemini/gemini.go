package gemini

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/geminicommon"
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

const (
	errInvalidLocationFormatKey = "gemini_invalid_location_format"
	locationSeparator           = "/"
	langCodeSeparator           = "_"
	langCodeNormalizedSep       = "-"

	modelPrefix           = "models/"
	modelTypeTTS          = "tts"
	modelTypePreviewTTS   = "preview-tts"
	modelTypeTextToSpeech = "text-to-speech"
)

var langCodeRegex = regexp.MustCompile(`^[a-z]{2}(-[A-Z]{2})?$`)

func NewClient() (ret *Client) {
	vendorName := "Gemini"
	ret = &Client{}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, nil)

	ret.ApiKey = ret.PluginBase.AddSetupQuestion("API key", true)

	return
}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func (o *Client) ListModels(_ context.Context) (ret []string, err error) {
	ctx := context.Background()
	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	resp, err := client.Models.List(ctx, &genai.ListModelsConfig{})
	if err != nil {
		return nil, err
	}

	for _, model := range resp.Items {
		modelName := strings.TrimPrefix(model.Name, "models/")
		ret = append(ret, modelName)
	}
	return
}

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	if o.isTTSModel(opts.Model) {
		if !opts.AudioOutput {
			err = fmt.Errorf(i18n.T("tts_model_requires_audio_output"), opts.Model)
			return
		}

		return o.generateTTSAudio(ctx, msgs, opts)
	}

	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	contents := geminicommon.ConvertMessages(msgs)

	cfg, err := o.buildGenerateContentConfig(opts)
	if err != nil {
		return "", err
	}

	response, err := client.Models.GenerateContent(ctx, o.buildModelNameFull(opts.Model), contents, cfg)
	if err != nil {
		return "", err
	}

	ret = geminicommon.ExtractTextWithCitations(response)
	return
}

func (o *Client) SendStream(_ context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) (err error) {
	ctx := context.Background()
	defer close(channel)

	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return
	}

	contents := geminicommon.ConvertMessages(msgs)

	cfg, err := o.buildGenerateContentConfig(opts)
	if err != nil {
		return err
	}

	stream := client.Models.GenerateContentStream(ctx, o.buildModelNameFull(opts.Model), contents, cfg)

	for response, err := range stream {
		if err != nil {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeError,
				Content: fmt.Sprintf(i18n.T("gemini_stream_error"), err),
			}
			return err
		}

		text := geminicommon.ExtractTextWithCitations(response)
		if text != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: text,
			}
		}

		if response.UsageMetadata != nil {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(response.UsageMetadata.PromptTokenCount),
					OutputTokens: int(response.UsageMetadata.CandidatesTokenCount),
					TotalTokens:  int(response.UsageMetadata.TotalTokenCount),
				},
			}
		}
	}

	return
}

func parseThinkingConfig(level domain.ThinkingLevel) (*genai.ThinkingConfig, bool) {
	lower := strings.ToLower(strings.TrimSpace(string(level)))
	switch domain.ThinkingLevel(lower) {
	case "", domain.ThinkingOff:
		return nil, false
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		if budget, ok := domain.ThinkingBudgets[domain.ThinkingLevel(lower)]; ok {
			b := int32(budget)
			return &genai.ThinkingConfig{IncludeThoughts: true, ThinkingBudget: &b}, true
		}
	default:
		if tokens, err := strconv.ParseInt(lower, 10, 32); err == nil && tokens > 0 {
			t := int32(tokens)
			return &genai.ThinkingConfig{IncludeThoughts: true, ThinkingBudget: &t}, true
		}
	}
	return nil, false
}

// buildGenerateContentConfig builds the request config. When opts.Search is
// set, it adds the Google Search tool. The search location accepts two forms:
//   - "Continent/City", for example "America/Los_Angeles"
//   - an ISO language code "ll" or "ll-CC", for example "en" or "en-US"
//
// In a language code, the first underscore becomes a hyphen. An invalid
// location returns an error.
func (o *Client) buildGenerateContentConfig(opts *domain.ChatOptions) (*genai.GenerateContentConfig, error) {
	temperature := float32(opts.Temperature)
	topP := float32(opts.TopP)
	var maxTokens int32
	if opts.MaxTokens > 0 {
		maxTokens = int32(opts.MaxTokens)
	}
	cfg := &genai.GenerateContentConfig{
		Temperature:     &temperature,
		TopP:            &topP,
		MaxOutputTokens: maxTokens,
	}

	if opts.Search {
		cfg.Tools = []*genai.Tool{{GoogleSearch: &genai.GoogleSearch{}}}
		if loc := opts.SearchLocation; loc != "" {
			if isValidLocationFormat(loc) {
				loc = normalizeLocation(loc)
				cfg.ToolConfig = &genai.ToolConfig{
					RetrievalConfig: &genai.RetrievalConfig{LanguageCode: loc},
				}
			} else {
				return nil, fmt.Errorf(i18n.T(errInvalidLocationFormatKey), loc)
			}
		}
	}

	if tc, ok := parseThinkingConfig(opts.Thinking); ok {
		cfg.ThinkingConfig = tc
	}

	return cfg, nil
}

// buildModelNameFull adds the "models/" prefix when the name lacks it.
func (o *Client) buildModelNameFull(modelName string) string {
	if strings.HasPrefix(modelName, modelPrefix) {
		return modelName
	}
	return modelPrefix + modelName
}

func isValidLocationFormat(location string) bool {
	if strings.Contains(location, locationSeparator) {
		parts := strings.Split(location, locationSeparator)
		return len(parts) == 2 && parts[0] != "" && parts[1] != ""
	}
	return isValidLanguageCode(location)
}

func normalizeLocation(location string) string {
	if strings.Contains(location, locationSeparator) {
		return location
	}
	return strings.Replace(location, langCodeSeparator, langCodeNormalizedSep, 1)
}

// isValidLanguageCode reports whether code is an ISO 639-1 language code with
// an optional ISO 3166-1 country code. It replaces the first underscore with a
// hyphen and ignores case.
func isValidLanguageCode(code string) bool {
	normalized := strings.Replace(code, langCodeSeparator, langCodeNormalizedSep, 1)
	parts := strings.Split(normalized, langCodeNormalizedSep)
	switch len(parts) {
	case 1:
		return langCodeRegex.MatchString(strings.ToLower(parts[0]))
	case 2:
		formatted := strings.ToLower(parts[0]) + langCodeNormalizedSep + strings.ToUpper(parts[1])
		return langCodeRegex.MatchString(formatted)
	default:
		return false
	}
}

func (o *Client) isTTSModel(modelName string) bool {
	lowerModel := strings.ToLower(modelName)
	return strings.Contains(lowerModel, modelTypeTTS) ||
		strings.Contains(lowerModel, modelTypePreviewTTS) ||
		strings.Contains(lowerModel, modelTypeTextToSpeech)
}

// extractTextForTTS returns the content of the last non-empty user message.
func (o *Client) extractTextForTTS(msgs []*chat.ChatCompletionMessage) (string, error) {
	for _, msg := range slices.Backward(msgs) {
		if msg.Role == chat.ChatMessageRoleUser && msg.Content != "" {
			return msg.Content, nil
		}
	}
	return "", errors.New(i18n.T("gemini_no_text_for_tts"))
}

func (o *Client) createGenaiClient(ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  o.ApiKey.Value,
		Backend: genai.BackendGeminiAPI,
	})
}

func (o *Client) generateTTSAudio(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	textToSpeak, err := o.extractTextForTTS(msgs)
	if err != nil {
		return "", err
	}

	if opts.Voice != "" && !IsValidGeminiVoice(opts.Voice) {
		validVoices := GetGeminiVoiceNames()
		return "", fmt.Errorf(i18n.T("gemini_invalid_voice"), opts.Voice, validVoices)
	}

	client, err := o.createGenaiClient(ctx)
	if err != nil {
		return "", err
	}

	return o.performTTSGeneration(ctx, client, textToSpeak, opts)
}

func (o *Client) performTTSGeneration(ctx context.Context, client *genai.Client, textToSpeak string, opts *domain.ChatOptions) (string, error) {
	contents := []*genai.Content{{
		Parts: []*genai.Part{{Text: textToSpeak}},
	}}

	voiceName := opts.Voice
	if voiceName == "" {
		voiceName = "Kore"
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

	response, err := client.Models.GenerateContent(ctx, o.buildModelNameFull(opts.Model), contents, config)
	if err != nil {
		return "", fmt.Errorf(i18n.T("gemini_tts_failed"), err)
	}

	if len(response.Candidates) > 0 && response.Candidates[0].Content != nil && len(response.Candidates[0].Content.Parts) > 0 {
		part := response.Candidates[0].Content.Parts[0]
		if part.InlineData != nil && len(part.InlineData.Data) > 0 {
			if part.InlineData.MIMEType != "" && !strings.HasPrefix(part.InlineData.MIMEType, "audio/") {
				return "", fmt.Errorf(i18n.T("gemini_unexpected_data_type"), part.InlineData.MIMEType)
			}

			pcmData := part.InlineData.Data
			if len(pcmData) < MinAudioDataSize {
				return "", fmt.Errorf(i18n.T("gemini_audio_data_too_small"), len(pcmData), MinAudioDataSize)
			}

			wavData, err := o.generateWAVFile(pcmData)
			if err != nil {
				return "", fmt.Errorf(i18n.T("gemini_wav_generation_failed"), err)
			}

			if len(wavData) < WAVHeaderSize {
				return "", fmt.Errorf(i18n.T("gemini_wav_data_invalid"), len(wavData), WAVHeaderSize)
			}

			// The CLI detects this prefix and writes the bytes after it to the output file.
			return AudioDataPrefix + string(wavData), nil
		}
	}

	return "", errors.New(i18n.T("gemini_no_audio_data"))
}

// generateWAVFile adds a WAV header to pcmData.
func (o *Client) generateWAVFile(pcmData []byte) ([]byte, error) {
	if len(pcmData) == 0 {
		return nil, errors.New(i18n.T("gemini_empty_pcm_data"))
	}
	if len(pcmData) > MaxAudioDataSize {
		return nil, fmt.Errorf(i18n.T("gemini_pcm_data_too_large"), len(pcmData), MaxAudioDataSize)
	}

	// Gemini TTS returns 16-bit mono PCM at 24 kHz.
	channels := DefaultChannels
	sampleRate := DefaultSampleRate
	bitsPerSample := DefaultBitsPerSample

	byteRate := sampleRate * channels * bitsPerSample / 8
	blockAlign := channels * bitsPerSample / 8
	dataLen := uint32(len(pcmData))
	riffSize := RIFFHeaderSize + dataLen

	totalSize := int(riffSize + 8) // riffSize excludes the "RIFF" tag and the size field
	buf := bytes.NewBuffer(make([]byte, 0, totalSize))

	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, riffSize)
	buf.WriteString("WAVE")

	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16)) // fmt chunk size
	binary.Write(buf, binary.LittleEndian, uint16(1))  // PCM
	binary.Write(buf, binary.LittleEndian, uint16(channels))
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))
	binary.Write(buf, binary.LittleEndian, uint32(byteRate))
	binary.Write(buf, binary.LittleEndian, uint16(blockAlign))
	binary.Write(buf, binary.LittleEndian, uint16(bitsPerSample))

	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, dataLen)

	buf.Write(pcmData)

	result := buf.Bytes()
	if len(result) < WAVHeaderSize {
		return nil, fmt.Errorf(i18n.T("gemini_wav_data_invalid"), len(result), WAVHeaderSize)
	}

	return result, nil
}
