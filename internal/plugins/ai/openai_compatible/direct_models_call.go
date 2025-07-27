package openai_compatible

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Model represents a model returned by the API
type Model struct {
	ID string `json:"id"`
}

// ErrorResponseLimit defines the maximum length of error response bodies for truncation.
const errorResponseLimit = 1024 // Limit for error response body size

// DirectlyGetModels is used to fetch models directly from the API
// when the standard OpenAI SDK method fails due to a nonstandard format.
// This is useful for providers like Together that return a direct array of models.
func (c *Client) DirectlyGetModels(ctx context.Context) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	baseURL := c.ApiBaseURL.Value
	if baseURL == "" {
		return nil, fmt.Errorf("API base URL not configured for provider %s", c.GetName())
	}

	// Build the /models endpoint URL
	fullURL, err := url.JoinPath(baseURL, "models")
	if err != nil {
		return nil, fmt.Errorf("failed to create models URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey.Value))
	req.Header.Set("Accept", "application/json")

	// TODO: Consider reusing a single http.Client instance (e.g., as a field on Client) instead of allocating a new one for each request.

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body for debugging
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		if len(bodyString) > errorResponseLimit { // Truncate if too large
			bodyString = bodyString[:errorResponseLimit] + "..."
		}
		return nil, fmt.Errorf("unexpected status code: %d from provider %s, response body: %s",
			resp.StatusCode, c.GetName(), bodyString)
	}

	// Read the response body once
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try to parse as an object with data field (OpenAI format)
	var openAIFormat struct {
		Data []Model `json:"data"`
	}
	// Try to parse as a direct array (Together format)
	var directArray []Model

	if err := json.Unmarshal(bodyBytes, &openAIFormat); err == nil && len(openAIFormat.Data) > 0 {
		return extractModelIDs(openAIFormat.Data), nil
	}

	if err := json.Unmarshal(bodyBytes, &directArray); err == nil && len(directArray) > 0 {
		return extractModelIDs(directArray), nil
	}

	var truncatedBody string
	if len(bodyBytes) > errorResponseLimit {
		truncatedBody = string(bodyBytes[:errorResponseLimit]) + "..."
	} else {
		truncatedBody = string(bodyBytes)
	}
	return nil, fmt.Errorf("unable to parse models response; raw response: %s", truncatedBody)
}

func extractModelIDs(models []Model) []string {
	modelIDs := make([]string, 0, len(models))
	for _, model := range models {
		modelIDs = append(modelIDs, model.ID)
	}
	return modelIDs
}
