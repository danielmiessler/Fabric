package codex

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/shared/constant"
)

func TestBuildAuthorizeURLIncludesPKCE(t *testing.T) {
	pkce := pkceCodes{
		CodeVerifier:  "verifier",
		CodeChallenge: "challenge",
	}
	redirectURL := fmt.Sprintf("http://localhost:%d/auth/callback", defaultCallbackPort)

	authURL, err := buildAuthorizeURL(defaultAuthBaseURL, redirectURL, pkce, "state-123")
	if err != nil {
		t.Fatalf("buildAuthorizeURL() error = %v", err)
	}

	parsed, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}

	if got := parsed.Query().Get("client_id"); got != oauthClientID {
		t.Fatalf("client_id = %q, want %q", got, oauthClientID)
	}
	if got := parsed.Query().Get("code_challenge"); got != pkce.CodeChallenge {
		t.Fatalf("code_challenge = %q, want %q", got, pkce.CodeChallenge)
	}
	if got := parsed.Query().Get("state"); got != "state-123" {
		t.Fatalf("state = %q, want %q", got, "state-123")
	}
	if got := parsed.Query().Get("redirect_uri"); got != redirectURL {
		t.Fatalf("redirect_uri = %q, want %q", got, redirectURL)
	}
	if got := parsed.Query().Get("originator"); got != defaultOriginator {
		t.Fatalf("originator = %q, want %q", got, defaultOriginator)
	}
}

func TestRunOAuthFlowCompletesWithCallback(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			http.NotFound(w, r)
			return
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("ParseForm() error = %v", err)
		}
		if got := r.Form.Get("grant_type"); got != "authorization_code" {
			t.Fatalf("grant_type = %q, want authorization_code", got)
		}

		_ = json.NewEncoder(w).Encode(oauthTokens{
			IDToken:      testJWT("acct_oauth", time.Now().Add(time.Hour)),
			AccessToken:  testJWT("acct_oauth", time.Now().Add(time.Hour)),
			RefreshToken: "refresh-oauth",
		})
	}))
	defer authServer.Close()

	client := NewClient()
	client.AuthBaseURL.Value = authServer.URL

	openBrowserFn := func(authURL string) error {
		parsed, err := url.Parse(authURL)
		if err != nil {
			return err
		}
		callbackURL := parsed.Query().Get("redirect_uri")
		callbackParsed, err := url.Parse(callbackURL)
		if err != nil {
			t.Fatalf("url.Parse(callbackURL) error = %v", err)
		}
		if callbackParsed.Path != oauthCallbackPath {
			t.Fatalf("callback path = %q, want %q", callbackParsed.Path, oauthCallbackPath)
		}
		host, port, err := net.SplitHostPort(callbackParsed.Host)
		if err != nil {
			t.Fatalf("SplitHostPort(%q) error = %v", callbackParsed.Host, err)
		}
		if host != "localhost" {
			t.Fatalf("callback host = %q, want %q", host, "localhost")
		}
		if port != fmt.Sprintf("%d", defaultCallbackPort) {
			t.Fatalf("callback port = %q, want %d", port, defaultCallbackPort)
		}
		state := parsed.Query().Get("state")

		go func() {
			_, _ = http.Get(callbackURL + "?code=auth-code&state=" + url.QueryEscape(state))
		}()

		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokens, err := client.runOAuthFlow(ctx, openBrowserFn)
	if err != nil {
		t.Fatalf("runOAuthFlow() error = %v", err)
	}

	if tokens.RefreshToken != "refresh-oauth" {
		t.Fatalf("RefreshToken = %q, want %q", tokens.RefreshToken, "refresh-oauth")
	}
	accountID, err := client.extractAccountID(tokens.IDToken, tokens.AccessToken)
	if err != nil {
		t.Fatalf("extractAccountID() error = %v", err)
	}
	if accountID != "acct_oauth" {
		t.Fatalf("accountID = %q, want %q", accountID, "acct_oauth")
	}
}

func TestBuildCodexResponseParamsMovesSystemPromptToInstructions(t *testing.T) {
	client := NewClient()

	req := client.buildCodexResponseParams([]*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "System guidance"},
		{Role: chat.ChatMessageRoleDeveloper, Content: "Developer guidance"},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}, &domain.ChatOptions{
		Model:       "gpt-5.4",
		Temperature: 0.7,
	})

	if got := req.Instructions.Value; got != "System guidance\n\nDeveloper guidance" {
		t.Fatalf("instructions = %q, want concatenated system/developer guidance", got)
	}
	if len(req.Input.OfInputItemList) != 1 {
		t.Fatalf("input length = %d, want 1 user message", len(req.Input.OfInputItemList))
	}
}

func TestListModelsFiltersSupportedVisibleModels(t *testing.T) {
	modelsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got == "" {
			t.Fatalf("Authorization header missing")
		}
		if got := r.Header.Get("ChatGPT-Account-ID"); got != "acct_models" {
			t.Fatalf("ChatGPT-Account-ID = %q, want %q", got, "acct_models")
		}
		if got := r.URL.Query().Get("client_version"); got == "" {
			t.Fatalf("client_version query parameter missing")
		}

		_ = json.NewEncoder(w).Encode(modelsResponse{
			Models: []modelInfo{
				{Slug: "gpt-5.4", SupportedInAPI: true, Visibility: "list"},
				{Slug: "gpt-5-hidden", SupportedInAPI: true, Visibility: "hide"},
				{Slug: "gpt-5-disabled", SupportedInAPI: false, Visibility: "list"},
			},
		})
	}))
	defer modelsServer.Close()

	client := newConfiguredTestClient(t, modelsServer.URL, "acct_models", testJWT("acct_models", time.Now().Add(time.Hour)))

	models, err := client.ListModels(context.Background())
	if err != nil {
		t.Fatalf("ListModels() error = %v", err)
	}
	if len(models) != 1 || models[0] != "gpt-5.4" {
		t.Fatalf("ListModels() = %#v, want []string{\"gpt-5.4\"}", models)
	}
}

func TestNormalizeSemverLikeVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "prefixed version", input: "v1.4.434", want: "1.4.434"},
		{name: "plain semver", input: "1.2.3", want: "1.2.3"},
		{name: "suffix trimmed", input: "1.2.3-dev", want: "1.2.3"},
		{name: "devel build ignored", input: "(devel)", want: ""},
		{name: "short version rejected", input: "v1.2", want: ""},
		{name: "invalid version rejected", input: "invalid", want: ""},
		{name: "whitespace trimmed", input: " v2.3.4 \n", want: "2.3.4"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := normalizeSemverLikeVersion(tc.input); got != tc.want {
				t.Fatalf("normalizeSemverLikeVersion(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestMapRequestErrorPreservesCodexAPIErrorMessage(t *testing.T) {
	client := NewClient()
	apiErr := &openaiapi.Error{StatusCode: http.StatusBadRequest}
	if err := apiErr.UnmarshalJSON([]byte(`{"message":"The requested model is not supported.","type":"invalid_request_error","param":"model","code":"invalid_value"}`)); err != nil {
		t.Fatalf("apiErr.UnmarshalJSON() error = %v", err)
	}

	err := client.mapRequestError(apiErr)
	if err == nil {
		t.Fatal("mapRequestError() returned nil")
	}
	if got := err.Error(); got != "codex request failed with status 400" {
		t.Fatalf("mapRequestError() = %q, want %q", got, "codex request failed with status 400")
	}
	if unwrapped := errors.Unwrap(err); unwrapped == nil || !strings.Contains(unwrapped.Error(), "The requested model is not supported.") {
		t.Fatalf("wrapped error = %v, want provider detail", unwrapped)
	}
}

func TestMapRequestErrorReadsAPIErrorResponseBodyWhenRawJSONMissing(t *testing.T) {
	client := NewClient()
	apiErr := &openaiapi.Error{
		StatusCode: http.StatusBadRequest,
		Response: &http.Response{
			Body: io.NopCloser(strings.NewReader(`{"detail":"The requested model is not supported for Codex."}`)),
		},
	}

	err := client.mapRequestError(apiErr)
	if err == nil {
		t.Fatal("mapRequestError() returned nil")
	}
	if got := err.Error(); got != "codex request failed with status 400" {
		t.Fatalf("mapRequestError() = %q, want %q", got, "codex request failed with status 400")
	}
	if unwrapped := errors.Unwrap(err); unwrapped == nil || !strings.Contains(unwrapped.Error(), "The requested model is not supported for Codex.") {
		t.Fatalf("wrapped error = %v, want provider detail", unwrapped)
	}
}

func TestSendRefreshesAfterUnauthorized(t *testing.T) {
	var responseCalls atomic.Int32
	var captureMu sync.Mutex
	var seenAuthHeaders []string
	var requestBodies []string

	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(refreshResponse{
			IDToken:      testJWT("acct_refresh", time.Now().Add(2*time.Hour)),
			AccessToken:  testJWT("acct_refresh", time.Now().Add(2*time.Hour)),
			RefreshToken: "refresh-new",
		})
	}))
	defer authServer.Close()

	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/responses" {
			http.NotFound(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("ReadAll(r.Body) error = %v", err)
		}
		captureMu.Lock()
		seenAuthHeaders = append(seenAuthHeaders, r.Header.Get("Authorization"))
		requestBodies = append(requestBodies, string(body))
		captureMu.Unlock()
		call := responseCalls.Add(1)
		if call == 1 {
			http.Error(w, `{"error":{"message":"expired","code":"invalid_token"}}`, http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatalf("response writer does not implement http.Flusher")
		}
		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type":  string(constant.ResponseOutputTextDelta("").Default()),
			"delta": "hello from codex",
		}))
		flusher.Flush()
		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type": "response.completed",
			"response": map[string]any{
				"output": []any{
					map[string]any{
						"type": "message",
						"content": []any{
							map[string]any{
								"type":        "output_text",
								"text":        "hello from codex",
								"annotations": []any{},
							},
						},
					},
				},
			},
		}))
		flusher.Flush()
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	defer apiServer.Close()

	client := NewClient()
	client.ApiBaseURL.Value = apiServer.URL
	client.AuthBaseURL.Value = authServer.URL
	client.RefreshToken.Value = "refresh-old"
	client.AccessToken.Value = testJWT("acct_refresh", time.Now().Add(time.Hour))
	client.AccountID.Value = "acct_refresh"

	if err := client.configure(); err != nil {
		t.Fatalf("configure() error = %v", err)
	}

	message, err := client.Send(context.Background(), []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}, &domain.ChatOptions{
		Model:       "gpt-5.4",
		Temperature: 0.7,
	})
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}

	if message != "hello from codex" {
		t.Fatalf("Send() = %q, want %q", message, "hello from codex")
	}
	if responseCalls.Load() != 2 {
		t.Fatalf("response call count = %d, want 2", responseCalls.Load())
	}
	if len(seenAuthHeaders) != 2 {
		t.Fatalf("seenAuthHeaders length = %d, want 2", len(seenAuthHeaders))
	}
	if seenAuthHeaders[0] == seenAuthHeaders[1] {
		t.Fatalf("expected second request to use refreshed bearer token, got %#v", seenAuthHeaders)
	}
	if len(requestBodies) != 2 {
		t.Fatalf("requestBodies length = %d, want 2", len(requestBodies))
	}
	if !strings.Contains(requestBodies[1], `"instructions":"You are a helpful assistant."`) {
		t.Fatalf("request body missing fallback instructions: %s", requestBodies[1])
	}
	if !strings.Contains(requestBodies[1], `"stream":true`) {
		t.Fatalf("stream request body missing stream=true: %s", requestBodies[1])
	}
}

func TestSendIncludesSourcesFromAnnotatedResponse(t *testing.T) {
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/responses" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatalf("response writer does not implement http.Flusher")
		}
		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type": "response.completed",
			"response": map[string]any{
				"output": []any{
					map[string]any{
						"type": "message",
						"content": []any{
							map[string]any{
								"type": "output_text",
								"text": "hello from codex",
								"annotations": []any{
									map[string]any{
										"type":  "url_citation",
										"title": "Example",
										"url":   "https://example.com",
									},
								},
							},
						},
					},
				},
			},
		}))
		flusher.Flush()
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	defer apiServer.Close()

	client := newConfiguredTestClient(t, apiServer.URL, "acct_sources", testJWT("acct_sources", time.Now().Add(time.Hour)))

	message, err := client.Send(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}, &domain.ChatOptions{
		Model:       "gpt-5.4",
		Temperature: 0.7,
		Search:      true,
	})
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}

	if !strings.Contains(message, "hello from codex") {
		t.Fatalf("Send() missing response text: %q", message)
	}
	if !strings.Contains(message, "## Sources") {
		t.Fatalf("Send() missing sources section: %q", message)
	}
	if !strings.Contains(message, "[Example](https://example.com)") {
		t.Fatalf("Send() missing expected citation: %q", message)
	}
}

func TestSendFallsBackToDeltaWhenCompletedResponseHasNoText(t *testing.T) {
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/responses" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatalf("response writer does not implement http.Flusher")
		}
		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type":  string(constant.ResponseOutputTextDelta("").Default()),
			"delta": "hello from delta",
		}))
		flusher.Flush()
		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type": "response.completed",
			"response": map[string]any{
				"output": []any{
					map[string]any{
						"type": "message",
						"content": []any{
							map[string]any{
								"type": "output_text",
								"text": "",
							},
						},
					},
				},
			},
		}))
		flusher.Flush()
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	defer apiServer.Close()

	client := newConfiguredTestClient(t, apiServer.URL, "acct_delta_fallback", testJWT("acct_delta_fallback", time.Now().Add(time.Hour)))

	message, err := client.Send(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}, &domain.ChatOptions{
		Model:       "gpt-5.4",
		Temperature: 0.7,
	})
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}

	if message != "hello from delta" {
		t.Fatalf("Send() = %q, want %q", message, "hello from delta")
	}
}

func TestSendStreamReadsCodexSSE(t *testing.T) {
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/responses" {
			http.NotFound(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("ReadAll(r.Body) error = %v", err)
		}
		if !strings.Contains(string(body), `"instructions":"Follow the system prompt"`) {
			t.Fatalf("request body missing system instructions: %s", string(body))
		}
		if strings.Contains(string(body), `"role":"system"`) {
			t.Fatalf("request body should not keep system messages in input: %s", string(body))
		}

		w.Header().Set("Content-Type", "text/event-stream")
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatalf("response writer does not implement http.Flusher")
		}

		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type":  string(constant.ResponseOutputTextDelta("").Default()),
			"delta": "hello",
		}))
		flusher.Flush()

		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type":  string(constant.ResponseOutputTextDelta("").Default()),
			"delta": " world",
		}))
		flusher.Flush()

		fmt.Fprintf(w, "data: %s\n\n", marshalJSON(t, map[string]any{
			"type": string(constant.ResponseOutputTextDone("").Default()),
			"text": "hello world",
		}))
		flusher.Flush()
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	defer apiServer.Close()

	client := newConfiguredTestClient(t, apiServer.URL, "acct_stream", testJWT("acct_stream", time.Now().Add(time.Hour)))

	updates := make(chan domain.StreamUpdate, 8)
	err := client.SendStream(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "Follow the system prompt"},
		{Role: "user", Content: "Hello"},
	}, &domain.ChatOptions{
		Model:       "gpt-5.4",
		Temperature: 0.7,
	}, updates)
	if err != nil {
		t.Fatalf("SendStream() error = %v", err)
	}

	var builder strings.Builder
	for update := range updates {
		if update.Type == domain.StreamTypeContent {
			builder.WriteString(update.Content)
		}
	}

	if builder.String() != "hello world\n" {
		t.Fatalf("streamed content = %q, want %q", builder.String(), "hello world\n")
	}
}

func TestSendStreamClosesChannelAndMapsHTTPError(t *testing.T) {
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/responses" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, `{"error":{"message":"usage limit reached"}}`, http.StatusTooManyRequests)
	}))
	defer apiServer.Close()

	client := newConfiguredTestClient(t, apiServer.URL, "acct_stream_error", testJWT("acct_stream_error", time.Now().Add(time.Hour)))

	updates := make(chan domain.StreamUpdate, 1)
	err := client.SendStream(context.Background(), []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}, &domain.ChatOptions{
		Model: "gpt-5.4",
	}, updates)
	if err == nil {
		t.Fatal("SendStream() error = nil, want mapped HTTP error")
	}
	if got := err.Error(); got != "codex usage limit reached" {
		t.Fatalf("SendStream() error = %q, want %q", got, "codex usage limit reached")
	}

	update, ok := <-updates
	if ok {
		t.Fatalf("expected closed channel after stream error, got update %#v", update)
	}
}

func newConfiguredTestClient(t *testing.T, apiBaseURL string, accountID string, accessToken string) *Client {
	t.Helper()

	client := NewClient()
	client.ApiBaseURL.Value = apiBaseURL
	client.AuthBaseURL.Value = defaultAuthBaseURL
	client.RefreshToken.Value = "refresh-test"
	client.AccessToken.Value = accessToken
	client.AccountID.Value = accountID

	if err := client.configure(); err != nil {
		t.Fatalf("configure() error = %v", err)
	}

	return client
}

func testJWT(accountID string, expiry time.Time) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none","typ":"JWT"}`))
	payloadBytes, _ := json.Marshal(map[string]any{
		"exp": expiry.Unix(),
		"https://api.openai.com/auth": map[string]any{
			"chatgpt_account_id": accountID,
			"chatgpt_plan_type":  "plus",
		},
	})
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	signature := base64.RawURLEncoding.EncodeToString([]byte("sig"))
	return header + "." + payload + "." + signature
}

func marshalJSON(t *testing.T, value any) string {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	return string(encoded)
}
