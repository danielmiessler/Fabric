package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

// modelResponse represents a minimal model returned by the API.
// This mirrors the shape used by OpenAI-compatible providers that return
// either an array of models or an object with a `data` field.
type modelResponse struct {
	ID string `json:"id"`
}

// errorResponseLimit defines the maximum length of error response bodies for truncation.
const errorResponseLimit = 1024

// maxResponseSize defines the maximum size of response bodies to prevent memory exhaustion.
const maxResponseSize = 10 * 1024 * 1024 // 10MB

// FetchModelsDirectly is used to fetch models directly from the API when the
// standard OpenAI SDK method fails due to a nonstandard format. This is useful
// for providers that return a direct array of models (e.g., GitHub Models) or
// other OpenAI-compatible implementations.
func FetchModelsDirectly(ctx context.Context, baseURL, apiKey, providerName string) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if baseURL == "" {
		return nil, fmt.Errorf(i18n.T("openai_api_base_url_not_configured"), providerName)
	}

	// Build the /models endpoint URL
	fullURL, err := url.JoinPath(baseURL, "models")
	if err != nil {
		return nil, fmt.Errorf(i18n.T("openai_failed_to_create_models_url"), err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Accept", "application/json")

	// TODO: Consider reusing a single http.Client instance (e.g., as a field on Client) instead of allocating a new one for
	// each request.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body for debugging, but limit the number of bytes read
		bodyBytes, readErr := io.ReadAll(io.LimitReader(resp.Body, errorResponseLimit))
		if readErr != nil {
			return nil, fmt.Errorf(i18n.T("openai_unexpected_status_code_read_error"),
				resp.StatusCode, providerName, readErr)
		}
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf(i18n.T("openai_unexpected_status_code_with_body"),
			resp.StatusCode, providerName, bodyString)
	}

	// Read the response body once, with a size limit to prevent memory exhaustion
	// Read up to maxResponseSize + 1 bytes to detect truncation
	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseSize+1))
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > maxResponseSize {
		return nil, fmt.Errorf(i18n.T("openai_models_response_too_large"), providerName, maxResponseSize)
	}

	// Try to parse as an object with data field (OpenAI format)
	var openAIFormat struct {
		Data []modelResponse `json:"data"`
	}
	// Try to parse as a direct array
	var directArray []modelResponse

	if err := json.Unmarshal(bodyBytes, &openAIFormat); err == nil {
		debuglog.Debug(debuglog.Detailed, "Successfully parsed models response from %s using OpenAI format (found %d models)\n", providerName, len(openAIFormat.Data))
		return extractModelIDs(openAIFormat.Data), nil
	}

	if err := json.Unmarshal(bodyBytes, &directArray); err == nil {
		debuglog.Debug(debuglog.Detailed, "Successfully parsed models response from %s using direct array format (found %d models)\n", providerName, len(directArray))
		return extractModelIDs(directArray), nil
	}

	var truncatedBody string
	if len(bodyBytes) > errorResponseLimit {
		truncatedBody = string(bodyBytes[:errorResponseLimit]) + "..."
	} else {
		truncatedBody = string(bodyBytes)
	}
	return nil, fmt.Errorf(i18n.T("openai_unable_to_parse_models_response"), truncatedBody)
}

func extractModelIDs(models []modelResponse) []string {
	modelIDs := make([]string, 0, len(models))
	for _, model := range models {
		modelIDs = append(modelIDs, model.ID)
	}
	return modelIDs
}
