package azurecommon

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestParseDeployments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"single deployment", "gpt-4o", []string{"gpt-4o"}},
		{"multiple deployments", "gpt-4o,gpt-5", []string{"gpt-4o", "gpt-5"}},
		{"with spaces", " gpt-4o , gpt-5 ", []string{"gpt-4o", "gpt-5"}},
		{"empty string", "", nil},
		{"only commas", ",,", nil},
		{"trailing comma", "gpt-4o,", []string{"gpt-4o"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseDeployments(tt.input)
			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d deployments, got %d", len(tt.expected), len(result))
			}
			for i, d := range tt.expected {
				if result[i] != d {
					t.Errorf("expected deployment %q, got %q", d, result[i])
				}
			}
		})
	}
}

func TestBuildEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"without trailing slash", "https://example.openai.azure.com", "https://example.openai.azure.com/openai/"},
		{"with trailing slash", "https://example.openai.azure.com/", "https://example.openai.azure.com/openai/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildEndpoint(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestAzureDeploymentMiddlewareChatCompletions(t *testing.T) {
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

	_, err = AzureDeploymentMiddleware(req, mockNext)
	if err != nil {
		t.Fatalf("Middleware returned error: %v", err)
	}

	expected := "/openai/deployments/gpt-4o/chat/completions"
	if capturedPath != expected {
		t.Errorf("Expected path %q, got %q", expected, capturedPath)
	}
}

func TestAzureDeploymentMiddlewareResponses(t *testing.T) {
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

	_, err = AzureDeploymentMiddleware(req, mockNext)
	if err != nil {
		t.Fatalf("Middleware returned error: %v", err)
	}

	expected := "/openai/deployments/gpt-5/responses"
	if capturedPath != expected {
		t.Errorf("Expected path %q, got %q", expected, capturedPath)
	}
}

func TestExtractDeploymentFromBody(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		expected  string
		expectErr bool
	}{
		{"valid model", `{"model": "gpt-4o"}`, "gpt-4o", false},
		{"empty model", `{"model": ""}`, "", true},
		{"missing model", `{"temperature": 0.5}`, "", true},
		{"invalid json", `not json`, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "https://example.com", io.NopCloser(bytes.NewReader([]byte(tt.body))))
			result, err := ExtractDeploymentFromBody(req)
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestExtractDeploymentFromBodyNilBody(t *testing.T) {
	req, _ := http.NewRequest("POST", "https://example.com", nil)
	_, err := ExtractDeploymentFromBody(req)
	if err == nil {
		t.Fatal("expected error for nil body, got nil")
	}
}
