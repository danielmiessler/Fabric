package azure

import (
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("Azure", "", ret.configure)
	ret.ApiDeployments = ret.AddSetupQuestionCustom("deployments", true,
		"Enter your Azure deployments (comma separated)")
	ret.ApiVersion = ret.AddSetupQuestionCustom("API Version", false,
		"Enter the Azure API version (optional)")

	return
}

type Client struct {
	*openai.Client
	ApiDeployments *plugins.SetupQuestion
	ApiVersion     *plugins.SetupQuestion

	apiDeployments []string
}

func (oi *Client) configure() (err error) {
	oi.apiDeployments = strings.Split(oi.ApiDeployments.Value, ",")
	opts := []option.RequestOption{option.WithAPIKey(oi.ApiKey.Value)}
	if oi.ApiBaseURL.Value != "" {
		opts = append(opts, option.WithBaseURL(oi.ApiBaseURL.Value))
	}
	if oi.ApiVersion.Value != "" {
		opts = append(opts, option.WithQuery("api-version", oi.ApiVersion.Value))
	}
	client := openaiapi.NewClient(opts...)
	oi.ApiClient = &client
	return
}

func (oi *Client) ListModels() (ret []string, err error) {
	ret = oi.apiDeployments
	return
}

func (oi *Client) NeedsRawMode(modelName string) bool {
	return false
}
