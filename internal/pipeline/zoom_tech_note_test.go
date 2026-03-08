package pipeline

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const zoomTechNoteSampleCaption = `
[Harkirat Singh] 00:00:01
Today we will talk about neural networks from scratch.

[Student] 00:00:05
Okay.

[Harkirat Singh] 00:00:07
A neuron combines inputs, weights, and bias before activation.
`

type testSessionContext struct {
	SourceMode      string `json:"source_mode"`
	SourceReference string `json:"source_reference"`
	InvocationDir   string `json:"invocation_dir"`
	SessionDir      string `json:"session_dir"`
	InputPath       string `json:"input_path"`
	SessionName     string `json:"session_name"`
	SessionID       string `json:"session_id"`
}

type zoomTechNoteFixture struct {
	RepoRoot      string
	SessionDir    string
	CaptionPath   string
	InvocationDir string
	PublishedPath string
}

func TestMain(m *testing.M) {
	if handled, err := CleanupRunDirFromEnv(); handled {
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestZoomTechNotePipelinesLoadAndValidate(t *testing.T) {
	t.Parallel()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	for _, name := range []string{"zoom-tech-note", "zoom-tech-note-deep-pass"} {
		pipe, err := loader.LoadNamed(name)
		require.NoError(t, err, name)
		require.NoError(t, Validate(pipe), name)
	}
}

func TestZoomTechNoteDeterministicWrappers(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "2026-01-18 10.00.00 Fast-tracking the AI course")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))
	captionPath := filepath.Join(sessionDir, "meeting_saved_closed_caption.txt")
	require.NoError(t, os.WriteFile(captionPath, []byte(strings.TrimSpace(`
[Harkirat Singh] 00:00:01
Today we will talk about neural networks from scratch.

[Student] 00:00:05
Okay.

[Harkirat Singh] 00:00:07
A neuron combines inputs, weights, and bias before activation.
`)+"\n"), 0o644))

	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "prepare_source.py"),
		nil,
		sessionDir,
		map[string]string{
			"FABRIC_PIPELINE_RUN_DIR":          runDir,
			"FABRIC_PIPELINE_INVOCATION_DIR":   invocationDir,
			"FABRIC_PIPELINE_SOURCE_MODE":      "source",
			"FABRIC_PIPELINE_SOURCE_REFERENCE": sessionDir,
		},
	)
	require.Equal(t, canonicalPath(t, captionPath), canonicalPath(t, strings.TrimSpace(prepareStdout)))

	sessionContextPath := filepath.Join(runDir, "session_context.json")
	stage1InputPath := filepath.Join(runDir, "stage1_input.md")
	require.FileExists(t, sessionContextPath)
	require.FileExists(t, stage1InputPath)

	var sessionContext testSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, sessionDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, "2026-01-18 10.00.00 Fast-tracking the AI course", sessionContext.SessionName)
	require.Equal(t, "2026-01-18-10-00-00-fast-tracking-the-ai-course", sessionContext.SessionID)

	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "run_ingest_stage.py"),
		[]string{"--session-context", sessionContextPath},
		captionPath+"\n",
		nil,
	)
	ledgerPath := filepath.Join(sessionDir, ".pipeline", "segment_ledger.jsonl")
	manifestPath := filepath.Join(sessionDir, ".pipeline", "segment_manifest.jsonl")
	require.FileExists(t, ledgerPath)
	require.FileExists(t, manifestPath)

	ledgerRows := readJSONLines(t, ledgerPath)
	require.Len(t, ledgerRows, 3)
	contentIDs := []string{
		ledgerRows[0]["segment_id"].(string),
		ledgerRows[2]["segment_id"].(string),
	}
	noiseID := ledgerRows[1]["segment_id"].(string)

	stage1Output := strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/refined_transcript.md>>>",
		"---",
		"[source: " + contentIDs[0] + "] Today we will talk about neural networks from scratch.",
		"---",
		"[source: " + contentIDs[1] + "] A neuron combines inputs, weights, and bias before activation.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/topic_inventory.json>>>",
		`{"concepts":["neural networks","neurons"],"technical_terms":["weights","bias","activation"],"code_or_commands":[],"qa_items":[],"named_entities":[]}`,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/corrections_log.csv>>>",
		"segment_id,raw_text,corrected_text,confidence_tier,confidence_score,reasoning",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/uncertainty_report.json>>>",
		"[]",
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "materialize_stage1.py"),
		[]string{"--session-context", sessionContextPath},
		stage1Output,
		map[string]string{"FABRIC_PIPELINE_RUN_DIR": runDir},
	)
	require.FileExists(t, filepath.Join(sessionDir, ".pipeline", "refined_transcript.md"))
	require.FileExists(t, filepath.Join(runDir, "stage2_input.md"))

	coverage := map[string]any{
		contentIDs[0]: map[string]any{"sections": []string{"S1"}, "status": "covered"},
		contentIDs[1]: map[string]any{"sections": []string{"S2"}, "status": "covered"},
		noiseID:       map[string]any{"sections": []string{}, "status": "noise"},
	}
	coverageJSON, err := json.MarshalIndent(coverage, "", "  ")
	require.NoError(t, err)

	stage2Output := strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/structured_notes.md>>>",
		"## Neural Networks",
		"[source: " + contentIDs[0] + "] Neural networks are computational systems built from simple neurons.",
		"## Neurons",
		"[source: " + contentIDs[1] + "] A neuron combines inputs, weights, and bias before activation.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/coverage_matrix.json>>>",
		string(coverageJSON),
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "materialize_stage2.py"),
		[]string{"--session-context", sessionContextPath},
		stage2Output,
		map[string]string{"FABRIC_PIPELINE_RUN_DIR": runDir},
	)
	require.FileExists(t, filepath.Join(sessionDir, ".pipeline", "structured_notes.md"))
	require.FileExists(t, filepath.Join(runDir, "stage3_input.md"))

	finalNotes := strings.Join([]string{
		"---",
		`title: "AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI"`,
		"---",
		"",
		"# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI",
		"",
		"## 🧠 Session Focus",
		"Build an intuition-first view of neural networks.",
		"",
		"## 🎯 Prerequisites",
		"- Basic algebra",
		"- Comfort reading simple functions",
		"",
		"## ✅ Learning Outcomes",
		"- Explain what a neuron is",
		"- Explain why weights and bias matter",
		"",
		"## 🧭 Topic Index",
		"- Neural networks",
		"- Neurons",
		"",
		"## 🗺️ Conceptual Roadmap",
		"```mermaid",
		"flowchart TD",
		`  A["Inputs"] --> B["Weights and Bias"]`,
		`  B --> C["Activation"]`,
		"```",
		"",
		"## 🏗️ Systems Visualization",
		"```mermaid",
		"graph LR",
		`  A["Input Layer"] --> B["Neuron"]`,
		`  B --> C["Output"]`,
		"```",
		"",
		"## 🌆 Skyline Intuition Diagram",
		"Inputs -> Weighted Sum -> Activation -> Output",
		"",
		"## 📚 Core Concepts (Intuition First)",
		"The core intuition is that a neuron learns a weighted response. This intuition matters because it explains how the model bends output. A second intuition is that bias shifts the response. A third intuition is that activation changes the shape of the decision rule.",
		"",
		"## ➗ Mathematical Intuition",
		"z = wx + b. The symbols describe a weighted input plus a shift.",
		"",
		"## 💻 Coding Walkthroughs",
		"- Beginner example: compute a weighted sum",
		"- Advanced example: add an activation function",
		"",
		"## 🚀 Advanced Real-World Scenario",
		"Use a neuron to score whether a feature should activate.",
		"",
		"## 🧩 HOTS (High-Order Thinking)",
		"1. Why can bias matter even if weights are correct?",
		"2. What would happen if activation never changed?",
		"",
		"## ❓ FAQ",
		"Q: Why do we need weights?",
		"A: They scale contribution.",
		"",
		"Q: Why do we need bias?",
		"A: It shifts the response.",
		"",
		"## 🛠️ Practice Roadmap",
		"1. Compute a weighted sum by hand.",
		"2. Write a small neuron in code.",
		"",
		"## 🔭 Next Improvements",
		"- Add multiple neurons",
		"",
		"## 🔗 Related Notes",
		"- Activation functions",
		"",
		"## 🧾 Traceability",
		"- See `.pipeline/coverage_matrix.json`",
	}, "\n")
	stage3Output := strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/enhanced_notes.md>>>",
		"[ENHANCED: Added intuition builders.]\n\n## Neural Networks\n[source: " + contentIDs[0] + "] Neural networks are intuitive weighted systems.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:final_notes.md>>>",
		finalNotes,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:bootcamp_index.md>>>",
		"# Bootcamp Index\n\n- [[AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI]]",
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
	stage3Stdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "materialize_stage3.py"),
		[]string{"--session-context", sessionContextPath},
		stage3Output,
		nil,
	)
	require.Contains(t, stage3Stdout, "# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI")
	require.FileExists(t, filepath.Join(sessionDir, ".pipeline", "enhanced_notes.md"))
	require.FileExists(t, filepath.Join(sessionDir, "final_notes.md"))
	require.FileExists(t, filepath.Join(sessionDir, "bootcamp_index.md"))

	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "deep_pass.py"),
		[]string{"--session-context", sessionContextPath},
		"",
		nil,
	)
	require.FileExists(t, filepath.Join(sessionDir, ".pipeline", "deep_pass_report.md"))

	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "run_validate_stage.py"),
		[]string{"--session-context", sessionContextPath},
		"",
		nil,
	)
	require.FileExists(t, filepath.Join(sessionDir, ".pipeline", "validation_report.md"))

	siblingSessionDir := filepath.Join(rootDir, "2026-01-10 09.00.00 Neural Networks from Scratch")
	require.NoError(t, os.MkdirAll(siblingSessionDir, 0o755))
	require.NoError(
		t,
		os.WriteFile(
			filepath.Join(siblingSessionDir, "final_notes.md"),
			[]byte("# sibling session\n"),
			0o644,
		),
	)

	runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "run_publish_stage.py"),
		[]string{"--session-context", sessionContextPath},
		"",
		nil,
	)
	publishedPath := filepath.Join(sessionDir, "AI-ML Class 01 [18-01-2026] - Fast-tracking the Course of AI.md")
	require.FileExists(t, publishedPath)
	publishedBody, err := os.ReadFile(publishedPath)
	require.NoError(t, err)
	require.Contains(t, string(publishedBody), "# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI")
}

func TestZoomPrepareSourceStdinUsesRunDirArtifacts(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "zoom-tech-note", "prepare_source.py"),
		nil,
		strings.TrimSpace(zoomTechNoteSampleCaption)+"\n",
		map[string]string{
			"FABRIC_PIPELINE_RUN_DIR":        runDir,
			"FABRIC_PIPELINE_INVOCATION_DIR": invocationDir,
			"FABRIC_PIPELINE_SOURCE_MODE":    "stdin",
		},
	)

	inputPath := strings.TrimSpace(prepareStdout)
	require.Equal(t, canonicalPath(t, filepath.Join(runDir, "meeting_saved_closed_caption.txt")), canonicalPath(t, inputPath))
	require.FileExists(t, inputPath)
	require.NoFileExists(t, filepath.Join(invocationDir, "meeting_saved_closed_caption.txt"))

	sessionContextPath := filepath.Join(runDir, "session_context.json")
	require.FileExists(t, sessionContextPath)

	var sessionContext testSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, invocationDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, canonicalPath(t, inputPath), canonicalPath(t, sessionContext.InputPath))
}

func TestZoomTechNoteRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newZoomTechNoteFixture(t)
	pipe, result, stdout, stderr := runZoomPipelineWithPatternStubs(
		t,
		"zoom-tech-note",
		fixture,
		RunOptions{
			InvocationDir:  fixture.InvocationDir,
			DisableCleanup: true,
		},
	)

	require.Contains(t, stdout, "# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[1/11] prepare_source ........ PASS")
	require.Contains(t, stderr, "[11/11] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertZoomRunManifest(t, result.RunDir, pipe)
	assertZoomCommonArtifacts(t, fixture, result.RunDir)
	require.NoFileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "deep_pass_report.md"))
}

func TestZoomTechNoteDeepPassRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newZoomTechNoteFixture(t)
	pipe, result, stdout, stderr := runZoomPipelineWithPatternStubs(
		t,
		"zoom-tech-note-deep-pass",
		fixture,
		RunOptions{
			InvocationDir: fixture.InvocationDir,
			CleanupDelay:  200 * time.Millisecond,
		},
	)

	require.Contains(t, stdout, "# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[9/12] deep_pass ........ PASS")
	require.Contains(t, stderr, "[12/12] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertZoomRunManifest(t, result.RunDir, pipe)
	assertZoomCommonArtifacts(t, fixture, result.RunDir)
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "deep_pass_report.md"))

	require.Eventually(t, func() bool {
		_, err := os.Stat(result.RunDir)
		return os.IsNotExist(err)
	}, 5*time.Second, 50*time.Millisecond)
	_, err := os.Stat(filepath.Join(fixture.InvocationDir, ".pipeline"))
	require.True(t, os.IsNotExist(err), "expected cleanup to remove .pipeline directory, got err=%v", err)
}

func repoRootFromTest(t *testing.T) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func runPythonScript(t *testing.T, repoRoot string, scriptPath string, args []string, stdin string, extraEnv map[string]string) string {
	t.Helper()

	cmdArgs := append([]string{scriptPath}, args...)
	cmd := exec.Command("python3", cmdArgs...)
	cmd.Dir = repoRoot
	cmd.Stdin = strings.NewReader(stdin)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()
	for key, value := range extraEnv {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	err := cmd.Run()
	require.NoErrorf(t, err, "script failed: %s\nstdout:\n%s\nstderr:\n%s", scriptPath, stdout.String(), stderr.String())
	return stdout.String()
}

func readJSONFile(t *testing.T, path string, target any) {
	t.Helper()

	content, err := os.ReadFile(path)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(content, target))
}

func readJSONLines(t *testing.T, path string) []map[string]any {
	t.Helper()

	content, err := os.ReadFile(path)
	require.NoError(t, err)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	rows := make([]map[string]any, 0, len(lines))
	for _, line := range lines {
		var row map[string]any
		require.NoError(t, json.Unmarshal([]byte(line), &row))
		rows = append(rows, row)
	}
	return rows
}

func canonicalPath(t *testing.T, path string) string {
	t.Helper()

	evaluated, err := filepath.EvalSymlinks(path)
	if err == nil {
		return evaluated
	}
	absPath, absErr := filepath.Abs(path)
	require.NoError(t, absErr)
	return absPath
}

func newZoomTechNoteFixture(t *testing.T) zoomTechNoteFixture {
	t.Helper()

	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "2026-01-18 10.00.00 Fast-tracking the AI course")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))
	captionPath := filepath.Join(sessionDir, "meeting_saved_closed_caption.txt")
	require.NoError(t, os.WriteFile(captionPath, []byte(strings.TrimSpace(zoomTechNoteSampleCaption)+"\n"), 0o644))

	return zoomTechNoteFixture{
		RepoRoot:      repoRootFromTest(t),
		SessionDir:    sessionDir,
		CaptionPath:   captionPath,
		InvocationDir: t.TempDir(),
		PublishedPath: filepath.Join(sessionDir, "AI-ML Class 01 [18-01-2026] - Fast-tracking the Course of AI.md"),
	}
}

func runZoomPipelineWithPatternStubs(t *testing.T, pipelineName string, fixture zoomTechNoteFixture, opts RunOptions) (*Pipeline, *RunResult, string, string) {
	t.Helper()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed(pipelineName)
	require.NoError(t, err)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	runner.stageExecutionOverride = zoomPatternStageOverride(t)

	result, err := runner.Run(context.Background(), pipe, RunSource{
		Mode:      SourceModeSource,
		Reference: fixture.SessionDir,
	}, opts)
	require.NoErrorf(t, err, "stderr:\n%s\nstdout:\n%s", stderr.String(), stdout.String())
	require.NotNil(t, result)

	return pipe, result, stdout.String(), stderr.String()
}

func zoomPatternStageOverride(t *testing.T) func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error) {
	t.Helper()

	return func(_ context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, bool, error) {
		switch stage.ID {
		case "stage1_refine_generate":
			sessionContext := loadZoomSessionContext(t, runtimeCtx)
			contentIDs, _ := loadZoomSegmentIDs(t, sessionContext.SessionDir)
			return &StageExecutionResult{Stdout: buildZoomStage1Output(contentIDs)}, true, nil
		case "stage2_synthesize_generate":
			sessionContext := loadZoomSessionContext(t, runtimeCtx)
			contentIDs, noiseID := loadZoomSegmentIDs(t, sessionContext.SessionDir)
			return &StageExecutionResult{Stdout: buildZoomStage2Output(t, contentIDs, noiseID)}, true, nil
		case "stage3_enhance_generate":
			sessionContext := loadZoomSessionContext(t, runtimeCtx)
			contentIDs, _ := loadZoomSegmentIDs(t, sessionContext.SessionDir)
			return &StageExecutionResult{Stdout: buildZoomStage3Output(contentIDs[0])}, true, nil
		default:
			return nil, false, nil
		}
	}
}

func loadZoomSessionContext(t *testing.T, runtimeCtx StageRuntimeContext) testSessionContext {
	t.Helper()

	sessionContextPath := filepath.Join(runtimeCtx.RunDir, "session_context.json")
	require.FileExists(t, sessionContextPath)

	var sessionContext testSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	return sessionContext
}

func loadZoomSegmentIDs(t *testing.T, sessionDir string) ([]string, string) {
	t.Helper()

	ledgerRows := readJSONLines(t, filepath.Join(sessionDir, ".pipeline", "segment_ledger.jsonl"))
	require.Len(t, ledgerRows, 3)

	return []string{
			ledgerRows[0]["segment_id"].(string),
			ledgerRows[2]["segment_id"].(string),
		},
		ledgerRows[1]["segment_id"].(string)
}

func buildZoomStage1Output(contentIDs []string) string {
	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/refined_transcript.md>>>",
		"---",
		"[source: " + contentIDs[0] + "] Today we will talk about neural networks from scratch.",
		"---",
		"[source: " + contentIDs[1] + "] A neuron combines inputs, weights, and bias before activation.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/topic_inventory.json>>>",
		`{"concepts":["neural networks","neurons"],"technical_terms":["weights","bias","activation"],"code_or_commands":[],"qa_items":[],"named_entities":[]}`,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/corrections_log.csv>>>",
		"segment_id,raw_text,corrected_text,confidence_tier,confidence_score,reasoning",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/uncertainty_report.json>>>",
		"[]",
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func buildZoomStage2Output(t *testing.T, contentIDs []string, noiseID string) string {
	t.Helper()

	coverage := map[string]any{
		contentIDs[0]: map[string]any{"sections": []string{"S1"}, "status": "covered"},
		contentIDs[1]: map[string]any{"sections": []string{"S2"}, "status": "covered"},
		noiseID:       map[string]any{"sections": []string{}, "status": "noise"},
	}
	coverageJSON, err := json.MarshalIndent(coverage, "", "  ")
	require.NoError(t, err)

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/structured_notes.md>>>",
		"## Neural Networks",
		"[source: " + contentIDs[0] + "] Neural networks are computational systems built from simple neurons.",
		"## Neurons",
		"[source: " + contentIDs[1] + "] A neuron combines inputs, weights, and bias before activation.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:.pipeline/coverage_matrix.json>>>",
		string(coverageJSON),
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func buildZoomStage3Output(firstContentID string) string {
	finalNotes := strings.Join([]string{
		"---",
		`title: "AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI"`,
		"---",
		"",
		"# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI",
		"",
		"## 🧠 Session Focus",
		"Build an intuition-first view of neural networks.",
		"",
		"## 🎯 Prerequisites",
		"- Basic algebra",
		"- Comfort reading simple functions",
		"",
		"## ✅ Learning Outcomes",
		"- Explain what a neuron is",
		"- Explain why weights and bias matter",
		"",
		"## 🧭 Topic Index",
		"- Neural networks",
		"- Neurons",
		"",
		"## 🗺️ Conceptual Roadmap",
		"```mermaid",
		"flowchart TD",
		`  A["Inputs"] --> B["Weights and Bias"]`,
		`  B --> C["Activation"]`,
		"```",
		"",
		"## 🏗️ Systems Visualization",
		"```mermaid",
		"graph LR",
		`  A["Input Layer"] --> B["Neuron"]`,
		`  B --> C["Output"]`,
		"```",
		"",
		"## 🌆 Skyline Intuition Diagram",
		"Inputs -> Weighted Sum -> Activation -> Output",
		"",
		"## 📚 Core Concepts (Intuition First)",
		"The core intuition is that a neuron learns a weighted response. This intuition matters because it explains how the model bends output. A second intuition is that bias shifts the response. A third intuition is that activation changes the shape of the decision rule.",
		"",
		"## ➗ Mathematical Intuition",
		"z = wx + b. The symbols describe a weighted input plus a shift.",
		"",
		"## 💻 Coding Walkthroughs",
		"- Beginner example: compute a weighted sum",
		"- Advanced example: add an activation function",
		"",
		"## 🚀 Advanced Real-World Scenario",
		"Use a neuron to score whether a feature should activate.",
		"",
		"## 🧩 HOTS (High-Order Thinking)",
		"1. Why can bias matter even if weights are correct?",
		"2. What would happen if activation never changed?",
		"",
		"## ❓ FAQ",
		"Q: Why do we need weights?",
		"A: They scale contribution.",
		"",
		"Q: Why do we need bias?",
		"A: It shifts the response.",
		"",
		"## 🛠️ Practice Roadmap",
		"1. Compute a weighted sum by hand.",
		"2. Write a small neuron in code.",
		"",
		"## 🔭 Next Improvements",
		"- Add multiple neurons",
		"",
		"## 🔗 Related Notes",
		"- Activation functions",
		"",
		"## 🧾 Traceability",
		"- See `.pipeline/coverage_matrix.json`",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:.pipeline/enhanced_notes.md>>>",
		"[ENHANCED: Added intuition builders.]\n\n## Neural Networks\n[source: " + firstContentID + "] Neural networks are intuitive weighted systems.",
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:final_notes.md>>>",
		finalNotes,
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:bootcamp_index.md>>>",
		"# Bootcamp Index\n\n- [[AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI]]",
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func assertZoomRunManifest(t *testing.T, runDir string, pipe *Pipeline) {
	t.Helper()

	var manifest RunManifest
	readJSONFile(t, filepath.Join(runDir, "run_manifest.json"), &manifest)
	require.Equal(t, "passed", manifest.Status)
	require.NotNil(t, manifest.FinalOutput)
	require.Equal(t, "stage3_enhance_materialize", manifest.FinalOutput.StageID)
	require.Len(t, manifest.Stages, len(pipe.Stages))
	for i, stage := range pipe.Stages {
		require.Equal(t, stage.ID, manifest.Stages[i].ID)
		require.Equal(t, "passed", manifest.Stages[i].Status)
	}

	var sourceManifest SourceManifest
	readJSONFile(t, filepath.Join(runDir, "source_manifest.json"), &sourceManifest)
	require.Equal(t, SourceModeSource, sourceManifest.Mode)

	var runState RunState
	readJSONFile(t, filepath.Join(runDir, "run.json"), &runState)
	require.Equal(t, "completed", runState.Status)
}

func assertZoomCommonArtifacts(t *testing.T, fixture zoomTechNoteFixture, runDir string) {
	t.Helper()

	for _, path := range []string{
		filepath.Join(runDir, "run_manifest.json"),
		filepath.Join(runDir, "run.json"),
		filepath.Join(runDir, "source_manifest.json"),
		filepath.Join(runDir, "session_context.json"),
		filepath.Join(runDir, "stage1_input.md"),
		filepath.Join(runDir, "stage2_input.md"),
		filepath.Join(runDir, "stage3_input.md"),
		filepath.Join(runDir, "publish_manifest.json"),
		filepath.Join(fixture.SessionDir, ".pipeline", "segment_ledger.jsonl"),
		filepath.Join(fixture.SessionDir, ".pipeline", "segment_manifest.jsonl"),
		filepath.Join(fixture.SessionDir, ".pipeline", "refined_transcript.md"),
		filepath.Join(fixture.SessionDir, ".pipeline", "topic_inventory.json"),
		filepath.Join(fixture.SessionDir, ".pipeline", "corrections_log.csv"),
		filepath.Join(fixture.SessionDir, ".pipeline", "uncertainty_report.json"),
		filepath.Join(fixture.SessionDir, ".pipeline", "structured_notes.md"),
		filepath.Join(fixture.SessionDir, ".pipeline", "coverage_matrix.json"),
		filepath.Join(fixture.SessionDir, ".pipeline", "enhanced_notes.md"),
		filepath.Join(fixture.SessionDir, ".pipeline", "validation_report.md"),
		filepath.Join(fixture.SessionDir, "final_notes.md"),
		filepath.Join(fixture.SessionDir, "bootcamp_index.md"),
		fixture.PublishedPath,
	} {
		require.FileExists(t, path)
	}

	publishedBody, err := os.ReadFile(fixture.PublishedPath)
	require.NoError(t, err)
	require.Contains(t, string(publishedBody), "# 🎓 AI/ML Class 01 [18/01/2026] - Fast-tracking the Course of AI")
}
