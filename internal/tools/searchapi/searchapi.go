package searchapi

// see https://www.searchapi.io/docs/google for more information

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

var searchURL = "https://www.searchapi.io/api/v1/search"

var httpClient = &http.Client{Timeout: 30 * time.Second}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "SearchApi"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("searchapi_label"),
			SetupDescription: i18n.T("searchapi_setup_description") + " " + i18n.T("optional_marker"),
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
	OrganicResults []struct {
		Title   string `json:"title"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"organic_results"`
}

// Search returns the Google results for query as markdown, one section per result.
// Google no longer honours a result count, so one call returns one page of results.
func (c *Client) Search(query string) (ret string, err error) {
	params := url.Values{}
	params.Set("engine", "google")
	params.Set("q", query)

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, searchURL+"?"+params.Encode(), nil); err != nil {
		err = fmt.Errorf(i18n.T("searchapi_failed_create_request"), err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.ApiKey.Value)

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("searchapi_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("searchapi_api_request_failed"), resp.StatusCode, string(body))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("searchapi_failed_decode_response"), err)
		return
	}

	var sb strings.Builder
	for _, r := range result.OrganicResults {
		fmt.Fprintf(&sb, "## %s\n%s\n\n%s\n\n", r.Title, r.Link, r.Snippet)
	}
	ret = strings.TrimSpace(sb.String())
	return
}
