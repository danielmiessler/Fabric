package restapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// TestRecoveryDumpsRedactCredentials checks the request dump that gin writes
// when a handler panics. The debug-mode panic dump and the broken-pipe dump
// both print the request headers, and the key is in one of them.
func TestRecoveryDumpsRedactCredentials(t *testing.T) {
	const (
		apiKey = "synthetic-api-key-0123456789"
		cookie = "session=synthetic-cookie-value"
	)

	cases := []struct {
		name   string
		mode   string
		panics any
	}{
		{"debug-mode panic dump", gin.DebugMode, "synthetic panic"},
		{"broken-pipe dump", gin.ReleaseMode, http.ErrAbortHandler},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("OPENAI_API_KEY", "")
			oldWriter := gin.DefaultErrorWriter
			t.Cleanup(func() {
				gin.DefaultErrorWriter = oldWriter
				gin.SetMode(gin.TestMode)
			})
			var logged bytes.Buffer
			gin.DefaultErrorWriter = &logged

			registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}
			r := newSecuredEngine(registry, apiKey, nil, "no key")
			r.GET("/boom", func(c *gin.Context) { panic(tc.panics) })
			gin.SetMode(tc.mode) // the engine sets release mode; the panic dump needs debug

			req := httptest.NewRequest(http.MethodGet, "/boom", nil)
			req.Header.Set(APIKeyHeader, apiKey)
			req.Header.Set("Cookie", cookie)
			req.Header.Set("User-Agent", "synthetic-agent")
			r.ServeHTTP(httptest.NewRecorder(), req)

			out := logged.String()
			if !strings.Contains(out, "synthetic-agent") {
				t.Fatalf("the dump should still show the other headers, got:\n%s", out)
			}
			for _, secret := range []string{apiKey, "synthetic-cookie-value"} {
				if strings.Contains(out, secret) {
					t.Errorf("the recovery dump leaks %q:\n%s", secret, out)
				}
			}
		})
	}
}
