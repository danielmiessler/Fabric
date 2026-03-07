package cli

import (
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
			InvocationDir:  invocationDir,
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
