// Package parallel provides opt-in web search through Parallel's Unified MCP server.
package parallel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const endpoint = "https://mcp.parallel.ai/mcp"
const userAgent = "Fabric/ParallelSearch (https://github.com/danielmiessler/Fabric)"

type Client struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
	// endpoint is private so normal use always connects to the Unified MCP server.
	endpoint string
}

func NewClient() *Client {
	c := &Client{
		PluginBase: &plugins.PluginBase{
			Name:             "Parallel",
			SetupDescription: "Parallel web search (optional; anonymous access needs no API key)",
			EnvNamePrefix:    plugins.BuildEnvVariablePrefix("Parallel"),
		},
		endpoint: endpoint,
	}
	c.ApiKey = c.AddSetupQuestion("API Key", false)
	return c
}

// requestTransport applies identification and optional credentials to every MCP
// request, including initialization, notifications, and session termination.
type requestTransport struct {
	base   http.RoundTripper
	apiKey string
}

func (t requestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("User-Agent", userAgent)
	if t.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+t.apiKey)
	}
	return t.base.RoundTrip(req)
}

// Search returns source URLs and excerpts as Markdown. Anonymous searches use
// the server's default Fast mode. Explicit credentials are never retried anonymously.
func (c *Client) Search(query string) (string, error) {
	if strings.TrimSpace(query) == "" {
		return "", fmt.Errorf("Parallel search requires a non-empty query")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	client := mcp.NewClient(&mcp.Implementation{Name: "fabric", Version: "1"},
		&mcp.ClientOptions{Capabilities: &mcp.ClientCapabilities{}})
	// Use the established session protocol supported by Unified MCP.
	session, err := client.Connect(ctx, &mcp.StreamableClientTransport{
		Endpoint: c.endpoint,
		HTTPClient: &http.Client{
			Timeout:   90 * time.Second,
			Transport: requestTransport{base: http.DefaultTransport, apiKey: c.ApiKey.Value},
			// Never forward credentials or queries to a redirected endpoint.
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
		DisableStandaloneSSE: true,
		MaxRetries:           -1,
	}, &mcp.ClientSessionOptions{ProtocolVersion: "2025-11-25"})
	if err != nil {
		return "", fmt.Errorf("Parallel MCP connection: %w", err)
	}
	defer session.Close()

	found := false
	for tool, err := range session.Tools(ctx, nil) {
		if err != nil {
			return "", fmt.Errorf("Parallel MCP discovery: %w", err)
		}
		if tool.Name == "web_search" {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("Parallel MCP does not offer web_search")
	}
	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "web_search",
		Arguments: map[string]any{"objective": query, "search_queries": []string{query}},
	})
	if err != nil {
		return "", fmt.Errorf("Parallel search: %w", err)
	}
	texts := []string{}
	for _, content := range result.Content {
		if text, ok := content.(*mcp.TextContent); ok {
			texts = append(texts, text.Text)
		}
	}
	if result.IsError {
		return "", fmt.Errorf("Parallel search: %s", strings.Join(texts, "\n"))
	}
	var payload []byte
	if result.StructuredContent != nil {
		payload, err = json.Marshal(result.StructuredContent)
	} else {
		payload = []byte(strings.Join(texts, "\n"))
	}
	if err != nil {
		return "", fmt.Errorf("Parallel search response: %w", err)
	}
	return formatResults(payload)
}

func formatResults(payload []byte) (string, error) {
	var response struct {
		Results []struct {
			Title    string   `json:"title"`
			URL      string   `json:"url"`
			Excerpts []string `json:"excerpts"`
		} `json:"results"`
		Warnings []struct {
			Message string `json:"message"`
		} `json:"warnings"`
	}
	if err := json.Unmarshal(payload, &response); err != nil {
		return "", fmt.Errorf("Parallel search response: %w", err)
	}
	sections := []string{}
	for _, r := range response.Results {
		title := r.Title
		if title == "" {
			title = r.URL
		}
		sections = append(sections, fmt.Sprintf("## %s\n%s\n\n%s", title, r.URL, strings.Join(r.Excerpts, "\n\n")))
	}
	if len(sections) == 0 {
		sections = append(sections, "No search results.")
	}
	for _, warning := range response.Warnings {
		sections = append(sections, "Parallel warning: "+warning.Message)
	}
	return strings.Join(sections, "\n\n---\n\n"), nil
}
