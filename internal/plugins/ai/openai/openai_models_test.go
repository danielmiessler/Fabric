package openai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensures we can fetch models directly when a provider returns a direct array of models
// instead of the standard OpenAI list response structure.
func TestFetchModelsDirectly_DirectArray(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[{"id":"github-model"}]`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, "github-model", models[0])
}

// Ensures we can fetch models when a provider returns the standard OpenAI format
func TestFetchModelsDirectly_OpenAIFormat(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"data":[{"id":"openai-model"}]}`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, "openai-model", models[0])
}

// Ensures we handle empty model lists correctly
func TestFetchModelsDirectly_EmptyArray(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[]`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(models))
}
