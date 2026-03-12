package openai

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSend_UsesChatCompletionsAPI(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/chat/completions", r.URL.Path)
		assert.Equal(t, "Bearer chat-key", r.Header.Get("Authorization"))

		body := decodeJSONBody(t, r.Body)
		assert.Equal(t, "test-chat-model", body["model"])

		messages := requireJSONArray(t, body["messages"])
		require.Len(t, messages, 1)

		first := requireJSONObject(t, messages[0])
		assert.Equal(t, "user", first["role"])
		assert.Equal(t, "hello from chat api", first["content"])

		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, `{"id":"chatcmpl_123","object":"chat.completion","created":0,"model":"test-chat-model","choices":[{"index":0,"message":{"role":"assistant","content":"chat completion reply"},"finish_reason":"stop"}]}`)
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredRuntimeTestClient(t, "Groq", srv.URL, "chat-key", false)

	result, err := client.Send(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello from chat api"},
	}, &domain.ChatOptions{
		Model:       "test-chat-model",
		Temperature: 0.4,
	})
	require.NoError(t, err)
	assert.Equal(t, "chat completion reply", result)
}

func TestSend_UsesResponsesAPI(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/responses", r.URL.Path)
		assert.Equal(t, "Bearer responses-key", r.Header.Get("Authorization"))

		body := decodeJSONBody(t, r.Body)
		assert.Equal(t, "test-response-model", body["model"])

		input := requireJSONArray(t, body["input"])
		require.Len(t, input, 1)

		first := requireJSONObject(t, input[0])
		assert.Equal(t, "user", first["role"])
		assert.Equal(t, "hello from responses api", first["content"])

		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, `{"id":"resp_123","object":"response","created_at":0,"status":"completed","error":null,"incomplete_details":null,"instructions":null,"max_output_tokens":null,"model":"test-response-model","output":[{"id":"msg_123","type":"message","role":"assistant","status":"completed","content":[{"type":"output_text","text":"responses api reply","annotations":[]}]}],"parallel_tool_calls":false,"temperature":0.2,"tool_choice":"auto","tools":[],"top_p":1,"truncation":"disabled","usage":{"input_tokens":3,"input_tokens_details":{"cached_tokens":0},"output_tokens":4,"output_tokens_details":{"reasoning_tokens":0},"total_tokens":7},"user":null,"metadata":{}}`)
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredRuntimeTestClient(t, "OpenAI", srv.URL, "responses-key", true)

	result, err := client.Send(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello from responses api"},
	}, &domain.ChatOptions{
		Model:       "test-response-model",
		Temperature: 0.2,
	})
	require.NoError(t, err)
	assert.Equal(t, "responses api reply", result)
}

func TestSendStream_UsesChatCompletionsAPI(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/chat/completions", r.URL.Path)
		assert.Equal(t, "Bearer stream-chat-key", r.Header.Get("Authorization"))

		body := decodeJSONBody(t, r.Body)
		assert.Equal(t, true, body["stream"])

		w.Header().Set("Content-Type", "text/event-stream")
		_, err := io.WriteString(w, strings.Join([]string{
			`data: {"id":"chatcmpl_123","object":"chat.completion.chunk","created":0,"model":"test-chat-model","choices":[{"index":0,"delta":{"content":"hello "},"finish_reason":null}]}`,
			"",
			`data: {"id":"chatcmpl_123","object":"chat.completion.chunk","created":0,"model":"test-chat-model","choices":[{"index":0,"delta":{"content":"world"},"finish_reason":null}],"usage":{"prompt_tokens":3,"completion_tokens":2,"total_tokens":5}}`,
			"",
			"data: [DONE]",
			"",
		}, "\n"))
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredRuntimeTestClient(t, "Groq", srv.URL, "stream-chat-key", false)

	updates := collectStreamUpdates(t, func(ch chan domain.StreamUpdate) error {
		return client.SendStream([]*chat.ChatCompletionMessage{
			{Role: chat.ChatMessageRoleUser, Content: "stream please"},
		}, &domain.ChatOptions{
			Model: "test-chat-model",
		}, ch)
	})

	require.Len(t, updates, 4)
	assert.Equal(t, domain.StreamTypeContent, updates[0].Type)
	assert.Equal(t, "hello ", updates[0].Content)
	assert.Equal(t, domain.StreamTypeContent, updates[1].Type)
	assert.Equal(t, "world", updates[1].Content)
	assert.Equal(t, domain.StreamTypeUsage, updates[2].Type)
	require.NotNil(t, updates[2].Usage)
	assert.Equal(t, 3, updates[2].Usage.InputTokens)
	assert.Equal(t, 2, updates[2].Usage.OutputTokens)
	assert.Equal(t, 5, updates[2].Usage.TotalTokens)
	assert.Equal(t, domain.StreamTypeContent, updates[3].Type)
	assert.Equal(t, "\n", updates[3].Content)
}

func TestSendStream_UsesResponsesAPI(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/responses", r.URL.Path)
		assert.Equal(t, "Bearer stream-responses-key", r.Header.Get("Authorization"))

		body := decodeJSONBody(t, r.Body)
		assert.Equal(t, true, body["stream"])

		w.Header().Set("Content-Type", "text/event-stream")
		_, err := io.WriteString(w, strings.Join([]string{
			`data: {"type":"response.output_text.delta","delta":"alpha"}`,
			"",
			`data: {"type":"response.output_text.delta","delta":" beta"}`,
			"",
			`data: {"type":"response.output_text.done","text":"alpha beta"}`,
			"",
			"data: [DONE]",
			"",
		}, "\n"))
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredRuntimeTestClient(t, "OpenAI", srv.URL, "stream-responses-key", true)

	updates := collectStreamUpdates(t, func(ch chan domain.StreamUpdate) error {
		return client.SendStream([]*chat.ChatCompletionMessage{
			{Role: chat.ChatMessageRoleUser, Content: "stream via responses"},
		}, &domain.ChatOptions{
			Model: "test-response-model",
		}, ch)
	})

	require.Len(t, updates, 3)
	assert.Equal(t, domain.StreamTypeContent, updates[0].Type)
	assert.Equal(t, "alpha", updates[0].Content)
	assert.Equal(t, domain.StreamTypeContent, updates[1].Type)
	assert.Equal(t, " beta", updates[1].Content)
	assert.Equal(t, domain.StreamTypeContent, updates[2].Type)
	assert.Equal(t, "\n", updates[2].Content)
}

func newConfiguredRuntimeTestClient(t *testing.T, vendorName, baseURL, apiKey string, implementsResponses bool) *Client {
	t.Helper()

	client := NewClientCompatibleWithResponses(vendorName, baseURL, implementsResponses, nil)
	client.ApiKey.Value = apiKey
	client.ApiBaseURL.Value = baseURL
	require.NoError(t, client.configure())
	return client
}

func collectStreamUpdates(t *testing.T, run func(chan domain.StreamUpdate) error) []domain.StreamUpdate {
	t.Helper()

	ch := make(chan domain.StreamUpdate, 16)
	errCh := make(chan error, 1)

	go func() {
		errCh <- run(ch)
	}()

	var updates []domain.StreamUpdate
	for update := range ch {
		updates = append(updates, update)
	}

	require.NoError(t, <-errCh)
	return updates
}

func decodeJSONBody(t *testing.T, body io.ReadCloser) map[string]any {
	t.Helper()
	defer body.Close()

	payload, err := io.ReadAll(body)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(payload, &decoded))
	return decoded
}

func requireJSONArray(t *testing.T, value any) []any {
	t.Helper()

	items, ok := value.([]any)
	require.True(t, ok, "expected JSON array")
	return items
}

func requireJSONObject(t *testing.T, value any) map[string]any {
	t.Helper()

	obj, ok := value.(map[string]any)
	require.True(t, ok, "expected JSON object")
	return obj
}
