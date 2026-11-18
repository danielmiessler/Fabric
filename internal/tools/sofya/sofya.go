package sofya

// see https://sofya.co/docs for more information

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

var searchURL = "https://sofya.co/v1/search"

// A basic search also reads each result page, so it takes longer than a plain results lookup.
var httpClient = &http.Client{Timeout: 60 * time.Second}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "Sofya"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("sofya_label"),
			SetupDescription: i18n.T("sofya_setup_description") + " " + i18n.T("optional_marker"),
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
		URL         string `json:"url"`
		Content     string `json:"content"`
		Description string `json:"description"`
	} `json:"results"`
}

// Search returns the top web results for query as markdown, one section per result,
// with the page text of each result when Sofya could read the page.
func (c *Client) Search(query string) (ret string, err error) {
	body := map[string]any{
		"query":        query,
		"search_depth": "basic",
		"max_results":  5,
	}

	payload, _ := json.Marshal(body)

	var req *http.Request
	if req, err = http.NewRequest(http.MethodPost, searchURL, bytes.NewReader(payload)); err != nil {
		err = fmt.Errorf(i18n.T("sofya_failed_create_request"), err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.ApiKey.Value)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("sofya_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("sofya_api_request_failed"), resp.StatusCode, string(respBody))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("sofya_failed_decode_response"), err)
		return
	}

	// Pages bring their own headings, so a rule between results keeps them apart.
	sections := make([]string, 0, len(result.Results))
	for _, r := range result.Results {
		content := r.Content
		if content == "" {
			content = r.Description
		}
		sections = append(sections, fmt.Sprintf("## %s\n%s\n\n%s", r.Title, r.URL, strings.TrimSpace(content)))
	}
	ret = strings.Join(sections, "\n\n---\n\n")
	return
}
