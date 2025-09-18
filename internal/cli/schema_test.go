package cli

import (
	"testing"
)

func TestValidateOutputWithSchema(t *testing.T) {
	tests := []struct {
		name          string
		output        string
		schemaContent string
		expectedErr   bool
		errMsg        string
	}{
		{
			name:          "Valid output and schema",
			output:        `{"name": "test", "age": 30}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   false,
		},
		{
			name:          "Valid output in json code block",
			output:        "```json\n{\"name\": \"test\", \"age\": 30}\n```",
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   false,
		},
		{
			name:          "Valid output in plain code block",
			output:        "```\n{\"name\": \"test\", \"age\": 30}\n```",
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   false,
		},
		{
			name:          "Valid output in code block without newlines",
			output:        "```json{\"name\": \"test\", \"age\": 30}```",
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   false,
		},
		{
			name:          "Invalid output against schema (missing required field)",
			output:        `{"name": "test"}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   true,
			errMsg:        "output failed schema validation:\n- (root): age is required",
		},
		{
			name:          "Invalid output in code block (missing required field)",
			output:        "```json\n{\"name\": \"test\"}\n```",
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   true,
			errMsg:        "output failed schema validation:\n- (root): age is required",
		},
		{
			name:          "Invalid output against schema (wrong type)",
			output:        `{"name": "test", "age": "thirty"}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "integer"}}, "required": ["name", "age"]}`,
			expectedErr:   true,
			errMsg:        "output failed schema validation:\n- age: Invalid type. Expected: integer, given: string",
		},
		{
			name:          "Empty schema content",
			output:        `{"name": "test"}`,
			schemaContent: ``,
			expectedErr:   false, // No validation should occur
		},
		{
			name:          "Invalid JSON output",
			output:        `{"name": "test", "age":}`, // Malformed JSON
			schemaContent: `{"type": "object"}`,
			expectedErr:   true,
			errMsg:        "error during schema validation", // Expecting an error from gojsonschema.Validate
		},
		{
			name:          "Invalid JSON in code block",
			output:        "```json\n{\"name\": \"test\", \"age\":}\n```", // Malformed JSON in code block
			schemaContent: `{"type": "object"}`,
			expectedErr:   true,
			errMsg:        "error during schema validation", // Expecting an error from gojsonschema.Validate
		},
		{
			name:          "Invalid JSON schema content",
			output:        `{"name": "test"}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"`, // Malformed JSON schema
			expectedErr:   true,
			errMsg:        "error during schema validation", // Expecting an error from gojsonschema.Validate
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOutputWithSchema(tt.output, tt.schemaContent)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected an error but got none")
				} else if tt.errMsg != "" && err.Error() != tt.errMsg && !contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error message to contain '%s', but got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got: %v", err)
				}
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}

func TestExtractJSONFromCodeBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "JSON in json code block with newlines",
			input:    "```json\n{\"key\": \"value\"}\n```",
			expected: "{\"key\": \"value\"}",
		},
		{
			name:     "JSON in plain code block with newlines",
			input:    "```\n{\"key\": \"value\"}\n```",
			expected: "{\"key\": \"value\"}",
		},
		{
			name:     "JSON in code block without newlines",
			input:    "```json{\"key\": \"value\"}```",
			expected: "{\"key\": \"value\"}",
		},
		{
			name:     "Raw JSON without code block",
			input:    "{\"key\": \"value\"}",
			expected: "{\"key\": \"value\"}",
		},
		{
			name:     "JSON with extra whitespace in code block",
			input:    "```json\n\n  {\"key\": \"value\"}  \n\n```",
			expected: "{\"key\": \"value\"}",
		},
		{
			name:     "Complex nested JSON in code block",
			input:    "```json\n{\n  \"nested\": {\n    \"array\": [1, 2, 3]\n  }\n}\n```",
			expected: "{\n  \"nested\": {\n    \"array\": [1, 2, 3]\n  }\n}",
		},
		{
			name:     "Text with no code block or JSON",
			input:    "This is plain text without JSON",
			expected: "This is plain text without JSON",
		},
		{
			name:     "Multiple code blocks (only first is extracted)",
			input:    "```json\n{\"first\": true}\n```\n\n```json\n{\"second\": true}\n```",
			expected: "{\"first\": true}",
		},
		{
			name:     "JSON with text before code block",
			input:    "Here is the JSON output:\n```json\n{\"result\": \"success\"}\n```",
			expected: "{\"result\": \"success\"}",
		},
		{
			name:     "JSON with text after code block",
			input:    "```json\n{\"data\": 123}\n```\nThat's your structured output!",
			expected: "{\"data\": 123}",
		},
		{
			name:     "JSON with text before and after code block",
			input:    "Processing complete. Here's the result:\n\n```json\n{\"status\": \"ok\"}\n```\n\nPlease review the output above.",
			expected: "{\"status\": \"ok\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSONFromCodeBlock(tt.input)
			if result != tt.expected {
				t.Errorf("extractJSONFromCodeBlock() = %q, want %q", result, tt.expected)
			}
		})
	}
}
