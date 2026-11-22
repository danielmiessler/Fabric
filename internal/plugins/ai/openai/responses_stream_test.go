package openai

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/stretchr/testify/require"
)

// responsesSSE answers a Responses API request with the given server-sent events.
func responsesSSE(events ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		for _, event := range events {
			_, _ = fmt.Fprint(w, event)
		}
	}
}

func sseEvent(name, data string) string {
	return fmt.Sprintf("event: %s\ndata: %s\n\n", name, data)
}

func streamResponses(t *testing.T, handler http.HandlerFunc) (updates []domain.StreamUpdate, err error) {
	t.Helper()

	server := httptest.NewServer(handler)
	defer server.Close()

	sdk := openai.NewClient(option.WithAPIKey("test-key"), option.WithBaseURL(server.URL))
	client := &Client{ApiClient: &sdk, ImplementsResponses: true}

	channel := make(chan domain.StreamUpdate, 16)
	err = client.SendStream(
		context.Background(),
		[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}},
		&domain.ChatOptions{Model: "gpt-5.6-terra"},
		channel,
	)
	for update := range channel {
		updates = append(updates, update)
	}
	return updates, err
}

func TestSendStreamResponsesReturnsErrorWhenResponseFails(t *testing.T) {
	_, err := streamResponses(t, responsesSSE(
		sseEvent("response.output_text.delta", `{"type":"response.output_text.delta","sequence_number":1,"item_id":"i","output_index":0,"content_index":0,"delta":"Hel","logprobs":[]}`),
		sseEvent("response.failed", `{"type":"response.failed","sequence_number":2,"response":{"id":"resp_1","object":"response","created_at":0,"status":"failed","model":"gpt-5.6-terra","output":[],"error":{"code":"server_error","message":"The model failed to generate a response."},"parallel_tool_calls":false,"tool_choice":"auto","tools":[]}}`),
	))

	require.Error(t, err)
	require.Contains(t, err.Error(), "The model failed to generate a response.")
}

func TestSendStreamResponsesCompletedStreamHasNoError(t *testing.T) {
	updates, err := streamResponses(t, responsesSSE(
		sseEvent("response.output_text.delta", `{"type":"response.output_text.delta","sequence_number":1,"item_id":"i","output_index":0,"content_index":0,"delta":"Hello","logprobs":[]}`),
		sseEvent("response.completed", `{"type":"response.completed","sequence_number":2,"response":{"id":"resp_1","object":"response","created_at":0,"status":"completed","model":"gpt-5.6-terra","output":[],"parallel_tool_calls":false,"tool_choice":"auto","tools":[]}}`),
	))

	require.NoError(t, err)
	require.Equal(t, "Hello", updates[0].Content)
}

func TestSendStreamResponsesReportsUsage(t *testing.T) {
	updates, err := streamResponses(t, responsesSSE(
		sseEvent("response.output_text.delta", `{"type":"response.output_text.delta","sequence_number":1,"item_id":"i","output_index":0,"content_index":0,"delta":"Hello","logprobs":[]}`),
		sseEvent("response.completed", `{"type":"response.completed","sequence_number":2,"response":{"id":"resp_1","object":"response","created_at":0,"status":"completed","model":"gpt-5.6-terra","output":[],"parallel_tool_calls":false,"tool_choice":"auto","tools":[],"usage":{"input_tokens":11,"output_tokens":7,"total_tokens":18,"input_tokens_details":{"cached_tokens":0},"output_tokens_details":{"reasoning_tokens":0}}}}`),
	))

	require.NoError(t, err)

	usage := reportedUsage(updates)
	require.NotNil(t, usage, "the usage of the completed response was not emitted")
	require.Equal(t, domain.UsageMetadata{InputTokens: 11, OutputTokens: 7, TotalTokens: 18}, *usage)
}

func TestSendStreamResponsesReportsUsageWhenIncomplete(t *testing.T) {
	updates, err := streamResponses(t, responsesSSE(
		sseEvent("response.output_text.delta", `{"type":"response.output_text.delta","sequence_number":1,"item_id":"i","output_index":0,"content_index":0,"delta":"Hello","logprobs":[]}`),
		sseEvent("response.incomplete", `{"type":"response.incomplete","sequence_number":2,"response":{"id":"resp_1","object":"response","created_at":0,"status":"incomplete","incomplete_details":{"reason":"max_output_tokens"},"model":"gpt-5.6-terra","output":[],"parallel_tool_calls":false,"tool_choice":"auto","tools":[],"usage":{"input_tokens":11,"output_tokens":7,"total_tokens":18,"input_tokens_details":{"cached_tokens":0},"output_tokens_details":{"reasoning_tokens":0}}}}`),
	))

	require.NoError(t, err)
	require.Equal(t, "Hello", updates[0].Content)

	usage := reportedUsage(updates)
	require.NotNil(t, usage, "the usage of the incomplete response was not emitted")
	require.Equal(t, domain.UsageMetadata{InputTokens: 11, OutputTokens: 7, TotalTokens: 18}, *usage)
}

func TestSendStreamResponsesReportsUsageWhenFailed(t *testing.T) {
	updates, err := streamResponses(t, responsesSSE(
		sseEvent("response.output_text.delta", `{"type":"response.output_text.delta","sequence_number":1,"item_id":"i","output_index":0,"content_index":0,"delta":"Hel","logprobs":[]}`),
		sseEvent("response.failed", `{"type":"response.failed","sequence_number":2,"response":{"id":"resp_1","object":"response","created_at":0,"status":"failed","model":"gpt-5.6-terra","output":[],"error":{"code":"server_error","message":"The model failed to generate a response."},"parallel_tool_calls":false,"tool_choice":"auto","tools":[],"usage":{"input_tokens":11,"output_tokens":4,"total_tokens":15,"input_tokens_details":{"cached_tokens":0},"output_tokens_details":{"reasoning_tokens":0}}}}`),
	))

	require.Error(t, err)
	require.Contains(t, err.Error(), "The model failed to generate a response.")

	usage := reportedUsage(updates)
	require.NotNil(t, usage, "the usage of the failed response was not emitted")
	require.Equal(t, domain.UsageMetadata{InputTokens: 11, OutputTokens: 4, TotalTokens: 15}, *usage)
}

func reportedUsage(updates []domain.StreamUpdate) *domain.UsageMetadata {
	for _, update := range slices.Backward(updates) {
		if update.Type == domain.StreamTypeUsage {
			return update.Usage
		}
	}
	return nil
}
