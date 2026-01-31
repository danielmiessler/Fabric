package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSectionAware_Steps(t *testing.T) {
	input := `# IDENTITY
You are an expert.

# STEPS
- Extract the main ideas
- Write a summary
- Create a list
`
	result := ParseMarkdownPrompt(input)

	// All items under STEPS should be steps
	if len(result.Steps) != 3 {
		t.Errorf("Steps = %d, want 3", len(result.Steps))
	}
}

func TestSectionAware_SkipExample(t *testing.T) {
	input := `# STEPS
- Do the work

# EXAMPLE OUTPUT
- This is example content that should be skipped
- More example content

# OUTPUT INSTRUCTIONS
- Use markdown
`
	result := ParseMarkdownPrompt(input)

	// Example content should NOT be extracted
	if len(result.Steps) != 1 {
		t.Errorf("Steps = %d, want 1 (example should be skipped)", len(result.Steps))
	}
}

func TestSectionAware_OutputSection(t *testing.T) {
	input := `# OUTPUT
- Output the results in markdown format
- Do not include opinions
`
	result := ParseMarkdownPrompt(input)

	// "Do not" should be restriction
	if len(result.Restrictions) < 1 {
		t.Errorf("Restrictions = %d, want 1+", len(result.Restrictions))
	}
}

func TestCorpusValidation(t *testing.T) {
	files, _ := filepath.Glob("../../data/patterns/*/system.md")
	if len(files) == 0 {
		t.Skip("no patterns found")
	}

	var passed, failed int
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			continue
		}

		toon := PromptToTOON(ParseMarkdownPrompt(string(content)))
		if strings.Contains(toon, "role:") && strings.Contains(toon, "steps[") {
			passed++
		} else {
			failed++
		}
	}

	t.Logf("Corpus: %d passed, %d failed", passed, failed)
	if failed > 10 {
		t.Errorf("Too many failures: %d", failed)
	}
}

func TestCLI(t *testing.T) {
	if err := exec.Command("go", "build", "-o", "md2toon_test", ".").Run(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("md2toon_test")

	cmd := exec.Command("./md2toon_test")
	cmd.Stdin = strings.NewReader("# IDENTITY\nYou are an expert.\n\n# STEPS\n- Do the work\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(out), "role:") || !strings.Contains(string(out), "steps[") {
		t.Error("Missing expected fields")
	}
}
