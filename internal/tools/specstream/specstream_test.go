package specstream

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

func TestSearchSendsKeyAndFormatsPassages(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "Bearer secret", r.Header.Get("Authorization"))
		require.Equal(t, "golang generics", r.URL.Query().Get("q"))
		require.Equal(t, "10", r.URL.Query().Get("num"))
		require.Equal(t, "fast", r.URL.Query().Get("speed"))
		require.Equal(t, "true", r.URL.Query().Get("passages"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"sources":[` +
			`{"title":"Go generics","url":"https://go.dev/doc/tutorial/generics","snippet":"Getting started with generics.","highlights":["This tutorial introduces the basics of generics in Go.","You'll declare two simple non-generic functions."]},` +
			`{"title":"Generics - Go by Example","url":"https://gobyexample.com/generics","snippet":"Starting with version 1.18, Go has added support for generics."}` +
			`]},"error":null}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.Equal(t, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nThis tutorial introduces the basics of generics in Go.\n\nYou'll declare two simple non-generic functions."+
		"\n\n---\n\n"+
		"## Generics - Go by Example\nhttps://gobyexample.com/generics\n\nStarting with version 1.18, Go has added support for generics.", got)
}

func TestSearchReturnsEmptyForNoSources(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"data":{"sources":[]},"error":null}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"

	got, err := client.Search("nothing")
	require.NoError(t, err)
	require.Equal(t, "", got)
}

func TestSearchReturnsAPIError(t *testing.T) {
	newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"data":null,"error":{"code":"unauthorized","message":"Invalid API key"}}`))
	})

	client := NewClient()
	client.ApiKey.Value = "bad"

	_, err := client.Search("anything")
	require.ErrorContains(t, err, "Invalid API key")
}

func TestIsConfiguredRequiresApiKey(t *testing.T) {
	client := NewClient()
	require.False(t, client.IsConfigured())

	client.ApiKey.Value = "secret"
	require.True(t, client.IsConfigured())
}
