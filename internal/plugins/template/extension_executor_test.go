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

	testScript := filepath.Join(tmpDir, "test-script.sh")
	scriptContent := `#!/bin/bash
case "$1" in
    "stdout")
        echo "Hello, $2!"
        ;;
    "echo")
        echo "$2"
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

	registry := NewExtensionRegistry(tmpDir)
	executor := NewExtensionExecutor(registry)

	// Shell metacharacters in the value must stay literal. Without escaping,
	// "sh -c" runs a value such as "; touch /tmp/pwned" as a second command.
	t.Run("ShellInjectionBlocked", func(t *testing.T) {
		// Use a marker file to detect if injection succeeded.
		markerFile := filepath.Join(tmpDir, "injection-marker")
		_ = os.Remove(markerFile)

		configPath := filepath.Join(tmpDir, "inject-test.yaml")
		configContent := `name: inject-test
executable: ` + testScript + `
type: executable
timeout: 5s
operations:
  echo:
    cmd_template: "{{executable}} echo {{value}}"
config:
  output:
    method: stdout`

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to create config: %v", err)
		}

		if err := registry.Register(configPath); err != nil {
			t.Fatalf("Failed to register extension: %v", err)
		}

		maliciousValue := "hello; touch " + markerFile

		output, err := executor.Execute("inject-test", "echo", maliciousValue)
		if err != nil {
			t.Fatalf("Execute failed: %v", err)
		}

		// A literal semicolon in the output shows the shell did not interpret it.
		if !strings.Contains(output, "hello; touch") {
			t.Errorf("Expected literal value in output, got: %q", output)
		}

		if _, err := os.Stat(markerFile); !os.IsNotExist(err) {
			t.Error("SECURITY: command injection succeeded — marker file was created")
		}
	})

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

	t.Run("ExecutionErrors", func(t *testing.T) {
		_, err := executor.Execute("nonexistent", "test", "value")
		if err == nil {
			t.Error("Expected error executing non-existent extension, got nil")
		}

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

	// createExtension writes an extension config file and registers it.
	createExtension := func(name, opName, cmdTemplate string, config map[string]any) error {
		configPath := filepath.Join(tmpDir, name+".yaml")
		var configContent strings.Builder
		configContent.WriteString(`name: ` + name + `
executable: ` + testScript + `
type: executable
timeout: 30s
operations:
  ` + opName + `:
    cmd_template: "` + cmdTemplate + `"
config:
  output:
    method: file
    file_config:`)

		for k, v := range config {
			configContent.WriteString("\n      " + k + ": " + strings.TrimSpace(v.(string)))
		}

		if err := os.WriteFile(configPath, []byte(configContent.String()), 0644); err != nil {
			return err
		}

		return registry.Register(configPath)
	}

	t.Run("BasicFixedFile", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "output.txt")
		config := map[string]any{
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

	t.Run("NoWorkDir", func(t *testing.T) {
		config := map[string]any{
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

	t.Run("CleanupBehavior", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "cleanup-test.txt")

		config := map[string]any{
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

		if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
			t.Error("Expected output file to be cleaned up")
		}

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

		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Error("Expected output file to remain")
		}
	})

	t.Run("ErrorCases", func(t *testing.T) {
		outputFile := filepath.Join(tmpDir, "error-test.txt")
		config := map[string]any{
			"output_file": `"error-test.txt"`,
			"work_dir":    `"` + tmpDir + `"`,
			"cleanup":     "true",
		}

		err := createExtension("error-test", "error",
			"{{executable}} error {{1}} "+outputFile, config)
		if err != nil {
			t.Fatalf("Failed to create extension: %v", err)
		}

		_, err = executor.Execute("error-test", "error", "World")
		if err == nil {
			t.Error("Expected error from failing command, got nil")
		}

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

	t.Run("MissingOutputFile", func(t *testing.T) {
		config := map[string]any{
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
