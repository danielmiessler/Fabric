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

	// The Anthropic client authenticates with Application Default Credentials.
	vertexOpt := vertex.WithGoogleAuth(ctx, region, projectID, cloudPlatformScope)
	client := anthropic.NewClient(vertexOpt)
	c.client = &client

	return nil
}

func (c *Client) ListModels(_ context.Context) ([]string, error) {
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, cloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_failed_google_credentials"), err)
	}
	httpClient := oauth2.NewClient(ctx, creds.TokenSource)

	type result struct {
		models    []string
		err       error
		publisher string
	}
	// One extra slot for the static Gemini list.
	results := make(chan result, len(publishers)+1)

	for _, pub := range publishers {
		go func(publisher string) {
			models, err := listPublisherModels(ctx, httpClient, c.Region.Value, c.ProjectID.Value, publisher)
			results <- result{models: models, err: err, publisher: publisher}
		}(pub)
	}

	go func() {
		results <- result{models: getKnownGeminiModels(), err: nil, publisher: "gemini"}
	}()

	var allModels []string
	for range len(publishers) + 1 {
		r := <-results
		if r.err != nil {
			// One failed source does not stop the listing.
			debuglog.Debug(debuglog.Basic, "Failed to list %s models: %v\n", r.publisher, r.err)
			continue
		}
		allModels = append(allModels, r.models...)
	}

	if len(allModels) == 0 {
		return nil, errors.New(i18n.T("vertexai_no_models_found"))
	}

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

	anthropicMessages := c.toMessages(msgs)
	if len(anthropicMessages) == 0 {
		return "", errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	params := anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: getMaxTokens(opts),
		Messages:  anthropicMessages,
	}

	// Some Claude models reject a request that sets both Temperature and TopP.
	if opts.TopP != domain.DefaultTopP {
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	response, err := c.client.Messages.New(ctx, params)
	if err != nil {
		return "", err
	}

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

func (c *Client) SendStream(_ context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
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

	anthropicMessages := c.toMessages(msgs)
	if len(anthropicMessages) == 0 {
		return errors.New(i18n.T("vertexai_no_valid_messages"))
	}

	params := anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: getMaxTokens(opts),
		Messages:  anthropicMessages,
	}

	// Some Claude models reject a request that sets both Temperature and TopP.
	if opts.TopP != domain.DefaultTopP {
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	stream := c.client.Messages.NewStreaming(ctx, params)

	for stream.Next() {
		event := stream.Current()

		if event.Delta.Text != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.Delta.Text,
			}
		}

		// message_start carries usage in Message.Usage. message_delta carries it in Usage.
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

// getGeminiRegion returns "global" for preview models. Vertex AI serves some
// preview models only on the global endpoint.
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

func (c *Client) buildGeminiConfig(opts *domain.ChatOptions) *genai.GenerateContentConfig {
	temperature := float32(opts.Temperature)
	topP := float32(opts.TopP)
	config := &genai.GenerateContentConfig{
		Temperature:     &temperature,
		TopP:            &topP,
		MaxOutputTokens: int32(getMaxTokens(opts)),
	}

	if opts.Search {
		config.Tools = []*genai.Tool{{GoogleSearch: &genai.GoogleSearch{}}}
	}

	if tc := parseGeminiThinking(opts.Thinking); tc != nil {
		config.ThinkingConfig = tc
	}

	return config
}

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

func (c *Client) toMessages(msgs []*chat.ChatCompletionMessage) []anthropic.MessageParam {
	// System content goes into the first user message. The code adds filler
	// messages to keep user and assistant roles alternating.

	var anthropicMessages []anthropic.MessageParam
	var systemContent string

	isFirstUserMessage := true
	lastRoleWasUser := false

	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" {
			continue
		}

		switch msg.Role {
		case chat.ChatMessageRoleSystem:
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
				anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock("Okay.")))
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(userContent)))
			lastRoleWasUser = true
		case chat.ChatMessageRoleAssistant:
			if isFirstUserMessage && systemContent != "" {
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
				lastRoleWasUser = true
				isFirstUserMessage = false
			} else if !lastRoleWasUser && len(anthropicMessages) > 0 {
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock("Hi")))
				lastRoleWasUser = true
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(msg.Content)))
			lastRoleWasUser = false
		default:
			continue
		}
	}

	if len(anthropicMessages) == 0 && systemContent != "" {
		anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
	}

	return anthropicMessages
}
