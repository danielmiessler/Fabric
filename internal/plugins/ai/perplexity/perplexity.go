// Package perplexity connects Fabric to the Perplexity Agent API. The Agent
// API uses the OpenAI Responses format, so the OpenAI client does the work.
package perplexity

import (
	"context"
	"maps"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaisdk "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/openai/openai-go/v3/responses"
)

// presets are the Agent API presets. Fabric shows them as models.
var presets = []string{"fast", "low", "medium", "high", "xhigh"}

// legacyModels maps the Sonar API model names to Agent API names, as the
// Perplexity migration guide recommends. Old configurations continue to work.
var legacyModels = map[string]string{
	"sonar":               "perplexity/sonar",
	"sonar-pro":           "fast",
	"sonar-reasoning-pro": "low",
	"sonar-deep-research": "high",
}

type Client struct {
	*openai.Client
}

func NewClient() *Client {
	c := &Client{openai.NewClientCompatibleWithResponses("Perplexity", "https://api.perplexity.ai/v1", true, nil)}
	c.SetWebSearchToolName("web_search")
	c.SetResponseParamsHook(agentParams)
	return c
}

// ListModels adds the presets to the models of the Agent API.
func (c *Client) ListModels(ctx context.Context) ([]string, error) {
	models, err := c.Client.ListModels(ctx)
	return append(models, presets...), err
}

// anthropicMaxTokens is the output limit for Anthropic models when the user
// does not set one. The Agent API requires a limit for these models. The
// value is the same as the Fabric Anthropic default.
const anthropicMaxTokens = 4096

// agentParams changes a request for the Agent API. A preset goes in the
// "preset" field and the model stays empty. Presets set their own sampling,
// and some presets reject temperature, so Fabric does not send it with a
// preset. The request always has the web_search tool, because Sonar always
// searched the web.
func agentParams(p *responses.ResponseNewParams) {
	model := string(p.Model)
	if name, ok := legacyModels[model]; ok {
		model = name
	}
	p.Model = responses.ResponsesModel(model)
	if slices.Contains(presets, model) {
		p.Model = ""
		extra := maps.Clone(p.ExtraFields())
		if extra == nil {
			extra = map[string]any{}
		}
		extra["preset"] = model
		p.SetExtraFields(extra)
		p.Temperature, p.TopP = param.Opt[float64]{}, param.Opt[float64]{}
	}
	if strings.HasPrefix(model, "anthropic/") && !p.MaxOutputTokens.Valid() {
		p.MaxOutputTokens = openaisdk.Int(anthropicMaxTokens)
	}
	if !slices.ContainsFunc(p.Tools, func(t responses.ToolUnionParam) bool { return t.OfWebSearch != nil }) {
		p.Tools = append(p.Tools, responses.ToolParamOfWebSearch("web_search"))
	}
}
