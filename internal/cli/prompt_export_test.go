package cli

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	langtool "github.com/danielmiessler/fabric/internal/tools/lang"
)

func newPromptExportRegistry(t *testing.T) *core.PluginRegistry {
	t.Helper()

	db := fsdb.NewDb(t.TempDir())
	if err := os.WriteFile(db.EnvFilePath, []byte{}, 0o644); err != nil {
		t.Fatalf("failed to create env file: %v", err)
	}
	if err := db.Configure(); err != nil {
		t.Fatalf("failed to configure db: %v", err)
	}

	return &core.PluginRegistry{
		Db:       db,
		Language: langtool.NewLanguage(),
	}
}

func createTestPattern(t *testing.T, db *fsdb.Db, name, content string) {
	t.Helper()

	patternDir := filepath.Join(db.Patterns.Dir, name)
	if err := os.MkdirAll(patternDir, 0o755); err != nil {
		t.Fatalf("failed to create pattern directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(patternDir, db.Patterns.SystemPatternFile), []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write pattern: %v", err)
	}
}

func createTestContext(t *testing.T, db *fsdb.Db, name, content string) {
	t.Helper()

	if err := os.WriteFile(filepath.Join(db.Contexts.Dir, name), []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write context: %v", err)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = w

	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close stdout writer: %v", err)
	}
	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	return string(output)
}

func TestRenderPromptExportPattern(t *testing.T) {
	registry := newPromptExportRegistry(t)
	createTestPattern(t, registry.Db, "test-pattern", "PATTERN\n{{input}}")

	flags := &Flags{
		PrintPrompt: true,
		Pattern:     "test-pattern",
		Message:     "user input",
	}

	got, err := renderPromptExport(flags, registry, "meta", "")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "System:\nPATTERN\nuser input\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportRawPattern(t *testing.T) {
	registry := newPromptExportRegistry(t)
	createTestPattern(t, registry.Db, "test-pattern", "PATTERN\n{{input}}")

	flags := &Flags{
		PrintPrompt: true,
		Pattern:     "test-pattern",
		Message:     "user input",
		Raw:         true,
	}

	got, err := renderPromptExport(flags, registry, "meta", "")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "User:\nPATTERN\nuser input\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportContextAndUser(t *testing.T) {
	registry := newPromptExportRegistry(t)
	createTestContext(t, registry.Db, "test-context", "CONTEXT")

	flags := &Flags{
		PrintPrompt: true,
		Context:     "test-context",
		Message:     "user input",
	}

	got, err := renderPromptExport(flags, registry, "meta", "")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "System:\nCONTEXT\n\nUser:\nuser input\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportNamedSessionHistory(t *testing.T) {
	registry := newPromptExportRegistry(t)

	session := &fsdb.Session{
		Name: "existing",
		Messages: []*chat.ChatCompletionMessage{
			{Role: chat.ChatMessageRoleUser, Content: "earlier"},
			{Role: chat.ChatMessageRoleAssistant, Content: "reply"},
		},
	}
	if err := registry.Db.Sessions.SaveSession(session); err != nil {
		t.Fatalf("failed to save session: %v", err)
	}

	flags := &Flags{
		PrintPrompt: true,
		Session:     "existing",
		Message:     "next",
	}

	got, err := renderPromptExport(flags, registry, "meta", "")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "User:\nearlier\n\nAssistant:\nreply\n\nUser:\nnext\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportMissingSessionDoesNotPrint(t *testing.T) {
	registry := newPromptExportRegistry(t)

	flags := &Flags{
		PrintPrompt: true,
		Session:     "missing",
		Message:     "hello",
	}

	var (
		got string
		err error
	)
	output := captureStdout(t, func() {
		got, err = renderPromptExport(flags, registry, "meta", "")
	})
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}
	if output != "" {
		t.Fatalf("expected no stdout output, got %q", output)
	}

	want := "User:\nhello\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportConsumesToolInput(t *testing.T) {
	registry := newPromptExportRegistry(t)
	createTestPattern(t, registry.Db, "test-pattern", "PATTERN\n{{input}}")

	flags := &Flags{
		PrintPrompt: true,
		Pattern:     "test-pattern",
	}

	got, err := renderPromptExport(flags, registry, "meta", "tool data")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "System:\nPATTERN\ntool data\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestRenderPromptExportWithoutVendorOrModel(t *testing.T) {
	registry := newPromptExportRegistry(t)
	flags := &Flags{
		PrintPrompt: true,
		Message:     "hello",
	}

	got, err := renderPromptExport(flags, registry, "meta", "")
	if err != nil {
		t.Fatalf("renderPromptExport() error = %v", err)
	}

	want := "User:\nhello\n\n"
	if got != want {
		t.Fatalf("renderPromptExport() = %q, want %q", got, want)
	}
}

func TestOutputPromptExport(t *testing.T) {
	prompt := "System:\nhello\n\n"

	t.Run("stdout only", func(t *testing.T) {
		var stdout bytes.Buffer
		copied := ""

		err := outputPromptExport(prompt, "", true, &stdout, func(content string) error {
			copied = content
			return nil
		})
		if err != nil {
			t.Fatalf("outputPromptExport() error = %v", err)
		}
		if stdout.String() != prompt {
			t.Fatalf("stdout = %q, want %q", stdout.String(), prompt)
		}
		if copied != prompt {
			t.Fatalf("copied content = %q, want %q", copied, prompt)
		}
	})

	t.Run("file only", func(t *testing.T) {
		var stdout bytes.Buffer
		outputPath := filepath.Join(t.TempDir(), "prompt.txt")

		err := outputPromptExport(prompt, outputPath, false, &stdout, func(string) error {
			t.Fatal("copy should not be called")
			return nil
		})
		if err != nil {
			t.Fatalf("outputPromptExport() error = %v", err)
		}
		if stdout.Len() != 0 {
			t.Fatalf("expected no stdout output, got %q", stdout.String())
		}

		content, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatalf("failed to read prompt file: %v", err)
		}
		if string(content) != prompt {
			t.Fatalf("file content = %q, want %q", string(content), prompt)
		}

		err = outputPromptExport(prompt, outputPath, false, &stdout, func(string) error { return nil })
		if err == nil {
			t.Fatal("expected error when output file already exists")
		}
	})
}

func TestValidatePromptExportFlags(t *testing.T) {
	tests := []struct {
		name  string
		flags *Flags
	}{
		{
			name: "dry run incompatible",
			flags: &Flags{
				PrintPrompt: true,
				DryRun:      true,
			},
		},
		{
			name: "output session incompatible",
			flags: &Flags{
				PrintPrompt:   true,
				OutputSession: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePromptExportFlags(tt.flags); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}
