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

const technicalStudyGuideSample = `
Today we are studying binary search trees.
Each node keeps left children smaller and right children larger.
Insertion, lookup, and traversal all depend on that invariant.
`

type technicalStudyGuideFixture struct {
	RepoRoot      string
	SessionDir    string
	SourcePath    string
	InvocationDir string
}

type technicalStudyGuideSessionContext struct {
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

func TestTechnicalStudyGuidePipelineLoadsAndValidates(t *testing.T) {
	t.Parallel()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("technical-study-guide")
	require.NoError(t, err)
	require.NoError(t, Validate(pipe))
}

func TestTechnicalStudyGuidePrepareSourceStdinUsesRunDirArtifacts(t *testing.T) {
	t.Parallel()

	repoRoot := repoRootFromTest(t)
	invocationDir := t.TempDir()
	runDir := filepath.Join(invocationDir, ".pipeline", "test-run")
	require.NoError(t, os.MkdirAll(runDir, 0o755))

	prepareStdout := runPythonScript(
		t,
		repoRoot,
		filepath.Join(repoRoot, "scripts", "pipelines", "technical-study-guide", "prepare_source.py"),
		nil,
		strings.TrimSpace(technicalStudyGuideSample)+"\n",
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

	var sessionContext technicalStudyGuideSessionContext
	readJSONFile(t, sessionContextPath, &sessionContext)
	require.Equal(t, canonicalPath(t, invocationDir), canonicalPath(t, sessionContext.SessionDir))
	require.Equal(t, canonicalPath(t, runDir), canonicalPath(t, sessionContext.StablePipelineDir))
	require.Equal(
		t,
		filepath.Join(sessionContext.StablePipelineDir, "final_notes.md"),
		sessionContext.FinalOutputPath,
	)
}

func TestTechnicalStudyGuideRunnerEndToEnd(t *testing.T) {
	t.Parallel()

	fixture := newTechnicalStudyGuideFixture(t)
	pipe, result, stdout, stderr := runTechnicalStudyGuideWithPatternStubs(
		t,
		fixture,
		RunOptions{
			InvocationDir:  fixture.InvocationDir,
			DisableCleanup: true,
		},
	)

	require.Contains(t, stdout, "# 🎓 Technical Study Guide")
	require.Equal(t, strings.TrimSpace(result.FinalOutput), strings.TrimSpace(stdout))
	require.Contains(t, stderr, "[1/7] prepare_source ........ PASS")
	require.Contains(t, stderr, "[7/7] write_publish_manifest ........ PASS")
	require.Contains(t, stderr, "run summary: status=passed")

	assertTechnicalStudyGuideManifest(t, result.RunDir, pipe)
	assertTechnicalStudyGuideArtifacts(t, fixture, result.RunDir)
}

func newTechnicalStudyGuideFixture(t *testing.T) technicalStudyGuideFixture {
	t.Helper()

	rootDir := t.TempDir()
	sessionDir := filepath.Join(rootDir, "BST Session")
	require.NoError(t, os.MkdirAll(sessionDir, 0o755))
	sourcePath := filepath.Join(sessionDir, "transcript.md")
	require.NoError(t, os.WriteFile(sourcePath, []byte(strings.TrimSpace(technicalStudyGuideSample)+"\n"), 0o644))

	return technicalStudyGuideFixture{
		RepoRoot:      repoRootFromTest(t),
		SessionDir:    sessionDir,
		SourcePath:    sourcePath,
		InvocationDir: t.TempDir(),
	}
}

func runTechnicalStudyGuideWithPatternStubs(t *testing.T, fixture technicalStudyGuideFixture, opts RunOptions) (*Pipeline, *RunResult, string, string) {
	t.Helper()

	loader, err := NewDefaultLoader()
	require.NoError(t, err)

	pipe, err := loader.LoadNamed("technical-study-guide")
	require.NoError(t, err)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runner := NewRunner(&stdout, &stderr, nil)
	runner.stageExecutionOverride = technicalStudyGuidePatternOverride(t)

	result, err := runner.Run(context.Background(), pipe, RunSource{
		Mode:      SourceModeSource,
		Reference: fixture.SessionDir,
	}, opts)
	require.NoErrorf(t, err, "stderr:\n%s\nstdout:\n%s", stderr.String(), stdout.String())
	require.NotNil(t, result)

	return pipe, result, stdout.String(), stderr.String()
}

func technicalStudyGuidePatternOverride(t *testing.T) func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error) {
	t.Helper()

	return func(_ context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, bool, error) {
		switch stage.ID {
		case "semantic_map_generate":
			return &StageExecutionResult{Stdout: buildTechnicalSemanticMapOutput(t)}, true, nil
		case "render_generate":
			return &StageExecutionResult{Stdout: buildTechnicalRenderOutput()}, true, nil
		default:
			return nil, false, nil
		}
	}
}

func buildTechnicalSemanticMapOutput(t *testing.T) string {
	t.Helper()

	semanticMap := map[string]any{
		"topic":                  "Binary Search Trees",
		"session_focus":          "Build intuition for BST invariants and operations.",
		"prerequisites":          []string{"Basic recursion", "Understanding of tree nodes"},
		"learning_outcomes":      []string{"Explain BST invariants", "Describe search and insertion"},
		"topic_index":            []string{"BST invariant", "Lookup", "Insertion", "Traversal"},
		"conceptual_roadmap":     []string{"Start from ordering", "Apply ordering to traversal", "Extend to update operations"},
		"systems_visualization":  []string{"Root node", "Left subtree smaller", "Right subtree larger"},
		"skyline_intuition":      []string{"Tree as ordered skyline", "Walking left or right narrows search space"},
		"core_concepts":          []string{"Ordering invariant", "Recursive decomposition"},
		"mathematical_intuition": []string{"Each comparison halves the search frontier in a balanced tree"},
		"coding_walkthroughs":    []string{"Lookup pseudocode", "Insertion pseudocode"},
		"advanced_scenario":      []string{"Keeping latency stable as the tree becomes skewed"},
		"hots":                   []string{"Why does skew degrade search time?", "How do traversals expose ordering?"},
		"faq":                    []string{"Why not always use arrays?", "When does balancing matter?"},
		"practice_roadmap":       []string{"Implement search", "Implement insert", "Print inorder traversal"},
		"next_improvements":      []string{"Balance the tree", "Compare against AVL or Red-Black trees"},
		"related_notes":          []string{"Recursion patterns", "Tree traversals"},
	}
	content, err := json.MarshalIndent(semanticMap, "", "  ")
	require.NoError(t, err)

	renderInput := strings.Join([]string{
		"# Rendering Brief",
		"",
		"- Topic: Binary Search Trees",
		"- Learner level: early intermediate",
		"",
		"## Key Intuitions",
		"- Ordering invariant",
		"- Comparisons route the search",
		"",
		"## Sections",
		"- Session focus",
		"- Systems visualization",
		"- Coding walkthroughs",
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

func buildTechnicalRenderOutput() string {
	finalNotes := strings.Join([]string{
		"# 🎓 Technical Study Guide - Binary Search Trees",
		"",
		"## 🧠 Session Focus",
		"Build an intuition-first mental model for binary search tree ordering and the operations that depend on it.",
		"",
		"## 🎯 Prerequisites",
		"- Basic recursion",
		"- Comfort with tree nodes and pointers",
		"",
		"## ✅ Learning Outcomes",
		"- Explain the BST ordering invariant",
		"- Trace lookup and insertion decisions",
		"",
		"## 🧭 Topic Index",
		"- BST invariant",
		"- Search path",
		"- Insert path",
		"- Traversal behavior",
		"",
		"## 🗺️ Conceptual Roadmap",
		"```mermaid",
		"flowchart TD",
		`  A["Ordering Invariant"] --> B["Directed Comparisons"]`,
		`  B --> C["Lookup and Insertion"]`,
		"```",
		"",
		"## 🏗️ Systems Visualization",
		"```mermaid",
		"graph TD",
		`  A["Root"] --> B["Left < Root"]`,
		`  A --> C["Right > Root"]`,
		"```",
		"",
		"## 🌆 Skyline Intuition Diagram",
		"Think of the tree as an ordered skyline where each comparison moves you toward the right building instead of scanning every building in sequence.",
		"",
		"## 📚 Core Concepts (Intuition First)",
		"The BST works because every node acts like a local sorting promise. Once that promise holds, each comparison tells you which entire half of the tree can be ignored.",
		"",
		"## ➗ Mathematical Intuition",
		"In a balanced tree, each comparison removes roughly half of the remaining search frontier, which is why lookup trends toward logarithmic depth.",
		"",
		"## 💻 Coding Walkthroughs",
		"- Lookup: compare target with current node, then recurse or iterate left/right.",
		"- Insert: follow the same comparison path until an empty child slot is found.",
		"",
		"## 🚀 Advanced Real-World Scenario",
		"If inserts arrive in sorted order, the BST can degenerate into a linked list. That failure mode explains why self-balancing trees matter in production systems.",
		"",
		"## 🧩 HOTS (High-Order Thinking)",
		"1. Why does sorted insertion destroy the performance intuition of a BST?",
		"2. How does inorder traversal reveal whether the invariant has been preserved?",
		"",
		"## ❓ FAQ",
		"Q: Why not just use an array?",
		"A: Arrays are great for contiguous storage, but naive insertions can be expensive.",
		"",
		"Q: When does balancing become important?",
		"A: As soon as insert order can create long skewed paths.",
		"",
		"## 🛠️ Practice Roadmap",
		"1. Implement iterative search.",
		"2. Implement recursive insert.",
		"3. Print inorder traversal and verify sorted output.",
		"",
		"## 🔭 Next Improvements",
		"- Compare BST behavior with AVL trees.",
		"- Measure lookup cost for balanced vs skewed inputs.",
		"",
		"## 🔗 Related Notes",
		"- Tree traversals",
		"- Recursion patterns",
	}, "\n")

	return strings.Join([]string{
		"<<<BEGIN_ARTIFACT:final_notes.md>>>",
		finalNotes,
		"<<<END_ARTIFACT>>>",
	}, "\n") + "\n"
}

func assertTechnicalStudyGuideManifest(t *testing.T, runDir string, pipe *Pipeline) {
	t.Helper()

	var manifest RunManifest
	readJSONFile(t, filepath.Join(runDir, "run_manifest.json"), &manifest)
	require.Equal(t, pipe.Name, manifest.PipelineName)
	require.Equal(t, "passed", manifest.Status)
	require.Len(t, manifest.Stages, len(pipe.Stages))
}

func assertTechnicalStudyGuideArtifacts(t *testing.T, fixture technicalStudyGuideFixture, runDir string) {
	t.Helper()

	require.FileExists(t, filepath.Join(runDir, "run_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "source_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "publish_manifest.json"))
	require.FileExists(t, filepath.Join(runDir, "session_context.json"))
	require.FileExists(t, filepath.Join(runDir, "semantic_map_input.md"))
	require.FileExists(t, filepath.Join(runDir, "semantic_map.json"))
	require.FileExists(t, filepath.Join(runDir, "render_input.md"))
	require.FileExists(t, filepath.Join(runDir, "validation_report.md"))

	stablePipelineDir := filepath.Join(fixture.SessionDir, ".pipeline")
	require.FileExists(t, filepath.Join(stablePipelineDir, "semantic_map.json"))
	require.FileExists(t, filepath.Join(stablePipelineDir, "validation_report.md"))
	require.FileExists(t, filepath.Join(fixture.SessionDir, "final_notes.md"))

	finalNotes, err := os.ReadFile(filepath.Join(fixture.SessionDir, "final_notes.md"))
	require.NoError(t, err)
	require.Contains(t, string(finalNotes), "## 🏗️ Systems Visualization")
	require.Contains(t, string(finalNotes), "```mermaid")
}
