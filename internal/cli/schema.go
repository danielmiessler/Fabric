package cli

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// extractJSONFromCodeBlock attempts to extract JSON content from markdown code blocks.
// It handles both ```json and ``` wrapped content, as well as raw JSON.
func extractJSONFromCodeBlock(output string) string {
	// Try to match JSON code blocks with optional language specifier
	// Matches: ```json{content}``` or ```{content}```
	codeBlockPattern := regexp.MustCompile("(?s)```(?:json)?\\s*\\n?(.+?)\\n?```")
	matches := codeBlockPattern.FindStringSubmatch(output)

	if len(matches) > 1 {
		// Return the extracted JSON content
		return strings.TrimSpace(matches[1])
	}

	// If no code block found, return the original output
	// (it might already be valid JSON)
	return output
}

// validateOutputWithSchema validates the AI's JSON output against a given JSON schema.
// It takes the AI's output and the schema content as strings.
// It returns an error if the output does not conform to the schema.
func validateOutputWithSchema(output, schemaContent string) error {
	if schemaContent == "" {
		return nil // No schema to validate against
	}

	// Extract JSON from potential code blocks
	jsonOutput := extractJSONFromCodeBlock(output)

	schemaLoader := gojsonschema.NewStringLoader(schemaContent)
	documentLoader := gojsonschema.NewStringLoader(jsonOutput)

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
