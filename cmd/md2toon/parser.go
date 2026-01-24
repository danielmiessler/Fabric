package main

import (
	"regexp"
	"strings"
)

var (
	identityHeaders    = []string{"IDENTITY AND PURPOSE", "IDENTITY", "PURPOSE", "IDENTITY & PURPOSE"}
	stepHeaders        = []string{"STEPS", "ACTIONS", "TASK", "PROCESS"}
	outputHeaders      = []string{"OUTPUT INSTRUCTIONS", "OUTPUT", "FORMAT"}
	restrictionHeaders = []string{"RESTRICTIONS", "CONSTRAINTS", "RULES", "LIMITATIONS"}

	headerRe       = regexp.MustCompile(`^#+\s*(.+)$`)
	sentenceRe     = regexp.MustCompile(`(?s)[.!?]\s+`)
	expertiseRe    = regexp.MustCompile(`(?i)(?:expert in|specialize[sd]? in|skilled in|proficient in)\s+([^.]+)`)
	expertiseSplit = regexp.MustCompile(`,\s*(?:and\s+)?|\s+and\s+`)
	bulletRe       = regexp.MustCompile(`^[-*•]\s+(.+)$`)
	numberedRe     = regexp.MustCompile(`^\d+[.)]\s+(.+)$`)
	boldRe         = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	sectionNameRe  = regexp.MustCompile(`(?i)(?:in a section called|under the heading|in a subsection called)\s*["']?([A-Z][A-Z\s_-]+)["']?`)
	bulletLineRe   = regexp.MustCompile(`(?m)^\s*[-*]\s+(.+)$`)
	bulletAllRe    = regexp.MustCompile(`(?m)^\s*[-*•]\s+(.+)$`)

	outputPatterns = []struct {
		re     *regexp.Regexp
		result func([]string) string
	}{
		{regexp.MustCompile(`(?i)output.*markdown`), func(_ []string) string { return "Output in Markdown format" }},
		{regexp.MustCompile(`(?i)output.*json`), func(_ []string) string { return "Output in JSON format" }},
		{regexp.MustCompile(`(?i)do not use.*bold`), func(_ []string) string { return "Do not use bold formatting" }},
		{regexp.MustCompile(`(?i)do not use.*italic`), func(_ []string) string { return "Do not use italic formatting" }},
		{regexp.MustCompile(`(?i)use bulleted lists`), func(_ []string) string { return "Use bulleted lists" }},
		{regexp.MustCompile(`(?i)(\d+)\s*words?\s*(?:or\s*)?(?:less|max)`), func(m []string) string { return "Maximum " + m[1] + " words" }},
		{regexp.MustCompile(`(?i)(\d+)\s*bullets?`), func(m []string) string { return "Use " + m[1] + " bullet points" }},
	}

	purposeKeywords = []string{"purpose", "goal", "aim", "objective", "task"}
)

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

func ParseMarkdownPrompt(content string) *FabricPrompt {
	sections := splitSections(content)

	p := &FabricPrompt{
		Expertise:          []string{},
		Steps:              []map[string]any{},
		OutputFormat:       "markdown",
		OutputSections:     []map[string]any{},
		OutputInstructions: []map[string]any{},
		Restrictions:       []map[string]any{},
	}

	for _, key := range identityHeaders {
		if text, ok := sections[key]; ok {
			p.Role, p.Expertise, p.Purpose = parseIdentity(text)
			if hasThinkingInstruction(text) {
				p.ThinkingInstruction = "Think step by step"
			}
			break
		}
	}

	for _, key := range stepHeaders {
		if text, ok := sections[key]; ok {
			p.Steps = parseSteps(text)
			break
		}
	}

	for _, key := range outputHeaders {
		if text, ok := sections[key]; ok {
			p.OutputSections, p.OutputInstructions = parseOutput(text)
			break
		}
	}

	for _, key := range restrictionHeaders {
		if text, ok := sections[key]; ok {
			p.Restrictions = parseRestrictions(text)
			break
		}
	}

	return p
}

func splitSections(content string) map[string]string {
	sections := make(map[string]string)
	var header string
	var lines []string

	for _, line := range strings.Split(content, "\n") {
		if m := headerRe.FindStringSubmatch(line); m != nil {
			if header != "" {
				sections[header] = strings.TrimSpace(strings.Join(lines, "\n"))
			}
			header = strings.ToUpper(strings.TrimSpace(m[1]))
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}
	if header != "" {
		sections[header] = strings.TrimSpace(strings.Join(lines, "\n"))
	}
	return sections
}

func hasThinkingInstruction(text string) bool {
	lower := strings.ToLower(text)
	return strings.Contains(lower, "step by step") || strings.Contains(lower, "think step")
}

func parseIdentity(text string) (role string, expertise []string, purpose string) {
	parts := sentenceRe.Split(strings.TrimSpace(text), -1)
	var sentences []string
	for _, s := range parts {
		if s = strings.TrimSpace(s); s != "" {
			sentences = append(sentences, s)
		}
	}
	if len(sentences) == 0 {
		return
	}

	first := sentences[0]
	if strings.Contains(strings.ToLower(first), "you are") || strings.HasPrefix(first, "You") {
		role = first
		expertise = extractExpertise(first)
	}

	for _, sentence := range sentences[1:] {
		lower := strings.ToLower(sentence)
		for _, kw := range purposeKeywords {
			if strings.Contains(lower, kw) {
				purpose = sentence
				break
			}
		}
		if strings.Contains(lower, "specialize") {
			expertise = append(expertise, extractExpertise(sentence)...)
		}
	}

	if purpose == "" && len(sentences) > 1 {
		purpose = sentences[len(sentences)-1]
	}
	return
}

func extractExpertise(text string) []string {
	var result []string
	for _, match := range expertiseRe.FindAllStringSubmatch(text, -1) {
		for _, item := range expertiseSplit.Split(match[1], -1) {
			if item = strings.TrimSpace(item); item != "" {
				result = append(result, item)
			}
		}
	}
	return result
}

func parseSteps(text string) []map[string]any {
	var steps []map[string]any
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if m := bulletRe.FindStringSubmatch(line); m != nil {
			steps = append(steps, map[string]any{"action": stripBold(m[1])})
		} else if m := numberedRe.FindStringSubmatch(line); m != nil {
			steps = append(steps, map[string]any{"action": strings.TrimSpace(m[1])})
		}
	}
	if len(steps) == 0 {
		for i, s := range strings.Split(text, ".") {
			if s = strings.TrimSpace(s); len(s) > 10 && i < 5 {
				steps = append(steps, map[string]any{"action": s})
			}
		}
	}
	return steps
}

func stripBold(text string) string {
	return strings.TrimSpace(boldRe.ReplaceAllString(text, "$1"))
}

func parseOutput(text string) ([]map[string]any, []map[string]any) {
	var sections []map[string]any
	for _, m := range sectionNameRe.FindAllStringSubmatch(text, -1) {
		name := strings.TrimRight(strings.TrimSpace(m[1]), ":")
		sections = append(sections, map[string]any{"name": name, "description": "Output section: " + name})
	}

	var instructions []map[string]any
	for _, p := range outputPatterns {
		if m := p.re.FindStringSubmatch(text); m != nil {
			instructions = append(instructions, map[string]any{"instruction": p.result(m)})
		}
	}
	for _, m := range bulletLineRe.FindAllStringSubmatch(text, -1) {
		if bullet := strings.TrimSpace(m[1]); len(bullet) < 100 {
			instructions = append(instructions, map[string]any{"instruction": bullet})
		}
	}
	return sections, instructions
}

func parseRestrictions(text string) []map[string]any {
	var restrictions []map[string]any
	for _, m := range bulletAllRe.FindAllStringSubmatch(text, -1) {
		restrictions = append(restrictions, map[string]any{"rule": stripBold(m[1])})
	}
	return restrictions
}
