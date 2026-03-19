// Package codex provides a subscription-backed OpenAI OAuth vendor that talks
// to the private Codex backend for supported models.
package codex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	plugins "github.com/danielmiessler/fabric/internal/plugins"
	openaivendor "github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	openaiapi "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared/constant"
)

const (
	vendorName            = "Codex"
	defaultBaseURL        = "https://chatgpt.com/backend-api/codex"
	defaultAuthBaseURL    = "https://auth.openai.com"
	defaultClientVersion  = "1.0.0"
	defaultOriginator     = "codex_cli_rs"
	defaultUserAgent      = "codex_cli_rs/fabric"
	defaultCallbackPort   = 1455
	oauthClientID         = "app_EMoamEEZ73f0CkXaXp7hrann"
	oauthCallbackPath     = "/auth/callback"
	oauthStateBytes       = 32
	oauthVerifierBytes    = 32
	oauthTimeout          = 5 * time.Minute
	tokenRefreshLeeway    = 5 * time.Minute
	modelsRequestTimeout  = 30 * time.Second
	defaultRoundTripLimit = 4096
)

const oauthScope = "openid profile email offline_access api.connectors.read api.connectors.invoke"

type Client struct {
	*openaivendor.Client

	AccessToken  *plugins.Setting
	RefreshToken *plugins.Setting
	AccountID    *plugins.Setting
	AuthBaseURL  *plugins.SetupQuestion

	authHTTPClient *http.Client
	apiHTTPClient  *http.Client

	tokenMu sync.Mutex
}

type modelInfo struct {
	Slug           string `json:"slug"`
	SupportedInAPI bool   `json:"supported_in_api"`
	Visibility     string `json:"visibility"`
}

type modelsResponse struct {
	Models []modelInfo `json:"models"`
}

// NewClient creates a new Codex vendor client.
func NewClient() *Client {
	client := &Client{}
	client.Client = openaivendor.NewClientCompatibleNoSetupQuestions(vendorName, client.configure)
	client.ImplementsResponses = true

	client.AccessToken = client.AddSetting("Access Token", false)
	client.RefreshToken = client.AddSetting("Refresh Token", true)
	client.AccountID = client.AddSetting("Account ID", true)

	client.ApiBaseURL = client.AddSetupQuestionWithEnvName("Base URL", false,
		"Enter your Codex API base URL")
	client.ApiBaseURL.Value = defaultBaseURL

	client.AuthBaseURL = client.AddSetupQuestionWithEnvName("Auth Base URL", false,
		"Enter your Codex OAuth base URL")
	client.AuthBaseURL.Value = defaultAuthBaseURL

	client.authHTTPClient = &http.Client{Timeout: modelsRequestTimeout}
	return client
}

// Setup runs interactive Codex configuration, including browser-based OAuth.
func (c *Client) Setup() error {
	if err := c.ApiBaseURL.Ask(vendorName); err != nil {
		return err
	}
	if err := c.AuthBaseURL.Ask(vendorName); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), oauthTimeout)
	defer cancel()

	fmt.Println()
	fmt.Println(i18n.T("codex_starting_browser_login"))
	debuglog.Debug(debuglog.Detailed, "Codex setup: starting OAuth flow against %s\n", c.AuthBaseURL.Value)

	tokens, err := c.runOAuthFlow(ctx, openBrowser)
	if err != nil {
		return err
	}

	accountID, err := c.extractAccountID(tokens.IDToken, tokens.AccessToken)
	if err != nil {
		return err
	}

	c.setSettingValue(c.AccessToken, tokens.AccessToken)
	c.setSettingValue(c.RefreshToken, tokens.RefreshToken)
	c.setSettingValue(c.AccountID, accountID)

	return c.configure()
}

func (c *Client) configure() error {
	c.authHTTPClient = &http.Client{Timeout: modelsRequestTimeout}

	if strings.TrimSpace(c.ApiBaseURL.Value) == "" {
		c.ApiBaseURL.Value = defaultBaseURL
	}
	if strings.TrimSpace(c.AuthBaseURL.Value) == "" {
		c.AuthBaseURL.Value = defaultAuthBaseURL
	}
	if strings.TrimSpace(c.RefreshToken.Value) == "" {
		return errors.New(i18n.T("codex_refresh_token_required"))
	}

	if _, _, err := c.ensureAccessToken(context.Background(), false); err != nil {
		return err
	}

	transport := &authTransport{
		client:  c,
		wrapped: http.DefaultTransport,
	}
	c.apiHTTPClient = &http.Client{Transport: transport}

	apiClient := openaiapi.NewClient(
		option.WithBaseURL(strings.TrimRight(c.ApiBaseURL.Value, "/")),
		option.WithHTTPClient(c.apiHTTPClient),
	)
	c.ApiClient = &apiClient
	debuglog.Debug(debuglog.Detailed, "Codex configure: authenticated account=%s base_url=%s\n", c.AccountID.Value, c.ApiBaseURL.Value)

	return nil
}

// ListModels returns the Codex models available to the configured account.
func (c *Client) ListModels() ([]string, error) {
	if c.apiHTTPClient == nil {
		if err := c.configure(); err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), modelsRequestTimeout)
	defer cancel()

	modelsURL := strings.TrimRight(c.ApiBaseURL.Value, "/") + "/models"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, modelsURL, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("client_version", codexClientVersion())
	req.URL.RawQuery = query.Encode()
	debuglog.Debug(debuglog.Trace, "Codex ListModels request: %s\n", req.URL.String())

	resp, err := c.apiHTTPClient.Do(req)
	if err != nil {
		return nil, c.mapRequestError(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, c.errorFromHTTPResponse(resp.StatusCode, body)
	}

	var decoded modelsResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, fmt.Errorf("failed to decode Codex models response: %w", err)
	}

	models := make([]string, 0, len(decoded.Models))
	for _, model := range decoded.Models {
		if model.Slug == "" || !model.SupportedInAPI || model.Visibility != "list" {
			continue
		}
		models = append(models, model.Slug)
	}

	return models, nil
}

// Send sends a request to Codex and returns the final text output.
func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if opts.ImageFile != "" {
		return "", errors.New(i18n.T("codex_image_file_not_supported"))
	}
	if c.ApiClient == nil {
		if err := c.configure(); err != nil {
			return "", err
		}
	}

	req := c.buildCodexResponseParams(msgs, opts)
	stream := c.ApiClient.Responses.NewStreaming(ctx, req)
	defer stream.Close()

	var (
		builder       strings.Builder
		completedResp *responses.Response
	)
	for stream.Next() {
		event := stream.Current()
		switch event.Type {
		case string(constant.ResponseOutputTextDelta("").Default()):
			builder.WriteString(event.AsResponseOutputTextDelta().Delta)
		case "response.completed":
			resp := event.AsResponseCompleted().Response
			completedResp = &resp
		}
	}

	if err := c.mapRequestError(stream.Err()); err != nil {
		return "", err
	}
	if completedResp != nil {
		return c.ExtractText(completedResp), nil
	}

	return builder.String(), nil
}

// SendStream sends a request to Codex and streams the response text updates.
func (c *Client) SendStream(
	msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) error {
	defer close(channel)

	if opts.ImageFile != "" {
		return errors.New(i18n.T("codex_image_file_not_supported"))
	}
	if c.ApiClient == nil {
		if err := c.configure(); err != nil {
			return err
		}
	}

	req := c.buildCodexResponseParams(msgs, opts)
	stream := c.ApiClient.Responses.NewStreaming(context.Background(), req)
	defer stream.Close()
	for stream.Next() {
		event := stream.Current()
		switch event.Type {
		case string(constant.ResponseOutputTextDelta("").Default()):
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.AsResponseOutputTextDelta().Delta,
			}
		case string(constant.ResponseOutputTextDone("").Default()):
			continue
		}
	}

	if stream.Err() == nil {
		channel <- domain.StreamUpdate{
			Type:    domain.StreamTypeContent,
			Content: "\n",
		}
	}

	return c.mapRequestError(stream.Err())
}

func (c *Client) buildCodexResponseParams(
	msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions,
) responses.ResponseNewParams {
	instructions, filteredMsgs := codexInstructionsAndMessages(msgs)
	req := c.BuildResponseParams(filteredMsgs, opts)
	req.Instructions = openaiapi.String(instructions)
	req.Store = openaiapi.Bool(false)
	return req
}

func codexInstructionsAndMessages(
	msgs []*chat.ChatCompletionMessage,
) (string, []*chat.ChatCompletionMessage) {
	filtered := make([]*chat.ChatCompletionMessage, 0, len(msgs))
	instructions := make([]string, 0, len(msgs))

	for _, msg := range msgs {
		if msg == nil {
			continue
		}
		switch msg.Role {
		case chat.ChatMessageRoleSystem, chat.ChatMessageRoleDeveloper:
			if text := codexMessageText(*msg); text != "" {
				instructions = append(instructions, text)
			}
		default:
			filtered = append(filtered, msg)
		}
	}

	if len(instructions) == 0 {
		return "You are a helpful assistant.", filtered
	}

	return strings.Join(instructions, "\n\n"), filtered
}

func codexMessageText(msg chat.ChatCompletionMessage) string {
	if text := strings.TrimSpace(msg.Content); text != "" {
		return text
	}

	if len(msg.MultiContent) == 0 {
		return ""
	}

	parts := make([]string, 0, len(msg.MultiContent))
	for _, part := range msg.MultiContent {
		if part.Type == chat.ChatMessagePartTypeText {
			if text := strings.TrimSpace(part.Text); text != "" {
				parts = append(parts, text)
			}
		}
	}

	return strings.Join(parts, "\n")
}
