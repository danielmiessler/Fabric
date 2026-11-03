package keenable

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func withServer(t *testing.T, handler http.HandlerFunc) {
	server := httptest.NewServer(handler)
	oldPublic, oldKeyed := publicSearchURL, keyedSearchURL
	publicSearchURL, keyedSearchURL = server.URL+"/public", server.URL+"/keyed"
	t.Cleanup(func() {
		publicSearchURL, keyedSearchURL = oldPublic, oldKeyed
		server.Close()
	})
}

func TestSearchWithoutKeyUsesPublicEndpoint(t *testing.T) {
	withServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/public", r.URL.Path)
		require.Empty(t, r.Header.Get("X-API-Key"))
		require.Equal(t, "Fabric", r.Header.Get("X-Keenable-Title"))
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		require.Equal(t, "golang generics", body["query"])
		require.EqualValues(t, 10, body["max_results"])
		_, _ = w.Write([]byte(`{"results":[
			{"title":"Go generics","url":"https://go.dev/doc/tutorial/generics","description":"Getting started with generics.","snippet":"long"},
			{"title":"No URL","description":"dropped"},
			{"title":"Snippet only","url":"https://example.com","snippet":"` + strings.Repeat("x", 1500) + `"}
		]}`))
	})

	client := NewClient()
	got, err := client.Search("golang generics")
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(got, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nGetting started with generics.\n\n## Snippet only\nhttps://example.com\n\n"))
	require.NotContains(t, got, "No URL")
	require.True(t, strings.HasSuffix(got, strings.Repeat("x", 1000)+"…"))
}

func TestSearchWithKeyUsesKeyedEndpoint(t *testing.T) {
	withServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/keyed", r.URL.Path)
		require.Equal(t, "secret", r.Header.Get("X-API-Key"))
		_, _ = w.Write([]byte(`{"results":[]}`))
	})

	client := NewClient()
	client.ApiKey.Value = "secret"
	_, err := client.Search("anything")
	require.NoError(t, err)
}

func TestSearchReturnsAPIError(t *testing.T) {
	withServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail":"rate limited"}`))
	})

	_, err := NewClient().Search("anything")
	require.ErrorContains(t, err, "rate limited")
}
