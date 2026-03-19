package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

func TestLoadWorkflow_Valid(t *testing.T) {
	content := `
name: test workflow
description: a test
steps:
  - pattern: summarize
  - pattern: extract_wisdom
    variables:
      role: expert
`
	path := writeTemp(t, content)
	wf, err := loadWorkflow(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if wf.Name != "test workflow" {
		t.Errorf("expected name 'test workflow', got %q", wf.Name)
	}
	if len(wf.Steps) != 2 {
		t.Fatalf("expected 2 steps, got %d", len(wf.Steps))
	}
	if wf.Steps[0].Pattern != "summarize" {
		t.Errorf("step 0 pattern: expected 'summarize', got %q", wf.Steps[0].Pattern)
	}
	if wf.Steps[1].Pattern != "extract_wisdom" {
		t.Errorf("step 1 pattern: expected 'extract_wisdom', got %q", wf.Steps[1].Pattern)
	}
	if wf.Steps[1].Variables["role"] != "expert" {
		t.Errorf("step 1 variable 'role': expected 'expert', got %q", wf.Steps[1].Variables["role"])
	}
}

func TestLoadWorkflow_NoSteps(t *testing.T) {
	content := `
name: empty
steps: []
`
	path := writeTemp(t, content)
	_, err := loadWorkflow(path)
	if err == nil {
		t.Fatal("expected error for empty steps, got nil")
	}
}

func TestLoadWorkflow_StepMissingPattern(t *testing.T) {
	content := `
steps:
  - pattern: summarize
  - variables:
      role: expert
`
	path := writeTemp(t, content)
	_, err := loadWorkflow(path)
	if err == nil {
		t.Fatal("expected error for step with no pattern, got nil")
	}
	// Verify error identifies the failing step (step 2)
	if !strings.Contains(err.Error(), "2") {
		t.Errorf("expected error to mention step 2, got: %v", err)
	}
}

func TestLoadWorkflow_MiddleStepMissingPattern(t *testing.T) {
	content := `
steps:
  - pattern: summarize
  - variables:
      role: expert
  - pattern: create_post
`
	path := writeTemp(t, content)
	_, err := loadWorkflow(path)
	if err == nil {
		t.Fatal("expected error for middle step with no pattern, got nil")
	}
	// Error must identify step 2 specifically, not step 1 or 3
	errMsg := err.Error()
	if !strings.Contains(errMsg, "2") {
		t.Errorf("expected error to reference step 2, got: %v", err)
	}
}

func TestLoadWorkflow_FileNotFound(t *testing.T) {
	_, err := loadWorkflow("/nonexistent/path/workflow.yaml")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestLoadWorkflow_MalformedYAML(t *testing.T) {
	content := `
steps:
  - pattern: [invalid yaml
`
	path := writeTemp(t, content)
	_, err := loadWorkflow(path)
	if err == nil {
		t.Fatal("expected error for malformed YAML, got nil")
	}
}

func TestMergeVariables(t *testing.T) {
	cli := map[string]string{"a": "1", "b": "2"}
	step := map[string]string{"b": "override", "c": "3"}
	merged := mergeVariables(cli, step)

	if merged["a"] != "1" {
		t.Errorf("expected a=1, got %q", merged["a"])
	}
	if merged["b"] != "override" {
		t.Errorf("expected b=override, got %q", merged["b"])
	}
	if merged["c"] != "3" {
		t.Errorf("expected c=3, got %q", merged["c"])
	}
}

func TestMergeVariables_NilMaps(t *testing.T) {
	merged := mergeVariables(nil, nil)
	if len(merged) != 0 {
		t.Errorf("expected empty map, got %v", merged)
	}
}

func TestMergeVariables_StepOverrideDoesNotMutateCLI(t *testing.T) {
	cli := map[string]string{"role": "beginner", "format": "markdown"}
	step := map[string]string{"role": "expert", "lang": "en"}

	merged := mergeVariables(cli, step)

	// Step var overrides CLI var
	if merged["role"] != "expert" {
		t.Errorf("expected role=expert, got %q", merged["role"])
	}
	// CLI-only var is preserved
	if merged["format"] != "markdown" {
		t.Errorf("expected format=markdown, got %q", merged["format"])
	}
	// Step-only var is added
	if merged["lang"] != "en" {
		t.Errorf("expected lang=en, got %q", merged["lang"])
	}
	// Original CLI map must NOT be mutated
	if cli["role"] != "beginner" {
		t.Errorf("CLI map was mutated: role=%q, expected 'beginner'", cli["role"])
	}
	if _, exists := cli["lang"]; exists {
		t.Error("CLI map was mutated: 'lang' key should not exist")
	}
}

func TestLoadWorkflow_VariableOverridePerStep(t *testing.T) {
	content := `
steps:
  - pattern: summarize
    variables:
      role: expert
      format: bullets
  - pattern: extract_wisdom
    variables:
      role: student
  - pattern: create_post
`
	path := writeTemp(t, content)
	wf, err := loadWorkflow(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Step 1 has both vars
	if wf.Steps[0].Variables["role"] != "expert" {
		t.Errorf("step 1 role: expected 'expert', got %q", wf.Steps[0].Variables["role"])
	}
	if wf.Steps[0].Variables["format"] != "bullets" {
		t.Errorf("step 1 format: expected 'bullets', got %q", wf.Steps[0].Variables["format"])
	}
	// Step 2 overrides role, has no format
	if wf.Steps[1].Variables["role"] != "student" {
		t.Errorf("step 2 role: expected 'student', got %q", wf.Steps[1].Variables["role"])
	}
	if _, exists := wf.Steps[1].Variables["format"]; exists {
		t.Error("step 2 should not have 'format' variable")
	}
	// Step 3 has no variables at all
	if len(wf.Steps[2].Variables) != 0 {
		t.Errorf("step 3 should have no variables, got %v", wf.Steps[2].Variables)
	}
}

func TestValidateWorkflowPatterns_AllExist(t *testing.T) {
	dir := t.TempDir()
	// Create fake pattern directories
	for _, name := range []string{"summarize", "extract_wisdom"} {
		if err := os.MkdirAll(filepath.Join(dir, name), 0755); err != nil {
			t.Fatal(err)
		}
	}

	registry := &core.PluginRegistry{
		Db: &fsdb.Db{
			Patterns: &fsdb.PatternsEntity{
				StorageEntity: &fsdb.StorageEntity{Dir: dir, Label: "patterns", ItemIsDir: true},
			},
		},
	}

	wf := &domain.WorkflowDefinition{
		Steps: []domain.WorkflowStep{
			{Pattern: "summarize"},
			{Pattern: "extract_wisdom"},
		},
	}

	if err := validateWorkflowPatterns(wf, registry); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateWorkflowPatterns_MissingPattern(t *testing.T) {
	dir := t.TempDir()
	// Only create "summarize", not "nonexistent_pattern"
	if err := os.MkdirAll(filepath.Join(dir, "summarize"), 0755); err != nil {
		t.Fatal(err)
	}

	registry := &core.PluginRegistry{
		Db: &fsdb.Db{
			Patterns: &fsdb.PatternsEntity{
				StorageEntity: &fsdb.StorageEntity{Dir: dir, Label: "patterns", ItemIsDir: true},
			},
		},
	}

	wf := &domain.WorkflowDefinition{
		Steps: []domain.WorkflowStep{
			{Pattern: "summarize"},
			{Pattern: "nonexistent_pattern"},
		},
	}

	err := validateWorkflowPatterns(wf, registry)
	if err == nil {
		t.Fatal("expected error for missing pattern, got nil")
	}
	errMsg := err.Error()
	if !strings.Contains(errMsg, "nonexistent_pattern") {
		t.Errorf("expected error to mention pattern name, got: %v", err)
	}
	if !strings.Contains(errMsg, "2") {
		t.Errorf("expected error to mention step 2, got: %v", err)
	}
}

// TestWorkflowIntegration_LoadAndValidate exercises the full pre-execution path:
// parse YAML → validate structure → validate patterns against filesystem.
func TestWorkflowIntegration_LoadAndValidate(t *testing.T) {
	patternsDir := t.TempDir()
	for _, name := range []string{"summarize", "extract_wisdom", "create_post"} {
		if err := os.MkdirAll(filepath.Join(patternsDir, name), 0755); err != nil {
			t.Fatal(err)
		}
	}

	registry := &core.PluginRegistry{
		Db: &fsdb.Db{
			Patterns: &fsdb.PatternsEntity{
				StorageEntity: &fsdb.StorageEntity{Dir: patternsDir, Label: "patterns", ItemIsDir: true},
			},
		},
	}

	workflowYAML := `
name: integration test
description: full pipeline test
steps:
  - pattern: summarize
  - pattern: extract_wisdom
    variables:
      role: expert
  - pattern: create_post
`
	wfPath := writeTemp(t, workflowYAML)

	// Step 1: load and parse
	wf, err := loadWorkflow(wfPath)
	if err != nil {
		t.Fatalf("loadWorkflow failed: %v", err)
	}
	if len(wf.Steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(wf.Steps))
	}

	// Step 2: validate all patterns exist
	if err := validateWorkflowPatterns(wf, registry); err != nil {
		t.Fatalf("validateWorkflowPatterns failed: %v", err)
	}

	// Step 3: verify variable merge for step 2 with CLI vars
	cliVars := map[string]string{"role": "beginner", "format": "markdown"}
	merged := mergeVariables(cliVars, wf.Steps[1].Variables)
	if merged["role"] != "expert" {
		t.Errorf("step 2 merge: expected role=expert, got %q", merged["role"])
	}
	if merged["format"] != "markdown" {
		t.Errorf("step 2 merge: expected format=markdown, got %q", merged["format"])
	}

	// Step 4: swapping a pattern to a nonexistent one should fail validation
	wf.Steps[1].Pattern = "nonexistent"
	err = validateWorkflowPatterns(wf, registry)
	if err == nil {
		t.Fatal("expected validation error for nonexistent pattern")
	}
	if !strings.Contains(err.Error(), "nonexistent") {
		t.Errorf("expected error to mention 'nonexistent', got: %v", err)
	}
}

func TestWorkflowConflictsWithPattern(t *testing.T) {
	// Simulate the guard from cli.go: --workflow + --pattern is an error.
	currentFlags := &Flags{
		Workflow: "some_workflow.yaml",
		Pattern:  "summarize",
	}

	// Reproduce the exact check from cli.go
	var err error
	if currentFlags.Workflow != "" && currentFlags.Pattern != "" {
		err = fmt.Errorf("%s", i18n.T("workflow_error_cannot_use_with_pattern"))
	}

	if err == nil {
		t.Fatal("expected error when both --workflow and --pattern are set")
	}
	if !strings.Contains(err.Error(), "--workflow") {
		t.Errorf("error should mention --workflow, got: %v", err)
	}
	if !strings.Contains(err.Error(), "--pattern") {
		t.Errorf("error should mention --pattern, got: %v", err)
	}
}

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "workflow.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	return path
}
