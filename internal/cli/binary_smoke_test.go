package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFabricBinaryHelp(t *testing.T) {
	h := newBinaryHarness(t)
	result := h.runFabric(t, "", "", nil, "--help")

	if result.exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(result.stdout, "Usage:") {
		t.Fatalf("expected help output to include Usage, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "Application Options:") {
		t.Fatalf("expected help output to include Application Options, got: %s", result.stdout)
	}
}

func TestFabricBinaryVersion(t *testing.T) {
	h := newBinaryHarness(t)
	result := h.runFabric(t, "", "", nil, "--version")

	if result.exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	versionOutput := strings.TrimSpace(result.stdout)
	if versionOutput == "" {
		t.Fatal("expected non-empty version output")
	}
	if !strings.HasPrefix(versionOutput, "v") {
		t.Fatalf("expected version output to start with 'v', got %q", versionOutput)
	}
}

func TestFabricBinaryListVendors(t *testing.T) {
	h := newBinaryHarness(t)
	result := h.runFabric(t, "", "", nil, "--listvendors")

	if result.exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(result.stdout, "Available Vendors:") {
		t.Fatalf("expected vendors header, got: %s", result.stdout)
	}
	for _, vendor := range []string{"ClaudeCode", "OpenAI", "LM Studio"} {
		if !strings.Contains(result.stdout, vendor) {
			t.Fatalf("expected vendor %q in output: %s", vendor, result.stdout)
		}
	}
}

func TestFabricBinaryListPatterns(t *testing.T) {
	h := newBinaryHarness(t, "summarize", "create_coding_feature")
	result := h.runFabric(t, "", "", nil, "--listpatterns")

	if result.exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	for _, patternName := range []string{"summarize", "create_coding_feature"} {
		if !strings.Contains(result.stdout, patternName) {
			t.Fatalf("expected pattern %q in output: %s", patternName, result.stdout)
		}
	}
}

func TestFabricBinaryDryRunPattern(t *testing.T) {
	h := newBinaryHarness(t, "summarize")
	input := "Summarize this short note."
	result := h.runFabric(t, input, "", nil, "--dry-run", "--pattern", "summarize")

	if result.exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(result.stdout, "Dry run: Would send the following request:") {
		t.Fatalf("expected dry-run request output, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "System:") {
		t.Fatalf("expected system prompt in dry-run output, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, input) {
		t.Fatalf("expected user input in dry-run output, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "Dry run: Fake response sent by DryRun plugin") {
		t.Fatalf("expected dry-run response marker, got: %s", result.stdout)
	}
}

func TestCode2ContextPipelineDryRun(t *testing.T) {
	h := newBinaryHarness(t, "create_coding_feature")
	code2contextBinary := buildGoBinary(t, "cmd/code2context")

	projectDir := filepath.Join(t.TempDir(), "project")
	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		t.Fatalf("create project dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(projectDir, "main.go"), []byte("package main\n\nfunc main() {}\n"), 0o644); err != nil {
		t.Fatalf("write sample file: %v", err)
	}

	scan := runCommand(t, code2contextBinary, "", h.repoRoot, h.baseEnv(), projectDir, "Create a Hello World C program")
	if scan.exitCode != 0 {
		t.Fatalf("expected code2context to succeed, got %d\nstderr: %s", scan.exitCode, scan.stderr)
	}
	if !strings.Contains(scan.stdout, "\"instructions\"") {
		t.Fatalf("expected code2context JSON output, got: %s", scan.stdout)
	}

	result := h.runFabric(t, scan.stdout, projectDir, nil, "--dry-run", "--pattern", "create_coding_feature")
	if result.exitCode != 0 {
		t.Fatalf("expected fabric dry-run to succeed, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(result.stdout, "Dry run: Would send the following request:") {
		t.Fatalf("expected dry-run output, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "\"type\": \"instructions\"") {
		t.Fatalf("expected code2context payload in request, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "Successfully applied file changes.") {
		t.Fatalf("expected create_coding_feature side-effect message, got: %s", result.stdout)
	}
}

func TestGenerateChangelogHelp(t *testing.T) {
	h := newBinaryHarness(t)
	changelogBinary := buildGoBinary(t, "cmd/generate_changelog")

	result := runCommand(t, changelogBinary, "", h.repoRoot, h.baseEnv(), "--help")
	if result.exitCode != 0 {
		t.Fatalf("expected generate_changelog help to succeed, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(result.stdout, "Usage:") {
		t.Fatalf("expected help usage output, got: %s", result.stdout)
	}
	if !strings.Contains(result.stdout, "--ai-summarize") {
		t.Fatalf("expected ai-summarize flag in help, got: %s", result.stdout)
	}
}
