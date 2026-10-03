package i18n

import (
	"os"
	"testing"
)

func TestDetectSystemLocale(t *testing.T) {
	originalLC_ALL := os.Getenv("LC_ALL")
	originalLC_MESSAGES := os.Getenv("LC_MESSAGES")
	originalLANG := os.Getenv("LANG")

	defer func() {
		os.Setenv("LC_ALL", originalLC_ALL)
		os.Setenv("LC_MESSAGES", originalLC_MESSAGES)
		os.Setenv("LANG", originalLANG)
	}()

	tests := []struct {
		name        string
		LC_ALL      string
		LC_MESSAGES string
		LANG        string
		expected    string
		description string
	}{
		{
			name:        "LC_ALL takes highest priority",
			LC_ALL:      "fr_FR.UTF-8",
			LC_MESSAGES: "de_DE.UTF-8",
			LANG:        "es_ES.UTF-8",
			expected:    "fr-FR",
			description: "LC_ALL should override all other variables",
		},
		{
			name:        "LC_MESSAGES used when LC_ALL empty",
			LC_ALL:      "",
			LC_MESSAGES: "ja_JP.UTF-8",
			LANG:        "ko_KR.UTF-8",
			expected:    "ja-JP",
			description: "LC_MESSAGES should be used when LC_ALL is not set",
		},
		{
			name:        "LANG used when LC_ALL and LC_MESSAGES empty",
			LC_ALL:      "",
			LC_MESSAGES: "",
			LANG:        "zh_CN.GB2312",
			expected:    "zh-CN",
			description: "LANG should be fallback when others are not set",
		},
		{
			name:        "Empty when no valid locale set",
			LC_ALL:      "",
			LC_MESSAGES: "",
			LANG:        "",
			expected:    "",
			description: "Should return empty when no environment variables set",
		},
		{
			name:        "Handle C locale",
			LC_ALL:      "C",
			LC_MESSAGES: "",
			LANG:        "",
			expected:    "",
			description: "C locale should be treated as invalid (fallback to default)",
		},
		{
			name:        "Handle POSIX locale",
			LC_ALL:      "",
			LC_MESSAGES: "POSIX",
			LANG:        "",
			expected:    "",
			description: "POSIX locale should be treated as invalid (fallback to default)",
		},
		{
			name:        "Handle locale with modifiers",
			LC_ALL:      "",
			LC_MESSAGES: "",
			LANG:        "de_DE.UTF-8@euro",
			expected:    "de-DE",
			description: "Should strip encoding and modifiers",
		},
		{
			name:        "Skip invalid locale and use next priority",
			LC_ALL:      "invalid_locale",
			LC_MESSAGES: "fr_CA.UTF-8",
			LANG:        "en_US.UTF-8",
			expected:    "fr-CA",
			description: "Should skip invalid high-priority locale and use next valid one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("LC_ALL", tt.LC_ALL)
			os.Setenv("LC_MESSAGES", tt.LC_MESSAGES)
			os.Setenv("LANG", tt.LANG)

			result := detectSystemLocale()
			if result != tt.expected {
				t.Errorf("%s: expected %q, got %q", tt.description, tt.expected, result)
			}
		})
	}
}

func TestNormalizeLocale(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// POSIX locale strings
		{"en_US.UTF-8", "en-US"},
		{"fr_FR.ISO8859-1", "fr-FR"},
		{"de_DE@euro", "de-DE"},
		{"zh_CN.GB2312", "zh-CN"},
		{"ja_JP.eucJP@traditional", "ja-JP"},

		// Already normalized
		{"en-US", "en-US"},
		{"fr-CA", "fr-CA"},

		// Language only
		{"en", "en"},
		{"fr", "fr"},
		{"zh", "zh"},

		// C and POSIX mean no locale
		{"C", ""},
		{"POSIX", ""},
		{"", ""},

		// Modifiers with values and scripts
		{"pt_BR.UTF-8@currency=BRL", "pt-BR"},
		{"sr_RS.UTF-8@latin", "sr-RS"},
		{"uz_UZ.UTF-8@cyrillic", "uz-UZ"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeLocale(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeLocale(%q): expected %q, got %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestIsValidLocale(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"en", true},
		{"en-US", true},
		{"fr-FR", true},
		{"zh-CN", true},
		{"ja-JP", true},
		{"pt-BR", true},
		{"es-MX", true},

		{"", false},
		{"invalid", false},
		{"123", false},

		// language.Parse also accepts "en-ZZ" (unknown region) and "en_US" (underscore),
		// so those are not invalid cases.
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isValidLocale(tt.input)
			if result != tt.expected {
				t.Errorf("isValidLocale(%q): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}

func TestGetPreferredLocale(t *testing.T) {
	originalLC_ALL := os.Getenv("LC_ALL")
	originalLC_MESSAGES := os.Getenv("LC_MESSAGES")
	originalLANG := os.Getenv("LANG")

	defer func() {
		os.Setenv("LC_ALL", originalLC_ALL)
		os.Setenv("LC_MESSAGES", originalLC_MESSAGES)
		os.Setenv("LANG", originalLANG)
	}()

	tests := []struct {
		name         string
		explicitLang string
		LC_ALL       string
		LC_MESSAGES  string
		LANG         string
		expected     string
		description  string
	}{
		{
			name:         "Explicit language takes precedence",
			explicitLang: "es-ES",
			LC_ALL:       "fr_FR.UTF-8",
			LC_MESSAGES:  "de_DE.UTF-8",
			LANG:         "ja_JP.UTF-8",
			expected:     "es-ES",
			description:  "Explicit language should override environment variables",
		},
		{
			name:         "Use environment when no explicit language",
			explicitLang: "",
			LC_ALL:       "it_IT.UTF-8",
			LC_MESSAGES:  "ru_RU.UTF-8",
			LANG:         "pl_PL.UTF-8",
			expected:     "it-IT",
			description:  "Should detect from environment when no explicit language",
		},
		{
			name:         "Empty when no explicit and no environment",
			explicitLang: "",
			LC_ALL:       "",
			LC_MESSAGES:  "",
			LANG:         "",
			expected:     "",
			description:  "Should return empty when nothing is set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("LC_ALL", tt.LC_ALL)
			os.Setenv("LC_MESSAGES", tt.LC_MESSAGES)
			os.Setenv("LANG", tt.LANG)

			result := getPreferredLocale(tt.explicitLang)
			if result != tt.expected {
				t.Errorf("%s: expected %q, got %q", tt.description, tt.expected, result)
			}
		})
	}
}

func TestIntegrationWithInit(t *testing.T) {
	originalLC_ALL := os.Getenv("LC_ALL")
	originalLANG := os.Getenv("LANG")

	defer func() {
		os.Setenv("LC_ALL", originalLC_ALL)
		os.Setenv("LANG", originalLANG)
		translator = nil
	}()

	os.Setenv("LC_ALL", "es_ES.UTF-8")
	os.Setenv("LANG", "fr_FR.UTF-8")

	localizer, err := Init("")
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if localizer == nil {
		t.Error("Expected non-nil localizer")
	}

	// T must run Init itself when translator is nil.
	translator = nil
	os.Setenv("LC_ALL", "")
	os.Setenv("LANG", "es_ES.UTF-8")

	result := T("html_readability_error")
	if result == "" {
		t.Error("Expected non-empty translation result")
	}
}
