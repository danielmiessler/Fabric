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

const nontechnicalStudyGuideSample = `
This session explores how social institutions shape personal choices.
It compares individual responsibility with structural constraints.
The discussion includes family expectations, economic pressure, and identity.
`

type nontechnicalStudyGuideFixture struct {
	RepoRoot      string
	SessionDir    string
	SourcePath    string
	InvocationDir string
}

type nontechnicalStudyGuideSessionContext struct {
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

func TestNontechnicalStudyGuidePipelineLoadsAndValidates(t *testing.T) {
	t.Parallel()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("nontechnical-study-guide")
	require.NoError(t, err)
	require.NoError(t, Validate(pipe))
}

func TestNontechnicalStudyGuidePrepareSourceStdinUsesRunDirArtifacts(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "nontechnical-study-guide", "prepare_source.py"),
		nil,
		strings.TrimSpace(nontechnicalStudyGuideSample)+"\n",
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

	var sessionContext nontechnicalStudyGuideSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, invocationDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, canonicalPath(t, runDir), canonicalPath(t, sessionContext.StablePipelineDir))
	require.Equal(
		t,
		filepath.Join(sessionContext.StablePipelineDir, "final_notes.md"),
		sessionContext.FinalOutputPath,
	)
}

func TestNontechnicalStudyGuideRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newNontechnicalStudyGuideFixture(t)
	pipe, result, stdout, stderr := runNontechnicalStudyGuideWithPatternStubs(
		t,
		fixture,
		RunOptions{
			InvocationDir:  fixture.InvocationDir,
			DisableCleanup: true,
		},
	)

	require.Contains(t, stdout, "# 📝 Nontechnical Study Guide - Institutions and Agency")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[1/7] prepare_source ........ PASS")
	require.Contains(t, stderr, "[7/7] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertNontechnicalStudyGuideManifest(t, result.RunDir, pipe)
	assertNontechnicalStudyGuideArtifacts(t, fixture, result.RunDir)
}

func newNontechnicalStudyGuideFixture(t *testing.T) nontechnicalStudyGuideFixture {
	t.Helper()

	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "Context Session")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))
	sourcePath := filepath.Join(sessionDir, "transcript.md")
	require.NoError(t, os.WriteFile(sourcePath, []byte(strings.TrimSpace(nontechnicalStudyGuideSample)+"\n"), 0o644))

	return nontechnicalStudyGuideFixture{
		RepoRoot:      repoRootFromTest(t),
		SessionDir:    sessionDir,
		SourcePath:    sourcePath,
		InvocationDir: t.TempDir(),
	}
}

func runNontechnicalStudyGuideWithPatternStubs(t *testing.T, fixture nontechnicalStudyGuideFixture, opts RunOptions) (*Pipeline, *RunResult, string, string) {
	t.Helper()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("nontechnical-study-guide")
	require.NoError(t, err)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	runner.stageExecutionOverride = nontechnicalStudyGuidePatternOverride(t)

	result, err := runner.Run(context.Background(), pipe, RunSource{
		Mode:      SourceModeSource,
		Reference: fixture.SessionDir,
	}, opts)
	require.NoErrorf(t, err, "stderr:\n%s\nstdout:\n%s", stderr.String(), stdout.String())
	require.NotNil(t, result)

	return pipe, result, stdout.String(), stderr.String()
}

func nontechnicalStudyGuidePatternOverride(t *testing.T) func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error) {
	t.Helper()

	return func(_ context.Context, stage Stage, _ StageRuntimeContext) (*StageExecutionResult, bool, error) {
		switch stage.ID {
		case "semantic_map_generate":
			return &StageExecutionResult{Stdout: buildNontechnicalSemanticMapOutput(t)}, true, nil
		case "render_generate":
			return &StageExecutionResult{Stdout: buildNontechnicalRenderOutput()}, true, nil
		default:
			return nil, false, nil
		}
	}
}

func buildNontechnicalSemanticMapOutput(t *testing.T) string {
	t.Helper()

	semanticMap := map[string]any{
		"topic":                           "Institutions and Agency",
		"session_focus":                   "Understand how social systems constrain and enable individual action.",
		"context_you_need_before_this":    []string{"Basic familiarity with social roles", "Awareness of economic pressure in daily life"},
		"big_ideas":                       []string{"Institutions shape options", "Agency operates inside constraints", "Identity is negotiated socially"},
		"what_was_actually_said":          []string{"Family expectations influence choices", "Economic pressure can narrow decisions"},
		"why_it_matters":                  []string{"Explains why motivation alone is insufficient", "Supports realistic planning"},
		"intuition_examples":              []string{"Job choice under family pressure", "Education decisions under budget limits", "Identity shifts across contexts"},
		"argument_or_theme_map":           []string{"Structure influences incentives", "Incentives influence behavior", "Behavior reshapes identity"},
		"difficult_terms":                 []map[string]string{{"term": "institution", "plain_language_meaning": "A stable social setup that shapes behavior"}, {"term": "agency", "plain_language_meaning": "Your practical ability to choose and act"}, {"term": "constraint", "plain_language_meaning": "A limit that narrows available options"}},
		"misunderstandings_to_avoid":      []string{"This is not pure determinism", "Context does not erase responsibility"},
		"reflection_questions":            []string{"Where do constraints help me?", "Where do they limit me?", "What choices can I still influence?", "Which assumptions are inherited?"},
		"revision_summary":                []string{"Context before judgment", "Map constraints before planning", "Use examples to test interpretations"},
		"connections_to_other_topics":     []string{"Behavioral economics", "Sociology of identity", "Decision-making under uncertainty"},
		"profile_extensions":              map[string]any{"viewpoints": []string{"individualist", "structural"}, "open_questions": []string{"What changes first: beliefs or structures?"}},
	}
	content, err := json.MarshalIndent(semanticMap, "", "  ")
	require.NoError(t, err)

	renderInput := strings.Join([]string{
		"# Rendering Brief",
		"",
		"- Topic: Institutions and Agency",
		"- Reader level: general nontechnical learner",
		"",
		"## Section Order",
		"- Context first",
		"- Then ideas and claims",
		"- Then examples, terms, and reflection",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:semantic_map.json>>>",
		string(content),
		"<<<END_ARTIFACT>>>",
		"<<<BEGIN_ARTIFACT:render_input.md>>>",
		renderInput,
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func buildNontechnicalRenderOutput() string {
	finalNotes := strings.Join([]string{
		"# 📝 Nontechnical Study Guide - Institutions and Agency",
		"",
		"## Session Focus",
		"This session explains how social structures and personal choice interact in everyday decisions.",
		"",
		"## Context You Need Before This",
		"- Social roles can shape expectations before choices are made.",
		"- Economic pressure can narrow practical options.",
		"",
		"## Big Ideas",
		"- Institutions influence incentives.",
		"- People still exercise agency within limits.",
		"- Identity changes across social contexts.",
		"",
		"## What Was Actually Said",
		"- Family expectations were described as decision drivers.",
		"- Economic pressure was described as reducing option space.",
		"- Personal responsibility remained important but contextualized.",
		"",
		"## Why It Matters",
		"- It prevents simplistic blame narratives.",
		"- It helps build realistic plans under constraints.",
		"",
		"## Intuition and Real-Life Examples",
		"- **Example:** A student chooses a local job due to household duties.",
		"- **Example:** Someone delays a risky career move because of debt pressure.",
		"- **Example:** A person adopts different behaviors at work and at home to fit expectations.",
		"",
		"## Argument or Theme Map",
		"```mermaid",
		"flowchart TD",
		`  A["Institutions"] --> B["Incentives"]`,
		`  B --> C["Choices"]`,
		`  C --> D["Identity"]`,
		"```",
		"",
		"## Difficult Terms Explained Simply",
		"- **Institution:** A stable social setup that guides behavior.",
		"- **Agency:** Your practical ability to choose and act.",
		"- **Constraint:** A real limit on available choices.",
		"",
		"## Misunderstandings to Avoid",
		"- Context is not an excuse for all outcomes.",
		"- Responsibility is not erased by social pressure.",
		"",
		"## Reflection Questions",
		"- Where do your choices feel most constrained?",
		"- Which constraints are external versus internalized?",
		"- What small actions increase your agency this week?",
		"- Which assumptions came from family or institutions?",
		"",
		"## Revision Summary",
		"- Start with context before making judgments.",
		"- Track incentives before explaining behavior.",
		"- Use concrete examples to test abstract claims.",
		"",
		"## Connections to Other Topics",
		"- Behavioral economics",
		"- Identity formation",
		"- Decision-making under pressure",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:final_notes.md>>>",
		finalNotes,
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func assertNontechnicalStudyGuideManifest(t *testing.T, runDir string, pipe *Pipeline) {
	t.Helper()

	var manifest RunManifest
	readJSONFile(t, filepath.Join(runDir, "run_manifest.json"), &manifest)
	require.Equal(t, pipe.Name, manifest.PipelineName)
	require.Equal(t, "passed", manifest.Status)
	require.Len(t, manifest.Stages, len(pipe.Stages))
}

func assertNontechnicalStudyGuideArtifacts(t *testing.T, fixture nontechnicalStudyGuideFixture, runDir string) {
	t.Helper()

	require.FileExists(t, filepath.Join(runDir, "run_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "source_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "publish_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "session_context.json"))
	require.FileExists(t, filepath.Join(runDir, "semantic_map.json"))
	require.FileExists(t, filepath.Join(runDir, "render_input.md"))
	require.FileExists(t, filepath.Join(runDir, "validation_report.md"))
	require.FileExists(t, filepath.Join(runDir, "final_notes.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "semantic_map.json"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, ".pipeline", "validation_report.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, "final_notes.md"))
}
