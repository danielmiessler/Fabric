package template

import (
	"os"
	"path/filepath"
	"testing"
)

// TestExtensionManager is the main test suite for ExtensionManager
func TestExtensionManager(t *testing.T) {
	// Create temporary directory for tests
	tmpDir, err := os.MkdirTemp("", "fabric-ext-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test extension config
	testConfig := filepath.Join(tmpDir, "test-extension.yaml")
	testScript := filepath.Join(tmpDir, "test-script.sh")

	// Create test script
	scriptContent := `#!/bin/bash
if [ "$1" = "echo" ]; then
    echo "Hello, $2!"
fi`

	err = os.WriteFile(testScript, []byte(scriptContent), 0755)
	if err != nil {
		t.Fatalf("Failed to create test script: %v", err)
	}

	// Create test config
	configContent := `name: test-extension
executable: ` + testScript + `
type: executable
timeout: 30s
description: "Test extension"
version: "1.0.0"
operations:
  echo:
    cmd_template: "{{executable}} echo {{1}}"
`

	err = os.WriteFile(testConfig, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	// Initialize manager
	manager := NewExtensionManager(tmpDir)

	// Test cases
	t.Run("RegisterExtension", func(t *testing.T) {
		err := manager.RegisterExtension(testConfig)
		if err != nil {
			t.Errorf("Failed to register extension: %v", err)
		}
	})

	t.Run("ListExtensions", func(t *testing.T) {
		err := manager.ListExtensions()
		if err != nil {
			t.Errorf("Failed to list extensions: %v", err)
		}
		// Note: Output validation would require capturing stdout
	})

	t.Run("ProcessExtension", func(t *testing.T) {
		output, err := manager.ProcessExtension("test-extension", "echo", "World")
		if err != nil {
			t.Errorf("Failed to process extension: %v", err)
		}
		expected := "Hello, World!\n"
		if output != expected {
			t.Errorf("Expected output %q, got %q", expected, output)
		}
	})

	t.Run("RemoveExtension", func(t *testing.T) {
		err := manager.RemoveExtension("test-extension")
		if err != nil {
			t.Errorf("Failed to remove extension: %v", err)
		}

		// Verify extension is removed by trying to process it
		_, err = manager.ProcessExtension("test-extension", "echo", "World")
		if err == nil {
			t.Error("Expected error processing removed extension, got nil")
		}
	})
}

// TestExtensionManagerErrors tests error cases
func TestExtensionManagerErrors(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fabric-ext-test-errors-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	manager := NewExtensionManager(tmpDir)

	t.Run("RegisterNonexistentConfig", func(t *testing.T) {
		err := manager.RegisterExtension("/nonexistent/config.yaml")
		if err == nil {
			t.Error("Expected error registering nonexistent config, got nil")
		}
	})

	t.Run("ProcessNonexistentExtension", func(t *testing.T) {
		_, err := manager.ProcessExtension("nonexistent", "echo", "test")
		if err == nil {
			t.Error("Expected error processing nonexistent extension, got nil")
		}
	})

	t.Run("RemoveNonexistentExtension", func(t *testing.T) {
		err := manager.RemoveExtension("nonexistent")
		if err == nil {
			t.Error("Expected error removing nonexistent extension, got nil")
		}
	})
}

// TestExtensionManagerWithInvalidConfig tests handling of invalid configurations
func TestExtensionManagerWithInvalidConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fabric-ext-test-invalid-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	invalidConfig := filepath.Join(tmpDir, "invalid-extension.yaml")

	// Test cases with different invalid configurations
	testCases := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name: "MissingExecutable",
			config: `name: invalid-extension
type: executable
timeout: 30s`,
			wantErr: true,
		},
		{
			name: "InvalidTimeout",
			config: `name: invalid-extension
executable: /bin/echo
type: executable
timeout: invalid`,
			wantErr: true,
		},
		{
			name: "EmptyName",
			config: `name: ""
executable: /bin/echo
type: executable
timeout: 30s`,
			wantErr: true,
		},
	}

	manager := NewExtensionManager(tmpDir)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := os.WriteFile(invalidConfig, []byte(tc.config), 0644)
			if err != nil {
				t.Fatalf("Failed to create invalid config file: %v", err)
			}

			err = manager.RegisterExtension(invalidConfig)
			if tc.wantErr && err == nil {
				t.Error("Expected error registering invalid config, got nil")
			} else if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error registering config: %v", err)
			}
		})
	}
}
