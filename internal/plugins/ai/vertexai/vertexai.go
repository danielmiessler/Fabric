package vertexai

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/vertex"
	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/geminicommon"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/genai"
)

const (
	cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
	defaultRegion      = "global"
	defaultMaxTokens   = 4096
)

// NewClient creates a new Vertex AI client for accessing Claude models via Google Cloud
func NewClient() (ret *Client) {
	vendorName := "VertexAI"
	ret = &Client{}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, ret.configure)

	ret.ProjectID = ret.AddSetupQuestion("Project ID", true)
	ret.Region = ret.AddSetupQuestion("Region", false)
	ret.Region.Value = defaultRegion

	return
}

// Client implements the ai.Vendor interface for Google Cloud Vertex AI with Anthropic models
type Client struct {
	*plugins.PluginBase
	ProjectID *plugins.SetupQuestion
	Region    *plugins.SetupQuestion

	client *anthropic.Client
}

func (c *Client) configure() error {
	ctx := context.Background()
	projectID := c.ProjectID.Value
	region := c.Region.Value

	// Initialize Anthropic client for Claude models via Vertex AI using Google ADC
	vertexOpt := vertex.WithGoogleAuth(ctx, region, projectID, cloudPlatformScope)
	client := anthropic.NewClient(vertexOpt)
	c.client = &client

	return nil
}

func (c *Client) ListModels() ([]string, error) {
	ctx := context.Background()

	// Get ADC credentials for API authentication
	creds, err := google.FindDefaultCredentials(ctx, cloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_failed_google_credentials"), err)
	}
	httpClient := oauth2.NewClient(ctx, creds.TokenSource)

	// Query all publishers in parallel for better performance
	type result struct {
		models    []string
		err       error
		publisher string
	}
	// +1 for known Gemini models (no API to list them)
	results := make(chan result, len(publishers)+1)

	// Query Model Garden API for third-party models
	for _, pub := range publishers {
		go func(publisher string) {
			models, err := listPublisherModels(ctx, httpClient, c.Region.Value, c.ProjectID.Value, publisher)
			results <- result{models: models, err: err, publisher: publisher}
		}(pub)
	}

	// Add known Gemini models (Vertex AI doesn't have a list API for Gemini)
	go func() {
		results <- result{models: getKnownGeminiModels(), err: nil, publisher: "gemini"}
	}()

	// Collect results from all sources
	var allModels []string
	for range len(publishers) + 1 {
		r := <-results
		if r.err != nil {
			// Log warning but continue - some sources may not be available
			debuglog.Debug(debuglog.Basic, "Failed to list %s models: %v\n", r.publisher, r.err)
			continue
		}
		allModels = append(allModels, r.models...)
	}

	if len(allModels) == 0 {
		return nil, errors.New(i18n.T("vertexai_no_models_found"))
	}

	// Filter to only conversational models and sort
	filtered := filterConversationalModels(allModels)
	if len(filtered) == 0 {
		return nil, errors.New(i18n.T("vertexai_no_conversational_models"))
	}

	return sortModels(filtered), nil
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if isGeminiModel(opts.Model) {
		return c.sendGemini(ctx, msgs, opts)
	}
	return c.sendClaude(ctx, msgs, opts)
}

// getMaxTokens returns the max output tokens to use for a request
func getMaxTokens(opts *domain.ChatOptions) int64 {
	if opts.MaxTokens > 0 {
		return int64(opts.MaxTokens)
	}
	return int64(defaultMaxTokens)
}

func (c *Client) sendClaude(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if c.client == nil {
		return "", errors.New(i18n.T("vertexai_client_not_initialized"))
	}

	// Convert chat messages to Anthropic format
	anthropicMessages := c.toMessages(msgs)
	if len(anthropicMessages) == 0 {
		return "", errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	// Build request params
	params := anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: getMaxTokens(opts),
		Messages:  anthropicMessages,
	}

	// Only set one of Temperature or TopP as some models don't allow both
	// (following anthropic.go pattern)
	if opts.TopP != domain.DefaultTopP {
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	response, err := c.client.Messages.New(ctx, params)
	if err != nil {
		return "", err
	}

	// Extract text from response
	var textParts []string
	for _, block := range response.Content {
		if block.Type == "text" && block.Text != "" {
			textParts = append(textParts, block.Text)
		}
	}

	if len(textParts) == 0 {
		return "", errors.New(i18n.T("vertexai_no_content_in_response"))
	}

	return strings.Join(textParts, ""), nil
}

func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	if isGeminiModel(opts.Model) {
		return c.sendStreamGemini(msgs, opts, channel)
	}
	return c.sendStreamClaude(msgs, opts, channel)
}

func (c *Client) sendStreamClaude(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	if c.client == nil {
		close(channel)
		return errors.New(i18n.T("vertexai_client_not_initialized"))
	}

	defer close(channel)
	ctx := context.Background()

	// Convert chat messages to Anthropic format
	anthropicMessages := c.toMessages(msgs)
	if len(anthropicMessages) == 0 {
		return errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	// Build request params
	params := anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: getMaxTokens(opts),
		Messages:  anthropicMessages,
	}

	// Only set one of Temperature or TopP as some models don't allow both
	if opts.TopP != domain.DefaultTopP {
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	// Create streaming request
	stream := c.client.Messages.NewStreaming(ctx, params)

	// Process stream
	for stream.Next() {
		event := stream.Current()

		// Handle Content
		if event.Delta.Text != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.Delta.Text,
			}
		}

		// Handle Usage
		if event.Message.Usage.InputTokens != 0 || event.Message.Usage.OutputTokens != 0 {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(event.Message.Usage.InputTokens),
					OutputTokens: int(event.Message.Usage.OutputTokens),
					TotalTokens:  int(event.Message.Usage.InputTokens + event.Message.Usage.OutputTokens),
				},
			}
		} else if event.Usage.InputTokens != 0 || event.Usage.OutputTokens != 0 {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(event.Usage.InputTokens),
					OutputTokens: int(event.Usage.OutputTokens),
					TotalTokens:  int(event.Usage.InputTokens + event.Usage.OutputTokens),
				},
			}
		}
	}

	return stream.Err()
}

// Gemini methods using genai SDK with Vertex AI backend

// getGeminiRegion returns the appropriate region for a Gemini model.
// Preview models are often only available on the global endpoint.
func (c *Client) getGeminiRegion(model string) string {
	if strings.Contains(strings.ToLower(model), "preview") {
		return "global"
	}
	return c.Region.Value
}

func (c *Client) sendGemini(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  c.ProjectID.Value,
		Location: c.getGeminiRegion(opts.Model),
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return "", fmt.Errorf(i18n.T("vertexai_failed_gemini_client"), err)
	}

	contents := geminicommon.ConvertMessages(msgs)
	if len(contents) == 0 {
		return "", errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	config := c.buildGeminiConfig(opts)

	response, err := client.Models.GenerateContent(ctx, opts.Model, contents, config)
	if err != nil {
		return "", err
	}

	return geminicommon.ExtractTextWithCitations(response), nil
}

// buildGeminiConfig creates the generation config for Gemini models
// following the gemini.go pattern for feature parity
func (c *Client) buildGeminiConfig(opts *domain.ChatOptions) *genai.GenerateContentConfig {
	temperature := float32(opts.Temperature)
	topP := float32(opts.TopP)
	config := &genai.GenerateContentConfig{
		Temperature:     &temperature,
		TopP:            &topP,
		MaxOutputTokens: int32(getMaxTokens(opts)),
	}

	// Add web search support
	if opts.Search {
		config.Tools = []*genai.Tool{{GoogleSearch: &genai.GoogleSearch{}}}
	}

	// Add thinking support
	if tc := parseGeminiThinking(opts.Thinking); tc != nil {
		config.ThinkingConfig = tc
	}

	return config
}

// parseGeminiThinking converts thinking level to Gemini thinking config
func parseGeminiThinking(level domain.ThinkingLevel) *genai.ThinkingConfig {
	lower := strings.ToLower(strings.TrimSpace(string(level)))
	switch domain.ThinkingLevel(lower) {
	case "", domain.ThinkingOff:
		return nil
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		if budget, ok := domain.ThinkingBudgets[domain.ThinkingLevel(lower)]; ok {
			b := int32(budget)
			return &genai.ThinkingConfig{IncludeThoughts: true, ThinkingBudget: &b}
		}
	default:
		// Try parsing as integer token count
		var tokens int
		if _, err := fmt.Sscanf(lower, "%d", &tokens); err == nil && tokens > 0 {
			t := int32(tokens)
			return &genai.ThinkingConfig{IncludeThoughts: true, ThinkingBudget: &t}
		}
	}
	return nil
}

func (c *Client) sendStreamGemini(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  c.ProjectID.Value,
		Location: c.getGeminiRegion(opts.Model),
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return fmt.Errorf(i18n.T("vertexai_failed_gemini_client"), err)
	}

	contents := geminicommon.ConvertMessages(msgs)
	if len(contents) == 0 {
		return errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	config := c.buildGeminiConfig(opts)

	stream := client.Models.GenerateContentStream(ctx, opts.Model, contents, config)

	for response, err := range stream {
		if err != nil {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeError,
				Content: fmt.Sprintf(i18n.T("vertexai_stream_error"), err),
			}
			return err
		}

		text := geminicommon.ExtractText(response)
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

	return nil
}

// Claude message conversion

func (c *Client) toMessages(msgs []*chat.ChatCompletionMessage) []anthropic.MessageParam {
	// Convert messages to Anthropic format with proper role handling
	// - System messages become part of the first user message
	// - Messages must alternate user/assistant
	// - Skip empty messages

	var anthropicMessages []anthropic.MessageParam
	var systemContent string

	isFirstUserMessage := true
	lastRoleWasUser := false

	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" {
			continue // Skip empty messages
		}

		switch msg.Role {
		case chat.ChatMessageRoleSystem:
			// Accumulate system content to prepend to first user message
			if systemContent != "" {
				systemContent += "\\n" + msg.Content
			} else {
				systemContent = msg.Content
			}
		case chat.ChatMessageRoleUser:
			userContent := msg.Content
			if isFirstUserMessage && systemContent != "" {
				userContent = systemContent + "\\n\\n" + userContent
				isFirstUserMessage = false
			}
			if lastRoleWasUser {
				// Enforce alternation: add a minimal assistant message
				anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock("Okay.")))
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(userContent)))
			lastRoleWasUser = true
		case chat.ChatMessageRoleAssistant:
			// If first message is assistant and we have system content, prepend user message
			if isFirstUserMessage && systemContent != "" {
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
				lastRoleWasUser = true
				isFirstUserMessage = false
			} else if !lastRoleWasUser && len(anthropicMessages) > 0 {
				// Enforce alternation: add a minimal user message
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock("Hi")))
				lastRoleWasUser = true
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(msg.Content)))
			lastRoleWasUser = false
		default:
			// Other roles are ignored for Anthropic's message structure
			continue
		}
	}

	// If only system content was provided, create a user message with it
	if len(anthropicMessages) == 0 && systemContent != "" {
		anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
	}

	return anthropicMessages
}
