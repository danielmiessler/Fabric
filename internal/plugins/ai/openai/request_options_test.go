package openai

import (
	"context"
	"net/http"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
	"github.com/stretchr/testify/assert"
)

// Ensures providers that don't configure session/UA headers keep default behavior.
func TestRequestOptions_DisabledByDefault(t *testing.T) {
	client := &Client{}
	assert.Empty(t, client.requestOptions("session-123"))
}

// Ensures a configured session header and User-Agent are attached to requests.
func TestRequestOptions_SendsSessionAndUserAgent(t *testing.T) {
	client := &Client{
		sessionHeaderName: "x-opencode-session",
		userAgent:         "fabric/test",
	}

	rt := &captureRoundTripper{body: `{
		"id": "chatcmpl-1",
		"object": "chat.completion",
		"created": 0,
		"model": "glm-5.1",
		"choices": [{"index": 0, "message": {"role": "assistant", "content": "hi"}, "finish_reason": "stop"}]
	}`}

	sdk := openai.NewClient(
		option.WithAPIKey("test-key"),
		option.WithBaseURL("https://opencode.test/v1"),
		option.WithHTTPClient(&http.Client{Transport: rt}),
	)

	_, err := sdk.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model:    shared.ChatModel("glm-5.1"),
		Messages: []openai.ChatCompletionMessageParamUnion{openai.UserMessage("hi")},
	}, client.requestOptions("session-123")...)

	assert.NoError(t, err)
	assert.Equal(t, "session-123", rt.req.Header.Get("x-opencode-session"))
	assert.Equal(t, "fabric/test", rt.req.Header.Get("User-Agent"))
}

// Ensures the session header is omitted when no session ID is available while
// the User-Agent override still applies.
func TestRequestOptions_EmptySessionOmitsHeader(t *testing.T) {
	client := &Client{
		sessionHeaderName: "x-opencode-session",
		userAgent:         "fabric/test",
	}

	rt := &captureRoundTripper{body: `{
		"id": "chatcmpl-1",
		"object": "chat.completion",
		"created": 0,
		"model": "glm-5.1",
		"choices": [{"index": 0, "message": {"role": "assistant", "content": "hi"}, "finish_reason": "stop"}]
	}`}

	sdk := openai.NewClient(
		option.WithAPIKey("test-key"),
		option.WithBaseURL("https://opencode.test/v1"),
		option.WithHTTPClient(&http.Client{Transport: rt}),
	)

	_, err := sdk.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model:    shared.ChatModel("glm-5.1"),
		Messages: []openai.ChatCompletionMessageParamUnion{openai.UserMessage("hi")},
	}, client.requestOptions("")...)

	assert.NoError(t, err)
	assert.Empty(t, rt.req.Header.Get("x-opencode-session"))
	assert.Equal(t, "fabric/test", rt.req.Header.Get("User-Agent"))
}
