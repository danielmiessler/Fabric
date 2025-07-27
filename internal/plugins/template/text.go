// Package template provides text transformation operations for the template system.
package template

import (
	"fmt"
	"strings"
	"unicode"
)

// TextPlugin provides string manipulation operations
type TextPlugin struct{}

// toTitle capitalizes a letter if it follows a non-letter, unless next char is space
func toTitle(s string) string {
	// First lowercase everything
	lower := strings.ToLower(s)
	runes := []rune(lower)

	for i := 0; i < len(runes); i++ {
		// Capitalize if previous char is non-letter AND
		// (we're at the end OR next char is not space)
		if i == 0 || !unicode.IsLetter(runes[i-1]) {
			if i == len(runes)-1 || !unicode.IsSpace(runes[i+1]) {
				runes[i] = unicode.ToUpper(runes[i])
			}
		}
	}

	return string(runes)
}

// Apply executes the requested text operation on the provided value
func (p *TextPlugin) Apply(operation string, value string) (string, error) {
	debugf("TextPlugin: operation=%s value=%q", operation, value)

	if value == "" {
		return "", fmt.Errorf("text: empty input for operation %q", operation)
	}

	switch operation {
	case "upper":
		result := strings.ToUpper(value)
		debugf("TextPlugin: upper result=%q", result)
		return result, nil

	case "lower":
		result := strings.ToLower(value)
		debugf("TextPlugin: lower result=%q", result)
		return result, nil

	case "title":
		result := toTitle(value)
		debugf("TextPlugin: title result=%q", result)
		return result, nil

	case "trim":
		result := strings.TrimSpace(value)
		debugf("TextPlugin: trim result=%q", result)
		return result, nil

	default:
		return "", fmt.Errorf("text: unknown text operation %q (supported: upper, lower, title, trim)", operation)
	}
}
