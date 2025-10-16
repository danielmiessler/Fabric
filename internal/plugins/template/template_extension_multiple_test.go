package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMultipleExtensionsWithInput ensures multiple extension calls each using {{input}} get proper substitution.
func TestMultipleExtensionsWithInput(t *testing.T) {
	input := "DATA"
	variables := map[string]string{}

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

	scriptPath := filepath.Join(binDir, "multi-echo.sh")
	script := "#!/bin/sh\necho ECHO=$1\n"
	if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}

	configYAML := "" +
		"name: multi-echo\n" +
		"type: executable\n" +
		"executable: " + scriptPath + "\n" +
		"description: multi echo extension\n" +
		"version: 1.0.0\n" +
		"timeout: 5s\n" +
		"operations:\n" +
		"  echo:\n" +
		"    cmd_template: '{{executable}} {{value}}'\n"
	if err := os.WriteFile(filepath.Join(configsDir, "multi-echo.yaml"), []byte(configYAML), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	mgr := NewExtensionManager(configDir)
	if err := mgr.RegisterExtension(filepath.Join(configsDir, "multi-echo.yaml")); err != nil {
		t.Fatalf("register: %v", err)
	}
	prev := extensionManager
	extensionManager = mgr
	defer func() { extensionManager = prev }()

	tmpl := strings.Join([]string{
		"First: {{ext:multi-echo:echo:{{input}}}}",
		"Second: {{ext:multi-echo:echo:{{input}}}}",
		"Third: {{ext:multi-echo:echo:{{input}}}}",
	}, " | ")

	out, err := ApplyTemplate(tmpl, variables, input)
	if err != nil {
		t.Fatalf("ApplyTemplate error: %v", err)
	}

	wantCount := 3
	occ := strings.Count(out, "ECHO=DATA")
	if occ != wantCount {
		t.Fatalf("expected %d occurrences of ECHO=DATA, got %d; output=%q", wantCount, occ, out)
	}
}
