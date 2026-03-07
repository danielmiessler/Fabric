package pipeline

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateRequiresMatchingPipelineNameAndSingleFinalOutput(t *testing.T) {
	t.Parallel()

	p := &Pipeline{
		Version:  1,
		Name:     "tech-note",
		FileStem: "other-name",
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

	err := Validate(p)
	require.Error(t, err)
	require.Contains(t, err.Error(), "must match filename stem")
}

func TestValidateRejectsDuplicateStageIDs(t *testing.T) {
	t.Parallel()

	p := &Pipeline{
		Version:  1,
		Name:     "tech-note",
		FileStem: "tech-note",
		Stages: []Stage{
			{ID: "same", Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "noop"}},
			{ID: "same", Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "passthrough"}, FinalOutput: true, PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout}},
		},
	}

	err := Validate(p)
	require.Error(t, err)
	require.Contains(t, err.Error(), "duplicate stage id")
}

func TestLoaderListsUserOverrides(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")
	userDir := filepath.Join(tempDir, "user")
	require.NoError(t, os.MkdirAll(builtInDir, 0o755))
	require.NoError(t, os.MkdirAll(userDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "tech-note.yaml"), []byte(validPipelineYAML("tech-note")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(userDir, "tech-note.yaml"), []byte(validPipelineYAML("tech-note")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "passthrough.yaml"), []byte(validPipelineYAML("passthrough")), 0o644))

	loader := &Loader{BuiltInDir: builtInDir, UserDir: userDir}
	entries, err := loader.List()
	require.NoError(t, err)
	require.Len(t, entries, 2)

	var overrideEntry *DiscoveryEntry
	for i := range entries {
		if entries[i].Name == "tech-note" {
			overrideEntry = &entries[i]
			break
		}
	}
	require.NotNil(t, overrideEntry)
	require.Equal(t, DefinitionSourceUser, overrideEntry.DefinitionSource)
	require.True(t, overrideEntry.OverridesBuiltIn)
}

func TestPreflightRejectsMissingEnvironmentVariables(t *testing.T) {
	t.Parallel()

	p := &Pipeline{
		Version:  1,
		Name:     "env-check",
		FileStem: "env-check",
		Stages: []Stage{
			{
				ID:       "run",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: "echo",
					Args:    []string{"${MISSING_PIPELINE_TEST_VAR}"},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	err := Preflight(context.Background(), p, PreflightOptions{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "MISSING_PIPELINE_TEST_VAR")
}

func TestRunnerCreatesRunArtifactsForPassthroughPipeline(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "passthrough",
		FileStem: "passthrough",
		FilePath: filepath.Join(tempDir, "passthrough.yaml"),
		Stages: []Stage{
			{
				ID:            "passthrough",
				Executor:      ExecutorBuiltin,
				Builtin:       &BuiltinConfig{Name: "passthrough"},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	var stdout, stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "hello world"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.NoError(t, err)
	require.NotEmpty(t, result.RunDir)
	require.FileExists(t, filepath.Join(result.RunDir, "run_manifest.json"))
	require.FileExists(t, filepath.Join(result.RunDir, "run.json"))
	require.FileExists(t, filepath.Join(result.RunDir, "source_manifest.json"))
	require.Equal(t, "hello world", result.FinalOutput)
	require.Contains(t, stdout.String(), "hello world")
	require.Contains(t, stderr.String(), "PASS")
	require.Contains(t, stderr.String(), "warning: pipeline passthrough has no validate stage")
}

func TestRunnerExecutesCommandStageAndCapturesStdout(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "command-stdout",
		FileStem: "command-stdout",
		FilePath: filepath.Join(tempDir, "command-stdout.yaml"),
		Stages: []Stage{
			{
				ID:       "run",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args: []string{
						"-test.run=TestPipelineHelperProcess",
						"--",
						"stdout",
						"hello-from-command",
					},
					Env: map[string]string{
						"GO_WANT_HELPER_PROCESS": "1",
					},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	var stdout, stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "ignored"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.NoError(t, err)
	require.Equal(t, "hello-from-command", result.FinalOutput)
	require.Contains(t, stdout.String(), "hello-from-command")
}

func TestRunnerUsesArtifactPrimaryOutput(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "artifact-output",
		FileStem: "artifact-output",
		FilePath: filepath.Join(tempDir, "artifact-output.yaml"),
		Stages: []Stage{
			{
				ID:       "generate",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args: []string{
						"-test.run=TestPipelineHelperProcess",
						"--",
						"artifact",
						"artifact-note",
					},
					Env: map[string]string{
						"GO_WANT_HELPER_PROCESS": "1",
					},
				},
				Artifacts: []ArtifactDeclaration{
					{Name: "note", Path: "note.md"},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputArtifact, Artifact: "note"},
			},
		},
	}

	var stdout, stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "ignored"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.NoError(t, err)
	require.Equal(t, "artifact-note", result.FinalOutput)
	require.Contains(t, stdout.String(), "artifact-note")
	require.FileExists(t, filepath.Join(result.RunDir, "note.md"))
}

func TestRunnerValidationStageBlocksFinalStdoutOnFailure(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "validate-blocks-output",
		FileStem: "validate-blocks-output",
		FilePath: filepath.Join(tempDir, "validate-blocks-output.yaml"),
		Stages: []Stage{
			{
				ID:            "render",
				Executor:      ExecutorBuiltin,
				Builtin:       &BuiltinConfig{Name: "passthrough"},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
			{
				ID:       "validate",
				Role:     StageRoleValidate,
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args:    []string{"-test.run=TestPipelineHelperProcess", "--", "fail"},
					Env:     map[string]string{"GO_WANT_HELPER_PROCESS": "1"},
				},
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	var stdout, stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "candidate-final-output"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.Nil(t, result)
	require.Error(t, err)
	require.Empty(t, stdout.String())
	require.Contains(t, stderr.String(), "validate ........ FAIL")
}

func TestRunnerPublishFailureStillEmitsValidatedOutput(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "publish-failure-still-emits",
		FileStem: "publish-failure-still-emits",
		FilePath: filepath.Join(tempDir, "publish-failure-still-emits.yaml"),
		Stages: []Stage{
			{
				ID:            "render",
				Executor:      ExecutorBuiltin,
				Builtin:       &BuiltinConfig{Name: "passthrough"},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
			{
				ID:       "validate",
				Executor: ExecutorBuiltin,
				Builtin:  &BuiltinConfig{Name: "validate_declared_outputs"},
			},
			{
				ID:       "publish",
				Role:     StageRolePublish,
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args:    []string{"-test.run=TestPipelineHelperProcess", "--", "fail"},
					Env:     map[string]string{"GO_WANT_HELPER_PROCESS": "1"},
				},
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	var stdout, stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "validated-output"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.Nil(t, result)
	require.Error(t, err)
	require.Contains(t, stdout.String(), "validated-output")
	require.Contains(t, stderr.String(), "publish ........ FAIL")
}

func TestCommandStageInterpolatesSourceReference(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	sourcePath := filepath.Join(tempDir, "session.txt")
	require.NoError(t, os.WriteFile(sourcePath, []byte("ignored"), 0o644))

	p := &Pipeline{
		Version:  1,
		Name:     "source-interpolation",
		FileStem: "source-interpolation",
		FilePath: filepath.Join(tempDir, "source-interpolation.yaml"),
		Stages: []Stage{
			{
				ID:       "normalize",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args: []string{
						"-test.run=TestPipelineHelperProcess",
						"--",
						"stdout",
						"{{source}}",
					},
					Env: map[string]string{"GO_WANT_HELPER_PROCESS": "1"},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
			},
		},
	}

	var stdout bytes.Buffer
	runner := NewRunner(&stdout, io.Discard, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeSource, Reference: sourcePath, Payload: "ignored"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.NoError(t, err)
	require.Equal(t, sourcePath, result.FinalOutput)
	require.Contains(t, stdout.String(), sourcePath)
}

func TestValidateBuiltinWritesValidationReport(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	p := &Pipeline{
		Version:  1,
		Name:     "validation-report",
		FileStem: "validation-report",
		FilePath: filepath.Join(tempDir, "validation-report.yaml"),
		Stages: []Stage{
			{
				ID:       "render",
				Executor: ExecutorCommand,
				Command: &CommandConfig{
					Program: os.Args[0],
					Args:    []string{"-test.run=TestPipelineHelperProcess", "--", "artifact", "artifact-note"},
					Env:     map[string]string{"GO_WANT_HELPER_PROCESS": "1"},
				},
				Artifacts: []ArtifactDeclaration{
					{Name: "note", Path: "note.md"},
				},
				FinalOutput:   true,
				PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputArtifact, Artifact: "note"},
			},
			{
				ID:       "validate",
				Executor: ExecutorBuiltin,
				Builtin:  &BuiltinConfig{Name: "validate_declared_outputs"},
			},
		},
	}

	runner := NewRunner(io.Discard, io.Discard, nil)
	result, err := runner.Run(context.Background(), p, RunSource{Mode: SourceModeStdin, Payload: "ignored"}, RunOptions{InvocationDir: tempDir, DisableCleanup: true})
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(result.RunDir, "validation_report.md"))
}

func TestCleanupExpiredRunsRemovesExpiredDirectories(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	runRoot := filepath.Join(tempDir, ".pipeline")
	runDir := filepath.Join(runRoot, "expired-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	expiredAt := time.Now().UTC().Add(-1 * time.Minute)
	state := &RunState{
		RunID:       "expired-run",
		Status:      "completed",
		StartedAt:   expiredAt.Add(-1 * time.Minute),
		UpdatedAt:   expiredAt,
		CompletedAt: &expiredAt,
		ExpiresAt:   &expiredAt,
		RunDir:      runDir,
	}
	require.NoError(t, writeJSON(filepath.Join(runDir, "run.json"), state))

	require.NoError(t, cleanupExpiredRuns(runRoot, time.Now().UTC()))
	_, err := os.Stat(runRoot)
	require.True(t, os.IsNotExist(err))
}

func TestPipelineHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	args := os.Args
	sep := -1
	for i := range args {
		if args[i] == "--" {
			sep = i
			break
		}
	}
	if sep == -1 || sep+1 >= len(args) {
		os.Exit(2)
	}

	mode := args[sep+1]
	switch mode {
	case "stdout":
		fmt.Print(args[sep+2])
	case "artifact":
		target := os.Getenv("FABRIC_PIPELINE_ARTIFACT_NOTE")
		if target == "" {
			fmt.Fprintln(os.Stderr, "missing FABRIC_PIPELINE_ARTIFACT_NOTE")
			os.Exit(3)
		}
		if err := os.WriteFile(target, []byte(args[sep+2]), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(4)
		}
	case "fail":
		fmt.Fprintln(os.Stderr, "intentional helper failure")
		os.Exit(7)
	default:
		os.Exit(5)
	}
	os.Exit(0)
}

func validPipelineYAML(name string) string {
	return "version: 1\nname: " + name + "\nstages:\n  - id: passthrough\n    executor: builtin\n    builtin:\n      name: passthrough\n    final_output: true\n    primary_output:\n      from: stdout\n"
}
