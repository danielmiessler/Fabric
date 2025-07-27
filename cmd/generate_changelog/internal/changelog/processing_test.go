package changelog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/config"
)

func TestDetectVersion(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name              string
		versionNixContent string
		expectedVersion   string
		shouldError       bool
	}{
		{
			name:              "valid version.nix",
			versionNixContent: `"1.2.3"`,
			expectedVersion:   "1.2.3",
			shouldError:       false,
		},
		{
			name:              "version with extra whitespace",
			versionNixContent: `"1.2.3"   `,
			expectedVersion:   "1.2.3",
			shouldError:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create version.nix file
			versionNixPath := filepath.Join(tempDir, "version.nix")
			if err := os.WriteFile(versionNixPath, []byte(tt.versionNixContent), 0644); err != nil {
				t.Fatalf("Failed to create version.nix: %v", err)
			}

			cfg := &config.Config{
				RepoPath: tempDir,
			}

			g := &Generator{cfg: cfg}

			version, err := g.detectVersion()
			if tt.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if version != tt.expectedVersion {
				t.Errorf("Expected version '%s', got '%s'", tt.expectedVersion, version)
			}

			// Clean up
			os.Remove(versionNixPath)
		})
	}
}

func TestInsertVersionAtTop_ImprovedRobustness(t *testing.T) {
	tempDir := t.TempDir()
	changelogPath := filepath.Join(tempDir, "CHANGELOG.md")

	cfg := &config.Config{
		RepoPath: tempDir,
	}

	g := &Generator{cfg: cfg}

	tests := []struct {
		name            string
		existingContent string
		entry           string
		expectedContent string
	}{
		{
			name:            "header with trailing spaces",
			existingContent: "# Changelog   \n\n## v1.0.0\n- Old content",
			entry:           "## v2.0.0\n- New content",
			expectedContent: "# Changelog   \n\n## v2.0.0\n- New content\n## v1.0.0\n- Old content",
		},
		{
			name:            "header with different line endings",
			existingContent: "# Changelog\r\n\r\n## v1.0.0\r\n- Old content",
			entry:           "## v2.0.0\n- New content",
			expectedContent: "# Changelog\r\n\r\n## v2.0.0\n- New content\n## v1.0.0\r\n- Old content",
		},
		{
			name:            "no existing header",
			existingContent: "Some existing content without header",
			entry:           "## v1.0.0\n- New content",
			expectedContent: "# Changelog\n\n## v1.0.0\n- New content\n\nSome existing content without header",
		},
		{
			name:            "new file creation",
			existingContent: "",
			entry:           "## v1.0.0\n- Initial release",
			expectedContent: "# Changelog\n\n## v1.0.0\n- Initial release\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write existing content (or create empty file)
			if tt.existingContent != "" {
				if err := os.WriteFile(changelogPath, []byte(tt.existingContent), 0644); err != nil {
					t.Fatalf("Failed to write existing content: %v", err)
				}
			} else {
				// Remove file if it exists to test new file creation
				os.Remove(changelogPath)
			}

			// Insert new version
			if err := g.insertVersionAtTop(tt.entry); err != nil {
				t.Fatalf("insertVersionAtTop failed: %v", err)
			}

			// Read result
			result, err := os.ReadFile(changelogPath)
			if err != nil {
				t.Fatalf("Failed to read result: %v", err)
			}

			if string(result) != tt.expectedContent {
				t.Errorf("Expected:\n%q\nGot:\n%q", tt.expectedContent, string(result))
			}
		})
	}
}

func TestProcessIncomingPRs_FileAggregation(t *testing.T) {
	tempDir := t.TempDir()
	incomingDir := filepath.Join(tempDir, "incoming")

	// Create incoming directory and files
	if err := os.MkdirAll(incomingDir, 0755); err != nil {
		t.Fatalf("Failed to create incoming dir: %v", err)
	}

	// Create test incoming files
	file1Content := "## PR #1\n- Feature A"
	file2Content := "## PR #2\n- Feature B"

	if err := os.WriteFile(filepath.Join(incomingDir, "1.txt"), []byte(file1Content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(incomingDir, "2.txt"), []byte(file2Content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test file aggregation logic by calling the internal functions
	files, err := filepath.Glob(filepath.Join(incomingDir, "*.txt"))
	if err != nil {
		t.Fatalf("Failed to glob files: %v", err)
	}

	if len(files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(files))
	}

	// Test content aggregation
	var content strings.Builder
	var processingErrors []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			processingErrors = append(processingErrors, err.Error())
			continue
		}
		content.WriteString(string(data))
		content.WriteString("\n")
	}

	if len(processingErrors) > 0 {
		t.Fatalf("Unexpected processing errors: %v", processingErrors)
	}

	aggregatedContent := content.String()
	if !strings.Contains(aggregatedContent, "Feature A") {
		t.Errorf("Aggregated content should contain 'Feature A'")
	}
	if !strings.Contains(aggregatedContent, "Feature B") {
		t.Errorf("Aggregated content should contain 'Feature B'")
	}
}

func TestFileProcessing_ErrorHandling(t *testing.T) {
	tempDir := t.TempDir()
	incomingDir := filepath.Join(tempDir, "incoming")

	// Create incoming directory with one good file and one unreadable file
	if err := os.MkdirAll(incomingDir, 0755); err != nil {
		t.Fatalf("Failed to create incoming dir: %v", err)
	}

	// Create a good file
	if err := os.WriteFile(filepath.Join(incomingDir, "1.txt"), []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create an unreadable file (simulate permission error)
	unreadableFile := filepath.Join(incomingDir, "2.txt")
	if err := os.WriteFile(unreadableFile, []byte("content"), 0000); err != nil {
		t.Fatalf("Failed to create unreadable file: %v", err)
	}
	defer os.Chmod(unreadableFile, 0644) // Clean up

	// Test error aggregation logic
	files, err := filepath.Glob(filepath.Join(incomingDir, "*.txt"))
	if err != nil {
		t.Fatalf("Failed to glob files: %v", err)
	}

	var content strings.Builder
	var processingErrors []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			processingErrors = append(processingErrors, err.Error())
			continue
		}
		content.WriteString(string(data))
		content.WriteString("\n")
	}

	if len(processingErrors) == 0 {
		t.Errorf("Expected processing errors due to unreadable file")
	}

	// Verify error message format
	errorMsg := strings.Join(processingErrors, "; ")
	if !strings.Contains(errorMsg, "2.txt") {
		t.Errorf("Error message should mention the problematic file")
	}
}

func TestEnsureIncomingDirCreation(t *testing.T) {
	tempDir := t.TempDir()
	incomingDir := filepath.Join(tempDir, "incoming")

	cfg := &config.Config{
		IncomingDir: incomingDir,
	}

	g := &Generator{cfg: cfg}

	err := g.ensureIncomingDir()
	if err != nil {
		t.Fatalf("ensureIncomingDir failed: %v", err)
	}

	if _, err := os.Stat(incomingDir); os.IsNotExist(err) {
		t.Errorf("Incoming directory was not created")
	}
}
