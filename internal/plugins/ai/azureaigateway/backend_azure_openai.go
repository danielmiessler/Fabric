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

// AzureOpenAIBackend implements the Backend interface for Azure OpenAI through Azure APIM Gateway
type AzureOpenAIBackend struct {
	subscriptionKey string
}

// NewAzureOpenAIBackend creates a new Azure OpenAI backend handler
func NewAzureOpenAIBackend(subscriptionKey string) *AzureOpenAIBackend {
	return &AzureOpenAIBackend{subscriptionKey: subscriptionKey}
}

// ListModels returns the list of models available through Azure OpenAI.
// These are deployment names that must exist in your Azure OpenAI resource.
func (b *AzureOpenAIBackend) ListModels() ([]string, error) {
	return []string{
		"DeepSeek-R1",
		"gpt-4o",
		"gpt-4o-mini",
		"gpt-4-turbo",
		"gpt-35-turbo",
		"o1",
		"o1-mini",
	}, nil
}

// BuildEndpoint constructs the Azure OpenAI API endpoint URL
func (b *AzureOpenAIBackend) BuildEndpoint(baseURL, deploymentName string) string {
	return fmt.Sprintf("%s/openai/deployments/%s/chat/completions?api-version=2024-10-21",
		strings.TrimSuffix(baseURL, "/"), url.PathEscape(deploymentName))
}

// AuthHeader returns the Azure OpenAI auth header
func (b *AzureOpenAIBackend) AuthHeader() (string, string) {
	return "api-key", b.subscriptionKey
}

// PrepareRequest converts messages to Azure OpenAI (OpenAI-compatible) API format
func (b *AzureOpenAIBackend) PrepareRequest(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) ([]byte, error) {
	var messages []map[string]string
	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" {
			continue
		}
		messages = append(messages, map[string]string{
			"role":    string(msg.Role),
			"content": msg.Content,
		})
	}

	debuglog.Debug(debuglog.Basic, "Azure OpenAI backend: %d input → %d API messages\n", len(msgs), len(messages))

	body := map[string]any{
		"messages": messages,
	}
	if opts.TopP != domain.DefaultTopP {
		body["top_p"] = opts.TopP
	}
	if opts.Temperature != domain.DefaultTemperature {
		body["temperature"] = opts.Temperature
	}

	return json.Marshal(body)
}

// ParseResponse parses Azure OpenAI API response (OpenAI chat completions format)
func (b *AzureOpenAIBackend) ParseResponse(body []byte) (string, error) {
	var resp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", fmt.Errorf("failed to parse Azure OpenAI response: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in Azure OpenAI response")
	}
	return resp.Choices[0].Message.Content, nil
}
