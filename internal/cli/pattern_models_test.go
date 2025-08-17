package cli

import (
	"path/filepath"
	"testing"
)

func TestGetPatternModelFile(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tmp)
	path, err := getPatternModelFile()
	if err != nil {
		t.Fatalf("getPatternModelFile returned error: %v", err)
	}
	expected := filepath.Join(tmp, "fabric", "pattern_models.yaml")
	if path != expected {
		t.Fatalf("expected %s, got %s", expected, path)
	}
}
