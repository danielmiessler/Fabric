package pipeline

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFilePathUsesUserDefinitionSource(t *testing.T) {
	tempDir := t.TempDir()
	pipelinePath := filepath.Join(tempDir, "sample.yaml")
	require.NoError(t, os.WriteFile(pipelinePath, []byte(validPipelineYAML("sample")), 0o644))

	loader := &Loader{}
	pipe, err := loader.LoadFilePath(pipelinePath)
	require.NoError(t, err)
	require.Equal(t, DefinitionSourceUser, pipe.DefinitionSource)
	require.Equal(t, "sample.yaml", pipe.FileName)
	require.Equal(t, "sample", pipe.FileStem)
}

func TestPreflightHonorsCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	p := &Pipeline{
		Version:  1,
		Name:     "cancelled-preflight",
		FileStem: "cancelled-preflight",
		Stages: []Stage{
			{
				ID:            "render",
				Executor:      ExecutorBuiltin,
				Builtin:       &BuiltinConfig{Name: "passthrough"},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	err := Preflight(ctx, p, PreflightOptions{})
	require.ErrorIs(t, err, context.Canceled)
}

func TestPreflightSkipsPatternValidationWhenRequested(t *testing.T) {
	p := &Pipeline{
		Version:  1,
		Name:     "skip-pattern-validation",
		FileStem: "skip-pattern-validation",
		Stages: []Stage{
			{
				ID:            "render",
				Executor:      ExecutorFabricPattern,
				Pattern:       "missing-pattern",
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	err := Preflight(context.Background(), p, PreflightOptions{
		SkipPatternValidation: func(stage Stage) bool {
			return stage.ID == "render"
		},
	})
	require.NoError(t, err)
}
