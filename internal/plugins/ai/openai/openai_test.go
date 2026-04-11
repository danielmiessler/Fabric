package openai

import (
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
	"github.com/stretchr/testify/assert"
)

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

	assert.NotNil(t, params.Tools, "Expected tools when search is enabled")
	assert.Len(t, params.Tools, 1, "Expected exactly one tool")

	tool := params.Tools[0]
	assert.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool")
	assert.Equal(t, responses.WebSearchToolType("web_search_preview"), tool.OfWebSearchPreview.Type)
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

	assert.NotNil(t, params.Tools, "Expected tools when search is enabled")
	tool := params.Tools[0]
	assert.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool")

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

	assert.NotNil(t, params.Tools, "Expected tools when search is enabled")
	assert.Len(t, params.Tools, 2, "Expected web_search plus x_search tools")

	webSearchTool := params.Tools[0]
	assert.NotNil(t, webSearchTool.OfWebSearchPreview, "Expected web search tool slot")
	assert.Equal(t, responses.WebSearchToolType("web_search"), webSearchTool.OfWebSearchPreview.Type)

	xSearchTool := params.Tools[1]
	assert.NotNil(t, xSearchTool.OfWebSearchPreview, "Expected x_search tool slot")
	assert.Equal(t, responses.WebSearchToolType("x_search"), xSearchTool.OfWebSearchPreview.Type)
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

	assert.NotNil(t, params.Tools, "Expected tools when search is enabled")
	assert.Len(t, params.Tools, 1, "Expected exactly one tool for default provider")

	tool := params.Tools[0]
	assert.NotNil(t, tool.OfWebSearchPreview, "Expected web search tool slot")
	assert.Equal(t, responses.WebSearchToolType("web_search_preview"), tool.OfWebSearchPreview.Type)
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

func TestCitationFormatting(t *testing.T) {
	// Test the citation formatting logic by simulating the citation extraction
	var textParts []string
	var citations []string
	citationMap := make(map[string]bool)

	// Simulate text content
	textParts = append(textParts, "Based on recent research, artificial intelligence is advancing rapidly.")

	// Simulate citations (as they would be extracted from OpenAI response)
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

	// Verify the result contains the expected text
	expectedText := "Based on recent research, artificial intelligence is advancing rapidly."
	assert.Contains(t, result, expectedText, "Expected result to contain original text")

	// Verify citations are included
	assert.Contains(t, result, "## Sources", "Expected result to contain Sources section")
	assert.Contains(t, result, "[AI Research Advances 2025](https://example.com/ai-research)", "Expected result to contain first citation")
	assert.Contains(t, result, "[Technology News Today](https://another-source.com/tech-news)", "Expected result to contain second citation")

	// Verify deduplication - should only have 2 unique citations, not 3
	citationCount := strings.Count(result, "- [")
	assert.Equal(t, 2, citationCount, "Expected 2 unique citations")
}
