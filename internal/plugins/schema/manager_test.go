package schema

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
)

func TestNewManager(t *testing.T) {
	manager := NewManager()
	if manager == nil {
		t.Fatal("NewManager() returned nil")
	}
	if manager.plugin == nil {
		t.Fatal("Manager plugin is nil")
	}
}

func TestSupportsStructuredOutput(t *testing.T) {
	manager := NewManager()

	tests := []struct {
		provider string
		expected bool
	}{
		{"anthropic", true},
		{"openai", true},
		{"gemini", true},
		{"ollama", true},
		{"dryrun", true},
		{"perplexity", true},
		{"lmstudio", true},
		{"unknown", false},
	}

	for _, test := range tests {
		result := manager.SupportsStructuredOutput(test.provider)
		if result != test.expected {
			t.Errorf("SupportsStructuredOutput(%q) = %v, expected %v",
				test.provider, result, test.expected)
		}
	}
}

func TestHandleSchemaTransformation_NoSchema(t *testing.T) {
	manager := NewManager()
	opts := &domain.ChatOptions{
		SchemaContent: "", // No schema
	}

	err := manager.HandleSchemaTransformation("anthropic", opts)
	if err != nil {
		t.Errorf("Expected no error for empty schema, got: %v", err)
	}
}

func TestHandleSchemaTransformation_UnsupportedProvider(t *testing.T) {
	manager := NewManager()
	opts := &domain.ChatOptions{
		SchemaContent: `{"type": "object"}`,
	}

	err := manager.HandleSchemaTransformation("unsupported", opts)
	if err == nil {
		t.Error("Expected error for unsupported provider, got nil")
	}
}

func TestHandleSchemaTransformation_ValidSchema(t *testing.T) {
	manager := NewManager()
	opts := &domain.ChatOptions{
		SchemaContent: `{"type": "object", "properties": {"name": {"type": "string"}}}`,
	}

	err := manager.HandleSchemaTransformation("anthropic", opts)
	if err != nil {
		t.Errorf("Expected no error for valid schema transformation, got: %v", err)
	}

	if opts.TransformedSchema == nil {
		t.Error("Expected TransformedSchema to be set, got nil")
	}
}

func TestHandleSchemaTransformationWithContext_OpenAI(t *testing.T) {
	manager := NewManager()
	schema := `{"type": "object", "properties": {"name": {"type": "string"}}}`

	tests := []struct {
		name        string
		context     map[string]interface{}
		expectField string
	}{
		{
			name:        "Responses API",
			context:     map[string]interface{}{"api_type": "responses"},
			expectField: "text",
		},
		{
			name:        "Chat Completions API",
			context:     map[string]interface{}{"api_type": "completions"},
			expectField: "type", // Should have the base format
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := &domain.ChatOptions{SchemaContent: schema}

			err := manager.HandleSchemaTransformationWithContext("openai", opts, test.context)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if opts.TransformedSchema == nil {
				t.Fatal("TransformedSchema is nil")
			}

			transformed, ok := opts.TransformedSchema.(map[string]interface{})
			if !ok {
				t.Fatal("TransformedSchema is not a map")
			}

			if _, exists := transformed[test.expectField]; !exists {
				t.Errorf("Expected field '%s' not found in transformed schema", test.expectField)
			}
		})
	}
}

func TestOpenAI_ResponsesAPI_AdditionalProperties(t *testing.T) {
	manager := NewManager()

	tests := []struct {
		name                  string
		inputSchema           string
		expectAdditionalProps bool
		expectError           bool
	}{
		{
			name:                  "Schema without additionalProperties",
			inputSchema:           `{"type": "object", "properties": {"name": {"type": "string"}}}`,
			expectAdditionalProps: true, // Should be added
			expectError:           false,
		},
		{
			name:                  "Schema with additionalProperties: true",
			inputSchema:           `{"type": "object", "properties": {"name": {"type": "string"}}, "additionalProperties": true}`,
			expectAdditionalProps: false, // Should preserve existing value
			expectError:           false,
		},
		{
			name:                  "Schema with additionalProperties: false",
			inputSchema:           `{"type": "object", "properties": {"name": {"type": "string"}}, "additionalProperties": false}`,
			expectAdditionalProps: false, // Should preserve existing value
			expectError:           false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := &domain.ChatOptions{SchemaContent: test.inputSchema}
			context := map[string]interface{}{"api_type": "responses"}

			err := manager.HandleSchemaTransformationWithContext("openai", opts, context)
			if test.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !test.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if test.expectError {
				return
			}

			// Navigate to the actual schema in the transformed structure
			// Responses API: { "text": { "format": { "type": "json_schema", "name": "...", "schema": {...} } } }
			transformed, ok := opts.TransformedSchema.(map[string]interface{})
			if !ok {
				t.Fatal("TransformedSchema is not a map")
			}

			textSection, ok := transformed["text"].(map[string]interface{})
			if !ok {
				t.Fatal("text section is not a map")
			}

			formatSection, ok := textSection["format"].(map[string]interface{})
			if !ok {
				t.Fatal("format section is not a map")
			}

			actualSchema, ok := formatSection["schema"].(map[string]interface{})
			if !ok {
				t.Fatal("schema section is not a map")
			}

			// Check additionalProperties
			additionalProps, exists := actualSchema["additionalProperties"]
			if test.expectAdditionalProps {
				if !exists {
					t.Error("Expected additionalProperties to be set but it wasn't found")
				} else if additionalProps != false {
					t.Errorf("Expected additionalProperties to be false, got %v", additionalProps)
				}
			} else {
				// For schemas that already had additionalProperties, verify original value is preserved
				if test.inputSchema == `{"type": "object", "properties": {"name": {"type": "string"}}, "additionalProperties": true}` {
					if !exists || additionalProps != true {
						t.Errorf("Expected to preserve additionalProperties: true, got %v", additionalProps)
					}
				}
			}
		})
	}
}

func TestDryRun_SchemaTransformation(t *testing.T) {
	manager := NewManager()
	schema := `{"type": "object", "properties": {"name": {"type": "string"}}}`

	opts := &domain.ChatOptions{SchemaContent: schema}

	err := manager.HandleSchemaTransformation("dryrun", opts)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if opts.TransformedSchema == nil {
		t.Fatal("TransformedSchema is nil")
	}

	// For DryRun, the transformed schema should be the same as the original
	// since it just passes through for validation and display purposes
	transformedSchema, ok := opts.TransformedSchema.(map[string]interface{})
	if !ok {
		t.Fatal("TransformedSchema is not a map")
	}

	if transformedSchema["type"] != "object" {
		t.Errorf("Expected type 'object', got %v", transformedSchema["type"])
	}

	if _, exists := transformedSchema["properties"]; !exists {
		t.Error("Expected 'properties' field to exist in transformed schema")
	}
}

func TestPerplexity_SchemaTransformation(t *testing.T) {
	manager := NewManager()
	schema := `{"type": "object", "properties": {"name": {"type": "string"}}}`

	opts := &domain.ChatOptions{SchemaContent: schema}

	err := manager.HandleSchemaTransformation("perplexity", opts)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if opts.TransformedSchema == nil {
		t.Fatal("TransformedSchema is nil")
	}

	// For Perplexity, the transformed schema should be wrapped in the json_schema format
	transformedSchema, ok := opts.TransformedSchema.(map[string]interface{})
	if !ok {
		t.Fatal("TransformedSchema is not a map")
	}

	if transformedSchema["type"] != "json_schema" {
		t.Errorf("Expected type 'json_schema', got %v", transformedSchema["type"])
	}

	if jsonSchemaSection, exists := transformedSchema["json_schema"]; exists {
		if jsonSchemaMap, ok := jsonSchemaSection.(map[string]interface{}); ok {
			if actualSchema, exists := jsonSchemaMap["schema"]; exists {
				if schemaMap, ok := actualSchema.(map[string]interface{}); ok {
					if schemaMap["type"] != "object" {
						t.Errorf("Expected inner schema type 'object', got %v", schemaMap["type"])
					}
					if _, exists := schemaMap["properties"]; !exists {
						t.Error("Expected 'properties' field to exist in inner schema")
					}
				} else {
					t.Error("Expected inner schema to be a map")
				}
			} else {
				t.Error("Expected 'schema' field to exist in json_schema section")
			}
		} else {
			t.Error("Expected json_schema section to be a map")
		}
	} else {
		t.Error("Expected 'json_schema' field to exist in transformed schema")
	}
}

func TestValidateOutput(t *testing.T) {
	manager := NewManager()

	tests := []struct {
		name          string
		output        string
		schemaContent string
		expectError   bool
	}{
		{
			name:          "No schema",
			output:        "any output",
			schemaContent: "",
			expectError:   false,
		},
		{
			name:          "Valid JSON against schema",
			output:        `{"name": "John", "age": 30}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "number"}}, "required": ["name"]}`,
			expectError:   false,
		},
		{
			name:          "Invalid JSON against schema",
			output:        `{"name": "John", "age": "thirty"}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "number"}}, "required": ["name"]}`,
			expectError:   true,
		},
		{
			name:          "Missing required field",
			output:        `{"age": 30}`,
			schemaContent: `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "number"}}, "required": ["name"]}`,
			expectError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := manager.ValidateOutput(test.output, test.schemaContent, "test")
			if test.expectError && err == nil {
				t.Error("Expected validation error, got nil")
			}
			if !test.expectError && err != nil {
				t.Errorf("Expected no validation error, got: %v", err)
			}
		})
	}
}

func TestValidateOutput_DryRun(t *testing.T) {
	manager := NewManager()

	// Test that validation is skipped for dry run provider
	invalidOutput := `{"name": "John", "age": "thirty"}` // age should be number, not string
	schema := `{"type": "object", "properties": {"name": {"type": "string"}, "age": {"type": "number"}}, "required": ["name"]}`

	// Should fail validation for normal provider
	err := manager.ValidateOutput(invalidOutput, schema, "openai")
	if err == nil {
		t.Error("Expected validation error for normal provider, got nil")
	}

	// Should skip validation for dry run provider
	err = manager.ValidateOutput(invalidOutput, schema, "dryrun")
	if err != nil {
		t.Errorf("Expected no validation error for dryrun provider, got: %v", err)
	}
}
