package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFilePlugin(t *testing.T) {
	plugin := &FilePlugin{}

	// Create temp test files
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	content := "line1\nline2\nline3\nline4\nline5\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	bigFile := filepath.Join(tmpDir, "big.txt")
	err = os.WriteFile(bigFile, []byte(strings.Repeat("x", MaxFileSize+1)), 0644)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		operation   string
		value       string
		want        string
		wantErr     bool
		errContains string
		validate    func(string) bool
	}{
		{
			name:      "read file",
			operation: "read",
			value:     testFile,
			want:      content,
		},
		{
			name:      "tail file",
			operation: "tail",
			value:     testFile + "|3",
			want:      "line3\nline4\nline5",
		},
		{
			name:      "exists true",
			operation: "exists",
			value:     testFile,
			want:      "true",
		},
		{
			name:      "exists false",
			operation: "exists",
			value:     filepath.Join(tmpDir, "nonexistent.txt"),
			want:      "false",
		},
		{
			name:      "size",
			operation: "size",
			value:     testFile,
			want:      "30",
		},
		{
			name:      "modified",
			operation: "modified",
			value:     testFile,
			validate: func(got string) bool {
				_, err := time.Parse(time.RFC3339, got)
				return err == nil
			},
		},
		// Error cases
		{
			name:        "read non-existent",
			operation:   "read",
			value:       filepath.Join(tmpDir, "nonexistent.txt"),
			wantErr:     true,
			errContains: "could not stat file",
		},
		{
			name:        "invalid operation",
			operation:   "invalid",
			value:       testFile,
			wantErr:     true,
			errContains: "unknown operation",
		},
		{
			name:        "path traversal attempt",
			operation:   "read",
			value:       "../../../etc/passwd",
			wantErr:     true,
			errContains: "cannot contain '..'",
		},
		{
			name:        "file too large",
			operation:   "read",
			value:       bigFile,
			wantErr:     true,
			errContains: "exceeds limit",
		},
		{
			name:        "invalid tail format",
			operation:   "tail",
			value:       testFile,
			wantErr:     true,
			errContains: "requires format path|lines",
		},
		{
			name:        "invalid tail count",
			operation:   "tail",
			value:       testFile + "|invalid",
			wantErr:     true,
			errContains: "invalid line count",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := plugin.Apply(tt.operation, tt.value)

			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("FilePlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errContains != "" {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error %q should contain %q", err.Error(), tt.errContains)
				}
				return
			}

			// Check success cases
			if err == nil {
				if tt.validate != nil {
					if !tt.validate(got) {
						t.Errorf("FilePlugin.Apply() returned invalid result: %q", got)
					}
				} else if tt.want != "" && got != tt.want {
					t.Errorf("FilePlugin.Apply() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
