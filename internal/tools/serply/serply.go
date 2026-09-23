package serply

// see https://serply.io/docs for more information

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

const userAgent = "Fabric"

var searchURL = "https://api.serply.io/v1/search"

type Client struct {
	*plugins.PluginBase
	ApiKey     *plugins.SetupQuestion
	HttpClient *http.Client
}

func NewClient() (ret *Client) {

	label := "Serply"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("serply_label"),
			SetupDescription: i18n.T("serply_setup_description") + " " + i18n.T("optional_marker"),
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		},
	}

	ret.ApiKey = ret.AddSetupQuestion("API Key", false)

	return
}

// IsConfigured reports whether an API key is set; the key is optional during setup.
func (c *Client) IsConfigured() bool {
	return c.ApiKey.Value != ""
}

type searchResponse struct {
	Results []struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
	} `json:"results"`
}

// Search returns the Google results for query as markdown, one section per result.
func (c *Client) Search(query string) (ret string, err error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("num", "10")

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, searchURL+"?"+params.Encode(), nil); err != nil {
		err = fmt.Errorf(i18n.T("serply_failed_create_request"), err)
		return
	}
	req.Header.Set("X-Api-Key", c.ApiKey.Value)
	req.Header.Set("User-Agent", userAgent)

	client := c.HttpClient
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("serply_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("serply_api_request_failed"), resp.StatusCode, string(body))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("serply_failed_decode_response"), err)
		return
	}

	var sb strings.Builder
	for _, r := range result.Results {
		fmt.Fprintf(&sb, "## %s\n%s\n\n%s\n\n", r.Title, r.Link, r.Description)
	}
	ret = strings.TrimSpace(sb.String())
	return
}
