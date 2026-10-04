package azurecommon

import (
	"testing"
)

func TestParseDeployments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"single deployment", "gpt-4o", []string{"gpt-4o"}},
		{"multiple deployments", "gpt-4o,gpt-5", []string{"gpt-4o", "gpt-5"}},
		{"with spaces", " gpt-4o , gpt-5 ", []string{"gpt-4o", "gpt-5"}},
		{"empty string", "", nil},
		{"only commas", ",,", nil},
		{"trailing comma", "gpt-4o,", []string{"gpt-4o"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseDeployments(tt.input)
			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d deployments, got %d", len(tt.expected), len(result))
			}
			for i, d := range tt.expected {
				if result[i] != d {
					t.Errorf("expected deployment %q, got %q", d, result[i])
				}
			}
		})
	}
}
