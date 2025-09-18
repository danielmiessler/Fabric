package i18n

import (
	"testing"

	gi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestTranslation(t *testing.T) {
	testCases := []struct {
		lang     string
		expected string
	}{
		{"es", "usa la entrada original, porque no se puede aplicar la legibilidad de html"},
		{"en", "use original input, because can't apply html readability"},
		{"zh", "使用原始输入，因为无法应用 HTML 可读性处理"},
		{"de", "verwende ursprüngliche Eingabe, da HTML-Lesbarkeit nicht angewendet werden kann"},
		{"ja", "HTML可読性を適用できないため、元の入力を使用します"},
		{"fr", "utilise l'entrée originale, car la lisibilité HTML ne peut pas être appliquée"},
		{"pt", "usa a entrada original, porque não é possível aplicar a legibilidade HTML"},
		{"fa", "از ورودی اصلی استفاده کن، چون نمی‌توان خوانایی HTML را اعمال کرد"},
		{"it", "usa l'input originale, perché non è possibile applicare la leggibilità HTML"},
	}

	for _, tc := range testCases {
		t.Run(tc.lang, func(t *testing.T) {
			loc, err := Init(tc.lang)
			if err != nil {
				t.Fatalf("init failed for %s: %v", tc.lang, err)
			}
			msg, err := loc.Localize(&gi18n.LocalizeConfig{MessageID: "html_readability_error"})
			if err != nil {
				t.Fatalf("localize failed for %s: %v", tc.lang, err)
			}
			if msg != tc.expected {
				t.Fatalf("unexpected translation for %s: got %s, expected %s", tc.lang, msg, tc.expected)
			}
		})
	}
}
