package azure_entra

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
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
	ret.Client = openai.NewClientCompatibleNoSetupQuestions("AzureEntra", ret.configure)

	ret.ApiBaseURL = ret.AddSetupQuestionCustom("API Base URL", true,
		i18n.T("azure_base_url_question"))
	ret.ApiDeployments = ret.AddSetupQuestionCustom("deployments", true,
		i18n.T("azure_deployments_question"))
	ret.ApiVersion = ret.AddSetupQuestionCustom("API Version", false,
		i18n.T("azure_api_version_question"))

	return
}

type Client struct {
	*openai.Client // ApiBaseURL is inherited from the embedded openai.Client
	ApiDeployments *plugins.SetupQuestion
	ApiVersion     *plugins.SetupQuestion

	apiDeployments []string
}

func (c *Client) configure() error {
	c.apiDeployments = azurecommon.ParseDeployments(c.ApiDeployments.Value)
	if len(c.apiDeployments) == 0 {
		return fmt.Errorf("%s", i18n.T("azure_deployments_required"))
	}

	baseURL := strings.TrimSpace(c.ApiBaseURL.Value)
	if baseURL == "" {
		return fmt.Errorf("%s", i18n.T("azure_base_url_required"))
	}

	apiVersion := strings.TrimSpace(c.ApiVersion.Value)
	if apiVersion == "" {
		apiVersion = azurecommon.DefaultAPIVersion
		c.ApiVersion.Value = apiVersion
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("%s: %w", i18n.T("azure_credential_failure"), err)
	}

	endpoint := azurecommon.BuildEndpoint(baseURL)

	client := openaiapi.NewClient(
		azure.WithTokenCredential(credential),
		option.WithBaseURL(endpoint),
		option.WithQueryAdd("api-version", apiVersion),
		option.WithMiddleware(azurecommon.AzureDeploymentMiddleware),
	)
	c.ApiClient = &client
	return nil
}

func (c *Client) ListModels() (ret []string, err error) {
	ret = c.apiDeployments
	return
}
