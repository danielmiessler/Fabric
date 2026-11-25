package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebSearchDoesNotAllowMissingEnvForOtherCommands(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })

	os.Args = []string{"fabric", "--listvendors"}
	require.ErrorContains(t, Cli("test"), ".env")
	os.Args = append(os.Args, "--web-search=query")
	require.ErrorContains(t, Cli("test"), ".env")
}

func TestWebSearchRejectsNullProviderInDefaultConfig(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })
	configDir := filepath.Join(home, ".config", "fabric")
	require.NoError(t, os.MkdirAll(configDir, 0700))
	configPath := filepath.Join(configDir, "config.yaml")

	for _, value := range []string{"", "null", "\"\""} {
		require.NoError(t, os.WriteFile(configPath, []byte("webSearchProvider: "+value+"\n"), 0600))
		os.Args = []string{"fabric", "--web-search=query"}
		flags, err := Init()
		require.NoError(t, err)
		require.ErrorContains(t, flags.resolveWebSearch(), "Unknown standalone web search provider")
		os.Args = append(os.Args, "--web-search-provider=off")
		flags, err = Init()
		require.NoError(t, err)
		require.ErrorContains(t, flags.resolveWebSearch(), "disabled")
	}
}

func TestWebSearchReportsDanglingDefaultConfig(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })
	configDir := filepath.Join(home, ".config", "fabric")
	require.NoError(t, os.MkdirAll(configDir, 0700))
	if err := os.Symlink(filepath.Join(home, "offline.yaml"), filepath.Join(configDir, "config.yaml")); err != nil {
		t.Skipf("symbolic links unavailable: %v", err)
	}

	os.Args = []string{"fabric", "--web-search=query", "--web-search-provider=off"}
	require.ErrorContains(t, Cli("test"), "config.yaml")
}

func TestWebSearchRejectsExplicitNamedSearch(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })

	for _, option := range []string{
		"--parallel_search=", "--firecrawl_search=query", "--keenable_search=query",
		"--searchapi_search=query", "--serply_search=query", "--sofya_search=query", "-q=query",
	} {
		os.Args = []string{"fabric", "--web-search=query", option}
		flags, err := Init()
		require.NoError(t, err)
		require.ErrorContains(t, flags.resolveWebSearch(), "provider-specific search flags")
	}
}
