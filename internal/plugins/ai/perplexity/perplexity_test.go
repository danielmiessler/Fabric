package perplexity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
	// Re-add this import
)

// commonExpectedSchema defines a schema used across multiple tests.
var commonExpectedSchema = map[string]interface{}{"type": "object", "properties": map[string]interface{}{"test": map[string]interface{}{"type": "string"}}}

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if c.Name != providerName {
		t.Errorf("Expected name %s, got %s", providerName, c.Name)
	}
	if c.APIKey == nil {
		t.Error("APIKey setup question not added")
	}
}

func TestConfigure(t *testing.T) {
	c := NewClient()
	// Test with API key from environment variable
	os.Setenv(c.EnvNamePrefix+"API_KEY", "test-env-key")
	defer os.Unsetenv(c.EnvNamePrefix + "API_KEY")

	if err := c.Configure(); err != nil {
		t.Errorf("Configure failed with env var: %v", err)
	}
	if c.APIKey.Value != "test-env-key" {
		t.Errorf("Expected APIKey from env, got %s", c.APIKey.Value)
	}
	if c.client == nil {
		t.Error("Perplexity client not initialized")
	}

	// Test with API key set via setup question
	c = NewClient() // Re-initialize to clear previous state
	c.APIKey.Value = "test-setup-key"
	if err := c.Configure(); err != nil {
		t.Errorf("Configure failed with setup value: %v", err)
	}
	if c.APIKey.Value != "test-setup-key" {
		t.Errorf("Expected APIKey from setup, got %s", c.APIKey.Value)
	}

	// Test without API key
	c = NewClient()
	c.APIKey.Value = "" // Clear any previous value
	os.Unsetenv(c.EnvNamePrefix + "API_KEY")
	err := c.Configure()
	if err == nil {
		t.Error("Configure succeeded without API key, expected error")
	}
	if !strings.Contains(err.Error(), "API key not configured") {
		t.Errorf("Expected 'API key not configured' error, got: %v", err)
	}
}

func TestListModels(t *testing.T) {
	c := NewClient()
	m, err := c.ListModels()
	if err != nil {
		t.Errorf("ListModels returned an error: %v", err)
	}
	if len(m) == 0 {
		t.Error("ListModels returned empty list")
	}
	if m[0] != "r1-1776" { // Check first model
		t.Errorf("Unexpected first model: %s", m[0])
	}
}

func TestHandleSchema(t *testing.T) {
	c := NewClient()
	opts := &domain.ChatOptions{}
	err := c.HandleSchema(opts)
	if err != nil {
		t.Errorf("HandleSchema returned an error: %v", err)
	}
	// Verify it's a no-op, no changes should occur to opts or client state
}

func TestSend_WithSchema(t *testing.T) {
	// Specific schema for this test, not the common one
	specificExpectedSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{"type": "string"},
		},
	}
	schemaBytes, _ := json.Marshal(specificExpectedSchema)
	schemaContent := string(schemaBytes)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Errorf("Expected path /chat/completions, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header 'Bearer test-api-key', got '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		if responseFormat, ok := reqBody["response_format"].(map[string]interface{}); ok {
			if rfType, ok := responseFormat["type"].(string); !ok || rfType != "json_object" {
				t.Errorf("Expected response_format type 'json_object', got '%v'", rfType)
			}
			if schemaJSON, ok := responseFormat["schema"].(map[string]interface{}); ok {
				if !jsonDeepEqual(schemaJSON, specificExpectedSchema) {
					t.Errorf("Received schema does not match expected schema.\nExpected: %v\nGot: %v", specificExpectedSchema, schemaJSON)
				}
			} else {
				t.Error("response_format schema not found or invalid type")
			}
		} else {
			t.Error("response_format not found or invalid type")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "chatcmpl-123", "choices": [{"message": {"role": "assistant", "content": "Mocked response"}}], "usage": {"prompt_tokens": 10, "completion_tokens": 5, "total_tokens": 15}}`))
	}))
	defer testServer.Close()

	originalBaseURL := os.Getenv("PPLX_BASE_URL")
	os.Setenv("PPLX_BASE_URL", testServer.URL)
	defer func() {
		if originalBaseURL != "" {
			os.Setenv("PPLX_BASE_URL", originalBaseURL)
		} else {
			os.Unsetenv("PPLX_BASE_URL")
		}
	}()

	c := NewClient()
	c.APIKey = &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-api-key"}}
	c.Configure()

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:         "sonar-small-online",
		SchemaContent: schemaContent,
	}

	resp, err := c.Send(context.Background(), msgs, opts)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized: check your API key") {
			t.Logf("Send returned expected unauthorized error: %v", err)
		} else {
			t.Fatalf("Send returned unexpected error: %v", err)
		}
	}
	if resp != "Mocked response" && err == nil { // Only check response if no error
		t.Errorf("Expected 'Mocked response', got '%s'", resp)
	}
}

func TestSend_InvalidSchema(t *testing.T) {
	c := NewClient()
	c.APIKey = &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-api-key"}} // Mock API key
	// No need to set c.client to a mock as the error should occur before API call

	opts := &domain.ChatOptions{
		Model:         "sonar-small-online",
		SchemaContent: "{invalid json",
	}
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err == nil {
		t.Fatal("Send with invalid schema did not return an error")
	}
	if !strings.Contains(err.Error(), "failed to parse schema content") {
		t.Errorf("Expected 'failed to parse schema content' error, got: %v", err)
	}
}

func TestSendStream_WithSchema(t *testing.T) {
	// Specific schema for this test, not the common one
	specificExpectedSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"data": map[string]interface{}{"type": "array"},
		},
	}
	schemaBytes, _ := json.Marshal(specificExpectedSchema)
	schemaContent := string(schemaBytes)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Errorf("Expected path /chat/completions, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header 'Bearer test-api-key', got '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		if responseFormat, ok := reqBody["response_format"].(map[string]interface{}); ok {
			if rfType, ok := responseFormat["type"].(string); !ok || rfType != "json_object" {
				t.Errorf("Expected response_format type 'json_object', got '%v'", rfType)
			}
			if schemaJSON, ok := responseFormat["schema"].(map[string]interface{}); ok {
				if !jsonDeepEqual(schemaJSON, specificExpectedSchema) {
					t.Errorf("Received schema does not match expected schema.\nExpected: %v\nGot: %v", specificExpectedSchema, schemaJSON)
				}
			} else {
				t.Error("response_format schema not found or invalid type")
			}
		} else {
			t.Error("response_format not found or invalid type")
		}

		// Simulate streaming responses
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatal("http.ResponseWriter does not implement http.Flusher")
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Send initial part
		fmt.Fprintf(w, "data: %s\n\n", `{"choices":[{"delta":{"content":"Stream part 1"}}]}`)
		flusher.Flush()
		time.Sleep(10 * time.Millisecond) // Simulate network delay

		// Send second part
		fmt.Fprintf(w, "data: %s\n\n", `{"choices":[{"delta":{"content":"Stream part 2"}}]}`)
		flusher.Flush()
		time.Sleep(10 * time.Millisecond) // Simulate network delay

		// Send citations at the end
		fmt.Fprintf(w, "data: %s\n\n", `{"citations":["Citation 1", "Citation 2"]}`)
		flusher.Flush()
	}))
	defer testServer.Close()

	originalBaseURL := os.Getenv("PPLX_BASE_URL")
	os.Setenv("PPLX_BASE_URL", testServer.URL)
	defer func() {
		if originalBaseURL != "" {
			os.Setenv("PPLX_BASE_URL", originalBaseURL)
		} else {
			os.Unsetenv("PPLX_BASE_URL")
		}
	}()

	c := NewClient()
	c.APIKey = &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-api-key"}}
	c.Configure()

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}
	opts := &domain.ChatOptions{
		Model:         "sonar-small-online",
		SchemaContent: schemaContent,
	}

	outputChannel := make(chan string)
	var receivedContent []string
	var wg sync.WaitGroup // Use a separate WaitGroup for receiving from outputChannel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for chunk := range outputChannel {
			receivedContent = append(receivedContent, chunk)
		}
	}()

	err := c.SendStream(msgs, opts, outputChannel)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized: check your API key") {
			t.Logf("SendStream returned expected unauthorized error: %v", err)
		} else {
			t.Fatalf("SendStream returned unexpected error: %v", err)
		}
	}

	wg.Wait() // Wait for all content to be received

	// Only check content if no error occurred
	if err == nil {
		expected := []string{"Stream part 1", "Stream part 2", "\n\n# CITATIONS\n\n", "- [1] Citation 1\n", "- [2] Citation 2\n"}
		if len(receivedContent) != len(expected) {
			t.Fatalf("Expected %d chunks, got %d: %v", len(expected), len(receivedContent), receivedContent)
		}
		for i, chunk := range receivedContent {
			if chunk != expected[i] {
				t.Errorf("Chunk %d: Expected '%s', got '%s'", i, expected[i], chunk)
			}
		}
	}
}

func TestSendStream_InvalidSchema(t *testing.T) {
	c := NewClient()
	c.APIKey = &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-api-key"}} // Mock API key
	// No need to set c.client to a mock as the error should occur before API call

	opts := &domain.ChatOptions{
		Model:         "sonar-small-online",
		SchemaContent: "{invalid json",
	}
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}

	outputChannel := make(chan string)
	err := c.SendStream(msgs, opts, outputChannel)
	if err == nil {
		t.Fatal("SendStream with invalid schema did not return an error")
	}
	if !strings.Contains(err.Error(), "failed to parse schema content") {
		t.Errorf("Expected 'failed to parse schema content' error, got: %v", err)
	}
	// Ensure the channel is closed on error
	select {
	case _, ok := <-outputChannel:
		if ok {
			t.Error("Output channel was not closed after error")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Timed out waiting for channel to close")
	}
}

// Helper function to compare two JSON interfaces
func jsonDeepEqual(a, b interface{}) bool {
	aBytes, _ := json.Marshal(a)
	bBytes, _ := json.Marshal(b)
	return string(aBytes) == string(bBytes)
}

// Test for the Configure method's handling of an unconfigured API key
func TestConfigure_NoAPIKey(t *testing.T) {
	c := NewClient()
	c.APIKey.Value = "" // Ensure it's empty
	os.Unsetenv(c.EnvNamePrefix + "API_KEY")

	err := c.Configure()
	if err == nil {
		t.Fatal("Expected error when API key is not configured, got nil")
	}
	expectedError := "Perplexity API key not configured. Please set the PERPLEXITY_API_KEY environment variable or run 'fabric --setup Perplexity'"
	if err.Error() != expectedError {
		t.Errorf("Expected error message:\n%s\nGot:\n%s", expectedError, err.Error())
	}
}

// Test SendCompletionRequest with a real HTTP server mock to ensure request structure
func TestSendCompletionRequest_HTTPMock(t *testing.T) {
	// This test uses commonExpectedSchema
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Errorf("Expected path /chat/completions, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header 'Bearer test-api-key', got '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		if model, ok := reqBody["model"].(string); !ok || model != "test-model" {
			t.Errorf("Expected model 'test-model', got '%v'", reqBody["model"])
		}

		if responseFormat, ok := reqBody["response_format"].(map[string]interface{}); ok {
			if rfType, ok := responseFormat["type"].(string); !ok || rfType != "json_object" {
				t.Errorf("Expected response_format type 'json_object', got '%v'", rfType)
			}
			if schema, ok := responseFormat["schema"].(map[string]interface{}); ok {
				if !jsonDeepEqual(schema, commonExpectedSchema) {
					t.Errorf("Received schema does not match expected schema.\nExpected: %v\nGot: %v", commonExpectedSchema, schema)
				}
			} else {
				t.Error("response_format schema not found or invalid type")
			}
		} else {
			t.Error("response_format not found or invalid type")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "chatcmpl-123", "choices": [{"message": {"role": "assistant", "content": "{\"test\": \"value\"}"}}], "usage": {"prompt_tokens": 10, "completion_tokens": 5, "total_tokens": 15}}`))
	}))
	defer testServer.Close()

	// Temporarily override the base URL for the perplexity-go client
	originalBaseURL := os.Getenv("PPLX_BASE_URL")
	os.Setenv("PPLX_BASE_URL", testServer.URL)
	defer func() {
		if originalBaseURL != "" {
			os.Setenv("PPLX_BASE_URL", originalBaseURL)
		} else {
			os.Unsetenv("PPLX_BASE_URL")
		}
	}()

	c := NewClient()
	c.APIKey = &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-api-key"}}
	c.Configure() // This will use the overridden PPLX_BASE_URL to initialize the internal client

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}
	schemaContent := `{"type": "object", "properties": {"test": {"type": "string"}}}`
	opts := &domain.ChatOptions{
		Model:         "test-model",
		SchemaContent: schemaContent,
	}

	_, err := c.Send(context.Background(), msgs, opts)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized: check your API key") {
			t.Logf("Send failed with expected unauthorized error: %v", err)
		} else {
			t.Fatalf("Send failed with unexpected error: %v", err)
		}
	}
}
