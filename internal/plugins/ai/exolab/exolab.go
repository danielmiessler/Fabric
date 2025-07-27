package exolab

import (
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.Client = openai.NewClientCompatibleNoSetupQuestions("Exolab", ret.configure)

	ret.ApiKey = ret.AddSetupQuestion("API Key", false)
	ret.ApiBaseURL = ret.AddSetupQuestion("API Base URL", true)
	ret.ApiBaseURL.Value = "http://localhost:52415"

	ret.ApiModels = ret.AddSetupQuestionCustom("models", true,
		"Enter your deployed Exolab models (comma separated)")

	return
}

type Client struct {
	*openai.Client
	ApiModels *plugins.SetupQuestion

	apiModels []string
}

func (oi *Client) configure() (err error) {
	oi.apiModels = strings.Split(oi.ApiModels.Value, ",")

	opts := []option.RequestOption{option.WithAPIKey(oi.ApiKey.Value)}
	if oi.ApiBaseURL.Value != "" {
		opts = append(opts, option.WithBaseURL(oi.ApiBaseURL.Value))
	}
	client := openaiapi.NewClient(opts...)
	oi.ApiClient = &client
	return
}

func (oi *Client) ListModels() (ret []string, err error) {
	ret = oi.apiModels
	return
}

func (oi *Client) NeedsRawMode(modelName string) bool {
	return false
}
