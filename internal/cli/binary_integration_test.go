//go:build integration

package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func requireClaudeIntegration(t *testing.T) []string {
	t.Helper()

	if os.Getenv("FABRIC_RUN_CLAUDE_INTEGRATION") != "1" {
		t.Skip("set FABRIC_RUN_CLAUDE_INTEGRATION=1 to run real ClaudeCode integration tests")
	}

	claudeBinary := os.Getenv("CLAUDECODE_BINARY_PATH")
	if claudeBinary == "" {
		var err error
		claudeBinary, err = exec.LookPath("claude")
		if err != nil {
			t.Skip("claude binary not found in PATH and CLAUDECODE_BINARY_PATH is not set")
		}
	}

	originalHome, err := os.UserHomeDir()
	if err != nil || originalHome == "" {
		t.Skip("unable to determine the real user home directory for Claude auth")
	}

	wrapperPath := filepath.Join(t.TempDir(), "claude-wrapper.sh")
	wrapper := fmt.Sprintf("#!/bin/sh\nHOME=%q USERPROFILE=%q exec %q \"$@\"\n", originalHome, originalHome, claudeBinary)
	if err := os.WriteFile(wrapperPath, []byte(wrapper), 0o700); err != nil {
		t.Fatalf("write Claude wrapper: %v", err)
	}
	if err := os.Chmod(wrapperPath, 0o700); err != nil {
		t.Fatalf("chmod Claude wrapper: %v", err)
	}

	return []string{"CLAUDECODE_BINARY_PATH=" + wrapperPath}
}

func TestIntegrationClaudeCodePatternExecution(t *testing.T) {
	env := requireClaudeIntegration(t)
	h := newBinaryHarness(t, "summarize")
	input := "Summarize this sentence in one short line: Fabric ships a compiled CLI."

	result := h.runFabric(t, input, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--pattern", "summarize")
	if result.exitCode != 0 {
		t.Fatalf("expected ClaudeCode execution to succeed, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if strings.TrimSpace(result.stdout) == "" {
		t.Fatal("expected ClaudeCode output to be non-empty")
	}
	if strings.Contains(result.stdout, "Dry run:") {
		t.Fatalf("expected real Claude output, got dry-run response: %s", result.stdout)
	}
}

func TestIntegrationClaudeCodeStreaming(t *testing.T) {
	env := requireClaudeIntegration(t)
	h := newBinaryHarness(t)
	input := "Reply with exactly the text: streaming-ok"

	result := h.runFabric(t, input, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--stream")
	if result.exitCode != 0 {
		t.Fatalf("expected ClaudeCode stream to succeed, got %d\nstderr: %s", result.exitCode, result.stderr)
	}
	if !strings.Contains(strings.ToLower(result.stdout), "streaming-ok") {
		t.Fatalf("expected streaming output to include target text, got: %s", result.stdout)
	}
}

func TestIntegrationWriteLatexToPDFPipeline(t *testing.T) {
	env := requireClaudeIntegration(t)
	if _, err := exec.LookPath("pdflatex"); err != nil {
		t.Skip("pdflatex is required for the write_latex -> to_pdf integration test")
	}

	h := newBinaryHarness(t, "write_latex")
	toPDFBinary := buildGoBinary(t, "cmd/to_pdf")

	latex := h.runFabric(
		t,
		"Write a minimal LaTeX document with a title and one short paragraph about Fabric.",
		"",
		env,
		"-V", "ClaudeCode",
		"-m", "claude-sonnet-4-6",
		"--pattern", "write_latex",
	)
	if latex.exitCode != 0 {
		t.Fatalf("expected write_latex command to succeed, got %d\nstderr: %s", latex.exitCode, latex.stderr)
	}
	if strings.TrimSpace(latex.stdout) == "" {
		t.Fatal("expected LaTeX output to be non-empty")
	}

	pdfPath := filepath.Join(t.TempDir(), "fabric.pdf")
	pdf := runCommand(t, toPDFBinary, latex.stdout, h.repoRoot, h.baseEnv(), pdfPath)
	if pdf.exitCode != 0 {
		t.Fatalf("expected to_pdf to succeed, got %d\nstderr: %s", pdf.exitCode, pdf.stderr)
	}
	if _, err := os.Stat(pdfPath); err != nil {
		t.Fatalf("expected PDF output at %s: %v", pdfPath, err)
	}
}

func TestIntegrationFabricPipelineCleanTextToSummarize(t *testing.T) {
	env := requireClaudeIntegration(t)
	h := newBinaryHarness(t, "clean_text", "summarize")

	input := "  Fabric   is   a   CLI   tool.\n\nIt helps with AI workflows.  "
	clean := h.runFabric(t, input, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--pattern", "clean_text")
	if clean.exitCode != 0 {
		t.Fatalf("expected clean_text command to succeed, got %d\nstderr: %s", clean.exitCode, clean.stderr)
	}
	if strings.TrimSpace(clean.stdout) == "" {
		t.Fatal("expected clean_text output to be non-empty")
	}

	summary := h.runFabric(t, clean.stdout, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--pattern", "summarize")
	if summary.exitCode != 0 {
		t.Fatalf("expected summarize command to succeed, got %d\nstderr: %s", summary.exitCode, summary.stderr)
	}
	if strings.TrimSpace(summary.stdout) == "" {
		t.Fatal("expected summarize output to be non-empty")
	}
	for _, marker := range []string{"ONE SENTENCE SUMMARY", "MAIN POINTS", "TAKEAWAYS"} {
		if !strings.Contains(summary.stdout, marker) {
			t.Fatalf("expected summarize output marker %q, got: %s", marker, summary.stdout)
		}
	}
}

func TestIntegrationFabricPipelineReviewCodeToSummarize(t *testing.T) {
	env := requireClaudeIntegration(t)
	h := newBinaryHarness(t, "review_code", "summarize")

	code := strings.Join([]string{
		"package main",
		"",
		"import \"fmt\"",
		"",
		"func main() {",
		"    values := []int{1, 2, 3}",
		"    for i := 0; i <= len(values); i++ {",
		"        fmt.Println(values[i])",
		"    }",
		"}",
	}, "\n")

	review := h.runFabric(t, code, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--pattern", "review_code")
	if review.exitCode != 0 {
		t.Fatalf("expected review_code command to succeed, got %d\nstderr: %s", review.exitCode, review.stderr)
	}
	if strings.TrimSpace(review.stdout) == "" {
		t.Fatal("expected review_code output to be non-empty")
	}

	summary := h.runFabric(t, review.stdout, "", env, "-V", "ClaudeCode", "-m", "claude-sonnet-4-6", "--pattern", "summarize")
	if summary.exitCode != 0 {
		t.Fatalf("expected summarize command to succeed, got %d\nstderr: %s", summary.exitCode, summary.stderr)
	}
	if strings.TrimSpace(summary.stdout) == "" {
		t.Fatal("expected summarize output to be non-empty")
	}
	for _, marker := range []string{"ONE SENTENCE SUMMARY", "MAIN POINTS", "TAKEAWAYS"} {
		if !strings.Contains(summary.stdout, marker) {
			t.Fatalf("expected summarize output marker %q, got: %s", marker, summary.stdout)
		}
	}
}

func TestIntegrationYouTubeSummarizeToExtractWisdomPipeline(t *testing.T) {
	env := requireClaudeIntegration(t)
	if os.Getenv("FABRIC_RUN_NETWORK_INTEGRATION") != "1" {
		t.Skip("set FABRIC_RUN_NETWORK_INTEGRATION=1 to run network-backed YouTube pipeline integration tests")
	}
	if _, err := exec.LookPath("yt-dlp"); err != nil {
		t.Skip("yt-dlp is required for YouTube transcript integration tests")
	}

	h := newBinaryHarness(t, "summarize", "extract_wisdom")
	videoURL := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	summary := h.runFabric(
		t,
		"",
		"",
		env,
		"-V", "ClaudeCode",
		"-m", "claude-sonnet-4-6",
		"-y", videoURL,
		"--transcript",
		"--pattern", "summarize",
	)
	if summary.exitCode != 0 {
		t.Fatalf("expected YouTube summarize command to succeed, got %d\nstderr: %s", summary.exitCode, summary.stderr)
	}
	if strings.TrimSpace(summary.stdout) == "" {
		t.Fatal("expected summarized transcript output to be non-empty")
	}

	wisdom := h.runFabric(
		t,
		summary.stdout,
		"",
		env,
		"-V", "ClaudeCode",
		"-m", "claude-sonnet-4-6",
		"--pattern", "extract_wisdom",
	)
	if wisdom.exitCode != 0 {
		t.Fatalf("expected extract_wisdom command to succeed, got %d\nstderr: %s", wisdom.exitCode, wisdom.stderr)
	}
	if strings.TrimSpace(wisdom.stdout) == "" {
		t.Fatal("expected extract_wisdom output to be non-empty")
	}
	if strings.Contains(strings.ToLower(wisdom.stderr), "pattern") && strings.Contains(strings.ToLower(wisdom.stderr), "not found") {
		t.Fatalf("unexpected pattern failure in pipeline stderr: %s", wisdom.stderr)
	}
}
