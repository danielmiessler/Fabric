package i18n

import (
	"os"
	"strings"

	"golang.org/x/text/language"
)

// detectSystemLocale detects the system locale using standard Unix environment variables.
// Follows the POSIX priority order for locale environment variables:
// 1. LC_ALL (highest priority - overrides all others)
// 2. LC_MESSAGES (for messages specifically)
// 3. LANG (general locale setting)
// 4. Returns empty string if none are set or valid
//
// This implementation follows POSIX standards and Unix best practices for locale detection.
func detectSystemLocale() string {
	// Check environment variables in priority order
	envVars := []string{"LC_ALL", "LC_MESSAGES", "LANG"}

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value != "" {
			locale := normalizeLocale(value)
			if locale != "" && isValidLocale(locale) {
				return locale
			}
		}
	}

	return ""
}

// normalizeLocale converts various locale formats to BCP 47 language tags.
// Examples:
//   - "en_US.UTF-8" -> "en-US"
//   - "fr_FR@euro" -> "fr-FR"
//   - "zh_CN.GB2312" -> "zh-CN"
//   - "C" or "POSIX" -> "" (invalid, falls back to default)
func normalizeLocale(locale string) string {
	// Handle special cases
	if locale == "C" || locale == "POSIX" || locale == "" {
		return ""
	}

	// Remove encoding and modifiers
	// Examples: en_US.UTF-8@euro -> en_US
	locale = strings.Split(locale, ".")[0] // Remove encoding (.UTF-8)
	locale = strings.Split(locale, "@")[0] // Remove modifiers (@euro)

	// Convert underscore to hyphen for BCP 47 compliance
	// en_US -> en-US
	locale = strings.ReplaceAll(locale, "_", "-")

	return locale
}

// isValidLocale checks if a locale string can be parsed as a valid language tag.
func isValidLocale(locale string) bool {
	if locale == "" {
		return false
	}

	// Use golang.org/x/text/language to validate
	_, err := language.Parse(locale)
	return err == nil
}

// getPreferredLocale returns the best locale to use based on user preferences.
// Priority order:
// 1. Explicit language flag (if provided)
// 2. System environment variables (LC_ALL, LC_MESSAGES, LANG)
// 3. Default fallback (empty string, which triggers "en" in Init)
func getPreferredLocale(explicitLang string) string {
	// If explicitly set via flag, use that
	if explicitLang != "" {
		return explicitLang
	}

	// Otherwise try to detect from system environment
	return detectSystemLocale()
}
