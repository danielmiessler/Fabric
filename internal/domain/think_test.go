package domain

import "testing"

func TestStripThinkBlocks(t *testing.T) {
	input := "<think>internal</think>\n\nresult"
	got := StripThinkBlocks(input, "<think>", "</think>")
	if got != "result" {
		t.Errorf("expected %q, got %q", "result", got)
	}
}

func TestStripThinkBlocksCustomTags(t *testing.T) {
	input := "[[t]]hidden[[/t]] visible"
	got := StripThinkBlocks(input, "[[t]]", "[[/t]]")
	if got != "visible" {
		t.Errorf("expected %q, got %q", "visible", got)
	}
}
