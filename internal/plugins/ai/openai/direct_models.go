package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

// modelResponse holds the only field this code reads from a model entry.
type modelResponse struct {
	ID string `json:"id"`
}

// errorResponseLimit caps the response bytes that go into an error message.
const errorResponseLimit = 1024

// maxResponseSize caps a models response body. A larger body is an error.
const maxResponseSize = 10 * 1024 * 1024

// FetchModelsDirectly is used to fetch models directly from the API when the
// standard OpenAI SDK method fails due to a nonstandard format. This is useful
// for providers that return a direct array of models instead of the OpenAI
// object format.
// If httpClient is nil, a new client with default settings will be created.
func FetchModelsDirectly(ctx context.Context, baseURL, apiKey, providerName string, httpClient *http.Client) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if baseURL == "" {
		return nil, fmt.Errorf(i18n.T("openai_api_base_url_not_configured"), providerName)
	}

	fullURL, err := url.JoinPath(baseURL, "models")
	if err != nil {
		return nil, fmt.Errorf(i18n.T("openai_failed_to_create_models_url"), err)
	}

	// Serve a fresh cached list first. Some discovery endpoints rate-limit requests.
	if models, ok := readModelsCache(providerName, fullURL, modelsCacheTTL); ok {
		debuglog.Debug(debuglog.Detailed, "Using cached models list for %s (%d models)\n", providerName, len(models))
		return models, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Accept", "application/json")

	client := httpClient
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		// On a network error, fall back to a cached list of any age.
		if models, ok := readModelsCache(providerName, fullURL, 0); ok {
			debuglog.Debug(debuglog.Basic, "Fetch failed for %s (%v); serving stale cached models\n", providerName, err)
			return models, nil
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// On a bad status, usually HTTP 429, fall back to a cached list of any age.
		if models, ok := readModelsCache(providerName, fullURL, 0); ok {
			debuglog.Debug(debuglog.Basic, "Status %d from %s; serving stale cached models\n", resp.StatusCode, providerName)
			return models, nil
		}

		bodyBytes, readErr := io.ReadAll(io.LimitReader(resp.Body, errorResponseLimit))
		if readErr != nil {
			return nil, fmt.Errorf(i18n.T("openai_unexpected_status_code_read_error"),
				resp.StatusCode, providerName, readErr)
		}

		// A rate-limit response often has an HTML body. Return a short message, not the page.
		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter := strings.TrimSpace(resp.Header.Get("Retry-After"))
			if retryAfter == "" {
				retryAfter = "60"
			}
			return nil, fmt.Errorf(i18n.T("openai_models_rate_limited"), providerName, retryAfter)
		}

		bodyString := string(bodyBytes)
		return nil, fmt.Errorf(i18n.T("openai_unexpected_status_code_with_body"),
			resp.StatusCode, providerName, bodyString)
	}

	// Read one byte more than maxResponseSize to detect an oversized body.
	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseSize+1))
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > maxResponseSize {
		return nil, fmt.Errorf(i18n.T("openai_models_response_too_large"), providerName, maxResponseSize)
	}

	var openAIFormat struct {
		Data []modelResponse `json:"data"`
	}
	var directArray []modelResponse

	if err := json.Unmarshal(bodyBytes, &openAIFormat); err == nil {
		debuglog.Debug(debuglog.Detailed, "Successfully parsed models response from %s using OpenAI format (found %d models)\n", providerName, len(openAIFormat.Data))
		ids := extractModelIDs(openAIFormat.Data)
		_ = writeModelsCache(providerName, fullURL, ids)
		return ids, nil
	}

	if err := json.Unmarshal(bodyBytes, &directArray); err == nil {
		debuglog.Debug(debuglog.Detailed, "Successfully parsed models response from %s using direct array format (found %d models)\n", providerName, len(directArray))
		ids := extractModelIDs(directArray)
		_ = writeModelsCache(providerName, fullURL, ids)
		return ids, nil
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
