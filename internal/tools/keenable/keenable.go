package keenable

// see https://keenable.ai for more information

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

// The public endpoint works without an API key; the keyed endpoint is used
// when an API key is configured.
var (
	publicSearchURL = "https://api.keenable.ai/v1/search/public"
	keyedSearchURL  = "https://api.keenable.ai/v1/search"
)

const (
	maxResults = 10
	// maxResponseBytes bounds a successful response before it is decoded.
	maxResponseBytes = 1 << 20
	// maxErrorBodyBytes bounds how much of an error response is read for diagnostics.
	maxErrorBodyBytes = 2 << 10
	titleMaxChars     = 300
	urlMaxChars       = 2048
	snippetMaxChars   = 1000
)

var (
	errRedirectRejected = errors.New("redirects are not followed")
	errResponseTooLarge = fmt.Errorf("response body exceeds %d bytes", maxResponseBytes)
)

// newHTTPClient never follows redirects: the API does not need them, and
// following one could send X-API-Key or the query to another host or over HTTP.
func newHTTPClient(transport http.RoundTripper) *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return errRedirectRejected
		},
	}
}

var httpClient = newHTTPClient(nil)

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
	if payload, err = json.Marshal(map[string]any{"query": query, "max_results": maxResults}); err != nil {
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
		err = fmt.Errorf(i18n.T("keenable_api_request_failed"), resp.StatusCode, readDiagnostic(resp.Body))
		return
	}

	var body []byte
	if body, err = readLimited(resp.Body, maxResponseBytes); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_decode_response"), err)
		return
	}

	var result searchResponse
	if err = json.Unmarshal(body, &result); err != nil {
		err = fmt.Errorf(i18n.T("keenable_failed_decode_response"), err)
		return
	}

	var sb strings.Builder
	count := 0
	for _, r := range result.Results {
		if count == maxResults {
			break
		}
		if r.URL == "" {
			continue
		}
		text := r.Description
		if text == "" {
			text = r.Snippet
		}
		fmt.Fprintf(&sb, "## %s\n%s\n\n%s\n\n",
			truncate(r.Title, titleMaxChars), truncate(r.URL, urlMaxChars), truncate(text, snippetMaxChars))
		count++
	}
	ret = strings.TrimSpace(sb.String())
	return
}

// readLimited reads at most limit bytes and reports errResponseTooLarge when the
// body is longer, even if the first limit bytes are valid JSON.
func readLimited(r io.Reader, limit int64) ([]byte, error) {
	body, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	if int64(len(body)) > limit {
		return nil, errResponseTooLarge
	}
	return body, nil
}

// readDiagnostic returns a bounded, printable excerpt of an error response.
func readDiagnostic(r io.Reader) string {
	body, err := io.ReadAll(io.LimitReader(r, maxErrorBodyBytes+1))
	truncated := len(body) > maxErrorBodyBytes
	if truncated {
		body = body[:maxErrorBodyBytes]
	}
	text := sanitize(string(body))
	if truncated {
		text += " …(truncated)"
	}
	if err != nil {
		text += fmt.Sprintf(" (reading response body: %v)", err)
	}
	return text
}

func sanitize(s string) string {
	return strings.Map(func(r rune) rune {
		if r == unicode.ReplacementChar || (unicode.IsControl(r) && r != ' ') {
			return ' '
		}
		return r
	}, strings.ToValidUTF8(s, " "))
}

func truncate(s string, max int) string {
	if runes := []rune(s); len(runes) > max {
		return string(runes[:max]) + "…"
	}
	return s
}
