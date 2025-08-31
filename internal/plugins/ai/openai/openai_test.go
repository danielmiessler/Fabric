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

	for i := 0; i < 2; i++ {
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

	for i := 0; i < 2; i++ {
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

func TestBuildChatCompletionParams_WithSchema(t *testing.T) {
	client := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Generate a JSON object."},
	}

	validSchema := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"age": {"type": "integer"}
		},
		"required": ["name", "age"]
	}`

	// Test with valid schema
	optsValid := &domain.ChatOptions{
		Model:         "gpt-3.5-turbo",
		SchemaContent: validSchema,
	}
	paramsValid := client.buildChatCompletionParams(msgs, optsValid)
	assert.NotNil(t, paramsValid.ResponseFormat.OfJSONSchema, "Expected OfJSONSchema to be set for valid schema")
	assert.Equal(t, "json_output", paramsValid.ResponseFormat.OfJSONSchema.JSONSchema.Name)
	assert.True(t, paramsValid.ResponseFormat.OfJSONSchema.JSONSchema.Strict.Value, "Expected strict to be true")
	assert.NotNil(t, paramsValid.ResponseFormat.OfJSONSchema.JSONSchema.Schema, "Expected schema content to be unmarshaled")

	// Test with invalid schema
	invalidSchema := `{ "type": "object", "properties": { "name": "string" }` // Malformed JSON
	optsInvalid := &domain.ChatOptions{
		Model:         "gpt-3.5-turbo",
		SchemaContent: invalidSchema,
	}
	paramsInvalid := client.buildChatCompletionParams(msgs, optsInvalid)
	assert.Nil(t, paramsInvalid.ResponseFormat.OfJSONSchema, "Expected OfJSONSchema to NOT be set for invalid schema")

	// Test without schema
	optsNoSchema := &domain.ChatOptions{
		Model: "gpt-3.5-turbo",
	}
	paramsNoSchema := client.buildChatCompletionParams(msgs, optsNoSchema)
	assert.Nil(t, paramsNoSchema.ResponseFormat.OfJSONSchema, "Expected OfJSONSchema to NOT be set when no schema is provided")
}

func TestBuildResponseParams_WithSchema(t *testing.T) {
	client := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Generate a JSON object."},
	}

	validSchema := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"age": {"type": "integer"}
		},
		"required": ["name", "age"]
	}`

	optsValid := &domain.ChatOptions{
		Model:         "gpt-4o", // Use a model that supports Responses API
		SchemaContent: validSchema,
	}
	paramsValid := client.buildResponseParams(msgs, optsValid)

	// Assert that extraFields is set
	assert.NotNil(t, paramsValid.ExtraFields(), "Expected ExtraFields to be set for valid schema")

	// Assert the "text" field within extraFields
	text, ok := paramsValid.ExtraFields()["text"].(map[string]interface{})
	assert.True(t, ok, "Expected 'text' field to be a map")
	assert.NotNil(t, text, "Expected 'text' map not to be nil")

	// Assert the "format" field within "text"
	format, ok := text["format"].(map[string]interface{})
	assert.True(t, ok, "Expected 'format' field to be a map")
	assert.NotNil(t, format, "Expected 'format' map not to be nil")

	// Assert content of "format"
	assert.Equal(t, "json_schema", format["type"], "Expected format type to be json_schema")
	assert.Equal(t, "json_output", format["name"], "Expected format name to be json_output")
	assert.True(t, format["strict"].(bool), "Expected strict to be true")

	// Assert schema content
	schema, ok := format["schema"].(map[string]interface{})
	assert.True(t, ok, "Expected 'schema' field to be a map")
	assert.NotNil(t, schema, "Expected schema content to be unmarshaled")
	assert.Equal(t, "object", schema["type"], "Expected schema type to be object")
	assert.NotNil(t, schema["properties"], "Expected schema properties to be present")
	assert.NotNil(t, schema["required"], "Expected schema required fields to be present")

	// Test with invalid schema
	invalidSchema := `{ "type": "object", "properties": { "name": "string" }` // Malformed JSON
	optsInvalid := &domain.ChatOptions{
		Model:         "gpt-4o",
		SchemaContent: invalidSchema,
	}
	paramsInvalid := client.buildResponseParams(msgs, optsInvalid)
	assert.Nil(t, paramsInvalid.ExtraFields(), "Expected ExtraFields NOT to be set for invalid schema")

	// Test without schema
	optsNoSchema := &domain.ChatOptions{
		Model: "gpt-4o",
	}
	paramsNoSchema := client.buildResponseParams(msgs, optsNoSchema)
	assert.Nil(t, paramsNoSchema.ExtraFields(), "Expected ExtraFields NOT to be set when no schema is provided")
}
