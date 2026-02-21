package azurecommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/openai/openai-go/option"
)

// DefaultAPIVersion is the default Azure OpenAI API version.
const DefaultAPIVersion = "2025-04-01-preview"

// deploymentRoutes defines the API paths that require deployment name injection.
var deploymentRoutes = map[string]bool{
	"/chat/completions":     true,
	"/completions":          true,
	"/embeddings":           true,
	"/audio/speech":         true,
	"/audio/transcriptions": true,
	"/audio/translations":   true,
	"/images/generations":   true,
	"/responses":            true,
}

// ParseDeployments splits a comma-separated deployment string into a slice,
// trimming whitespace and discarding empty entries.
func ParseDeployments(value string) []string {
	parts := strings.Split(value, ",")
	var deployments []string
	for _, part := range parts {
		if deployment := strings.TrimSpace(part); deployment != "" {
			deployments = append(deployments, deployment)
		}
	}
	return deployments
}

// BuildEndpoint appends the /openai/ suffix to the given base URL.
func BuildEndpoint(baseURL string) string {
	return strings.TrimSuffix(baseURL, "/") + "/openai/"
}

// AzureDeploymentMiddleware transforms Azure OpenAI API paths to include
// the deployment name. Azure requires URLs like:
// /openai/deployments/{deployment-name}/chat/completions
// but the SDK sends paths like /chat/completions
func AzureDeploymentMiddleware(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {

	path := req.URL.Path

	// Remove /openai prefix if present (SDK may add it via base URL)
	trimmedPath := strings.TrimPrefix(path, "/openai")
	if !strings.HasPrefix(trimmedPath, "/") {
		trimmedPath = "/" + trimmedPath
	}

	if deploymentRoutes[trimmedPath] {
		deploymentName, err := ExtractDeploymentFromBody(req)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", i18n.T("azure_failed_extract_deployment"), err)
		}

		newPath := "/openai/deployments/" + url.PathEscape(deploymentName) + trimmedPath
		req.URL.Path = newPath
		req.URL.RawPath = "" // Clear RawPath to ensure Path is used
	}

	return next(req)
}

// ExtractDeploymentFromBody reads the model field from the JSON request body
// and restores the body for subsequent use.
func ExtractDeploymentFromBody(req *http.Request) (string, error) {
	if req.Body == nil {
		return "", fmt.Errorf("%s", i18n.T("azure_request_body_nil"))
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	// Restore body for subsequent reads
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	var payload struct {
		Model string `json:"model"`
	}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return "", err
	}

	if payload.Model == "" {
		return "", fmt.Errorf("%s", i18n.T("azure_model_field_empty"))
	}

	return payload.Model, nil
}
