package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtensionExecutor(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fabric-ext-executor-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test script that has both stdout and file output modes
	testScript := filepath.Join(tmpDir, "test-script.sh")
	scriptContent := `#!/bin/bash
case "$1" in
    "stdout")
        echo "Hello, $2!"
        ;;
    "file")
        echo "Hello, $2!" > "$3"
        echo "$3"  # Print the filename for path_from_stdout
        ;;
    *)
        echo "Unknown command" >&2
        exit 1
        ;;
esac`

	if err := os.WriteFile(testScript, []byte(scriptContent), 0755); err != nil {
		t.Fatalf("Failed to create test script: %v", err)
	}

	// Create registry and register our test extensions
	registry := NewExtensionRegistry(tmpDir)
	executor := NewExtensionExecutor(registry)

	// Test stdout-based extension
	t.Run("StdoutExecution", func(t *testing.T) {
		configPath := filepath.Join(tmpDir, "stdout-extension.yaml")
		configContent := `name: stdout-test
executable: ` + testScript + `
type: executable
timeout: 30s
operations:
  greet:
    cmd_template: "{{executable}} stdout {{1}}"
config:
  output:
    method: stdout`

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to create config: %v", err)
		}

		if err := registry.Register(configPath); err != nil {
			t.Fatalf("Failed to register extension: %v", err)
		}

		output, err := executor.Execute("stdout-test", "greet", "World")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}

		expected := "Hello, World!\n"
		if output != expected {
			t.Errorf("Expected output %q, got %q", expected, output)
		}
	})

	// Test file-based extension
	t.Run("FileExecution", func(t *testing.T) {
		configPath := filepath.Join(tmpDir, "file-extension.yaml")
		configContent := `name: file-test
executable: ` + testScript + `
type: executable
timeout: 30s
operations:
  greet:
    cmd_template: "{{executable}} file {{1}} {{2}}"
config:
  output:
    method: file
    file_config:
      cleanup: true
      path_from_stdout: true`

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to create config: %v", err)
		}

		if err := registry.Register(configPath); err != nil {
			t.Fatalf("Failed to register extension: %v", err)
		}

		output, err := executor.Execute("file-test", "greet", "World|/tmp/test.txt")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}

		expected := "Hello, World!\n"
		if output != expected {
			t.Errorf("Expected output %q, got %q", expected, output)
		}
	})

	// Test execution errors
	t.Run("ExecutionErrors", func(t *testing.T) {
		// Test with non-existent extension
		_, err := executor.Execute("nonexistent", "test", "value")
		if err == nil {
			t.Error("Expected error executing non-existent extension, got nil")
		}

		// Test with invalid command that should exit non-zero
		configPath := filepath.Join(tmpDir, "error-extension.yaml")
		configContent := `name: error-test
executable: ` + testScript + `
type: executable
timeout: 30s
operations:
  invalid:
    cmd_template: "{{executable}} invalid {{1}}"
config:
  output:
    method: stdout`

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to create config: %v", err)
		}

		if err := registry.Register(configPath); err != nil {
			t.Fatalf("Failed to register extension: %v", err)
		}

		_, err = executor.Execute("error-test", "invalid", "test")
		if err == nil {
			t.Error("Expected error from invalid command, got nil")
		}
		if !strings.Contains(err.Error(), "Unknown command") {
			t.Errorf("Expected 'Unknown command' in error, got: %v", err)
		}
	})
}

func TestFixedFileExtensionExecutor(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fabric-ext-executor-fixed-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test script
	testScript := filepath.Join(tmpDir, "test-script.sh")
	scriptContent := `#!/bin/bash
case "$1" in
    "write")
        echo "Hello, $2!" > "$3"
        ;;
    "append")
        echo "Hello, $2!" >> "$3"
        ;;
    "large")
        for i in {1..1000}; do
            echo "Line $i" >> "$3"
        done
        ;;
    "error")
        echo "Error message" >&2
        exit 1
        ;;
    *)
        echo "Unknown command" >&2
        exit 1
        ;;
esac`

	if err := os.WriteFile(testScript, []byte(scriptContent), 0755); err != nil {
		t.Fatalf("Failed to create test script: %v", err)
	}

	registry := NewExtensionRegistry(tmpDir)
	executor := NewExtensionExecutor(registry)

	// Helper function to create and register extension
	createExtension := func(name, opName, cmdTemplate string, config map[string]interface{}) error {
		configPath := filepath.Join(tmpDir, name+".yaml")
		configContent := `name: ` + name + `
executable: ` + testScript + `
type: executable
timeout: 30s
operations:
  ` + opName + `:
    cmd_template: "` + cmdTemplate + `"
config:
  output:
    method: file
    file_config:`

		// Add config options
		for k, v := range config {
			configContent += "\n      " + k + ": " + strings.TrimSpace(v.(string))
		}

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			return err
		}

		return registry.Register(configPath)
	}

	// Test basic fixed file output
	t.Run("BasicFixedFile", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "output.txt")
		config := map[string]interface{}{
			"output_file": `"output.txt"`,
			"work_dir":    `"` + tmpDir + `"`,
			"cleanup":     "true",
		}

		err := createExtension("basic-test", "write",
			"{{executable}} write {{1}} "+outputFile, config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		output, err := executor.Execute("basic-test", "write", "World")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}

		expected := "Hello, World!\n"
		if output != expected {
			t.Errorf("Expected output %q, got %q", expected, output)
		}
	})

	// Test no work_dir specified
	t.Run("NoWorkDir", func(t *testing.T) {
		config := map[string]interface{}{
			"output_file": `"direct-output.txt"`,
			"cleanup":     "true",
		}

		err := createExtension("no-workdir-test", "write",
			"{{executable}} write {{1}} direct-output.txt", config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("no-workdir-test", "write", "World")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}
	})

	// Test cleanup behavior
	t.Run("CleanupBehavior", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "cleanup-test.txt")

		// Test with cleanup enabled
		config := map[string]interface{}{
			"output_file": `"cleanup-test.txt"`,
			"work_dir":    `"` + tmpDir + `"`,
			"cleanup":     "true",
		}

		err := createExtension("cleanup-test", "write",
			"{{executable}} write {{1}} "+outputFile, config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("cleanup-test", "write", "World")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}

		// File should be deleted after execution
		if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
			t.Error("Expected output file to be cleaned up")
		}

		// Test with cleanup disabled
		config["cleanup"] = "false"
		err = createExtension("no-cleanup-test", "write",
			"{{executable}} write {{1}} "+outputFile, config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("no-cleanup-test", "write", "World")
		if err != nil {
			t.Errorf("Failed to execute: %v", err)
		}

		// File should remain after execution
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Error("Expected output file to remain")
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "error-test.txt")
		config := map[string]interface{}{
			"output_file": `"error-test.txt"`,
			"work_dir":    `"` + tmpDir + `"`,
			"cleanup":     "true",
		}

		// Test command error
		err := createExtension("error-test", "error",
			"{{executable}} error {{1}} "+outputFile, config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("error-test", "error", "World")
		if err == nil {
			t.Error("Expected error from failing command, got nil")
		}

		// Test invalid work_dir
		config["work_dir"] = `"/nonexistent/directory"`
		err = createExtension("invalid-dir-test", "write",
			"{{executable}} write {{1}} output.txt", config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("invalid-dir-test", "write", "World")
		if err == nil {
			t.Error("Expected error from invalid work_dir, got nil")
		}
	})

	// Test with missing output_file
	t.Run("MissingOutputFile", func(t *testing.T) {
		config := map[string]interface{}{
			"work_dir": `"` + tmpDir + `"`,
			"cleanup":  "true",
		}

		err := createExtension("missing-output-test", "write",
			"{{executable}} write {{1}} output.txt", config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("missing-output-test", "write", "World")
		if err == nil {
			t.Error("Expected error from missing output_file, got nil")
		}
	})
}
