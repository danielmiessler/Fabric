package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// fakePatterns implements patternResolver for validation tests.
type fakePatterns struct {
	known map[string]bool
}

func (f *fakePatterns) GetRaw(name string) (*fsdb.Pattern, error) {
	if f.known[name] {
		return &fsdb.Pattern{Name: name, Pattern: "# " + name}, nil
	}
	return nil, fmt.Errorf("pattern %q not found", name)
}

func TestLoadWorkflowYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "wf.yaml")
	content := `name: demo
steps:
  - pattern: summarize
  - pattern: analyze_claims
    variables:
      role: expert
  - pattern: create_social_posts
    model: gpt-4o
`
	os.WriteFile(path, []byte(content), 0o644)

	wf, err := LoadWorkflow(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if wf.Name != "demo" || len(wf.Steps) != 3 {
		t.Fatalf("unexpected workflow: %+v", wf)
	}
	if wf.Steps[1].Variables["role"] != "expert" {
		t.Errorf("step 1 variables = %v", wf.Steps[1].Variables)
	}
	if wf.Steps[2].Model != "gpt-4o" {
		t.Errorf("step 2 model = %q", wf.Steps[2].Model)
	}
}

func TestLoadWorkflowJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "wf.json")
	os.WriteFile(path, []byte(`{"steps":[{"pattern":"a"},{"pattern":"b"}]}`), 0o644)

	wf, err := LoadWorkflow(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(wf.Steps) != 2 || wf.Steps[1].Pattern != "b" {
		t.Fatalf("unexpected steps: %+v", wf.Steps)
	}
}

func TestLoadWorkflowBadPath(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"absolute missing", "/definitely/does/not/exist.yaml"},
		{"relative missing", "./missing_workflow.yaml"},
		{"directory not file", t.TempDir()},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := LoadWorkflow(tc.path)
			if err == nil {
				t.Fatalf("expected error for %q, got nil", tc.path)
			}
			if !strings.Contains(err.Error(), "workflow") {
				t.Errorf("error %q should mention 'workflow'", err.Error())
			}
		})
	}
}

func TestLoadWorkflowMalformedYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.yaml")
	os.WriteFile(path, []byte("steps: [this is: not: valid"), 0o644)

	_, err := LoadWorkflow(path)
	if err == nil || !strings.Contains(err.Error(), "parsing workflow YAML") {
		t.Fatalf("expected YAML parse error, got %v", err)
	}
}

func TestValidate(t *testing.T) {
	store := &fakePatterns{known: map[string]bool{
		"summarize":  true,
		"improve_it": true,
	}}

	tests := []struct {
		name    string
		wf      *Workflow
		store   patternResolver
		wantErr string // substring that must appear in the error, "" = no error
	}{
		{
			name:  "happy path",
			wf:    &Workflow{Steps: []WorkflowStep{{Pattern: "summarize"}, {Pattern: "improve_it"}}},
			store: store,
		},
		{
			name:    "nil workflow",
			wf:      nil,
			store:   store,
			wantErr: "workflow is nil",
		},
		{
			name:    "no steps",
			wf:      &Workflow{},
			store:   store,
			wantErr: "no steps",
		},
		{
			name:    "empty pattern",
			wf:      &Workflow{Steps: []WorkflowStep{{Pattern: "summarize"}, {Pattern: "  "}}},
			store:   store,
			wantErr: "[step 2/2 ?] missing required field 'pattern'",
		},
		{
			name:    "duplicate consecutive",
			wf:      &Workflow{Steps: []WorkflowStep{{Pattern: "summarize"}, {Pattern: "summarize"}}},
			store:   store,
			wantErr: "[step 2/2 summarize] repeats",
		},
		{
			name:    "unknown pattern",
			wf:      &Workflow{Steps: []WorkflowStep{{Pattern: "summarize"}, {Pattern: "does_not_exist"}}},
			store:   store,
			wantErr: "[step 2/2 does_not_exist] pattern not found",
		},
		{
			name:  "file path skips existence check",
			wf:    &Workflow{Steps: []WorkflowStep{{Pattern: "./local/prompt.md"}}},
			store: store,
		},
		{
			name: "nil store skips existence check",
			wf:   &Workflow{Steps: []WorkflowStep{{Pattern: "anything"}}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.wf.Validate(tc.store)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error %q does not contain %q", err.Error(), tc.wantErr)
			}
		})
	}
}

func TestStepLabel(t *testing.T) {
	got := stepLabel(0, 3, "summarize")
	want := "[step 1/3 summarize]"
	if got != want {
		t.Fatalf("stepLabel = %q, want %q", got, want)
	}
}

func TestExtractStepOutput(t *testing.T) {
	ok := &fsdb.Session{}
	ok.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleAssistant, Content: "  hello  "})

	blank := &fsdb.Session{}
	blank.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleAssistant, Content: "   \n\t  "})

	tests := []struct {
		name    string
		sess    *fsdb.Session
		want    string
		wantErr string
	}{
		{"trims and returns content", ok, "hello", ""},
		{"nil session", nil, "", "no session"},
		{"empty session", &fsdb.Session{}, "", "no message"},
		{"whitespace only", blank, "", "empty output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := extractStepOutput(tc.sess)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if out != tc.want {
					t.Fatalf("got %q, want %q", out, tc.want)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("got err=%v, want substring %q", err, tc.wantErr)
			}
		})
	}
}

func TestResolveStepInputWithOverride(t *testing.T) {
	carried := "previous step output"

	tests := []struct {
		name         string
		step         WorkflowStep
		wantInput    string
		wantOverride bool
	}{
		{
			name:      "no input field -> carried through",
			step:      WorkflowStep{Pattern: "p"},
			wantInput: carried,
		},
		{
			name:         "explicit input overrides carried",
			step:         WorkflowStep{Pattern: "p", Input: "  custom text  "},
			wantInput:    "custom text",
			wantOverride: true,
		},
		{
			name:      "whitespace-only input is NOT an override",
			step:      WorkflowStep{Pattern: "p", Input: "  \n\t  "},
			wantInput: carried,
		},
		{
			name:      "empty string is NOT an override",
			step:      WorkflowStep{Pattern: "p", Input: ""},
			wantInput: carried,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, overridden := resolveStepInputWithOverride(tc.step, carried)
			if got != tc.wantInput {
				t.Errorf("input = %q, want %q", got, tc.wantInput)
			}
			if overridden != tc.wantOverride {
				t.Errorf("overridden = %v, want %v", overridden, tc.wantOverride)
			}
		})
	}
}

func TestResolveStepInputWithOverrideChain(t *testing.T) {
	// 4-step chain proving: override does NOT leak into later steps and
	// whitespace-only input never overrides.
	steps := []WorkflowStep{
		{Pattern: "a"},                  // uses stdin
		{Pattern: "b", Input: "custom"}, // overrides
		{Pattern: "c", Input: "   "},    // whitespace -> use carried (out-b)
		{Pattern: "d"},                  // uses out-c
	}
	outputs := []string{"out-a", "out-b", "out-c", "out-d"}
	wantIn := []string{"stdin", "custom", "out-b", "out-c"}
	wantOv := []bool{false, true, false, false}

	carried := "stdin"
	for i, step := range steps {
		gotIn, gotOv := resolveStepInputWithOverride(step, carried)
		if gotIn != wantIn[i] || gotOv != wantOv[i] {
			t.Fatalf("step %d: got (%q, %v), want (%q, %v)",
				i+1, gotIn, gotOv, wantIn[i], wantOv[i])
		}
		carried = outputs[i] // simulate model response
	}
}

func TestMergeVars(t *testing.T) {
	got := mergeVars(
		map[string]string{"a": "1", "b": "2"},
		map[string]string{"b": "override", "c": "3"},
	)
	if got["a"] != "1" || got["b"] != "override" || got["c"] != "3" {
		t.Errorf("unexpected merge result: %v", got)
	}
	if mergeVars(nil, nil) != nil {
		t.Error("expected nil for empty merge")
	}
}
