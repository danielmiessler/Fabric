package openai

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensures we can fetch models directly when a provider returns a direct array of models
// instead of the standard OpenAI list response structure.
func TestFetchModelsDirectly_DirectArray(t *testing.T) {
	withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[{"id":"github-model"}]`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider", nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, "github-model", models[0])
}

// Ensures we can fetch models when a provider returns the standard OpenAI format
func TestFetchModelsDirectly_OpenAIFormat(t *testing.T) {
	withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"data":[{"id":"openai-model"}]}`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider", nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(models))
	assert.Equal(t, "openai-model", models[0])
}

// Ensures non-GitHub hosts do not receive the GitHub-specific API version header.
func TestFetchModelsDirectly_NoGitHubHeaderForOtherHosts(t *testing.T) {
	withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.Header.Get("X-GitHub-Api-Version"))
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[{"id":"some-model"}]`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	_, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider", nil)
	assert.NoError(t, err)
}

// captureRoundTripper records the outgoing request and returns a canned response,
// allowing assertions against requests sent to real hosts without networking.
type captureRoundTripper struct {
	req  *http.Request
	body string
}

func (c *captureRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	c.req = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(c.body)),
		Request:    req,
	}, nil
}

// Ensures the GitHub Models catalog host receives the documented API version
// header so it is not throttled by GitHub's edge rate limiter.
func TestFetchModelsDirectly_GitHubHostSendsAPIVersionHeader(t *testing.T) {
	withTempModelsCache(t)
	rt := &captureRoundTripper{body: `[{"id":"github-model"}]`}
	client := &http.Client{Transport: rt}

	models, err := FetchModelsDirectly(context.Background(), "https://models.github.ai/catalog", "test-key", "GitHub", client)
	assert.NoError(t, err)
	assert.Equal(t, []string{"github-model"}, models)
	assert.Equal(t, "2022-11-28", rt.req.Header.Get("X-GitHub-Api-Version"))
}

// Ensures we handle empty model lists correctly
func TestFetchModelsDirectly_EmptyArray(t *testing.T) {
	withTempModelsCache(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[]`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	models, err := FetchModelsDirectly(context.Background(), srv.URL, "test-key", "TestProvider", nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(models))
}
