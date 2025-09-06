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
			name:          "Invalid output against schema (missing required field)",
			output:        `{"name": "test"}`,
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
