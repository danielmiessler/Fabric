package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanFiles(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "test1.go")
	file2 := filepath.Join(tmpDir, "test2.go")
	subDir := filepath.Join(tmpDir, "subdir")
	file3 := filepath.Join(subDir, "test3.go")

	require.NoError(t, os.WriteFile(file1, []byte("package main\n"), 0644))
	require.NoError(t, os.WriteFile(file2, []byte("package main\n\nfunc main() {}\n"), 0644))
	require.NoError(t, os.MkdirAll(subDir, 0755))
	require.NoError(t, os.WriteFile(file3, []byte("package subdir\n"), 0644))

	files := []string{file1, file3}
	instructions := "Test instructions"

	jsonData, err := ScanFiles(files, instructions)
	require.NoError(t, err)

	var result []any
	err = json.Unmarshal(jsonData, &result)
	require.NoError(t, err)
	assert.Len(t, result, 3) // directory, report, instructions

	report := result[1].(map[string]any)
	assert.Equal(t, "report", report["type"])
	assert.Equal(t, float64(2), report["files"])

	instr := result[2].(map[string]any)
	assert.Equal(t, "instructions", instr["type"])
	assert.Equal(t, "Test instructions", instr["details"])
}

func TestScanFilesSkipsDirectories(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "test.go")
	subDir := filepath.Join(tmpDir, "subdir")

	require.NoError(t, os.WriteFile(file1, []byte("package main\n"), 0644))
	require.NoError(t, os.MkdirAll(subDir, 0755))

	files := []string{file1, subDir}

	jsonData, err := ScanFiles(files, "test")
	require.NoError(t, err)

	var result []any
	err = json.Unmarshal(jsonData, &result)
	require.NoError(t, err)

	report := result[1].(map[string]any)
	assert.Equal(t, float64(1), report["files"])
}

func TestScanFilesNonExistentFile(t *testing.T) {
	files := []string{"/nonexistent/file.go"}
	_, err := ScanFiles(files, "test")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error accessing file")
}

func TestScanDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "main.go")
	require.NoError(t, os.WriteFile(file1, []byte("package main\n"), 0644))

	jsonData, err := ScanDirectory(tmpDir, 3, "Test instructions", []string{})
	require.NoError(t, err)

	var result []any
	err = json.Unmarshal(jsonData, &result)
	require.NoError(t, err)
	assert.Len(t, result, 3)

	instr := result[2].(map[string]any)
	assert.Equal(t, "Test instructions", instr["details"])
}
