package firecrawl

// see https://docs.firecrawl.dev/features/search?utm_source=fabric&utm_medium=integration for more information

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

var searchURL = "https://api.firecrawl.dev/v2/search"

// Search also scrapes each result page. Firecrawl times out a search after 60s by default,
// so the client waits a bit longer to get Firecrawl's own error instead of a client timeout.
var httpClient = &http.Client{Timeout: 90 * time.Second}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "Firecrawl"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("firecrawl_label"),
			SetupDescription: i18n.T("firecrawl_setup_description") + " " + i18n.T("optional_marker"),
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
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    struct {
		Web []struct {
			Title       string `json:"title"`
			URL         string `json:"url"`
			Description string `json:"description"`
			Markdown    string `json:"markdown"`
		} `json:"web"`
	} `json:"data"`
}

// Search returns the top web results for query as markdown, one section per result,
// with the page content of each result when Firecrawl could scrape it.
func (c *Client) Search(query string) (ret string, err error) {
	body := map[string]any{
		"query":   query,
		"limit":   5,
		"sources": []string{"web"},
		"origin":  "fabric",
		"scrapeOptions": map[string]any{
			"formats":         []string{"markdown"},
			"onlyMainContent": true,
		},
	}

	payload, _ := json.Marshal(body)

	var req *http.Request
	if req, err = http.NewRequest(http.MethodPost, searchURL, bytes.NewReader(payload)); err != nil {
		err = fmt.Errorf(i18n.T("firecrawl_failed_create_request"), err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.ApiKey.Value)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("firecrawl_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("firecrawl_api_request_failed"), resp.StatusCode, string(respBody))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("firecrawl_failed_decode_response"), err)
		return
	}
	if !result.Success {
		err = fmt.Errorf(i18n.T("firecrawl_api_request_failed"), resp.StatusCode, result.Error)
		return
	}

	// Pages bring their own headings, so a rule between results keeps them apart.
	sections := make([]string, 0, len(result.Data.Web))
	for _, r := range result.Data.Web {
		content := r.Markdown
		if content == "" {
			content = r.Description
		}
		sections = append(sections, fmt.Sprintf("## %s\n%s\n\n%s", r.Title, r.URL, strings.TrimSpace(content)))
	}
	ret = strings.Join(sections, "\n\n---\n\n")
	return
}
