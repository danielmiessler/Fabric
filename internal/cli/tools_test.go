package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/tools/firecrawl"
	"github.com/danielmiessler/fabric/internal/tools/jina"
	"github.com/danielmiessler/fabric/internal/tools/parallel"
	"github.com/danielmiessler/fabric/internal/tools/serply"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchToolPrecedenceAndOptIn(t *testing.T) {
	registry := &core.PluginRegistry{Parallel: parallel.NewClient(), Serply: serply.NewClient(), Firecrawl: firecrawl.NewClient()}

	// Explicit provider errors take precedence: Serply and Firecrawl configured checks run first.
	for _, flags := range []*Flags{{SerplySearch: "query", ParallelSearch: "query"}, {FirecrawlSearch: "query", ParallelSearch: "query"}} {
		_, _, err := handleToolProcessing(flags, registry)
		require.ErrorContains(t, err, "not configured")
	}

	// No opt-in (no ParallelSearch flag) makes no search requests.
	for _, flags := range []*Flags{{}, {Search: true}, {DryRun: true}} {
		_, _, err := handleToolProcessing(flags, registry)
		require.NoError(t, err)
	}
}

func TestWebSearchRejectsEmptyQuery(t *testing.T) {
	for _, query := range []string{"", "  "} {
		flags := &Flags{WebSearch: &query}
		require.ErrorContains(t, flags.resolveWebSearch(), "non-empty query")
	}
}

func TestWebSearchRejectsDisabledUnknownAndPreviewRequests(t *testing.T) {
	for _, tc := range []struct {
		flags Flags
		want  string
	}{
		{Flags{WebSearchProvider: new("off")}, "disabled"},
		{Flags{WebSearchProvider: new("missing")}, "Unknown standalone web search provider"},
		{Flags{DryRun: true}, "--dry-run"},
		{Flags{PrintPrompt: true}, "--print-prompt"},
	} {
		tc.flags.WebSearch = new("query")
		require.ErrorContains(t, tc.flags.resolveWebSearch(), tc.want)
	}
}

func TestWebSearchUsesParallelAndPreservesOutputRouting(t *testing.T) {
	var searchCalls atomic.Int32
	server := mcp.NewServer(&mcp.Implementation{Name: "test", Version: "1"}, nil)
	server.AddTool(&mcp.Tool{Name: "web_search", InputSchema: json.RawMessage(`{"type":"object"}`)},
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			searchCalls.Add(1)
			assert.JSONEq(t, `{"objective":"Go generics","search_queries":["Go generics"]}`, string(req.Params.Arguments))
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{
				Text: `{"results":[{"title":"Go","url":"https://go.dev/","excerpts":["A passage"]}]}`,
			}}}, nil
		})
	endpoint := httptest.NewServer(mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, nil))
	t.Cleanup(endpoint.Close)
	localURL, err := url.Parse(endpoint.URL)
	require.NoError(t, err)

	// Redirect the Parallel client to the test server without letting any real requests through.
	transport := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = transport })
	http.DefaultTransport = searchRoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != "https://mcp.parallel.ai/mcp" {
			return nil, fmt.Errorf("unexpected search endpoint: %s", req.URL)
		}
		assert.Empty(t, req.Header.Get("Authorization"))
		forwarded := req.Clone(req.Context())
		forwarded.URL.Scheme, forwarded.URL.Host = localURL.Scheme, localURL.Host
		forwarded.Host = localURL.Host
		return transport.RoundTrip(forwarded)
	})

	registry := &core.PluginRegistry{Parallel: parallel.NewClient()}
	const expected = "## Go\nhttps://go.dev/\n\nA passage"
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("PARALLEL_API_KEY", "")
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })

	output := filepath.Join(home, "results.md")
	os.Args = []string{"fabric", "--web-search", "Go generics", "-o", output}
	require.NoError(t, Cli("test"))
	written, err := os.ReadFile(output)
	require.NoError(t, err)
	require.Equal(t, expected+"\n", string(written))
	require.NoFileExists(t, filepath.Join(home, ".config", "fabric", ".env"))
	require.EqualValues(t, 1, searchCalls.Load())

	os.Args = []string{"fabric", "--web-search=Go generics", "--web-search-provider=off"}
	require.ErrorContains(t, Cli("test"), "disabled")
	os.Args = []string{"fabric", "-p", "summarize", "private input"}
	require.ErrorContains(t, Cli("test"), ".env")
	require.EqualValues(t, 1, searchCalls.Load())

	flags := &Flags{WebSearch: new("Go generics"), Pattern: "summarize", Message: "private input",
		Output: filepath.Join(home, "summary.md")}
	require.NoError(t, flags.resolveWebSearch())
	got, handled, err := handleToolProcessing(flags, registry)
	require.NoError(t, err)
	require.False(t, handled)
	require.Equal(t, expected, got)
	require.NoFileExists(t, flags.Output)
	require.EqualValues(t, 2, searchCalls.Load())
}

func TestWebSearchJinaSurfacesProviderFailure(t *testing.T) {
	const response = "Jina rejected the request"
	expectedURL := "https://s.jina.ai/C%23%20generics%3F"
	transport := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = transport })
	http.DefaultTransport = searchRoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, expectedURL, req.URL.String())
		return &http.Response{
			Status:     "429 Too Many Requests",
			StatusCode: http.StatusTooManyRequests,
			Body:       io.NopCloser(strings.NewReader(response)),
		}, nil
	})
	registry := &core.PluginRegistry{Jina: jina.NewClient()}

	generic := &Flags{WebSearch: new("C# generics?"), WebSearchProvider: new("jina")}
	require.NoError(t, generic.resolveWebSearch())
	got, _, err := handleToolProcessing(generic, registry)
	require.ErrorContains(t, err, "429")
	require.Empty(t, got)

	expectedURL = "https://s.jina.ai/query"
	legacy := &Flags{ScrapeQuestion: "query", Pattern: "summarize"}
	got, _, err = handleToolProcessing(legacy, registry)
	require.NoError(t, err)
	require.Equal(t, response, got)
}

func TestStandaloneWebSearchCompletesWithNoResults(t *testing.T) {
	transport := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = transport })
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("KEENABLE_API_KEY", "")
	t.Setenv("JINA_AI_API_KEY", "")
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })

	for _, tc := range []struct{ provider, endpoint, body string }{
		{provider: "keenable", endpoint: "https://api.keenable.ai/v1/search/public", body: `{"results":[]}`},
		{provider: "jina", endpoint: "https://s.jina.ai/query"},
	} {
		t.Run(tc.provider, func(t *testing.T) {
			var searchCalls int
			http.DefaultTransport = searchRoundTripperFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.String() != tc.endpoint {
					return nil, fmt.Errorf("unexpected endpoint: %s", req.URL)
				}
				searchCalls++
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tc.body)),
				}, nil
			})

			output := filepath.Join(t.TempDir(), "results.md")
			os.Args = []string{"fabric", "--web-search=query", "--web-search-provider=" + tc.provider, "--search", "-o", output}
			err := Cli("test")
			require.Equal(t, 1, searchCalls)
			written, readErr := os.ReadFile(output)
			require.NoError(t, readErr)
			require.Equal(t, "\n", string(written))
			require.NoError(t, err)
		})
	}
}

type searchRoundTripperFunc func(*http.Request) (*http.Response, error)

func (f searchRoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
