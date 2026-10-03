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
			name:     "Existing provider - Eden AI",
			provider: "Eden AI",
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
			name:     "Existing provider - OrcaRouter",
			provider: "OrcaRouter",
			exists:   true,
		},
		{
			name:     "Existing provider - MiniMax",
			provider: "MiniMax",
			exists:   true,
		},
		{
			name:     "Existing provider - Cheaper Inference",
			provider: "Cheaper Inference",
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
			name:     "Existing provider - TrustedRouter",
			provider: "TrustedRouter",
			exists:   true,
		},
		{
			name:     "Existing provider - FuturMix",
			provider: "FuturMix",
			exists:   true,
		},
		{
			name:     "New Chinese provider - Aliyun DashScope",
			provider: "Aliyun DashScope",
			exists:   true,
		},
		{
			name:     "New Chinese provider - Zhipu AI",
			provider: "Zhipu AI",
			exists:   true,
		},
		{
			name:     "New Chinese provider - ByteDance Ark",
			provider: "ByteDance Ark",
			exists:   true,
		},
		{
			name:     "Existing provider - llmman",
			provider: "llmman",
			exists:   true,
		},
		{
			name:     "Local provider - Apple Foundation Models",
			provider: "Apple Foundation Models",
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

func TestIsConfiguredOptionalKey(t *testing.T) {
	const envVar = "APPLE_FOUNDATION_MODELS_API_BASE_URL"
	t.Setenv(envVar, "")

	apple, _ := CreateClient("Apple Foundation Models")
	if apple.IsConfigured() {
		t.Fatal("optional-key provider must not be configured until its base URL is in the environment")
	}

	t.Setenv(envVar, "http://localhost:1976/v1")
	if !apple.IsConfigured() {
		t.Fatal("optional-key provider must be configured once its base URL is in the environment")
	}

	groq, _ := CreateClient("Groq")
	if groq.IsConfigured() {
		t.Fatal("required-key provider must not be configured without a key")
	}
}
