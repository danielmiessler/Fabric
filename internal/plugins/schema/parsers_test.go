package schema

import (
	"encoding/json"
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go"
	"google.golang.org/genai"
)

// Test data
const testSchema = `{
  "type": "object",
  "properties": {
    "name": {"type": "string"},
    "age": {"type": "number"}
  },
  "required": ["name", "age"]
}`

const testStructuredJSON = `{
  "name": "John Doe",
  "age": 30
}`

const invalidJSON = `{
  "name": "John Doe",
  "age": 
}`

func TestAnthropicParser_ParseResponse(t *testing.T) {
	parser := NewAnthropicParser()

	tests := []struct {
		name         string
		response     interface{}
		opts         *domain.ChatOptions
		expectedText string
		expectError  bool
	}{
		{
			name:         "String response without schema",
			response:     "Hello world",
			opts:         &domain.ChatOptions{},
			expectedText: "Hello world",
			expectError:  false,
		},
		{
			name:         "Structured output with mock tool",
			response:     createMockAnthropicToolResponse(testStructuredJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  false, // Will use fallback parsing
		},
		{
			name:         "Unsupported response type",
			response:     123,
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseResponse(tt.response, tt.opts)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && tt.expectedText != "" && result != tt.expectedText {
				t.Errorf("expected %q, got %q", tt.expectedText, result)
			}
		})
	}
}

func TestAnthropicParser_IsStructuredResponse(t *testing.T) {
	parser := NewAnthropicParser()

	tests := []struct {
		name     string
		response interface{}
		opts     *domain.ChatOptions
		expected bool
	}{
		{
			name:     "No schema content",
			response: createMockAnthropicToolResponse(testStructuredJSON),
			opts:     &domain.ChatOptions{},
			expected: false,
		},
		{
			name:     "With schema content",
			response: createMockAnthropicToolResponse(testStructuredJSON),
			opts:     &domain.ChatOptions{SchemaContent: testSchema},
			expected: true,
		},
		{
			name:     "String response with schema",
			response: "text response",
			opts:     &domain.ChatOptions{SchemaContent: testSchema},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.IsStructuredResponse(tt.response, tt.opts)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOpenAIParser_ParseResponse(t *testing.T) {
	parser := NewOpenAIParser()

	tests := []struct {
		name         string
		response     interface{}
		opts         *domain.ChatOptions
		expectedText string
		expectError  bool
	}{
		{
			name:         "String response",
			response:     "Hello world",
			opts:         &domain.ChatOptions{},
			expectedText: "Hello world",
			expectError:  false,
		},
		{
			name:         "Chat completion structured response",
			response:     createMockChatCompletion(testStructuredJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: testStructuredJSON,
			expectError:  false,
		},
		{
			name:         "Chat completion invalid JSON",
			response:     createMockChatCompletion(invalidJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true,
		},
		{
			name:         "Responses API response (fallback)",
			response:     createMockResponsesAPIResponse(testStructuredJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true, // Will fail since not actual SDK type
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseResponse(tt.response, tt.opts)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && result != tt.expectedText {
				t.Errorf("expected %q, got %q", tt.expectedText, result)
			}
		})
	}
}

func TestGeminiParser_ParseResponse(t *testing.T) {
	parser := NewGeminiParser()

	tests := []struct {
		name         string
		response     interface{}
		opts         *domain.ChatOptions
		expectedText string
		expectError  bool
	}{
		{
			name:         "String response",
			response:     "Hello world",
			opts:         &domain.ChatOptions{},
			expectedText: "Hello world",
			expectError:  false,
		},
		{
			name:         "Gemini structured response",
			response:     createMockGeminiResponse(testStructuredJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: testStructuredJSON,
			expectError:  false,
		},
		{
			name:         "Gemini invalid JSON",
			response:     createMockGeminiResponse(invalidJSON),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true,
		},
		{
			name:         "Empty Gemini response",
			response:     createMockGeminiResponse(""),
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseResponse(tt.response, tt.opts)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && result != tt.expectedText {
				t.Errorf("expected %q, got %q", tt.expectedText, result)
			}
		})
	}
}

func TestOllamaParser_ParseResponse(t *testing.T) {
	parser := NewOllamaParser()

	tests := []struct {
		name         string
		response     interface{}
		opts         *domain.ChatOptions
		expectedText string
		expectError  bool
	}{
		{
			name:         "String response without schema",
			response:     "Hello world",
			opts:         &domain.ChatOptions{},
			expectedText: "Hello world",
			expectError:  false,
		},
		{
			name:         "Structured JSON response",
			response:     testStructuredJSON,
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: testStructuredJSON,
			expectError:  false,
		},
		{
			name:         "Invalid JSON with schema",
			response:     invalidJSON,
			opts:         &domain.ChatOptions{SchemaContent: testSchema},
			expectedText: "",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseResponse(tt.response, tt.opts)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && result != tt.expectedText {
				t.Errorf("expected %q, got %q", tt.expectedText, result)
			}
		})
	}
}

// Helper functions to create mock responses

// Create mock responses using map[string]interface{} to simulate SDK types
func createMockAnthropicToolResponse(content string) map[string]interface{} {
	var toolInput interface{}
	if err := json.Unmarshal([]byte(content), &toolInput); err != nil {
		toolInput = map[string]interface{}{"result": content}
	}

	return map[string]interface{}{
		"content": []interface{}{
			map[string]interface{}{
				"type":  "tool_use",
				"name":  "get_structured_output",
				"input": toolInput,
			},
		},
	}
}

func createMockChatCompletion(content string) *openai.ChatCompletion {
	return &openai.ChatCompletion{
		ID:     "chatcmpl-123",
		Object: "chat.completion",
		Model:  "gpt-4",
		Choices: []openai.ChatCompletionChoice{
			{
				Index: 0,
				Message: openai.ChatCompletionMessage{
					Role:    "assistant",
					Content: content,
				},
				FinishReason: "stop",
			},
		},
	}
}

func createMockResponsesAPIResponse(content string) map[string]interface{} {
	return map[string]interface{}{
		"output": []interface{}{
			map[string]interface{}{
				"type": "message",
				"content": []interface{}{
					map[string]interface{}{
						"type": "output_text",
						"text": content,
					},
				},
			},
		},
	}
}

func createMockGeminiResponse(content string) *genai.GenerateContentResponse {
	part := &genai.Part{
		Text: content,
	}

	candidate := &genai.Candidate{
		Content: &genai.Content{
			Parts: []*genai.Part{part},
		},
	}

	return &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{candidate},
	}
}
