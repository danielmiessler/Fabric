package strategy

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	debuglog "github.com/danielmiessler/fabric/internal/log"
)

func TestLoadStrategy_ValidName(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}

	strategyPath := filepath.Join(strategyDir, "test-strategy.json")
	if err := os.WriteFile(strategyPath, []byte(`{"name":"test","description":"desc","prompt":"PROMPT"}`), 0o644); err != nil {
		t.Fatalf("failed to write strategy: %v", err)
	}

	s, err := LoadStrategy("test-strategy")
	if err != nil {
		t.Fatalf("LoadStrategy returned error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil strategy")
	}
	if s.Prompt != "PROMPT" {
		t.Errorf("expected prompt %q, got %q", "PROMPT", s.Prompt)
	}
}

func TestLoadStrategy_EmptyName(t *testing.T) {
	s, err := LoadStrategy("")
	if err != nil {
		t.Fatalf("expected no error for empty name, got: %v", err)
	}
	if s != nil {
		t.Fatal("expected nil strategy for empty name")
	}
}

func TestLoadStrategy_PathTraversal(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}

	// Create a file outside the strategy directory that an attacker might target
	outsideFile := filepath.Join(homeDir, ".config", "fabric", "secret.json")
	if err := os.WriteFile(outsideFile, []byte(`{"prompt":"STOLEN"}`), 0o644); err != nil {
		t.Fatalf("failed to write outside file: %v", err)
	}

	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "dot-dot traversal",
			filename: "../secret",
		},
		{
			name:     "deep traversal",
			filename: "../../etc/passwd",
		},
		{
			name:     "dot-dot with json extension would match",
			filename: "../secret.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := LoadStrategy(tt.filename)
			if err == nil {
				t.Fatalf("expected error for traversal filename %q, but got strategy: %+v", tt.filename, s)
			}
			if !strings.Contains(err.Error(), "outside the strategy directory") {
				// It's also fine if it's "not found" — the point is it doesn't succeed
				t.Logf("error was: %v (acceptable if not a path traversal success)", err)
			}
			if s != nil {
				t.Fatalf("expected nil strategy for traversal attempt, got: %+v", s)
			}
		})
	}
}

func TestLoadStrategy_WithoutExtension(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}

	strategyPath := filepath.Join(strategyDir, "bare-strategy")
	if err := os.WriteFile(strategyPath, []byte(`{"name":"bare","description":"no ext","prompt":"BARE PROMPT"}`), 0o644); err != nil {
		t.Fatalf("failed to write strategy: %v", err)
	}

	s, err := LoadStrategy("bare-strategy")
	if err != nil {
		t.Fatalf("LoadStrategy returned error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil strategy")
	}
	if s.Prompt != "BARE PROMPT" {
		t.Errorf("expected prompt %q, got %q", "BARE PROMPT", s.Prompt)
	}
}

func TestLoadStrategy_InvalidJSON(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}

	strategyPath := filepath.Join(strategyDir, "bad.json")
	if err := os.WriteFile(strategyPath, []byte(`{not valid json`), 0o644); err != nil {
		t.Fatalf("failed to write strategy: %v", err)
	}

	_, err := LoadStrategy("bad")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestLoadStrategy_NotFound(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}

	_, err := LoadStrategy("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent strategy")
	}
}

func TestLoadStrategy_NameWithSeparatorFailsBeforeStat(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(homeDir, ".config", "fabric", "x.json"), []byte(`{}`), 0o644); err != nil {
		t.Fatalf("failed to write outside file: %v", err)
	}

	// A file that exists and a file that does not exist must give the same
	// error, so the error does not show which files exist.
	_, errExists := LoadStrategy("../x")
	_, errMissing := LoadStrategy("../missing")
	if errExists == nil || errMissing == nil {
		t.Fatalf("expected errors, got %v and %v", errExists, errMissing)
	}
	if strings.ReplaceAll(errExists.Error(), "../x", "N") != strings.ReplaceAll(errMissing.Error(), "../missing", "N") {
		t.Errorf("errors differ: %q and %q", errExists, errMissing)
	}
}

// TestLoadAllFiles_SkipsFileThatDoesNotLoad checks that a strategy file that
// does not load does not hide the strategies after it.
func TestLoadAllFiles_SkipsFileThatDoesNotLoad(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy dir: %v", err)
	}
	// LoadStrategy refuses the name "a." because it ends with a dot, and
	// bad.json is not JSON. Both come before good.json in the walk.
	files := map[string]string{
		"a..json":   `{"description":"desc","prompt":"PROMPT"}`,
		"bad.json":  `{`,
		"good.json": `{"description":"desc","prompt":"PROMPT"}`,
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(strategyDir, name), []byte(content), 0o644); err != nil {
			t.Fatalf("failed to write %s: %v", name, err)
		}
	}

	var logged bytes.Buffer
	debuglog.SetOutput(&logged)
	t.Cleanup(func() { debuglog.SetOutput(os.Stderr) })

	strategies, err := LoadAllFiles()
	if err != nil {
		t.Fatalf("LoadAllFiles returned error: %v", err)
	}
	if _, ok := strategies["good"]; !ok || len(strategies) != 1 {
		t.Errorf("strategies = %v, want only good", strategies)
	}
	for _, name := range []string{"a..json", "bad.json"} {
		if !strings.Contains(logged.String(), name) {
			t.Errorf("log %q does not tell about %s", logged.String(), name)
		}
	}
}
