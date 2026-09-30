// Package template provides text transformation operations for the template system.
package template

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// TextPlugin provides string manipulation operations
type TextPlugin struct{}

// toTitle lowercases s, then uppercases each rune that starts the string or follows a non-letter.
// Such a rune stays lowercase when a space follows it.
func toTitle(s string) string {
	lower := strings.ToLower(s)
	runes := []rune(lower)

	for i := range runes {
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
		return "", fmt.Errorf(i18n.T("template_text_empty_input"), operation)
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
		return "", fmt.Errorf(i18n.T("template_text_unknown_operation"), operation)
	}
}
