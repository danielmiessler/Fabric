package azure

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai/azurecommon"
	openaiapi "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// Test generated using Keploy
func TestNewClientInitialization(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatalf("Expected non-nil client, got nil")
	}
	if client.ApiDeployments == nil {
		t.Errorf("Expected ApiDeployments to be initialized, got nil")
	}
	if client.ApiVersion == nil {
		t.Errorf("Expected ApiVersion to be initialized, got nil")
	}
	if client.Client == nil {
		t.Errorf("Expected Client to be initialized, got nil")
	}
}

// Test generated using Keploy
func TestClientConfigure(t *testing.T) {
	client := NewClient()
	client.ApiDeployments.Value = "deployment1,deployment2"
	client.ApiKey.Value = "test-api-key"
	client.ApiBaseURL.Value = "https://example.com"
	client.ApiVersion.Value = "2025-04-01-preview"

	err := client.configure()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedDeployments := []string{"deployment1", "deployment2"}
	if len(client.apiDeployments) != len(expectedDeployments) {
		t.Errorf("Expected %d deployments, got %d", len(expectedDeployments), len(client.apiDeployments))
	}
	for i, deployment := range expectedDeployments {
		if client.apiDeployments[i] != deployment {
			t.Errorf("Expected deployment %s, got %s", deployment, client.apiDeployments[i])
		}
	}

	if client.ApiClient == nil {
		t.Errorf("Expected ApiClient to be initialized, got nil")
	}

	if client.ApiVersion.Value != "2025-04-01-preview" {
		t.Errorf("Expected API version to be '2025-04-01-preview', got %s", client.ApiVersion.Value)
	}
}

func TestClientConfigureDefaultAPIVersion(t *testing.T) {
	client := NewClient()
	client.ApiDeployments.Value = "deployment1"
	client.ApiKey.Value = "test-api-key"
	client.ApiBaseURL.Value = "https://example.com"

	if err := client.configure(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client.ApiVersion.Value != azurecommon.DefaultAPIVersion {
		t.Errorf("Expected API version to default to %s, got %s", azurecommon.DefaultAPIVersion, client.ApiVersion.Value)
	}
}

// Test generated using Keploy
func TestListModels(t *testing.T) {
	client := NewClient()
	client.apiDeployments = []string{"deployment1", "deployment2"}

	models, err := client.ListModels(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedModels := []string{"deployment1", "deployment2"}
	if len(models) != len(expectedModels) {
		t.Errorf("Expected %d models, got %d", len(expectedModels), len(models))
	}
	for i, model := range expectedModels {
		if models[i] != model {
			t.Errorf("Expected model %s, got %s", model, models[i])
		}
	}
}

func TestNeedsRawModeInheritsFromParent(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		model    string
		expected bool
	}{
		{"o1 model", "o1", true},
		{"o1-preview", "o1-preview", true},
		{"o3-mini", "o3-mini", true},
		{"o4-mini", "o4-mini", true},
		{"gpt-5", "gpt-5", true},
		{"gpt-5-turbo", "gpt-5-turbo", true},
		{"gpt-4o", "gpt-4o", false},
		{"gpt-4", "gpt-4", false},
		{"regular deployment", "my-deployment", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.NeedsRawMode(tt.model)
			if result != tt.expected {
				t.Errorf("NeedsRawMode(%q) = %v, want %v", tt.model, result, tt.expected)
			}
		})
	}
}

func TestConfigureRoutesChatToDeployment(t *testing.T) {
	var path, apiVersion, apiKey string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, apiVersion, apiKey = r.URL.Path, r.URL.Query().Get("api-version"), r.Header.Get("Api-Key")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[]}`)
	}))
	defer server.Close()

	client := NewClient()
	client.ApiDeployments.Value = "gpt-4o"
	client.ApiKey.Value = "test-api-key"
	client.ApiBaseURL.Value = server.URL
	if err := client.configure(); err != nil {
		t.Fatalf("configure() error = %v", err)
	}

	// openai-go v3 sends Azure credentials only over HTTPS. server.Client accepts the test certificate.
	_, err := client.ApiClient.Chat.Completions.New(context.Background(), openaiapi.ChatCompletionNewParams{
		Model:    "gpt-4o",
		Messages: []openaiapi.ChatCompletionMessageParamUnion{openaiapi.UserMessage("Hello")},
	}, option.WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatalf("Chat.Completions.New() error = %v", err)
	}
	if path != "/openai/deployments/gpt-4o/chat/completions" {
		t.Errorf("request path = %q, want /openai/deployments/gpt-4o/chat/completions", path)
	}
	if apiVersion != azurecommon.DefaultAPIVersion {
		t.Errorf("api-version = %q, want %q", apiVersion, azurecommon.DefaultAPIVersion)
	}
	if apiKey != "test-api-key" {
		t.Errorf("Api-Key header = %q, want test-api-key", apiKey)
	}
}
