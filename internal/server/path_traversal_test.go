package restapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
)

// Each storage route that gets a name must reject a traversal name, not
// only DELETE. Most routes share storageError through the fsdb layer.
// The exists route validates in the handler, because its storage
// contract returns only a bool.
func TestStorageHandler_RejectsTraversalOnAllRoutes(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	root := t.TempDir()
	contextsDir := filepath.Join(root, "contexts")
	if err := os.MkdirAll(contextsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(root, "keep.txt")
	if err := os.WriteFile(marker, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}

	r := gin.New()
	NewContextsHandler(r, &fsdb.ContextsEntity{
		StorageEntity: &fsdb.StorageEntity{Label: "Contexts", Dir: contextsDir},
	})

	for _, tc := range []struct{ method, path string }{
		{http.MethodGet, "/contexts/%2e%2e"},
		{http.MethodDelete, "/contexts/%2e%2e"},
		{http.MethodDelete, "/contexts/.."},
		{http.MethodPost, "/contexts/%2e%2e"},
		{http.MethodPut, "/contexts/rename/%2e%2e/ok"},
		{http.MethodPut, "/contexts/rename/ok/%2e%2e"},
		{http.MethodGet, "/contexts/exists/%2e%2e"},
	} {
		var body io.Reader
		if tc.method == http.MethodPost {
			body = strings.NewReader("x")
		}
		w := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, tc.path, body)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("%s %s: got %d, want 400", tc.method, tc.path, w.Code)
		}
		if w.Header().Get("Strict-Transport-Security") == "" {
			t.Fatalf("%s %s: validation 400 lacks the HSTS header", tc.method, tc.path)
		}
	}

	if _, err := os.Stat(marker); err != nil {
		t.Fatalf("parent marker was deleted: %v", err)
	}
	if _, err := os.Stat(contextsDir); err != nil {
		t.Fatalf("contexts dir was deleted: %v", err)
	}
}

// A non-validation failure stays a 500. Its body must not show the
// filesystem path that is in the wrapped *os.PathError.
func TestStorageHandler_GenericErrorHidesPaths(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	dir := t.TempDir()
	r := gin.New()
	NewContextsHandler(r, &fsdb.ContextsEntity{
		StorageEntity: &fsdb.StorageEntity{Label: "Contexts", Dir: dir},
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/contexts/missing", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("got %d, want 500", w.Code)
	}
	if strings.Contains(w.Body.String(), dir) {
		t.Fatalf("500 body leaks the storage path: %s", w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "internal error") {
		t.Fatalf("500 body is not the generic message: %s", w.Body.String())
	}
}

// A pattern backend failure must use the shared storageError mapping.
// That is a JSON envelope with the generic message, never the wrapped
// os.PathError.
func TestPatternsHandler_BackendErrorHidesPaths(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	dir := t.TempDir()
	r := gin.New()
	NewPatternsHandler(r, &fsdb.PatternsEntity{
		StorageEntity:     &fsdb.StorageEntity{Label: "Patterns", Dir: dir, ItemIsDir: true},
		SystemPatternFile: "system.md",
	})

	for _, req := range []*http.Request{
		httptest.NewRequest(http.MethodGet, "/patterns/missing", nil),
		httptest.NewRequest(http.MethodPost, "/patterns/missing/apply", strings.NewReader(`{"input":"x"}`)),
	} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Fatalf("%s %s: got %d, want 500", req.Method, req.URL.Path, w.Code)
		}
		if strings.Contains(w.Body.String(), dir) {
			t.Fatalf("%s %s: 500 body leaks the storage path: %s", req.Method, req.URL.Path, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "internal error") {
			t.Fatalf("%s %s: 500 body is not the generic envelope: %s", req.Method, req.URL.Path, w.Body.String())
		}
	}
}

func TestPatternsHandler_RejectsPathTraversalSave(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	root := t.TempDir()
	patternsDir := filepath.Join(root, "patterns")
	if err := os.MkdirAll(patternsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	r := gin.New()
	NewPatternsHandler(r, &fsdb.PatternsEntity{
		StorageEntity:     &fsdb.StorageEntity{Label: "Patterns", Dir: patternsDir, ItemIsDir: true},
		SystemPatternFile: "system.md",
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/patterns/%2e%2e", strings.NewReader("pwned"))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("POST /patterns/%%2e%%2e: got %d, want 400", w.Code)
	}
	if _, err := os.Stat(filepath.Join(root, "system.md")); err == nil {
		t.Fatal("wrote system.md in the parent directory")
	}
}

func TestChatHandler_RejectsUnsafeNames(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	// The zero-value handler is an intentional seam. A request that goes
	// through pre-validation causes a nil panic in HandleChat. These
	// tests cannot pass on a request that got no validation.
	r := gin.New()
	r.POST("/chat", (&ChatHandler{}).HandleChat)

	cases := map[string]string{
		// Every prefix class loadPattern treats as a filesystem path
		"absolute pattern name":  `{"prompts":[{"patternName":"/etc/hosts","userInput":"x"}]}`,
		"home pattern name":      `{"prompts":[{"patternName":"~/secret.md","userInput":"x"}]}`,
		"relative pattern name":  `{"prompts":[{"patternName":"./secret.md","userInput":"x"}]}`,
		"backslash pattern name": `{"prompts":[{"patternName":"\\secret.md","userInput":"x"}]}`,
		// Context and session names get the same pre-validation
		"traversal context name": `{"prompts":[{"userInput":"x","contextName":"../keep.txt"}]}`,
		"traversal session name": `{"prompts":[{"userInput":"x","sessionName":".."}]}`,
		// The rejection loop stops at the first bad name, at each depth
		"second prompt path-like": `{"prompts":[{"userInput":"x"},{"patternName":"/etc/hosts","userInput":"y"}]}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/chat", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("got status %d, want 400", w.Code)
			}
		})
	}
}

func TestPatternsHandler_RejectsUnsafeNamesOnReadRoutes(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	r := gin.New()
	NewPatternsHandler(r, &fsdb.PatternsEntity{
		StorageEntity:     &fsdb.StorageEntity{Label: "Patterns", Dir: t.TempDir(), ItemIsDir: true},
		SystemPatternFile: "system.md",
	})

	for _, req := range []*http.Request{
		httptest.NewRequest(http.MethodGet, "/patterns/%2e%2e", nil),
		httptest.NewRequest(http.MethodPost, "/patterns/%2e%2e/apply", strings.NewReader(`{"input":"x"}`)),
		// Names that fail only ValidateStorageName, not the file-path check
		httptest.NewRequest(http.MethodGet, "/patterns/foo:bar", nil),
		httptest.NewRequest(http.MethodGet, "/patterns/NUL", nil),
		httptest.NewRequest(http.MethodPost, "/patterns/foo:bar/apply", strings.NewReader(`{"input":"x"}`)),
	} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("%s %s: got %d, want 400", req.Method, req.URL.Path, w.Code)
		}
	}
}

// A safe request must go through pre-validation. With the zero-value
// handler seam, a request that goes through causes a nil panic, and
// Recovery changes that into a 500. A 400 shows that validation
// rejected valid names.
func TestChatHandler_AcceptsBenignNames(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/chat", (&ChatHandler{}).HandleChat)

	w := httptest.NewRecorder()
	body := `{"prompts":[{"userInput":"x","patternName":"summarize","contextName":"myctx","sessionName":"mysession"}]}`
	req := httptest.NewRequest(http.MethodPost, "/chat", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code == http.StatusBadRequest {
		t.Fatalf("benign request was rejected with 400: %s", w.Body.String())
	}
}

// A non-loopback bind without an API key must fail closed before the
// server starts. A loopback bind operates without a key.
func TestRequireAPIKeyForBind(t *testing.T) {
	tests := []struct {
		address string
		apiKey  string
		wantErr bool
	}{
		{"127.0.0.1:8080", "", false},
		{"localhost:8080", "", false},
		{"localhost", "", false},
		{"[::1]:8080", "", false},
		{":8080", "", true}, // wildcard bind exposes every interface
		{"0.0.0.0:8080", "", true},
		{"[::]:8080", "", true},
		{"192.168.1.50:8080", "", true},
		{"example.com:8080", "", true},
		{":8080", "secret", false},
		{"0.0.0.0:8080", "secret", false},
	}
	for _, tt := range tests {
		t.Run(tt.address, func(t *testing.T) {
			err := requireAPIKeyForBind(tt.address, tt.apiKey)
			if (err != nil) != tt.wantErr {
				t.Fatalf("requireAPIKeyForBind(%q, %q) error = %v, wantErr %v", tt.address, tt.apiKey, err, tt.wantErr)
			}
		})
	}
}

// Serve and ServeOllama must return the fail-closed error and must not
// start an unauthenticated server on a non-loopback bind. The registry
// is nil, and a check that does not occur first causes a panic.
func TestServeFailsClosedOnNonLoopbackBind(t *testing.T) {
	if err := Serve(nil, ":0", ""); err == nil {
		t.Fatal("Serve on a wildcard bind without a key did not fail")
	}
	if err := ServeOllama(nil, ":0", "v", ""); err == nil {
		t.Fatal("ServeOllama on a wildcard bind without a key did not fail")
	}
}

func TestAPIKeyMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(APIKeyMiddleware("secret"))
	r.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })

	t.Run("missing key", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Fatalf("got %d, want 401", w.Code)
		}
	})

	t.Run("wrong key", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		req.Header.Set(APIKeyHeader, "wrong")
		r.ServeHTTP(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Fatalf("got %d, want 401", w.Code)
		}
	})

	t.Run("valid key", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		req.Header.Set(APIKeyHeader, "secret")
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("got %d, want 200", w.Code)
		}
	})
}
