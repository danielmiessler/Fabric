package digitalocean

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
)

const (
	defaultInferenceBaseURL = "https://inference.do-ai.run/v1"
	controlPlaneModelsURL   = "https://api.digitalocean.com/v2/gen-ai/models"
	errorResponseLimit      = 1024
	maxResponseSize         = 10 * 1024 * 1024
)

type Client struct {
	*openai.Client
	ControlPlaneToken *plugins.SetupQuestion
	httpClient        *http.Client
}

type modelsResponse struct {
	Models []modelDetails `json:"models"`
}

type modelDetails struct {
	InferenceName string `json:"inference_name"`
	Name          string `json:"name"`
	UUID          string `json:"uuid"`
}

func NewClient() *Client {
	base := openai.NewClientCompatibleNoSetupQuestions("DigitalOcean", nil)
	base.ApiKey = base.AddSetupQuestion("Inference Key", true)
	base.ApiBaseURL = base.AddSetupQuestion("Inference Base URL", false)
	base.ApiBaseURL.Value = defaultInferenceBaseURL
	base.ImplementsResponses = false

	client := &Client{
		Client: base,
	}
	client.ControlPlaneToken = client.AddSetupQuestion("Token", false)
	return client
}

func (c *Client) ListModels() ([]string, error) {
	if c.ControlPlaneToken.Value == "" {
		models, err := c.Client.ListModels()
		if err == nil && len(models) > 0 {
			return models, nil
		}
		if err != nil {
			return nil, fmt.Errorf(
				"DigitalOcean model list unavailable: %w. Set DIGITALOCEAN_TOKEN to fetch models from the control plane",
				err,
			)
		}
		return nil, fmt.Errorf("%s", i18n.T("digitalocean_model_list_unavailable"))
	}
	return c.fetchModelsFromControlPlane(context.Background())
}

func (c *Client) fetchModelsFromControlPlane(ctx context.Context) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	fullURL, err := url.Parse(controlPlaneModelsURL)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("digitalocean_failed_parse_control_plane_url"), err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ControlPlaneToken.Value))
	req.Header.Set("Accept", "application/json")

	client := c.httpClient
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, readErr := io.ReadAll(io.LimitReader(resp.Body, errorResponseLimit))
		if readErr != nil {
			return nil, fmt.Errorf(
				"DigitalOcean models request failed with status %d: %w",
				resp.StatusCode,
				readErr,
			)
		}
		return nil, fmt.Errorf(
			"DigitalOcean models request failed with status %d: %s",
			resp.StatusCode,
			string(bodyBytes),
		)
	}

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseSize+1))
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > maxResponseSize {
		return nil, fmt.Errorf(i18n.T("openai_models_response_too_large"), c.GetName(), maxResponseSize)
	}

	var payload modelsResponse
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return nil, err
	}

	models := make([]string, 0, len(payload.Models))
	seen := make(map[string]struct{}, len(payload.Models))
	for _, model := range payload.Models {
		var value string
		switch {
		case model.InferenceName != "":
			value = model.InferenceName
		case model.Name != "":
			value = model.Name
		case model.UUID != "":
			value = model.UUID
		}
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		models = append(models, value)
	}
	return models, nil
}
