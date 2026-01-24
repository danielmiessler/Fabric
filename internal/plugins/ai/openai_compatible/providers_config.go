package openai_compatible

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

// ProviderConfig defines the configuration for an OpenAI-compatible API provider
type ProviderConfig struct {
	Name                string
	BaseURL             string
	ModelsURL           string // Optional: Custom endpoint for listing models (if different from BaseURL/models)
	ImplementsResponses bool   // Whether the provider supports OpenAI's new Responses API
}

// Client is the common structure for all OpenAI-compatible providers
type Client struct {
	*openai.Client
	modelsURL string // Custom URL for listing models (if different from BaseURL/models)
}

// NewClient creates a new OpenAI-compatible client for the specified provider
func NewClient(providerConfig ProviderConfig) *Client {
	client := &Client{
		modelsURL: providerConfig.ModelsURL,
	}
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
	// If a custom models URL is provided, handle it
	if c.modelsURL != "" {
		// Check for static model list
		if strings.HasPrefix(c.modelsURL, "static:") {
			return c.getStaticModels(c.modelsURL)
		}
		// TODO: Handle context properly in Fabric by accepting and propagating a context.Context
		// instead of creating a new one here.
		return openai.FetchModelsDirectly(context.Background(), c.modelsURL, c.Client.ApiKey.Value, c.GetName(), nil)
	}

	// First try the standard OpenAI SDK approach
	models, err := c.Client.ListModels()
	if err == nil && len(models) > 0 { // only return if OpenAI SDK returns models
		return models, nil
	}

	// Fall back to direct API fetch
	return c.DirectlyGetModels(context.Background())
}

// NeedsRawMode overrides the parent implementation to handle provider-specific raw mode requirements
func (c *Client) NeedsRawMode(modelName string) bool {
	// MiniMax models require raw mode for proper message formatting
	if c.GetName() == "MiniMax" {
		return true
	}
	// Fall back to parent OpenAI client implementation for other providers
	return c.Client.NeedsRawMode(modelName)
}

// getStaticModels returns a predefined list of models for providers that don't support model discovery
func (c *Client) getStaticModels(modelsKey string) ([]string, error) {
	switch modelsKey {
	case "static:abacus":
		return []string{
			"route-llm",
			"gpt-4o-2024-11-20",
			"gpt-4o-mini",
			"o4-mini",
			"o3-pro",
			"o3",
			"o3-mini",
			"gpt-4.1",
			"gpt-4.1-mini",
			"gpt-4.1-nano",
			"gpt-5",
			"gpt-5-mini",
			"gpt-5-nano",
			"gpt-5.1",
			"gpt-5.1-chat-latest",
			"openai/gpt-oss-120b",
			"claude-3-7-sonnet-20250219",
			"claude-sonnet-4-20250514",
			"claude-opus-4-20250514",
			"claude-opus-4-1-20250805",
			"claude-sonnet-4-5-20250929",
			"claude-haiku-4-5-20251001",
			"claude-opus-4-5-20251101",
			"meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8",
			"meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo",
			"meta-llama/Meta-Llama-3.1-70B-Instruct",
			"meta-llama/Meta-Llama-3.1-8B-Instruct",
			"llama-3.3-70b-versatile",
			"gemini-2.0-flash-001",
			"gemini-2.0-pro-exp-02-05",
			"gemini-2.5-pro",
			"gemini-2.5-flash",
			"gemini-3-pro-preview",
			"qwen-2.5-coder-32b",
			"Qwen/Qwen2.5-72B-Instruct",
			"Qwen/QwQ-32B",
			"Qwen/Qwen3-235B-A22B-Instruct-2507",
			"Qwen/Qwen3-32B",
			"qwen/qwen3-coder-480b-a35b-instruct",
			"qwen/qwen3-Max",
			"grok-4-0709",
			"grok-4-fast-non-reasoning",
			"grok-4-1-fast-non-reasoning",
			"grok-code-fast-1",
			"kimi-k2-turbo-preview",
			"deepseek/deepseek-v3.1",
			"deepseek-ai/DeepSeek-V3.1-Terminus",
			"deepseek-ai/DeepSeek-R1",
			"deepseek-ai/DeepSeek-V3.2",
			"zai-org/glm-4.5",
			"zai-org/glm-4.6",
		}, nil
	case "static:minimax":
		return []string{
			"MiniMax-M2",
			"MiniMax-M2.1",
			"MiniMax-M2.1-lightning",
		}, nil
	default:
		return nil, fmt.Errorf("unknown static model list: %s", modelsKey)
	}
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
	"GitHub": {
		Name:                "GitHub",
		BaseURL:             "https://models.github.ai/inference",
		ModelsURL:           "https://models.github.ai/catalog", // FetchModelsDirectly will append /models
		ImplementsResponses: false,
	},
	"Infermatic": {
		Name:                "Infermatic",
		BaseURL:             "https://api.totalgpt.ai/v1",
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
	"MiniMax": {
		Name:                "MiniMax",
		BaseURL:             "https://api.minimax.io/v1",
		ModelsURL:           "static:minimax",
		ImplementsResponses: false,
	},
	"Mistral": {
		Name:                "Mistral",
		BaseURL:             "https://api.mistral.ai/v1",
		ImplementsResponses: false,
	},
	"Novita AI": {
		Name:                "Novita AI",
		BaseURL:             "https://api.novita.ai/openai/v1",
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
	"Z AI": {
		Name:                "Z AI",
		BaseURL:             "https://api.z.ai/api/paas/v4",
		ImplementsResponses: false,
	},
	"Abacus": {
		Name:                "Abacus",
		BaseURL:             "https://routellm.abacus.ai/v1/",
		ModelsURL:           "static:abacus", // Special marker for static model list
		ImplementsResponses: false,
	},
	"Mammouth": {
		Name:                "Mammouth",
		BaseURL:             "https://api.mammouth.ai/v1",
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
