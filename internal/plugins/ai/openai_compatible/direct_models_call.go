package openai_compatible

import (
	"context"

	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

// DirectlyGetModels is used to fetch models directly from the API when the
// standard OpenAI SDK method fails due to a nonstandard format.
func (c *Client) DirectlyGetModels(ctx context.Context) ([]string, error) {
	return openai.FetchModelsDirectly(ctx, c.ApiBaseURL.Value, c.ApiKey.Value, c.GetName())
}
