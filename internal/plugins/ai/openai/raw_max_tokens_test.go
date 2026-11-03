package openai

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/assert"
)

func rawMaxTokensMessages() []*chat.ChatCompletionMessage {
	return []*chat.ChatCompletionMessage{{Role: "User", Content: "My msg"}}
}

// Raw mode is switched on for reasoning models (Chatter.NeedsRawMode). It must drop the sampling
// options they reject, and keep the output limit the user asked for with --maxTokens.
func TestBuildResponseParams_RawKeepsMaxOutputTokens(t *testing.T) {
	opts := &domain.ChatOptions{Model: "gpt-5", Raw: true, MaxTokens: 50, Temperature: 0.8, TopP: 0.9}

	params := NewClient().buildResponseParams(rawMaxTokensMessages(), opts)

	assert.True(t, params.MaxOutputTokens.Valid())
	assert.EqualValues(t, 50, params.MaxOutputTokens.Value)
	assert.False(t, params.Temperature.Valid(), "raw mode must not send temperature")
	assert.False(t, params.TopP.Valid(), "raw mode must not send top_p")
}

func TestBuildResponseParams_RawWithoutMaxTokensSendsNoLimit(t *testing.T) {
	opts := &domain.ChatOptions{Model: "gpt-5", Raw: true}

	params := NewClient().buildResponseParams(rawMaxTokensMessages(), opts)

	assert.False(t, params.MaxOutputTokens.Valid())
}

func TestBuildChatCompletionParams_RawReasoningModelUsesMaxCompletionTokens(t *testing.T) {
	for _, model := range []string{"gpt-5", "gpt-6-luna", "o1", "o3-mini", "o4-mini"} {
		opts := &domain.ChatOptions{Model: model, Raw: true, MaxTokens: 50, Temperature: 0.8, TopP: 0.9}

		params := NewClient().buildChatCompletionParams(rawMaxTokensMessages(), opts)

		assert.True(t, params.MaxCompletionTokens.Valid(), model)
		assert.EqualValues(t, 50, params.MaxCompletionTokens.Value, model)
		assert.False(t, params.MaxTokens.Valid(), "%s: reasoning models reject max_tokens", model)
		assert.False(t, params.Temperature.Valid(), "%s: raw mode must not send temperature", model)
		assert.False(t, params.TopP.Valid(), "%s: raw mode must not send top_p", model)
	}
}

func TestBuildChatCompletionParams_RawOtherModelKeepsMaxTokens(t *testing.T) {
	opts := &domain.ChatOptions{Model: "glm-4.6", Raw: true, MaxTokens: 50, Temperature: 0.8}

	params := NewClient().buildChatCompletionParams(rawMaxTokensMessages(), opts)

	assert.True(t, params.MaxTokens.Valid())
	assert.EqualValues(t, 50, params.MaxTokens.Value)
	assert.False(t, params.MaxCompletionTokens.Valid())
	assert.False(t, params.Temperature.Valid())
}

func TestBuildChatCompletionParams_NotRawStillSendsMaxTokensAndSampling(t *testing.T) {
	opts := &domain.ChatOptions{Model: "gpt-4.1", MaxTokens: 50, Temperature: 0.8, TopP: 0.9, Seed: 7}

	params := NewClient().buildChatCompletionParams(rawMaxTokensMessages(), opts)

	assert.EqualValues(t, 50, params.MaxTokens.Value)
	assert.False(t, params.MaxCompletionTokens.Valid())
	assert.EqualValues(t, 0.8, params.Temperature.Value)
	assert.EqualValues(t, 0.9, params.TopP.Value)
	assert.EqualValues(t, 7, params.Seed.Value)
}

func TestBuildChatCompletionParams_RawStillDropsPenaltiesAndSeed(t *testing.T) {
	opts := &domain.ChatOptions{Model: "gpt-5", Raw: true, PresencePenalty: 0.5, FrequencyPenalty: 0.5, Seed: 7}

	params := NewClient().buildChatCompletionParams(rawMaxTokensMessages(), opts)

	assert.False(t, params.PresencePenalty.Valid())
	assert.False(t, params.FrequencyPenalty.Valid())
	assert.False(t, params.Seed.Valid())
}
