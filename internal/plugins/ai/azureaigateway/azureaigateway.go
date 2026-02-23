// Package azureaigateway provides a plugin for Azure API Management (APIM) Gateway
// fronting multiple AI backends (AWS Bedrock, Azure OpenAI, Google Vertex AI).
//
// Each backend only defines what differs: endpoint path, auth header, request/response format.
// The shared HTTP plumbing lives in the Client.
package azureaigateway

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
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
	APIVersion      *plugins.SetupQuestion

	backend    Backend
	httpClient *http.Client
}

// NewClient creates a new Azure AI Gateway client
func NewClient() *Client {
	vendorName := "AzureAIGateway"
	client := &Client{}

	client.PluginBase = plugins.NewVendorPluginBase(vendorName, client.configure)

	client.BackendType = client.AddSetupQuestionCustom("backend", true,
		i18n.T("azureaigateway_backend_type_question"))
	client.GatewayURL = client.AddSetupQuestionCustom("gateway_url", true,
		i18n.T("azureaigateway_gateway_url_question"))
	client.SubscriptionKey = client.AddSetupQuestionCustom("subscription_key", true,
		i18n.T("azureaigateway_subscription_key_question"))
	client.APIVersion = client.AddSetupQuestionCustom("api_version", false,
		i18n.T("azureaigateway_api_version_question"))

	return client
}

// configure initializes the HTTP client and instantiates the appropriate backend
func (c *Client) configure() error {
	if c.GatewayURL.Value == "" {
		return errors.New(i18n.T("azureaigateway_gateway_url_required"))
	}
	parsed, err := url.Parse(c.GatewayURL.Value)
	if err != nil {
		return fmt.Errorf(i18n.T("azureaigateway_invalid_gateway_url"), err)
	}
	if parsed.Scheme != "https" {
		return errors.New(i18n.T("azureaigateway_gateway_url_https_required"))
	}
	if c.SubscriptionKey.Value == "" {
		return errors.New(i18n.T("azureaigateway_subscription_key_required"))
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
		c.backend = NewAzureOpenAIBackend(c.SubscriptionKey.Value, c.APIVersion.Value)
	case "vertex-ai":
		c.backend = NewVertexAIBackend(c.SubscriptionKey.Value)
	default:
		return fmt.Errorf(i18n.T("azureaigateway_unsupported_backend"), backendType)
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
		return nil, errors.New(i18n.T("azureaigateway_backend_not_initialized"))
	}
	return c.backend.ListModels()
}

// Send sends a non-streaming request through the APIM gateway.
// This is the single implementation of HTTP plumbing shared by all backends.
func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if c.backend == nil {
		return "", errors.New(i18n.T("azureaigateway_backend_not_initialized"))
	}

	bodyBytes, err := c.backend.PrepareRequest(msgs, opts)
	if err != nil {
		return "", fmt.Errorf(i18n.T("azureaigateway_prepare_request_failed"), err)
	}

	endpoint := c.backend.BuildEndpoint(c.GatewayURL.Value, opts.Model)
	debuglog.Debug(debuglog.Detailed, "AzureAIGateway request to %s\n", endpoint)

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf(i18n.T("azureaigateway_failed_create_request"), err)
	}

	req.Header.Set("Content-Type", "application/json")
	headerName, headerValue := c.backend.AuthHeader()
	req.Header.Set(headerName, headerValue)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf(i18n.T("azureaigateway_http_request_failed"), err)
	}
	defer resp.Body.Close()

	// Limit response body size to 10MB to prevent memory exhaustion
	limitedBody := io.LimitReader(resp.Body, 10*1024*1024)
	respBody, err := io.ReadAll(limitedBody)
	if err != nil {
		return "", fmt.Errorf(i18n.T("azureaigateway_failed_read_response"), err)
	}

	debuglog.Debug(debuglog.Detailed, "AzureAIGateway response status: %d\n", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		debugBody := string(respBody)
		if len(debugBody) > 2000 {
			debugBody = debugBody[:2000] + "...[truncated]"
		}
		debuglog.Debug(debuglog.Detailed, "AzureAIGateway error body: %s\n", debugBody)
		errMsg := string(respBody)
		if len(errMsg) > 500 {
			errMsg = errMsg[:500] + "..."
		}
		return "", fmt.Errorf(i18n.T("azureaigateway_http_error"), resp.StatusCode, errMsg)
	}

	return c.backend.ParseResponse(respBody)
}

// SendStream falls back to non-streaming (APIM gateway doesn't support SSE pass-through).
//
// NOTE: This method uses context.Background() because the ai.Vendor interface does not
// accept a context parameter for SendStream. If the caller disconnects, this request will
// continue until the gateway timeout (300s). A future update to the ai.Vendor interface
// should add context propagation to SendStream.
func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)
	if c.backend == nil {
		return errors.New(i18n.T("azureaigateway_backend_not_initialized"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), gatewayTimeout)
	defer cancel()

	result, err := c.Send(ctx, msgs, opts)
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
