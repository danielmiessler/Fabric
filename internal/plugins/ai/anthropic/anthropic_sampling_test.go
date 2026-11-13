package anthropic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/anthropics/anthropic-sdk-go/v2/option"
	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/require"
)

const okResponse = `{"id":"msg_1","type":"message","role":"assistant","model":"m","content":[{"type":"text","text":"ok"}],"stop_reason":"end_turn","usage":{"input_tokens":1,"output_tokens":1}}`

// sentRequest sends one message through Send and returns the API version
// header and the JSON body of each request that reached the server.
func sentRequest(t *testing.T, an *Client, opts *domain.ChatOptions) (versions []string, bodies []map[string]any) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var body map[string]any
		require.NoError(t, json.Unmarshal(raw, &body))
		versions = append(versions, r.Header.Get("anthropic-version"))
		bodies = append(bodies, body)
		// Reject a request with a beta header, to make Send retry without it.
		if r.Header.Get("anthropic-beta") != "" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"type":"error","error":{"type":"invalid_request_error","message":"beta"}}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, okResponse)
	}))
	t.Cleanup(server.Close)

	an.maxTokens = 16
	an.client = anthropic.NewClient(option.WithBaseURL(server.URL), option.WithAPIKey("k"), option.WithMaxRetries(0))
	_, err := an.Send(context.Background(), []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}, opts)
	require.NoError(t, err)
	return
}

func TestSendSamplingParams(t *testing.T) {
	tests := []struct {
		name      string
		opts      domain.ChatOptions
		wantKey   string
		wantValue float64
	}{
		{"older model sends temperature",
			domain.ChatOptions{Model: "claude-sonnet-4-6", Temperature: domain.DefaultTemperature, TopP: domain.DefaultTopP},
			"temperature", domain.DefaultTemperature},
		{"older model sends top_p, not temperature, when top_p is set",
			domain.ChatOptions{Model: "claude-sonnet-4-6", Temperature: 0.3, TopP: 0.5},
			"top_p", 0.5},
		{"Vertex AI model ID",
			domain.ChatOptions{Model: "claude-opus-4-5@20251101", Temperature: 0.3, TopP: domain.DefaultTopP},
			"temperature", 0.3},
		{"Claude Sonnet 4 dated ID",
			domain.ChatOptions{Model: "claude-sonnet-4-20250514", Temperature: 0.3, TopP: domain.DefaultTopP},
			"temperature", 0.3},
		{"Claude Opus 4 on Vertex AI",
			domain.ChatOptions{Model: "claude-opus-4@20250514", Temperature: 0.3, TopP: domain.DefaultTopP},
			"temperature", 0.3},
		{"Claude 3.5 alias",
			domain.ChatOptions{Model: "claude-3-5-haiku-latest", Temperature: 0.3, TopP: domain.DefaultTopP},
			"temperature", 0.3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions, bodies := sentRequest(t, &Client{}, &tt.opts)
			require.Equal(t, []string{LegacyAPIVersion}, versions)
			for _, key := range []string{"temperature", "top_p", "top_k"} {
				if key == tt.wantKey {
					require.Equal(t, tt.wantValue, bodies[0][key])
				} else {
					require.NotContains(t, bodies[0], key)
				}
			}
		})
	}

	// These models reject sampling values, so their requests send none and
	// use the SDK's current API version. A model that is not in the allowlist,
	// such as a future one, is handled the same way.
	for _, model := range []string{"claude-opus-4-7", "claude-opus-5-5", "claude-haiku-5-5", "claude-fable-5-2", "claude-mythos-5-2", "claude-mythos-preview", "claude-opus-6"} {
		t.Run(model, func(t *testing.T) {
			versions, bodies := sentRequest(t, &Client{}, &domain.ChatOptions{Model: model, Temperature: 0.8, TopP: 0.8})
			require.Equal(t, []string{"2026-10-21"}, versions)
			for _, key := range []string{"temperature", "top_p", "top_k"} {
				require.NotContains(t, bodies[0], key)
			}
		})
	}
}

// When a request with a beta header fails, Send retries without the header.
// The retry must keep the sampling value and the legacy API version.
func TestSendBetaRetryKeepsSamplingParams(t *testing.T) {
	an := &Client{modelBetas: map[string][]string{"claude-sonnet-4-6": {"context-1m-2025-08-07"}}}
	versions, bodies := sentRequest(t, an, &domain.ChatOptions{Model: "claude-sonnet-4-6", Temperature: 0.3, TopP: domain.DefaultTopP})
	require.Equal(t, []string{LegacyAPIVersion, LegacyAPIVersion}, versions)
	for _, body := range bodies {
		require.Equal(t, 0.3, body["temperature"])
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// The Claude API gets the key as "Authorization: Bearer". A custom base URL
// gets it as X-Api-Key only, because a proxy can read only that header.
func TestConfigureKeyHeader(t *testing.T) {
	// The SDK client uses http.DefaultTransport as it is when that transport is
	// not an *http.Transport. Replace it, so that no request leaves the test.
	var auth, apiKey string
	saved := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = saved })
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		auth, apiKey = r.Header.Get("Authorization"), r.Header.Get("X-Api-Key")
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader(okResponse)),
			Request:    r,
		}, nil
	})

	for _, tt := range []struct {
		baseURL, wantAuth, wantAPIKey string
	}{
		{defaultBaseUrl, "Bearer test-key", ""},
		{"https://api.anthropic.com", "Bearer test-key", ""},
		{"", "Bearer test-key", ""},
		{"https://proxy.example.com/anthropic/", "", "test-key"},
	} {
		t.Run(tt.baseURL, func(t *testing.T) {
			auth, apiKey = "", ""
			an := NewClient()
			an.ApiBaseURL.Value = tt.baseURL
			an.ApiKey.Value = "test-key"
			require.NoError(t, an.configure())
			_, err := an.Send(context.Background(), []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}, &domain.ChatOptions{Model: "claude-opus-5-5"})
			require.NoError(t, err)
			require.Equal(t, tt.wantAuth, auth)
			require.Equal(t, tt.wantAPIKey, apiKey)
		})
	}
}
