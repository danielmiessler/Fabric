package restapi

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORSMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	newEngine := func(origins []string) *gin.Engine {
		r := gin.New()
		r.Use(CORSMiddleware(origins))
		r.Use(APIKeyMiddleware("secret"))
		r.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })
		return r
	}
	send := func(r *gin.Engine, method, origin, key string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(method, "/ping", nil)
		req.Header.Set("Origin", origin)
		if key != "" {
			req.Header.Set(APIKeyHeader, key)
		}
		r.ServeHTTP(w, req)
		return w
	}

	r := newEngine([]string{"http://localhost:1420"})

	w := send(r, http.MethodGet, "http://localhost:1420", "secret")
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:1420" {
		t.Fatalf("listed origin: Allow-Origin = %q, want the origin", got)
	}
	if got := w.Header().Get("Access-Control-Allow-Headers"); got != "Content-Type, "+APIKeyHeader {
		t.Fatalf("Allow-Headers = %q", got)
	}
	if got := w.Header().Get("Vary"); got != "Origin" {
		t.Fatalf("Vary = %q, want Origin", got)
	}

	w = send(r, http.MethodGet, "http://evil.example", "secret")
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("unlisted origin: Allow-Origin = %q, want empty", got)
	}

	w = send(r, http.MethodOptions, "http://evil.example", "")
	if got := w.Header().Get("Access-Control-Allow-Origin"); w.Code != http.StatusNoContent || got != "" {
		t.Fatalf("unlisted preflight: status = %d, Allow-Origin = %q, want 204 and empty", w.Code, got)
	}

	// A preflight request has no API key. It must get 204, not 401.
	w = send(r, http.MethodOptions, "http://localhost:1420", "")
	if w.Code != http.StatusNoContent {
		t.Fatalf("preflight: status = %d, want 204", w.Code)
	}

	w = send(newEngine([]string{"*"}), http.MethodGet, "http://any.example", "secret")
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "http://any.example" {
		t.Fatalf("wildcard: Allow-Origin = %q, want the origin", got)
	}

	w = send(newEngine([]string{"*"}), http.MethodGet, "null", "secret")
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("wildcard with null origin: Allow-Origin = %q, want empty", got)
	}
}

func TestCleanCORSOrigins(t *testing.T) {
	got, err := cleanCORSOrigins([]string{"http://a.example", " http://b.example/", "", " "}, "")
	if err != nil || !slices.Equal(got, []string{"http://a.example", "http://b.example"}) {
		t.Fatalf("got %q, %v", got, err)
	}
	if got, err = cleanCORSOrigins([]string{""}, ""); err != nil || len(got) != 0 {
		t.Fatalf("blank list: got %q, %v, want empty", got, err)
	}
	if _, err = cleanCORSOrigins([]string{" *"}, ""); err == nil {
		t.Fatal("wildcard without an API key did not fail")
	}
	if _, err = cleanCORSOrigins([]string{"*"}, "secret"); err != nil {
		t.Fatalf("wildcard with an API key: %v", err)
	}
	// Both servers refuse the wildcard before Run, also on a loopback bind.
	if err = Serve(nil, "127.0.0.1:0", "", []string{"*"}); err == nil {
		t.Fatal("Serve with a wildcard origin and no key did not fail")
	}
	if err = ServeOllama(nil, "127.0.0.1:0", "v", "", []string{"*"}); err == nil {
		t.Fatal("ServeOllama with a wildcard origin and no key did not fail")
	}
}
