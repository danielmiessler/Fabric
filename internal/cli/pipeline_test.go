package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/pipeline"
	"github.com/stretchr/testify/require"
)

func TestHandlePipelineCommandsWritesOutputFileOnPublishFailure(t *testing.T) {
	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")
	require.NoError(t, os.MkdirAll(builtInDir, 0o755))

	pipelineYAML := "version: 1\n" +
		"name: publish-cli\n" +
		"stages:\n" +
		"  - id: render\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: passthrough\n" +
		"    final_output: true\n" +
		"    primary_output:\n" +
		"      from: stdout\n" +
		"  - id: validate\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: validate_declared_outputs\n" +
		"  - id: publish\n" +
		"    role: publish\n" +
		"    executor: command\n" +
		"    command:\n" +
		"      program: " + os.Args[0] + "\n" +
		"      args:\n" +
		"        - -test.run=TestPipelineCLIHelperProcess\n" +
		"        - --\n" +
		"        - fail\n" +
		"      env:\n" +
		"        GO_WANT_HELPER_PROCESS: \"1\"\n" +
		"    primary_output:\n" +
		"      from: stdout\n"
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "publish-cli.yaml"), []byte(pipelineYAML), 0o644))

	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", builtInDir)
	outputPath := filepath.Join(tempDir, "final.md")

	originalRunOptions := pipelineRunOptions
	pipelineRunOptions = func(invocationDir string) pipeline.RunOptions {
		return pipeline.RunOptions{
			InvocationDir:  tempDir,
			DisableCleanup: true,
		}
	}
	t.Cleanup(func() { pipelineRunOptions = originalRunOptions })

	flags := &Flags{
		Pipeline:      "publish-cli",
		Output:        outputPath,
		stdinProvided: true,
		stdinMessage:  "validated-output",
	}

	handled, err := handlePipelineCommands(flags, nil)
	require.True(t, handled)
	require.Error(t, err)
	require.FileExists(t, outputPath)
	content, readErr := os.ReadFile(outputPath)
	require.NoError(t, readErr)
	require.Equal(t, "validated-output\n", string(content))
}

func TestHandlePipelineCommandsDryRunEmitsPlanWithoutExecutingStages(t *testing.T) {
	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")
	require.NoError(t, os.MkdirAll(builtInDir, 0o755))

	pipelineYAML := "version: 1\n" +
		"name: dry-run-cli\n" +
		"stages:\n" +
		"  - id: fail-if-executed\n" +
		"    executor: command\n" +
		"    command:\n" +
		"      program: " + os.Args[0] + "\n" +
		"      args:\n" +
		"        - -test.run=TestPipelineCLIHelperProcess\n" +
		"        - --\n" +
		"        - fail\n" +
		"      env:\n" +
		"        GO_WANT_HELPER_PROCESS: \"1\"\n" +
		"    final_output: true\n" +
		"    primary_output:\n" +
		"      from: stdout\n"
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "dry-run-cli.yaml"), []byte(pipelineYAML), 0o644))
	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", builtInDir)

	flags := &Flags{
		Pipeline:      "dry-run-cli",
		DryRun:        true,
		stdinProvided: true,
		stdinMessage:  "input",
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handlePipelineCommands(flags, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)

	var plan map[string]any
	require.NoError(t, json.Unmarshal([]byte(stdout), &plan))
	require.Equal(t, "dry-run-cli", plan["pipeline"])
	selected, ok := plan["selected_stage_ids"].([]any)
	require.True(t, ok)
	require.Len(t, selected, 1)
	require.Equal(t, "fail-if-executed", selected[0])
}

func TestHandlePipelineCommandsRejectsRuntimeFlagsWithoutPipeline(t *testing.T) {
	flags := &Flags{FromStage: "prepare"}
	handled, err := handlePipelineCommands(flags, nil)
	require.True(t, handled)
	require.Error(t, err)
	require.Contains(t, err.Error(), "require --pipeline")
}

func TestHandlePipelineCommandsRejectsValidateOnlyWithValidatePipeline(t *testing.T) {
	tempDir := t.TempDir()
	pipelinePath := filepath.Join(tempDir, "pipeline.yaml")
	pipelineYAML := "version: 1\n" +
		"name: validate-conflict\n" +
		"stages:\n" +
		"  - id: pass\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: passthrough\n" +
		"    final_output: true\n" +
		"    primary_output:\n" +
		"      from: stdout\n"
	require.NoError(t, os.WriteFile(pipelinePath, []byte(pipelineYAML), 0o644))

	flags := &Flags{
		ValidatePipeline: pipelinePath,
		ValidateOnly:     true,
	}
	handled, err := handlePipelineCommands(flags, nil)
	require.True(t, handled)
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot be used with --validate-pipeline")
}

func TestHandlePipelineCommandsDryRunRespectsStageSelection(t *testing.T) {
	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")
	require.NoError(t, os.MkdirAll(builtInDir, 0o755))

	pipelineYAML := "version: 1\n" +
		"name: dry-run-selection\n" +
		"stages:\n" +
		"  - id: one\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: passthrough\n" +
		"  - id: two\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: passthrough\n" +
		"    final_output: true\n" +
		"    primary_output:\n" +
		"      from: stdout\n" +
		"  - id: three\n" +
		"    role: publish\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: write_publish_manifest\n"
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "dry-run-selection.yaml"), []byte(pipelineYAML), 0o644))
	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", builtInDir)

	flags := &Flags{
		Pipeline:      "dry-run-selection",
		DryRun:        true,
		FromStage:     "two",
		ToStage:       "three",
		stdinProvided: true,
		stdinMessage:  "input",
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handlePipelineCommands(flags, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)

	var plan struct {
		SelectedStageIDs []string `json:"selected_stage_ids"`
		SkippedStageIDs  []string `json:"skipped_stage_ids"`
	}
	require.NoError(t, json.Unmarshal([]byte(stdout), &plan))
	require.Equal(t, []string{"two", "three"}, plan.SelectedStageIDs)
	require.Equal(t, []string{"one"}, plan.SkippedStageIDs)
}

func captureStdout(run func() error) (string, error) {
	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = w

	runErr := run()
	closeErr := w.Close()
	os.Stdout = original

	var buffer bytes.Buffer
	if _, copyErr := io.Copy(&buffer, r); copyErr != nil {
		return "", copyErr
	}
	if err := r.Close(); err != nil {
		return "", err
	}
	if runErr != nil {
		return "", runErr
	}
	if closeErr != nil {
		return "", closeErr
	}
	return buffer.String(), nil
}

func TestPipelineCLIHelperProcess(t *testing.T) {
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

	switch args[sep+1] {
	case "fail":
		_, _ = os.Stderr.WriteString("intentional helper failure\n")
		os.Exit(7)
	default:
		os.Exit(5)
	}
}
