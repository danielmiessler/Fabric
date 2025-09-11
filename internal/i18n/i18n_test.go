package i18n

import (
	"testing"

	gi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestTranslation(t *testing.T) {
	loc, err := Init("es")
	if err != nil {
		t.Fatalf("init failed: %v", err)
	}
	msg, err := loc.Localize(&gi18n.LocalizeConfig{MessageID: "html_readability_error"})
	if err != nil {
		t.Fatalf("localize failed: %v", err)
	}
	expected := "usa la entrada original, porque no se puede aplicar la legibilidad de html"
	if msg != expected {
		t.Fatalf("unexpected translation: %s", msg)
	}
}
