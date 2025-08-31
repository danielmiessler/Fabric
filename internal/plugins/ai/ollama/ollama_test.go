package ollama

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

func TestHandleSchema(t *testing.T) {
	client := NewClient()

	// Test case 1: SchemaContent is not empty (should return nil now)
	optsWithSchema := &domain.ChatOptions{
		SchemaContent: "some schema content",
	}
	err := client.HandleSchema(optsWithSchema)
	if err != nil {
		t.Errorf("Expected no error when SchemaContent is not empty, but got %v", err)
	}

	// Test case 2: SchemaContent is empty (should return nil)
	optsWithoutSchema := &domain.ChatOptions{
		SchemaContent: "",
	}
	err = client.HandleSchema(optsWithoutSchema)
	if err != nil {
		t.Errorf("Expected no error when SchemaContent is empty, but got %v", err)
	}
}

func TestCreateChatRequestWithSchema(t *testing.T) {
	client := NewClient()
	testSchema := `{"type": "object", "properties": {"name": {"type": "string"}}}`
	opts := &domain.ChatOptions{
		Model:         "test-model",
		SchemaContent: testSchema,
	}
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Hello"},
	}

	req := client.createChatRequest(msgs, opts)

	if string(req.Format) != testSchema {
		t.Errorf("Expected Format to be %q, got %q", testSchema, string(req.Format))
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("NewClient returned nil")
	}

	if client.PluginBase == nil {
		t.Error("PluginBase is nil")
	}

	if client.PluginBase.Name != "Ollama" {
		t.Errorf("Expected PluginBase.Name to be 'Ollama', got %q", client.PluginBase.Name)
	}

	if client.ApiUrl == nil {
		t.Error("ApiUrl is nil")
	} else if client.ApiUrl.Value != DefaultBaseUrl {
		t.Errorf("Expected ApiUrl.Value to be %q, got %q", DefaultBaseUrl, client.ApiUrl.Value)
	}

	if client.ApiKey == nil {
		t.Error("ApiKey is nil")
	} else if client.ApiKey.Value != "" {
		t.Errorf("Expected ApiKey.Value to be empty, got %q", client.ApiKey.Value)
	}

	if client.ApiHttpTimeout == nil {
		t.Error("ApiHttpTimeout is nil")
	} else if client.ApiHttpTimeout.Value != "20m" {
		t.Errorf("Expected ApiHttpTimeout.Value to be '20m', got %q", client.ApiHttpTimeout.Value)
	}
}
