// Package copilot provides integration with Microsoft 365 Copilot Chat API.
// This vendor allows Fabric to interact with Microsoft 365 Copilot, which provides
// AI capabilities grounded in your organization's Microsoft 365 data.
//
// Requirements:
// - Microsoft 365 Copilot license for each user
// - Microsoft 365 E3 or E5 subscription (or equivalent)
// - Azure AD app registration with appropriate permissions
//
// The Chat API is currently in preview and requires delegated (work or school account)
// permissions. Application permissions are not supported.
package copilot

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"golang.org/x/oauth2"
)

const (
	vendorName = "Copilot"

	// Microsoft Graph API endpoints
	defaultBaseURL    = "https://graph.microsoft.com/beta/copilot"
	conversationsPath = "/conversations"

	// OAuth2 endpoints for Microsoft identity platform
	microsoftAuthURL  = "https://login.microsoftonline.com/%s/oauth2/v2.0/authorize"
	microsoftTokenURL = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"

	// Default scopes required for Copilot Chat API
	// These are the minimum required permissions
	defaultScopes = "Sites.Read.All Mail.Read People.Read.All OnlineMeetingTranscript.Read.All Chat.Read ChannelMessage.Read.All ExternalItem.Read.All offline_access"

	// Model name exposed by Copilot (single model)
	copilotModelName = "microsoft-365-copilot"
)

// NewClient creates a new Microsoft 365 Copilot client.
func NewClient() *Client {
	c := &Client{}

	c.PluginBase = &plugins.PluginBase{
		Name:            vendorName,
		EnvNamePrefix:   plugins.BuildEnvVariablePrefix(vendorName),
		ConfigureCustom: c.configure,
	}

	// Setup questions for configuration
	c.TenantID = c.AddSetupQuestion("Tenant ID", true)
	c.TenantID.Question = "Enter your Azure AD Tenant ID (e.g., contoso.onmicrosoft.com or GUID)"

	c.ClientID = c.AddSetupQuestion("Client ID", true)
	c.ClientID.Question = "Enter your Azure AD Application (Client) ID"

	c.ClientSecret = c.AddSetupQuestion("Client Secret", false)
	c.ClientSecret.Question = "Enter your Azure AD Client Secret (optional, for confidential clients)"

	c.AccessToken = c.AddSetupQuestion("Access Token", false)
	c.AccessToken.Question = "Enter a pre-obtained OAuth2 Access Token (optional, for testing)"

	c.RefreshToken = c.AddSetupQuestion("Refresh Token", false)
	c.RefreshToken.Question = "Enter a pre-obtained OAuth2 Refresh Token (optional)"

	c.ApiBaseURL = c.AddSetupQuestion("API Base URL", false)
	c.ApiBaseURL.Value = defaultBaseURL

	c.TimeZone = c.AddSetupQuestion("Time Zone", false)
	c.TimeZone.Value = "America/New_York"
	c.TimeZone.Question = "Enter your timezone (e.g., America/New_York, Europe/London)"

	return c
}

// Client represents a Microsoft 365 Copilot API client.
type Client struct {
	*plugins.PluginBase

	// Configuration
	TenantID     *plugins.SetupQuestion
	ClientID     *plugins.SetupQuestion
	ClientSecret *plugins.SetupQuestion
	AccessToken  *plugins.SetupQuestion
	RefreshToken *plugins.SetupQuestion
	ApiBaseURL   *plugins.SetupQuestion
	TimeZone     *plugins.SetupQuestion

	// Runtime state
	httpClient   *http.Client
	oauth2Config *oauth2.Config
	token        *oauth2.Token
}

// configure initializes the client with OAuth2 configuration.
func (c *Client) configure() error {
	if c.TenantID.Value == "" || c.ClientID.Value == "" {
		return errors.New(i18n.T("copilot_tenant_client_id_required"))
	}

	// Build OAuth2 configuration
	c.oauth2Config = &oauth2.Config{
		ClientID:     c.ClientID.Value,
		ClientSecret: c.ClientSecret.Value,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf(microsoftAuthURL, c.TenantID.Value),
			TokenURL: fmt.Sprintf(microsoftTokenURL, c.TenantID.Value),
		},
		Scopes: strings.Split(defaultScopes, " "),
	}

	// If we have pre-configured tokens, use them
	if c.AccessToken.Value != "" {
		c.token = &oauth2.Token{
			AccessToken:  c.AccessToken.Value,
			RefreshToken: c.RefreshToken.Value,
			TokenType:    "Bearer",
		}
		// If we have a refresh token, set expiry in the past to trigger refresh
		if c.RefreshToken.Value != "" && c.ClientSecret.Value != "" {
			c.token.Expiry = time.Now().Add(-time.Hour)
		}
	}

	// Create HTTP client with OAuth2 token source
	if c.token != nil {
		tokenSource := c.oauth2Config.TokenSource(context.Background(), c.token)
		c.httpClient = oauth2.NewClient(context.Background(), tokenSource)
	} else {
		// No tokens available - will need device code flow or manual token
		c.httpClient = &http.Client{Timeout: 120 * time.Second}
	}

	return nil
}

// IsConfigured returns true if the client has valid configuration.
func (c *Client) IsConfigured() bool {
	// Minimum requirement: tenant ID and client ID
	if c.TenantID.Value == "" || c.ClientID.Value == "" {
		return false
	}
	// Must have either an access token or ability to get one
	return c.AccessToken.Value != "" || (c.RefreshToken.Value != "" && c.ClientSecret.Value != "")
}

// ListModels returns the available models.
// Microsoft 365 Copilot exposes a single model - the Copilot service itself.
func (c *Client) ListModels() ([]string, error) {
	// Copilot doesn't expose multiple models - it's a unified service
	// We expose it as a single "model" for consistency with Fabric's architecture
	return []string{copilotModelName}, nil
}

// Send sends a message to Copilot and returns the response.
func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	// Create a conversation
	conversationID, err := c.createConversation(ctx)
	if err != nil {
		return "", fmt.Errorf(i18n.T("copilot_failed_create_conversation"), err)
	}

	// Build the message content from chat messages
	messageText := c.buildMessageText(msgs)

	// Send the chat message
	response, err := c.sendChatMessage(ctx, conversationID, messageText)
	if err != nil {
		return "", fmt.Errorf(i18n.T("copilot_failed_send_message"), err)
	}

	return response, nil
}

// SendStream sends a message to Copilot and streams the response.
func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)

	ctx := context.Background()

	// Create a conversation
	conversationID, err := c.createConversation(ctx)
	if err != nil {
		return fmt.Errorf(i18n.T("copilot_failed_create_conversation"), err)
	}

	// Build the message content from chat messages
	messageText := c.buildMessageText(msgs)

	// Send the streaming chat message
	if err := c.sendChatMessageStream(ctx, conversationID, messageText, channel); err != nil {
		return fmt.Errorf(i18n.T("copilot_failed_stream_message"), err)
	}

	return nil
}

// buildMessageText combines chat messages into a single prompt for Copilot.
func (c *Client) buildMessageText(msgs []*chat.ChatCompletionMessage) string {
	var parts []string

	for _, msg := range msgs {
		content := strings.TrimSpace(msg.Content)
		if content == "" {
			continue
		}

		switch msg.Role {
		case chat.ChatMessageRoleSystem:
			// Prepend system messages as context
			parts = append([]string{content}, parts...)
		case chat.ChatMessageRoleUser, chat.ChatMessageRoleAssistant:
			parts = append(parts, content)
		}
	}

	return strings.Join(parts, "\n\n")
}

// createConversation creates a new Copilot conversation.
func (c *Client) createConversation(ctx context.Context) (string, error) {
	url := c.ApiBaseURL.Value + conversationsPath

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString("{}"))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthHeader(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf(i18n.T("copilot_error_create_conversation"), resp.Status, string(body))
	}

	var result conversationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	debuglog.Debug(debuglog.Detailed, i18n.T("copilot_debug_created_conversation")+"\n", result.ID)
	return result.ID, nil
}

// sendChatMessage sends a message to an existing conversation (synchronous).
func (c *Client) sendChatMessage(ctx context.Context, conversationID, messageText string) (string, error) {
	url := fmt.Sprintf("%s%s/%s/chat", c.ApiBaseURL.Value, conversationsPath, conversationID)

	reqBody := chatRequest{
		Message: messageParam{
			Text: messageText,
		},
		LocationHint: locationHint{
			TimeZone: c.TimeZone.Value,
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthHeader(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf(i18n.T("copilot_error_chat_request"), resp.Status, string(body))
	}

	var result conversationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Extract the assistant's response from messages
	return c.extractResponseText(result.Messages), nil
}

// sendChatMessageStream sends a message and streams the response via SSE.
func (c *Client) sendChatMessageStream(ctx context.Context, conversationID, messageText string, channel chan domain.StreamUpdate) error {
	url := fmt.Sprintf("%s%s/%s/chatOverStream", c.ApiBaseURL.Value, conversationsPath, conversationID)

	reqBody := chatRequest{
		Message: messageParam{
			Text: messageText,
		},
		LocationHint: locationHint{
			TimeZone: c.TimeZone.Value,
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	c.addAuthHeader(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(i18n.T("copilot_error_stream_request"), resp.Status, string(body))
	}

	// Parse SSE stream
	return c.parseSSEStream(resp.Body, channel)
}

// parseSSEStream parses the Server-Sent Events stream from Copilot.
func (c *Client) parseSSEStream(reader io.Reader, channel chan domain.StreamUpdate) error {
	scanner := bufio.NewScanner(reader)
	var lastMessageText string

	for scanner.Scan() {
		line := scanner.Text()

		// SSE format: "data: {...json...}"
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		jsonData := strings.TrimPrefix(line, "data: ")
		if jsonData == "" {
			continue
		}

		var event conversationResponse
		if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
			debuglog.Debug(debuglog.Detailed, i18n.T("copilot_debug_failed_parse_sse_event")+"\n", err)
			continue
		}

		// Extract new text from the response
		newText := c.extractResponseText(event.Messages)
		if newText != "" && newText != lastMessageText {
			// Send only the delta (new content)
			if delta, ok := strings.CutPrefix(newText, lastMessageText); ok {
				if delta != "" {
					channel <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: delta}
				}
			} else {
				// Complete message replacement
				channel <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: newText}
			}
			lastMessageText = newText
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(i18n.T("copilot_error_reading_stream"), err)
	}

	channel <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: "\n"}
	return nil
}

// extractResponseText extracts the assistant's response from messages.
func (c *Client) extractResponseText(messages []responseMessage) string {
	// Find the last assistant message (Copilot's response)
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		// Response messages from Copilot have the copilotConversationResponseMessage type
		if msg.ODataType == "#microsoft.graph.copilotConversationResponseMessage" {
			if msg.Text != "" {
				return msg.Text
			}
		}
	}
	return ""
}

// addAuthHeader adds the authorization header to a request.
func (c *Client) addAuthHeader(req *http.Request) {
	if c.token != nil && c.token.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	} else if c.AccessToken.Value != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken.Value)
	}
}

// API request/response types

type chatRequest struct {
	Message             messageParam         `json:"message"`
	LocationHint        locationHint         `json:"locationHint"`
	AdditionalContext   []contextMessage     `json:"additionalContext,omitempty"`
	ContextualResources *contextualResources `json:"contextualResources,omitempty"`
}

type messageParam struct {
	Text string `json:"text"`
}

type locationHint struct {
	TimeZone string `json:"timeZone"`
}

type contextMessage struct {
	Text string `json:"text"`
}

type contextualResources struct {
	Files      []fileResource `json:"files,omitempty"`
	WebContext *webContext    `json:"webContext,omitempty"`
}

type fileResource struct {
	URI string `json:"uri"`
}

type webContext struct {
	IsWebEnabled bool `json:"isWebEnabled"`
}

type conversationResponse struct {
	ID              string            `json:"id"`
	CreatedDateTime string            `json:"createdDateTime"`
	DisplayName     string            `json:"displayName"`
	State           string            `json:"state"`
	TurnCount       int               `json:"turnCount"`
	Messages        []responseMessage `json:"messages,omitempty"`
}

type responseMessage struct {
	ODataType       string        `json:"@odata.type"`
	ID              string        `json:"id"`
	Text            string        `json:"text"`
	CreatedDateTime string        `json:"createdDateTime"`
	AdaptiveCards   []any         `json:"adaptiveCards,omitempty"`
	Attributions    []attribution `json:"attributions,omitempty"`
}

type attribution struct {
	AttributionType     string `json:"attributionType"`
	ProviderDisplayName string `json:"providerDisplayName"`
	AttributionSource   string `json:"attributionSource"`
	SeeMoreWebURL       string `json:"seeMoreWebUrl"`
}
