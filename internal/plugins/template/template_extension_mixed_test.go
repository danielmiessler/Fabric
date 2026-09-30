package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestExtensionValueMixedInputAndVariable ensures an extension value mixing {{input}} and another template variable is processed.
func TestExtensionValueMixedInputAndVariable(t *testing.T) {
	input := "PRIMARY"
	variables := map[string]string{
		"suffix": "SUF",
	}

	tmp := t.TempDir()
	configDir := filepath.Join(tmp, ".config", "fabric")
	extsDir := filepath.Join(configDir, "extensions")
	binDir := filepath.Join(extsDir, "bin")
	configsDir := filepath.Join(extsDir, "configs")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatalf("mkdir bin: %v", err)
	}
	if err := os.MkdirAll(configsDir, 0o755); err != nil {
		t.Fatalf("mkdir configs: %v", err)
	}

	scriptPath := filepath.Join(binDir, "mix-echo.sh")
	script := "#!/bin/sh\necho VAL=$1\n"
	if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}

	configYAML := "" +
		"name: mix-echo\n" +
		"type: executable\n" +
		"executable: " + scriptPath + "\n" +
		"description: mixed input/variable test\n" +
		"version: 1.0.0\n" +
		"timeout: 5s\n" +
		"operations:\n" +
		"  echo:\n" +
		"    cmd_template: '{{executable}} {{value}}'\n"
	if err := os.WriteFile(filepath.Join(configsDir, "mix-echo.yaml"), []byte(configYAML), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	mgr := NewExtensionManager(configDir)
	if err := mgr.RegisterExtension(filepath.Join(configsDir, "mix-echo.yaml")); err != nil {
		if strings.Contains(err.Error(), "operation not permitted") {
			t.Skipf("skipping due to exec restriction: %v", err)
		}
		t.Fatalf("register: %v", err)
	}

	prevMgr := extensionManager
	extensionManager = mgr
	defer func() { extensionManager = prevMgr }()

	tmpl := "{{ext:mix-echo:echo:pre-{{input}}-mid-{{suffix}}-post}}"

	out, err := ApplyTemplate(tmpl, variables, input)
	if err != nil {
		t.Fatalf("ApplyTemplate error: %v", err)
	}

	if !strings.Contains(out, "VAL=pre-PRIMARY-mid-SUF-post") {
		t.Fatalf("unexpected output: %q", out)
	}
}
