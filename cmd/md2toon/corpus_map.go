package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// VerbStats tracks where a verb appears in the corpus
type VerbStats struct {
	InIdentity int
	InSteps    int
	InOutput   int
	InRestrict int
}

// CorpusMap holds verb frequency statistics derived from the corpus
type CorpusMap struct {
	Verbs map[string]*VerbStats
}

var (
	corpusSectionHeaderRe = regexp.MustCompile(`(?i)^#+\s*(.+)$`)
	firstWordRe           = regexp.MustCompile(`(?i)^(\w+)`)
)

// BuildCorpusMap scans the corpus and builds a frequency map of verbs by section type
func BuildCorpusMap(patternsDir string) (*CorpusMap, error) {
	cm := &CorpusMap{
		Verbs: make(map[string]*VerbStats),
	}

	files, err := filepath.Glob(filepath.Join(patternsDir, "*/system.md"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		cm.scanFile(string(content))
	}

	return cm, nil
}

func (cm *CorpusMap) scanFile(content string) {
	lines := strings.Split(content, "\n")
	var currentSection string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Detect section header
		if m := corpusSectionHeaderRe.FindStringSubmatch(trimmed); m != nil {
			header := strings.ToUpper(m[1])
			switch {
			case strings.Contains(header, "IDENTITY") || strings.Contains(header, "ROLE"):
				currentSection = "identity"
			case strings.Contains(header, "STEP") || strings.Contains(header, "TASK") || strings.Contains(header, "ACTION"):
				currentSection = "steps"
			case strings.Contains(header, "OUTPUT") || strings.Contains(header, "FORMAT"):
				currentSection = "output"
			case strings.Contains(header, "RESTRICT") || strings.Contains(header, "CONSTRAINT") || strings.Contains(header, "RULE"):
				currentSection = "restrict"
			default:
				currentSection = "other"
			}
			continue
		}

		// Extract first word from bullets
		if strings.HasPrefix(trimmed, "-") || strings.HasPrefix(trimmed, "*") {
			bullet := strings.TrimLeft(trimmed, "-* ")
			if m := firstWordRe.FindStringSubmatch(bullet); m != nil {
				verb := strings.ToLower(m[1])
				cm.recordVerb(verb, currentSection)
			}
		}
	}
}

func (cm *CorpusMap) recordVerb(verb string, section string) {
	if _, ok := cm.Verbs[verb]; !ok {
		cm.Verbs[verb] = &VerbStats{}
	}

	switch section {
	case "identity":
		cm.Verbs[verb].InIdentity++
	case "steps":
		cm.Verbs[verb].InSteps++
	case "output":
		cm.Verbs[verb].InOutput++
	case "restrict":
		cm.Verbs[verb].InRestrict++
	}
}

// ClassifyVerb returns the most likely section type for a verb based on corpus stats
func (cm *CorpusMap) ClassifyVerb(verb string) string {
	verb = strings.ToLower(verb)
	stats, ok := cm.Verbs[verb]
	if !ok {
		return "unknown"
	}

	// Find max
	max := stats.InSteps
	result := "steps"

	if stats.InOutput > max {
		max = stats.InOutput
		result = "output"
	}
	if stats.InRestrict > max {
		max = stats.InRestrict
		result = "restrict"
	}
	if stats.InIdentity > max {
		result = "identity"
	}

	return result
}

// Confidence returns how confident we are in the classification (0.0 - 1.0)
func (cm *CorpusMap) Confidence(verb string) float64 {
	verb = strings.ToLower(verb)
	stats, ok := cm.Verbs[verb]
	if !ok {
		return 0.0
	}

	total := stats.InIdentity + stats.InSteps + stats.InOutput + stats.InRestrict
	if total == 0 {
		return 0.0
	}

	max := stats.InSteps
	if stats.InOutput > max {
		max = stats.InOutput
	}
	if stats.InRestrict > max {
		max = stats.InRestrict
	}
	if stats.InIdentity > max {
		max = stats.InIdentity
	}

	return float64(max) / float64(total)
}

// TopVerbs returns the most common verbs for each section type
func (cm *CorpusMap) TopVerbs(n int) map[string][]string {
	result := map[string][]string{
		"steps":    {},
		"output":   {},
		"restrict": {},
	}

	// Simple approach: for each verb, add to its dominant section
	for verb, stats := range cm.Verbs {
		total := stats.InSteps + stats.InOutput + stats.InRestrict
		if total < 3 {
			continue // Skip rare verbs
		}

		section := cm.ClassifyVerb(verb)
		if section != "unknown" && section != "identity" {
			result[section] = append(result[section], verb)
		}
	}

	// Truncate to n
	for k, v := range result {
		if len(v) > n {
			result[k] = v[:n]
		}
	}

	return result
}
