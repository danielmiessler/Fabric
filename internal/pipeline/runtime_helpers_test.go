package pipeline

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/dryrun"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
	"github.com/stretchr/testify/require"
)

func TestPreflightFabricPatternStage(t *testing.T) {
	t.Run("accepts configured pattern and context", func(t *testing.T) {
		registry := newDryRunRegistryForPatternTests(t)
		writeNamedPattern(t, registry.Db, "example-pattern", "Pattern src={{src}}\n{{input}}")
		writeNamedContext(t, registry.Db, "example-context", "Context guidance")

		p := &Pipeline{
			Version: 1,
			Name:    "pattern-preflight",
			Stages: []Stage{
				{
					ID:            "render",
					Executor:      ExecutorFabricPattern,
					Pattern:       "example-pattern",
					Context:       "example-context",
					Variables:     map[string]string{"src": "{{source_reference}}"},
					FinalOutput:   true,
					PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
				},
			},
		}

		require.NoError(t, Preflight(context.Background(), p, PreflightOptions{Registry: registry}))
	})

	t.Run("rejects missing registry", func(t *testing.T) {
		p := &Pipeline{
			Version: 1,
			Name:    "pattern-preflight",
			Stages: []Stage{
				{
					ID:            "render",
					Executor:      ExecutorFabricPattern,
					Pattern:       "example-pattern",
					FinalOutput:   true,
					PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
				},
			},
		}

		err := Preflight(context.Background(), p, PreflightOptions{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot validate patterns without plugin registry")
	})

	t.Run("rejects missing pattern", func(t *testing.T) {
		registry := newDryRunRegistryForPatternTests(t)
		p := &Pipeline{
			Version: 1,
			Name:    "pattern-preflight",
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

		err := Preflight(context.Background(), p, PreflightOptions{Registry: registry})
		require.Error(t, err)
		require.Contains(t, err.Error(), "missing-pattern")
	})

	t.Run("rejects missing declared pattern variables", func(t *testing.T) {
		registry := newDryRunRegistryForPatternTests(t)
		writeNamedPattern(t, registry.Db, "example-pattern", "Pattern src={{src}}\n{{input}}")

		p := &Pipeline{
			Version: 1,
			Name:    "pattern-preflight",
			Stages: []Stage{
				{
					ID:            "render",
					Executor:      ExecutorFabricPattern,
					Pattern:       "example-pattern",
					FinalOutput:   true,
					PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
				},
			},
		}

		err := Preflight(context.Background(), p, PreflightOptions{Registry: registry})
		require.Error(t, err)
		require.Contains(t, err.Error(), "missing required variable: src")
	})
}

func TestExecutePatternStageWithDryRunVendor(t *testing.T) {
	registry := newDryRunRegistryForPatternTests(t)
	writeNamedPattern(t, registry.Db, "example-pattern", "Pattern src={{src}}\nPrompt input={{input}}")
	writeNamedContext(t, registry.Db, "example-context", "Context guidance")

	runner := NewRunner(&bytes.Buffer{}, &bytes.Buffer{}, registry)
	stage := Stage{
		ID:       "render",
		Executor: ExecutorFabricPattern,
		Pattern:  "example-pattern",
		Context:  "example-context",
		Variables: map[string]string{
			"src": "{{source_reference}}",
		},
	}
	runtimeCtx := StageRuntimeContext{
		Pipeline: &Pipeline{Name: "pattern-exec"},
		Stage:    stage,
		Source: RunSource{
			Mode:      SourceModeSource,
			Reference: "/tmp/source.md",
			Payload:   "pipeline input",
		},
		InputPayload:  "pipeline input",
		InvocationDir: t.TempDir(),
		RunDir:        t.TempDir(),
		RunID:         "run-1",
	}

	result, err := runner.executePatternStage(context.Background(), stage, runtimeCtx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, result.Stdout, "Dry run: Would send the following request")
	require.Contains(t, result.Stdout, dryrun.DryRunResponse)
	require.Contains(t, result.Stdout, "Model: dry-run-model")
	require.Contains(t, result.Stdout, "Context guidance")
	require.Contains(t, result.Stdout, "Pattern src=/tmp/source.md")
	require.Contains(t, result.Stdout, "Prompt input=pipeline input")
}

func TestResolveRuntimePlaceholder(t *testing.T) {
	t.Parallel()

	runtimeCtx := StageRuntimeContext{
		Stage:          Stage{ID: "render"},
		Source:         RunSource{Reference: "/tmp/source.txt", Payload: "source payload"},
		InputPayload:   "input payload",
		InvocationDir:  "/tmp/invocation",
		RunDir:         "/tmp/run",
		RunID:          "run-123",
		StageArtifacts: map[string]map[string]string{"generate": {"note": "/tmp/run/note.md"}},
	}

	value, err := resolveRuntimePlaceholder("source", runtimeCtx)
	require.NoError(t, err)
	require.Equal(t, "/tmp/source.txt", value)

	runtimeCtx.Source.Reference = ""
	value, err = resolveRuntimePlaceholder("source", runtimeCtx)
	require.NoError(t, err)
	require.Equal(t, "source payload", value)

	testCases := map[string]string{
		"source_reference":       "",
		"source_payload":         "input payload",
		"input":                  "input payload",
		"previous":               "input payload",
		"run_dir":                "/tmp/run",
		"run_id":                 "run-123",
		"invocation_dir":         "/tmp/invocation",
		"stage_id":               "render",
		"artifact:generate:note": "/tmp/run/note.md",
	}
	for key, expected := range testCases {
		resolved, resolveErr := resolveRuntimePlaceholder(key, runtimeCtx)
		require.NoError(t, resolveErr, key)
		require.Equal(t, expected, resolved, key)
	}

	_, err = resolveRuntimePlaceholder("artifact:generate", runtimeCtx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid runtime placeholder")

	_, err = resolveRuntimePlaceholder("artifact:missing:note", runtimeCtx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown stage")

	_, err = resolveRuntimePlaceholder("artifact:generate:missing", runtimeCtx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown artifact")

	_, err = resolveRuntimePlaceholder("unknown", runtimeCtx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown runtime placeholder")
}

func TestFindLastValidateStageIndex(t *testing.T) {
	t.Parallel()

	stages := []Stage{
		{ID: "render", Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "passthrough"}},
		{ID: "validate-one", Role: StageRoleValidate, Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "noop"}},
		{ID: "publish", Role: StageRolePublish, Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "write_publish_manifest"}},
		{ID: "validate-two", Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "validate_declared_outputs"}},
	}

	require.Equal(t, 3, findLastValidateStageIndex(stages))
	require.Equal(t, -1, findLastValidateStageIndex([]Stage{{ID: "render", Executor: ExecutorBuiltin, Builtin: &BuiltinConfig{Name: "passthrough"}}}))
}

func TestResolveStageInput(t *testing.T) {
	t.Parallel()

	stages := []Stage{
		{ID: "prepare"},
		{ID: "render"},
	}
	stagePayloads := map[string]string{"prepare": "previous output"}

	value, err := resolveStageInput(Stage{ID: "prepare"}, 0, "source payload", stages, map[string]string{}, nil)
	require.NoError(t, err)
	require.Equal(t, "source payload", value)

	value, err = resolveStageInput(Stage{ID: "render"}, 1, "source payload", stages, stagePayloads, nil)
	require.NoError(t, err)
	require.Equal(t, "previous output", value)

	value, err = resolveStageInput(Stage{
		ID:    "explicit-source",
		Input: &StageInput{From: StageInputSource},
	}, 1, "source payload", stages, stagePayloads, nil)
	require.NoError(t, err)
	require.Equal(t, "source payload", value)

	tempDir := t.TempDir()
	artifactPath := filepath.Join(tempDir, "note.md")
	require.NoError(t, os.WriteFile(artifactPath, []byte("artifact content"), 0o644))

	value, err = resolveStageInput(Stage{
		ID: "from-artifact",
		Input: &StageInput{
			From:     StageInputArtifact,
			Stage:    "prepare",
			Artifact: "note",
		},
	}, 1, "source payload", stages, stagePayloads, map[string]map[string]string{
		"prepare": {"note": artifactPath},
	})
	require.NoError(t, err)
	require.Equal(t, "artifact content", value)

	_, err = resolveStageInput(Stage{ID: "render"}, 1, "source payload", stages, map[string]string{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "requires previous stage")

	_, err = resolveStageInput(Stage{
		ID: "from-missing-artifact-stage",
		Input: &StageInput{
			From:     StageInputArtifact,
			Stage:    "missing",
			Artifact: "note",
		},
	}, 1, "source payload", stages, stagePayloads, map[string]map[string]string{})
	require.Error(t, err)
	require.Contains(t, err.Error(), `stage "missing" artifacts are not available`)

	_, err = resolveStageInput(Stage{
		ID: "from-missing-artifact",
		Input: &StageInput{
			From:     StageInputArtifact,
			Stage:    "prepare",
			Artifact: "missing",
		},
	}, 1, "source payload", stages, stagePayloads, map[string]map[string]string{"prepare": {}})
	require.Error(t, err)
	require.Contains(t, err.Error(), `artifact "missing" is not available`)

	require.NoError(t, os.Remove(artifactPath))
	_, err = resolveStageInput(Stage{
		ID: "from-deleted-artifact",
		Input: &StageInput{
			From:     StageInputArtifact,
			Stage:    "prepare",
			Artifact: "note",
		},
	}, 1, "source payload", stages, stagePayloads, map[string]map[string]string{
		"prepare": {"note": artifactPath},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "read artifact")

	_, err = resolveStageInput(Stage{
		ID:    "bad",
		Input: &StageInput{From: "elsewhere"},
	}, 1, "source payload", stages, stagePayloads, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported input.from")
}

func TestEmitFinalOutput(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	runner := NewRunner(&stdout, &bytes.Buffer{}, nil)

	runner.emitFinalOutput("")
	require.Empty(t, stdout.String())

	runner.emitFinalOutput("hello world\n\n")
	require.Equal(t, "hello world\n", stdout.String())
}

func TestResolveStageOutputs(t *testing.T) {
	t.Parallel()

	runDir := t.TempDir()
	stage := Stage{
		ID: "render",
		Artifacts: []ArtifactDeclaration{
			{Name: "note", Path: "note.md"},
			{Name: "optional", Path: "optional.md", Required: boolPtr(false)},
		},
		PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputArtifact, Artifact: "note"},
	}

	artifactPaths, writtenFiles, primaryPayload, err := resolveStageOutputs(stage, runDir, "note body", nil)
	require.NoError(t, err)
	require.Equal(t, "note body", primaryPayload)
	require.Equal(t, filepath.Join(runDir, "note.md"), artifactPaths["note"])
	require.NotContains(t, artifactPaths, "optional")
	require.Len(t, writtenFiles, 1)
	require.FileExists(t, filepath.Join(runDir, "note.md"))

	_, _, _, err = resolveStageOutputs(Stage{
		ID: "missing-required",
		Artifacts: []ArtifactDeclaration{
			{Name: "note", Path: "missing.md"},
		},
		PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputStdout},
	}, runDir, "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), `required artifact "note" was not produced`)

	_, _, _, err = resolveStageOutputs(Stage{
		ID: "missing-primary",
		Artifacts: []ArtifactDeclaration{
			{Name: "note", Path: "missing-primary.md", Required: boolPtr(false)},
		},
		PrimaryOutput: &PrimaryOutputConfig{From: PrimaryOutputArtifact, Artifact: "note"},
	}, runDir, "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), `primary artifact "note" was not produced`)
}

func newDryRunRegistryForPatternTests(t *testing.T) *core.PluginRegistry {
	t.Helper()

	db := fsdb.NewDb(t.TempDir())
	require.NoError(t, db.Patterns.Configure())
	require.NoError(t, db.Contexts.Configure())
	require.NoError(t, db.Sessions.Configure())

	vendorManager := ai.NewVendorsManager()
	vendorManager.AddVendors(dryrun.NewClient())

	return &core.PluginRegistry{
		Db:            db,
		VendorManager: vendorManager,
		Defaults: &tools.Defaults{
			PluginBase: &plugins.PluginBase{Name: "Defaults"},
			Vendor:     &plugins.Setting{Value: "DryRun"},
			Model: &plugins.SetupQuestion{
				Setting: &plugins.Setting{Value: "dry-run-model"},
			},
			ModelContextLength: &plugins.SetupQuestion{
				Setting: &plugins.Setting{Value: "0"},
			},
		},
	}
}

func writeNamedPattern(t *testing.T, db *fsdb.Db, name string, content string) {
	t.Helper()

	patternDir := filepath.Join(db.Patterns.Dir, name)
	require.NoError(t, os.MkdirAll(patternDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(patternDir, db.Patterns.SystemPatternFile), []byte(content), 0o644))
}

func writeNamedContext(t *testing.T, db *fsdb.Db, name string, content string) {
	t.Helper()

	require.NoError(t, db.Contexts.Save(name, []byte(content)))
}

func boolPtr(v bool) *bool {
	return &v
}
