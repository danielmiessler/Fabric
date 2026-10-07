package restapi

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// TestServeEngineMiddleware checks the middleware chain that Serve installs.
// If a r.Use line goes out of newSecuredEngine, a case here fails.
func TestServeEngineMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("OPENAI_API_KEY", "")
	registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}
	withoutKey := newServeEngine(registry, "", nil)
	withKey := newServeEngine(registry, "secret", nil)

	big := bytes.Repeat([]byte("a"), MaxRequestBodyBytes+1)
	cases := []struct {
		name   string
		r      *gin.Engine
		host   string
		origin string
		body   io.Reader
		want   int
	}{
		{"body above the limit", withoutKey, "localhost", "", bytes.NewReader(big), http.StatusRequestEntityTooLarge},
		{"other host with no key", withoutKey, "evil.example", "", strings.NewReader("{}"), http.StatusForbidden},
		{"other site origin", withoutKey, "localhost", "http://evil.example", strings.NewReader("{}"), http.StatusForbidden},
		{"key set but not sent", withKey, "localhost", "", strings.NewReader("{}"), http.StatusUnauthorized},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/chat", tc.body)
			req.Host = tc.host
			req.Header.Set("Content-Type", "application/json")
			if tc.origin != "" {
				req.Header.Set("Origin", tc.origin)
			}
			w := httptest.NewRecorder()
			tc.r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Errorf("got %d, want %d", w.Code, tc.want)
			}
		})
	}
}

// TestServeEngineSwaggerExemption checks that a server with an API key
// serves /swagger/ with no key, and that the exemption covers only that
// prefix.
func TestServeEngineSwaggerExemption(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := newServeEngine(&core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}, "secret", nil)

	for path, want := range map[string]int{
		"/swagger/index.html": http.StatusOK,
		"/swaggerx":           http.StatusUnauthorized,
		"/patterns/names":     http.StatusUnauthorized,
	} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Host = "localhost"
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != want {
			t.Errorf("GET %s: got %d, want %d", path, w.Code, want)
		}
	}
}

// TestServeEngineWithCORSOrigins checks the CORS middleware and the loopback
// guard together, on a server with no API key.
func TestServeEngineWithCORSOrigins(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("OPENAI_API_KEY", "")
	r := newServeEngine(&core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}, "", []string{"http://localhost:5173"})

	cases := []struct {
		name, method, host, origin string
		wantCode                   int
		wantAllowOrigin            string
	}{
		{"listed origin", http.MethodPost, "localhost:8080", "http://localhost:5173", http.StatusUnsupportedMediaType, "http://localhost:5173"},
		{"other site origin", http.MethodPost, "localhost:8080", "http://evil.example", http.StatusForbidden, ""},
		{"other host", http.MethodPost, "evil.example", "http://localhost:5173", http.StatusForbidden, "http://localhost:5173"},
		{"preflight from listed origin", http.MethodOptions, "localhost:8080", "http://localhost:5173", http.StatusNoContent, "http://localhost:5173"},
		{"preflight from other site", http.MethodOptions, "evil.example", "http://evil.example", http.StatusNoContent, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// text/plain gets past the Origin check only to requireJSON, which
			// answers 415. Thus 415 shows that both middleware accepted it.
			req := httptest.NewRequest(tc.method, "/chat", strings.NewReader("{}"))
			req.Host = tc.host
			req.Header.Set("Content-Type", "text/plain")
			req.Header.Set("Origin", tc.origin)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Errorf("status = %d, want %d", w.Code, tc.wantCode)
			}
			if got := w.Header().Get("Access-Control-Allow-Origin"); got != tc.wantAllowOrigin {
				t.Errorf("Allow-Origin = %q, want %q", got, tc.wantAllowOrigin)
			}
		})
	}
}

// TestEnginesRefuseSystemPlugins checks that a pattern that a client saves
// cannot read the server environment through the REST or the Ollama server.
func TestEnginesRefuseSystemPlugins(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("OPENAI_API_KEY", "sk-test-secret")
	engines := map[string]func(*core.PluginRegistry) *gin.Engine{
		"serve": func(reg *core.PluginRegistry) *gin.Engine { return newServeEngine(reg, "", nil) },
		"ollama": func(reg *core.PluginRegistry) *gin.Engine {
			return newOllamaEngine(reg, "localhost:11434", "test", "", nil)
		},
	}
	for engineName, newEngine := range engines {
		t.Run(engineName, func(t *testing.T) {
			r := newEngine(&core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())})

			send := func(method, path, contentType, body string) *httptest.ResponseRecorder {
				req := httptest.NewRequest(method, path, strings.NewReader(body))
				req.Host = "localhost:8080"
				req.Header.Set("Content-Type", contentType)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				return w
			}

			for name, body := range map[string]string{
				"leak":  "K={{plugin:sys:env:OPENAI_API_KEY}}",
				"upper": "{{plugin:text:upper:ok}}",
			} {
				if w := send(http.MethodPost, "/patterns/"+name, "text/plain", body); w.Code != http.StatusOK {
					t.Fatalf("save %s: status %d", name, w.Code)
				}
			}

			w := send(http.MethodPost, "/patterns/leak/apply", "application/json", "{}")
			if w.Code == http.StatusOK || strings.Contains(w.Body.String(), "sk-test-secret") {
				t.Errorf("apply leak: status %d, body %q; want an error and no key", w.Code, w.Body)
			}
			w = send(http.MethodPost, "/patterns/upper/apply", "application/json", "{}")
			if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "OK") {
				t.Errorf("apply upper: status %d, body %q; want 200 with OK", w.Code, w.Body)
			}
		})
	}
}

// TestSecuredEngineIgnoresForwardedFor checks that a client cannot set the
// IP address in the logs with the X-Forwarded-For header.
func TestSecuredEngineIgnoresForwardedFor(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := newSecuredEngine(&core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}, "secret", nil, "")
	r.GET("/ip", func(c *gin.Context) { c.String(http.StatusOK, c.ClientIP()) })

	req := httptest.NewRequest(http.MethodGet, "/ip", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	req.Header.Set("X-Forwarded-For", "1.2.3.4")
	req.Header.Set(APIKeyHeader, "secret")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if got := w.Body.String(); got != "127.0.0.1" {
		t.Errorf("ClientIP() = %q, want %q", got, "127.0.0.1")
	}
}
