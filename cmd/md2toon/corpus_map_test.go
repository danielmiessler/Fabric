package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestBuildCorpusMap(t *testing.T) {
	cm, err := BuildCorpusMap("../../data/patterns")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total unique verbs discovered: %d", len(cm.Verbs))

	// Show top verbs by section
	t.Log("\n=== CORPUS-DERIVED VERB FREQUENCIES ===")

	// Collect and sort by frequency
	type verbFreq struct {
		verb  string
		stats *VerbStats
		total int
	}

	var all []verbFreq
	for v, s := range cm.Verbs {
		total := s.InSteps + s.InOutput + s.InRestrict + s.InIdentity
		if total >= 3 { // Only verbs appearing 3+ times
			all = append(all, verbFreq{v, s, total})
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].total > all[j].total
	})

	t.Log("\nTop 20 verbs by total frequency:")
	for i, vf := range all {
		if i >= 20 {
			break
		}
		classification := cm.ClassifyVerb(vf.verb)
		confidence := cm.Confidence(vf.verb)
		t.Logf("  %-15s total=%3d steps=%3d output=%3d restrict=%3d -> %s (%.0f%%)",
			vf.verb, vf.total, vf.stats.InSteps, vf.stats.InOutput, vf.stats.InRestrict,
			classification, confidence*100)
	}

	// Show verbs that are strong indicators
	t.Log("\n=== STRONG INDICATORS (>80% confidence) ===")

	t.Log("\nSTEPS indicators:")
	for _, vf := range all {
		if cm.ClassifyVerb(vf.verb) == "steps" && cm.Confidence(vf.verb) > 0.8 && vf.stats.InSteps >= 5 {
			t.Logf("  %s (%d occurrences, %.0f%% confidence)", vf.verb, vf.stats.InSteps, cm.Confidence(vf.verb)*100)
		}
	}

	t.Log("\nOUTPUT indicators:")
	for _, vf := range all {
		if cm.ClassifyVerb(vf.verb) == "output" && cm.Confidence(vf.verb) > 0.8 && vf.stats.InOutput >= 5 {
			t.Logf("  %s (%d occurrences, %.0f%% confidence)", vf.verb, vf.stats.InOutput, cm.Confidence(vf.verb)*100)
		}
	}

	t.Log("\nRESTRICT indicators:")
	for _, vf := range all {
		if cm.ClassifyVerb(vf.verb) == "restrict" && cm.Confidence(vf.verb) > 0.8 && vf.stats.InRestrict >= 3 {
			t.Logf("  %s (%d occurrences, %.0f%% confidence)", vf.verb, vf.stats.InRestrict, cm.Confidence(vf.verb)*100)
		}
	}
}

func TestCompareStaticVsDynamic(t *testing.T) {
	cm, err := BuildCorpusMap("../../data/patterns")
	if err != nil {
		t.Fatal(err)
	}

	// Test cases - verbs we expect to classify correctly
	testCases := []struct {
		verb     string
		expected string
	}{
		{"extract", "steps"},
		{"create", "steps"},
		{"write", "steps"},
		{"output", "output"},
		{"do", "restrict"},
		{"never", "restrict"},
		{"use", "output"},
	}

	t.Log("\n=== STATIC vs DYNAMIC CLASSIFICATION ===")
	for _, tc := range testCases {
		dynamic := cm.ClassifyVerb(tc.verb)
		confidence := cm.Confidence(tc.verb)
		stats := cm.Verbs[tc.verb]

		match := "✓"
		if dynamic != tc.expected {
			match = "✗"
		}

		if stats != nil {
			t.Logf("%s %-10s expected=%-10s dynamic=%-10s conf=%.0f%% (steps=%d output=%d restrict=%d)",
				match, tc.verb, tc.expected, dynamic, confidence*100,
				stats.InSteps, stats.InOutput, stats.InRestrict)
		} else {
			t.Logf("%s %-10s expected=%-10s dynamic=%-10s (not in corpus)",
				match, tc.verb, tc.expected, dynamic)
		}
	}
}

func TestGenerateVerbList(t *testing.T) {
	cm, err := BuildCorpusMap("../../data/patterns")
	if err != nil {
		t.Fatal(err)
	}

	// Generate Go code for a static verb list based on corpus
	t.Log("\n=== GENERATED VERB LISTS FROM CORPUS ===")

	var stepsVerbs, outputVerbs, restrictVerbs []string

	for verb, stats := range cm.Verbs {
		total := stats.InSteps + stats.InOutput + stats.InRestrict
		if total < 3 {
			continue
		}

		section := cm.ClassifyVerb(verb)
		conf := cm.Confidence(verb)

		if conf >= 0.6 { // 60%+ confidence
			switch section {
			case "steps":
				stepsVerbs = append(stepsVerbs, verb)
			case "output":
				outputVerbs = append(outputVerbs, verb)
			case "restrict":
				restrictVerbs = append(restrictVerbs, verb)
			}
		}
	}

	sort.Strings(stepsVerbs)
	sort.Strings(outputVerbs)
	sort.Strings(restrictVerbs)

	fmt.Printf("\nstepsVerbs := []string{%q}\n", stepsVerbs)
	fmt.Printf("\noutputVerbs := []string{%q}\n", outputVerbs)
	fmt.Printf("\nrestrictVerbs := []string{%q}\n", restrictVerbs)
}
