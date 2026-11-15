package cli

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/tools/firecrawl"
	"github.com/danielmiessler/fabric/internal/tools/parallel"
	"github.com/danielmiessler/fabric/internal/tools/serply"
	"github.com/stretchr/testify/require"
)

func TestSearchToolPrecedenceAndOptIn(t *testing.T) {
	registry := &core.PluginRegistry{Parallel: parallel.NewClient(), Serply: serply.NewClient(), Firecrawl: firecrawl.NewClient()}

	// Explicit provider errors take precedence: Serply and Firecrawl configured checks run first.
	for _, flags := range []*Flags{{SerplySearch: "query", ParallelSearch: "query"}, {FirecrawlSearch: "query", ParallelSearch: "query"}} {
		_, err := handleToolProcessing(flags, registry)
		require.ErrorContains(t, err, "not configured")
	}

	// No opt-in (no ParallelSearch flag) makes no search requests.
	for _, flags := range []*Flags{{}, {Search: true}, {DryRun: true}} {
		_, err := handleToolProcessing(flags, registry)
		require.NoError(t, err)
	}
}
