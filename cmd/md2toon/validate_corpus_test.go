package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func validateTOON(toon string) []string {
	var errs []string
	required := []string{"role:", "steps[", "output_format:", "output_instructions[", "restrictions["}
	for _, r := range required {
		if !strings.Contains(toon, r) {
			errs = append(errs, "missing "+r)
		}
	}
	if len(strings.TrimSpace(toon)) < 50 {
		errs = append(errs, "output too short")
	}
	for _, line := range strings.Split(toon, "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "[") && strings.Contains(line, "]") {
			if strings.Count(line, "[") != strings.Count(line, "]") {
				errs = append(errs, "unbalanced brackets")
				break
			}
		}
	}
	return errs
}

func TestFullCorpusValidation(t *testing.T) {
	files, err := filepath.Glob("../../data/patterns/*/system.md")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no pattern files found")
	}

	var success, failed, mdChars, toonChars int
	var failures []string

	for _, f := range files {
		name := filepath.Base(filepath.Dir(f))
		md, err := os.ReadFile(f)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", name, err))
			failed++
			continue
		}
		mdChars += len(md)

		toon := PromptToTOON(ParseMarkdownPrompt(string(md)))
		toonChars += len(toon)

		if errs := validateTOON(toon); len(errs) > 0 {
			failures = append(failures, fmt.Sprintf("%s: %v", name, errs))
			failed++
		} else {
			success++
		}
	}

	savings := float64(mdChars-toonChars) / float64(mdChars) * 100
	t.Logf("Patterns: %d, Passed: %d, Failed: %d, Savings: %.1f%%", len(files), success, failed, savings)

	if failed > 0 {
		for _, f := range failures {
			t.Logf("  FAIL: %s", f)
		}
		t.Fatalf("%d/%d failed", failed, len(files))
	}
}

func TestCLI(t *testing.T) {
	if err := exec.Command("go", "build", "-o", "md2toon_test", ".").Run(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("md2toon_test")

	cmd := exec.Command("./md2toon_test")
	cmd.Stdin = strings.NewReader("# IDENTITY\nYou are an expert.\n\n# STEPS\n- Step one\n\n# OUTPUT\n- Markdown")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	toon := string(out)
	if !strings.Contains(toon, "role:") || !strings.Contains(toon, "steps[") {
		t.Error("missing expected fields in output")
	}
}
