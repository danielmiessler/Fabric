package restapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func newConfigTestEngine(t *testing.T) (*gin.Engine, *fsdb.Db) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	// The handler calls os.Setenv. t.Setenv puts back the old value after the test.
	t.Setenv("OPENAI_API_KEY", "")
	db := fsdb.NewDb(t.TempDir())
	r := gin.New()
	NewConfigHandler(r, db)
	return r, db
}

func postConfig(r *gin.Engine, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/config/update", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// An update of one key keeps the keys that the config form does not know.
func TestUpdateConfig_KeepsOtherKeys(t *testing.T) {
	r, db := newConfigTestEngine(t)

	seed := map[string]string{"CODEX_OAUTH_TOKEN": "keep-me", "OPENAI_API_KEY": "sk-old"}
	if err := godotenv.Write(seed, db.EnvFilePath); err != nil {
		t.Fatalf("seed: %v", err)
	}

	if w := postConfig(r, `{"openai_api_key":"sk-newkey12345"}`); w.Code != http.StatusOK {
		t.Fatalf("update: got %d, want 200 (%s)", w.Code, w.Body.String())
	}

	final, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("read .env: %v", err)
	}
	if final["CODEX_OAUTH_TOKEN"] != "keep-me" {
		t.Errorf("other key changed: %q", final["CODEX_OAUTH_TOKEN"])
	}
	if final["OPENAI_API_KEY"] != "sk-newkey12345" {
		t.Errorf("known key not written: %q", final["OPENAI_API_KEY"])
	}
	info, err := os.Stat(db.EnvFilePath)
	if err != nil {
		t.Fatalf("stat .env: %v", err)
	}
	if mode := info.Mode().Perm(); mode != 0o600 {
		t.Errorf("mode = %o, want 600", mode)
	}
}

// A value with a newline, a carriage return or a NUL gets 400, and the file does not change.
func TestUpdateConfig_RejectsLineBreakInValue(t *testing.T) {
	for _, value := range []string{`sk-abc\nEXTRA_KEY=1`, `sk-abc\rEXTRA_KEY=1`, `sk-abc\u0000`} {
		t.Run(value, func(t *testing.T) {
			r, db := newConfigTestEngine(t)

			if err := godotenv.Write(map[string]string{"OPENAI_API_KEY": "sk-old"}, db.EnvFilePath); err != nil {
				t.Fatalf("seed: %v", err)
			}
			before, err := os.ReadFile(db.EnvFilePath)
			if err != nil {
				t.Fatalf("read .env: %v", err)
			}

			if w := postConfig(r, `{"openai_api_key":"`+value+`"}`); w.Code != http.StatusBadRequest {
				t.Fatalf("got %d, want 400", w.Code)
			}

			after, err := os.ReadFile(db.EnvFilePath)
			if err != nil {
				t.Fatalf("read .env: %v", err)
			}
			if !bytes.Equal(before, after) {
				t.Errorf(".env changed:\nbefore: %q\nafter:  %q", before, after)
			}
		})
	}
}

// A value of only spaces changes neither the .env file nor the process
// environment.
func TestUpdateConfig_SkipsBlankValue(t *testing.T) {
	r, db := newConfigTestEngine(t)
	if err := godotenv.Write(map[string]string{"OPENAI_API_KEY": "sk-old"}, db.EnvFilePath); err != nil {
		t.Fatalf("seed: %v", err)
	}

	if w := postConfig(r, `{"openai_api_key":"   "}`); w.Code != http.StatusOK {
		t.Fatalf("update: got %d, want 200 (%s)", w.Code, w.Body.String())
	}

	final, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("read .env: %v", err)
	}
	if final["OPENAI_API_KEY"] != "sk-old" {
		t.Errorf(".env key changed: %q", final["OPENAI_API_KEY"])
	}
	if got := os.Getenv("OPENAI_API_KEY"); strings.TrimSpace(got) == "" && got != "" {
		t.Errorf("process environment got the blank value %q", got)
	}
}

func getConfig(t *testing.T, r *gin.Engine) map[string]string {
	t.Helper()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/config", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("get: got %d, want 200 (%s)", w.Code, w.Body.String())
	}
	var got map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return got
}

// GET /config masks each set key and returns the URLs as they are.
func TestGetConfig_MasksKeys(t *testing.T) {
	r, db := newConfigTestEngine(t)
	if err := godotenv.Write(map[string]string{}, db.EnvFilePath); err != nil {
		t.Fatalf("seed: %v", err)
	}
	env := map[string]string{
		"openai":        "OPENAI_API_KEY",
		"anthropic":     "ANTHROPIC_API_KEY",
		"groq":          "GROQ_API_KEY",
		"mistral":       "MISTRAL_API_KEY",
		"gemini":        "GEMINI_API_KEY",
		"openrouter":    "OPENROUTER_API_KEY",
		"trustedrouter": "TRUSTEDROUTER_API_KEY",
		"silicon":       "SILICON_API_KEY",
		"deepseek":      "DEEPSEEK_API_KEY",
		"grokai":        "GROKAI_API_KEY",
	}
	for _, name := range env {
		t.Setenv(name, "sk-secret-"+name)
	}
	t.Setenv("OLLAMA_URL", "http://localhost:11434")
	t.Setenv("LM_STUDIO_API_BASE_URL", "http://localhost:1234/v1")

	got := getConfig(t, r)
	if len(got) != 12 {
		t.Errorf("got %d keys, want 12: %v", len(got), got)
	}
	for field := range env {
		if got[field] != maskedValue {
			t.Errorf("%s: got %q, want the mask", field, got[field])
		}
	}
	if got["ollama"] != "http://localhost:11434" || got["lmstudio"] != "http://localhost:1234/v1" {
		t.Errorf("URLs changed: ollama %q, lmstudio %q", got["ollama"], got["lmstudio"])
	}
}

// With no .env file, GET /config returns the same 12 keys, all empty.
func TestGetConfig_NoEnvFile(t *testing.T) {
	r, _ := newConfigTestEngine(t)
	got := getConfig(t, r)
	if len(got) != 12 {
		t.Errorf("got %d keys, want 12: %v", len(got), got)
	}
	for field, value := range got {
		if value != "" {
			t.Errorf("%s: got %q, want empty", field, value)
		}
	}
	if _, ok := got["lmstudio"]; !ok {
		t.Errorf("lmstudio key missing")
	}
}
