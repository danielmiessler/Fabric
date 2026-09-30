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

//go:embed locales/*.json
var localeFS embed.FS

var (
	translator *i18n.Localizer
	initOnce   sync.Once
)

// defaultLanguageVariants maps a base language to its fallback regional variant.
// getLocaleCandidates tries it after the requested locale and the base language.
var defaultLanguageVariants = map[string]string{
	"pt": "pt-BR", // "pt" meant Brazilian Portuguese before pt-PT.json existed
}

// Init initializes the i18n bundle and localizer. It loads the specified locale
// and falls back to English if loading fails.
// Translation files are searched in the user config directory and downloaded
// from GitHub if missing.
//
// If locale is empty, it will attempt to detect the system locale from
// environment variables (LC_ALL, LC_MESSAGES, LANG) following POSIX standards.
func Init(locale string) (*i18n.Localizer, error) {
	locale = getPreferredLocale(locale)
	locale = normalizeToBCP47(locale)
	if locale == "" {
		locale = "en"
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	locales := getLocaleCandidates(locale)

	embedded := false
	for _, candidate := range locales {
		if data, err := localeFS.ReadFile("locales/" + candidate + ".json"); err == nil {
			_, _ = bundle.ParseMessageFileBytes(data, candidate+".json")
			embedded = true
			locale = candidate // the disk path and the localizer use the loaded candidate
			break
		}
	}

	if !embedded {
		if data, err := localeFS.ReadFile("locales/en.json"); err == nil {
			_, _ = bundle.ParseMessageFileBytes(data, "en.json")
		}
	}

	path := filepath.Join(userLocaleDir(), locale+".json")
	if _, err := os.Stat(path); os.IsNotExist(err) && !embedded {
		if err := downloadLocale(path, locale); err != nil {
			// a failed download leaves the English fallback in place
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
			Init("")
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

// getErrorMessage reads messageID from the embedded file for the system locale,
// then from en.json, then returns fallback. Init uses it before translator exists.
func getErrorMessage(messageID, fallback string) string {
	systemLocale := getPreferredLocale("")
	if systemLocale == "" {
		systemLocale = "en"
	}

	if msg := tryGetMessage(systemLocale, messageID); msg != "" {
		return msg
	}

	if systemLocale != "en" {
		if msg := tryGetMessage("en", messageID); msg != "" {
			return msg
		}
	}

	return fallback
}

// tryGetMessage returns messageID from the embedded file for locale, or "" when absent.
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

// normalizeToBCP47 replaces underscores with hyphens, lowercases the language,
// uppercases the region, and drops any subtag after the region.
func normalizeToBCP47(locale string) string {
	if locale == "" {
		return ""
	}

	locale = strings.ReplaceAll(locale, "_", "-")

	parts := strings.Split(locale, "-")
	if len(parts) == 1 {
		return strings.ToLower(parts[0])
	} else if len(parts) >= 2 {
		parts[0] = strings.ToLower(parts[0])
		parts[1] = strings.ToUpper(parts[1])
		return strings.Join(parts[:2], "-")
	}

	return locale
}

// getLocaleCandidates returns the requested locale, then its base language, then
// the default variant for that language, without duplicates.
// For "pt-PT" it returns ["pt-PT", "pt", "pt-BR"].
func getLocaleCandidates(locale string) []string {
	candidates := []string{}

	if locale == "" {
		return candidates
	}

	candidates = append(candidates, locale)

	if baseLang, _, found := strings.Cut(locale, "-"); found {
		candidates = append(candidates, baseLang)

		if defaultVariant, exists := defaultLanguageVariants[baseLang]; exists {
			if defaultVariant != locale {
				candidates = append(candidates, defaultVariant)
			}
		}
	} else {
		if defaultVariant, exists := defaultLanguageVariants[locale]; exists {
			candidates = append(candidates, defaultVariant)
		}
	}

	return candidates
}
