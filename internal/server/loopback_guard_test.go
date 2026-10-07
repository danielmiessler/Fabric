package restapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

func guardEngine(corsOrigins []string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(LoopbackSecurityMiddleware(corsOrigins))
	r.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.POST("/patterns/:name", func(c *gin.Context) { c.Status(http.StatusOK) })
	return r
}

func TestLoopbackGuard_HostHeader(t *testing.T) {
	r := guardEngine(nil)

	cases := []struct {
		name string
		host string
		want int
	}{
		{"loopback name", "localhost:8080", http.StatusOK},
		{"loopback ipv4", "127.0.0.1:8080", http.StatusOK},
		{"loopback ipv6", "[::1]:8080", http.StatusOK},
		{"other host name", "other.example", http.StatusForbidden},
		{"public ip host", "203.0.113.5:8080", http.StatusForbidden},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			req.Host = tc.host
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Errorf("Host %q: got %d, want %d", tc.host, w.Code, tc.want)
			}
		})
	}
}

func TestLoopbackGuard_Origin(t *testing.T) {
	r := guardEngine([]string{"http://app.localhost:3000"})

	cases := []struct {
		name   string
		origin string
		want   int
	}{
		{"no origin (curl/CLI)", "", http.StatusOK},
		{"same-origin", "http://localhost:8080", http.StatusOK},
		{"web UI dev server with text/plain", "http://localhost:5173", http.StatusForbidden},
		{"loopback origin on a different port", "http://127.0.0.1:9999", http.StatusForbidden},
		{"same host on a different port", "http://localhost:3000", http.StatusForbidden},
		{"allowlisted origin", "http://app.localhost:3000", http.StatusOK},
		{"other site origin", "http://other.example", http.StatusForbidden},
		{"null origin", "null", http.StatusForbidden},
		{"origin with no scheme", "foo", http.StatusForbidden},
		{"origin with no host", "http://", http.StatusForbidden},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// A browser sends a text/plain POST to a different site with
			// no CORS preflight. Thus the Origin check must stop it.
			req := httptest.NewRequest(http.MethodPost, "/patterns/name", nil)
			req.Host = "localhost:8080"
			req.Header.Set("Content-Type", "text/plain")
			if tc.origin != "" {
				req.Header.Set("Origin", tc.origin)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Errorf("Origin %q: got %d, want %d", tc.origin, w.Code, tc.want)
			}
		})
	}
}

// The web UI sends JSON. The guard accepts the web UI origin for a JSON
// request, because a browser must do a preflight to send it cross-origin.
func TestLoopbackGuard_WebUIJSON(t *testing.T) {
	r := guardEngine(nil)
	req := httptest.NewRequest(http.MethodPost, "/patterns/name", nil)
	req.Host = "localhost:8080"
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Origin", "http://127.0.0.1:5173")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d, want %d", w.Code, http.StatusOK)
	}
}

// The guard lets a preflight and the Swagger documents through from any
// Host. It accepts the Host of a CORS origin, but only for that host.
func TestLoopbackGuard_Bypass(t *testing.T) {
	r := guardEngine([]string{"http://app.example:3000"})
	r.OPTIONS("/patterns/:name", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	r.GET("/swagger/*any", func(c *gin.Context) { c.Status(http.StatusOK) })

	cases := []struct {
		name, method, path, host string
		want                     int
	}{
		{"preflight from other host", http.MethodOptions, "/patterns/name", "evil.example", http.StatusNoContent},
		{"swagger from other host", http.MethodGet, "/swagger/index.html", "evil.example", http.StatusOK},
		{"swagger prefix only", http.MethodGet, "/swaggerx", "evil.example", http.StatusForbidden},
		{"CORS origin host", http.MethodGet, "/ping", "app.example:8080", http.StatusOK},
		{"other host", http.MethodGet, "/ping", "evil.example", http.StatusForbidden},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			req.Host = tc.host
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Errorf("%s %s Host %q: got %d, want %d", tc.method, tc.path, tc.host, w.Code, tc.want)
			}
		})
	}
}

// requireJSON accepts application/json with parameters and in any case. It
// rejects each other Content-Type and a request with no Content-Type.
func TestRequireJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/chat", requireJSON, func(c *gin.Context) { c.Status(http.StatusOK) })

	for contentType, want := range map[string]int{
		"application/json":                  http.StatusOK,
		"application/json; charset=utf-8":   http.StatusOK,
		"Application/JSON":                  http.StatusOK,
		"text/plain":                        http.StatusUnsupportedMediaType,
		"application/x-www-form-urlencoded": http.StatusUnsupportedMediaType,
		"":                                  http.StatusUnsupportedMediaType,
	} {
		req := httptest.NewRequest(http.MethodPost, "/chat", strings.NewReader("{}"))
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != want {
			t.Errorf("Content-Type %q: got %d, want %d", contentType, w.Code, want)
		}
	}
}

// Each route that reads a JSON body must use requireJSON. A save route
// takes a raw body with any Content-Type.
func TestJSONRoutesRequireJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}
	r := newOllamaEngine(registry, ":0", "test-version", "", nil)
	NewYouTubeHandler(r, registry)

	post := func(path string) int {
		req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{}"))
		req.Host = "localhost"
		req.Header.Set("Content-Type", "text/plain")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w.Code
	}
	for _, path := range []string{"/chat", "/patterns/name/apply", "/config/update", "/youtube/transcript", "/api/chat"} {
		if code := post(path); code != http.StatusUnsupportedMediaType {
			t.Errorf("POST %s: got %d, want 415", path, code)
		}
	}
	if code := post("/patterns/name"); code != http.StatusOK {
		t.Errorf("POST /patterns/name: got %d, want 200", code)
	}
}
