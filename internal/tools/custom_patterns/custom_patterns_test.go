package custom_patterns

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCustomPatterns(t *testing.T) {
	plugin := NewCustomPatterns()

	assert.NotNil(t, plugin)
	assert.Equal(t, "Custom Patterns", plugin.GetName())
	assert.Equal(t, "Custom Patterns - Set directory for your custom patterns (optional)", plugin.GetSetupDescription())
	assert.False(t, plugin.IsConfigured()) // Should not be configured initially
}
func TestCustomPatterns_Configure(t *testing.T) {
	plugin := NewCustomPatterns()

	// Test with empty directory (should work)
	plugin.CustomPatternsDir.Value = ""
	err := plugin.configure()
	assert.NoError(t, err)

	// Test with home directory expansion
	plugin.CustomPatternsDir.Value = "~/test-patterns"
	err = plugin.configure()
	assert.NoError(t, err)

	homeDir, _ := os.UserHomeDir()
	expectedPath := filepath.Join(homeDir, "test-patterns")
	absExpected, _ := filepath.Abs(expectedPath)
	assert.Equal(t, absExpected, plugin.CustomPatternsDir.Value)

	// Clean up
	os.RemoveAll(plugin.CustomPatternsDir.Value)
}

func TestCustomPatterns_ConfigureWithTempDir(t *testing.T) {
	plugin := NewCustomPatterns()

	// Test with a temporary directory
	tmpDir, err := os.MkdirTemp("", "test-custom-patterns-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	plugin.CustomPatternsDir.Value = tmpDir
	err = plugin.configure()
	assert.NoError(t, err)

	absPath, _ := filepath.Abs(tmpDir)
	assert.Equal(t, absPath, plugin.CustomPatternsDir.Value)

	// Verify directory exists
	info, err := os.Stat(plugin.CustomPatternsDir.Value)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())

	// Should be configured now
	assert.True(t, plugin.IsConfigured())
}

func TestCustomPatterns_IsConfigured(t *testing.T) {
	plugin := NewCustomPatterns()

	// Initially not configured
	assert.False(t, plugin.IsConfigured())

	// Set a directory
	plugin.CustomPatternsDir.Value = "/some/path"
	assert.True(t, plugin.IsConfigured())

	// Clear the directory
	plugin.CustomPatternsDir.Value = ""
	assert.False(t, plugin.IsConfigured())
}
