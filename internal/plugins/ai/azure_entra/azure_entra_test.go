package azure_entra

import (
	"testing"
)

func TestNewClientInitialization(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatalf("Expected non-nil client, got nil")
	}
	if client.ApiBaseURL == nil {
		t.Errorf("Expected ApiBaseURL to be initialized, got nil")
	}
	if client.ApiDeployments == nil {
		t.Errorf("Expected ApiDeployments to be initialized, got nil")
	}
	if client.ApiVersion == nil {
		t.Errorf("Expected ApiVersion to be initialized, got nil")
	}
	if client.Client == nil {
		t.Errorf("Expected embedded Client to be initialized, got nil")
	}
}

func TestPluginName(t *testing.T) {
	client := NewClient()
	if name := client.GetName(); name != "AzureEntra" {
		t.Errorf("Expected plugin name %q, got %q", "AzureEntra", name)
	}
}

func TestEnvPrefix(t *testing.T) {
	client := NewClient()
	// Setup questions should use the AZUREENTRA_ prefix.
	// The base URL question key is "API Base URL" under vendor "AzureEntra",
	// which results in env var AZUREENTRA_API_BASE_URL.
	if client.ApiBaseURL == nil {
		t.Fatal("ApiBaseURL setup question is nil")
	}
	if client.ApiDeployments == nil {
		t.Fatal("ApiDeployments setup question is nil")
	}
}

func TestListModels(t *testing.T) {
	client := NewClient()
	client.apiDeployments = []string{"gpt-4o", "gpt-5"}

	models, err := client.ListModels()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"gpt-4o", "gpt-5"}
	if len(models) != len(expected) {
		t.Fatalf("Expected %d models, got %d", len(expected), len(models))
	}
	for i, m := range expected {
		if models[i] != m {
			t.Errorf("Expected model %q, got %q", m, models[i])
		}
	}
}

func TestConfigureMissingBaseURL(t *testing.T) {
	client := NewClient()
	client.ApiDeployments.Value = "gpt-4o"
	client.ApiBaseURL.Value = ""

	err := client.configure()
	if err == nil {
		t.Fatal("Expected error for missing base URL, got nil")
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
