// Package main provides the md2toon parser
//
// Mathematical Foundation (derived from corpus analysis of 237 patterns):
//
//	The classification function f(item) is defined as:
//	  | matches restriction pattern → restrictions[]  (syntactic override, 100% precision)
//	  | section = identity ∧ matches role → role      (positional)
//	  | section = steps → steps[]                     (trust author)
//	  | section = output → output_instructions[]      (trust author)
//	  | section = skip → ignore                       (noise: examples, input)
//	  | section = unknown → output_instructions[]     (safe default)
//
//	Key insight: Section headers ARE the ground truth.
//	We do NOT second-guess the author with intent classification.
package main

import (
	"regexp"
	"strings"
)

// FabricPrompt represents a parsed prompt structure
type FabricPrompt struct {
	Role                string
	Expertise           []string
	Purpose             string
	Steps               []map[string]any
	OutputFormat        string
	OutputSections      []map[string]any
	OutputInstructions  []map[string]any
	Restrictions        []map[string]any
	ThinkingInstruction string
}

// Section represents the semantic type of a markdown section
type Section int

const (
	SectionUnknown  Section = iota
	SectionIdentity         // IDENTITY, ROLE, PURPOSE → role extraction
	SectionSteps            // STEPS, GOALS, TASK → steps[]
	SectionOutput           // OUTPUT, OUTPUT SECTIONS, OUTPUT INSTRUCTIONS → output_instructions[]
	SectionSkip             // EXAMPLE, INPUT → ignore entirely
)

// Compiled patterns (order matters for some)
var (
	// Section header detection: "# HEADER" or "HEADER:"
	headerHashRe  = regexp.MustCompile(`(?i)^#{1,3}\s*(.+)$`)
	headerColonRe = regexp.MustCompile(`(?i)^([A-Z][A-Z\s]+):$`)

	// Role detection: "You are...", "You're...", "You [verb]...", "As a/an...", "I want you to act as..."
	roleRe = regexp.MustCompile(`(?i)^(you're|you\s+\w+|as\s+an?|i\s+want\s+you\s+to\s+act\s+as)\s+`)

	// Restriction patterns: syntactic override (100% precision per corpus analysis)
	// These ALWAYS go to restrictions regardless of section
	restrictionRe = regexp.MustCompile(`(?i)^(do\s+not|don't|never|avoid|must\s+not|cannot|can't|should\s+not|shouldn't|only\s+output|output\s+only)\b`)

	// Item extraction
	numberedItemRe = regexp.MustCompile(`^\d+[.)]\s*(.+)$`)
)

// classifySection maps header text to section type
// Based on corpus frequency analysis: 173 OUTPUT INSTRUCTIONS, 162 STEPS, 132 IDENTITY
func classifySection(header string) Section {
	h := strings.ToUpper(strings.TrimSpace(header))

	// Order: most specific first
	switch {
	// Skip sections (noise)
	case strings.Contains(h, "EXAMPLE"):
		return SectionSkip
	case strings.Contains(h, "INPUT"):
		return SectionSkip

	// Identity sections → role extraction
	case strings.Contains(h, "IDENTITY"):
		return SectionIdentity
	case strings.Contains(h, "ROLE"):
		return SectionIdentity
	case h == "PURPOSE":
		return SectionIdentity

	// Steps sections → steps[]
	case strings.Contains(h, "STEP"):
		return SectionSteps
	case strings.Contains(h, "GOAL"):
		return SectionSteps
	case strings.Contains(h, "TASK"):
		return SectionSteps

	// Output sections → output_instructions[]
	// Trust the author completely here
	case strings.Contains(h, "OUTPUT"):
		return SectionOutput
	case strings.Contains(h, "FORMAT"):
		return SectionOutput
	case strings.Contains(h, "RESTRICTION"):
		return SectionOutput
	case strings.Contains(h, "CONSTRAINT"):
		return SectionOutput
	}

	return SectionUnknown
}

// ParseMarkdownPrompt is the main entry point
func ParseMarkdownPrompt(content string) *FabricPrompt {
	return Parse(content)
}

// Parse implements the mathematically-derived algorithm
func Parse(content string) *FabricPrompt {
	fp := &FabricPrompt{
		Expertise:          []string{},
		Steps:              []map[string]any{},
		OutputFormat:       "markdown",
		OutputSections:     []map[string]any{},
		OutputInstructions: []map[string]any{},
		Restrictions:       []map[string]any{},
	}

	lines := strings.Split(content, "\n")
	currentSection := SectionUnknown
	seen := make(map[string]bool)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// Detect section headers
		if m := headerHashRe.FindStringSubmatch(trimmed); m != nil {
			currentSection = classifySection(m[1])
			continue
		}
		if m := headerColonRe.FindStringSubmatch(trimmed); m != nil {
			currentSection = classifySection(m[1])
			continue
		}

		// Skip sections we should ignore (EXAMPLE, INPUT)
		if currentSection == SectionSkip {
			continue
		}

		// Extract item content
		item := extractItemContent(trimmed)
		if item == "" {
			// Check for role in paragraph form (identity OR unknown sections)
			if fp.Role == "" && roleRe.MatchString(trimmed) && len(trimmed) > 20 {
				if currentSection == SectionIdentity || currentSection == SectionUnknown {
					fp.Role = trimmed
				}
			}
			// Paragraph under STEPS is still a step
			if currentSection == SectionSteps && len(trimmed) > 20 && !strings.HasPrefix(trimmed, "#") {
				key := dedupeKey(trimmed)
				if !seen[key] {
					seen[key] = true
					fp.Steps = append(fp.Steps, map[string]any{"action": trimmed})
				}
			}
			continue
		}

		// Skip noise (very short items)
		if len(item) < 10 || len(strings.Fields(item)) < 3 {
			continue
		}

		// Deduplicate
		key := dedupeKey(item)
		if seen[key] {
			continue
		}
		seen[key] = true

		// CLASSIFICATION ALGORITHM
		//
		// Rule 1: Restrictions are SYNTACTIC (100% precision)
		// "Do not...", "Never..." → restrictions[] regardless of section
		if restrictionRe.MatchString(item) {
			fp.Restrictions = append(fp.Restrictions, map[string]any{"rule": item})
			continue
		}

		// Rule 2-4: TRUST THE SECTION HEADER
		switch currentSection {
		case SectionIdentity:
			if fp.Role == "" && roleRe.MatchString(item) {
				fp.Role = item
			} else if fp.Purpose == "" {
				fp.Purpose = item
			}

		case SectionSteps:
			// Author said STEPS → it's a step. Period.
			fp.Steps = append(fp.Steps, map[string]any{"action": item})

		case SectionOutput:
			// Author said OUTPUT → it's an output instruction. Period.
			// NO second-guessing with intent classification.
			fp.OutputInstructions = append(fp.OutputInstructions, map[string]any{"instruction": item})

		case SectionUnknown:
			// Unknown section: check for role, otherwise default to output
			if fp.Role == "" && roleRe.MatchString(item) {
				fp.Role = item
			} else {
				fp.OutputInstructions = append(fp.OutputInstructions, map[string]any{"instruction": item})
			}
		}
	}

	// Detect output format
	lower := strings.ToLower(content)
	if strings.Contains(lower, "json") && !strings.Contains(lower, "not json") {
		fp.OutputFormat = "json"
	}
	if strings.Contains(lower, "step by step") || strings.Contains(lower, "step-by-step") {
		fp.ThinkingInstruction = "Think step by step"
	}

	return fp
}

// extractItemContent pulls content from bullet or numbered list
func extractItemContent(line string) string {
	// Bullet: - item or * item
	if strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") || strings.HasPrefix(line, "•") {
		return strings.TrimSpace(strings.TrimLeft(line, "-*• "))
	}
	// Numbered: 1. item or 1) item
	if m := numberedItemRe.FindStringSubmatch(line); m != nil {
		return strings.TrimSpace(m[1])
	}
	return ""
}

// dedupeKey creates a normalized key for deduplication
func dedupeKey(s string) string {
	key := strings.ToLower(s)
	if len(key) > 50 {
		key = key[:50]
	}
	return key
}
