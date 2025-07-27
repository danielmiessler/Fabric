package openai_compatible

import (
	"testing"
)

func TestCreateClient(t *testing.T) {
	testCases := []struct {
		name     string
		provider string
		exists   bool
	}{
		{
			name:     "Existing provider - Mistral",
			provider: "Mistral",
			exists:   true,
		},
		{
			name:     "Existing provider - Groq",
			provider: "Groq",
			exists:   true,
		},
		{
			name:     "Non-existent provider",
			provider: "NonExistent",
			exists:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, exists := CreateClient(tc.provider)
			if exists != tc.exists {
				t.Errorf("Expected exists=%v for provider %s, got %v",
					tc.exists, tc.provider, exists)
			}
			if exists && client == nil {
				t.Errorf("Expected non-nil client for provider %s", tc.provider)
			}
		})
	}
}
