package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
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

func TestBuildResponseParams_WithSearch(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "gpt-4o",
		Temperature: 0.7,
		Search:      true,
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "What's the weather today?"},
	}

	params := client.buildResponseParams(msgs, opts)

	require.Len(t, params.Tools, 1, "Expected exactly one tool")

	tool := params.Tools[0]
	require.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool")
	assert.Equal(t, responses.WebSearchPreviewToolTypeWebSearchPreview, tool.OfWebSearchPreview.Type)
}

func TestBuildResponseParams_WithSearchAndLocation(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:          "gpt-4o",
		Temperature:    0.7,
		Search:         true,
		SearchLocation: "America/Los_Angeles",
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "What's the weather in San Francisco?"},
	}

	params := client.buildResponseParams(msgs, opts)

	require.NotEmpty(t, params.Tools, "Expected tools when search is enabled")
	tool := params.Tools[0]
	require.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool")

	userLocation := tool.OfWebSearchPreview.UserLocation
	assert.Equal(t, "approximate", string(userLocation.Type))
	assert.True(t, userLocation.Timezone.Valid(), "Expected timezone to be set")
	assert.Equal(t, opts.SearchLocation, userLocation.Timezone.Value)
}

// TestBuildResponseParams_GrokAI_WithSearch verifies that a client
// configured with a custom web search tool name and x_search enabled
// emits both tool entries with the xAI-expected type strings.
func TestBuildResponseParams_GrokAI_WithSearch(t *testing.T) {
	client := NewClient()
	client.SetWebSearchToolName("web_search")
	client.SetEnableXSearch(true)

	opts := &domain.ChatOptions{
		Model:       "grok-4-fast-reasoning",
		Temperature: 0.7,
		Search:      true,
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "What happened in the news today?"},
	}

	params := client.buildResponseParams(msgs, opts)

	require.Len(t, params.Tools, 2, "Expected web_search plus x_search tools")

	webSearchTool := params.Tools[0]
	require.NotNil(t, webSearchTool.OfWebSearch, "Expected web search tool slot")
	assert.Equal(t, responses.WebSearchToolType("web_search"), webSearchTool.OfWebSearch.Type)

	xSearchTool := params.Tools[1]
	require.NotNil(t, xSearchTool.OfWebSearch, "Expected x_search tool slot")
	assert.Equal(t, responses.WebSearchToolType("x_search"), xSearchTool.OfWebSearch.Type)
}

// TestBuildResponseParams_DefaultProvider_Unchanged guards backwards
// compatibility. A client that does not set the new override fields
// must continue emitting a single web_search_preview tool entry.
func TestBuildResponseParams_DefaultProvider_Unchanged(t *testing.T) {
	client := NewClient()

	opts := &domain.ChatOptions{
		Model:       "gpt-4o",
		Temperature: 0.7,
		Search:      true,
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "What is the capital of France?"},
	}

	params := client.buildResponseParams(msgs, opts)

	require.Len(t, params.Tools, 1, "Expected exactly one tool for default provider")

	tool := params.Tools[0]
	require.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool slot")
	assert.Equal(t, responses.WebSearchPreviewToolTypeWebSearchPreview, tool.OfWebSearchPreview.Type)
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

func TestProviderErrorMessageIsPreserved(t *testing.T) {
	const providerMessage = "The model 'gpt-nope' does not exist"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error":{"message":%q,"type":"invalid_request_error","param":"model","code":"model_not_found"}}`, providerMessage)
	}))
	defer server.Close()

	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "Hello"}}
	opts := &domain.ChatOptions{Model: "gpt-nope"}
	check := func(t *testing.T, err error) {
		require.Error(t, err)
		assert.Contains(t, err.Error(), "404 Not Found")
		assert.Contains(t, err.Error(), providerMessage)
		assert.Contains(t, err.Error(), "model_not_found")
		var apiErr *openai.Error
		require.ErrorAs(t, err, &apiErr)
	}

	for _, responses := range []bool{true, false} {
		client := newConfiguredOpenAITestClient(t, server.URL, responses)
		t.Run(fmt.Sprintf("Send/responses=%v", responses), func(t *testing.T) {
			_, err := client.Send(context.Background(), msgs, opts)
			check(t, err)
		})
		t.Run(fmt.Sprintf("SendStream/responses=%v", responses), func(t *testing.T) {
			check(t, client.SendStream(context.Background(), msgs, opts, make(chan domain.StreamUpdate, 1)))
		})
	}
	t.Run("TranscribeFile", func(t *testing.T) {
		audio := filepath.Join(t.TempDir(), "audio.mp3")
		require.NoError(t, os.WriteFile(audio, []byte("audio"), 0o600))
		_, err := newConfiguredOpenAITestClient(t, server.URL, false).
			TranscribeFile(context.Background(), audio, AllowedTranscriptionModels[0], false)
		check(t, err)
	})
}

func newConfiguredOpenAITestClient(t *testing.T, baseURL string, implementsResponses bool) *Client {
	t.Helper()
	client := NewClientCompatibleWithResponses("Test", baseURL, implementsResponses, nil)
	client.ApiKey.Value = "test-key"
	require.NoError(t, client.configure())
	return client
}

func writeSSE(w http.ResponseWriter, frame string) {
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

func TestSendStreamSkipsKeepAliveFrames(t *testing.T) {
	keepAlives := []string{": keep-alive", "event: ping"}
	tests := []struct {
		name                string
		implementsResponses bool
		delta               func(text string) string
		end                 string
	}{
		{
			name:                "responses",
			implementsResponses: true,
			delta: func(text string) string {
				return fmt.Sprintf(`data: {"type":"response.output_text.delta","delta":%q}`, text)
			},
			end: `data: {"type":"response.completed"}`,
		},
		{
			name: "chat completions",
			delta: func(text string) string {
				return fmt.Sprintf(`data: {"choices":[{"index":0,"delta":{"content":%q}}]}`, text)
			},
			end: "data: [DONE]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "text/event-stream")
				writeSSE(w, tt.delta("one "))
				writeSSE(w, keepAlives[0])
				writeSSE(w, tt.delta("two "))
				writeSSE(w, keepAlives[1])
				writeSSE(w, tt.delta("three"))
				writeSSE(w, tt.end)
			}))
			defer server.Close()

			client := newConfiguredOpenAITestClient(t, server.URL, tt.implementsResponses)
			channel := make(chan domain.StreamUpdate)
			errCh := make(chan error, 1)
			go func() {
				errCh <- client.SendStream(context.Background(),
					[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "Hello"}},
					&domain.ChatOptions{Model: "gpt-4o"}, channel)
			}()

			var content strings.Builder
			for update := range channel {
				if update.Type == domain.StreamTypeContent {
					content.WriteString(update.Content)
				}
			}
			require.NoError(t, <-errCh)
			assert.Equal(t, "one two three\n", content.String())
		})
	}
}

func TestBuildResponseParams_WebSearchToolsJSON(t *testing.T) {
	tests := []struct {
		name     string
		toolName string
		location string
		want     string
	}{
		{"default", "", "", `[{"type":"web_search_preview"}]`},
		{"default with location", "", "Europe/Paris",
			`[{"type":"web_search_preview","user_location":{"type":"approximate","timezone":"Europe/Paris"}}]`},
		{"custom name", "web_search", "", `[{"type":"web_search"}]`},
		{"custom name with location", "web_search", "Europe/Paris",
			`[{"type":"web_search","user_location":{"type":"approximate","timezone":"Europe/Paris"}}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient()
			client.SetWebSearchToolName(tt.toolName)
			params := client.buildResponseParams(
				[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "Hello"}},
				&domain.ChatOptions{Model: "gpt-4o", Search: true, SearchLocation: tt.location})

			got, err := json.Marshal(params.Tools)
			require.NoError(t, err)
			assert.JSONEq(t, tt.want, string(got))
		})
	}
}
