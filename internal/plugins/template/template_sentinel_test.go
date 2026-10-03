package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// withTestExtension creates a temporary test extension and runs the test function
func withTestExtension(t *testing.T, name string, scriptContent string, testFunc func(*ExtensionManager, string)) {
	t.Helper()

	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "fabric")
	extensionsDir := filepath.Join(configDir, "extensions")
	binDir := filepath.Join(extensionsDir, "bin")
	configsDir := filepath.Join(extensionsDir, "configs")

	err := os.MkdirAll(binDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create bin directory: %v", err)
	}
	err = os.MkdirAll(configsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create configs directory: %v", err)
	}

	scriptPath := filepath.Join(binDir, name+".sh")
	err = os.WriteFile(scriptPath, []byte(scriptContent), 0755)
	if err != nil {
		t.Fatalf("Failed to create test script: %v", err)
	}

	configPath := filepath.Join(configsDir, name+".yaml")
	configContent := fmt.Sprintf(`name: %s
executable: %s
type: executable
timeout: "5s"
description: "Test extension"
version: "1.0.0"

operations:
  echo:
    cmd_template: "{{executable}} {{value}}"

config:
  output:
    method: stdout
`, name, scriptPath)
	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create extension config: %v", err)
	}

	mgr := NewExtensionManager(configDir)

	err = mgr.RegisterExtension(configPath)
	if err != nil {
		t.Fatalf("Failed to register extension: %v", err)
	}

	testFunc(mgr, name)
}

// TestSentinelTokenReplacement tests the fix for the {{input}} sentinel token bug
// This test verifies that when {{input}} is used inside an extension call,
// the actual input is passed to the extension, not the sentinel token.
func TestSentinelTokenReplacement(t *testing.T) {
	scriptContent := `#!/bin/bash
echo "RECEIVED: $@"
`

	withTestExtension(t, "echo-test", scriptContent, func(mgr *ExtensionManager, name string) {
		oldManager := extensionManager
		defer func() { extensionManager = oldManager }()
		extensionManager = mgr

		tests := []struct {
			name           string
			template       string
			input          string
			wantContain    string
			wantNotContain string
		}{
			{
				name:           "sentinel token with {{input}} in extension value",
				template:       "{{ext:echo-test:echo:__FABRIC_INPUT_SENTINEL_TOKEN__}}",
				input:          "test input data",
				wantContain:    "RECEIVED: test input data",
				wantNotContain: "__FABRIC_INPUT_SENTINEL_TOKEN__",
			},
			{
				name:           "direct input variable replacement",
				template:       "{{ext:echo-test:echo:{{input}}}}",
				input:          "Hello World",
				wantContain:    "RECEIVED: Hello World",
				wantNotContain: "{{input}}",
			},
			{
				name:           "sentinel with complex input",
				template:       "Result: {{ext:echo-test:echo:__FABRIC_INPUT_SENTINEL_TOKEN__}}",
				input:          "What is AI?",
				wantContain:    "RECEIVED: What is AI?",
				wantNotContain: "__FABRIC_INPUT_SENTINEL_TOKEN__",
			},
			{
				name:           "multiple words in input",
				template:       "{{ext:echo-test:echo:{{input}}}}",
				input:          "Multiple word input string",
				wantContain:    "RECEIVED: Multiple word input string",
				wantNotContain: "{{input}}",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ApplyTemplate(tt.template, map[string]string{}, tt.input)
				if err != nil {
					t.Errorf("ApplyTemplate() error = %v", err)
					return
				}

				if !strings.Contains(got, tt.wantContain) {
					t.Errorf("ApplyTemplate() = %q, should contain %q", got, tt.wantContain)
				}

				if strings.Contains(got, tt.wantNotContain) {
					t.Errorf("ApplyTemplate() = %q, should NOT contain %q", got, tt.wantNotContain)
				}
			})
		}
	})
}

// TestSentinelInVariableProcessing tests that the sentinel token is handled
// correctly in regular variable processing (not just extensions)
// Note: The sentinel is only replaced when it appears in extension values,
// not when used as a standalone variable (which would be a user error)
func TestSentinelInVariableProcessing(t *testing.T) {
	tests := []struct {
		name     string
		template string
		vars     map[string]string
		input    string
		want     string
	}{
		{
			name:     "input variable works normally",
			template: "Value: {{input}}",
			input:    "actual input",
			want:     "Value: actual input",
		},
		{
			name:     "multiple input references",
			template: "First: {{input}}, Second: {{input}}",
			input:    "test",
			want:     "First: test, Second: test",
		},
		{
			name:     "input with variables",
			template: "Var: {{name}}, Input: {{input}}",
			vars:     map[string]string{"name": "TestVar"},
			input:    "input value",
			want:     "Var: TestVar, Input: input value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyTemplate(tt.template, tt.vars, tt.input)
			if err != nil {
				t.Errorf("ApplyTemplate() error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("ApplyTemplate() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestExtensionValueWithSentinel specifically tests the extension value
// sentinel replacement logic
func TestExtensionValueWithSentinel(t *testing.T) {
	scriptContent := `#!/bin/bash
# Output each argument on a separate line
for arg in "$@"; do
    echo "ARG: $arg"
done
`

	withTestExtension(t, "arg-test", scriptContent, func(mgr *ExtensionManager, name string) {
		oldManager := extensionManager
		defer func() { extensionManager = oldManager }()
		extensionManager = mgr

		template := "{{ext:arg-test:echo:prefix-__FABRIC_INPUT_SENTINEL_TOKEN__-suffix}}"
		input := "MYINPUT"

		got, err := ApplyTemplate(template, map[string]string{}, input)
		if err != nil {
			t.Fatalf("ApplyTemplate() error = %v", err)
		}

		expectedContain := "ARG: prefix-MYINPUT-suffix"
		if !strings.Contains(got, expectedContain) {
			t.Errorf("ApplyTemplate() = %q, should contain %q", got, expectedContain)
		}

		if strings.Contains(got, "__FABRIC_INPUT_SENTINEL_TOKEN__") {
			t.Errorf("ApplyTemplate() = %q, should NOT contain sentinel token", got)
		}
	})
}

// TestNestedInputInExtension tests the original bug case:
// {{ext:name:op:{{input}}}} should pass the actual input, not the sentinel
func TestNestedInputInExtension(t *testing.T) {
	scriptContent := `#!/bin/bash
echo "NESTED_TEST: $*"
`

	withTestExtension(t, "nested-test", scriptContent, func(mgr *ExtensionManager, name string) {
		oldManager := extensionManager
		defer func() { extensionManager = oldManager }()
		extensionManager = mgr

		// The inner {{input}} resolves first. The extension then runs with the expanded value.
		template := "{{ext:nested-test:echo:{{input}}}}"
		input := "What is Artificial Intelligence"

		got, err := ApplyTemplate(template, map[string]string{}, input)
		if err != nil {
			t.Fatalf("ApplyTemplate() error = %v", err)
		}

		expectedContain := "NESTED_TEST: What is Artificial Intelligence"
		if !strings.Contains(got, expectedContain) {
			t.Errorf("ApplyTemplate() = %q, should contain %q", got, expectedContain)
		}

		if strings.Contains(got, "__FABRIC_INPUT_SENTINEL_TOKEN__") {
			t.Errorf("ApplyTemplate() output contains sentinel token (BUG NOT FIXED): %q", got)
		}

		if strings.Contains(got, "{{input}}") {
			t.Errorf("ApplyTemplate() output contains unresolved {{input}}: %q", got)
		}
	})
}
