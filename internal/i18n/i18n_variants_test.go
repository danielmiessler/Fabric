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
		// Already normalized
		{"pt", "pt"},
		{"pt-BR", "pt-BR"},
		{"pt-PT", "pt-PT"},

		// Underscore to hyphen
		{"pt_BR", "pt-BR"},
		{"pt_PT", "pt-PT"},
		{"en_US", "en-US"},

		// Mixed case
		{"pt-br", "pt-BR"},
		{"PT-BR", "pt-BR"},
		{"Pt-Br", "pt-BR"},
		{"pT-bR", "pt-BR"},

		// Language only
		{"EN", "en"},
		{"Pt", "pt"},
		{"ZH", "zh"},

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
		{"pt-PT", []string{"pt-PT", "pt", "pt-BR"}}, // pt-BR is the default for pt
		{"pt-BR", []string{"pt-BR", "pt"}},          // the request is already the default
		{"pt", []string{"pt", "pt-BR"}},

		// Arabic variants
		{"ar-SA", []string{"ar-SA", "ar", "ar-BH"}}, // ar-BH is the default for ar
		{"ar-BH", []string{"ar-BH", "ar"}},
		{"ar", []string{"ar", "ar-BH"}},

		// Other languages without default variants
		{"en-US", []string{"en-US", "en"}},
		{"en", []string{"en"}},
		{"fr-FR", []string{"fr-FR", "fr"}},
		{"zh-CN", []string{"zh-CN", "zh"}},

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

			msg := localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "help_message"})
			if msg == "" {
				t.Errorf("Failed to localize message for locale %q", tc.locale)
			}
		})
	}
}

func TestPortugueseVariantDistinction(t *testing.T) {
	localizerBR, err := Init("pt-BR")
	if err != nil {
		t.Fatalf("Failed to init pt-BR: %v", err)
	}

	localizerPT, err := Init("pt-PT")
	if err != nil {
		t.Fatalf("Failed to init pt-PT: %v", err)
	}

	msgBR := localizerBR.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})
	msgPT := localizerPT.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})

	if msgBR == msgPT {
		t.Errorf("pt-BR and pt-PT returned the same translation for 'output_to_file': %q", msgBR)
	}

	if msgBR != "Exportar para arquivo" {
		t.Errorf("pt-BR 'output_to_file' = %q; want 'Exportar para arquivo'", msgBR)
	}
	if msgPT != "Saída para ficheiro" {
		t.Errorf("pt-PT 'output_to_file' = %q; want 'Saída para ficheiro'", msgPT)
	}
}

func TestBackwardCompatibility(t *testing.T) {
	localizerPT, err := Init("pt")
	if err != nil {
		t.Fatalf("Failed to init 'pt': %v", err)
	}

	localizerBR, err := Init("pt-BR")
	if err != nil {
		t.Fatalf("Failed to init 'pt-BR': %v", err)
	}

	msgPT := localizerPT.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})
	msgBR := localizerBR.MustLocalize(&goi18n.LocalizeConfig{MessageID: "output_to_file"})

	if msgPT != msgBR {
		t.Errorf("'pt' and 'pt-BR' returned different translations: %q vs %q", msgPT, msgBR)
	}

	if msgPT != "Exportar para arquivo" {
		t.Errorf("'pt' did not default to Brazilian Portuguese. Got %q, want 'Exportar para arquivo'", msgPT)
	}
}
