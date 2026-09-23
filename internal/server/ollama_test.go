package restapi

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

func TestBuildFabricChatURL(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		want    string
		wantErr bool
	}{
		{
			name:    "empty address",
			addr:    "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "valid http URL",
			addr:    "http://localhost:8080",
			want:    "http://localhost:8080",
			wantErr: false,
		},
		{
			name:    "valid https URL",
			addr:    "https://api.example.com",
			want:    "https://api.example.com",
			wantErr: false,
		},
		{
			name:    "http URL with trailing slash",
			addr:    "http://localhost:8080/",
			want:    "http://localhost:8080",
			wantErr: false,
		},
		{
			name:    "malformed URL - missing host",
			addr:    "http://",
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed URL - port only with http",
			addr:    "https://:8080",
			want:    "",
			wantErr: true,
		},
		{
			name:    "colon-prefixed port",
			addr:    ":8080",
			want:    "http://127.0.0.1:8080",
			wantErr: false,
		},
		{
			name:    "bare host:port",
			addr:    "localhost:8080",
			want:    "http://localhost:8080",
			wantErr: false,
		},
		{
			name:    "bare hostname",
			addr:    "localhost",
			want:    "http://localhost",
			wantErr: false,
		},
		{
			name:    "IP address with port",
			addr:    "192.168.1.1:3000",
			want:    "http://192.168.1.1:3000",
			wantErr: false,
		},
		{
			name:    "bare address with path - invalid",
			addr:    "localhost:8080/some/path",
			want:    "",
			wantErr: true,
		},
		{
			name:    "bare hostname with path - invalid",
			addr:    "localhost/api",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildFabricChatURL(tt.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildFabricChatURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildFabricChatURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseOllamaNumCtx(t *testing.T) {
	tests := []struct {
		name    string
		options map[string]any
		want    int
		wantErr bool
		errMsg  string
	}{
		// --- Valid inputs ---
		{
			name:    "nil options",
			options: nil,
			want:    0,
			wantErr: false,
		},
		{
			name:    "empty options",
			options: map[string]any{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "num_ctx not present",
			options: map[string]any{"other_key": 123},
			want:    0,
			wantErr: false,
		},
		{
			name:    "num_ctx is null",
			options: map[string]any{"num_ctx": nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "valid int",
			options: map[string]any{"num_ctx": 4096},
			want:    4096,
			wantErr: false,
		},
		{
			name:    "valid float64 (whole number)",
			options: map[string]any{"num_ctx": float64(8192)},
			want:    8192,
			wantErr: false,
		},
		{
			name:    "valid float32 (whole number)",
			options: map[string]any{"num_ctx": float32(2048)},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "valid json.Number",
			options: map[string]any{"num_ctx": json.Number("16384")},
			want:    16384,
			wantErr: false,
		},
		{
			name:    "valid string",
			options: map[string]any{"num_ctx": "32768"},
			want:    32768,
			wantErr: false,
		},
		{
			name:    "valid int64",
			options: map[string]any{"num_ctx": int64(65536)},
			want:    65536,
			wantErr: false,
		},
		// --- Invalid inputs ---
		{
			name:    "float64 with fractional part",
			options: map[string]any{"num_ctx": 4096.5},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be an integer, got float with fractional part",
		},
		{
			name:    "float32 with fractional part",
			options: map[string]any{"num_ctx": float32(2048.75)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be an integer, got float with fractional part",
		},
		{
			name:    "negative int",
			options: map[string]any{"num_ctx": -100},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
		{
			name:    "zero int",
			options: map[string]any{"num_ctx": 0},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
		{
			name:    "negative float64",
			options: map[string]any{"num_ctx": float64(-500)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
		{
			name:    "negative float32",
			options: map[string]any{"num_ctx": float32(-250)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
		{
			name:    "non-numeric string",
			options: map[string]any{"num_ctx": "not-a-number"},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a valid number",
		},
		{
			name:    "invalid json.Number",
			options: map[string]any{"num_ctx": json.Number("invalid")},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a valid number",
		},
		{
			name:    "exceeds maximum allowed value",
			options: map[string]any{"num_ctx": 2000000},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx exceeds maximum allowed value",
		},
		{
			name:    "unsupported type (bool)",
			options: map[string]any{"num_ctx": true},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a number, got invalid type",
		},
		{
			name:    "unsupported type (slice)",
			options: map[string]any{"num_ctx": []int{1, 2, 3}},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a number, got invalid type",
		},
		// --- Edge cases ---
		{
			name:    "minimum valid value",
			options: map[string]any{"num_ctx": 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "maximum allowed value",
			options: map[string]any{"num_ctx": 1000000},
			want:    1000000,
			wantErr: false,
		},
		{
			name:    "very large float64 (overflow)",
			options: map[string]any{"num_ctx": float64(math.MaxFloat64)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx value out of range",
		},
		{
			name:    "large int64 exceeding maxInt on 32-bit",
			options: map[string]any{"num_ctx": int64(1 << 40)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx", // either "too large" or "exceeds maximum"
		},
		{
			name:    "long string gets truncated in error",
			options: map[string]any{"num_ctx": "this-is-a-very-long-string-that-should-be-truncated-in-the-error-message"},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a valid number",
		},
		// --- Special float values ---
		{
			name:    "float64 NaN",
			options: map[string]any{"num_ctx": math.NaN()},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		{
			name:    "float64 positive infinity",
			options: map[string]any{"num_ctx": math.Inf(1)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		{
			name:    "float64 negative infinity",
			options: map[string]any{"num_ctx": math.Inf(-1)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		{
			name:    "float32 NaN",
			options: map[string]any{"num_ctx": float32(math.NaN())},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		{
			name:    "float32 positive infinity",
			options: map[string]any{"num_ctx": float32(math.Inf(1))},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		{
			name:    "float32 negative infinity",
			options: map[string]any{"num_ctx": float32(math.Inf(-1))},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be a finite number",
		},
		// --- Negative int64 (32-bit wraparound prevention) ---
		{
			name:    "negative int64",
			options: map[string]any{"num_ctx": int64(-1000)},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
		{
			name:    "negative json.Number",
			options: map[string]any{"num_ctx": json.Number("-500")},
			want:    0,
			wantErr: true,
			errMsg:  "num_ctx must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOllamaNumCtx(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOllamaNumCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errMsg != "" {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("parseOllamaNumCtx() error message = %q, want to contain %q", err.Error(), tt.errMsg)
				}
			}
			if got != tt.want {
				t.Errorf("parseOllamaNumCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOllamaEngine_APIKeyWiring(t *testing.T) {
	gin.SetMode(gin.TestMode)
	registry := &core.PluginRegistry{Db: fsdb.NewDb(t.TempDir())}

	getVersion := func(r *gin.Engine, key string) int {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/version", nil)
		if key != "" {
			req.Header.Set(APIKeyHeader, key)
		}
		r.ServeHTTP(w, req)
		return w.Code
	}

	withKey := newOllamaEngine(registry, ":0", "test-version", "secret")
	if code := getVersion(withKey, ""); code != http.StatusUnauthorized {
		t.Fatalf("no key presented: got %d, want 401", code)
	}
	if code := getVersion(withKey, "secret"); code != http.StatusOK {
		t.Fatalf("valid key presented: got %d, want 200", code)
	}

	withoutKey := newOllamaEngine(registry, ":0", "test-version", "")
	if code := getVersion(withoutKey, ""); code != http.StatusOK {
		t.Fatalf("no key configured: got %d, want 200", code)
	}
}

func TestOllamaChat_ForwardsAPIKeyToChat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Make the loopback /chat route with the middleware installed, the
	// same as newOllamaEngine makes it when --api-key is set.
	upstream := gin.New()
	upstream.Use(APIKeyMiddleware("secret"))
	upstream.POST("/chat", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		fmt.Fprint(c.Writer, "data: {\"type\":\"content\",\"format\":\"markdown\",\"content\":\"hi\"}\n\n")
	})
	server := httptest.NewServer(upstream)
	defer server.Close()

	chatRequest := func(key string) int {
		r := gin.New()
		conv := APIConvert{addr: &server.URL, apiKey: key}
		r.POST("/api/chat", conv.ollamaChat)

		w := httptest.NewRecorder()
		body := `{"model":"test:latest","messages":[{"role":"user","content":"hi"}]}`
		req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	if code := chatRequest("secret"); code != http.StatusOK {
		t.Fatalf("matching key: got %d, want 200", code)
	}
	if code := chatRequest("wrong"); code != http.StatusUnauthorized {
		t.Fatalf("wrong key: got %d, want 401", code)
	}
}

// The self-forward client must not use a proxy. A proxy gets the
// configured API key, and the operator did not configure that host.
// The redirect test below covers the no-redirect property.
func TestFabricChatClient_NoProxy(t *testing.T) {
	transport, ok := fabricChatClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport is %T, want *http.Transport", fabricChatClient.Transport)
	}
	if transport.Proxy != nil {
		t.Fatal("self-forward transport has a proxy configured")
	}
}

// An upstream redirect must show as an upstream error. The client must
// not go to the redirect target with the API key.
func TestOllamaChat_DoesNotFollowUpstreamRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	redirectTargetHit := false
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		redirectTargetHit = true
	}))
	defer target.Close()

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL+"/chat", http.StatusFound)
	}))
	defer upstream.Close()

	r := gin.New()
	conv := APIConvert{addr: &upstream.URL, apiKey: "secret"}
	r.POST("/api/chat", conv.ollamaChat)

	w := httptest.NewRecorder()
	body := `{"model":"test:latest","messages":[{"role":"user","content":"hi"}]}`
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if redirectTargetHit {
		t.Fatal("the redirect target was contacted")
	}
	if w.Code != http.StatusFound {
		t.Fatalf("got %d, want the upstream 302 surfaced as an error", w.Code)
	}
}

// Malformed client JSON is a client error. The answer is a 400 with a
// stable generic message, not a 500.
func TestOllamaChat_MalformedJSONIs400(t *testing.T) {
	gin.SetMode(gin.TestMode)

	addr := ":0"
	r := gin.New()
	conv := APIConvert{addr: &addr}
	r.POST("/api/chat", conv.ollamaChat)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader("{not json"))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want 400", w.Code)
	}
}

// An upstream that is not available is a 500. The body must not contain
// the raw transport error, because that error shows the internal
// upstream URL.
func TestOllamaChat_UpstreamFailureHidesDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := httptest.NewServer(http.NotFoundHandler())
	url := server.URL
	server.Close() // nothing listens on url anymore

	r := gin.New()
	conv := APIConvert{addr: &url}
	r.POST("/api/chat", conv.ollamaChat)

	w := httptest.NewRecorder()
	body := `{"model":"test:latest","messages":[{"role":"user","content":"hi"}]}`
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("got %d, want 500", w.Code)
	}
	got := w.Body.String()
	if strings.Contains(got, "dial tcp") || strings.Contains(got, strings.TrimPrefix(url, "http://")) {
		t.Fatalf("500 body leaks transport details: %s", got)
	}
}

// With no configured key, the forwarded request must not contain the header.
func TestOllamaChat_NoKeyOmitsHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var gotHeader string
	upstream := gin.New()
	upstream.POST("/chat", func(c *gin.Context) {
		gotHeader = c.GetHeader(APIKeyHeader)
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		fmt.Fprint(c.Writer, "data: {\"type\":\"content\",\"format\":\"markdown\",\"content\":\"hi\"}\n\n")
	})
	server := httptest.NewServer(upstream)
	defer server.Close()

	r := gin.New()
	conv := APIConvert{addr: &server.URL, apiKey: ""}
	r.POST("/api/chat", conv.ollamaChat)

	w := httptest.NewRecorder()
	body := `{"model":"test:latest","messages":[{"role":"user","content":"hi"}]}`
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("got %d, want 200", w.Code)
	}
	if gotHeader != "" {
		t.Fatalf("X-API-Key was forwarded with no key configured: %q", gotHeader)
	}
}
