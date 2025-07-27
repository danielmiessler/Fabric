package core

import (
	"os"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

func TestSaveEnvFile(t *testing.T) {
	db := fsdb.NewDb(os.TempDir())
	registry, err := NewPluginRegistry(db)
	if err != nil {
		t.Fatalf("NewPluginRegistry() error = %v", err)
	}

	err = registry.SaveEnvFile()
	if err != nil {
		t.Fatalf("SaveEnvFile() error = %v", err)
	}
}
