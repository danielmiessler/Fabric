package openai_compatible

import (
	"context"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

// ProviderConfig defines the configuration for an OpenAI-compatible API provider
type ProviderConfig struct {
	Name                string
	BaseURL             string
	ImplementsResponses bool // Whether the provider supports OpenAI's new Responses API
}

// Client is the common structure for all OpenAI-compatible providers
type Client struct {
	*openai.Client
}

// NewClient creates a new OpenAI-compatible client for the specified provider
func NewClient(providerConfig ProviderConfig) *Client {
	client := &Client{}
	client.Client = openai.NewClientCompatibleWithResponses(
		providerConfig.Name,
		providerConfig.BaseURL,
		providerConfig.ImplementsResponses,
		nil,
	)
	return client
}

// ListModels overrides the default ListModels to handle different response formats
func (c *Client) ListModels() ([]string, error) {
	// First try the standard OpenAI SDK approach
	models, err := c.Client.ListModels()
	if err == nil && len(models) > 0 { // only return if OpenAI SDK returns models
		return models, nil
	}

	// TODO: Handle context properly in Fabric by accepting and propagating a context.Context
	// instead of creating a new one here.
	return c.DirectlyGetModels(context.Background())
}

// ProviderMap is a map of provider name to ProviderConfig for O(1) lookup
var ProviderMap = map[string]ProviderConfig{
	"AIML": {
		Name:                "AIML",
		BaseURL:             "https://api.aimlapi.com/v1",
		ImplementsResponses: false,
	},
	"Cerebras": {
		Name:                "Cerebras",
		BaseURL:             "https://api.cerebras.ai/v1",
		ImplementsResponses: false,
	},
	"DeepSeek": {
		Name:                "DeepSeek",
		BaseURL:             "https://api.deepseek.com",
		ImplementsResponses: false,
	},
	"GrokAI": {
		Name:                "GrokAI",
		BaseURL:             "https://api.x.ai/v1",
		ImplementsResponses: false,
	},
	"Groq": {
		Name:                "Groq",
		BaseURL:             "https://api.groq.com/openai/v1",
		ImplementsResponses: false,
	},
	"Langdock": {
		Name:                "Langdock",
		BaseURL:             "https://api.langdock.com/openai/{{REGION=us}}/v1",
		ImplementsResponses: false,
	},
	"LiteLLM": {
		Name:                "LiteLLM",
		BaseURL:             "http://localhost:4000",
		ImplementsResponses: false,
	},
	"Mistral": {
		Name:                "Mistral",
		BaseURL:             "https://api.mistral.ai/v1",
		ImplementsResponses: false,
	},
	"OpenRouter": {
		Name:                "OpenRouter",
		BaseURL:             "https://openrouter.ai/api/v1",
		ImplementsResponses: false,
	},
	"SiliconCloud": {
		Name:                "SiliconCloud",
		BaseURL:             "https://api.siliconflow.cn/v1",
		ImplementsResponses: false,
	},
	"Together": {
		Name:                "Together",
		BaseURL:             "https://api.together.xyz/v1",
		ImplementsResponses: false,
	},
	"Venice AI": {
		Name:                "Venice AI",
		BaseURL:             "https://api.venice.ai/api/v1",
		ImplementsResponses: false,
	},
}

// GetProviderByName returns the provider configuration for a given name with O(1) lookup
func GetProviderByName(name string) (ProviderConfig, bool) {
	provider, found := ProviderMap[name]
	if strings.Contains(provider.BaseURL, "{{") && strings.Contains(provider.BaseURL, "}}") {
		// Extract the template variable and default value
		start := strings.Index(provider.BaseURL, "{{")
		end := strings.Index(provider.BaseURL, "}}") + 2
		template := provider.BaseURL[start:end]

		// Parse the template to get variable name and default value
		inner := template[2 : len(template)-2] // Remove {{ and }}
		parts := strings.Split(inner, "=")
		if len(parts) == 2 {
			varName := strings.TrimSpace(parts[0])
			defaultValue := strings.TrimSpace(parts[1])

			// Create environment variable name
			envVarName := strings.ToUpper(provider.Name) + "_" + varName

			// Get value from environment or use default
			envValue := os.Getenv(envVarName)
			if envValue == "" {
				envValue = defaultValue
			}

			// Replace the template with the actual value
			provider.BaseURL = strings.Replace(provider.BaseURL, template, envValue, 1)
		}
	}
	return provider, found
}

// CreateClient creates a new client for a provider by name
func CreateClient(providerName string) (*Client, bool) {
	providerConfig, found := GetProviderByName(providerName)
	if !found {
		return nil, false
	}
	return NewClient(providerConfig), true
}
