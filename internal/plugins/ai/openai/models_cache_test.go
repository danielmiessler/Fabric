package openai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// withTempModelsCache redirects the models cache to a temporary directory for
// the duration of a test and restores the original afterwards.
func withTempModelsCache(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	orig := modelsCacheDir
	modelsCacheDir = func() (string, error) { return dir, nil }
	t.Cleanup(func() { modelsCacheDir = orig })
	return dir
}

func modelsURLFor(t *testing.T, baseURL string) string {
	t.Helper()
	full, err := url.JoinPath(baseURL, "models")
	require.NoError(t, err)
	return full
}

// writeAgedCache writes a cache entry with an explicit age so freshness and
// stale-fallback paths can be exercised deterministically.
func writeAgedCache(t *testing.T, dir, provider, fullURL string, models []string, age time.Duration) {
	t.Helper()
	entry := modelsCacheEntry{URL: fullURL, FetchedAt: time.Now().Add(-age), Models: models}
	data, err := json.Marshal(entry)
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll(dir, 0o755))
	require.NoError(t, os.WriteFile(modelsCacheFile(dir, provider, fullURL), data, 0o600))
}

func TestModelsCache_WriteThenReadFresh(t *testing.T) {
	withTempModelsCache(t)
	const provider, fullURL = "GitHub", "https://models.github.ai/catalog/models"

	require.NoError(t, writeModelsCache(provider, fullURL, []string{"a", "b"}))

	models, ok := readModelsCache(provider, fullURL, modelsCacheTTL)
	assert.True(t, ok)
	assert.Equal(t, []string{"a", "b"}, models)
}

func TestModelsCache_EmptyListNotCached(t *testing.T) {
	withTempModelsCache(t)
	const provider, fullURL = "GitHub", "https://models.github.ai/catalog/models"

	require.NoError(t, writeModelsCache(provider, fullURL, nil))

	_, ok := readModelsCache(provider, fullURL, 0)
	assert.False(t, ok)
}

func TestModelsCache_ExpiredMissesWithTTLButHitsWithoutAgeLimit(t *testing.T) {
	dir := withTempModelsCache(t)
	const provider, fullURL = "GitHub", "https://models.github.ai/catalog/models"
	writeAgedCache(t, dir, provider, fullURL, []string{"old"}, 48*time.Hour)

	_, ok := readModelsCache(provider, fullURL, modelsCacheTTL)
	assert.False(t, ok, "entry older than TTL should be a miss")

	models, ok := readModelsCache(provider, fullURL, 0)
	assert.True(t, ok, "maxAge<=0 should accept any age")
	assert.Equal(t, []string{"old"}, models)
}

func TestModelsCache_DifferentURLDoesNotCollide(t *testing.T) {
	withTempModelsCache(t)
	require.NoError(t, writeModelsCache("GitHub", "https://a/models", []string{"a"}))

	_, ok := readModelsCache("GitHub", "https://b/models", 0)
	assert.False(t, ok)
}

// A fresh cache short-circuits before any network call.
func TestFetchModelsDirectly_ServesFreshCacheWithoutRequest(t *testing.T) {
	dir := withTempModelsCache(t)
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	writeAgedCache(t, dir, "GitHub", modelsURLFor(t, srv.URL), []string{"cached-model"}, time.Minute)

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "key", "GitHub", nil)
	assert.NoError(t, err)
	assert.Equal(t, []string{"cached-model"}, models)
	assert.False(t, called, "fresh cache should prevent the network call")
}

// A 429 with a stale cache present returns the stale list rather than erroring.
func TestFetchModelsDirectly_429ServesStaleCache(t *testing.T) {
	dir := withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte("<html>Whoa there!</html>"))
	}))
	defer srv.Close()

	// Stale so the TTL check misses and the request is actually made.
	writeAgedCache(t, dir, "GitHub", modelsURLFor(t, srv.URL), []string{"stale-model"}, 48*time.Hour)

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "key", "GitHub", nil)
	assert.NoError(t, err)
	assert.Equal(t, []string{"stale-model"}, models)
}

// A 429 with no cache yields a concise message, not the raw HTML body.
func TestFetchModelsDirectly_429NoCacheCleanError(t *testing.T) {
	withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte("<html><title>Rate limit</title>Whoa there!</html>"))
	}))
	defer srv.Close()

	_, err := FetchModelsDirectly(context.Background(), srv.URL, "key", "GitHub", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rate limit")
	assert.Contains(t, err.Error(), "60")
	assert.NotContains(t, err.Error(), "<html>")
	assert.NotContains(t, err.Error(), "Whoa there")
}

// A successful fetch is cached, so a later failure is served from cache.
func TestFetchModelsDirectly_WritesCacheOnSuccess(t *testing.T) {
	withTempModelsCache(t)
	fail := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fail {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":"m1"}]`))
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "key", "GitHub", nil)
	require.NoError(t, err)
	assert.Equal(t, []string{"m1"}, models)

	// Cache is fresh now, so even with the server failing we get the cached list.
	fail = true
	models, err = FetchModelsDirectly(context.Background(), srv.URL, "key", "GitHub", nil)
	require.NoError(t, err)
	assert.Equal(t, []string{"m1"}, models)
}
