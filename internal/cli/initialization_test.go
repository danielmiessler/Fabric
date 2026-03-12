package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnsureEnvFileCreatesAndPreservesFile(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	envPath := filepath.Join(homeDir, ".config", "fabric", ".env")

	require.NoError(t, ensureEnvFile())
	require.FileExists(t, envPath)

	require.NoError(t, os.WriteFile(envPath, []byte("EXISTING=1\n"), 0o644))
	require.NoError(t, ensureEnvFile())

	content, err := os.ReadFile(envPath)
	require.NoError(t, err)
	require.Equal(t, "EXISTING=1\n", string(content))
}

func TestInitializeFabricCreatesRegistryForTempHome(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	require.NoError(t, ensureEnvFile())

	registry, err := initializeFabric()
	require.NoError(t, err)
	require.NotNil(t, registry)
	require.NotNil(t, registry.Db)
	require.DirExists(t, filepath.Join(homeDir, ".config", "fabric"))
	require.DirExists(t, registry.Db.Patterns.Dir)
	require.DirExists(t, registry.Db.Sessions.Dir)
	require.DirExists(t, registry.Db.Contexts.Dir)
}
