package schema

import (
	"github.com/danielmiessler/fabric/internal/domain"
)

// SchemaPlugin handles structured output schema processing for different AI providers
type SchemaPlugin interface {
	// Transform schema content for provider-specific format requirements
	Transform(schemaContent string, providerName string) (interface{}, error)

	// Transform schema with additional context (e.g., API type)
	TransformWithContext(schemaContent string, providerName string, context map[string]interface{}) (interface{}, error)

	// Validate output against the provided JSON schema
	Validate(output, schemaContent string, providerName string) error

	// Check if provider supports structured outputs
	SupportsStructuredOutput(providerName string) bool

	// Get provider-specific schema format requirements
	GetProviderRequirements(providerName string) *ProviderRequirements

	// Parse response using provider-specific logic
	ParseResponse(rawResponse interface{}, providerName string, opts *domain.ChatOptions) (string, error)

	// Parse streaming response events using provider-specific logic
	ParseStreamResponse(rawEvent interface{}, providerName string, opts *domain.ChatOptions) (string, error)

	// Check if provider requires custom response parsing for structured outputs
	RequiresCustomParsing(providerName string) bool
}

// ProviderRequirements defines how each provider handles structured outputs
type ProviderRequirements struct {
	// Schema transformation requirements
	RequiresTools          bool     // Anthropic uses "tools" approach
	RequiresResponseFormat bool     // OpenAI uses response_format
	SupportedFormats       []string // ["json_schema", "json_object", etc.]
	MaxSchemaSize          int      // Provider schema size limits

	// Response parsing requirements
	ResponseType          ResponseType
	ToolName              string // For tool-based providers (e.g., "get_structured_output")
	StreamEventType       string // For streaming (e.g., "ResponseOutputTextDelta")
	CustomParsingRequired bool   // Whether this provider needs custom parsing logic
}

// ResponseType indicates how the provider returns structured output responses
type ResponseType int

const (
	ResponseTypeText       ResponseType = iota // Standard text response
	ResponseTypeTool                           // Tool-based response (Anthropic)
	ResponseTypeStructured                     // Structured response format (OpenAI)
	ResponseTypeStream                         // Special streaming handling required
)

// ResponseParser handles provider-specific response parsing
type ResponseParser interface {
	// Parse a completed response from the provider
	ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error)

	// Parse a streaming event from the provider
	ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error)

	// Check if this response contains structured output
	IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool
}
