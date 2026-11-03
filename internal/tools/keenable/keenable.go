package keenable

// see https://keenable.ai for more information

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

// The public endpoint works without an API key; the keyed endpoint is used
// when an API key is configured.
var (
	publicSearchURL = "https://api.keenable.ai/v1/search/public"
	keyedSearchURL  = "https://api.keenable.ai/v1/search"
)

const snippetMaxChars = 1000

var httpClient = &http.Client{Timeout: 30 * time.Second}

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
}

func NewClient() (ret *Client) {

	label := "Keenable"

	ret = &Client{
		PluginBase: &plugins.PluginBase{
			Name:             i18n.T("keenable_label"),
			SetupDescription: i18n.T("keenable_setup_description") + " " + i18n.T("optional_marker"),
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		},
	}

	ret.ApiKey = ret.AddSetupQuestion("API Key", false)

	return
}

type searchResponse struct {
	Results []struct {
		Title       string `json:"title"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Snippet     string `json:"snippet"`
	} `json:"results"`
}

// Search returns the web results for query as markdown, one section per result.
// It works without an API key via Keenable's public endpoint.
func (c *Client) Search(query string) (ret string, err error) {
	var payload []byte
	if payload, err = json.Marshal(map[string]any{"query": query, "max_results": 10}); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_create_request"), err)
		return
	}

	endpoint := publicSearchURL
	apiKey := strings.TrimSpace(c.ApiKey.Value)
	if apiKey != "" {
		endpoint = keyedSearchURL
	}

	var req *http.Request
	if req, err = http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(payload)); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_create_request"), err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Keenable-Title", "Fabric")
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_execute_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf(i18n.T("keenable_api_request_failed"), resp.StatusCode, string(body))
		return
	}

	var result searchResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_decode_response"), err)
		return
	}

	var sb strings.Builder
	for _, r := range result.Results {
		if r.URL == "" {
			continue
		}
		text := r.Description
		if text == "" {
			text = r.Snippet
			if runes := []rune(text); len(runes) > snippetMaxChars {
				text = string(runes[:snippetMaxChars]) + "…"
			}
		}
		fmt.Fprintf(&sb, "## %s\n%s\n\n%s\n\n", r.Title, r.URL, text)
	}
	ret = strings.TrimSpace(sb.String())
	return
}
