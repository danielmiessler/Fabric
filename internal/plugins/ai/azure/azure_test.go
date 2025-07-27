package azure

import (
	"testing"
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
	client.ApiVersion.Value = "2021-01-01"

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

	if client.ApiVersion.Value != "2021-01-01" {
		t.Errorf("Expected API version to be '2021-01-01', got %s", client.ApiVersion.Value)
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
