package anthropic

import (
	"context"
	"strings"
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/danielmiessler/fabric/internal/chat"
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

	models, err := client.ListModels(context.Background())
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
	models, err := client.ListModels(context.Background())

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
		Temperature: 0.8,                // Not the default, so the value must come from opts
		TopP:        domain.DefaultTopP, // The default, so the code sends temperature
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

func TestBuildMessageParams_UsesConfiguredMaxTokensByDefault(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}
	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.MaxTokens != int64(client.maxTokens) {
		t.Errorf("Expected default max_tokens %d, got %d", client.maxTokens, params.MaxTokens)
	}
}

func TestBuildMessageParams_UsesChatOptionsMaxTokens(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
		MaxTokens:   8192,
	}
	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.MaxTokens != int64(opts.MaxTokens) {
		t.Errorf("Expected max_tokens %d, got %d", opts.MaxTokens, params.MaxTokens)
	}
}

func TestBuildMessageParams_WithSearch(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: 0.8,
		TopP:        domain.DefaultTopP,
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
		Temperature:    0.8,
		TopP:           domain.DefaultTopP,
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

func TestBuildMessageParams_Opus47OmitsSamplingParams(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:       string(anthropic.ModelClaudeOpus4_7),
		Temperature: 0.8,
		TopP:        0.8,
		Search:      false,
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Temperature.Value != 0 {
		t.Errorf("expected temperature to be omitted for %s, got %f", opts.Model, params.Temperature.Value)
	}
	if params.TopP.Value != 0 {
		t.Errorf("expected top_p to be omitted for %s, got %f", opts.Model, params.TopP.Value)
	}
}

func TestModelBetasConfiguration(t *testing.T) {
	client := NewClient()
	model := string(anthropic.ModelClaudeSonnet5)
	betas, ok := client.modelBetas[model]
	if !ok || len(betas) != 1 || betas[0] != "context-1m-2025-08-07" {
		t.Errorf("expected beta mapping for %s", model)
	}
}

func TestCitationFormatting(t *testing.T) {
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
						URL:       "https://example.com/ai-research", // Same URL and title as the first citation
						Title:     "AI Research Advances 2025",
						CitedText: "machine learning models",
					},
				},
			},
		},
	}

	// This is a copy of the citation code in Send. The test does not call Send.
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

	expectedText := "Based on recent research, artificial intelligence is advancing rapidly. Machine learning models are becoming more sophisticated."
	if !strings.Contains(result, expectedText) {
		t.Errorf("Expected result to contain text: %s", expectedText)
	}

	if !strings.Contains(result, "## Sources") {
		t.Error("Expected result to contain Sources section")
	}

	if !strings.Contains(result, "[AI Research Advances 2025](https://example.com/ai-research)") {
		t.Error("Expected result to contain first citation")
	}

	if !strings.Contains(result, "[Technology News Today](https://another-source.com/tech-news)") {
		t.Error("Expected result to contain second citation")
	}

	citationCount := strings.Count(result, "- [")
	if citationCount != 2 {
		t.Errorf("Expected 2 unique citations, got %d", citationCount)
	}
}

func TestBuildMessageParams_DefaultValues(t *testing.T) {
	client := NewClient()

	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: domain.DefaultTemperature, // 0.7, sent because the API default is 1.0
		TopP:        domain.DefaultTopP,        // 0.9, the default, so the code sends temperature
		Search:      false,
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Temperature.Value != opts.Temperature {
		t.Errorf("Expected temperature %f, got %f", opts.Temperature, params.Temperature.Value)
	}

	if params.TopP.Value != 0 {
		t.Errorf("Expected TopP to not be set (0), but got %f", params.TopP.Value)
	}
}

func TestBuildMessageParams_ExplicitTopP(t *testing.T) {
	client := NewClient()

	opts := &domain.ChatOptions{
		Model:       "claude-3-5-sonnet-latest",
		Temperature: domain.DefaultTemperature, // 0.7, not sent because TopP is set
		TopP:        0.5,                       // Not the default, so the code sends it
		Search:      false,
	}

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock("Hello")),
	}

	params := client.buildMessageParams(messages, opts)

	if params.Temperature.Value != 0 {
		t.Errorf("Expected temperature to not be set (0), but got %f", params.Temperature.Value)
	}

	if params.TopP.Value != opts.TopP {
		t.Errorf("Expected TopP %f, got %f", opts.TopP, params.TopP.Value)
	}
}

func TestToMessages_MultiContentPDFAttachment(t *testing.T) {
	client := NewClient()
	msg := &chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{
				Type: chat.ChatMessagePartTypeText,
				Text: "Summarize this document.",
			},
			{
				Type: chat.ChatMessagePartTypeImageURL,
				ImageURL: &chat.ChatMessageImageURL{
					URL: "data:application/pdf;base64,SGVsbG8=",
				},
			},
		},
	}

	messages := client.toMessages([]*chat.ChatCompletionMessage{msg})
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(messages))
	}
	if len(messages[0].Content) != 2 {
		t.Fatalf("Expected 2 content blocks, got %d", len(messages[0].Content))
	}
	if messages[0].Content[0].OfText == nil || messages[0].Content[0].OfText.Text != "Summarize this document." {
		t.Fatalf("Expected first content block to be text, got %#v", messages[0].Content[0])
	}
	document := messages[0].Content[1].OfDocument
	if document == nil || document.Source.OfBase64 == nil {
		t.Fatalf("Expected second content block to be a base64 document, got %#v", messages[0].Content[1])
	}
	if document.Source.OfBase64.Data != "SGVsbG8=" {
		t.Fatalf("Expected document data to match base64 payload, got %s", document.Source.OfBase64.Data)
	}
}

func TestParseThinking_AdaptiveModelUsesAdaptivePlusEffort(t *testing.T) {
	for _, tc := range []struct {
		level  domain.ThinkingLevel
		effort anthropic.OutputConfigEffort
	}{
		{domain.ThinkingLow, anthropic.OutputConfigEffortLow},
		{domain.ThinkingMedium, anthropic.OutputConfigEffortMedium},
		{domain.ThinkingHigh, anthropic.OutputConfigEffortHigh},
	} {
		thinking, effort, ok := parseThinking(tc.level, string(anthropic.ModelClaudeSonnet5))
		if !ok {
			t.Fatalf("parseThinking(%q) returned ok=false", tc.level)
		}
		if thinking.OfAdaptive == nil {
			t.Errorf("level %q: expected OfAdaptive to be set on an adaptive model", tc.level)
		}
		if thinking.OfEnabled != nil {
			t.Errorf("level %q: OfEnabled must NOT be set on an adaptive model; the API rejects it", tc.level)
		}
		if effort != tc.effort {
			t.Errorf("level %q: expected effort %q, got %q", tc.level, tc.effort, effort)
		}
	}
}

func TestParseThinking_LegacyModelKeepsBudgetTokens(t *testing.T) {
	thinking, effort, ok := parseThinking(domain.ThinkingMedium, string(anthropic.ModelClaudeOpus4_7))
	if !ok {
		t.Fatal("parseThinking returned ok=false")
	}
	if thinking.OfEnabled == nil {
		t.Fatal("expected OfEnabled (budget_tokens) to be set on a non-adaptive model")
	}
	if thinking.OfEnabled.BudgetTokens != domain.TokenBudgetMedium {
		t.Errorf("expected budget %d, got %d", domain.TokenBudgetMedium, thinking.OfEnabled.BudgetTokens)
	}
	if effort != "" {
		t.Errorf("effort must be empty for a non-adaptive model, got %q", effort)
	}
}

func TestParseThinking_OffIsDisabledOnBothGenerations(t *testing.T) {
	for _, model := range []string{
		string(anthropic.ModelClaudeSonnet5),
		string(anthropic.ModelClaudeOpus4_7),
	} {
		thinking, effort, ok := parseThinking(domain.ThinkingOff, model)
		if !ok {
			t.Fatalf("model %s: parseThinking returned ok=false", model)
		}
		if thinking.OfDisabled == nil {
			t.Errorf("model %s: expected OfDisabled to be set", model)
		}
		if effort != "" {
			t.Errorf("model %s: effort must be empty when thinking is off, got %q", model, effort)
		}
	}
}

func TestParseThinking_NumericBudgetBucketsToEffortOnAdaptiveModel(t *testing.T) {
	for _, tc := range []struct {
		tokens string
		effort anthropic.OutputConfigEffort
	}{
		{"512", anthropic.OutputConfigEffortLow},
		{"1024", anthropic.OutputConfigEffortLow},
		{"2048", anthropic.OutputConfigEffortMedium},
		{"9000", anthropic.OutputConfigEffortHigh},
	} {
		thinking, effort, ok := parseThinking(domain.ThinkingLevel(tc.tokens), string(anthropic.ModelClaudeSonnet5))
		if !ok {
			t.Fatalf("tokens %s: parseThinking returned ok=false", tc.tokens)
		}
		if thinking.OfAdaptive == nil {
			t.Errorf("tokens %s: expected OfAdaptive on an adaptive model", tc.tokens)
		}
		if effort != tc.effort {
			t.Errorf("tokens %s: expected effort %q, got %q", tc.tokens, tc.effort, effort)
		}
	}
}

func TestBuildMessageParams_ThinkingShapeFollowsModel(t *testing.T) {
	for _, tc := range []struct {
		model        anthropic.Model
		wantAdaptive bool
		wantEffort   anthropic.OutputConfigEffort
	}{
		{anthropic.ModelClaudeSonnet5, true, anthropic.OutputConfigEffortMedium},
		{anthropic.ModelClaudeOpus4_7, false, ""},
	} {
		params := NewClient().buildMessageParams(
			[]anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock("Hello"))},
			&domain.ChatOptions{Model: string(tc.model), Thinking: domain.ThinkingMedium},
		)
		if got := params.Thinking.OfAdaptive != nil; got != tc.wantAdaptive {
			t.Errorf("model %s: adaptive=%v, want %v", tc.model, got, tc.wantAdaptive)
		}
		if got := params.Thinking.OfEnabled != nil; got == tc.wantAdaptive {
			t.Errorf("model %s: budget_tokens set=%v, want %v", tc.model, got, !tc.wantAdaptive)
		}
		if params.OutputConfig.Effort != tc.wantEffort {
			t.Errorf("model %s: effort=%q, want %q", tc.model, params.OutputConfig.Effort, tc.wantEffort)
		}
	}
}
