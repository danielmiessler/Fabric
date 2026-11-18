package azureaigateway

import (
	"encoding/json"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/require"
)

func prepareBedrockBody(t *testing.T, model string, topP float64) map[string]any {
	t.Helper()

	opts := &domain.ChatOptions{Model: model, Temperature: domain.DefaultTemperature, TopP: topP}
	raw, err := NewBedrockBackend("key").PrepareRequest(
		[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}, opts)
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(raw, &body))
	return body
}

// Claude Opus 4.7 and later, and the Claude 5 models, reject temperature and top_p with a 400.
// A model that is not in the allowlist, such as a future one, gets neither.
func TestBedrockPrepareRequestOmitsSamplingForModelsThatRejectIt(t *testing.T) {
	for _, model := range []string{
		"us.anthropic.claude-opus-5-5-v1:0",
		"global.anthropic.claude-sonnet-5-5",
		"anthropic.claude-opus-4-7-v1:0",
		"us.anthropic.claude-opus-6-v1:0",
	} {
		t.Run(model, func(t *testing.T) {
			for _, topP := range []float64{domain.DefaultTopP, 0.5} {
				body := prepareBedrockBody(t, model, topP)
				require.NotContains(t, body, "temperature")
				require.NotContains(t, body, "top_p")
			}
		})
	}
}

func TestBedrockPrepareRequestKeepsSamplingForOtherModels(t *testing.T) {
	for _, model := range []string{
		"us.anthropic.claude-sonnet-4-5-20250929-v1:0",
		"us.anthropic.claude-sonnet-4-20250514-v1:0",
		"us.anthropic.claude-3-haiku-20240307-v1:0",
	} {
		body := prepareBedrockBody(t, model, domain.DefaultTopP)
		require.Contains(t, body, "temperature", model)
		require.NotContains(t, body, "top_p", model)
	}

	body := prepareBedrockBody(t, "us.anthropic.claude-sonnet-4-5-20250929-v1:0", domain.DefaultTopP)
	require.Contains(t, body, "temperature")
	require.NotContains(t, body, "top_p")

	body = prepareBedrockBody(t, "us.anthropic.claude-sonnet-4-5-20250929-v1:0", 0.5)
	require.Contains(t, body, "top_p")
	require.NotContains(t, body, "temperature")
}
