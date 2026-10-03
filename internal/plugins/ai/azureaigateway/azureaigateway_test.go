package azureaigateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
)

// TestMain pins the locale to English so that i18n.T() assertions
// are deterministic regardless of the CI machine's system locale.
func TestMain(m *testing.M) {
	if _, err := i18n.Init("en"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to init i18n: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestBedrockBuildEndpoint(t *testing.T) {
	b := NewBedrockBackend("key")
	got := b.BuildEndpoint("https://gw.example.com", "us.anthropic.claude-3-haiku-20240307-v1:0")
	// url.PathEscape keeps the colon, which is valid in a path segment.
	want := "https://gw.example.com/model/us.anthropic.claude-3-haiku-20240307-v1:0/invoke"
	if got != want {
		t.Errorf("BuildEndpoint() = %q, want %q", got, want)
	}
}

func TestBedrockBuildEndpointTrailingSlash(t *testing.T) {
	b := NewBedrockBackend("key")
	got := b.BuildEndpoint("https://gw.example.com/", "model-id")
	want := "https://gw.example.com/model/model-id/invoke"
	if got != want {
		t.Errorf("BuildEndpoint() = %q, want %q", got, want)
	}
}

func TestBedrockAuthHeader(t *testing.T) {
	b := NewBedrockBackend("my-key")
	name, value := b.AuthHeader()
	if name != "Authorization" {
		t.Errorf("AuthHeader name = %q, want %q", name, "Authorization")
	}
	if value != "Bearer my-key" {
		t.Errorf("AuthHeader value = %q, want %q", value, "Bearer my-key")
	}
}

func TestBedrockListModels(t *testing.T) {
	b := NewBedrockBackend("key")
	models, err := b.ListModels(context.Background())
	if err != nil {
		t.Fatalf("ListModels() error = %v", err)
	}
	if len(models) == 0 {
		t.Fatal("ListModels() returned empty list")
	}
}

func TestBedrockPrepareRequestSystemMessages(t *testing.T) {
	b := NewBedrockBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "You are a helpful assistant."},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	systemField, ok := body["system"]
	if !ok {
		t.Fatal("expected 'system' field in request body")
	}
	if systemField != "You are a helpful assistant." {
		t.Errorf("system = %q, want %q", systemField, "You are a helpful assistant.")
	}

	messages := body["messages"].([]any)
	if len(messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(messages))
	}
	msg := messages[0].(map[string]any)
	if msg["role"] != "user" {
		t.Errorf("message role = %q, want %q", msg["role"], "user")
	}
}

func TestBedrockPrepareRequestMaxTokensDefault(t *testing.T) {
	b := NewBedrockBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	json.Unmarshal(bodyBytes, &body)

	maxTokens := int(body["max_tokens"].(float64))
	if maxTokens != 4096 {
		t.Errorf("max_tokens = %d, want 4096 (default)", maxTokens)
	}
}

func TestBedrockPrepareRequestMaxTokensCustom(t *testing.T) {
	b := NewBedrockBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
		MaxTokens:   8192,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	json.Unmarshal(bodyBytes, &body)

	maxTokens := int(body["max_tokens"].(float64))
	if maxTokens != 8192 {
		t.Errorf("max_tokens = %d, want 8192", maxTokens)
	}
}

func TestBedrockPrepareRequestSkipsEmptyMessages(t *testing.T) {
	b := NewBedrockBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
		{Role: chat.ChatMessageRoleUser, Content: "   "},
		{Role: chat.ChatMessageRoleUser, Content: ""},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	json.Unmarshal(bodyBytes, &body)

	messages := body["messages"].([]any)
	if len(messages) != 1 {
		t.Errorf("expected 1 message after filtering, got %d", len(messages))
	}
}

func TestBedrockParseResponse(t *testing.T) {
	b := NewBedrockBackend("key")
	respJSON := `{"content":[{"type":"text","text":"Hello world"}]}`
	result, err := b.ParseResponse([]byte(respJSON))
	if err != nil {
		t.Fatalf("ParseResponse() error = %v", err)
	}
	if result != "Hello world" {
		t.Errorf("ParseResponse() = %q, want %q", result, "Hello world")
	}
}

func TestBedrockParseResponseMultipleBlocks(t *testing.T) {
	b := NewBedrockBackend("key")
	respJSON := `{"content":[{"type":"text","text":"Hello "},{"type":"text","text":"world"}]}`
	result, err := b.ParseResponse([]byte(respJSON))
	if err != nil {
		t.Fatalf("ParseResponse() error = %v", err)
	}
	if result != "Hello world" {
		t.Errorf("ParseResponse() = %q, want %q", result, "Hello world")
	}
}

func TestBedrockParseResponseNoTextBlocks(t *testing.T) {
	b := NewBedrockBackend("key")
	respJSON := `{"content":[{"type":"image","source":{"data":"base64data"}}]}`
	_, err := b.ParseResponse([]byte(respJSON))
	if err == nil {
		t.Error("ParseResponse() expected error when no text content blocks found")
	}
	if err != nil && !strings.Contains(err.Error(), "no text content blocks") {
		t.Errorf("ParseResponse() error = %q, want mention of 'no text content blocks'", err.Error())
	}
}

func TestBedrockParseResponseInvalid(t *testing.T) {
	b := NewBedrockBackend("key")
	_, err := b.ParseResponse([]byte("not json"))
	if err == nil {
		t.Error("ParseResponse() expected error for invalid JSON")
	}
}

func TestAzureOpenAIBuildEndpoint(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	got := b.BuildEndpoint("https://gw.example.com", "gpt-4o")
	want := "https://gw.example.com/openai/deployments/gpt-4o/chat/completions?api-version=2025-04-01-preview"
	if got != want {
		t.Errorf("BuildEndpoint() = %q, want %q", got, want)
	}
}

func TestAzureOpenAIAuthHeader(t *testing.T) {
	b := NewAzureOpenAIBackend("my-key", "")
	name, value := b.AuthHeader()
	if name != "api-key" {
		t.Errorf("AuthHeader name = %q, want %q", name, "api-key")
	}
	if value != "my-key" {
		t.Errorf("AuthHeader value = %q, want %q", value, "my-key")
	}
}

func TestAzureOpenAIListModels(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	models, err := b.ListModels(context.Background())
	if err != nil {
		t.Fatalf("ListModels() error = %v", err)
	}
	if len(models) == 0 {
		t.Fatal("ListModels() returned empty list")
	}
}

func TestAzureOpenAIPrepareRequest(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "You are helpful."},
		{Role: chat.ChatMessageRoleUser, Content: "Hi"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	json.Unmarshal(bodyBytes, &body)

	messages := body["messages"].([]any)
	if len(messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(messages))
	}
	firstMsg := messages[0].(map[string]any)
	if firstMsg["role"] != "system" {
		t.Errorf("first message role = %q, want %q", firstMsg["role"], "system")
	}
}

func TestAzureOpenAIParseResponse(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	respJSON := `{"choices":[{"message":{"content":"Hello!"}}]}`
	result, err := b.ParseResponse([]byte(respJSON))
	if err != nil {
		t.Fatalf("ParseResponse() error = %v", err)
	}
	if result != "Hello!" {
		t.Errorf("ParseResponse() = %q, want %q", result, "Hello!")
	}
}

func TestAzureOpenAIParseResponseNoChoices(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	_, err := b.ParseResponse([]byte(`{"choices":[]}`))
	if err == nil {
		t.Error("ParseResponse() expected error for empty choices")
	}
}

func TestVertexAIBuildEndpoint(t *testing.T) {
	b := NewVertexAIBackend("key")
	got := b.BuildEndpoint("https://gw.example.com", "gemini-2.0-flash")
	want := "https://gw.example.com/publishers/google/models/gemini-2.0-flash:generateContent"
	if got != want {
		t.Errorf("BuildEndpoint() = %q, want %q", got, want)
	}
}

func TestVertexAIAuthHeader(t *testing.T) {
	b := NewVertexAIBackend("my-key")
	name, value := b.AuthHeader()
	if name != "x-goog-api-key" {
		t.Errorf("AuthHeader name = %q, want %q", name, "x-goog-api-key")
	}
	if value != "my-key" {
		t.Errorf("AuthHeader value = %q, want %q", value, "my-key")
	}
}

func TestVertexAIListModels(t *testing.T) {
	b := NewVertexAIBackend("key")
	models, err := b.ListModels(context.Background())
	if err != nil {
		t.Fatalf("ListModels() error = %v", err)
	}
	if len(models) == 0 {
		t.Fatal("ListModels() returned empty list")
	}
}

func TestVertexAIPrepareRequestSystemMessages(t *testing.T) {
	b := NewVertexAIBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "You are a helpful assistant."},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	si, ok := body["systemInstruction"]
	if !ok {
		t.Fatal("expected 'systemInstruction' field in request body")
	}
	siMap := si.(map[string]any)
	parts := siMap["parts"].([]any)
	firstPart := parts[0].(map[string]any)
	if firstPart["text"] != "You are a helpful assistant." {
		t.Errorf("systemInstruction text = %q, want %q", firstPart["text"], "You are a helpful assistant.")
	}

	contents := body["contents"].([]any)
	if len(contents) != 1 {
		t.Fatalf("expected 1 content entry, got %d", len(contents))
	}
	content := contents[0].(map[string]any)
	if content["role"] != "user" {
		t.Errorf("content role = %q, want %q", content["role"], "user")
	}
}

func TestVertexAIPrepareRequestAssistantRole(t *testing.T) {
	b := NewVertexAIBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
		{Role: chat.ChatMessageRoleAssistant, Content: "Hi there"},
	}
	opts := &domain.ChatOptions{
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	bodyBytes, err := b.PrepareRequest(msgs, opts)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	var body map[string]any
	json.Unmarshal(bodyBytes, &body)

	contents := body["contents"].([]any)
	if len(contents) != 2 {
		t.Fatalf("expected 2 contents, got %d", len(contents))
	}
	secondContent := contents[1].(map[string]any)
	if secondContent["role"] != "model" {
		t.Errorf("assistant role should be mapped to 'model', got %q", secondContent["role"])
	}
}

func TestVertexAIParseResponse(t *testing.T) {
	b := NewVertexAIBackend("key")
	respJSON := `{"candidates":[{"content":{"parts":[{"text":"Hello world"}]}}]}`
	result, err := b.ParseResponse([]byte(respJSON))
	if err != nil {
		t.Fatalf("ParseResponse() error = %v", err)
	}
	if result != "Hello world" {
		t.Errorf("ParseResponse() = %q, want %q", result, "Hello world")
	}
}

func TestVertexAIParseResponseNoCandidates(t *testing.T) {
	b := NewVertexAIBackend("key")
	_, err := b.ParseResponse([]byte(`{"candidates":[]}`))
	if err == nil {
		t.Error("ParseResponse() expected error for empty candidates")
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Fatal("NewClient() returned nil")
	}
	if c.BackendType == nil {
		t.Error("BackendType setup question not initialized")
	}
	if c.GatewayURL == nil {
		t.Error("GatewayURL setup question not initialized")
	}
	if c.SubscriptionKey == nil {
		t.Error("SubscriptionKey setup question not initialized")
	}
}

func TestConfigureRequiresGatewayURL(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = ""
	c.SubscriptionKey.Value = "key"

	err := c.configure()
	if err == nil {
		t.Error("configure() expected error for empty gateway URL")
	}
}

func TestConfigureRequiresHTTPS(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "http://gw.example.com"
	c.SubscriptionKey.Value = "key"

	err := c.configure()
	if err == nil {
		t.Error("configure() expected error for HTTP (non-HTTPS) URL")
	}
	if err != nil && !strings.Contains(err.Error(), "HTTPS") {
		t.Errorf("configure() error = %q, want mention of HTTPS", err.Error())
	}
}

func TestConfigureRequiresSubscriptionKey(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "https://gw.example.com"
	c.SubscriptionKey.Value = ""

	err := c.configure()
	if err == nil {
		t.Error("configure() expected error for empty subscription key")
	}
}

func TestConfigureDefaultsToBedrockBackend(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "https://gw.example.com"
	c.SubscriptionKey.Value = "key"
	c.BackendType.Value = ""

	err := c.configure()
	if err != nil {
		t.Fatalf("configure() error = %v", err)
	}
	if c.BackendType.Value != "bedrock" {
		t.Errorf("BackendType = %q, want %q", c.BackendType.Value, "bedrock")
	}
}

func TestConfigureAllBackendTypes(t *testing.T) {
	tests := []struct {
		name        string
		backendType string
	}{
		{"bedrock", "bedrock"},
		{"azure-openai", "azure-openai"},
		{"vertex-ai", "vertex-ai"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			c.GatewayURL.Value = "https://gw.example.com"
			c.SubscriptionKey.Value = "key"
			c.BackendType.Value = tt.backendType

			err := c.configure()
			if err != nil {
				t.Fatalf("configure(%s) error = %v", tt.backendType, err)
			}
			if c.backend == nil {
				t.Errorf("backend not initialized for type %q", tt.backendType)
			}
		})
	}
}

func TestConfigureInvalidBackend(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "https://gw.example.com"
	c.SubscriptionKey.Value = "key"
	c.BackendType.Value = "unsupported"

	err := c.configure()
	if err == nil {
		t.Error("configure() expected error for unsupported backend")
	}
}

func TestListModelsWithoutInit(t *testing.T) {
	c := NewClient()
	_, err := c.ListModels(context.Background())
	if err == nil {
		t.Error("ListModels() expected error when backend not initialized")
	}
}

func TestIsConfigured(t *testing.T) {
	c := NewClient()
	if c.IsConfigured() {
		t.Error("IsConfigured() = true for unconfigured client")
	}

	c.GatewayURL.Value = "https://gw.example.com"
	if c.IsConfigured() {
		t.Error("IsConfigured() = true with only gateway URL")
	}

	c.SubscriptionKey.Value = "key"
	if !c.IsConfigured() {
		t.Error("IsConfigured() = false for fully configured client")
	}
}

func TestNeedsRawMode(t *testing.T) {
	c := NewClient()
	if c.NeedsRawMode("any-model") {
		t.Error("NeedsRawMode() should always return false")
	}
}

func TestSendBedrockIntegration(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("wrong auth header: %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("wrong content type: %s", r.Header.Get("Content-Type"))
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]any
		json.Unmarshal(body, &req)

		if _, ok := req["system"]; !ok {
			t.Error("expected 'system' field in request")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"content": []map[string]any{
				{"type": "text", "text": "Response from Bedrock"},
			},
		})
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "Be helpful."},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "us.anthropic.claude-3-haiku-20240307-v1:0",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	result, err := c.Send(context.Background(), msgs, opts)
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	if result != "Response from Bedrock" {
		t.Errorf("Send() = %q, want %q", result, "Response from Bedrock")
	}
}

func TestSendErrorTruncation(t *testing.T) {
	longBody := strings.Repeat("x", 600)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(longBody))
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error for 500 response")
	}
	// The error holds a 26-byte prefix, 500 body bytes, and "...". The 600 limit leaves margin.
	if len(err.Error()) > 600 {
		t.Errorf("error message too long (%d chars), should be truncated", len(err.Error()))
	}
	if strings.Count(err.Error(), "x") > 500 {
		t.Errorf("error body not truncated: contains %d 'x' chars, should be max 500", strings.Count(err.Error(), "x"))
	}
}

func TestSendAuthenticationError(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid subscription key"}`))
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "invalid-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("invalid-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error for 401 response")
	}
	if !strings.Contains(err.Error(), "401") {
		t.Errorf("error should mention 401 status: %v", err)
	}
}

func TestSendModelNotFoundError(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Model not found"}`))
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "non-existent-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error for 404 response")
	}
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("error should mention 404 status: %v", err)
	}
}

// failingRoundTripper returns an error on every request to simulate a network failure.
type failingRoundTripper struct{}

func (f *failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("connection refused")
}

func TestSendNetworkError(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "https://gateway.example.com"
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = &http.Client{Transport: &failingRoundTripper{}}
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error for network failure")
	}
	if !strings.Contains(err.Error(), "HTTP request failed") {
		t.Errorf("error should mention HTTP request failure: %v", err)
	}
}

func TestSendMalformedResponseJSON(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error for malformed JSON")
	}
}

func TestSendWithoutBackendInit(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "https://gw.example.com"
	c.SubscriptionKey.Value = "test-key"

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send() expected error when backend not initialized")
	}
	if !strings.Contains(err.Error(), "backend not initialized") {
		t.Errorf("error should mention backend not initialized: %v", err)
	}
}

func TestSendStreamWithoutBackendInit(t *testing.T) {
	c := NewClient()

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	channel := make(chan domain.StreamUpdate, 1)
	err := c.SendStream(context.Background(), msgs, opts, channel)
	if err == nil {
		t.Fatal("SendStream() expected error when backend not initialized")
	}
	if !strings.Contains(err.Error(), "backend not initialized") {
		t.Errorf("error should mention backend not initialized: %v", err)
	}
}

func TestConfigureInvalidURL(t *testing.T) {
	c := NewClient()
	c.GatewayURL.Value = "://invalid-url"
	c.SubscriptionKey.Value = "test-key"

	err := c.configure()
	if err == nil {
		t.Fatal("configure() expected error for malformed URL")
	}
	if !strings.Contains(err.Error(), "invalid gateway URL") {
		t.Errorf("error should mention invalid URL: %v", err)
	}
}

func TestSendStreamFallback(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"content": []map[string]any{
				{"type": "text", "text": "Streaming response"},
			},
		})
	}))
	defer server.Close()

	c := NewClient()
	c.GatewayURL.Value = server.URL
	c.SubscriptionKey.Value = "test-key"
	c.BackendType.Value = "bedrock"
	c.httpClient = server.Client()
	c.backend = NewBedrockBackend("test-key")

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:       "test-model",
		Temperature: domain.DefaultTemperature,
		TopP:        domain.DefaultTopP,
	}

	channel := make(chan domain.StreamUpdate, 10)
	err := c.SendStream(context.Background(), msgs, opts, channel)
	if err != nil {
		t.Fatalf("SendStream() error = %v", err)
	}

	// The range ends only when SendStream closes the channel.
	updates := []domain.StreamUpdate{}
	for update := range channel {
		updates = append(updates, update)
	}

	if len(updates) != 1 {
		t.Fatalf("expected 1 stream update, got %d", len(updates))
	}
	if updates[0].Content != "Streaming response" {
		t.Errorf("unexpected content: %q", updates[0].Content)
	}
}

func TestAzureOpenAIAPIVersionCompatibility(t *testing.T) {
	// Reference: https://learn.microsoft.com/en-us/azure/foundry/openai/api-version-lifecycle
	// If you change the default API version, make sure that the APIM gateways accept the new version.

	b := NewAzureOpenAIBackend("key", "")
	endpoint := b.BuildEndpoint("https://gw.example.com", "gpt-4")

	if !strings.Contains(endpoint, "api-version=") {
		t.Error("endpoint should include api-version parameter")
	}

	if !strings.Contains(endpoint, "2025-04-01-preview") {
		t.Errorf("Default API version should be 2025-04-01-preview. Got: %s", endpoint)
	}
}

func TestAzureOpenAICustomAPIVersion(t *testing.T) {
	customVersion := "2024-08-01-preview"
	b := NewAzureOpenAIBackend("key", customVersion)
	endpoint := b.BuildEndpoint("https://gw.example.com", "gpt-4")

	if !strings.Contains(endpoint, "api-version="+customVersion) {
		t.Errorf("Custom API version not used. Expected %s in: %s", customVersion, endpoint)
	}
}

func TestAzureOpenAIBackwardCompatibility(t *testing.T) {
	b := NewAzureOpenAIBackend("key", "")
	endpoint := b.BuildEndpoint("https://gw.example.com", "gpt-4")

	if !strings.Contains(endpoint, "2025-04-01-preview") {
		t.Errorf("Empty API version should default to 2025-04-01-preview. Got: %s", endpoint)
	}
}

func TestBedrockTemperatureTopPMutualExclusivity(t *testing.T) {
	b := NewBedrockBackend("key")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}

	opts1 := &domain.ChatOptions{
		Temperature: 0.8,
		TopP:        domain.DefaultTopP,
	}
	bodyBytes1, err := b.PrepareRequest(msgs, opts1)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}
	var body1 map[string]any
	json.Unmarshal(bodyBytes1, &body1)

	if _, ok := body1["temperature"]; !ok {
		t.Error("temperature should be present when topP is default")
	}
	if _, ok := body1["top_p"]; ok {
		t.Error("top_p should not be present when using default value")
	}

	opts2 := &domain.ChatOptions{
		Temperature: 0.8,
		TopP:        0.95, // not the default 0.9
	}
	bodyBytes2, err := b.PrepareRequest(msgs, opts2)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}
	var body2 map[string]any
	json.Unmarshal(bodyBytes2, &body2)

	if _, ok := body2["top_p"]; !ok {
		t.Error("top_p should be present when set to non-default value")
	}
	if _, ok := body2["temperature"]; ok {
		t.Error("temperature should not be present when topP is non-default (mutual exclusivity)")
	}
}
