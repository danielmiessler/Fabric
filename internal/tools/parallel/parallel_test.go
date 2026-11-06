package parallel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

// Exercise the real SDK transport against a local MCP server, with both response
// encodings supported by Streamable HTTP. These are synthetic, not live results.
func TestSearchMCP(t *testing.T) {
	for _, jsonResponse := range []bool{false, true} {
		for _, key := range []string{"", "explicit-key"} {
			t.Run(fmt.Sprintf("json=%v/key=%v", jsonResponse, key != ""), func(t *testing.T) {
				server := mcp.NewServer(&mcp.Implementation{Name: "test", Version: "1"}, nil)
				server.AddTool(&mcp.Tool{Name: "web_search", InputSchema: json.RawMessage(`{"type":"object"}`)},
					func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
						require.JSONEq(t, `{"objective":"Go generics tutorial","search_queries":["Go generics tutorial"]}`, string(req.Params.Arguments))
						return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"results":[{"title":"Go","url":"https://go.dev/","excerpts":["A passage","Another passage"]}],"warnings":[{"message":"Limited results"}]}`}}}, nil
					})
				handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, &mcp.StreamableHTTPOptions{JSONResponse: jsonResponse})
				var mu sync.Mutex
				methods := map[string]int{}
				endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					require.Equal(t, userAgent, r.Header.Get("User-Agent"))
					if key == "" {
						require.Empty(t, r.Header.Get("Authorization"))
					} else {
						require.Equal(t, "Bearer "+key, r.Header.Get("Authorization"))
					}
					method := r.Method
					if r.Method == "POST" {
						payload, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						r.Body = io.NopCloser(bytes.NewReader(payload))
						var message struct {
							Method string `json:"method"`
						}
						require.NoError(t, json.Unmarshal(payload, &message))
						method = message.Method
					}
					mu.Lock()
					methods[method]++
					mu.Unlock()
					handler.ServeHTTP(w, r)
				}))
				defer endpoint.Close()
				client := NewClient()
				client.endpoint = endpoint.URL
				client.ApiKey.Value = key
				got, err := client.Search("Go generics tutorial")
				require.NoError(t, err)
				require.Equal(t, "## Go\nhttps://go.dev/\n\nA passage\n\nAnother passage\n\n---\n\nParallel warning: Limited results", got)
				mu.Lock()
				defer mu.Unlock()
				for _, method := range []string{"initialize", "notifications/initialized", "tools/list", "tools/call", "DELETE"} {
					require.Equal(t, 1, methods[method], "method %s; requests %v", method, methods)
				}
				require.Zero(t, methods["GET"])
			})
		}
	}
}

func TestSearchErrors(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests, http.StatusInternalServerError} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				require.Equal(t, "Bearer bad", r.Header.Get("Authorization"))
				w.WriteHeader(status)
			}))
			defer server.Close()
			c := NewClient()
			c.endpoint = server.URL
			c.ApiKey.Value = "bad"
			_, err := c.Search("query")
			require.Error(t, err)
			require.Equal(t, 1, calls, "must not retry anonymously or switch provider")
		})
	}
	for _, test := range []struct {
		name   string
		tool   bool
		result *mcp.CallToolResult
		want   string
	}{
		{name: "missing tool", want: "does not offer web_search"},
		{name: "tool error", tool: true, result: &mcp.CallToolResult{IsError: true, Content: []mcp.Content{&mcp.TextContent{Text: "rate limited"}}}, want: "rate limited"},
		{name: "malformed", tool: true, result: &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: "invalid JSON"}}}, want: "response"},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := mcp.NewServer(&mcp.Implementation{Name: "test", Version: "1"}, nil)
			if test.tool {
				server.AddTool(&mcp.Tool{Name: "web_search", InputSchema: json.RawMessage(`{"type":"object"}`)}, func(context.Context, *mcp.CallToolRequest) (*mcp.CallToolResult, error) { return test.result, nil })
			}
			endpoint := httptest.NewServer(mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, nil))
			defer endpoint.Close()
			c := NewClient()
			c.endpoint = endpoint.URL
			_, err := c.Search("query")
			require.ErrorContains(t, err, test.want)
		})
	}
}

func TestFormatResultsAndEmptyQuery(t *testing.T) {
	got, err := formatResults([]byte(`{"results":[]}`))
	require.NoError(t, err)
	require.Equal(t, "No search results.", got)
	got, err = formatResults([]byte(`{"results":[{"url":"https://go.dev","excerpts":["passage"]}]}`))
	require.NoError(t, err)
	require.Contains(t, got, "## https://go.dev")
	_, err = NewClient().Search("  ")
	require.ErrorContains(t, err, "non-empty query")
}

func TestOptionalCredentialConfiguration(t *testing.T) {
	t.Setenv("PARALLEL_API_KEY", "")
	c := NewClient()
	require.NoError(t, c.Configure())
	require.True(t, c.IsConfigured())
	require.Empty(t, c.ApiKey.Value)
	t.Setenv("PARALLEL_API_KEY", "explicit")
	require.NoError(t, c.Configure())
	require.Equal(t, "explicit", c.ApiKey.Value)
}
