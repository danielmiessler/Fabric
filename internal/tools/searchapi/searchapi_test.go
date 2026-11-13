package searchapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, handler http.HandlerFunc) {
	server := httptest.NewServer(handler)
	oldURL := searchURL
	searchURL = server.URL
	t.Cleanup(func() {
		searchURL = oldURL
		server.Close()
	})
}

func TestSearchSendsKeyAndFormatsResults(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "Bearer secret", r.Header.Get("Authorization"))
		require.Equal(t, "google", r.URL.Query().Get("engine"))
		require.Equal(t, "golang generics", r.URL.Query().Get("q"))
		require.Empty(t, r.URL.Query().Get("api_key"), "the key must travel in the header, not the URL")
		require.Empty(t, r.URL.Query().Get("num"), "Google no longer honours num, so it must not be sent")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"search_metadata": {"status": "Success"},
			"organic_results": [
				{"position": 1, "title": "Go generics", "link": "https://go.dev/doc/tutorial/generics", "snippet": "Getting started with generics."},
				{"position": 2, "title": "Type parameters", "link": "https://go.dev/blog/intro-generics", "snippet": "An introduction to generics."}
			]
		}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.Equal(t, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nGetting started with generics.\n\n"+
		"## Type parameters\nhttps://go.dev/blog/intro-generics\n\nAn introduction to generics.", got)
}

func TestSearchReturnsAPIError(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"Invalid API key."}`))
	})

	client := NewClient()
	client.ApiKey.Value = "bad"

	_, err := client.Search("anything")
	require.ErrorContains(t, err, "401")
	require.ErrorContains(t, err, "Invalid API key")
}

func TestSearchReturnsEmptyForNoResults(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"search_metadata": {"status": "Success"}}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("nothing here")
	require.NoError(t, err)
	require.Empty(t, got)
}

func TestIsConfiguredRequiresApiKey(t *testing.T) {
	client := NewClient()
	require.False(t, client.IsConfigured())
	require.Equal(t, "SEARCHAPI_API_KEY", client.ApiKey.EnvVariable)

	client.ApiKey.Value = "secret"
	require.True(t, client.IsConfigured())
}
