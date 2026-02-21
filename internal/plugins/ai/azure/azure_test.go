package azure

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai/azurecommon"
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

	models, err := client.ListModels()
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

func TestMiddlewareResponsesRoute(t *testing.T) {
	// Verify /responses is in the deployment routes by testing the middleware
	body := `{"model": "gpt-5"}`
	req, err := http.NewRequest("POST", "https://example.com/openai/responses", io.NopCloser(bytes.NewReader([]byte(body))))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var capturedPath string
	mockNext := func(req *http.Request) (*http.Response, error) {
		capturedPath = req.URL.Path
		return &http.Response{StatusCode: 200}, nil
	}

	_, err = azurecommon.AzureDeploymentMiddleware(req, mockNext)
	if err != nil {
		t.Fatalf("Middleware returned error: %v", err)
	}

	expected := "/openai/deployments/gpt-5/responses"
	if capturedPath != expected {
		t.Errorf("Expected path %q, got %q", expected, capturedPath)
	}
}

func TestMiddlewareChatCompletionsRoute(t *testing.T) {
	body := `{"model": "gpt-4o"}`
	req, err := http.NewRequest("POST", "https://example.com/openai/chat/completions", io.NopCloser(bytes.NewReader([]byte(body))))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var capturedPath string
	mockNext := func(req *http.Request) (*http.Response, error) {
		capturedPath = req.URL.Path
		return &http.Response{StatusCode: 200}, nil
	}

	_, err = azurecommon.AzureDeploymentMiddleware(req, mockNext)
	if err != nil {
		t.Fatalf("Middleware returned error: %v", err)
	}

	expected := "/openai/deployments/gpt-4o/chat/completions"
	if capturedPath != expected {
		t.Errorf("Expected path %q, got %q", expected, capturedPath)
	}
}
