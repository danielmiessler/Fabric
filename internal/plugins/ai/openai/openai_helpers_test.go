package openai

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/openai/openai-go/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientCompatibleRegistersOpenRouterSetupQuestions(t *testing.T) {
	t.Run("openrouter registers provider routing questions", func(t *testing.T) {
		client := NewClientCompatible(openRouterVendorName, "https://openrouter.ai/api/v1", nil)
		require.NotNil(t, client.ApiKey)
		require.NotNil(t, client.ApiBaseURL)
		require.NotNil(t, client.openRouterProviderOrder)
		require.NotNil(t, client.openRouterAllowFallbacks)
		assert.Equal(t, "https://openrouter.ai/api/v1", client.ApiBaseURL.Value)
	})

	t.Run("other vendors skip openrouter routing questions", func(t *testing.T) {
		client := NewClientCompatible("Groq", "https://api.groq.com/openai/v1", nil)
		require.NotNil(t, client.ApiKey)
		require.NotNil(t, client.ApiBaseURL)
		assert.Nil(t, client.openRouterProviderOrder)
		assert.Nil(t, client.openRouterAllowFallbacks)
	})
}

func TestSetResponsesAPIEnabledAndSupportsResponsesAPI(t *testing.T) {
	client := NewClientCompatibleWithResponses("OpenAI", "https://api.openai.com/v1", false, nil)
	assert.False(t, client.supportsResponsesAPI())

	client.SetResponsesAPIEnabled(true)
	assert.True(t, client.supportsResponsesAPI())
}

func TestNeedsRawMode(t *testing.T) {
	client := NewClient()

	assert.True(t, client.NeedsRawMode("gpt-5"))
	assert.True(t, client.NeedsRawMode("o3-mini"))
	assert.True(t, client.NeedsRawMode("gpt-4o-search-preview-2025-03-11"))
	assert.False(t, client.NeedsRawMode("claude-opus-4.6"))
}

func TestParseReasoningEffort(t *testing.T) {
	effort, ok := parseReasoningEffort(domain.ThinkingHigh)
	require.True(t, ok)
	assert.Equal(t, "high", string(effort))

	effort, ok = parseReasoningEffort(domain.ThinkingLevel("MEDIUM"))
	require.True(t, ok)
	assert.Equal(t, "medium", string(effort))

	effort, ok = parseReasoningEffort(domain.ThinkingLevel("invalid"))
	assert.False(t, ok)
	assert.Equal(t, "", string(effort))
}

func TestBuildChatCompletionParams_DeepseekSingleSystemMessageBecomesUser(t *testing.T) {
	client := NewClientCompatibleWithResponses("OpenRouter", "https://openrouter.ai/api/v1", false, nil)
	params := client.buildChatCompletionParams([]*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "system instruction"},
	}, &domain.ChatOptions{
		Model:       "deepseek/deepseek-chat",
		Temperature: 0.1,
	})

	body := marshalRequestBody(t, params)
	messages, ok := body["messages"].([]any)
	require.True(t, ok)
	require.Len(t, messages, 1)

	message, ok := messages[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, string(chat.ChatMessageRoleUser), message["role"])
}

func TestConvertChatMessage_MultiContent(t *testing.T) {
	client := NewClient()
	msg := chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "hello"},
			{Type: chat.ChatMessagePartTypeImageURL, ImageURL: &chat.ChatMessageImageURL{URL: "https://example.com/cat.png"}},
		},
	}

	body := marshalAnyJSON(t, client.convertChatMessage(msg))
	assert.Equal(t, string(chat.ChatMessageRoleUser), body["role"])

	content, ok := body["content"].([]any)
	require.True(t, ok)
	require.Len(t, content, 2)

	textPart, ok := content[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "text", textPart["type"])
	assert.Equal(t, "hello", textPart["text"])

	imagePart, ok := content[1].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "image_url", imagePart["type"])

	imageURL, ok := imagePart["image_url"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "https://example.com/cat.png", imageURL["url"])
}

func TestConvertMessage_MultiContent(t *testing.T) {
	msg := chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "hello"},
			{Type: chat.ChatMessagePartTypeImageURL, ImageURL: &chat.ChatMessageImageURL{URL: "https://example.com/cat.png"}},
		},
	}

	body := marshalAnyJSON(t, convertMessage(msg))
	assert.Equal(t, string(chat.ChatMessageRoleUser), body["role"])

	content, ok := body["content"].([]any)
	require.True(t, ok)
	require.Len(t, content, 2)

	textPart, ok := content[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "input_text", textPart["type"])
	assert.Equal(t, "hello", textPart["text"])

	imagePart, ok := content[1].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "input_image", imagePart["type"])
	assert.Equal(t, "https://example.com/cat.png", imagePart["image_url"])
	assert.Equal(t, "auto", imagePart["detail"])
}

func TestExtractText_AppendsDedupedCitations(t *testing.T) {
	var resp responses.Response
	require.NoError(t, json.Unmarshal([]byte(`{
		"output": [{
			"type": "message",
			"content": [{
				"type": "output_text",
				"text": "Result body",
				"annotations": [
					{"type": "url_citation", "url": "https://example.com/a", "title": "Example A"},
					{"type": "url_citation", "url": "https://example.com/a", "title": "Example A"},
					{"type": "url_citation", "url": "https://example.com/b", "title": "Example B"}
				]
			}]
		}]
	}`), &resp))

	client := NewClient()
	result := client.extractText(&resp)

	assert.Contains(t, result, "Result body")
	assert.Contains(t, result, "## Sources")
	assert.Contains(t, result, "- [Example A](https://example.com/a)")
	assert.Contains(t, result, "- [Example B](https://example.com/b)")
	assert.Equal(t, 2, countLinesWithPrefix(result, "- ["))
}

func marshalAnyJSON(t *testing.T, value any) map[string]any {
	t.Helper()

	data, err := json.Marshal(value)
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(data, &body))
	return body
}

func countLinesWithPrefix(input, prefix string) int {
	count := 0
	for _, line := range splitLines(input) {
		if strings.HasPrefix(line, prefix) {
			count++
		}
	}
	return count
}

func splitLines(input string) []string {
	if input == "" {
		return nil
	}
	return strings.Split(input, "\n")
}
