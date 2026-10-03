package vertexai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

const (
	maxResponseSize    = 10 * 1024 * 1024
	errorResponseLimit = 1024

	// The Model Garden listing needs a regional endpoint. "global" does not work.
	defaultModelGardenRegion = "us-central1"
)

// publishers lists the Model Garden publishers whose models Send can handle.
var publishers = []string{"google", "anthropic"}

// publisherModelsResponse is the publishers.models.list response body.
type publisherModelsResponse struct {
	PublisherModels []publisherModel `json:"publisherModels"`
	NextPageToken   string           `json:"nextPageToken"`
}

type publisherModel struct {
	Name string `json:"name"` // Format: publishers/{publisher}/models/{model}
}

// fetchModelsPage fetches one page. A separate function lets defer close the
// response body after each page.
func fetchModelsPage(ctx context.Context, httpClient *http.Client, url, projectID, publisher string) (*publisherModelsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_error_create_request"), err)
	}

	req.Header.Set("Accept", "application/json")
	// Vertex AI requires the quota project header.
	req.Header.Set("x-goog-user-project", projectID)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_error_request_failed"), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, errorResponseLimit))
		debuglog.Debug(debuglog.Basic, "API error for %s: status %d, url: %s, body: %s\n", publisher, resp.StatusCode, url, string(bodyBytes))
		return nil, fmt.Errorf(i18n.T("vertexai_error_api_status"), resp.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseSize+1))
	if err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_error_read_response"), err)
	}

	if len(bodyBytes) > maxResponseSize {
		return nil, fmt.Errorf(i18n.T("vertexai_error_response_too_large"), maxResponseSize)
	}

	var response publisherModelsResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf(i18n.T("vertexai_error_parse_response"), err)
	}

	return &response, nil
}

// listPublisherModels fetches models from a specific publisher via the Model Garden API
func listPublisherModels(ctx context.Context, httpClient *http.Client, region, projectID, publisher string) ([]string, error) {
	if region == "" || region == "global" {
		region = defaultModelGardenRegion
	}

	baseURL := fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1beta1/publishers/%s/models", region, publisher)

	var allModels []string
	pageToken := ""

	for {
		url := baseURL
		if pageToken != "" {
			url = fmt.Sprintf("%s?pageToken=%s", baseURL, pageToken)
		}

		response, err := fetchModelsPage(ctx, httpClient, url, projectID, publisher)
		if err != nil {
			return nil, err
		}

		for _, model := range response.PublisherModels {
			modelName := extractModelName(model.Name)
			if modelName != "" {
				allModels = append(allModels, modelName)
			}
		}

		if response.NextPageToken == "" {
			break
		}
		pageToken = response.NextPageToken
	}

	debuglog.Debug(debuglog.Detailed, "Listed %d models from publisher %s\n", len(allModels), publisher)
	return allModels, nil
}

// extractModelName returns "gemini-2.0-flash" for
// "publishers/google/models/gemini-2.0-flash".
func extractModelName(fullName string) string {
	parts := strings.Split(fullName, "/")
	if len(parts) >= 4 && parts[0] == "publishers" && parts[2] == "models" {
		return parts[3]
	}
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return fullName
}

// sortModels puts Gemini models first, then Claude, then the rest. Within a
// group, names sort alphabetically, ignoring case.
func sortModels(models []string) []string {
	sort.Slice(models, func(i, j int) bool {
		pi := modelPriority(models[i])
		pj := modelPriority(models[j])
		if pi != pj {
			return pi < pj
		}
		return strings.ToLower(models[i]) < strings.ToLower(models[j])
	})
	return models
}

// modelPriority returns the sort rank. A lower rank sorts first.
func modelPriority(model string) int {
	lower := strings.ToLower(model)
	switch {
	case strings.HasPrefix(lower, "gemini"):
		return 1
	case strings.HasPrefix(lower, "claude"):
		return 2
	default:
		return 3
	}
}

// knownGeminiModels lists the Gemini models on Vertex AI. The Model Garden
// listing does not return them.
// See: https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/google-models
var knownGeminiModels = []string{
	"gemini-3-pro-preview",
	"gemini-3-flash-preview",
	"gemini-2.5-pro",
	"gemini-2.5-flash",
	"gemini-2.5-flash-lite",
	"gemini-2.0-flash",
	"gemini-2.0-flash-lite",
}

func getKnownGeminiModels() []string {
	return knownGeminiModels
}

func isGeminiModel(modelName string) bool {
	return strings.HasPrefix(strings.ToLower(modelName), "gemini")
}

// isConversationalModel rejects image, video, audio, embedding, legacy, and
// medical model names.
func isConversationalModel(modelName string) bool {
	lower := strings.ToLower(modelName)

	excludePatterns := []string{
		"imagen", // Image generation models
		"imagegeneration",
		"imagetext",
		"image-segmentation",
		"embedding",
		"textembedding",
		"multimodalembedding",
		"text-bison", // Legacy completion models (not chat)
		"text-unicorn",
		"code-bison", // Legacy code models
		"code-gecko",
		"codechat-bison", // Deprecated chat model
		"chat-bison",     // Deprecated chat model
		"veo",            // Video generation
		"chirp",          // Audio/speech models
		"medlm",          // Medical models (restricted)
		"medical",
	}

	for _, pattern := range excludePatterns {
		if strings.Contains(lower, pattern) {
			return false
		}
	}

	return true
}

func filterConversationalModels(models []string) []string {
	var filtered []string
	for _, model := range models {
		if isConversationalModel(model) {
			filtered = append(filtered, model)
		}
	}
	return filtered
}
