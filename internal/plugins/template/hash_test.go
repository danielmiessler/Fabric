// template/hash_test.go
package template

import (
	"os"
	"path/filepath"
	"testing"
)

func TestComputeHash(t *testing.T) {
	// Create a temporary test file
	content := []byte("test content for hashing")
	tmpfile, err := os.CreateTemp("", "hashtest")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	tests := []struct {
		name    string
		path    string
		want    string // known hash for test content
		wantErr bool
	}{
		{
			name:    "valid file",
			path:    tmpfile.Name(),
			want:    "e25dd806d495b413931f4eea50b677a7a5c02d00460924661283f211a37f7e7f", // pre-computed hash of "test content for hashing"
			wantErr: false,
		},
		{
			name:    "nonexistent file",
			path:    filepath.Join(os.TempDir(), "nonexistent"),
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComputeHash(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.wantErr {
				t.Errorf("ComputeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeStringHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:  "simple string",
			input: "test",
			want:  "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			name:  "longer string with spaces",
			input: "this is a test string",
			want:  "f6774519d1c7a3389ef327e9c04766b999db8cdfb85d1346c471ee86d65885bc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeStringHash(tt.input); got != tt.want {
				t.Errorf("ComputeStringHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestHashConsistency ensures both hash functions produce same results for same content
func TestHashConsistency(t *testing.T) {
	content := "test content for consistency check"

	// Create a file with the test content
	tmpfile, err := os.CreateTemp("", "hashconsistency")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if err := os.WriteFile(tmpfile.Name(), []byte(content), 0644); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Get hashes using both methods
	fileHash, err := ComputeHash(tmpfile.Name())
	if err != nil {
		t.Fatalf("ComputeHash failed: %v", err)
	}

	stringHash := ComputeStringHash(content)

	// Compare results
	if fileHash != stringHash {
		t.Errorf("Hash inconsistency: file hash %v != string hash %v", fileHash, stringHash)
	}
}
