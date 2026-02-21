package azureaigateway

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

const bedrockAnthropicVersion = "bedrock-2023-05-31"

// BedrockBackend implements the Backend interface for AWS Bedrock through Azure APIM Gateway
type BedrockBackend struct {
	subscriptionKey string
}

// NewBedrockBackend creates a new Bedrock backend handler
func NewBedrockBackend(subscriptionKey string) *BedrockBackend {
	return &BedrockBackend{subscriptionKey: subscriptionKey}
}

// ListModels returns the list of available Bedrock inference profiles
func (b *BedrockBackend) ListModels() ([]string, error) {
	return []string{
		"us.anthropic.claude-3-haiku-20240307-v1:0",
		"us.anthropic.claude-3-opus-20240229-v1:0",
		"us.anthropic.claude-3-sonnet-20240229-v1:0",
		"us.anthropic.claude-3-5-haiku-20241022-v1:0",
		"us.anthropic.claude-3-5-sonnet-20240620-v1:0",
		"us.anthropic.claude-3-5-sonnet-20241022-v2:0",
		"us.anthropic.claude-3-7-sonnet-20250219-v1:0",
		"us.anthropic.claude-haiku-4-5-20251001-v1:0",
		"us.anthropic.claude-opus-4-20250514-v1:0",
		"us.anthropic.claude-opus-4-1-20250805-v1:0",
		"us.anthropic.claude-opus-4-5-20251101-v1:0",
		"us.anthropic.claude-opus-4-6-v1",
		"us.anthropic.claude-sonnet-4-20250514-v1:0",
		"us.anthropic.claude-sonnet-4-5-20250929-v1:0",
	}, nil
}

// BuildEndpoint constructs the Bedrock API endpoint URL
func (b *BedrockBackend) BuildEndpoint(baseURL, model string) string {
	return fmt.Sprintf("%s/model/%s/invoke", strings.TrimSuffix(baseURL, "/"), url.PathEscape(model))
}

// AuthHeader returns the Bedrock auth header (Bearer token)
func (b *BedrockBackend) AuthHeader() (string, string) {
	return "Authorization", "Bearer " + b.subscriptionKey
}

// PrepareRequest converts messages to Bedrock API format (Anthropic Messages API).
// System messages are extracted into the top-level "system" field per the Anthropic API spec.
func (b *BedrockBackend) PrepareRequest(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) ([]byte, error) {
	var systemParts []string
	var messages []map[string]any
	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" {
			continue
		}
		if msg.Role == chat.ChatMessageRoleSystem {
			systemParts = append(systemParts, msg.Content)
			continue
		}
		messages = append(messages, map[string]any{
			"role":    string(msg.Role),
			"content": msg.Content,
		})
	}

	debuglog.Debug(debuglog.Basic, "Bedrock backend: %d input → %d API messages, %d system parts\n", len(msgs), len(messages), len(systemParts))

	maxTokens := 4096
	if opts.MaxTokens > 0 {
		maxTokens = opts.MaxTokens
	}

	body := map[string]any{
		"anthropic_version": bedrockAnthropicVersion,
		"max_tokens":        maxTokens,
		"messages":          messages,
	}
	if len(systemParts) > 0 {
		body["system"] = strings.Join(systemParts, "\n\n")
	}
	if opts.TopP != domain.DefaultTopP {
		body["top_p"] = opts.TopP
	} else {
		body["temperature"] = opts.Temperature
	}

	return json.Marshal(body)
}

// ParseResponse parses Bedrock API response (Anthropic content blocks)
func (b *BedrockBackend) ParseResponse(body []byte) (string, error) {
	var resp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", fmt.Errorf("failed to parse Bedrock response: %w", err)
	}

	var parts []string
	for _, block := range resp.Content {
		if block.Type == "text" && block.Text != "" {
			parts = append(parts, block.Text)
		}
	}
	return strings.Join(parts, ""), nil
}
