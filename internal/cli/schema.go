package cli

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// validateOutputWithSchema validates the AI's JSON output against a given JSON schema.
// It takes the AI's output and the schema content as strings.
// It returns an error if the output does not conform to the schema.
func validateOutputWithSchema(output, schemaContent string) error {
	if schemaContent == "" {
		return nil // No schema to validate against
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
