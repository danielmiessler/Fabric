package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileChanges(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int // number of expected file changes
		wantErr bool
	}{
		{
			name:    "No " + FileChangesMarker + " section",
			input:   "This is a normal response with no file changes.",
			want:    0,
			wantErr: false,
		},
		{
			name: "Valid " + FileChangesMarker + " section",
			input: `Some text before.
` + FileChangesMarker + `
[
	{
		"operation": "create",
		"path": "test.txt",
		"content": "Hello, World!"
	},
	{
		"operation": "update",
		"path": "other.txt",
		"content": "Updated content"
	}
]
Some text after.`,
			want:    2,
			wantErr: false,
		},
		{
			name: "Invalid JSON in " + FileChangesMarker + " section",
			input: `Some text before.
` + FileChangesMarker + `
[
	{
		"operation": "create",
		"path": "test.txt",
		"content": "Hello, World!"
	},
	{
		"operation": "invalid",
		"path": "other.txt"
		"content": "Updated content"
	}
]`,
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid operation",
			input: `Some text before.
` + FileChangesMarker + `
[
	{
		"operation": "delete",
		"path": "test.txt",
		"content": ""
	}
]`,
			want:    0,
			wantErr: true,
		},
		{
			name: "Empty path",
			input: `Some text before.
` + FileChangesMarker + `
[
	{
		"operation": "create",
		"path": "",
		"content": "Hello, World!"
	}
]`,
			want:    0,
			wantErr: true,
		},
		{
			name: "Suspicious path with directory traversal",
			input: `Some text before.
` + FileChangesMarker + `
[
	{
		"operation": "create",
		"path": "../etc/passwd",
		"content": "Hello, World!"
	}
]`,
			want:    0,
			wantErr: true,
		},
		{
			name: "Control character in path",
			input: FileChangesMarker + `
[
	{
		"operation": "create",
		"path": "a\u001b7b\u0007.txt",
		"content": "Hello, World!"
	}
]`,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, err := ParseFileChanges(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFileChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.want {
				t.Errorf("ParseFileChanges() got %d file changes, want %d", len(got), tt.want)
			}
		})
	}
}

func TestApplyFileChanges(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "file-manager-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	changes := []FileChange{
		{
			Operation: "create",
			Path:      "test.txt",
			Content:   "Hello, World!",
		},
		{
			Operation: "create",
			Path:      "subdir/nested.txt",
			Content:   "Nested content",
		},
	}

	if err := ApplyFileChanges(tempDir, changes); err != nil {
		t.Fatalf("ApplyFileChanges() error = %v", err)
	}

	content, err := os.ReadFile(filepath.Join(tempDir, "test.txt"))
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}
	if string(content) != "Hello, World!" {
		t.Errorf("File content = %q, want %q", string(content), "Hello, World!")
	}

	content, err = os.ReadFile(filepath.Join(tempDir, "subdir/nested.txt"))
	if err != nil {
		t.Fatalf("Failed to read created nested file: %v", err)
	}
	if string(content) != "Nested content" {
		t.Errorf("Nested file content = %q, want %q", string(content), "Nested content")
	}

	updateChanges := []FileChange{
		{
			Operation: "update",
			Path:      "test.txt",
			Content:   "Updated content",
		},
	}

	if err := ApplyFileChanges(tempDir, updateChanges); err != nil {
		t.Fatalf("ApplyFileChanges() error = %v", err)
	}
	content, err = os.ReadFile(filepath.Join(tempDir, "test.txt"))
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}
	if string(content) != "Updated content" {
		t.Errorf("Updated file content = %q, want %q", string(content), "Updated content")
	}

	for _, bad := range []string{"../escape.txt", "/etc/escape.txt", "subdir/../../escape.txt"} {
		if err := ApplyFileChanges(tempDir, []FileChange{{Operation: "create", Path: bad, Content: "x"}}); err == nil {
			t.Errorf("ApplyFileChanges(%q) = nil, want error", bad)
		}
	}
}

// TestApplyFileChanges_RejectsGitAndSymlinks makes sure that ApplyFileChanges
// does not write into .git or through a symlink.
func TestApplyFileChanges_RejectsGitAndSymlinks(t *testing.T) {
	projectRoot := t.TempDir()
	outside := t.TempDir()

	if err := ApplyFileChanges(projectRoot, []FileChange{
		{Operation: "create", Path: ".GIT/hooks/pre-commit", Content: "#!/bin/sh\n"},
	}); err == nil {
		t.Errorf("write into .git should be rejected")
	}
	if _, err := os.Stat(filepath.Join(projectRoot, ".GIT")); err == nil {
		t.Errorf(".git directory was made despite rejection")
	}

	if err := os.Symlink(outside, filepath.Join(projectRoot, "link")); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}
	if err := ApplyFileChanges(projectRoot, []FileChange{
		{Operation: "create", Path: "link/out.txt", Content: "out"},
	}); err == nil {
		t.Errorf("write through a symlink out of the root should be rejected")
	}
	if _, err := os.Stat(filepath.Join(outside, "out.txt")); err == nil {
		t.Errorf("file was written out of the project root")
	}

	// A symlink that stays in the root, for example hooks -> .git/hooks.
	if err := os.MkdirAll(filepath.Join(projectRoot, ".git", "hooks"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(".git", "hooks"), filepath.Join(projectRoot, "hooks")); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"hooks/pre-commit", "hooks"} {
		if err := ApplyFileChanges(projectRoot, []FileChange{
			{Operation: "create", Path: path, Content: "#!/bin/sh\n"},
		}); err == nil {
			t.Errorf("write to %q through an in-root symlink should be rejected", path)
		}
	}
	if _, err := os.Stat(filepath.Join(projectRoot, ".git", "hooks", "pre-commit")); err == nil {
		t.Errorf("file was written into .git through a symlink")
	}
}

// TestApplyFileChanges_NoWriteBeforeBadPath checks that a bad path stops the
// change set before the first write.
func TestApplyFileChanges_NoWriteBeforeBadPath(t *testing.T) {
	projectRoot := t.TempDir()
	if err := ApplyFileChanges(projectRoot, []FileChange{
		{Operation: "create", Path: "good.txt", Content: "good"},
		{Operation: "create", Path: ".git/config", Content: "bad"},
	}); err == nil {
		t.Fatal("ApplyFileChanges() accepted a path into .git")
	}
	if _, err := os.Stat(filepath.Join(projectRoot, "good.txt")); err == nil {
		t.Error("good.txt was written before the bad path was found")
	}
}
