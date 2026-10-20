package perplexity

import (
	"encoding/json"
	"testing"

	openaisdk "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func agentJSON(t *testing.T, p responses.ResponseNewParams) map[string]any {
	t.Helper()
	agentParams(&p)
	data, err := json.Marshal(p)
	require.NoError(t, err)
	var ret map[string]any
	require.NoError(t, json.Unmarshal(data, &ret))
	return ret
}

func TestAgentParams(t *testing.T) {
	webSearch := []any{map[string]any{"type": "web_search"}}

	// A legacy name that maps to a preset: no model, no temperature.
	got := agentJSON(t, responses.ResponseNewParams{Model: "sonar-pro", Temperature: openaisdk.Float(0.7)})
	assert.Equal(t, "fast", got["preset"])
	assert.Empty(t, got["model"])
	assert.NotContains(t, got, "temperature")
	assert.Equal(t, webSearch, got["tools"])

	// A legacy name that maps to a model.
	got = agentJSON(t, responses.ResponseNewParams{Model: "sonar", Temperature: openaisdk.Float(0.7)})
	assert.Equal(t, "perplexity/sonar", got["model"])
	assert.NotContains(t, got, "preset")
	assert.InDelta(t, 0.7, got["temperature"], 1e-9)

	// The --search tool is kept and not added a second time.
	got = agentJSON(t, responses.ResponseNewParams{
		Model: "anthropic/claude-haiku-4-5",
		Tools: []responses.ToolUnionParam{responses.ToolParamOfWebSearch("web_search")},
	})
	assert.Equal(t, webSearch, got["tools"])
	assert.InDelta(t, anthropicMaxTokens, got["max_output_tokens"], 0)
}
