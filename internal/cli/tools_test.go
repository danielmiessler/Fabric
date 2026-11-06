package cli

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/tools/firecrawl"
	"github.com/danielmiessler/fabric/internal/tools/parallel"
	"github.com/danielmiessler/fabric/internal/tools/serply"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

type toolTestTransport func(*http.Request) (*http.Response, error)

func (f toolTestTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestParallelToolProcessing(t *testing.T) {
	server := mcp.NewServer(&mcp.Implementation{Name: "fixture", Version: "1"}, nil)
	calls := 0
	server.AddTool(&mcp.Tool{Name: "web_search", InputSchema: json.RawMessage(`{"type":"object"}`)}, func(context.Context, *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		calls++
		return &mcp.CallToolResult{StructuredContent: map[string]any{"results": []any{map[string]any{"title": "Go", "url": "https://go.dev", "excerpts": []string{"A useful passage"}}}}}, nil
	})
	endpoint := httptest.NewServer(mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, nil))
	defer endpoint.Close()
	original := http.DefaultTransport
	http.DefaultTransport = toolTestTransport(func(r *http.Request) (*http.Response, error) {
		require.Equal(t, "mcp.parallel.ai", r.URL.Host)
		local := r.Clone(r.Context())
		local.URL.Scheme = "http"
		local.URL.Host = endpoint.Listener.Addr().String()
		local.Host = local.URL.Host
		return original.RoundTrip(local)
	})
	t.Cleanup(func() { http.DefaultTransport = original })
	registry := &core.PluginRegistry{Parallel: parallel.NewClient(), Serply: serply.NewClient(), Firecrawl: firecrawl.NewClient()}
	for _, chat := range []bool{false, true} {
		flags := &Flags{ParallelSearch: "Go tutorial"}
		if chat {
			flags.Pattern = "summarize"
		}
		var message string
		printed := captureStdout(t, func() { var err error; message, err = handleToolProcessing(flags, registry); require.NoError(t, err) })
		require.Contains(t, message, "https://go.dev\n\nA useful passage")
		if chat {
			require.Empty(t, printed)
		} else {
			require.Equal(t, message+"\n", string(printed))
		}
	}
	require.Equal(t, 2, calls)
	// Existing explicit provider errors take precedence, with no Parallel request.
	for _, flags := range []*Flags{{SerplySearch: "query", ParallelSearch: "query"}, {FirecrawlSearch: "query", ParallelSearch: "query"}} {
		_, err := handleToolProcessing(flags, registry)
		require.ErrorContains(t, err, "not configured")
	}
	require.Equal(t, 2, calls)
	// No opt-in, including a model search selection or model dry run, makes no
	// Parallel requests. Model-native search is handled elsewhere as before.
	for _, flags := range []*Flags{{}, {Search: true}, {DryRun: true}} {
		_, err := handleToolProcessing(flags, registry)
		require.NoError(t, err)
	}
	require.Equal(t, 2, calls)
}
