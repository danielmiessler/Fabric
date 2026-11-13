package sofya

import (
	"encoding/json"
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

func TestSearchSendsRequestAndFormatsResults(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/", r.URL.Path)
		require.Equal(t, "Bearer secret", r.Header.Get("Authorization"))
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var body struct {
			Query       string `json:"query"`
			SearchDepth string `json:"search_depth"`
			MaxResults  int    `json:"max_results"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		require.Equal(t, "golang generics", body.Query)
		require.Equal(t, "basic", body.SearchDepth)
		require.Equal(t, 5, body.MaxResults)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"query":"golang generics","search_depth":"basic","results":[` +
			`{"title":"Go generics","url":"https://go.dev/doc/tutorial/generics","description":"Getting started with generics.","content":"# Tutorial: Getting started with generics\n\nThis tutorial introduces generics.\n","fetched":true},` +
			`{"title":"Generics - Go by Example","url":"https://gobyexample.com/generics","description":"Go by Example: Generics","content":"Starting with version 1.18, Go has added support for generics.","fetched":true}` +
			`],"credits_used":2,"credits_remaining":98}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.Equal(t, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\n# Tutorial: Getting started with generics\n\nThis tutorial introduces generics."+
		"\n\n---\n\n"+
		"## Generics - Go by Example\nhttps://gobyexample.com/generics\n\nStarting with version 1.18, Go has added support for generics.", got)
}

func TestSearchFallsBackToDescription(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"results":[{"title":"Go generics","url":"https://go.dev/doc/tutorial/generics","description":"Getting started with generics.","content":""}]}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.Equal(t, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nGetting started with generics.", got)
}

func TestSearchReturnsAPIError(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"detail":"Invalid API key"}`))
	})

	client := NewClient()
	client.ApiKey.Value = "bad"

	_, err := client.Search("anything")
	require.ErrorContains(t, err, "401")
	require.ErrorContains(t, err, "Invalid API key")
}

func TestSearchReturnsEmptyWhenNoResults(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"results":[]}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("nothing matches this")
	require.NoError(t, err)
	require.Empty(t, got)
}

func TestIsConfiguredRequiresApiKey(t *testing.T) {
	client := NewClient()
	require.False(t, client.IsConfigured())

	client.ApiKey.Value = "secret"
	require.True(t, client.IsConfigured())
}
