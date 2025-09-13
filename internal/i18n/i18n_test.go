package i18n

import (
    "testing"

    gi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestTranslation(t *testing.T) {
    cases := []struct {
        lang, expected string
    }{
        {"es", "usa la entrada original, porque no se puede aplicar la legibilidad de html"},
        {"de", "Originaleingabe wird beibehalten, da HTML nicht leserlich aufbereitet werden konnte"},
    }

    for _, tc := range cases {
        tc := tc
        t.Run(tc.lang, func(t *testing.T) {
            loc, err := Init(tc.lang)
            if err != nil {
                t.Fatalf("init failed: %v", err)
            }
            msg, err := loc.Localize(&gi18n.LocalizeConfig{MessageID: "html_readability_error"})
            if err != nil {
                t.Fatalf("localize failed: %v", err)
            }
            if msg != tc.expected {
                t.Fatalf("unexpected translation (%s): %q", tc.lang, msg)
            }
        })
    }
}

