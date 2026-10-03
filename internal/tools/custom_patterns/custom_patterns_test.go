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
	assert.False(t, plugin.IsConfigured())
}
func TestCustomPatterns_Configure(t *testing.T) {
	plugin := NewCustomPatterns()

	plugin.CustomPatternsDir.Value = ""
	err := plugin.configure()
	assert.NoError(t, err)

	plugin.CustomPatternsDir.Value = "~/test-patterns"
	err = plugin.configure()
	assert.NoError(t, err)

	homeDir, _ := os.UserHomeDir()
	expectedPath := filepath.Join(homeDir, "test-patterns")
	absExpected, _ := filepath.Abs(expectedPath)
	assert.Equal(t, absExpected, plugin.CustomPatternsDir.Value)

	os.RemoveAll(plugin.CustomPatternsDir.Value)
}

func TestCustomPatterns_ConfigureWithTempDir(t *testing.T) {
	plugin := NewCustomPatterns()

	tmpDir, err := os.MkdirTemp("", "test-custom-patterns-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	plugin.CustomPatternsDir.Value = tmpDir
	err = plugin.configure()
	assert.NoError(t, err)

	absPath, _ := filepath.Abs(tmpDir)
	assert.Equal(t, absPath, plugin.CustomPatternsDir.Value)

	info, err := os.Stat(plugin.CustomPatternsDir.Value)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())

	assert.True(t, plugin.IsConfigured())
}

func TestCustomPatterns_IsConfigured(t *testing.T) {
	plugin := NewCustomPatterns()

	assert.False(t, plugin.IsConfigured())

	plugin.CustomPatternsDir.Value = "/some/path"
	assert.True(t, plugin.IsConfigured())

	plugin.CustomPatternsDir.Value = ""
	assert.False(t, plugin.IsConfigured())
}
