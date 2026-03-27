package strategy

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
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

	// Create a strategy file WITHOUT .json extension
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
