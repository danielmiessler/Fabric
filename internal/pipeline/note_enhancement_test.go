package pipeline

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const noteEnhancementSample = `
Meeting focused on migration to the new analytics pipeline.
Team agreed to split the rollout by tenant tier.
Main risk is data drift during dual-write window.
`

type noteEnhancementFixture struct {
	RepoRoot      string
	SessionDir    string
	SourcePath    string
	InvocationDir string
}

type noteEnhancementSessionContext struct {
	SourceMode        string `json:"source_mode"`
	SourceReference   string `json:"source_reference"`
	InvocationDir     string `json:"invocation_dir"`
	SessionDir        string `json:"session_dir"`
	InputPath         string `json:"input_path"`
	SessionName       string `json:"session_name"`
	SessionID         string `json:"session_id"`
	StablePipelineDir string `json:"stable_pipeline_dir"`
	FinalOutputPath   string `json:"final_output_path"`
}

func TestNoteEnhancementPipelineLoadsAndValidates(t *testing.T) {
	t.Parallel()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("note-enhancement")
	require.NoError(t, err)
	require.NoError(t, Validate(pipe))
}

func TestNoteEnhancementPrepareSourceStdinUsesRunDirArtifacts(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "note-enhancement", "prepare_source.py"),
		nil,
		strings.TrimSpace(noteEnhancementSample)+"\n",
		map[string]string{
			"FABRIC_PIPELINE_RUN_DIR":        runDir,
			"FABRIC_PIPELINE_INVOCATION_DIR": invocationDir,
			"FABRIC_PIPELINE_SOURCE_MODE":    "stdin",
		},
	)

	inputPath := strings.TrimSpace(prepareStdout)
	require.Equal(t, canonicalPath(t, filepath.Join(runDir, "source_input.md")), canonicalPath(t, inputPath))
	require.FileExists(t, inputPath)
	require.NoFileExists(t, filepath.Join(invocationDir, "source_input.md"))

	sessionContextPath := filepath.Join(runDir, "session_context.json")
	require.FileExists(t, sessionContextPath)

	var sessionContext noteEnhancementSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, invocationDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, canonicalPath(t, runDir), canonicalPath(t, sessionContext.StablePipelineDir))
	require.Equal(
		t,
		filepath.Join(sessionContext.StablePipelineDir, "enhanced_notes.md"),
		sessionContext.FinalOutputPath,
	)
}

func TestNoteEnhancementRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newNoteEnhancementFixture(t)
	pipe, result, stdout, stderr := runNoteEnhancementWithPatternStub(
		t,
		fixture,
		RunOptions{
			InvocationDir:  fixture.InvocationDir,
			DisableCleanup: true,
		},
	)

	require.Contains(t, stdout, "# ✨ Enhanced Notes - Analytics Migration Rollout")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[1/5] prepare_source ........ PASS")
	require.Contains(t, stderr, "[5/5] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertNoteEnhancementManifest(t, result.RunDir, pipe)
	assertNoteEnhancementArtifacts(t, fixture, result.RunDir)
}

func newNoteEnhancementFixture(t *testing.T) noteEnhancementFixture {
	t.Helper()

	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "Migration Session")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))
	sourcePath := filepath.Join(sessionDir, "notes.md")
	require.NoError(t, os.WriteFile(sourcePath, []byte(strings.TrimSpace(noteEnhancementSample)+"\n"), 0o644))

	return noteEnhancementFixture{
		RepoRoot:      repoRootFromTest(t),
		SessionDir:    sessionDir,
		SourcePath:    sourcePath,
		InvocationDir: t.TempDir(),
	}
}

func runNoteEnhancementWithPatternStub(t *testing.T, fixture noteEnhancementFixture, opts RunOptions) (*Pipeline, *RunResult, string, string) {
	t.Helper()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("note-enhancement")
	require.NoError(t, err)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	runner.stageExecutionOverride = noteEnhancementPatternOverride()

	result, err := runner.Run(context.Background(), pipe, RunSource{
		Mode:      SourceModeSource,
		Reference: fixture.SessionDir,
	}, opts)
	require.NoErrorf(t, err, "stderr:\n%s\nstdout:\n%s", stderr.String(), stdout.String())
	require.NotNil(t, result)

	return pipe, result, stdout.String(), stderr.String()
}

func noteEnhancementPatternOverride() func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error) {
	return func(_ context.Context, stage Stage, _ StageRuntimeContext) (*StageExecutionResult, bool, error) {
		switch stage.ID {
		case "enhance_generate":
			return &StageExecutionResult{Stdout: buildNoteEnhancementOutput()}, true, nil
		default:
			return nil, false, nil
		}
	}
}

func buildNoteEnhancementOutput() string {
	enhanced := strings.Join([]string{
		"# ✨ Enhanced Notes - Analytics Migration Rollout",
		"",
		"## Improved Summary",
		"The team decided on a tiered rollout for analytics migration with explicit guardrails for drift detection.",
		"",
		"## Key Takeaways",
		"- Rollout order will follow tenant risk tier.",
		"- Dual-write period needs explicit reconciliation checks.",
		"- Drift metrics must be reviewed before each cohort expansion.",
		"",
		"## Clarifications Added",
		"- Identified that drift risk is highest during schema transition.",
		"- Clarified ownership for monitoring and rollback triggers.",
		"",
		"## Suggested Next Questions",
		"- Which drift threshold triggers automatic rollback?",
		"- What is the maximum dual-write window per tier?",
	}, "\n")
	editLog := strings.Join([]string{
		"# Edit Log",
		"",
		"## Structural Changes",
		"- Reorganized free-form notes into sectioned output.",
		"",
		"## Language Tightening",
		"- Converted vague statements into explicit operational wording.",
		"",
		"## Assumptions Avoided",
		"- Did not introduce migration dates not present in source.",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:enhanced_notes.md>>>",
		enhanced,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:edit_log.md>>>",
		editLog,
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func assertNoteEnhancementManifest(t *testing.T, runDir string, pipe *Pipeline) {
	t.Helper()

	var manifest RunManifest
	readJSONFile(t, filepath.Join(runDir, "run_manifest.json"), &manifest)
	require.Equal(t, pipe.Name, manifest.PipelineName)
	require.Equal(t, "passed", manifest.Status)
	require.Len(t, manifest.Stages, len(pipe.Stages))
}

func assertNoteEnhancementArtifacts(t *testing.T, fixture noteEnhancementFixture, runDir string) {
	t.Helper()

	require.FileExists(t, filepath.Join(runDir, "run_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "source_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "publish_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "session_context.json"))
	require.FileExists(t, filepath.Join(runDir, "enhancement_input.md"))
	require.FileExists(t, filepath.Join(runDir, "enhanced_notes.md"))
	require.FileExists(t, filepath.Join(runDir, "edit_log.md"))
	require.FileExists(t, filepath.Join(runDir, "validation_report.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "edit_log.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "validation_report.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, "enhanced_notes.md"))
}
