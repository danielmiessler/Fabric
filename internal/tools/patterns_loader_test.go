package tools

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// Configure runs on every fabric invocation via the plugin registry. It must
// not create the patterns temp directory; only PopulateDB uses it.
func TestConfigureDoesNotCreateTempDir(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("TMPDIR", tmp)
	t.Setenv("TMP", tmp)
	t.Setenv("TEMP", tmp)

	loader := NewPatternsLoader(&fsdb.PatternsEntity{
		StorageEntity: &fsdb.StorageEntity{Dir: t.TempDir()},
	})
	if err := loader.Configure(); err != nil {
		t.Fatalf("Configure() failed: %v", err)
	}

	matches, err := filepath.Glob(filepath.Join(tmp, "fabric-patterns-*"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 0 {
		t.Errorf("Configure() created temp directories: %v", matches)
	}
}

// PopulateDB must create the temp directory lazily and remove it when done,
// even on failure.
func TestPopulateDBCleansUpTempDir(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("TMPDIR", tmp)
	t.Setenv("TMP", tmp)
	t.Setenv("TEMP", tmp)

	loader := NewPatternsLoader(&fsdb.PatternsEntity{
		StorageEntity: &fsdb.StorageEntity{Dir: t.TempDir()},
	})
	if err := loader.Configure(); err != nil {
		t.Fatalf("Configure() failed: %v", err)
	}
	// Point at an invalid repo so PopulateDB fails fast without network.
	loader.DefaultGitRepoUrl.Value = filepath.Join(t.TempDir(), "no-such-repo")

	if err := loader.PopulateDB(); err == nil {
		t.Fatal("PopulateDB() unexpectedly succeeded with invalid repo")
	}

	entries, err := os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		t.Errorf("PopulateDB() left temp entry behind: %v", e.Name())
	}
}
