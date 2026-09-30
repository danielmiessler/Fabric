package jina

// See https://jina.ai for the API documentation.

import (
	"fmt"
	"io"
	"net/http"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "Jina AI"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("jina_label"),
			SetupDescription: i18n.T("jina_setup_description") + " " + i18n.T("optional_marker"),
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		},
	}

	ret.ApiKey = ret.AddSetupQuestion("API Key", false)

	return
}

// ScrapeURL return the main content of a webpage in clean, LLM-friendly text.
func (jc *Client) ScrapeURL(url string) (ret string, err error) {
	return jc.request(fmt.Sprintf("https://r.jina.ai/%s", url))
}

func (jc *Client) ScrapeQuestion(question string) (ret string, err error) {
	return jc.request(fmt.Sprintf("https://s.jina.ai/%s", question))
}

func (jc *Client) request(requestURL string) (ret string, err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", requestURL, nil); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("jina_error_creating_request"), err))
		return
	}

	if jc.ApiKey.Value != "" {
		req.Header.Set("Authorization", "Bearer "+jc.ApiKey.Value)
	}

	client := &http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("jina_error_sending_request"), err))
		return
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("jina_error_reading_response_body"), err))
		return
	}
	ret = string(body)
	return
}
