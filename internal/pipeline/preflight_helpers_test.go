package pipeline

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadRawPatternContent(t *testing.T) {
	registry := newDryRunRegistryForPatternTests(t)
	writeNamedPattern(t, registry.Db, "named-pattern", "Hello {{input}}")

	content, err := loadRawPatternContent("named-pattern", registry)
	require.NoError(t, err)
	require.Equal(t, "Hello {{input}}", content)

	patternFile := filepath.Join(t.TempDir(), "pattern.md")
	require.NoError(t, os.WriteFile(patternFile, []byte("File pattern {{input}}"), 0o644))

	content, err = loadRawPatternContent(patternFile, registry)
	require.NoError(t, err)
	require.Equal(t, "File pattern {{input}}", content)
}

func TestPreflightCommandStageResolvesRelativeProgramAndCwd(t *testing.T) {
	tempDir := t.TempDir()
	scriptPath := filepath.Join(tempDir, "script.sh")
	require.NoError(t, os.WriteFile(scriptPath, []byte("#!/bin/sh\nexit 0\n"), 0o755))

	workingDir := filepath.Join(tempDir, "subdir")
	require.NoError(t, os.MkdirAll(workingDir, 0o755))

	p := &Pipeline{
		Version:  1,
		Name:     "command-preflight",
		FileStem: "command-preflight",
		FilePath: filepath.Join(tempDir, "command-preflight.yaml"),
		Stages: []Stage{
			{
				ID:       "run",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: "./script.sh",
					Cwd:     "./subdir",
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	require.NoError(t, Preflight(context.Background(), p, PreflightOptions{}))
	require.Equal(t, scriptPath, p.Stages[0].Command.Program)
	require.Equal(t, workingDir, p.Stages[0].Command.Cwd)
}

func TestResolveCommandProgram(t *testing.T) {
	tempDir := t.TempDir()
	executable := filepath.Join(tempDir, "tool.sh")
	require.NoError(t, os.WriteFile(executable, []byte("#!/bin/sh\nexit 0\n"), 0o755))

	p := &Pipeline{FilePath: filepath.Join(tempDir, "pipeline.yaml")}

	resolved, err := resolveCommandProgram(p, executable)
	require.NoError(t, err)
	require.Equal(t, executable, resolved)

	resolved, err = resolveCommandProgram(p, "./tool.sh")
	require.NoError(t, err)
	require.Equal(t, executable, resolved)

	resolved, err = resolveCommandProgram(p, "sh")
	require.NoError(t, err)
	require.NotEmpty(t, resolved)

	_, err = resolveCommandProgram(p, "./missing.sh")
	require.Error(t, err)
}
