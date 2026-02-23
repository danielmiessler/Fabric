// Package azureaigateway - Vertex AI backend for Google Vertex AI using Gemini API format
package azureaigateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

// VertexAIBackend implements the Backend interface for Google Vertex AI (Gemini)
// through Azure APIM Gateway.
type VertexAIBackend struct {
	subscriptionKey string
}

// NewVertexAIBackend creates a new Vertex AI backend handler
func NewVertexAIBackend(subscriptionKey string) *VertexAIBackend {
	return &VertexAIBackend{subscriptionKey: subscriptionKey}
}

// ListModels returns the list of Gemini models available through Vertex AI
func (b *VertexAIBackend) ListModels() ([]string, error) {
	return []string{
		"gemini-3-pro-preview",
		"gemini-2.5-pro",
		"gemini-2.5-flash",
		"gemini-2.5-flash-lite",
		"gemini-2.0-flash",
		"gemini-2.0-flash-lite",
	}, nil
}

// BuildEndpoint constructs the Vertex AI API endpoint URL
// Uses /publishers/google/models/{model}:generateContent path per Azure APIM Gateway routing
// This is the APIM-specific path that proxies to Google's Vertex AI service
// (differs from direct Vertex AI API which uses /v1beta/models/{model}:generateContent)
func (b *VertexAIBackend) BuildEndpoint(baseURL, model string) string {
	return fmt.Sprintf("%s/publishers/google/models/%s:generateContent",
		strings.TrimSuffix(baseURL, "/"), url.PathEscape(model))
}

// AuthHeader returns the Vertex AI auth header (Google API key via APIM)
func (b *VertexAIBackend) AuthHeader() (string, string) {
	return "x-goog-api-key", b.subscriptionKey
}

// PrepareRequest converts messages to Gemini API format (contents/parts).
// System messages are extracted into the top-level "systemInstruction" field per the Gemini API spec.
func (b *VertexAIBackend) PrepareRequest(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) ([]byte, error) {
	var systemParts []string
	var contents []map[string]any
	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" {
			debuglog.Debug(debuglog.Basic, "Skipping empty message\n")
			continue
		}
		if msg.Role == chat.ChatMessageRoleSystem {
			systemParts = append(systemParts, msg.Content)
			continue
		}
		role := string(msg.Role)
		if msg.Role == chat.ChatMessageRoleAssistant {
			role = "model"
		}
		contents = append(contents, map[string]any{
			"role": role,
			"parts": []map[string]string{
				{"text": msg.Content},
			},
		})
	}

	debuglog.Debug(debuglog.Basic, "Vertex AI backend: %d input → %d API messages, %d system parts\n", len(msgs), len(contents), len(systemParts))

	if len(contents) == 0 {
		return nil, errors.New(i18n.T("azureaigateway_no_valid_messages"))
	}

	body := map[string]any{
		"contents": contents,
	}
	if len(systemParts) > 0 {
		body["systemInstruction"] = map[string]any{
			"parts": []map[string]string{
				{"text": strings.Join(systemParts, "\n\n")},
			},
		}
	}

	generationConfig := make(map[string]any)
	if opts.TopP != domain.DefaultTopP {
		generationConfig["topP"] = opts.TopP
	}
	if opts.Temperature != domain.DefaultTemperature {
		generationConfig["temperature"] = opts.Temperature
	}
	if len(generationConfig) > 0 {
		body["generationConfig"] = generationConfig
	}

	return json.Marshal(body)
}

// ParseResponse parses Gemini API response (candidates/content/parts)
func (b *VertexAIBackend) ParseResponse(body []byte) (string, error) {
	var resp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", fmt.Errorf(i18n.T("azureaigateway_vertexai_parse_response_failed"), err)
	}
	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New(i18n.T("azureaigateway_vertexai_no_content"))
	}

	var parts []string
	for _, part := range resp.Candidates[0].Content.Parts {
		if part.Text != "" {
			parts = append(parts, part.Text)
		}
	}
	return strings.Join(parts, ""), nil
}
