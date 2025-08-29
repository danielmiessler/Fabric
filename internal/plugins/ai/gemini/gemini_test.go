package gemini

import (
	"strings"
	"testing"

	"google.golang.org/genai"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Test buildModelNameFull method
func TestBuildModelNameFull(t *testing.T) {
	client := &Client{}

	tests := []struct {
		input    string
		expected string
	}{
		{"chat-bison-001", "models/chat-bison-001"},
		{"models/chat-bison-001", "models/chat-bison-001"},
		{"gemini-2.5-flash-preview-tts", "models/gemini-2.5-flash-preview-tts"},
	}

	for _, test := range tests {
		result := client.buildModelNameFull(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, got %v", test.input, test.expected, result)
		}
	}
}

// Test extractTextFromResponse method
func TestExtractTextFromResponse(t *testing.T) {
	client := &Client{}
	response := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []*genai.Part{
						{Text: "Hello, "},
						{Text: "world!"},
					},
				},
			},
		},
	}
	expected := "Hello, world!"

	result := client.extractTextFromResponse(response)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExtractTextFromResponse_Nil(t *testing.T) {
	client := &Client{}
	if got := client.extractTextFromResponse(nil); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestExtractTextFromResponse_EmptyGroundingChunks(t *testing.T) {
	client := &Client{}
	response := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content:           &genai.Content{Parts: []*genai.Part{{Text: "Hello"}}},
				GroundingMetadata: &genai.GroundingMetadata{GroundingChunks: nil},
			},
		},
	}
	if got := client.extractTextFromResponse(response); got != "Hello" {
		t.Fatalf("expected 'Hello', got %q", got)
	}
}

func TestBuildGenerateContentConfig_WithSearch(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Search: true}

	cfg, err := client.buildGenerateContentConfig(opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Tools == nil || len(cfg.Tools) != 1 || cfg.Tools[0].GoogleSearch == nil {
		t.Errorf("expected google search tool to be included")
	}
}

func TestBuildGenerateContentConfig_WithSearchAndLocation(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Search: true, SearchLocation: "America/Los_Angeles"}

	cfg, err := client.buildGenerateContentConfig(opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ToolConfig == nil || cfg.ToolConfig.RetrievalConfig == nil {
		t.Fatalf("expected retrieval config when search location provided")
	}
	if cfg.ToolConfig.RetrievalConfig.LanguageCode != opts.SearchLocation {
		t.Errorf("expected language code %s, got %s", opts.SearchLocation, cfg.ToolConfig.RetrievalConfig.LanguageCode)
	}
}

func TestBuildGenerateContentConfig_InvalidLocation(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Search: true, SearchLocation: "invalid"}

	_, err := client.buildGenerateContentConfig(opts)
	if err == nil {
		t.Fatalf("expected error for invalid location")
	}
}

func TestBuildGenerateContentConfig_LanguageCodeNormalization(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Search: true, SearchLocation: "en_US"}

	cfg, err := client.buildGenerateContentConfig(opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ToolConfig == nil || cfg.ToolConfig.RetrievalConfig.LanguageCode != "en-US" {
		t.Fatalf("expected normalized language code 'en-US', got %+v", cfg.ToolConfig)
	}
}

func TestBuildGenerateContentConfig_Thinking(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Thinking: domain.ThinkingLow}

	cfg, err := client.buildGenerateContentConfig(opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ThinkingConfig == nil || !cfg.ThinkingConfig.IncludeThoughts {
		t.Fatalf("expected thinking config with thoughts included")
	}
	if cfg.ThinkingConfig.ThinkingBudget == nil || *cfg.ThinkingConfig.ThinkingBudget != int32(domain.TokenBudgetLow) {
		t.Errorf("expected thinking budget %d, got %+v", domain.TokenBudgetLow, cfg.ThinkingConfig.ThinkingBudget)
	}
}

func TestBuildGenerateContentConfig_ThinkingTokens(t *testing.T) {
	client := &Client{}
	opts := &domain.ChatOptions{Thinking: domain.ThinkingLevel("123")}

	cfg, err := client.buildGenerateContentConfig(opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ThinkingConfig == nil || cfg.ThinkingConfig.ThinkingBudget == nil {
		t.Fatalf("expected thinking config with budget")
	}
	if *cfg.ThinkingConfig.ThinkingBudget != 123 {
		t.Errorf("expected thinking budget 123, got %d", *cfg.ThinkingConfig.ThinkingBudget)
	}
}

func TestCitationFormatting(t *testing.T) {
	client := &Client{}
	response := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{Parts: []*genai.Part{{Text: "Based on recent research, AI is advancing rapidly."}}},
				GroundingMetadata: &genai.GroundingMetadata{
					GroundingChunks: []*genai.GroundingChunk{
						{Web: &genai.GroundingChunkWeb{URI: "https://example.com/ai", Title: "AI Research"}},
						{Web: &genai.GroundingChunkWeb{URI: "https://news.com/tech", Title: "Tech News"}},
						{Web: &genai.GroundingChunkWeb{URI: "https://example.com/ai", Title: "AI Research"}}, // duplicate
					},
				},
			},
		},
	}

	result := client.extractTextFromResponse(response)
	if !strings.Contains(result, "## Sources") {
		t.Fatalf("expected sources section in result: %s", result)
	}
	if strings.Count(result, "- [") != 2 {
		t.Errorf("expected 2 unique citations, got %d", strings.Count(result, "- ["))
	}
}

// Test convertMessages handles role mapping correctly
func TestConvertMessagesRoles(t *testing.T) {
	client := &Client{}
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "user"},
		{Role: chat.ChatMessageRoleAssistant, Content: "assistant"},
		{Role: chat.ChatMessageRoleSystem, Content: "system"},
	}

	contents := client.convertMessages(msgs)

	expected := []string{"user", "model", "user"}

	if len(contents) != len(expected) {
		t.Fatalf("expected %d contents, got %d", len(expected), len(contents))
	}

	for i, c := range contents {
		if c.Role != expected[i] {
			t.Errorf("content %d expected role %s, got %s", i, expected[i], c.Role)
		}
	}
}

// Test isTTSModel method
func TestIsTTSModel(t *testing.T) {
	client := &Client{}

	tests := []struct {
		modelName string
		expected  bool
	}{
		{"gemini-2.5-flash-preview-tts", true},
		{"text-to-speech-model", true},
		{"TTS-MODEL", true},
		{"gemini-pro", false},
		{"chat-bison", false},
		{"", false},
	}

	for _, test := range tests {
		result := client.isTTSModel(test.modelName)
		if result != test.expected {
			t.Errorf("For model %v, expected %v, got %v", test.modelName, test.expected, result)
		}
	}
}

// Test generateWAVFile method (basic test)
func TestGenerateWAVFile(t *testing.T) {
	client := &Client{}

	// Test with minimal PCM data
	pcmData := []byte{0x00, 0x01, 0x02, 0x03}

	result, err := client.generateWAVFile(pcmData)
	if err != nil {
		t.Errorf("generateWAVFile failed: %v", err)
	}

	// Check that we got some data back
	if len(result) == 0 {
		t.Error("generateWAVFile returned empty data")
	}

	// Check that it starts with RIFF header
	if len(result) >= 4 && string(result[0:4]) != "RIFF" {
		t.Error("Generated WAV data doesn't start with RIFF header")
	}
}
