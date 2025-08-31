package lmstudio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Mock HTTP Client for testing
type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, fmt.Errorf("DoFunc not set")
}

func TestHandleSchema(t *testing.T) {
	client := NewClient()

	// Test case 1: Valid schema content
	validSchema := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"age": {"type": "integer"}
		},
		"required": ["name"]
	}`
	opts := &domain.ChatOptions{SchemaContent: validSchema}

	err := client.HandleSchema(opts)
	if err != nil {
		t.Fatalf("HandleSchema failed for valid schema: %v", err)
	}

	if client.ParsedSchema == nil {
		t.Fatal("ParsedSchema is nil after handling valid schema")
	}

	parsedType, ok := client.ParsedSchema["type"].(string)
	if !ok || parsedType != "json_schema" {
		t.Errorf("Expected type 'json_schema', got %v", parsedType)
	}

	jsonSchema, ok := client.ParsedSchema["json_schema"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected json_schema map, got nil")
	}

	schemaName, ok := jsonSchema["name"].(string)
	if !ok || schemaName != "structured_output_schema" {
		t.Errorf("Expected schema name 'structured_output_schema', got %v", schemaName)
	}

	strict, ok := jsonSchema["strict"].(bool)
	if !ok || !strict {
		t.Errorf("Expected strict true, got %v", strict)
	}

	schemaContent, ok := jsonSchema["schema"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected schema content map, got nil")
	}
	if _, found := schemaContent["properties"]; !found {
		t.Error("Expected 'properties' in schema content")
	}

	// Debugging: Print the parsed schema to understand the "name" discrepancy
	t.Logf("ParsedSchema: %+v", client.ParsedSchema)

	// Test case 2: Empty schema content
	optsEmpty := &domain.ChatOptions{SchemaContent: ""}
	err = client.HandleSchema(optsEmpty)
	if err != nil {
		t.Fatalf("HandleSchema failed for empty schema: %v", err)
	}
	if client.ParsedSchema != nil {
		t.Error("ParsedSchema should be nil for empty schema content")
	}

	// Test case 3: Invalid schema content (malformed JSON)
	invalidSchema := `{"type": "object", "properties": {"name": "string"` // Missing closing brace
	optsInvalid := &domain.ChatOptions{SchemaContent: invalidSchema}
	err = client.HandleSchema(optsInvalid)
	if err == nil {
		t.Fatal("HandleSchema did not return an error for invalid schema")
	}
	if !strings.Contains(err.Error(), "failed to parse schema content") {
		t.Errorf("Expected error about parsing schema, got %v", err)
	}
}

func TestSendWithSchema(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("Expected to request '/v1/chat/completions', got '%s'", r.URL.Path)
		}
		var reqPayload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqPayload)
		if err != nil {
			t.Fatalf("Failed to decode request payload: %v", err)
		}

		if _, ok := reqPayload["response_format"]; !ok {
			t.Error("Expected 'response_format' in request payload, but not found")
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"choices": [
				{
					"message": {
						"content": "{\"name\": \"test\", \"age\": 30}"
					}
				}
			]
		}`)
	}))
	defer ts.Close()

	client := NewClient()
	client.ApiUrl.Value = ts.URL + "/v1" // Correct the API URL to include /v1
	client.HttpClient = ts.Client()      // Use the test server's client

	validSchema := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"age": {"type": "integer"}
		},
		"required": ["name"]
	}`
	opts := &domain.ChatOptions{
		Model:         "test-model",
		SchemaContent: validSchema,
	}
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Generate a person"},
	}

	// Prepare schema
	err := client.HandleSchema(opts)
	if err != nil {
		t.Fatalf("HandleSchema failed: %v", err)
	}

	content, err := client.Send(context.Background(), msgs, opts)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}

	expectedContent := `{"name": "test", "age": 30}`
	if content != expectedContent {
		t.Errorf("Expected content '%s', got '%s'", expectedContent, content)
	}
}

func TestSendStreamWithSchema(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("Expected to request '/v1/chat/completions', got '%s'", r.URL.Path)
		}
		var reqPayload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqPayload)
		if err != nil {
			t.Fatalf("Failed to decode request payload: %v", err)
		}

		if _, ok := reqPayload["response_format"]; !ok {
			t.Error("Expected 'response_format' in request payload, but not found")
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Simulate streaming response
		// Simulate streaming response with single newlines
		// Simulate streaming response with valid JSON objects per line, containing fragments
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":"{"}}]}`+"\n")
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":"\"name\":"}}]}`+"\n")
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":"\"test\""}}]}`+"\n")
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":", \"age\":"}}]}`+"\n")
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":"30"}}]}`+"\n")
		fmt.Fprint(w, `data: {"choices":[{"delta":{"content":"}"}}]}`+"\n")
		fmt.Fprint(w, "data: [DONE]\n")
	}))
	defer ts.Close()

	client := NewClient()
	client.ApiUrl.Value = ts.URL + "/v1" // Correct the API URL to include /v1
	client.HttpClient = ts.Client()

	validSchema := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"age": {"type": "integer"}
		},
		"required": ["name"]
	}`
	opts := &domain.ChatOptions{
		Model:         "test-model",
		SchemaContent: validSchema,
	}
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Generate a person"},
	}

	// Prepare schema
	err := client.HandleSchema(opts)
	if err != nil {
		t.Fatalf("HandleSchema failed: %v", err)
	}

	channel := make(chan string)
	var receivedContent []string
	done := make(chan struct{})
	go func() {
		for c := range channel {
			t.Logf("Received content chunk: %s", c) // Add logging for debugging
			receivedContent = append(receivedContent, c)
		}
		close(done)
	}()

	err = client.SendStream(msgs, opts, channel)
	if err != nil {
		t.Fatalf("SendStream failed: %v", err)
	}
	<-done // Wait for all content to be received

	// The mock server sends escaped JSON fragments. json.Unmarshal in SendStream will unescape them.
	// So, the expected content is the unescaped, joined string.
	expectedContent := `{"name":"test", "age":30}` // Removed extra spaces to match actual JSON output
	combinedContent := strings.Join(receivedContent, "")
	if combinedContent != expectedContent {
		t.Errorf("Expected combined content '%s', got '%s'", expectedContent, combinedContent)
	}
}
