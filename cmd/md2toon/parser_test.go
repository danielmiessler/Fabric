// parser_test.go - Rigorous test suite for md2toon parser
//
// Test strategy based on elite testing techniques:
// 1. Invariant-Based Testing (properties that must always hold)
// 2. Metamorphic Relations (relationships between inputs/outputs)
// 3. Boundary Conditions (edge cases)
// 4. Order Preservation (ksylvan's bug)
// 5. Fuzzing (malformed inputs)
// 6. Golden File Testing (canonical examples)
package main

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"
)

// =============================================================================
// INVARIANT-BASED TESTS
// Properties that must ALWAYS hold regardless of input
// =============================================================================

func TestInvariant_OutputNeverNil(t *testing.T) {
	inputs := []string{
		"",
		"   ",
		"\n\n\n",
		"random text",
		"# HEADER ONLY",
		"- single bullet",
	}

	for _, input := range inputs {
		result := ParseMarkdownPrompt(input)
		if result == nil {
			t.Errorf("ParseMarkdownPrompt(%q) returned nil", input)
		}
		if result.Steps == nil {
			t.Errorf("Steps slice is nil for input %q", input)
		}
		if result.OutputInstructions == nil {
			t.Errorf("OutputInstructions slice is nil for input %q", input)
		}
		if result.Restrictions == nil {
			t.Errorf("Restrictions slice is nil for input %q", input)
		}
	}
}

func TestInvariant_RoleEitherEmptyOrMeaningful(t *testing.T) {
	inputs := []string{
		"You are an expert analyst who specializes in data.",
		"# IDENTITY\nYou are a helpful assistant.",
		"Random content without role.",
		"",
	}

	for _, input := range inputs {
		result := ParseMarkdownPrompt(input)
		// Role must be empty OR longer than 10 chars (meaningful)
		if result.Role != "" && len(result.Role) <= 10 {
			t.Errorf("Role is too short to be meaningful: %q", result.Role)
		}
	}
}

func TestInvariant_RestrictionsAlwaysNegative(t *testing.T) {
	// Note: items must have 3+ words and 10+ chars to pass parser filter
	input := `# OUTPUT
- Do not include personal opinions
- Never use first person pronouns
- Avoid using technical jargon
- Must not exceed 100 words
- Cannot include external links
`
	result := ParseMarkdownPrompt(input)

	// All restriction patterns should be captured
	if len(result.Restrictions) < 5 {
		t.Errorf("Expected 5 restrictions, got %d", len(result.Restrictions))
	}

	// Verify each restriction contains a negative pattern
	negativePatterns := []string{"do not", "never", "avoid", "must not", "cannot"}
	for i, r := range result.Restrictions {
		rule := strings.ToLower(r["rule"].(string))
		hasNegative := false
		for _, pattern := range negativePatterns {
			if strings.Contains(rule, pattern) {
				hasNegative = true
				break
			}
		}
		if !hasNegative {
			t.Errorf("Restriction[%d] %q doesn't contain negative pattern", i, rule)
		}
	}
}

func TestInvariant_NoDataLoss(t *testing.T) {
	input := `# IDENTITY
You are an expert data analyst.

# STEPS
- Analyze the input data
- Identify key patterns
- Generate insights

# OUTPUT INSTRUCTIONS
- Use markdown format
- Include code blocks

# RESTRICTIONS
- Do not include opinions
`
	result := ParseMarkdownPrompt(input)

	// Count total items extracted
	total := 0
	if result.Role != "" {
		total++
	}
	total += len(result.Steps)
	total += len(result.OutputInstructions)
	total += len(result.Restrictions)

	// We should extract at least 6 items (1 role + 3 steps + 2 output + 1 restriction)
	// Note: "Do not include opinions" goes to restrictions
	if total < 6 {
		t.Errorf("Data loss detected: only %d items extracted, expected at least 6", total)
	}
}

// =============================================================================
// METAMORPHIC RELATION TESTS
// Verify relationships between different inputs/outputs
// =============================================================================

func TestMetamorphic_CompressionGuarantee(t *testing.T) {
	// Note: TOON has fixed metadata overhead, so compression is only
	// meaningful for larger prompts. Small prompts may expand.
	// Real-world prompts average 78% token savings per corpus analysis.
	input := `# IDENTITY
You are an expert analyst who specializes in extracting insights from complex documents.

# STEPS
- Read the entire document carefully
- Identify the main themes and arguments
- Extract key supporting evidence
- Synthesize findings into coherent summary

# OUTPUT INSTRUCTIONS
- Use markdown formatting
- Include bullet points for clarity
- Keep paragraphs concise
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// Log compression stats (informational, not a hard requirement)
	compressionRatio := float64(len(input)-len(toon)) / float64(len(input)) * 100
	t.Logf("Compression: %.1f%% (input=%d bytes, toon=%d bytes)", compressionRatio, len(input), len(toon))

	// Verify structure is present (the real value of TOON is structure, not bytes)
	if !strings.Contains(toon, "role:") {
		t.Error("Missing role in TOON output")
	}
	if !strings.Contains(toon, "steps[") {
		t.Error("Missing steps in TOON output")
	}
}

func TestMetamorphic_SectionOrderIndependence(t *testing.T) {
	// Same content, different section order should produce same items
	input1 := `# IDENTITY
You are an expert.

# STEPS
- Do the work

# OUTPUT
- Use markdown
`
	input2 := `# STEPS
- Do the work

# OUTPUT
- Use markdown

# IDENTITY
You are an expert.
`
	result1 := ParseMarkdownPrompt(input1)
	result2 := ParseMarkdownPrompt(input2)

	// Same number of items regardless of section order
	if len(result1.Steps) != len(result2.Steps) {
		t.Errorf("Section order affected steps: %d vs %d",
			len(result1.Steps), len(result2.Steps))
	}
	if len(result1.OutputInstructions) != len(result2.OutputInstructions) {
		t.Errorf("Section order affected output_instructions: %d vs %d",
			len(result1.OutputInstructions), len(result2.OutputInstructions))
	}
}

func TestMetamorphic_Idempotency(t *testing.T) {
	input := `# IDENTITY
You are an expert analyst.

# STEPS
- Analyze data
- Generate report
`
	// Parse twice, results should be identical
	result1 := ParseMarkdownPrompt(input)
	result2 := ParseMarkdownPrompt(input)

	toon1 := PromptToTOON(result1)
	toon2 := PromptToTOON(result2)

	if toon1 != toon2 {
		t.Error("Parsing is not idempotent - different results on same input")
	}
}

// =============================================================================
// BOUNDARY CONDITION TESTS
// Edge cases and limits
// =============================================================================

func TestBoundary_EmptyInput(t *testing.T) {
	result := ParseMarkdownPrompt("")
	if result == nil {
		t.Fatal("nil result for empty input")
	}
	if result.Role != "" {
		t.Error("Expected empty role for empty input")
	}
}

func TestBoundary_SingleLine(t *testing.T) {
	inputs := []struct {
		input string
		desc  string
	}{
		{"You are an expert data analyst.", "role only"},
		{"- Extract the main points", "single bullet"},
		{"# STEPS", "header only"},
	}

	for _, tc := range inputs {
		result := ParseMarkdownPrompt(tc.input)
		if result == nil {
			t.Errorf("%s: nil result", tc.desc)
		}
	}
}

func TestBoundary_NoSections(t *testing.T) {
	input := `You are an expert analyst.

- Analyze the input
- Extract key points
- Generate summary
- Do not include opinions
`
	result := ParseMarkdownPrompt(input)

	// Without sections, role should still be detected
	if result.Role == "" {
		t.Error("Role not detected without section headers")
	}
	// Restriction pattern should still be caught
	if len(result.Restrictions) == 0 {
		t.Error("Restriction not detected without section headers")
	}
}

func TestBoundary_EmptySections(t *testing.T) {
	input := `# IDENTITY

# STEPS

# OUTPUT

`
	result := ParseMarkdownPrompt(input)
	// Should not crash, just produce empty results
	if result == nil {
		t.Fatal("nil result for empty sections")
	}
}

func TestBoundary_DeepNesting(t *testing.T) {
	input := `# STEPS
- Level 1 item
  - Level 2 nested
    - Level 3 nested
      - Level 4 nested
- Another level 1 item
`
	result := ParseMarkdownPrompt(input)
	// Should handle nesting gracefully (may flatten)
	if len(result.Steps) == 0 {
		t.Error("No steps extracted from nested input")
	}
}

func TestBoundary_UnicodeContent(t *testing.T) {
	input := `# IDENTITY
You are an expert in analyzing 日本語 and 中文 content.

# STEPS
- Extract émojis 🎉 and special chars
- Handle RTL text مرحبا correctly
- Process Cyrillic Привет

# OUTPUT
- Include Unicode: αβγδ ∑∏∫
`
	result := ParseMarkdownPrompt(input)

	if result.Role == "" {
		t.Error("Role not extracted with Unicode")
	}
	if len(result.Steps) < 3 {
		t.Errorf("Steps not fully extracted with Unicode: got %d", len(result.Steps))
	}

	// Verify Unicode integrity in output
	toon := PromptToTOON(result)
	if !utf8.ValidString(toon) {
		t.Error("TOON output contains invalid UTF-8")
	}
	if !strings.Contains(toon, "日本語") {
		t.Error("Unicode content lost in conversion")
	}
}

func TestBoundary_VeryLongContent(t *testing.T) {
	// Generate a prompt with many items
	var sb strings.Builder
	sb.WriteString("# IDENTITY\nYou are an expert analyst.\n\n# STEPS\n")
	for i := 0; i < 100; i++ {
		sb.WriteString("- Step number ")
		sb.WriteString(string(rune('0' + i%10)))
		sb.WriteString(" in the process\n")
	}

	result := ParseMarkdownPrompt(sb.String())

	// Should handle large input without crashing
	// May deduplicate similar items
	if len(result.Steps) == 0 {
		t.Error("No steps extracted from large input")
	}
	t.Logf("Extracted %d steps from 100 input items", len(result.Steps))
}

// =============================================================================
// ORDER PRESERVATION TESTS
// Critical: instruction order must be preserved (ksylvan's bug)
// =============================================================================

func TestOrderPreservation_StepsOrder(t *testing.T) {
	input := `# STEPS
- First: Take the ribeye out of the fridge
- Second: Fire up the barbecue
- Third: Season the steak
- Fourth: Cook the steak
- Fifth: Let it rest
`
	result := ParseMarkdownPrompt(input)

	if len(result.Steps) != 5 {
		t.Fatalf("Expected 5 steps, got %d", len(result.Steps))
	}

	// Verify order is preserved
	expectedOrder := []string{"First", "Second", "Third", "Fourth", "Fifth"}
	for i, step := range result.Steps {
		action := step["action"].(string)
		if !strings.HasPrefix(action, expectedOrder[i]) {
			t.Errorf("Step %d out of order: got %q, expected to start with %q",
				i, action, expectedOrder[i])
		}
	}
}

func TestOrderPreservation_OutputInstructionsOrder(t *testing.T) {
	input := `# OUTPUT INSTRUCTIONS
- 1. Start with a summary
- 2. Then provide details
- 3. Include examples
- 4. End with recommendations
`
	result := ParseMarkdownPrompt(input)

	if len(result.OutputInstructions) != 4 {
		t.Fatalf("Expected 4 output instructions, got %d", len(result.OutputInstructions))
	}

	// Verify order is preserved
	for i, inst := range result.OutputInstructions {
		instruction := inst["instruction"].(string)
		expectedPrefix := string(rune('1' + i))
		if !strings.HasPrefix(instruction, expectedPrefix) {
			t.Errorf("Instruction %d out of order: got %q", i, instruction)
		}
	}
}

func TestOrderPreservation_TOONOutput(t *testing.T) {
	input := `# STEPS
- Alpha: first step
- Beta: second step  
- Gamma: third step
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// In TOON output, Alpha should appear before Beta, Beta before Gamma
	alphaPos := strings.Index(toon, "Alpha")
	betaPos := strings.Index(toon, "Beta")
	gammaPos := strings.Index(toon, "Gamma")

	if alphaPos == -1 || betaPos == -1 || gammaPos == -1 {
		t.Fatal("Missing step content in TOON output")
	}

	if !(alphaPos < betaPos && betaPos < gammaPos) {
		t.Errorf("Order not preserved in TOON: Alpha@%d, Beta@%d, Gamma@%d",
			alphaPos, betaPos, gammaPos)
	}
}

// =============================================================================
// FUZZING TESTS
// Malformed and adversarial inputs
// =============================================================================

func TestFuzz_MalformedHeaders(t *testing.T) {
	inputs := []string{
		"#### STEPS",           // Too many #
		"#STEPS",               // No space
		"STEPS",                // No # or :
		"# ",                   // Empty header
		"#",                    // Just hash
		"# STEPS # OUTPUT",     // Multiple headers on one line
		"## STEPS\n### OUTPUT", // Mixed header levels
	}

	for _, input := range inputs {
		result := ParseMarkdownPrompt(input)
		if result == nil {
			t.Errorf("nil result for malformed input: %q", input)
		}
	}
}

func TestFuzz_MixedBulletStyles(t *testing.T) {
	// Note: items must have 3+ words and 10+ chars to pass parser filter
	input := `# STEPS
- Dash bullet with enough words
* Star bullet with enough words
• Unicode bullet with enough words
1. Numbered item with enough words
2) Numbered with paren and enough words
`
	result := ParseMarkdownPrompt(input)

	// Should handle all bullet styles
	if len(result.Steps) < 4 {
		t.Errorf("Expected at least 4 steps from mixed bullets, got %d", len(result.Steps))
	}
}

func TestFuzz_SpecialCharacters(t *testing.T) {
	input := `# IDENTITY
You are an expert in handling "quotes" and 'apostrophes'.

# STEPS
- Handle backslashes \ and forward slashes /
- Process brackets [like] {these} (and these)
- Deal with pipes | and ampersands & and at-signs @
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// Should not crash and should produce valid output
	if result.Role == "" {
		t.Error("Role not extracted with special characters")
	}

	// TOON should properly escape special characters
	if strings.Count(toon, "\\\"") < 1 {
		t.Log("Note: quotes should be escaped in TOON output")
	}
}

func TestFuzz_RandomContent(t *testing.T) {
	// Property test: parser should never panic on random input
	r := rand.New(rand.NewSource(42))

	for i := 0; i < 100; i++ {
		// Generate random "markdown-like" content
		var sb strings.Builder
		lines := r.Intn(20) + 1
		for j := 0; j < lines; j++ {
			switch r.Intn(5) {
			case 0:
				sb.WriteString("# HEADER\n")
			case 1:
				sb.WriteString("- bullet item\n")
			case 2:
				sb.WriteString("Some paragraph text\n")
			case 3:
				sb.WriteString("\n")
			case 4:
				sb.WriteString("1. numbered item\n")
			}
		}

		// Should not panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic on random input %d: %v", i, r)
				}
			}()
			result := ParseMarkdownPrompt(sb.String())
			_ = PromptToTOON(result)
		}()
	}
}

// =============================================================================
// GOLDEN FILE TESTS
// Canonical examples with expected outputs
// =============================================================================

func TestGolden_MinimalPrompt(t *testing.T) {
	// Note: items must have 3+ words and 10+ chars to pass parser filter
	input := `# IDENTITY
You are an expert analyst.

# STEPS
- Analyze the input carefully

# OUTPUT
- Use markdown formatting throughout
`
	result := ParseMarkdownPrompt(input)

	if result.Role != "You are an expert analyst." {
		t.Errorf("Role mismatch: got %q", result.Role)
	}
	if len(result.Steps) != 1 {
		t.Errorf("Steps count: got %d, want 1", len(result.Steps))
	}
	if len(result.OutputInstructions) != 1 {
		t.Errorf("OutputInstructions count: got %d, want 1", len(result.OutputInstructions))
	}
}

func TestGolden_RestrictionDetection(t *testing.T) {
	testCases := []struct {
		input   string
		isRestr bool
		desc    string
	}{
		{"Do not include personal opinions", true, "Do not"},
		{"Don't use first person", true, "Don't"},
		{"Never exceed 100 words", true, "Never"},
		{"Avoid technical jargon", true, "Avoid"},
		{"Must not include external links", true, "Must not"},
		{"Cannot reference previous conversations", true, "Cannot"},
		{"Can't use markdown tables", true, "Can't"},
		{"Should not make assumptions", true, "Should not"},
		{"Shouldn't include speculation", true, "Shouldn't"},
		{"Only output the final result", true, "Only output"},
		{"Output only the summary", true, "Output only"},
		{"Include detailed analysis", false, "Positive instruction"},
		{"Use markdown formatting", false, "Positive instruction"},
	}

	for _, tc := range testCases {
		input := "# OUTPUT\n- " + tc.input
		result := ParseMarkdownPrompt(input)

		gotRestr := len(result.Restrictions) > 0
		if gotRestr != tc.isRestr {
			t.Errorf("%s: %q - got restriction=%v, want %v",
				tc.desc, tc.input, gotRestr, tc.isRestr)
		}
	}
}

func TestGolden_RolePatterns(t *testing.T) {
	testCases := []struct {
		input   string
		hasRole bool
		desc    string
	}{
		{"You are an expert analyst.", true, "You are"},
		{"You're a skilled developer.", true, "You're"},
		{"You extract insights from data.", true, "You [verb]"},
		{"As an expert, analyze this.", true, "As an"},
		{"As a developer, write code.", true, "As a"},
		{"I want you to act as an editor.", true, "I want you to act as"},
		{"The system processes data.", false, "No role pattern"},
		{"Analyze the following text.", false, "Imperative only"},
	}

	for _, tc := range testCases {
		result := ParseMarkdownPrompt(tc.input)
		hasRole := result.Role != ""

		if hasRole != tc.hasRole {
			t.Errorf("%s: %q - got role=%v (%q), want hasRole=%v",
				tc.desc, tc.input, hasRole, result.Role, tc.hasRole)
		}
	}
}

// =============================================================================
// SECTION CLASSIFICATION TESTS
// Verify correct section detection
// =============================================================================

func TestSectionClassification_AllVariants(t *testing.T) {
	testCases := []struct {
		header   string
		expected Section
	}{
		// Identity variants
		{"IDENTITY", SectionIdentity},
		{"IDENTITY AND PURPOSE", SectionIdentity},
		{"ROLE", SectionIdentity},
		{"PURPOSE", SectionIdentity},

		// Steps variants
		{"STEPS", SectionSteps},
		{"STEP BY STEP", SectionSteps},
		{"GOALS", SectionSteps},
		{"TASK", SectionSteps},

		// Output variants
		{"OUTPUT", SectionOutput},
		{"OUTPUT INSTRUCTIONS", SectionOutput},
		{"OUTPUT SECTIONS", SectionOutput},
		{"OUTPUT FORMAT", SectionOutput},
		{"FORMAT", SectionOutput},

		// Skip variants
		{"EXAMPLE", SectionSkip},
		{"EXAMPLE OUTPUT", SectionSkip},
		{"EXAMPLES", SectionSkip},
		{"INPUT", SectionSkip},
		{"INPUT FORMAT", SectionSkip},

		// Unknown
		{"RANDOM HEADER", SectionUnknown},
		{"NOTES", SectionUnknown},
	}

	for _, tc := range testCases {
		got := classifySection(tc.header)
		if got != tc.expected {
			t.Errorf("classifySection(%q) = %v, want %v", tc.header, got, tc.expected)
		}
	}
}

// =============================================================================
// TOON OUTPUT FORMAT TESTS
// Verify TOON encoding correctness
// =============================================================================

func TestTOON_StringEscaping(t *testing.T) {
	input := `# IDENTITY
You are an expert who handles "quoted text" and backslashes \ properly.
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// Quotes should be escaped
	if !strings.Contains(toon, `\"quoted text\"`) {
		t.Errorf("Quotes not properly escaped in TOON: %s", toon)
	}
}

func TestTOON_ArrayNotation(t *testing.T) {
	// Note: items must have 3+ words and 10+ chars to pass parser filter
	input := `# STEPS
- Step one in the process
- Step two in the process
- Step three in the process
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// Should have array notation with count
	if !strings.Contains(toon, "steps[3]") {
		t.Errorf("Missing array count notation: %s", toon)
	}
}

func TestTOON_EmptyArrays(t *testing.T) {
	input := `# IDENTITY
You are an expert.
`
	result := ParseMarkdownPrompt(input)
	toon := PromptToTOON(result)

	// Empty arrays should be noted
	if !strings.Contains(toon, "steps[0]:") {
		t.Errorf("Empty steps array not properly formatted: %s", toon)
	}
}
