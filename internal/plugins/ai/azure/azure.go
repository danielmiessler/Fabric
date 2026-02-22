package azure

import (
	"errors"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/azurecommon"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/azure"
	"github.com/openai/openai-go/option"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("Azure", "", ret.configure)
	ret.ApiDeployments = ret.AddSetupQuestionCustom("deployments", true,
		i18n.T("azure_deployments_question"))
	ret.ApiVersion = ret.AddSetupQuestionCustom("API Version", false,
		i18n.T("azure_api_version_question"))

	return
}

type Client struct {
	*openai.Client
	ApiDeployments *plugins.SetupQuestion
	ApiVersion     *plugins.SetupQuestion

	apiDeployments []string
}

func (oi *Client) configure() error {
	oi.apiDeployments = azurecommon.ParseDeployments(oi.ApiDeployments.Value)
	if len(oi.apiDeployments) == 0 {
		return errors.New(i18n.T("azure_deployments_required"))
	}

	apiKey := strings.TrimSpace(oi.ApiKey.Value)
	if apiKey == "" {
		return errors.New(i18n.T("azure_api_key_required"))
	}

	baseURL := strings.TrimSpace(oi.ApiBaseURL.Value)
	if baseURL == "" {
		return errors.New(i18n.T("azure_base_url_required"))
	}

	apiVersion := strings.TrimSpace(oi.ApiVersion.Value)
	if apiVersion == "" {
		apiVersion = azurecommon.DefaultAPIVersion
		oi.ApiVersion.Value = apiVersion
	}

	endpoint := azurecommon.BuildEndpoint(baseURL)

	client := openaiapi.NewClient(
		azure.WithAPIKey(apiKey),
		option.WithBaseURL(endpoint),
		option.WithQueryAdd("api-version", apiVersion),
		option.WithMiddleware(azurecommon.AzureDeploymentMiddleware),
	)
	oi.ApiClient = &client
	return nil
}

func (oi *Client) ListModels() (ret []string, err error) {
	ret = oi.apiDeployments
	return
}
