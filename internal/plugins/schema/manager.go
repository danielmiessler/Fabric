package schema

import (
	"fmt"

	"github.com/danielmiessler/fabric/internal/domain"
)

// Manager orchestrates schema handling across different AI providers
type Manager struct {
	plugin SchemaPlugin
}

// NewManager creates a new schema manager with the default plugin
func NewManager() *Manager {
	return &Manager{
		plugin: NewDefaultSchemaPlugin(),
	}
}

// NewManagerWithPlugin creates a new schema manager with a custom plugin
func NewManagerWithPlugin(plugin SchemaPlugin) *Manager {
	return &Manager{
		plugin: plugin,
	}
}

// HandleSchemaTransformation prepares schema for provider-specific usage
// This should be called before sending requests to the AI provider
func (m *Manager) HandleSchemaTransformation(providerName string, opts *domain.ChatOptions) error {
	return m.HandleSchemaTransformationWithContext(providerName, opts, nil)
}

// HandleSchemaTransformationWithContext prepares schema with additional context
func (m *Manager) HandleSchemaTransformationWithContext(providerName string, opts *domain.ChatOptions, context map[string]interface{}) error {
	if opts.SchemaContent == "" {
		return nil // No schema to handle
	}

	// Check if provider supports structured outputs
	if !m.plugin.SupportsStructuredOutput(providerName) {
		return fmt.Errorf("provider '%s' does not support structured outputs", providerName)
	}

	// Transform schema for provider-specific format with context
	transformedSchema, err := m.plugin.TransformWithContext(opts.SchemaContent, providerName, context)
	if err != nil {
		return fmt.Errorf("failed to transform schema for provider '%s': %w", providerName, err)
	}

	// Store transformed schema in ChatOptions for provider to use
	opts.TransformedSchema = transformedSchema
	return nil
}

// HandleResponseParsing extracts structured output from provider-specific response format
// This should be called after receiving responses from the AI provider
func (m *Manager) HandleResponseParsing(providerName string, rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// If no schema or provider doesn't need custom parsing, return as-is
	if opts.SchemaContent == "" || !m.plugin.RequiresCustomParsing(providerName) {
		// For non-structured responses, we expect the provider to have already parsed the response
		if stringResponse, ok := rawResponse.(string); ok {
			return stringResponse, nil
		}
		return "", fmt.Errorf("expected string response for non-structured output")
	}

	// Use schema plugin for custom parsing
	parsedResponse, err := m.plugin.ParseResponse(rawResponse, providerName, opts)
	if err != nil {
		return "", fmt.Errorf("failed to parse structured response from provider '%s': %w", providerName, err)
	}

	return parsedResponse, nil
}

// HandleStreamResponseParsing extracts structured output from provider-specific stream events
func (m *Manager) HandleStreamResponseParsing(providerName string, rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// If no schema or provider doesn't need custom parsing, return as-is
	if opts.SchemaContent == "" || !m.plugin.RequiresCustomParsing(providerName) {
		if stringEvent, ok := rawEvent.(string); ok {
			return stringEvent, nil
		}
		return "", fmt.Errorf("expected string event for non-structured output")
	}

	// Use schema plugin for custom stream parsing
	parsedEvent, err := m.plugin.ParseStreamResponse(rawEvent, providerName, opts)
	if err != nil {
		return "", fmt.Errorf("failed to parse structured stream event from provider '%s': %w", providerName, err)
	}

	return parsedEvent, nil
}

// ValidateOutput validates the final output against the provided schema
func (m *Manager) ValidateOutput(output string, schemaContent string, providerName string) error {
	if schemaContent == "" {
		return nil // No schema to validate against
	}

	return m.plugin.Validate(output, schemaContent, providerName)
}

// GetProviderRequirements returns the schema requirements for a specific provider
func (m *Manager) GetProviderRequirements(providerName string) *ProviderRequirements {
	return m.plugin.GetProviderRequirements(providerName)
}

// SupportsStructuredOutput checks if a provider supports structured outputs
func (m *Manager) SupportsStructuredOutput(providerName string) bool {
	return m.plugin.SupportsStructuredOutput(providerName)
}
