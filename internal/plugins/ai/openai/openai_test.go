package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNeedsRawModeGPT6(t *testing.T) {
	if !NewClient().NeedsRawMode("gpt-6-astra") {
		t.Fatal("gpt-6 models should use raw mode")
	}
}

func TestBuildResponseRequestWithMaxTokens(t *testing.T) {

	var msgs []*chat.ChatCompletionMessage

	for range 2 {
		msgs = append(msgs, &chat.ChatCompletionMessage{
			Role:    "User",
			Content: "My msg",
		})
	}

	opts := &domain.ChatOptions{
		Temperature: 0.8,
		TopP:        0.9,
		Raw:         false,
		MaxTokens:   50,
	}

	var client = NewClient()
	request := client.buildResponseParams(msgs, opts)
	assert.Equal(t, shared.ResponsesModel(opts.Model), request.Model)
	assert.Equal(t, openai.Float(opts.Temperature), request.Temperature)
	assert.Equal(t, openai.Float(opts.TopP), request.TopP)
	assert.Equal(t, openai.Int(int64(opts.MaxTokens)), request.MaxOutputTokens)
}

func TestBuildResponseRequestNoMaxTokens(t *testing.T) {

	var msgs []*chat.ChatCompletionMessage

	for range 2 {
		msgs = append(msgs, &chat.ChatCompletionMessage{
			Role:    "User",
			Content: "My msg",
		})
	}

	opts := &domain.ChatOptions{
		Temperature: 0.8,
		TopP:        0.9,
		Raw:         false,
	}

	var client = NewClient()
	request := client.buildResponseParams(msgs, opts)
	assert.Equal(t, shared.ResponsesModel(opts.Model), request.Model)
	assert.Equal(t, openai.Float(opts.Temperature), request.Temperature)
	assert.Equal(t, openai.Float(opts.TopP), request.TopP)
	assert.False(t, request.MaxOutputTokens.Valid())
}

func TestBuildResponseParams_WithoutSearch(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "gpt-4o",
		Temperature: 0.7,
		Search:      false,
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}

	params := client.buildResponseParams(msgs, opts)

	assert.Nil(t, params.Tools, "Expected no tools when search is disabled")
	assert.Equal(t, shared.ResponsesModel(opts.Model), params.Model)
	assert.Equal(t, openai.Float(opts.Temperature), params.Temperature)
}

func TestBuildResponseParams_SearchToolsJSON(t *testing.T) {
	msgs := []*chat.ChatCompletionMessage{{Role: "user", Content: "What's the news?"}}
	tests := []struct {
		name      string
		toolName  string
		xSearch   bool
		location  string
		wantTools int
		wantJSON  string
	}{
		{
			name:      "OpenAI default",
			wantTools: 1,
			wantJSON:  `[{"type":"web_search_preview"}]`,
		},
		{
			name:      "OpenAI with location",
			location:  "America/Los_Angeles",
			wantTools: 1,
			wantJSON:  `[{"type":"web_search_preview","user_location":{"type":"approximate","timezone":"America/Los_Angeles"}}]`,
		},
		{
			name:      "xAI",
			toolName:  "web_search",
			xSearch:   true,
			wantTools: 2,
			wantJSON:  `[{"type":"web_search"},{"type":"x_search"}]`,
		},
		{
			name:      "xAI with location",
			toolName:  "web_search",
			xSearch:   true,
			location:  "America/Los_Angeles",
			wantTools: 2,
			wantJSON:  `[{"type":"web_search","user_location":{"type":"approximate","timezone":"America/Los_Angeles"}},{"type":"x_search"}]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient()
			client.SetWebSearchToolName(tc.toolName)
			client.SetEnableXSearch(tc.xSearch)
			params := client.buildResponseParams(msgs, &domain.ChatOptions{
				Model:          "gpt-4o",
				Search:         true,
				SearchLocation: tc.location,
			})

			require.Len(t, params.Tools, tc.wantTools)
			encoded, err := json.Marshal(params.Tools)
			require.NoError(t, err)
			assert.JSONEq(t, tc.wantJSON, string(encoded))
		})
	}
}

// TestBuildResponseParams_GrokAI_WithoutSearch confirms that a GrokAI
// style client without Search enabled does not append any tools.
// This protects the no-search path from regressions introduced by the
// new override logic.
func TestBuildResponseParams_GrokAI_WithoutSearch(t *testing.T) {
	client := NewClient()
	client.SetWebSearchToolName("web_search")
	client.SetEnableXSearch(true)

	opts := &domain.ChatOptions{
		Model:       "grok-4-fast-reasoning",
		Temperature: 0.7,
		Search:      false,
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}

	params := client.buildResponseParams(msgs, opts)

	assert.Nil(t, params.Tools, "Expected no tools when search is disabled")
}

func TestSendAndSendStreamPreserveProviderErrorMessage(t *testing.T) {
	const providerMessage = "The model 'gpt-nope' does not exist."
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error":{"message":%q,"type":"invalid_request_error","param":"model","code":"model_not_found"}}`, providerMessage)
	}))
	defer server.Close()

	client := newConfiguredOpenAITestClient(t, server.URL, true)
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "Hello"}}
	opts := &domain.ChatOptions{Model: "gpt-nope"}

	t.Run("Send", func(t *testing.T) {
		_, err := client.Send(context.Background(), msgs, opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), providerMessage)
		var apiErr *openai.Error
		require.ErrorAs(t, err, &apiErr)
	})

	t.Run("SendStream", func(t *testing.T) {
		err := client.SendStream(context.Background(), msgs, opts, make(chan domain.StreamUpdate, 1))
		require.Error(t, err)
		assert.Contains(t, err.Error(), providerMessage)
		var apiErr *openai.Error
		require.ErrorAs(t, err, &apiErr)
	})
}

func TestSendStreamIgnoresDataLessSSEEvents(t *testing.T) {
	tests := []struct {
		name                string
		path                string
		implementsResponses bool
		firstDelta          string
		secondDelta         string
	}{
		{
			name:                "Responses",
			path:                "/responses",
			implementsResponses: true,
			firstDelta:          `data: {"type":"response.output_text.delta","delta":"hello"}`,
			secondDelta:         `data: {"type":"response.output_text.delta","delta":" world"}`,
		},
		{
			name:        "Chat Completions",
			path:        "/chat/completions",
			firstDelta:  `data: {"id":"chatcmpl-test","object":"chat.completion.chunk","created":0,"model":"gpt-test","choices":[{"index":0,"delta":{"content":"hello"}}]}`,
			secondDelta: `data: {"id":"chatcmpl-test","object":"chat.completion.chunk","created":0,"model":"gpt-test","choices":[{"index":0,"delta":{"content":" world"}}]}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tc.path {
					t.Errorf("request path = %q, want %q", r.URL.Path, tc.path)
					http.NotFound(w, r)
					return
				}
				w.Header().Set("Content-Type", "text/event-stream")
				if _, ok := w.(http.Flusher); !ok {
					t.Errorf("response writer does not implement http.Flusher")
					return
				}
				writeTestSSEFrame(w, tc.firstDelta)
				writeTestSSEFrame(w, ": keep-alive")
				writeTestSSEFrame(w, "event: ping")
				writeTestSSEFrame(w, tc.secondDelta)
				writeTestSSEFrame(w, "data: [DONE]")
			}))
			defer server.Close()

			client := newConfiguredOpenAITestClient(t, server.URL, tc.implementsResponses)
			updates := make(chan domain.StreamUpdate)
			errCh := make(chan error, 1)
			go func() {
				errCh <- client.SendStream(context.Background(), []*chat.ChatCompletionMessage{
					{Role: chat.ChatMessageRoleUser, Content: "Hello"},
				}, &domain.ChatOptions{Model: "gpt-test"}, updates)
			}()

			var text strings.Builder
			for update := range updates {
				if update.Type == domain.StreamTypeContent {
					text.WriteString(update.Content)
				}
			}
			require.NoError(t, <-errCh)
			assert.Equal(t, "hello world\n", text.String())
		})
	}
}

func newConfiguredOpenAITestClient(t *testing.T, baseURL string, implementsResponses bool) *Client {
	t.Helper()
	client := NewClientCompatibleWithResponses("Test", baseURL, implementsResponses, nil)
	client.ApiKey.Value = "test-key"
	require.NoError(t, client.configure())
	return client
}

func writeTestSSEFrame(w http.ResponseWriter, frame string) {
	fmt.Fprintf(w, "%s\n\n", frame)
	w.(http.Flusher).Flush()
}

func TestCitationFormatting(t *testing.T) {
	var textParts []string
	var citations []string
	citationMap := make(map[string]bool)

	textParts = append(textParts, "Based on recent research, artificial intelligence is advancing rapidly.")

	mockCitations := []struct {
		URL   string
		Title string
	}{
		{"https://example.com/ai-research", "AI Research Advances 2025"},
		{"https://another-source.com/tech-news", "Technology News Today"},
		{"https://example.com/ai-research", "AI Research Advances 2025"}, // Duplicate to test deduplication
	}

	for _, citation := range mockCitations {
		citationKey := citation.URL + "|" + citation.Title
		if !citationMap[citationKey] {
			citationMap[citationKey] = true
			citationText := "- [" + citation.Title + "](" + citation.URL + ")"
			citations = append(citations, citationText)
		}
	}

	result := strings.Join(textParts, "")
	if len(citations) > 0 {
		result += "\n\n## Sources\n\n" + strings.Join(citations, "\n")
	}

	expectedText := "Based on recent research, artificial intelligence is advancing rapidly."
	assert.Contains(t, result, expectedText, "Expected result to contain original text")

	assert.Contains(t, result, "## Sources", "Expected result to contain Sources section")
	assert.Contains(t, result, "[AI Research Advances 2025](https://example.com/ai-research)", "Expected result to contain first citation")
	assert.Contains(t, result, "[Technology News Today](https://another-source.com/tech-news)", "Expected result to contain second citation")

	citationCount := strings.Count(result, "- [")
	assert.Equal(t, 2, citationCount, "Expected 2 unique citations")
}
