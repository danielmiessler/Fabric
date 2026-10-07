package openai_compatible

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

const abacusRouteLLMModelsURL = "https://routellm.abacus.ai/api/v0/_listRouteLLMModels"

// openCodeSessionHeader carries a stable per-conversation session ID that
// OpenCode uses to optimize routing and prompt caching.
const openCodeSessionHeader = "x-opencode-session"

// ProviderConfig defines the configuration for an OpenAI-compatible API provider
type ProviderConfig struct {
	Name                string
	BaseURL             string
	ModelsURL           string // Optional: Custom endpoint for listing models (if different from BaseURL/models)
	ImplementsResponses bool   // Whether the provider supports OpenAI's new Responses API
	// WebSearchToolName overrides the default "web_search_preview" tool name
	// emitted on the Responses API when Search is enabled. Leave empty to keep
	// the OpenAI default. xAI, for example, requires "web_search".
	WebSearchToolName string
	// EnableXSearch, when true, also appends an xAI "x_search" tool entry
	// alongside the web search tool when Search is enabled. Non-xAI
	// providers should leave this false.
	EnableXSearch bool
	// SessionHeader, when non-empty, is the request header used to carry a
	// stable per-conversation session ID. OpenCode, for example, expects
	// "x-opencode-session" so it can optimize routing and prompt caching.
	SessionHeader string
	// UserAgent, when non-empty, overrides the SDK's default User-Agent
	// header. Providers that ask clients to identify themselves (e.g.
	// OpenCode) should set this.
	UserAgent string
	// ApiKeyOptional makes the API key not required at setup. Local servers
	// such as Apple's fm serve accept requests without a key.
	ApiKeyOptional bool
}

// Client is the common structure for all OpenAI-compatible providers
type Client struct {
	*openai.Client
	modelsURL string // endpoint URL or a "static:" key
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
	client.Client.SetWebSearchToolName(providerConfig.WebSearchToolName)
	client.Client.SetEnableXSearch(providerConfig.EnableXSearch)
	client.Client.SetSessionHeaderName(providerConfig.SessionHeader)
	client.Client.SetUserAgent(providerConfig.UserAgent)
	client.Client.ApiKey.Required = !providerConfig.ApiKeyOptional
	return client
}

// IsConfigured reports whether the vendor is ready to use. A provider with an
// optional API key has no setting that can fail, so it would always count as
// configured. For those, require the base URL in the environment, which
// fabric -S writes when the user selects the provider. Ollama uses the same rule.
func (c *Client) IsConfigured() bool {
	if c.ApiKey.Required {
		return c.Client.IsConfigured()
	}
	return os.Getenv(c.ApiBaseURL.EnvVariable) != ""
}

// ListModels overrides the default ListModels to handle different response formats
func (c *Client) ListModels(ctx context.Context) ([]string, error) {
	if c.modelsURL != "" {
		if c.modelsURL == "static:abacus" {
			models, err := c.fetchAbacusModels()
			if err == nil && len(models) > 0 {
				return models, nil
			}
			return c.getStaticModels(c.modelsURL)
		}

		if strings.HasPrefix(c.modelsURL, "static:") {
			return c.getStaticModels(c.modelsURL)
		}
		// TODO: pass ctx instead of context.Background().
		return openai.FetchModelsDirectly(context.Background(), c.modelsURL, c.Client.ApiKey.Value, c.GetName(), nil)
	}

	models, err := c.Client.ListModels(ctx)
	if err == nil && len(models) > 0 {
		return models, nil
	}

	return c.DirectlyGetModels(ctx)
}

func (c *Client) fetchAbacusModels() ([]string, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, abacusRouteLLMModelsURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.Client.ApiKey.Value != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Client.ApiKey.Value))
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(i18n.T("abacus_models_endpoint_status"), resp.StatusCode)
	}

	var response struct {
		Result []struct {
			Name string `json:"name"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	models := make([]string, 0, len(response.Result))
	for _, item := range response.Result {
		if item.Name != "" {
			models = append(models, item.Name)
		}
	}

	return models, nil
}

// NeedsRawMode overrides the parent implementation to handle provider-specific raw mode requirements
func (c *Client) NeedsRawMode(modelName string) bool {
	// MiniMax models need raw mode for correct message formatting.
	if c.GetName() == "MiniMax" {
		return true
	}
	return c.Client.NeedsRawMode(modelName)
}

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
			"gpt-5-codex",
			"gpt-5.1",
			"gpt-5.1-codex",
			"gpt-5.1-codex-max",
			"gpt-5.1-chat-latest",
			"gpt-5.2",
			"gpt-5.2-chat-latest",
			"gpt-5.2-codex",
			"openai/gpt-oss-120b",
			"claude-sonnet-4-20250514",
			"claude-opus-4-20250514",
			"claude-opus-4-1-20250805",
			"claude-sonnet-4-5-20250929",
			"claude-haiku-4-5-20251001",
			"claude-opus-4-5-20251101",
			"claude-opus-4-6",
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
			"gemini-3-flash-preview",
			"qwen-2.5-coder-32b",
			"Qwen/Qwen2.5-72B-Instruct",
			"Qwen/QwQ-32B",
			"Qwen/Qwen3-235B-A22B-Instruct-2507",
			"Qwen/Qwen3-32B",
			"qwen/qwen3-coder-480b-a35b-instruct",
			"qwen3-max",
			"grok-4-0709",
			"grok-4-fast-non-reasoning",
			"grok-4-1-fast-non-reasoning",
			"grok-code-fast-1",
			"kimi-k2-turbo-preview",
			"kimi-k2.5",
			"deepseek/deepseek-v3.1",
			"deepseek-ai/DeepSeek-V3.1-Terminus",
			"deepseek-ai/DeepSeek-R1",
			"deepseek-ai/DeepSeek-V3.2",
			"zai-org/glm-4.5",
			"zai-org/glm-4.6",
			"zai-org/glm-4.7",
			"zai-org/glm-5",
		}, nil
	case "static:minimax":
		return []string{
			"MiniMax-M3",
			"MiniMax-M2.7",
			"MiniMax-M2.7-highspeed",
		}, nil
	default:
		return nil, fmt.Errorf(i18n.T("openai_compatible_unknown_static_model_list"), modelsKey)
	}
}

// ProviderMap is a map of provider name to ProviderConfig for O(1) lookup
var ProviderMap = map[string]ProviderConfig{
	"Abacus": {
		Name:                "Abacus",
		BaseURL:             "https://routellm.abacus.ai/v1/",
		ModelsURL:           "static:abacus",
		ImplementsResponses: false,
	},
	"AIML": {
		Name:                "AIML",
		BaseURL:             "https://api.aimlapi.com/v1",
		ImplementsResponses: false,
	},
	"Aliyun DashScope": {
		Name:                "Aliyun DashScope",
		BaseURL:             "https://dashscope.aliyuncs.com/compatible-mode/v1",
		ImplementsResponses: false,
	},
	"Apple Foundation Models": {
		Name:                "Apple Foundation Models",
		BaseURL:             "http://localhost:1976/v1",
		ImplementsResponses: false,
		ApiKeyOptional:      true,
	},
	"API Route": {
		Name:                "API Route",
		BaseURL:             "https://global.api-route.com/v1",
		ImplementsResponses: false,
	},
	"ByteDance Ark": {
		Name:                "ByteDance Ark",
		BaseURL:             "https://ark.cn-beijing.volces.com/api/v3",
		ImplementsResponses: false,
	},
	"Cerebras": {
		Name:                "Cerebras",
		BaseURL:             "https://api.cerebras.ai/v1",
		ImplementsResponses: false,
	},
	"Cheaper Inference": {
		Name:                "Cheaper Inference",
		BaseURL:             "https://api.cheaperinference.com/v1",
		ImplementsResponses: false,
	},
	"DeepSeek": {
		Name:                "DeepSeek",
		BaseURL:             "https://api.deepseek.com",
		ImplementsResponses: false,
	},
	"DemonRoute": {
		Name:                "DemonRoute",
		BaseURL:             "https://api.demonroute.com/v1",
		ImplementsResponses: false,
	},
	"Eden AI": {
		Name:                "Eden AI",
		BaseURL:             "https://api.edenai.run/v3",
		ImplementsResponses: false,
	},
	"FuturMix": {
		Name:                "FuturMix",
		BaseURL:             "https://futurmix.ai/v1",
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
		ImplementsResponses: true,
		// xAI's Responses API uses the "web_search" tool type, not "web_search_preview".
		// It also accepts an "x_search" tool entry.
		WebSearchToolName: "web_search",
		EnableXSearch:     true,
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
	"llmman": {
		Name:                "llmman",
		BaseURL:             "http://localhost:17434/v1",
		ImplementsResponses: false,
	},
	"Mammouth": {
		Name:                "Mammouth",
		BaseURL:             "https://api.mammouth.ai/v1",
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
	"OpenCode Go": {
		Name:                "OpenCode Go",
		BaseURL:             "https://opencode.ai/zen/go/v1",
		ImplementsResponses: false,
		SessionHeader:       openCodeSessionHeader,
		UserAgent:           "fabric",
	},
	"OpenCode Zen": {
		Name:                "OpenCode Zen",
		BaseURL:             "https://opencode.ai/zen/v1",
		ImplementsResponses: false,
		SessionHeader:       openCodeSessionHeader,
		UserAgent:           "fabric",
	},
	"OpenRouter": {
		Name:                "OpenRouter",
		BaseURL:             "https://openrouter.ai/api/v1",
		ImplementsResponses: false,
	},
	"Opper": {
		Name:                "Opper",
		BaseURL:             "https://api.opper.ai/v3/compat",
		ImplementsResponses: false,
	},
	"OrcaRouter": {
		Name:                "OrcaRouter",
		BaseURL:             "https://api.orcarouter.ai/v1",
		ImplementsResponses: false,
	},
	"Pzero": {
		Name:                "Pzero",
		BaseURL:             "https://api.pzero.studio/v1",
		ImplementsResponses: false,
	},
	"Requesty": {
		Name:                "Requesty",
		BaseURL:             "https://router.requesty.ai/v1",
		ImplementsResponses: false,
	},
	"SiliconCloud": {
		Name:                "SiliconCloud",
		BaseURL:             "https://api.siliconflow.cn/v1",
		ImplementsResponses: false,
	},
	"Synthorai": {
		Name:                "Synthorai",
		BaseURL:             "https://synthorai.io/v1",
		ImplementsResponses: false,
	},
	"Together": {
		Name:                "Together",
		BaseURL:             "https://api.together.xyz/v1",
		ImplementsResponses: false,
	},
	"TrustedRouter": {
		Name:                "TrustedRouter",
		BaseURL:             "https://api.trustedrouter.com/v1",
		ImplementsResponses: false,
	},
	"Venice AI": {
		Name:                "Venice AI",
		BaseURL:             "https://api.venice.ai/api/v1",
		ImplementsResponses: false,
	},
	"Y-API": {
		Name:                "Y-API",
		BaseURL:             "https://api.y-api.bestvirtualgoods.com/v1",
		ImplementsResponses: false,
	},
	"Z AI": {
		Name:                "Z AI",
		BaseURL:             "https://api.z.ai/api/paas/v4",
		ImplementsResponses: false,
	},
	"Zhipu AI": {
		Name:                "Zhipu AI",
		BaseURL:             "https://open.bigmodel.cn/api/paas/v4",
		ImplementsResponses: false,
	},
}

// GetProviderByName returns the provider configuration for a given name with O(1) lookup
func GetProviderByName(name string) (ProviderConfig, bool) {
	provider, found := ProviderMap[name]
	if strings.Contains(provider.BaseURL, "{{") && strings.Contains(provider.BaseURL, "}}") {
		// A {{VAR=default}} in BaseURL takes its value from the env var NAME_VAR, for example LANGDOCK_REGION.
		start := strings.Index(provider.BaseURL, "{{")
		end := strings.Index(provider.BaseURL, "}}") + 2
		template := provider.BaseURL[start:end]

		inner := template[2 : len(template)-2]
		parts := strings.Split(inner, "=")
		if len(parts) == 2 {
			varName := strings.TrimSpace(parts[0])
			defaultValue := strings.TrimSpace(parts[1])

			envVarName := strings.ToUpper(provider.Name) + "_" + varName

			envValue := os.Getenv(envVarName)
			if envValue == "" {
				envValue = defaultValue
			}

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
