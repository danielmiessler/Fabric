package pipeline

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const therapyConversationSample = `
Therapist: Last week you said work stress made sleep difficult.
Client: Yes, I wake up at 3am and start replaying mistakes.
Therapist: What do you notice in your body when that happens?
Client: Tight chest and shallow breathing.
Therapist: Let's map small grounding steps you can use in that moment.
`

type therapyConversationFixture struct {
	RepoRoot      string
	SessionDir    string
	SourcePath    string
	InvocationDir string
}

type therapyConversationSessionContext struct {
	SourceMode        string `json:"source_mode"`
	SourceReference   string `json:"source_reference"`
	InvocationDir     string `json:"invocation_dir"`
	SessionDir        string `json:"session_dir"`
	InputPath         string `json:"input_path"`
	SessionName       string `json:"session_name"`
	SessionID         string `json:"session_id"`
	StablePipelineDir string `json:"stable_pipeline_dir"`
	FinalOutputPath   string `json:"final_output_path"`
	ContextRoot       string `json:"context_root"`
}

func TestTherapyConversationNotesPipelineLoadsAndValidates(t *testing.T) {
	t.Parallel()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("therapy-conversation-notes")
	require.NoError(t, err)
	require.NoError(t, Validate(pipe))
}

func TestTherapyConversationNotesPrepareSourceStdinUsesRunDirArtifacts(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "therapy-conversation-notes", "prepare_source.py"),
		nil,
		strings.TrimSpace(therapyConversationSample)+"\n",
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

	var sessionContext therapyConversationSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, invocationDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, canonicalPath(t, runDir), canonicalPath(t, sessionContext.StablePipelineDir))
	require.Equal(
		t,
		filepath.Join(sessionContext.StablePipelineDir, "therapy_conversation_notes.md"),
		sessionContext.FinalOutputPath,
	)
}

func TestTherapyConversationNotesRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newTherapyConversationFixture(t)
	pipe, result, stdout, stderr := runTherapyConversationWithPatternStub(
		t,
		fixture,
		RunOptions{
			InvocationDir:  fixture.InvocationDir,
			DisableCleanup: true,
		},
	)

	require.Contains(t, stdout, "# 🧠 Therapy Conversation Notes - Sleep and Stress Loop")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[1/7] prepare_source ........ PASS")
	require.Contains(t, stderr, "[7/7] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertTherapyConversationManifest(t, result.RunDir, pipe)
	assertTherapyConversationArtifacts(t, fixture, result.RunDir)
}

func newTherapyConversationFixture(t *testing.T) therapyConversationFixture {
	t.Helper()

	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "Therapy Session 12")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))

	sourcePath := filepath.Join(sessionDir, "transcript.md")
	require.NoError(t, os.WriteFile(sourcePath, []byte(strings.TrimSpace(therapyConversationSample)+"\n"), 0o644))

	contextDir := filepath.Join(sessionDir, "context")
	require.NoError(t, os.MkdirAll(contextDir, 0o755))
	contextPath := filepath.Join(contextDir, "grounding_protocol.md")
	contextText := `
# Grounding Protocol

Use body-first reset:
- Name 5 things you can see.
- Slow exhale for 6 counts.
- Write one concrete next action.
`
	require.NoError(t, os.WriteFile(contextPath, []byte(strings.TrimSpace(contextText)+"\n"), 0o644))

	return therapyConversationFixture{
		RepoRoot:      repoRootFromTest(t),
		SessionDir:    sessionDir,
		SourcePath:    sourcePath,
		InvocationDir: t.TempDir(),
	}
}

func runTherapyConversationWithPatternStub(t *testing.T, fixture therapyConversationFixture, opts RunOptions) (*Pipeline, *RunResult, string, string) {
	t.Helper()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("therapy-conversation-notes")
	require.NoError(t, err)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	runner.stageExecutionOverride = therapyConversationPatternOverride(t)

	result, err := runner.Run(context.Background(), pipe, RunSource{
		Mode:      SourceModeSource,
		Reference: fixture.SessionDir,
	}, opts)
	require.NoErrorf(t, err, "stderr:\n%s\nstdout:\n%s", stderr.String(), stdout.String())
	require.NotNil(t, result)

	return pipe, result, stdout.String(), stderr.String()
}

func therapyConversationPatternOverride(t *testing.T) func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error) {
	t.Helper()

	return func(_ context.Context, stage Stage, _ StageRuntimeContext) (*StageExecutionResult, bool, error) {
		switch stage.ID {
		case "analyze_generate":
			return &StageExecutionResult{Stdout: buildTherapyAnalyzeOutput(t)}, true, nil
		default:
			return nil, false, nil
		}
	}
}

func buildTherapyAnalyzeOutput(t *testing.T) string {
	t.Helper()

	snapshot := map[string]any{
		"summary_points":    []string{"Night-time rumination follows work stress", "Body activation appears as chest tightness"},
		"emotions_observed": []string{"anxiety", "self-criticism", "overwhelm"},
		"patterns_observed": []string{"catastrophic replay", "sleep disruption loop"},
		"actions":           []string{"use breath reset", "write one concrete next step", "set a short worry window earlier in the evening"},
		"safety_flags":      []string{},
	}
	content, err := json.MarshalIndent(snapshot, "", "  ")
	require.NoError(t, err)

	finalNotes := strings.Join([]string{
		"# 🧠 Therapy Conversation Notes - Sleep and Stress Loop",
		"",
		"## Conversation Summary",
		"The conversation focused on recurring night-time wakeups linked to work-stress rumination.",
		"",
		"## Emotional Signals",
		"- Anxiety around perceived mistakes",
		"- Fear of repeating errors",
		"- Fatigue from fragmented sleep",
		"",
		"## Thought Patterns",
		"- Replaying worst-case outcomes at 3am",
		"- Treating one mistake as global failure",
		"- Difficulty shifting attention once activated",
		"",
		"## Actionable Reflection",
		"- Use a 60-second exhale-focused breathing reset when waking.",
		"- Write one realistic action for the next morning instead of solving everything at night.",
		"- Identify one thought distortion and restate it in neutral language.",
		"",
		"## Safety and Boundaries",
		"These notes are for reflection only and are not a substitute for professional care.",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:final_notes.md>>>",
		finalNotes,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:analysis_snapshot.json>>>",
		string(content),
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func assertTherapyConversationManifest(t *testing.T, runDir string, pipe *Pipeline) {
	t.Helper()

	var manifest RunManifest
	readJSONFile(t, filepath.Join(runDir, "run_manifest.json"), &manifest)
	require.Equal(t, pipe.Name, manifest.PipelineName)
	require.Equal(t, "passed", manifest.Status)
	require.Len(t, manifest.Stages, len(pipe.Stages))
}

func assertTherapyConversationArtifacts(t *testing.T, fixture therapyConversationFixture, runDir string) {
	t.Helper()

	require.FileExists(t, filepath.Join(runDir, "run_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "source_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "publish_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "session_context.json"))
	require.FileExists(t, filepath.Join(runDir, "conversation_input.md"))
	require.FileExists(t, filepath.Join(runDir, "context_pack.md"))
	require.FileExists(t, filepath.Join(runDir, "context_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "analysis_input.md"))
	require.FileExists(t, filepath.Join(runDir, "analysis_snapshot.json"))
	require.FileExists(t, filepath.Join(runDir, "validation_report.md"))
	require.FileExists(t, filepath.Join(runDir, "final_notes.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "context_pack.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "context_manifest.json"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "analysis_snapshot.json"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "validation_report.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, "therapy_conversation_notes.md"))
}
