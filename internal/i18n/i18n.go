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
	if locale == "" {
		locale = "en"
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// load embedded translations for the requested locale if available
	embedded := false
	if data, err := localeFS.ReadFile("locales/" + locale + ".json"); err == nil {
		_, _ = bundle.ParseMessageFileBytes(data, locale+".json")
		embedded = true
	} else if strings.Contains(locale, "-") {
		// Try base language if regional variant not found (e.g., es-ES -> es)
		baseLang := strings.Split(locale, "-")[0]
		if data, err := localeFS.ReadFile("locales/" + baseLang + ".json"); err == nil {
			_, _ = bundle.ParseMessageFileBytes(data, baseLang+".json")
			embedded = true
		}
	}
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
