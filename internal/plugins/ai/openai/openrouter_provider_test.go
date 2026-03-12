package openai

import (
	"encoding/json"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildChatCompletionParams_AddsOpenRouterProviderRouting(t *testing.T) {
	client := NewClientCompatibleWithResponses(openRouterVendorName, "https://openrouter.ai/api/v1", false, nil)
	require.NotNil(t, client.openRouterProviderOrder)

	client.openRouterProviderOrder.Value = " amazon-bedrock, another-provider "
	require.NoError(t, client.configureOpenRouterProviderRouting())

	params := client.buildChatCompletionParams([]*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
	}, &domain.ChatOptions{
		Model:       "anthropic/claude-opus-4.6",
		Temperature: 0.2,
	})

	body := marshalRequestBody(t, params)
	provider := requireProviderObject(t, body)

	assert.Equal(t, false, provider["allow_fallbacks"])
	assert.Equal(t, []string{"amazon-bedrock", "another-provider"}, toStringSlice(t, provider["order"]))
}

func TestBuildChatCompletionParams_OmitsProviderRoutingForOtherVendors(t *testing.T) {
	client := NewClientCompatibleWithResponses("Groq", "https://api.groq.com/openai/v1", false, nil)

	params := client.buildChatCompletionParams([]*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
	}, &domain.ChatOptions{
		Model:       "llama-3.3-70b-versatile",
		Temperature: 0.2,
	})

	body := marshalRequestBody(t, params)
	assert.NotContains(t, body, "provider")
}

func TestBuildResponseParams_MergesOpenRouterProviderRoutingWithExistingExtraFields(t *testing.T) {
	client := NewClientCompatibleWithResponses(openRouterVendorName, "https://openrouter.ai/api/v1", false, nil)
	require.NotNil(t, client.openRouterProviderOrder)
	require.NotNil(t, client.openRouterAllowFallbacks)

	client.openRouterProviderOrder.Value = "amazon-bedrock"
	client.openRouterAllowFallbacks.Value = "true"
	require.NoError(t, client.configureOpenRouterProviderRouting())

	params := client.buildResponseParams([]*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
	}, &domain.ChatOptions{
		Model:            "anthropic/claude-opus-4.6",
		Temperature:      0.2,
		PresencePenalty:  0.4,
		FrequencyPenalty: 0.7,
		Seed:             42,
	})

	body := marshalRequestBody(t, params)
	provider := requireProviderObject(t, body)

	assert.Equal(t, true, provider["allow_fallbacks"])
	assert.Equal(t, []string{"amazon-bedrock"}, toStringSlice(t, provider["order"]))
	assert.Equal(t, 0.4, body["presence_penalty"])
	assert.Equal(t, 0.7, body["frequency_penalty"])
	assert.Equal(t, float64(42), body["seed"])
}

func TestConfigureOpenRouterProviderRouting(t *testing.T) {
	t.Run("skips non openrouter vendors", func(t *testing.T) {
		client := NewClientCompatibleWithResponses("Groq", "https://api.groq.com/openai/v1", false, nil)
		require.NoError(t, client.configureOpenRouterProviderRouting())
		assert.Nil(t, client.openRouterProviderRouting)
	})

	t.Run("ignores empty provider order", func(t *testing.T) {
		client := NewClientCompatibleWithResponses(openRouterVendorName, "https://openrouter.ai/api/v1", false, nil)
		require.NotNil(t, client.openRouterProviderOrder)

		client.openRouterProviderOrder.Value = " ,  , "
		require.NoError(t, client.configureOpenRouterProviderRouting())
		assert.Nil(t, client.openRouterProviderRouting)
	})

	t.Run("parses provider order and explicit false fallback", func(t *testing.T) {
		client := NewClientCompatibleWithResponses(openRouterVendorName, "https://openrouter.ai/api/v1", false, nil)
		require.NotNil(t, client.openRouterProviderOrder)
		require.NotNil(t, client.openRouterAllowFallbacks)

		client.openRouterProviderOrder.Value = "amazon-bedrock,anthropic"
		client.openRouterAllowFallbacks.Value = "off"

		require.NoError(t, client.configureOpenRouterProviderRouting())
		require.NotNil(t, client.openRouterProviderRouting)
		assert.Equal(t, []string{"amazon-bedrock", "anthropic"}, client.openRouterProviderRouting.Order)
		assert.False(t, client.openRouterProviderRouting.AllowFallbacks)
	})

	t.Run("returns error for invalid fallback value", func(t *testing.T) {
		client := NewClientCompatibleWithResponses(openRouterVendorName, "https://openrouter.ai/api/v1", false, nil)
		require.NotNil(t, client.openRouterProviderOrder)
		require.NotNil(t, client.openRouterAllowFallbacks)

		client.openRouterProviderOrder.Value = "amazon-bedrock"
		client.openRouterAllowFallbacks.Value = "maybe"

		err := client.configureOpenRouterProviderRouting()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "allow_fallbacks")
		assert.Nil(t, client.openRouterProviderRouting)
	})
}

func TestParseCommaSeparatedValues(t *testing.T) {
	assert.Nil(t, parseCommaSeparatedValues(""))
	assert.Nil(t, parseCommaSeparatedValues("   ,   "))
	assert.Equal(t, []string{"amazon-bedrock", "anthropic"}, parseCommaSeparatedValues(" amazon-bedrock , , anthropic "))
}

func TestParseFlexibleBool(t *testing.T) {
	trueValues := []string{"1", "true", "yes", "on", " TRUE "}
	for _, value := range trueValues {
		parsed, err := parseFlexibleBool(value)
		require.NoError(t, err, value)
		assert.True(t, parsed, value)
	}

	falseValues := []string{"0", "false", "no", "off", " Off "}
	for _, value := range falseValues {
		parsed, err := parseFlexibleBool(value)
		require.NoError(t, err, value)
		assert.False(t, parsed, value)
	}

	_, err := parseFlexibleBool("sometimes")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid boolean value")
}

func marshalRequestBody(t *testing.T, v any) map[string]any {
	t.Helper()

	data, err := json.Marshal(v)
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(data, &body))
	return body
}

func requireProviderObject(t *testing.T, body map[string]any) map[string]any {
	t.Helper()

	rawProvider, ok := body["provider"]
	require.True(t, ok, "provider object missing from request body")

	provider, ok := rawProvider.(map[string]any)
	require.True(t, ok, "provider object has unexpected type")
	return provider
}

func toStringSlice(t *testing.T, value any) []string {
	t.Helper()

	rawItems, ok := value.([]any)
	require.True(t, ok, "expected array value")

	items := make([]string, 0, len(rawItems))
	for _, item := range rawItems {
		s, ok := item.(string)
		require.True(t, ok, "expected string array item")
		items = append(items, s)
	}
	return items
}
