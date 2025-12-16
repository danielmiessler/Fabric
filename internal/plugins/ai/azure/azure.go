package azure

import (
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/azure"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatible("Azure", "", ret.configure)
	ret.ApiDeployments = ret.AddSetupQuestionCustom("deployments", true,
		"Enter your Azure deployments (comma separated)")
	ret.ApiVersion = ret.AddSetupQuestionCustom("API Version", false,
		"Enter the Azure API version (optional)")

	return ret
}

type Client struct {
	*openai.Client
	ApiDeployments *plugins.SetupQuestion
	ApiVersion     *plugins.SetupQuestion

	apiDeployments []string
}

const defaultAPIVersion = "2024-05-01-preview"

func (oi *Client) configure() error {
	oi.apiDeployments = parseDeployments(oi.ApiDeployments.Value)

	apiKey := strings.TrimSpace(oi.ApiKey.Value)
	if apiKey == "" {
		return fmt.Errorf("Azure API key is required")
	}

	baseURL := strings.TrimSpace(oi.ApiBaseURL.Value)
	if baseURL == "" {
		return fmt.Errorf("Azure API base URL is required")
	}

	apiVersion := strings.TrimSpace(oi.ApiVersion.Value)
	if apiVersion == "" {
		apiVersion = defaultAPIVersion
		oi.ApiVersion.Value = apiVersion
	}

	client := openaiapi.NewClient(
		azure.WithAPIKey(apiKey),
		azure.WithEndpoint(baseURL, apiVersion),
	)
	oi.ApiClient = &client
	return nil
}

func parseDeployments(value string) []string {
	parts := strings.Split(value, ",")
	var deployments []string
	for _, part := range parts {
		if deployment := strings.TrimSpace(part); deployment != "" {
			deployments = append(deployments, deployment)
		}
	}
	return deployments
}

func (oi *Client) ListModels() (ret []string, err error) {
	ret = oi.apiDeployments
	return ret, err
}

func (oi *Client) NeedsRawMode(modelName string) bool {
	return false
}
