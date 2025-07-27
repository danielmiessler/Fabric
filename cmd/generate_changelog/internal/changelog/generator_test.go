package changelog

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/config"
)

func TestDetectVersionFromNix(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("version.nix exists", func(t *testing.T) {
		versionNixContent := `"1.2.3"`
		versionNixPath := filepath.Join(tempDir, "version.nix")
		err := os.WriteFile(versionNixPath, []byte(versionNixContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write version.nix: %v", err)
		}

		data, err := os.ReadFile(versionNixPath)
		if err != nil {
			t.Fatalf("Failed to read version.nix: %v", err)
		}

		versionRegex := regexp.MustCompile(`"([^"]+)"`)
		matches := versionRegex.FindStringSubmatch(string(data))

		if len(matches) <= 1 {
			t.Fatalf("No version found in version.nix")
		}

		version := matches[1]
		if version != "1.2.3" {
			t.Errorf("Expected version 1.2.3, got %s", version)
		}
	})
}

func TestEnsureIncomingDir(t *testing.T) {
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

func TestInsertVersionAtTop(t *testing.T) {
	tempDir := t.TempDir()
	changelogPath := filepath.Join(tempDir, "CHANGELOG.md")

	cfg := &config.Config{
		RepoPath: tempDir,
	}

	g := &Generator{cfg: cfg}

	t.Run("new changelog", func(t *testing.T) {
		entry := "## v1.0.0 (2025-01-01)\n\n- Initial release"

		err := g.insertVersionAtTop(entry)
		if err != nil {
			t.Fatalf("insertVersionAtTop failed: %v", err)
		}

		content, err := os.ReadFile(changelogPath)
		if err != nil {
			t.Fatalf("Failed to read changelog: %v", err)
		}

		expected := "# Changelog\n\n## v1.0.0 (2025-01-01)\n\n- Initial release\n"
		if string(content) != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(content))
		}
	})

	t.Run("existing changelog", func(t *testing.T) {
		existingContent := "# Changelog\n\n## v0.9.0 (2024-12-01)\n\n- Previous release"
		err := os.WriteFile(changelogPath, []byte(existingContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write existing changelog: %v", err)
		}

		entry := "## v1.0.0 (2025-01-01)\n\n- New release"

		err = g.insertVersionAtTop(entry)
		if err != nil {
			t.Fatalf("insertVersionAtTop failed: %v", err)
		}

		content, err := os.ReadFile(changelogPath)
		if err != nil {
			t.Fatalf("Failed to read changelog: %v", err)
		}

		expected := "# Changelog\n\n## v1.0.0 (2025-01-01)\n\n- New release\n## v0.9.0 (2024-12-01)\n\n- Previous release"
		if string(content) != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(content))
		}
	})
}
