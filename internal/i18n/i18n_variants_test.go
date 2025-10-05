package i18n

import (
	"testing"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestNormalizeToBCP47(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Basic cases
		{"pt", "pt"},
		{"pt-BR", "pt-BR"},
		{"pt-PT", "pt-PT"},

		// Underscore normalization
		{"pt_BR", "pt-BR"},
		{"pt_PT", "pt-PT"},
		{"en_US", "en-US"},

		// Mixed case normalization
		{"pt-br", "pt-BR"},
		{"PT-BR", "pt-BR"},
		{"Pt-Br", "pt-BR"},
		{"pT-bR", "pt-BR"},

		// Language only cases
		{"EN", "en"},
		{"Pt", "pt"},
		{"ZH", "zh"},

		// Empty string
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeToBCP47(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeToBCP47(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetLocaleCandidates(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		// Portuguese variants
		{"pt-PT", []string{"pt-PT", "pt", "pt-BR"}}, // pt-BR is default for pt
		{"pt-BR", []string{"pt-BR", "pt"}},          // pt-BR doesn't need default since it IS the default
		{"pt", []string{"pt", "pt-BR"}},             // pt defaults to pt-BR

		// Other languages without default variants
		{"en-US", []string{"en-US", "en"}},
		{"en", []string{"en"}},
		{"fr-FR", []string{"fr-FR", "fr"}},
		{"zh-CN", []string{"zh-CN", "zh"}},

		// Empty
		{"", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := getLocaleCandidates(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("getLocaleCandidates(%q) returned %d candidates; want %d",
					tt.input, len(result), len(tt.expected))
				t.Errorf("  got: %v", result)
				t.Errorf("  want: %v", tt.expected)
				return
			}
			for i, candidate := range result {
				if candidate != tt.expected[i] {
					t.Errorf("getLocaleCandidates(%q)[%d] = %q; want %q",
						tt.input, i, candidate, tt.expected[i])
				}
			}
		})
	}
}

func TestPortugueseVariantLoading(t *testing.T) {
	// Test that both Portuguese variants can be loaded
	testCases := []struct {
		locale string
		desc   string
	}{
		{"pt", "Portuguese (defaults to Brazilian)"},
		{"pt-BR", "Brazilian Portuguese"},
		{"pt-PT", "European Portuguese"},
		{"pt_BR", "Brazilian Portuguese with underscore"},
		{"pt_PT", "European Portuguese with underscore"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			localizer, err := Init(tc.locale)
			if err != nil {
				t.Errorf("Init(%q) failed: %v", tc.locale, err)
				return
			}
			if localizer == nil {
				t.Errorf("Init(%q) returned nil localizer", tc.locale)
			}

			// Try to get a message to verify it loaded correctly
			msg := localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "help_message"})
			if msg == "" {
				t.Errorf("Failed to localize message for locale %q", tc.locale)
			}
		})
	}
}

func TestPortugueseVariantDistinction(t *testing.T) {
	// Test that pt-BR and pt-PT return different translations
	localizerBR, err := Init("pt-BR")
	if err != nil {
		t.Fatalf("Failed to init pt-BR: %v", err)
	}

	localizerPT, err := Init("pt-PT")
	if err != nil {
		t.Fatalf("Failed to init pt-PT: %v", err)
	}

	// Check a key that should differ between variants
	// "output_to_file" should be "Exportar para arquivo" in pt-BR and "Saída para ficheiro" in pt-PT
	msgBR := localizerBR.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})
	msgPT := localizerPT.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})

	if msgBR == msgPT {
		t.Errorf("pt-BR and pt-PT returned the same translation for 'output_to_file': %q", msgBR)
	}

	// Verify specific expected values
	if msgBR != "Exportar para arquivo" {
		t.Errorf("pt-BR 'output_to_file' = %q; want 'Exportar para arquivo'", msgBR)
	}
	if msgPT != "Saída para ficheiro" {
		t.Errorf("pt-PT 'output_to_file' = %q; want 'Saída para ficheiro'", msgPT)
	}
}

func TestBackwardCompatibility(t *testing.T) {
	// Test that requesting "pt" still works and defaults to pt-BR
	localizerPT, err := Init("pt")
	if err != nil {
		t.Fatalf("Failed to init 'pt': %v", err)
	}

	localizerBR, err := Init("pt-BR")
	if err != nil {
		t.Fatalf("Failed to init 'pt-BR': %v", err)
	}

	// Both should return the same Brazilian Portuguese translation
	msgPT := localizerPT.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})
	msgBR := localizerBR.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})

	if msgPT != msgBR {
		t.Errorf("'pt' and 'pt-BR' returned different translations: %q vs %q", msgPT, msgBR)
	}

	if msgPT != "Exportar para arquivo" {
		t.Errorf("'pt' did not default to Brazilian Portuguese. Got %q, want 'Exportar para arquivo'", msgPT)
	}
}
