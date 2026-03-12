package cli

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/stretchr/testify/require"
)

func TestHandleSetupAndServerCommandsNoop(t *testing.T) {
	handled, err := handleSetupAndServerCommands(&Flags{}, nil, "test-version")
	require.False(t, handled)
	require.NoError(t, err)
}

func TestHandleToolProcessingNoop(t *testing.T) {
	toolsMessage, err := handleToolProcessing(&Flags{}, &core.PluginRegistry{})
	require.NoError(t, err)
	require.Empty(t, toolsMessage)
}

func TestHandleToolProcessingRejectsInvalidYouTubeURL(t *testing.T) {
	registry := newToolProcessingRegistry(t)

	toolsMessage, err := handleToolProcessing(&Flags{YouTube: "not-a-youtube-url"}, registry)
	require.Error(t, err)
	require.Empty(t, toolsMessage)
	require.Contains(t, err.Error(), "not-a-youtube-url")
}

func TestHandleToolProcessingRejectsInvalidSpotifyURL(t *testing.T) {
	registry := newToolProcessingRegistry(t)

	toolsMessage, err := handleToolProcessing(&Flags{Spotify: "not-a-spotify-url"}, registry)
	require.Error(t, err)
	require.Empty(t, toolsMessage)
	require.Contains(t, err.Error(), "not-a-spotify-url")
}

func TestHandleToolProcessingScrapeURLWritesOutput(t *testing.T) {
	registry := newToolProcessingRegistry(t)
	installHTTPTransport(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, "https://r.jina.ai/https://example.com/page", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("scraped page body")),
			Request:    req,
		}, nil
	})

	outputPath := filepath.Join(t.TempDir(), "scraped.md")
	stdout, err := captureStdout(func() error {
		toolsMessage, runErr := handleToolProcessing(&Flags{
			ScrapeURL: "https://example.com/page",
			Output:    outputPath,
		}, registry)
		require.Equal(t, "scraped page body", toolsMessage)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "scraped page body")

	content, readErr := os.ReadFile(outputPath)
	require.NoError(t, readErr)
	require.Equal(t, "scraped page body\n", string(content))
}

func TestHandleToolProcessingScrapeQuestionAppendsForChatRequests(t *testing.T) {
	registry := newToolProcessingRegistry(t)
	installHTTPTransport(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, "https://s.jina.ai/what-is-fabric", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("question answer")),
			Request:    req,
		}, nil
	})

	toolsMessage, err := handleToolProcessing(&Flags{
		ScrapeQuestion: "what-is-fabric",
		Message:        "chat-mode",
	}, registry)
	require.NoError(t, err)
	require.Equal(t, "question answer", toolsMessage)
}

func TestHandleToolProcessingScrapeSourcesAppendInOrder(t *testing.T) {
	registry := newToolProcessingRegistry(t)
	requests := make([]string, 0, 2)
	installHTTPTransport(t, func(req *http.Request) (*http.Response, error) {
		requests = append(requests, req.URL.String())
		body := "scraped url"
		if strings.Contains(req.URL.String(), "s.jina.ai") {
			body = "scraped answer"
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(body)),
			Request:    req,
		}, nil
	})

	toolsMessage, err := handleToolProcessing(&Flags{
		ScrapeURL:      "https://example.com/page",
		ScrapeQuestion: "what-is-fabric",
		Message:        "chat-mode",
	}, registry)
	require.NoError(t, err)
	require.Equal(t, "scraped url\nscraped answer", toolsMessage)
	require.Equal(t, []string{
		"https://r.jina.ai/https://example.com/page",
		"https://s.jina.ai/what-is-fabric",
	}, requests)
}

func newToolProcessingRegistry(t *testing.T) *core.PluginRegistry {
	t.Helper()

	db := newConfiguredTestDB(t)
	registry, err := core.NewPluginRegistry(db)
	require.NoError(t, err)
	return registry
}
