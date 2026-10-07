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

// The engine answers 413 for a body above MaxRequestBodyBytes, before the
// handler runs. A normal request still gets to the handler.
func TestBodyLimitOnEngine(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("OPENAI_API_KEY", "")
	registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}
	r := newOllamaEngine(registry, ":0", "test-version", "", nil)

	post := func(path string, body io.Reader) int {
		req := httptest.NewRequest(http.MethodPost, path, body)
		if _, ok := body.(*chunked); ok {
			req.ContentLength = -1
		}
		req.Host = "localhost"
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w.Code
	}

	big := bytes.Repeat([]byte("a"), MaxRequestBodyBytes+1)
	for _, path := range []string{"/config/update", "/chat", "/api/chat"} {
		if code := post(path, bytes.NewReader(big)); code != http.StatusRequestEntityTooLarge {
			t.Errorf("POST %s with large body: got %d, want 413", path, code)
		}
		if code := post(path, &chunked{bytes.NewReader(big)}); code != http.StatusRequestEntityTooLarge {
			t.Errorf("POST %s with large chunked body: got %d, want 413", path, code)
		}
	}
	if code := post("/config/update", strings.NewReader(`{"openai_api_key":"sk-test"}`)); code != http.StatusOK {
		t.Errorf("POST /config/update with normal body: got %d, want 200", code)
	}
}

// A body with no Content-Length cannot be read past the limit.
func TestBodyLimitWithoutContentLength(t *testing.T) {
	gin.SetMode(gin.TestMode)
	const limit = 1 << 10
	r := gin.New()
	r.Use(MaxBodyBytesMiddleware(limit))
	r.POST("/x", func(c *gin.Context) {
		// The middleware must answer 413 before the handler runs.
		if _, err := io.ReadAll(c.Request.Body); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})

	send := func(size int) int {
		req := httptest.NewRequest(http.MethodPost, "/x", bytes.NewReader(bytes.Repeat([]byte("a"), size)))
		req.ContentLength = -1
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w.Code
	}
	if code := send(limit + 1); code != http.StatusRequestEntityTooLarge {
		t.Errorf("large body: got %d, want 413", code)
	}
	if code := send(128); code != http.StatusOK {
		t.Errorf("small body: got %d, want 200", code)
	}
}

// The engine rejects a request with a wrong API key or a cross-site Origin
// before it reads the body.
func TestBodyLimitAfterAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}

	for _, tc := range []struct {
		name   string
		apiKey string
		origin string
		want   int
	}{
		{"wrong key", "secret", "", http.StatusUnauthorized},
		{"cross-site origin", "", "https://evil.example", http.StatusForbidden},
	} {
		r := newServeEngine(registry, tc.apiKey, nil)
		body := &readSpy{Reader: strings.NewReader(`{}`)}
		req := httptest.NewRequest(http.MethodPost, "/chat", body)
		req.ContentLength = -1
		req.Host = "localhost"
		req.Header.Set("Content-Type", "text/plain")
		req.Header.Set(APIKeyHeader, "wrong")
		if tc.origin != "" {
			req.Header.Set("Origin", tc.origin)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != tc.want {
			t.Errorf("%s: got %d, want %d", tc.name, w.Code, tc.want)
		}
		if body.read {
			t.Errorf("%s: the engine read the body of a rejected request", tc.name)
		}
	}
}

// readSpy records a read of the body.
type readSpy struct {
	io.Reader
	read bool
}

func (r *readSpy) Read(p []byte) (int, error) {
	r.read = true
	return r.Reader.Read(p)
}

// chunked is a body with no length, as a chunked request has.
type chunked struct{ io.Reader }
