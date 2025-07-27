package anthropic

import (
	"strings"
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Test generated using Keploy
func TestNewClient_DefaultInitialization(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("Expected client to be initialized, got nil")
	}

	if client.ApiBaseURL.Value != defaultBaseUrl {
		t.Errorf("Expected default API Base URL to be %s, got %s", defaultBaseUrl, client.ApiBaseURL.Value)
	}

	if client.maxTokens != 4096 {
		t.Errorf("Expected default maxTokens to be 4096, got %d", client.maxTokens)
	}

	if len(client.models) == 0 {
		t.Error("Expected models to be initialized with default values, got empty list")
	}
}

// Test generated using Keploy
func TestClientListModels(t *testing.T) {
	client := NewClient()

	models, err := client.ListModels()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(models) != len(client.models) {
		t.Errorf("Expected %d models, got %d", len(client.models), len(models))
	}

	for i, model := range models {
		if model != client.models[i] {
			t.Errorf("Expected model at index %d to be %s, got %s", i, client.models[i], model)
		}
	}
}

func TestClient_ListModels_ReturnsCorrectModels(t *testing.T) {
	client := NewClient()
	models, err := client.ListModels()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(models) != len(client.models) {
		t.Errorf("Expected %d models, got %d", len(client.models), len(models))
	}

	for i, model := range models {
		if model != client.models[i] {
			t.Errorf("Expected model %s at index %d, got %s", client.models[i], i, model)
		}
	}
}

func TestBuildMessageParams_WithoutSearch(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: 0.7,
		Search:      false,
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Tools != nil {
		t.Error("Expected no tools when search is disabled, got tools")
	}

	if params.Model != anthropic.Model(opts.Model) {
		t.Errorf("Expected model %s, got %s", opts.Model, params.Model)
	}

	if params.Temperature.Value != opts.Temperature {
		t.Errorf("Expected temperature %f, got %f", opts.Temperature, params.Temperature.Value)
	}
}

func TestBuildMessageParams_WithSearch(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: 0.7,
		Search:      true,
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("What's the weather today?")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Tools == nil {
		t.Fatal("Expected tools when search is enabled, got nil")
	}

	if len(params.Tools) != 1 {
		t.Errorf("Expected 1 tool, got %d", len(params.Tools))
	}

	webTool := params.Tools[0].OfWebSearchTool20250305
	if webTool == nil {
		t.Fatal("Expected web search tool, got nil")
	}

	if webTool.Name != "web_search" {
		t.Errorf("Expected tool name 'web_search', got %s", webTool.Name)
	}

	if webTool.Type != "web_search_20250305" {
		t.Errorf("Expected tool type 'web_search_20250305', got %s", webTool.Type)
	}
}

func TestBuildMessageParams_WithSearchAndLocation(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:          "claude-3-5-sonnet-latest",
		Temperature:    0.7,
		Search:         true,
		SearchLocation: "America/Los_Angeles",
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("What's the weather in San Francisco?")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Tools == nil {
		t.Fatal("Expected tools when search is enabled, got nil")
	}

	webTool := params.Tools[0].OfWebSearchTool20250305
	if webTool == nil {
		t.Fatal("Expected web search tool, got nil")
	}

	if webTool.UserLocation.Type != "approximate" {
		t.Errorf("Expected location type 'approximate', got %s", webTool.UserLocation.Type)
	}

	if webTool.UserLocation.Timezone.Value != opts.SearchLocation {
		t.Errorf("Expected timezone %s, got %s", opts.SearchLocation, webTool.UserLocation.Timezone.Value)
	}
}

func TestCitationFormatting(t *testing.T) {
	// Test the citation formatting logic by creating a mock message with citations
	message := &anthropic.Message{
		Content: []anthropic.ContentBlockUnion{
			{
				Type: "text",
				Text: "Based on recent research, artificial intelligence is advancing rapidly.",
				Citations: []anthropic.TextCitationUnion{
					{
						Type:      "web_search_result_location",
						URL:       "https://example.com/ai-research",
						Title:     "AI Research Advances 2025",
						CitedText: "artificial intelligence is advancing rapidly",
					},
					{
						Type:      "web_search_result_location",
						URL:       "https://another-source.com/tech-news",
						Title:     "Technology News Today",
						CitedText: "recent developments in AI",
					},
				},
			},
			{
				Type: "text",
				Text: " Machine learning models are becoming more sophisticated.",
				Citations: []anthropic.TextCitationUnion{
					{
						Type:      "web_search_result_location",
						URL:       "https://example.com/ai-research", // Duplicate URL should be deduplicated
						Title:     "AI Research Advances 2025",
						CitedText: "machine learning models",
					},
				},
			},
		},
	}

	// Extract text and citations using the same logic as the Send method
	var textParts []string
	var citations []string
	citationMap := make(map[string]bool)

	for _, block := range message.Content {
		if block.Type == "text" && block.Text != "" {
			textParts = append(textParts, block.Text)

			for _, citation := range block.Citations {
				if citation.Type == "web_search_result_location" {
					citationKey := citation.URL + "|" + citation.Title
					if !citationMap[citationKey] {
						citationMap[citationKey] = true
						citationText := "- [" + citation.Title + "](" + citation.URL + ")"
						if citation.CitedText != "" {
							citationText += " - \"" + citation.CitedText + "\""
						}
						citations = append(citations, citationText)
					}
				}
			}
		}
	}

	result := strings.Join(textParts, "")
	if len(citations) > 0 {
		result += "\n\n## Sources\n\n" + strings.Join(citations, "\n")
	}

	// Verify the result contains the expected text
	expectedText := "Based on recent research, artificial intelligence is advancing rapidly. Machine learning models are becoming more sophisticated."
	if !strings.Contains(result, expectedText) {
		t.Errorf("Expected result to contain text: %s", expectedText)
	}

	// Verify citations are included
	if !strings.Contains(result, "## Sources") {
		t.Error("Expected result to contain Sources section")
	}

	if !strings.Contains(result, "[AI Research Advances 2025](https://example.com/ai-research)") {
		t.Error("Expected result to contain first citation")
	}

	if !strings.Contains(result, "[Technology News Today](https://another-source.com/tech-news)") {
		t.Error("Expected result to contain second citation")
	}

	// Verify deduplication - should only have 2 unique citations, not 3
	citationCount := strings.Count(result, "- [")
	if citationCount != 2 {
		t.Errorf("Expected 2 unique citations, got %d", citationCount)
	}
}
