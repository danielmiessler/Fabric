package i18n

import (
	"os"
	"strings"

	"golang.org/x/text/language"
)

// detectSystemLocale checks LC_ALL, LC_MESSAGES, and LANG in POSIX priority order.
// It returns the first valid locale, or "" when none is valid.
func detectSystemLocale() string {
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

// normalizeLocale converts a POSIX locale string to a BCP 47 language-REGION tag.
// It drops the encoding and modifier suffixes and returns "" for "C" and "POSIX".
// Examples:
//   - "en_US.UTF-8" -> "en-US"
//   - "fr_FR@euro" -> "fr-FR"
//   - "zh_CN.GB2312" -> "zh-CN"
func normalizeLocale(locale string) string {
	if locale == "C" || locale == "POSIX" || locale == "" {
		return ""
	}

	locale = strings.Split(locale, ".")[0]
	locale = strings.Split(locale, "@")[0]

	locale = strings.ReplaceAll(locale, "_", "-")

	parts := strings.Split(locale, "-")
	if len(parts) >= 2 {
		parts[0] = strings.ToLower(parts[0])
		parts[1] = strings.ToUpper(parts[1])
		locale = strings.Join(parts[:2], "-")
	} else if len(parts) == 1 {
		locale = strings.ToLower(parts[0])
	}

	return locale
}

// isValidLocale reports whether language.Parse accepts locale.
func isValidLocale(locale string) bool {
	if locale == "" {
		return false
	}

	_, err := language.Parse(locale)
	return err == nil
}

// getPreferredLocale returns explicitLang when it is not empty.
// Otherwise it returns the result of detectSystemLocale.
func getPreferredLocale(explicitLang string) string {
	if explicitLang != "" {
		return explicitLang
	}

	return detectSystemLocale()
}
