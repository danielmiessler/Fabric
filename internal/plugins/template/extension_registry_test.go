package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRegistryPersistence(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fabric-ext-registry-persist-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test executable
	execPath := filepath.Join(tmpDir, "test-exec.sh")
	execContent := []byte("#!/bin/bash\necho \"test\"")
	err = os.WriteFile(execPath, execContent, 0755)
	if err != nil {
		t.Fatalf("Failed to create test executable: %v", err)
	}

	// Create valid config
	configContent := `name: test-extension
executable: ` + execPath + `
type: executable
timeout: 30s
operations:
  test:
    cmd_template: "{{executable}} {{operation}}"`

	configPath := filepath.Join(tmpDir, "test-extension.yaml")
	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	// Test registry persistence
	t.Run("SaveAndReload", func(t *testing.T) {
		// Create and populate first registry
		registry1 := NewExtensionRegistry(tmpDir)
		err := registry1.Register(configPath)
		if err != nil {
			t.Fatalf("Failed to register extension: %v", err)
		}

		// Create new registry instance and verify it loads the saved state
		registry2 := NewExtensionRegistry(tmpDir)
		ext, err := registry2.GetExtension("test-extension")
		if err != nil {
			t.Fatalf("Failed to get extension from reloaded registry: %v", err)
		}
		if ext.Name != "test-extension" {
			t.Errorf("Expected extension name 'test-extension', got %q", ext.Name)
		}
	})

	// Test hash verification
	t.Run("HashVerification", func(t *testing.T) {
		registry := NewExtensionRegistry(tmpDir)

		// Modify executable after registration
		modifiedExecContent := []byte("#!/bin/bash\necho \"modified\"")
		err := os.WriteFile(execPath, modifiedExecContent, 0755)
		if err != nil {
			t.Fatalf("Failed to modify executable: %v", err)
		}

		_, err = registry.GetExtension("test-extension")
		if err == nil {
			t.Error("Expected error when executable modified, got nil")
		}
	})
}
