package openai_compatible

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			name:     "Existing provider - Z AI",
			provider: "Z AI",
			exists:   true,
		},
		{
			name:     "Existing provider - Abacus",
			provider: "Abacus",
			exists:   true,
		},
		{
			name:     "Existing provider - Infermatic",
			provider: "Infermatic",
			exists:   true,
		},
		{
			name:     "Existing provider - MiniMax",
			provider: "MiniMax",
			exists:   true,
		},
		{
			name:     "Existing provider - OpenCode Zen",
			provider: "OpenCode Zen",
			exists:   true,
		},
		{
			name:     "Existing provider - OpenCode Go",
			provider: "OpenCode Go",
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

// Ensures both OpenCode providers carry the session-routing header and a
// client-specific User-Agent required by OpenCode Go/Zen.
func TestOpenCodeProvidersConfigureSessionRouting(t *testing.T) {
	for _, name := range []string{"OpenCode Zen", "OpenCode Go"} {
		provider, found := GetProviderByName(name)
		assert.True(t, found, "provider %s should exist", name)
		assert.Equal(t, "x-opencode-session", provider.SessionHeader, "provider %s session header", name)
		assert.NotEmpty(t, provider.UserAgent, "provider %s user agent", name)
	}
}
