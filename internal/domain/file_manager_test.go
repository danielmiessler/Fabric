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
	// Create a temporary directory for testing
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "file-manager-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	// Test file changes
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

	// Apply the changes
	if err := ApplyFileChanges(tempDir, changes); err != nil {
		t.Fatalf("ApplyFileChanges() error = %v", err)
	}

	// Verify the first file was created correctly
	content, err := os.ReadFile(filepath.Join(tempDir, "test.txt"))
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}
	if string(content) != "Hello, World!" {
		t.Errorf("File content = %q, want %q", string(content), "Hello, World!")
	}

	// Verify the nested file was created correctly
	content, err = os.ReadFile(filepath.Join(tempDir, "subdir/nested.txt"))
	if err != nil {
		t.Fatalf("Failed to read created nested file: %v", err)
	}
	if string(content) != "Nested content" {
		t.Errorf("Nested file content = %q, want %q", string(content), "Nested content")
	}

	// Test updating a file
	updateChanges := []FileChange{
		{
			Operation: "update",
			Path:      "test.txt",
			Content:   "Updated content",
		},
	}

	// Apply the update
	if err := ApplyFileChanges(tempDir, updateChanges); err != nil {
		t.Fatalf("ApplyFileChanges() error = %v", err)
	}
	// Verify the file was updated correctly
	content, err = os.ReadFile(filepath.Join(tempDir, "test.txt"))
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}
	if string(content) != "Updated content" {
		t.Errorf("Updated file content = %q, want %q", string(content), "Updated content")
	}
}
