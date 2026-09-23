package serply

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchSendsKeyAndFormatsResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "secret", r.Header.Get("X-Api-Key"))
		require.Equal(t, "Fabric", r.Header.Get("User-Agent"))
		require.Equal(t, "golang generics", r.URL.Query().Get("q"))
		require.Equal(t, "10", r.URL.Query().Get("num"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"results":[{"title":"Go generics","link":"https://go.dev/doc/tutorial/generics","description":"Getting started with generics."}]}`))
	}))
	defer server.Close()

	oldURL := searchURL
	searchURL = server.URL
	defer func() { searchURL = oldURL }()

	client := NewClient()
	client.ApiKey.Value = "secret"
	client.HttpClient = server.Client()

	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.Equal(t, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nGetting started with generics.", got)
}

func TestSearchReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"detail":"Invalid API key"}`))
	}))
	defer server.Close()

	oldURL := searchURL
	searchURL = server.URL
	defer func() { searchURL = oldURL }()

	client := NewClient()
	client.ApiKey.Value = "bad"
	client.HttpClient = server.Client()

	_, err := client.Search("anything")
	require.ErrorContains(t, err, "Invalid API key")
}

func TestIsConfiguredRequiresApiKey(t *testing.T) {
	client := NewClient()
	require.False(t, client.IsConfigured())

	client.ApiKey.Value = "secret"
	require.True(t, client.IsConfigured())
}
