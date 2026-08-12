package tools

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

func newTestPatternsLoader(t *testing.T) *PatternsLoader {
	t.Helper()

	patternsDir := filepath.Join(t.TempDir(), "patterns")
	return NewPatternsLoader(&fsdb.PatternsEntity{
		StorageEntity: &fsdb.StorageEntity{
			Label:     "Patterns",
			Dir:       patternsDir,
			ItemIsDir: true,
		},
	})
}

func countTempPatternsFolders(t *testing.T, dir string) int {
	t.Helper()

	matches, err := filepath.Glob(filepath.Join(dir, "fabric-patterns-*"))
	require.NoError(t, err)
	return len(matches)
}

// configure() runs on every plugin registry construction, i.e. on every fabric
// invocation. It must not create the scratch folder that only the pattern
// download path needs, otherwise each invocation leaks an empty temp directory.
func TestPatternsLoader_ConfigureDoesNotCreateTempFolder(t *testing.T) {
	tempRoot := t.TempDir()
	t.Setenv("TMPDIR", tempRoot)

	loader := newTestPatternsLoader(t)
	require.NoError(t, loader.configure())

	assert.Empty(t, loader.tempPatternsFolder)
	assert.Equal(t, 0, countTempPatternsFolders(t, tempRoot))

	// Repeated invocations must not accumulate directories either.
	for i := 0; i < 5; i++ {
		require.NoError(t, loader.configure())
	}
	assert.Equal(t, 0, countTempPatternsFolders(t, tempRoot))
}

func TestPatternsLoader_EnsureTempPatternsFolderIsLazyAndIdempotent(t *testing.T) {
	tempRoot := t.TempDir()
	t.Setenv("TMPDIR", tempRoot)

	loader := newTestPatternsLoader(t)
	require.NoError(t, loader.configure())

	require.NoError(t, loader.ensureTempPatternsFolder())
	created := loader.tempPatternsFolder
	require.NotEmpty(t, created)
	assert.Equal(t, 1, countTempPatternsFolders(t, tempRoot))

	info, err := os.Stat(created)
	require.NoError(t, err)
	assert.True(t, info.IsDir())

	// Calling it again reuses the existing folder rather than creating another.
	require.NoError(t, loader.ensureTempPatternsFolder())
	assert.Equal(t, created, loader.tempPatternsFolder)
	assert.Equal(t, 1, countTempPatternsFolders(t, tempRoot))
}
