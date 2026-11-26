package specstream

// see https://specstream.ai/agents.md for more information

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

var searchURL = "https://api.specstream.ai/v1/web-search"

// With passages=true the search also reads the top result pages, so allow more than a plain lookup.
var httpClient = &http.Client{Timeout: 60 * time.Second}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "SpecStream"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("specstream_label"),
			SetupDescription: i18n.T("specstream_setup_description") + " " + i18n.T("optional_marker"),
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
	Data struct {
		Sources []struct {
			Title      string   `json:"title"`
			URL        string   `json:"url"`
			Snippet    string   `json:"snippet"`
			Highlights []string `json:"highlights"`
		} `json:"sources"`
	} `json:"data"`
}

// Search returns the top web results for query as markdown, one section per result,
// with the passages of each page closest to the query when SpecStream read the page.
func (c *Client) Search(query string) (ret string, err error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("num", "10")
	params.Set("speed", "fast")
	params.Set("passages", "true")

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, searchURL+"?"+params.Encode(), nil); err != nil {
		err = fmt.Errorf(i18n.T("specstream_failed_create_request"), err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.ApiKey.Value)

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("specstream_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("specstream_api_request_failed"), resp.StatusCode, string(body))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("specstream_failed_decode_response"), err)
		return
	}

	// Passages can carry their own headings, so a rule between results keeps them apart.
	sections := make([]string, 0, len(result.Data.Sources))
	for _, s := range result.Data.Sources {
		content := strings.TrimSpace(strings.Join(s.Highlights, "\n\n"))
		if content == "" {
			content = strings.TrimSpace(s.Snippet)
		}
		sections = append(sections, fmt.Sprintf("## %s\n%s\n\n%s", s.Title, s.URL, content))
	}
	ret = strings.Join(sections, "\n\n---\n\n")
	return
}
