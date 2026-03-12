package openai

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchModelsDirectly_ShapesRequestAndTruncatesErrorBody(t *testing.T) {
	t.Parallel()

	providerName := "TestProvider"
	apiKey := "test-key"
	longBody := strings.Repeat("x", errorResponseLimit+64)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/models", r.URL.Path)
		assert.Equal(t, "Bearer "+apiKey, r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusTeapot)
		_, err := w.Write([]byte(longBody))
		require.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, apiKey, providerName, srv.Client())
	require.Error(t, err)
	assert.Nil(t, models)
	assert.ErrorContains(t, err, "418")
	assert.ErrorContains(t, err, providerName)
	assert.ErrorContains(t, err, strings.Repeat("x", 64))
	assert.NotContains(t, err.Error(), longBody)
}

func TestFetchModelsDirectly_RejectsOversizedResponse(t *testing.T) {
	t.Parallel()

	providerName := "SizeGuardProvider"
	oversizedBody := strings.Repeat("a", maxResponseSize+1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(oversizedBody))
		require.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", providerName, srv.Client())
	require.Error(t, err)
	assert.Nil(t, models)
	assert.ErrorContains(t, err, providerName)
	assert.ErrorContains(t, err, fmt.Sprintf("%d", maxResponseSize))
}

func TestListModels_UsesSDKForStandardOpenAIResponses(t *testing.T) {
	t.Parallel()

	var requestCount atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.Add(1)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/models", r.URL.Path)
		assert.Equal(t, "Bearer sdk-key", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"object":"list","data":[{"id":"sdk-model","object":"model","owned_by":"fabric"}]}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredModelsTestClient(t, "SDKProvider", srv.URL, "sdk-key")

	models, err := client.ListModels()
	require.NoError(t, err)
	assert.Equal(t, []string{"sdk-model"}, models)
	assert.Equal(t, int32(1), requestCount.Load())
}

func TestListModels_FallsBackToDirectFetchForDirectArrayProviders(t *testing.T) {
	t.Parallel()

	var requestCount atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := requestCount.Add(1)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/models", r.URL.Path)
		assert.Equal(t, "Bearer fallback-key", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		if current == 1 {
			_, err := w.Write([]byte(`{"data":`))
			require.NoError(t, err)
			return
		}
		_, err := w.Write([]byte(`[{"id":"fallback-model"}]`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := newConfiguredModelsTestClient(t, "FallbackProvider", srv.URL, "fallback-key")

	models, err := client.ListModels()
	require.NoError(t, err)
	assert.Equal(t, []string{"fallback-model"}, models)
	assert.Equal(t, int32(2), requestCount.Load())
}

func newConfiguredModelsTestClient(t *testing.T, vendorName, baseURL, apiKey string) *Client {
	t.Helper()

	client := NewClientCompatible(vendorName, baseURL, nil)
	client.ApiKey.Value = apiKey
	client.ApiBaseURL.Value = baseURL
	require.NoError(t, client.configure())
	return client
}
