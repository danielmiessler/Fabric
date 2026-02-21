// Package azureaigateway provides a plugin for Azure API Management (APIM) Gateway
// fronting multiple AI backends (AWS Bedrock, Azure OpenAI, Google Vertex AI).
//
// Each backend only defines what differs: endpoint path, auth header, request/response format.
// The shared HTTP plumbing lives in the Client.
package azureaigateway

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
)

const gatewayTimeout = 300 * time.Second

// Ensure Client implements the ai.Vendor interface
var _ ai.Vendor = (*Client)(nil)

// Backend defines the interface that all Azure AI Gateway backends must implement.
// Each backend only provides what is unique to its API format.
// The shared HTTP mechanics (request execution, error handling, streaming fallback)
// are handled by the Client.
type Backend interface {
	// ListModels returns the list of models available for this backend
	ListModels() ([]string, error)

	// BuildEndpoint constructs the full API endpoint URL for the given model
	BuildEndpoint(baseURL, model string) string

	// AuthHeader returns the header name and value for authentication.
	// Each APIM backend uses a different auth header:
	//   Bedrock:      "Authorization", "Bearer <key>"
	//   Azure OpenAI: "api-key", "<key>"
	//   Vertex AI:    "x-goog-api-key", "<key>"
	AuthHeader() (name, value string)

	// PrepareRequest prepares the HTTP request body for this backend's API format
	PrepareRequest(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) ([]byte, error)

	// ParseResponse parses the HTTP response body into text content
	ParseResponse(body []byte) (string, error)
}

// Client implements the Azure AI Gateway vendor for Fabric.
// It supports multiple backends (Bedrock, Azure OpenAI, Vertex AI) through
// a unified Azure APIM Gateway with shared subscription key authentication.
type Client struct {
	*plugins.PluginBase
	BackendType     *plugins.SetupQuestion
	GatewayURL      *plugins.SetupQuestion
	SubscriptionKey *plugins.SetupQuestion

	backend    Backend
	httpClient *http.Client
}

// NewClient creates a new Azure AI Gateway client
func NewClient() *Client {
	vendorName := "AzureAIGateway"
	client := &Client{}

	client.PluginBase = plugins.NewVendorPluginBase(vendorName, client.configure)

	client.BackendType = client.AddSetupQuestionCustom("backend", true,
		"Select backend type (bedrock, azure-openai, vertex-ai)")
	client.GatewayURL = client.AddSetupQuestionCustom("gateway_url", true,
		"Enter your Azure APIM Gateway base URL (e.g., https://gateway.company.com)")
	client.SubscriptionKey = client.AddSetupQuestionCustom("subscription_key", true,
		"Enter your Azure APIM subscription key")

	return client
}

// configure initializes the HTTP client and instantiates the appropriate backend
func (c *Client) configure() error {
	if c.GatewayURL.Value == "" {
		return fmt.Errorf("Azure APIM Gateway URL is required")
	}
	parsed, err := url.Parse(c.GatewayURL.Value)
	if err != nil {
		return fmt.Errorf("invalid gateway URL: %w", err)
	}
	if parsed.Scheme != "https" {
		return fmt.Errorf("gateway URL must use HTTPS scheme, got %q", parsed.Scheme)
	}
	if c.SubscriptionKey.Value == "" {
		return fmt.Errorf("Azure APIM subscription key is required")
	}

	// Normalize backend type; default to bedrock for backward compatibility
	backendType := strings.ToLower(strings.TrimSpace(c.BackendType.Value))
	if backendType == "" {
		backendType = "bedrock"
		c.BackendType.Value = backendType
	}

	c.httpClient = &http.Client{Timeout: gatewayTimeout}

	switch backendType {
	case "bedrock":
		c.backend = NewBedrockBackend(c.SubscriptionKey.Value)
	case "azure-openai":
		c.backend = NewAzureOpenAIBackend(c.SubscriptionKey.Value)
	case "vertex-ai":
		c.backend = NewVertexAIBackend(c.SubscriptionKey.Value)
	default:
		return fmt.Errorf("unsupported backend: %s (valid options: bedrock, azure-openai, vertex-ai)", backendType)
	}

	return nil
}

// IsConfigured returns true if both gateway URL and subscription key are configured
func (c *Client) IsConfigured() bool {
	return c.GatewayURL.Value != "" && c.SubscriptionKey.Value != ""
}

// ListModels delegates to the active backend
func (c *Client) ListModels() ([]string, error) {
	if c.backend == nil {
		return nil, fmt.Errorf("backend not initialized - run 'fabric --setup' to configure")
	}
	return c.backend.ListModels()
}

// Send sends a non-streaming request through the APIM gateway.
// This is the single implementation of HTTP plumbing shared by all backends.
func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if c.backend == nil {
		return "", fmt.Errorf("backend not initialized - run 'fabric --setup' to configure")
	}

	bodyBytes, err := c.backend.PrepareRequest(msgs, opts)
	if err != nil {
		return "", fmt.Errorf("AzureAIGateway: %w", err)
	}

	endpoint := c.backend.BuildEndpoint(c.GatewayURL.Value, opts.Model)
	debuglog.Debug(debuglog.Detailed, "AzureAIGateway request to %s\n", endpoint)

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("AzureAIGateway: failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	headerName, headerValue := c.backend.AuthHeader()
	req.Header.Set(headerName, headerValue)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("AzureAIGateway: HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("AzureAIGateway: failed to read response: %w", err)
	}

	debuglog.Debug(debuglog.Detailed, "AzureAIGateway response status: %d\n", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		debuglog.Debug(debuglog.Detailed, "AzureAIGateway error body: %s\n", string(respBody))
		errMsg := string(respBody)
		if len(errMsg) > 200 {
			errMsg = errMsg[:200] + "... (truncated)"
		}
		return "", fmt.Errorf("AzureAIGateway: HTTP %d: %s", resp.StatusCode, errMsg)
	}

	return c.backend.ParseResponse(respBody)
}

// SendStream falls back to non-streaming (APIM gateway doesn't support SSE pass-through).
func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)
	if c.backend == nil {
		return fmt.Errorf("backend not initialized - run 'fabric --setup' to configure")
	}

	result, err := c.Send(context.Background(), msgs, opts)
	if err != nil {
		return err
	}
	channel <- domain.StreamUpdate{
		Type:    domain.StreamTypeContent,
		Content: result,
	}
	return nil
}

// NeedsRawMode returns false as Azure AI Gateway doesn't require raw mode
func (c *Client) NeedsRawMode(modelName string) bool {
	return false
}
