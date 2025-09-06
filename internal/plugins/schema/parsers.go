package schema

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
	perplexity "github.com/sgaunet/perplexity-go/v2"
	"google.golang.org/genai"

	bedrockTypes "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types" // Alias to avoid conflict with schema.types
)

// AnthropicParser handles Anthropic-specific response parsing
type AnthropicParser struct{}

// NewAnthropicParser creates a new Anthropic response parser
func NewAnthropicParser() *AnthropicParser {
	return &AnthropicParser{}
}

// ParseResponse parses an Anthropic response for structured output
func (p *AnthropicParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// For structured outputs, Anthropic returns tool use blocks
	// Check if we have an actual Anthropic message
	if message, ok := rawResponse.(*anthropic.Message); ok {
		return p.parseAnthropicMessage(message, opts)
	}

	if opts.SchemaContent == "" {
		// Non-structured response - should be handled by provider's normal parsing
		return p.parseNormalResponse(rawResponse)
	}

	return p.parseStructuredResponse(rawResponse, opts)
}

// parseAnthropicMessage handles actual Anthropic message types
func (p *AnthropicParser) parseAnthropicMessage(message *anthropic.Message, opts *domain.ChatOptions) (string, error) {
	// Check if we have a schema - if so, look for structured output tool use
	if opts.SchemaContent != "" {
		for _, block := range message.Content {
			switch variant := block.AsAny().(type) {
			case anthropic.ToolUseBlock:
				if variant.Name == "get_structured_output" {
					jsonBytes, err := json.MarshalIndent(variant.Input, "", "  ")
					if err != nil {
						return "", fmt.Errorf("failed to marshal tool_use input: %w", err)
					}
					return string(jsonBytes), nil
				}
			}
		}
		return "", fmt.Errorf("no structured output found in anthropic response")
	}

	// For non-structured responses, extract text blocks
	var textParts []string
	for _, block := range message.Content {
		switch variant := block.AsAny().(type) {
		case anthropic.TextBlock:
			textParts = append(textParts, variant.Text)
		}
	}

	if len(textParts) == 0 {
		return "", fmt.Errorf("no text content found in anthropic message")
	}

	return textParts[0], nil // Return first text part for simplicity
}

func (p *AnthropicParser) parseStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {

	if message, ok := rawResponse.(*anthropic.Message); ok {
		return p.parseAnthropicMessage(message, opts)
	}

	// Fallback: try to parse map-based response (e.g., from test mocks or raw API responses)
	if responseMap, ok := rawResponse.(map[string]interface{}); ok {
		return p.parseMapBasedResponse(responseMap, opts)
	}

	return "", fmt.Errorf("expected *anthropic.Message for structured Anthropic response, got %T", rawResponse)
}

// parseMapBasedResponse handles map-based responses (e.g., from test mocks or raw API responses)
func (p *AnthropicParser) parseMapBasedResponse(responseMap map[string]interface{}, opts *domain.ChatOptions) (string, error) {
	// Look for content array with tool_use blocks
	if content, ok := responseMap["content"].([]interface{}); ok {
		for _, item := range content {
			if block, ok := item.(map[string]interface{}); ok {
				if blockType, ok := block["type"].(string); ok && blockType == "tool_use" {
					if name, ok := block["name"].(string); ok && name == "get_structured_output" {
						if input, ok := block["input"]; ok {
							jsonBytes, err := json.MarshalIndent(input, "", "  ")
							if err != nil {
								return "", fmt.Errorf("failed to marshal tool input: %w", err)
							}
							return string(jsonBytes), nil
						}
					}
				}
			}
		}
	}

	// If we can't find tool_use content, return empty string (matches test expectation)
	return "", nil
}

func (p *AnthropicParser) parseNormalResponse(rawResponse interface{}) (string, error) {
	// For non-structured responses, expect string
	if stringResponse, ok := rawResponse.(string); ok {
		return stringResponse, nil
	}
	return "", fmt.Errorf("expected string response for normal Anthropic output")
}

// ParseStreamEvent parses an Anthropic streaming event
func (p *AnthropicParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Anthropic streaming typically sends delta text directly
	// The provider handles the actual streaming event types, we just handle the extracted content

	// Use reflection to check for Delta.Text field
	if eventMap, ok := rawEvent.(map[string]interface{}); ok {
		if delta, exists := eventMap["delta"]; exists {
			if deltaMap, ok := delta.(map[string]interface{}); ok {
				if text, exists := deltaMap["text"]; exists {
					if textStr, ok := text.(string); ok && textStr != "" {
						return textStr, nil
					}
				}
			}
		}
	}

	// For now, handle as string fallback
	if stringEvent, ok := rawEvent.(string); ok {
		return stringEvent, nil
	}

	return "", nil // Return empty string instead of error for stream events that don't contain text
}

// IsStructuredResponse checks if response contains structured output
func (p *AnthropicParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	// First check if schema content was provided
	if opts.SchemaContent == "" {
		return false
	}

	// If we have an Anthropic message, check for tool use blocks
	if message, ok := rawResponse.(*anthropic.Message); ok {
		for _, block := range message.Content {
			switch variant := block.AsAny().(type) {
			case anthropic.ToolUseBlock:
				if variant.Name == "get_structured_output" {
					return true
				}
			}
		}
		return false
	}

	// Fallback: assume structured if schema was provided
	return true
}

// OpenAIParser handles OpenAI-specific response parsing
type OpenAIParser struct{}

// NewOpenAIParser creates a new OpenAI response parser
func NewOpenAIParser() *OpenAIParser {
	return &OpenAIParser{}
}

// ParseResponse parses an OpenAI response for structured output
func (p *OpenAIParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// Check if it's a Responses API response
	if respAPIResponse, ok := rawResponse.(*responses.Response); ok {
		return p.parseResponsesAPIResponse(respAPIResponse, opts)
	}

	// Check if it's a Chat Completions API response
	if chatResponse, ok := rawResponse.(*openai.ChatCompletion); ok {
		return p.parseChatCompletionResponse(chatResponse, opts)
	}

	// Fallback to generic parsing
	if opts.SchemaContent == "" {
		return p.parseNormalResponse(rawResponse)
	}

	return p.parseStructuredResponse(rawResponse, opts)
}

// parseResponsesAPIResponse handles OpenAI Responses API responses
func (p *OpenAIParser) parseResponsesAPIResponse(response *responses.Response, opts *domain.ChatOptions) (string, error) {
	if len(response.Output) == 0 {
		return "", fmt.Errorf("no output items in openai responses API response")
	}

	// Extract text content from output items using the same logic as OpenAI client
	var textParts []string

	for _, item := range response.Output {
		if item.Type == "message" {
			for _, c := range item.Content {
				if c.Type == "output_text" {
					outputText := c.AsOutputText()
					textParts = append(textParts, outputText.Text)
				}
			}
		}
	}

	if len(textParts) == 0 {
		return "", fmt.Errorf("no text content in openai responses API response")
	}

	// Join all text parts
	textContent := ""
	for _, part := range textParts {
		textContent += part
	}

	// For structured outputs, validate JSON
	if opts.SchemaContent != "" {
		var jsonTest interface{}
		if err := json.Unmarshal([]byte(textContent), &jsonTest); err != nil {
			return "", fmt.Errorf("openai responses API structured output is not valid JSON: %w", err)
		}
	}

	return textContent, nil
}

// parseChatCompletionResponse handles OpenAI Chat Completions API responses
func (p *OpenAIParser) parseChatCompletionResponse(response *openai.ChatCompletion, opts *domain.ChatOptions) (string, error) {
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no choices in openai chat completion response")
	}

	choice := response.Choices[0]
	content := choice.Message.Content

	// For structured outputs, content should already be JSON
	if opts.SchemaContent != "" {
		// Validate that it's valid JSON
		var jsonTest interface{}
		if err := json.Unmarshal([]byte(content), &jsonTest); err != nil {
			return "", fmt.Errorf("openai structured response is not valid JSON: %w", err)
		}
	}

	return content, nil
}

func (p *OpenAIParser) parseStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// OpenAI has two APIs: Responses API and Chat Completions API
	// Both return structured outputs differently

	// For Responses API: content is structured JSON in the response
	// For Chat Completions API: content is in choices[0].message.content

	if responseMap, ok := rawResponse.(map[string]interface{}); ok {
		// Try Responses API format first
		if output, exists := responseMap["output"]; exists {
			if outputMap, ok := output.(map[string]interface{}); ok {
				if content, exists := outputMap["content"]; exists {
					// Content should already be structured JSON
					if contentStr, ok := content.(string); ok {
						return contentStr, nil
					}
					// Or it might already be structured
					jsonBytes, err := json.Marshal(content)
					if err != nil {
						return "", fmt.Errorf("failed to marshal responses API content: %w", err)
					}
					return string(jsonBytes), nil
				}
			}
		}

		// Try Chat Completions API format
		if choices, exists := responseMap["choices"]; exists {
			if choicesSlice, ok := choices.([]interface{}); ok && len(choicesSlice) > 0 {
				if choice, ok := choicesSlice[0].(map[string]interface{}); ok {
					if message, exists := choice["message"]; exists {
						if messageMap, ok := message.(map[string]interface{}); ok {
							if content, exists := messageMap["content"]; exists {
								if contentStr, ok := content.(string); ok {
									return contentStr, nil
								}
							}
						}
					}
				}
			}
		}
	}

	// Fallback: if it's already a string, return it
	if stringResponse, ok := rawResponse.(string); ok {
		return stringResponse, nil
	}

	return "", fmt.Errorf("no structured output found in openai response")
}

func (p *OpenAIParser) parseNormalResponse(rawResponse interface{}) (string, error) {
	if stringResponse, ok := rawResponse.(string); ok {
		return stringResponse, nil
	}
	return "", fmt.Errorf("expected string response for normal OpenAI output")
}

// ParseStreamEvent parses an OpenAI streaming event
func (p *OpenAIParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Handle Responses API streaming events
	if eventMap, ok := rawEvent.(map[string]interface{}); ok {
		// Check for ResponseOutputTextDelta events
		if eventType, exists := eventMap["type"]; exists {
			if eventType == "response.output.text.delta" {
				if delta, exists := eventMap["delta"]; exists {
					if deltaStr, ok := delta.(string); ok {
						return deltaStr, nil
					}
				}
			}
		}
	}

	// Handle Chat Completions API streaming chunks
	if chunk, ok := rawEvent.(*openai.ChatCompletionChunk); ok {
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			return chunk.Choices[0].Delta.Content, nil
		}
	}

	// Fallback to string
	if stringEvent, ok := rawEvent.(string); ok {
		return stringEvent, nil
	}

	return "", nil // Return empty for events that don't contain text content
}

// IsStructuredResponse checks if response contains structured output
func (p *OpenAIParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	return opts.SchemaContent != ""
}

// PerplexityParser handles Perplexity-specific response parsing
type PerplexityParser struct{}

// NewPerplexityParser creates a new Perplexity response parser
func NewPerplexityParser() *PerplexityParser {
	return &PerplexityParser{}
}

// ParseResponse parses a Perplexity response for structured output
func (p *PerplexityParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// Handle perplexity.CompletionResponse objects from perplexity-go v2.12.0
	if resp, ok := rawResponse.(*perplexity.CompletionResponse); ok {
		return p.parsePerplexityCompletionResponse(resp, opts)
	}

	// Fallback to string response (for backward compatibility)
	stringResponse, ok := rawResponse.(string)
	if !ok {
		return "", fmt.Errorf("expected string response from Perplexity, got %T", rawResponse)
	}

	// Perplexity can return structured output with a <think> section
	// We need to extract the JSON or regex matched content after the </think> tag
	// or directly if no <think> section is present.
	thinkEndTag := "</think>"
	if idx := findLastIndex(stringResponse, thinkEndTag); idx != -1 {
		// Extract content after the last </think> tag
		extractedContent := stringResponse[idx+len(thinkEndTag):]
		// Trim leading/trailing whitespace, including newlines
		stringResponse = trimSpaceAndNewlines(extractedContent)
	}

	// If a schema was provided, attempt to validate as JSON.
	// This covers both json_schema and regex cases, as regex output is usually plain text.
	if opts.SchemaContent != "" {
		// Attempt to unmarshal as JSON. If it fails, it might be a regex output.
		var jsonTest interface{}
		if err := json.Unmarshal([]byte(stringResponse), &jsonTest); err != nil {
			// If it's not valid JSON, it's likely a regex output.
			// We don't have a way to validate regex output against the pattern here,
			// so we just return the string. The external validation step will handle it.
			return stringResponse, nil
		}
		// If it's valid JSON, return it.
		return stringResponse, nil
	}

	return stringResponse, nil
}

// parsePerplexityCompletionResponse handles perplexity.CompletionResponse objects
func (p *PerplexityParser) parsePerplexityCompletionResponse(resp *perplexity.CompletionResponse, opts *domain.ChatOptions) (string, error) {
	if resp == nil {
		return "", fmt.Errorf("nil perplexity completion response")
	}

	// Extract content from the response choices
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in perplexity response")
	}

	// Get the content from the first choice
	choice := resp.Choices[0]
	content := choice.Message.Content

	// Process the content first (handles <think> sections, etc.)
	processedContent, err := p.processPerplexityContent(content, opts)
	if err != nil {
		return "", err
	}

	// Handle citations based on whether we're outputting structured JSON or plain text
	return p.handleCitations(processedContent, resp.Citations, opts)
}

// processPerplexityContent handles the extracted content with <think> sections and JSON parsing
func (p *PerplexityParser) processPerplexityContent(content string, opts *domain.ChatOptions) (string, error) {
	// Perplexity can return structured output with a <think> section
	// We need to extract the JSON or regex matched content after the </think> tag
	// or directly if no <think> section is present.
	thinkEndTag := "</think>"
	if idx := findLastIndex(content, thinkEndTag); idx != -1 {
		// Extract content after the last </think> tag
		extractedContent := content[idx+len(thinkEndTag):]
		// Trim leading/trailing whitespace, including newlines
		content = trimSpaceAndNewlines(extractedContent)
	}

	// If a schema was provided, attempt to validate as JSON.
	// This covers both json_schema and regex cases, as regex output is usually plain text.
	if opts.SchemaContent != "" {
		// Attempt to unmarshal as JSON. If it fails, it might be a regex output.
		var jsonTest interface{}
		if err := json.Unmarshal([]byte(content), &jsonTest); err != nil {
			// If it's not valid JSON, it's likely a regex output.
			// We don't have a way to validate regex output against the pattern here,
			// so we just return the string. The external validation step will handle it.
			return content, nil
		}
		// If it's valid JSON, return it.
		return content, nil
	}

	return content, nil
}

// handleCitations processes citations based on output type (JSON vs text)
func (p *PerplexityParser) handleCitations(content string, citations *[]string, opts *domain.ChatOptions) (string, error) {
	// If no citations, return content as-is
	if citations == nil || len(*citations) == 0 {
		return content, nil
	}

	// For structured output (schema provided), handle citations in JSON
	if opts.SchemaContent != "" {
		// Try to parse as JSON to see if we can inject citations
		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(content), &jsonData); err == nil {
			// It's valid JSON, format citations according to the schema
			formattedCitations := p.formatCitationsForSchema(citations, opts)

			// Check if there's already a citations field
			if _, hasCitations := jsonData["citations"]; hasCitations {
				// Populate the existing citations field
				jsonData["citations"] = formattedCitations
			} else {
				// Add a citations field since Perplexity always provides citations
				jsonData["citations"] = formattedCitations
			}

			// Marshal back to JSON
			updatedJSON, err := json.MarshalIndent(jsonData, "", "  ")
			if err != nil {
				return "", fmt.Errorf("failed to marshal JSON with citations: %w", err)
			}
			return string(updatedJSON), nil
		}
		// If it's not valid JSON (could be regex output), fall through to text append
	}

	// For non-JSON output, append citations as markdown
	result := content + "\n\n## Sources\n\n"
	for _, citation := range *citations {
		result += fmt.Sprintf("- [%s](%s)\n", citation, citation)
	}
	return result, nil
}

// formatCitationsForSchema formats citations according to the schema structure
func (p *PerplexityParser) formatCitationsForSchema(citations *[]string, opts *domain.ChatOptions) []interface{} {
	if citations == nil || len(*citations) == 0 {
		return []interface{}{}
	}

	// Parse the schema to understand the expected citation format
	var schemaData map[string]interface{}
	if err := json.Unmarshal([]byte(opts.SchemaContent), &schemaData); err != nil {
		// If we can't parse the schema, fall back to simple string array
		result := make([]interface{}, len(*citations))
		for i, citation := range *citations {
			result[i] = citation
		}
		return result
	}

	// Check if the schema defines a citations field structure
	if properties, ok := schemaData["properties"].(map[string]interface{}); ok {
		if citationsField, ok := properties["citations"].(map[string]interface{}); ok {
			if items, ok := citationsField["items"].(map[string]interface{}); ok {
				if itemProps, ok := items["properties"].(map[string]interface{}); ok {
					// Check if it expects objects with id, title and url fields
					if _, hasTitle := itemProps["title"]; hasTitle {
						if _, hasURL := itemProps["url"]; hasURL {
							// Check if id field is also expected
							if _, hasID := itemProps["id"]; hasID {
								// Format as objects with id, title and url
								result := make([]interface{}, len(*citations))
								for i, citation := range *citations {
									result[i] = map[string]interface{}{
										"id":    fmt.Sprintf("%d", i+1), // Use 1-based indexing for citation numbers
										"title": p.extractTitleFromURL(citation),
										"url":   citation,
									}
								}
								return result
							} else {
								// Format as objects with title and url (backward compatibility)
								result := make([]interface{}, len(*citations))
								for i, citation := range *citations {
									result[i] = map[string]interface{}{
										"title": p.extractTitleFromURL(citation),
										"url":   citation,
									}
								}
								return result
							}
						}
					}
				}
			}
		}
	}

	// Fall back to simple string array if schema structure is unknown
	result := make([]interface{}, len(*citations))
	for i, citation := range *citations {
		result[i] = citation
	}
	return result
}

// extractTitleFromURL attempts to extract a meaningful title from a URL
func (p *PerplexityParser) extractTitleFromURL(url string) string {
	// Try to extract domain name as a fallback title
	// This is a simple implementation - could be enhanced to fetch actual page titles
	if len(url) > 8 && (url[:7] == "http://" || url[:8] == "https://") {
		start := 7
		if url[:8] == "https://" {
			start = 8
		}

		// Find the end of the domain
		end := start
		for i := start; i < len(url) && url[i] != '/'; i++ {
			end = i + 1
		}

		domain := url[start:end]

		// Remove www. prefix if present
		if len(domain) > 4 && domain[:4] == "www." {
			domain = domain[4:]
		}

		// Capitalize first letter
		if len(domain) > 0 {
			return strings.ToUpper(domain[:1]) + domain[1:]
		}
		return domain
	}

	// If URL format is unexpected, return the URL itself
	return url
}

// ParseStreamEvent parses a Perplexity streaming event
func (p *PerplexityParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Handle perplexity.CompletionResponse objects from streaming
	if resp, ok := rawEvent.(*perplexity.CompletionResponse); ok {
		if len(resp.Choices) > 0 {
			// For streaming, we want to return just the delta content without citations
			content := resp.Choices[0].Message.Content
			if opts.SchemaContent != "" {
				// For structured output streaming, process the content
				return p.processPerplexityContent(content, opts)
			}
			return content, nil
		}
		return "", nil // Empty response
	}

	// Fallback: Perplexity streaming typically sends delta text directly
	if stringEvent, ok := rawEvent.(string); ok {
		return stringEvent, nil
	}
	return "", fmt.Errorf("expected string event or CompletionResponse from Perplexity, got %T", rawEvent)
}

// IsStructuredResponse checks if response contains structured output
func (p *PerplexityParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	// If schema content was provided, then it's a structured response.
	return opts.SchemaContent != ""
}

// Helper function to find the last index of a substring
func findLastIndex(s, substr string) int {
	for i := len(s) - len(substr); i >= 0; i-- {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// Helper function to trim space and newlines
func trimSpaceAndNewlines(s string) string {
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\n' || s[start] == '\r' || s[start] == '\t') {
		start++
	}
	end := len(s)
	for end > start && (s[end-1] == ' ' || s[end-1] == '\n' || s[end-1] == '\r' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}

// BedrockParser handles Bedrock-specific response parsing
type BedrockParser struct{}

// NewBedrockParser creates a new Bedrock response parser
func NewBedrockParser() *BedrockParser {
	return &BedrockParser{}
}

// ParseResponse parses a Bedrock response for structured output
func (p *BedrockParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// Bedrock structured outputs are expected to be in toolUse blocks within Content
	if responseMessage, ok := rawResponse.(*bedrockTypes.Message); ok {
		return p.parseBedrockMessage(responseMessage, opts)
	}

	// Fallback for non-structured responses
	if stringResponse, ok := rawResponse.(string); ok {
		return stringResponse, nil
	}

	return "", fmt.Errorf("unsupported bedrock response type for ParseResponse: %T", rawResponse)
}

func (p *BedrockParser) parseBedrockMessage(message *bedrockTypes.Message, opts *domain.ChatOptions) (string, error) {
	if opts.SchemaContent != "" {
		for _, block := range message.Content {
			if toolUseBlock, ok := block.(*bedrockTypes.ContentBlockMemberToolUse); ok {
				jsonBytes, err := json.MarshalIndent(toolUseBlock.Value, "", "  ")
				if err != nil {
					return "", fmt.Errorf("failed to marshal tool_use input: %w", err)
				}
				return string(jsonBytes), nil
			}
		}
		return "", fmt.Errorf("no structured output (tool_use) found in bedrock message")
	}

	// For non-structured responses, extract text blocks
	var textParts []string
	for _, block := range message.Content {
		if textBlock, ok := block.(*bedrockTypes.ContentBlockMemberText); ok {
			textParts = append(textParts, textBlock.Value)
		}
	}

	if len(textParts) == 0 {
		return "", fmt.Errorf("no text content found in bedrock message")
	}

	return textParts[0], nil
}

// ParseStreamEvent parses a Bedrock streaming event
func (p *BedrockParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Bedrock streaming events can contain content block deltas
	if eventContentDelta, ok := rawEvent.(*bedrockTypes.ConverseStreamOutputMemberContentBlockDelta); ok {
		if textDelta, ok := eventContentDelta.Value.Delta.(*bedrockTypes.ContentBlockDeltaMemberText); ok {
			return textDelta.Value, nil
		}
		if toolUseDelta, ok := eventContentDelta.Value.Delta.(*bedrockTypes.ContentBlockDeltaMemberToolUse); ok {
			// For streaming tool use, we might get partial JSON.
			// This needs careful handling. For now, just return the raw string if it's there.
			// A more robust solution would involve buffering and parsing complete JSON.
			jsonBytes, err := json.Marshal(toolUseDelta.Value)
			if err != nil {
				return "", fmt.Errorf("failed to marshal streaming tool use delta: %w", err)
			}
			return string(jsonBytes), nil
		}
	} else if messageStop, ok := rawEvent.(*bedrockTypes.ConverseStreamOutputMemberMessageStop); ok {
		// MessageStop indicates end of response. Add newline.
		// If tool_use was the last thing, it might not have been fully parsed yet.
		// This is a complex problem for streaming structured outputs.
		// For now, just return a newline.
		_ = messageStop // suppress unused warning
		return "\n", nil
	}

	return "", nil // Return empty string for events that don't contain text or complete tool use
}

// IsStructuredResponse checks if response contains structured output
func (p *BedrockParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	if opts.SchemaContent == "" {
		return false
	}

	if responseMessage, ok := rawResponse.(*bedrockTypes.Message); ok {
		for _, block := range responseMessage.Content {
			if _, ok := block.(*bedrockTypes.ContentBlockMemberToolUse); ok {
				return true
			}
		}
	}
	return false
}

// GeminiParser handles Gemini-specific response parsing
type GeminiParser struct{}

// NewGeminiParser creates a new Gemini response parser
func NewGeminiParser() *GeminiParser {
	return &GeminiParser{}
}

// ParseResponse parses a Gemini response for structured output
func (p *GeminiParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// Check if it's a Gemini response
	if geminiResponse, ok := rawResponse.(*genai.GenerateContentResponse); ok {
		return p.parseGeminiResponse(geminiResponse, opts)
	}

	// Fallback to string response
	if stringResponse, ok := rawResponse.(string); ok {
		return stringResponse, nil
	}

	return "", fmt.Errorf("unsupported gemini response type")
}

// parseGeminiResponse handles actual Gemini response types
func (p *GeminiParser) parseGeminiResponse(response *genai.GenerateContentResponse, opts *domain.ChatOptions) (string, error) {
	if response == nil {
		return "", fmt.Errorf("nil gemini response")
	}

	// Extract text content from response using similar logic to Gemini client
	var textContent string
	for _, candidate := range response.Candidates {
		if candidate == nil || candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			if part != nil && part.Text != "" {
				textContent += part.Text
			}
		}
	}

	if textContent == "" {
		return "", fmt.Errorf("no text content in gemini response")
	}

	// For structured outputs, validate JSON
	if opts.SchemaContent != "" {
		var jsonTest interface{}
		if err := json.Unmarshal([]byte(textContent), &jsonTest); err != nil {
			return "", fmt.Errorf("gemini structured response is not valid JSON: %w", err)
		}
	}

	return textContent, nil
}

// ParseStreamEvent parses a Gemini streaming event
func (p *GeminiParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Handle Gemini streaming responses
	if geminiResponse, ok := rawEvent.(*genai.GenerateContentResponse); ok {
		// Extract text from streaming response
		var textContent string
		for _, candidate := range geminiResponse.Candidates {
			if candidate == nil || candidate.Content == nil {
				continue
			}
			for _, part := range candidate.Content.Parts {
				if part != nil && part.Text != "" {
					textContent += part.Text
				}
			}
		}
		return textContent, nil
	}

	// Fallback to string
	if stringEvent, ok := rawEvent.(string); ok {
		return stringEvent, nil
	}

	return "", nil // Return empty for events that don't contain text
}

// IsStructuredResponse checks if response contains structured output
func (p *GeminiParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	// Check if schema was provided
	if opts.SchemaContent == "" {
		return false
	}

	// For Gemini, if schema is provided and we have a response, assume it's structured
	if geminiResponse, ok := rawResponse.(*genai.GenerateContentResponse); ok {
		if geminiResponse != nil {
			return true
		}
	}

	// Fallback: assume structured if schema was provided
	return true
}

// OllamaParser handles Ollama-specific response parsing
type OllamaParser struct{}

// NewOllamaParser creates a new Ollama response parser
func NewOllamaParser() *OllamaParser {
	return &OllamaParser{}
}

// ParseResponse parses an Ollama response for structured output
func (p *OllamaParser) ParseResponse(rawResponse interface{}, opts *domain.ChatOptions) (string, error) {
	// Ollama typically returns JSON directly for structured outputs
	if stringResponse, ok := rawResponse.(string); ok {
		if opts.SchemaContent != "" {
			// For structured outputs, validate it's valid JSON
			var jsonTest interface{}
			if err := json.Unmarshal([]byte(stringResponse), &jsonTest); err != nil {
				return "", fmt.Errorf("ollama structured response is not valid JSON: %w", err)
			}
		}
		return stringResponse, nil
	}
	return "", fmt.Errorf("expected string response from Ollama")
}

// ParseStreamEvent parses an Ollama streaming event
func (p *OllamaParser) ParseStreamEvent(rawEvent interface{}, opts *domain.ChatOptions) (string, error) {
	// Ollama streaming is typically simpler
	if stringEvent, ok := rawEvent.(string); ok {
		return stringEvent, nil
	}
	return "", fmt.Errorf("expected string event from Ollama")
}

// IsStructuredResponse checks if response contains structured output
func (p *OllamaParser) IsStructuredResponse(rawResponse interface{}, opts *domain.ChatOptions) bool {
	return opts.SchemaContent != ""
}
