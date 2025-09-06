package schema

import (
	"encoding/json"
	"fmt"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/xeipuuv/gojsonschema"
	"google.golang.org/genai"
)

// DefaultSchemaPlugin provides the default implementation of SchemaPlugin
type DefaultSchemaPlugin struct {
	*plugins.PluginBase
	providerConfigs map[string]*ProviderRequirements
	parsers         map[string]ResponseParser
}

// NewDefaultSchemaPlugin creates a new default schema plugin with built-in provider support
func NewDefaultSchemaPlugin() *DefaultSchemaPlugin {
	ret := &DefaultSchemaPlugin{
		providerConfigs: map[string]*ProviderRequirements{
			"anthropic": {
				RequiresTools:          true,
				RequiresResponseFormat: false,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          10000,
				ResponseType:           ResponseTypeTool,
				ToolName:               "get_structured_output",
				CustomParsingRequired:  false,
			},
			"openai": {
				RequiresTools:          false,
				RequiresResponseFormat: true,
				SupportedFormats:       []string{"json_schema", "json_object"},
				MaxSchemaSize:          50000,
				ResponseType:           ResponseTypeStructured,
				StreamEventType:        "ResponseOutputTextDelta",
				CustomParsingRequired:  true,
			},
			"gemini": {
				RequiresTools:          false,
				RequiresResponseFormat: true,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          25000,
				ResponseType:           ResponseTypeStructured,
				CustomParsingRequired:  true,
			},
			"ollama": {
				RequiresTools:          false,
				RequiresResponseFormat: true,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          15000,
				ResponseType:           ResponseTypeStructured,
				CustomParsingRequired:  false,
			},
			"dryrun": {
				RequiresTools:          false,
				RequiresResponseFormat: false,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          100000, // No real limits for dry run
				ResponseType:           ResponseTypeText,
				CustomParsingRequired:  false,
			},
			"perplexity": {
				RequiresTools:          false,
				RequiresResponseFormat: true,
				SupportedFormats:       []string{"json_schema", "regex"},
				MaxSchemaSize:          50000,
				ResponseType:           ResponseTypeStructured,
				CustomParsingRequired:  true,
			},
			"lmstudio": {
				RequiresTools:          false,
				RequiresResponseFormat: true,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          50000,
				ResponseType:           ResponseTypeStructured,
				CustomParsingRequired:  false,
			},
			"bedrock": {
				RequiresTools:          true,
				RequiresResponseFormat: false,
				SupportedFormats:       []string{"json_schema"},
				MaxSchemaSize:          100000, // Adjust as needed for Bedrock
				ResponseType:           ResponseTypeTool,
				CustomParsingRequired:  true,
			},
		},
		parsers: make(map[string]ResponseParser),
	}

	ret.PluginBase = &plugins.PluginBase{
		Name:          "Schema",
		EnvNamePrefix: "FABRIC_SCHEMA_",
	}

	// Initialize parsers
	ret.parsers["anthropic"] = NewAnthropicParser()
	ret.parsers["openai"] = NewOpenAIParser()
	ret.parsers["gemini"] = NewGeminiParser()
	ret.parsers["ollama"] = NewOllamaParser()
	ret.parsers["bedrock"] = NewBedrockParser()
	ret.parsers["perplexity"] = NewPerplexityParser()

	return ret
}

// Transform transforms schema content for provider-specific format requirements
func (sp *DefaultSchemaPlugin) Transform(schemaContent, providerName string) (interface{}, error) {
	return sp.TransformWithContext(schemaContent, providerName, nil)
}

// TransformWithContext transforms schema with additional context information
func (sp *DefaultSchemaPlugin) TransformWithContext(schemaContent, providerName string, context map[string]interface{}) (interface{}, error) {
	requirements := sp.GetProviderRequirements(providerName)
	if requirements == nil {
		return nil, fmt.Errorf("unsupported provider: %s", providerName)
	}

	// Validate schema is valid JSON
	var schema interface{}
	if err := json.Unmarshal([]byte(schemaContent), &schema); err != nil {
		return nil, fmt.Errorf("invalid JSON schema: %w", err)
	}

	// Check schema size limits
	if requirements.MaxSchemaSize > 0 && len(schemaContent) > requirements.MaxSchemaSize {
		return nil, fmt.Errorf("schema size (%d bytes) exceeds provider limit (%d bytes)",
			len(schemaContent), requirements.MaxSchemaSize)
	}

	// Transform based on provider requirements
	switch providerName {
	case "anthropic":
		return sp.transformForAnthropic(schema)
	case "openai":
		return sp.transformForOpenAI(schema, context)
	case "gemini":
		return sp.transformForGemini(schema)
	case "ollama":
		return sp.transformForOllama(schema)
	case "dryrun":
		return sp.transformForDryRun(schema)
	case "perplexity":
		return sp.transformForPerplexity(schema)
	case "lmstudio":
		return sp.transformForLMStudio(schema)
	case "bedrock":
		return sp.transformForBedrock(schema)
	default:
		// Default passthrough for unsupported providers
		return schema, nil
	}
}

// Validate validates output against the provided JSON schema
func (sp *DefaultSchemaPlugin) Validate(output, schemaContent string, providerName string) error {
	if schemaContent == "" {
		return nil // No schema to validate against
	}

	// Skip validation for dry run providers
	if providerName == "dryrun" {
		return nil
	}

	schemaLoader := gojsonschema.NewStringLoader(schemaContent)
	documentLoader := gojsonschema.NewStringLoader(output)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("error during schema validation: %w", err)
	}

	if !result.Valid() {
		errorString := "output failed schema validation:"
		for _, desc := range result.Errors() {
			errorString += fmt.Sprintf("\n- %s", desc)
		}
		return fmt.Errorf("%s", errorString)
	}

	return nil
}

// SupportsStructuredOutput checks if provider supports structured outputs
func (sp *DefaultSchemaPlugin) SupportsStructuredOutput(providerName string) bool {
	requirements := sp.GetProviderRequirements(providerName)
	return requirements != nil
}

// GetProviderRequirements returns the schema requirements for a specific provider
func (sp *DefaultSchemaPlugin) GetProviderRequirements(providerName string) *ProviderRequirements {
	return sp.providerConfigs[providerName]
}

// ParseResponse parses response using provider-specific logic
func (sp *DefaultSchemaPlugin) ParseResponse(rawResponse interface{}, providerName string, opts *domain.ChatOptions) (string, error) {
	parser, exists := sp.parsers[providerName]
	if !exists {
		return "", fmt.Errorf("no parser available for provider: %s", providerName)
	}

	return parser.ParseResponse(rawResponse, opts)
}

// ParseStreamResponse parses streaming response events using provider-specific logic
func (sp *DefaultSchemaPlugin) ParseStreamResponse(rawEvent interface{}, providerName string, opts *domain.ChatOptions) (string, error) {
	parser, exists := sp.parsers[providerName]
	if !exists {
		return "", fmt.Errorf("no parser available for provider: %s", providerName)
	}

	return parser.ParseStreamEvent(rawEvent, opts)
}

// RequiresCustomParsing checks if provider requires custom response parsing
func (sp *DefaultSchemaPlugin) RequiresCustomParsing(providerName string) bool {
	requirements := sp.GetProviderRequirements(providerName)
	return requirements != nil && requirements.CustomParsingRequired
}

// Provider-specific transformation methods

func (sp *DefaultSchemaPlugin) transformForAnthropic(schema interface{}) (interface{}, error) {
	// Anthropic uses a "tool" approach where the schema becomes the input schema for a tool
	return map[string]interface{}{
		"name":         "get_structured_output",
		"description":  "Generate structured output according to the provided schema",
		"input_schema": schema,
	}, nil
}

func (sp *DefaultSchemaPlugin) transformForOpenAI(schema interface{}, context map[string]interface{}) (interface{}, error) {
	// Default to Chat Completions API format (most OpenAI-compatible providers use this)
	// The specific OpenAI client can override this by providing context
	isResponsesAPI := false
	if context != nil {
		if apiType, exists := context["api_type"]; exists {
			if apiTypeStr, ok := apiType.(string); ok {
				isResponsesAPI = apiTypeStr == "responses"
			}
		}
	}

	// For Responses API, ensure additionalProperties: false is set
	processedSchema := schema
	if isResponsesAPI {
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			// Create a copy to avoid modifying the original
			schemaCopy := make(map[string]interface{})
			for k, v := range schemaMap {
				schemaCopy[k] = v
			}

			// Check if additionalProperties is already set
			if _, exists := schemaCopy["additionalProperties"]; !exists {
				schemaCopy["additionalProperties"] = false
			}
			processedSchema = schemaCopy
		}
	}

	if isResponsesAPI {
		// For Responses API: use text.format structure
		return map[string]interface{}{
			"text": map[string]interface{}{
				"format": map[string]interface{}{
					"type":   "json_schema",
					"name":   "structured_output",
					"schema": processedSchema,
					"strict": true,
				},
			},
		}, nil
	} else {
		// For Chat Completions API: use response_format structure (default)
		return map[string]interface{}{
			"type":   "json_schema",
			"name":   "structured_output",
			"schema": processedSchema,
			"strict": true,
		}, nil
	}
}

func (sp *DefaultSchemaPlugin) transformForGemini(schema interface{}) (interface{}, error) {
	// Convert the JSON schema to genai.Schema struct
	return sp.convertToGenaiSchema(schema)
}

// convertToGenaiSchema converts a JSON schema (map) to a genai.Schema struct
func (sp *DefaultSchemaPlugin) convertToGenaiSchema(schema interface{}) (*genai.Schema, error) {
	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("schema is not a JSON object")
	}

	genaiSchema := &genai.Schema{}

	// Convert common fields
	if title, exists := schemaMap["title"]; exists {
		if titleStr, ok := title.(string); ok {
			genaiSchema.Title = titleStr
		}
	}

	if description, exists := schemaMap["description"]; exists {
		if descStr, ok := description.(string); ok {
			genaiSchema.Description = descStr
		}
	}

	if schemaType, exists := schemaMap["type"]; exists {
		if typeStr, ok := schemaType.(string); ok {
			genaiSchema.Type = sp.convertToGenaiType(typeStr)
		}
	}

	// Handle properties for object types
	if properties, exists := schemaMap["properties"]; exists {
		if propMap, ok := properties.(map[string]interface{}); ok {
			genaiSchema.Properties = make(map[string]*genai.Schema)
			for propName, propSchema := range propMap {
				if convertedProp, err := sp.convertToGenaiSchema(propSchema); err == nil {
					genaiSchema.Properties[propName] = convertedProp
				}
			}
		}
	}

	// Handle required fields
	if required, exists := schemaMap["required"]; exists {
		if reqList, ok := required.([]interface{}); ok {
			genaiSchema.Required = make([]string, len(reqList))
			for i, req := range reqList {
				if reqStr, ok := req.(string); ok {
					genaiSchema.Required[i] = reqStr
				}
			}
		}
	}

	// Handle enum values
	if enum, exists := schemaMap["enum"]; exists {
		if enumList, ok := enum.([]interface{}); ok {
			genaiSchema.Enum = make([]string, len(enumList))
			for i, enumVal := range enumList {
				if enumStr, ok := enumVal.(string); ok {
					genaiSchema.Enum[i] = enumStr
				}
			}
		}
	}

	// Handle items for array types
	if items, exists := schemaMap["items"]; exists {
		if itemSchema, err := sp.convertToGenaiSchema(items); err == nil {
			genaiSchema.Items = itemSchema
		}
	}

	return genaiSchema, nil
}

// convertToGenaiType converts a JSON schema type string to genai.Type
func (sp *DefaultSchemaPlugin) convertToGenaiType(typeStr string) genai.Type {
	switch typeStr {
	case "string":
		return genai.TypeString
	case "integer":
		return genai.TypeInteger
	case "number":
		return genai.TypeNumber
	case "boolean":
		return genai.TypeBoolean
	case "array":
		return genai.TypeArray
	case "object":
		return genai.TypeObject
	default:
		return genai.TypeUnspecified
	}
}

func (sp *DefaultSchemaPlugin) transformForBedrock(schema interface{}) (interface{}, error) {
	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("bedrock schema is not a JSON object")
	}

	toolName, ok := schemaMap["title"].(string)
	if !ok || toolName == "" {
		return nil, fmt.Errorf("bedrock schema must contain a 'title' field for the tool name")
	}

	toolDescription, _ := schemaMap["description"].(string) // Description is optional

	// Return a generic map that represents the Bedrock ToolConfiguration structure
	// This will be marshaled into bedrockruntime/types.ToolConfiguration in bedrock.go
	return map[string]interface{}{
		"tools": []interface{}{
			map[string]interface{}{
				"toolSpec": map[string]interface{}{
					"name":        toolName,
					"description": toolDescription,
					"inputSchema": map[string]interface{}{
						"json": schemaMap, // Pass the original schema map here
					},
				},
			},
		},
	}, nil
}

func (sp *DefaultSchemaPlugin) transformForOllama(schema interface{}) (interface{}, error) {
	// Ollama uses format parameter with the raw JSON schema
	return map[string]interface{}{
		"format": schema,
	}, nil
}

func (sp *DefaultSchemaPlugin) transformForDryRun(schema interface{}) (interface{}, error) {
	// DryRun just validates and passes through the original schema for display purposes
	return schema, nil
}

func (sp *DefaultSchemaPlugin) transformForPerplexity(schema interface{}) (interface{}, error) {
	// Perplexity uses the format: { "type": "json_schema", "json_schema": { "schema": object } }
	// or { "type": "regex", "regex": { "regex": str } }
	// The schema parameter here is the actual schema object or regex string.
	// We need to determine if it's a JSON schema or a regex.

	// Attempt to unmarshal as a JSON object to check for schema structure
	if _, ok := schema.(map[string]interface{}); ok {
		return map[string]interface{}{
			"type": "json_schema",
			"json_schema": map[string]interface{}{
				"schema": schema,
			},
		}, nil
	} else if regexStr, ok := schema.(string); ok {
		// If it's a string, assume it's a regex pattern
		return map[string]interface{}{
			"type": "regex",
			"regex": map[string]interface{}{
				"regex": regexStr,
			},
		}, nil
	}

	return nil, fmt.Errorf("unsupported schema format for perplexity: expected JSON object or string, got %T", schema)
}

func (sp *DefaultSchemaPlugin) transformForLMStudio(schema interface{}) (interface{}, error) {
	// LMStudio uses OpenAI-compatible format for response_format
	return map[string]interface{}{
		"type": "json_schema",
		"json_schema": map[string]interface{}{
			"name":   "structured_output_schema",
			"strict": true,
			"schema": schema,
		},
	}, nil
}
