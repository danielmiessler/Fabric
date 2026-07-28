package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// embedded default locales
//
//go:embed locales/*.json
var localeFS embed.FS

var (
	translator *i18n.Localizer
	initOnce   sync.Once
)

// defaultLanguageVariants maps language codes without regions to their default regional variants.
// This is used when a language without a base file is requested.
var defaultLanguageVariants = map[string]string{
	"pt": "pt-BR", // Portuguese defaults to Brazilian Portuguese for backward compatibility
	// Note: We currently have base files for these languages, but if we add regional variants
	// in the future, these defaults will be used:
	// "de": "de-DE", // German would default to Germany German
	// "en": "en-US", // English would default to US English
	// "es": "es-ES", // Spanish would default to Spain Spanish
	// "fa": "fa-IR", // Persian would default to Iran Persian
	// "fr": "fr-FR", // French would default to France French
	// "it": "it-IT", // Italian would default to Italy Italian
	// "ja": "ja-JP", // Japanese would default to Japan Japanese
	// "zh": "zh-CN", // Chinese would default to Simplified Chinese
}

// Init initializes the i18n bundle and localizer. It loads the specified locale
// and falls back to English if loading fails.
// Translation files are searched in the user config directory and downloaded
// from GitHub if missing.
//
// If locale is empty, it will attempt to detect the system locale from
// environment variables (LC_ALL, LC_MESSAGES, LANG) following POSIX standards.
func Init(locale string) (*i18n.Localizer, error) {
	// Use preferred locale detection if no explicit locale provided
	locale = getPreferredLocale(locale)
	// Normalize the locale to BCP 47 format (with hyphens)
	locale = normalizeToBCP47(locale)
	if locale == "" {
		locale = "en"
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Build a list of locale candidates to try
	locales := getLocaleCandidates(locale)

	// Try to load embedded translations for each candidate
	embedded := false
	for _, candidate := range locales {
		if data, err := localeFS.ReadFile("locales/" + candidate + ".json"); err == nil {
			_, _ = bundle.ParseMessageFileBytes(data, candidate+".json")
			embedded = true
			locale = candidate // Update locale to what was actually loaded
			break
		}
	}

	// Fall back to English if nothing was loaded
	if !embedded {
		if data, err := localeFS.ReadFile("locales/en.json"); err == nil {
			_, _ = bundle.ParseMessageFileBytes(data, "en.json")
		}
	}

	// load locale from disk or download when not embedded
	path := filepath.Join(userLocaleDir(), locale+".json")
	if _, err := os.Stat(path); os.IsNotExist(err) && !embedded {
		if err := downloadLocale(path, locale); err != nil {
			// if download fails, still continue with embedded translations
			fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf(getErrorMessage("i18n_download_failed", "Failed to download translation for language '%s': %v"), locale, err))
		}
	}
	if _, err := os.Stat(path); err == nil {
		if _, err := bundle.LoadMessageFile(path); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf(getErrorMessage("i18n_load_failed", "Failed to load translation file: %v"), err))
		}
	}

	translator = i18n.NewLocalizer(bundle, locale)
	return translator, nil
}

// T returns the localized string for the given message id.
// If the translator is not initialized, it will automatically initialize
// with system locale detection.
func T(messageID string) string {
	initOnce.Do(func() {
		if translator == nil {
			Init("") // Empty string triggers system locale detection
		}
	})
	return translator.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
}

func userLocaleDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	path := filepath.Join(dir, "fabric", "locales")
	os.MkdirAll(path, 0o755)
	return path
}

func downloadLocale(path, locale string) error {
	url := fmt.Sprintf("https://raw.githubusercontent.com/danielmiessler/Fabric/main/internal/i18n/locales/%s.json", locale)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// getErrorMessage tries to get a translated error message, falling back to system locale
// and then to the provided fallback message. This is used during initialization when
// the translator may not be fully ready.
func getErrorMessage(messageID, fallback string) string {
	// Try to get system locale for error messages
	systemLocale := getPreferredLocale("")
	if systemLocale == "" {
		systemLocale = "en"
	}

	// First try the system locale
	if msg := tryGetMessage(systemLocale, messageID); msg != "" {
		return msg
	}

	// Fall back to English
	if systemLocale != "en" {
		if msg := tryGetMessage("en", messageID); msg != "" {
			return msg
		}
	}

	// Final fallback to hardcoded message
	return fallback
}

// tryGetMessage attempts to get a message from embedded locale files
func tryGetMessage(locale, messageID string) string {
	if data, err := localeFS.ReadFile("locales/" + locale + ".json"); err == nil {
		var messages map[string]string
		if json.Unmarshal(data, &messages) == nil {
			if msg, exists := messages[messageID]; exists {
				return msg
			}
		}
	}
	return ""
}

// normalizeToBCP47 normalizes a locale string to BCP 47 format.
// Converts underscores to hyphens and ensures proper casing (language-REGION).
func normalizeToBCP47(locale string) string {
	if locale == "" {
		return ""
	}

	// Replace underscores with hyphens
	locale = strings.ReplaceAll(locale, "_", "-")

	// Split into parts
	parts := strings.Split(locale, "-")
	if len(parts) == 1 {
		// Language only, lowercase it
		return strings.ToLower(parts[0])
	} else if len(parts) >= 2 {
		// Language and region (and possibly more)
		// Lowercase language, uppercase region
		parts[0] = strings.ToLower(parts[0])
		parts[1] = strings.ToUpper(parts[1])
		return strings.Join(parts[:2], "-") // Return only language-REGION
	}

	return locale
}

// getLocaleCandidates returns a list of locale candidates to try, in order of preference.
// For example, for "pt-PT" it returns ["pt-PT", "pt", "pt-BR"] (where pt-BR is the default for pt).
func getLocaleCandidates(locale string) []string {
	candidates := []string{}

	if locale == "" {
		return candidates
	}

	// First candidate is always the requested locale
	candidates = append(candidates, locale)

	// If it's a regional variant, add the base language as a candidate
	if baseLang, _, found := strings.Cut(locale, "-"); found {
		candidates = append(candidates, baseLang)

		// Also check if the base language has a default variant
		if defaultVariant, exists := defaultLanguageVariants[baseLang]; exists {
			// Only add if it's different from what we already have
			if defaultVariant != locale {
				candidates = append(candidates, defaultVariant)
			}
		}
	} else {
		// If this is a base language without a region, check for default variant
		if defaultVariant, exists := defaultLanguageVariants[locale]; exists {
			candidates = append(candidates, defaultVariant)
		}
	}

	return candidates
}
