// Package codex provides a subscription-backed OpenAI OAuth vendor that talks
// to the private Codex backend for supported models.
package codex

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"slices"
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

var errReplayBodyUnavailable = errors.New("request body cannot be replayed for Codex re-authentication retry")

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

type oauthTokens struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type refreshRequest struct {
	ClientID     string `json:"client_id"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type refreshResponse struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type oauthResult struct {
	tokens oauthTokens
	err    error
}

type modelInfo struct {
	Slug           string `json:"slug"`
	SupportedInAPI bool   `json:"supported_in_api"`
	Visibility     string `json:"visibility"`
}

type modelsResponse struct {
	Models []modelInfo `json:"models"`
}

type tokenClaims struct {
	Exp     int64           `json:"exp"`
	Auth    tokenAuthClaims `json:"https://api.openai.com/auth"`
	Profile tokenProfile    `json:"https://api.openai.com/profile"`
	Email   string          `json:"email"`
}

type tokenAuthClaims struct {
	ChatGPTAccountID string `json:"chatgpt_account_id"`
	ChatGPTPlanType  string `json:"chatgpt_plan_type"`
	UserID           string `json:"user_id"`
	ChatGPTUserID    string `json:"chatgpt_user_id"`
}

type tokenProfile struct {
	Email string `json:"email"`
}

type authTransport struct {
	client  *Client
	wrapped http.RoundTripper
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

func (c *Client) runOAuthFlow(
	ctx context.Context,
	openBrowserFn func(string) error,
) (oauthTokens, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", defaultCallbackPort))
	if err != nil {
		return oauthTokens{}, fmt.Errorf("failed to start local OAuth callback server: %w", err)
	}
	defer listener.Close()
	debuglog.Debug(debuglog.Detailed, "Codex OAuth callback listener started on 127.0.0.1:%d\n", defaultCallbackPort)

	pkce, err := generatePKCECodes()
	if err != nil {
		return oauthTokens{}, err
	}

	state, err := randomBase64URL(oauthStateBytes)
	if err != nil {
		return oauthTokens{}, err
	}

	callbackURL := (&url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%d", defaultCallbackPort),
		Path:   oauthCallbackPath,
	}).String()
	authURL, err := buildAuthorizeURL(c.AuthBaseURL.Value, callbackURL, pkce, state)
	if err != nil {
		return oauthTokens{}, err
	}

	results := make(chan oauthResult, 1)
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.handleOAuthCallback(w, r, callbackURL, pkce, state, results)
		}),
	}

	serveDone := make(chan error, 1)
	go func() {
		err := server.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveDone <- err
			return
		}
		serveDone <- nil
	}()

	if err := openBrowserFn(authURL); err != nil {
		fmt.Printf("If your browser did not open, navigate to this URL to authenticate:\n%s\n", authURL)
	}

	select {
	case result := <-results:
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
		<-serveDone
		return result.tokens, result.err
	case err := <-serveDone:
		if err != nil {
			return oauthTokens{}, err
		}
		return oauthTokens{}, errors.New(i18n.T("codex_login_server_stopped"))
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
		<-serveDone
		return oauthTokens{}, errors.New(i18n.T("codex_login_timed_out"))
	}
}

func (c *Client) handleOAuthCallback(
	w http.ResponseWriter,
	r *http.Request,
	callbackURL string,
	pkce pkceCodes,
	expectedState string,
	results chan<- oauthResult,
) {
	if r.URL.Path != oauthCallbackPath {
		http.NotFound(w, r)
		return
	}

	if r.URL.Query().Get("state") != expectedState {
		http.Error(w, "State mismatch", http.StatusBadRequest)
		c.publishOAuthResult(results, oauthResult{
			err: errors.New(i18n.T("codex_login_state_mismatch")),
		})
		return
	}

	if callbackError := strings.TrimSpace(r.URL.Query().Get("error")); callbackError != "" {
		description := strings.TrimSpace(r.URL.Query().Get("error_description"))
		if description != "" {
			http.Error(w, description, http.StatusForbidden)
			c.publishOAuthResult(results, oauthResult{
				err: fmt.Errorf(i18n.T("codex_login_failed"), description),
			})
			return
		}
		http.Error(w, callbackError, http.StatusForbidden)
		c.publishOAuthResult(results, oauthResult{
			err: fmt.Errorf(i18n.T("codex_login_failed"), callbackError),
		})
		return
	}

	code := strings.TrimSpace(r.URL.Query().Get("code"))
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		c.publishOAuthResult(results, oauthResult{
			err: errors.New(i18n.T("codex_login_missing_auth_code")),
		})
		return
	}

	tokens, err := c.exchangeCodeForTokens(r.Context(), callbackURL, pkce, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		c.publishOAuthResult(results, oauthResult{err: err})
		return
	}

	if _, err := c.extractAccountID(tokens.IDToken, tokens.AccessToken); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		c.publishOAuthResult(results, oauthResult{err: err})
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte("<html><body><h1>Codex login completed</h1><p>Return to Fabric.</p></body></html>"))
	c.publishOAuthResult(results, oauthResult{tokens: tokens})
}

func (c *Client) publishOAuthResult(results chan<- oauthResult, result oauthResult) {
	select {
	case results <- result:
	default:
	}
}

func (c *Client) exchangeCodeForTokens(
	ctx context.Context,
	callbackURL string,
	pkce pkceCodes,
	code string,
) (oauthTokens, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", callbackURL)
	form.Set("client_id", oauthClientID)
	form.Set("code_verifier", pkce.CodeVerifier)

	tokenURL := strings.TrimRight(c.AuthBaseURL.Value, "/") + "/oauth/token"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return oauthTokens{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.authHTTPClient.Do(req)
	if err != nil {
		return oauthTokens{}, fmt.Errorf("Codex token exchange failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return oauthTokens{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return oauthTokens{}, c.errorFromHTTPResponse(resp.StatusCode, body)
	}

	var tokens oauthTokens
	if err := json.Unmarshal(body, &tokens); err != nil {
		return oauthTokens{}, fmt.Errorf("failed to decode Codex token exchange response: %w", err)
	}
	if strings.TrimSpace(tokens.AccessToken) == "" || strings.TrimSpace(tokens.RefreshToken) == "" {
		return oauthTokens{}, errors.New(i18n.T("codex_login_missing_tokens"))
	}

	return tokens, nil
}

func (c *Client) ensureAccessToken(ctx context.Context, forceRefresh bool) (string, string, error) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	accessToken := strings.TrimSpace(c.AccessToken.Value)
	accountID := strings.TrimSpace(c.AccountID.Value)

	if !forceRefresh && accessToken != "" && !tokenNeedsRefresh(accessToken, time.Now()) {
		if accountID == "" {
			parsedAccountID, err := extractAccountIDFromJWT(accessToken)
			if err == nil && parsedAccountID != "" {
				accountID = parsedAccountID
				c.setSettingValue(c.AccountID, accountID)
			}
		}
		if accountID != "" {
			return accessToken, accountID, nil
		}
	}

	refreshed, err := c.refreshAccessToken(ctx)
	if err != nil {
		return "", "", err
	}

	refreshedAccountID, err := c.extractAccountID(refreshed.IDToken, refreshed.AccessToken)
	if err != nil {
		return "", "", err
	}
	if accountID != "" && refreshedAccountID != "" && !strings.EqualFold(accountID, refreshedAccountID) {
		return "", "", errors.New(i18n.T("codex_login_account_changed"))
	}

	c.setSettingValue(c.AccessToken, refreshed.AccessToken)
	if strings.TrimSpace(refreshed.RefreshToken) != "" {
		c.setSettingValue(c.RefreshToken, refreshed.RefreshToken)
	}
	c.setSettingValue(c.AccountID, refreshedAccountID)
	debuglog.Debug(debuglog.Detailed, "Codex access token refreshed for account=%s\n", refreshedAccountID)

	return c.AccessToken.Value, c.AccountID.Value, nil
}

func (c *Client) refreshAccessToken(ctx context.Context) (oauthTokens, error) {
	payload := refreshRequest{
		ClientID:     oauthClientID,
		GrantType:    "refresh_token",
		RefreshToken: strings.TrimSpace(c.RefreshToken.Value),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return oauthTokens{}, err
	}

	tokenURL := strings.TrimRight(c.AuthBaseURL.Value, "/") + "/oauth/token"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(string(body)))
	if err != nil {
		return oauthTokens{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.authHTTPClient.Do(req)
	if err != nil {
		return oauthTokens{}, fmt.Errorf("failed to refresh Codex login: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return oauthTokens{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return oauthTokens{}, c.refreshErrorFromResponse(resp.StatusCode, responseBody)
	}

	var refreshed refreshResponse
	if err := json.Unmarshal(responseBody, &refreshed); err != nil {
		return oauthTokens{}, fmt.Errorf("failed to decode refreshed Codex token response: %w", err)
	}
	if strings.TrimSpace(refreshed.AccessToken) == "" {
		return oauthTokens{}, errors.New(i18n.T("codex_token_refresh_missing_access_token"))
	}

	return oauthTokens{
		IDToken:      strings.TrimSpace(refreshed.IDToken),
		AccessToken:  strings.TrimSpace(refreshed.AccessToken),
		RefreshToken: strings.TrimSpace(refreshed.RefreshToken),
	}, nil
}

func (c *Client) extractAccountID(idToken string, accessToken string) (string, error) {
	if accountID, err := extractAccountIDFromJWT(idToken); err == nil && accountID != "" {
		return accountID, nil
	}
	if accountID, err := extractAccountIDFromJWT(accessToken); err == nil && accountID != "" {
		return accountID, nil
	}
	return "", errors.New(i18n.T("codex_login_missing_account_claim"))
}

func (c *Client) setSettingValue(setting *plugins.Setting, value string) {
	setting.Value = value
	if setting.EnvVariable != "" {
		_ = os.Setenv(setting.EnvVariable, value)
	}
}

func (c *Client) errorFromHTTPResponse(statusCode int, body []byte) error {
	message := extractErrorMessage(body)
	if statusCode == http.StatusUnauthorized {
		return errors.New(i18n.T("codex_login_invalid"))
	}
	if isUsageLimitMessage(message) {
		return errors.New(message)
	}
	if message == "" {
		message = fmt.Sprintf("Codex request failed with status %d", statusCode)
	}
	return errors.New(message)
}

func (c *Client) refreshErrorFromResponse(statusCode int, body []byte) error {
	message := extractErrorMessage(body)
	code := strings.ToLower(extractErrorCode(body))

	if statusCode == http.StatusUnauthorized {
		switch code {
		case "refresh_token_expired", "refresh_token_reused", "refresh_token_invalidated":
			return errors.New(i18n.T("codex_login_revoked"))
		default:
			return errors.New(i18n.T("codex_login_refresh_failed"))
		}
	}

	if message == "" {
		message = fmt.Sprintf("failed to refresh Codex login (status %d)", statusCode)
	}
	return errors.New(message)
}

func (c *Client) mapRequestError(err error) error {
	if err == nil {
		return nil
	}

	var apiErr *openaiapi.Error
	if errors.As(err, &apiErr) {
		body := []byte(apiErr.RawJSON())
		if len(body) == 0 {
			body = readAPIErrorBody(apiErr)
		}
		return c.errorFromHTTPResponse(apiErr.StatusCode, body)
	}

	message := err.Error()
	lower := strings.ToLower(message)

	switch {
	case strings.Contains(lower, "status code 401"),
		strings.Contains(lower, "401 unauthorized"),
		strings.Contains(lower, "refresh token"),
		strings.Contains(lower, "chatgpt login"):
		return errors.New(i18n.T("codex_login_invalid"))
	case isUsageLimitMessage(message):
		return errors.New(message)
	default:
		return err
	}
}

func readAPIErrorBody(apiErr *openaiapi.Error) []byte {
	if apiErr == nil || apiErr.Response == nil || apiErr.Response.Body == nil {
		return nil
	}

	body, err := io.ReadAll(apiErr.Response.Body)
	if err != nil {
		return nil
	}
	apiErr.Response.Body = io.NopCloser(strings.NewReader(string(body)))
	return body
}

// RoundTrip adds Codex authentication headers and retries once after a 401.
func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.roundTrip(req, false)
}

func (t *authTransport) roundTrip(req *http.Request, retried bool) (*http.Response, error) {
	token, accountID, err := t.client.ensureAccessToken(req.Context(), false)
	if err != nil {
		return nil, err
	}

	clone, err := cloneRequest(req)
	if err != nil {
		return nil, err
	}
	clone.Header.Set(http.CanonicalHeaderKey("originator"), defaultOriginator)
	clone.Header.Set("User-Agent", defaultUserAgent)
	clone.Header.Set("Authorization", "Bearer "+token)
	clone.Header.Set("ChatGPT-Account-ID", accountID)

	resp, err := t.roundTripper().RoundTrip(clone)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusUnauthorized || retried {
		return resp, nil
	}

	drainAndClose(resp.Body)
	debuglog.Debug(debuglog.Detailed, "Codex request returned 401; attempting token refresh and one retry\n")

	if _, _, err := t.client.ensureAccessToken(req.Context(), true); err != nil {
		return nil, err
	}

	return t.roundTrip(req, true)
}

func (t *authTransport) roundTripper() http.RoundTripper {
	if t.wrapped != nil {
		return t.wrapped
	}
	return http.DefaultTransport
}

func cloneRequest(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.Body == nil || req.Body == http.NoBody {
		return clone, nil
	}
	// Codex retry logic assumes GetBody is available so the request can be replayed after refresh.
	if req.GetBody == nil {
		return nil, errReplayBodyUnavailable
	}

	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	clone.Body = body
	return clone, nil
}

func drainAndClose(body io.ReadCloser) {
	if body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(body, defaultRoundTripLimit))
	_ = body.Close()
}

func buildAuthorizeURL(authBaseURL string, callbackURL string, pkce pkceCodes, state string) (string, error) {
	issuer, err := url.Parse(strings.TrimRight(authBaseURL, "/"))
	if err != nil {
		return "", fmt.Errorf("invalid Codex auth base URL: %w", err)
	}

	issuer.Path = strings.TrimRight(issuer.Path, "/") + "/oauth/authorize"
	query := issuer.Query()
	query.Set("response_type", "code")
	query.Set("client_id", oauthClientID)
	query.Set("redirect_uri", callbackURL)
	query.Set("scope", oauthScope)
	query.Set("code_challenge", pkce.CodeChallenge)
	query.Set("code_challenge_method", "S256")
	query.Set("id_token_add_organizations", "true")
	query.Set("codex_cli_simplified_flow", "true")
	query.Set("state", state)
	query.Set("originator", defaultOriginator)
	issuer.RawQuery = query.Encode()

	return issuer.String(), nil
}

type pkceCodes struct {
	CodeVerifier  string
	CodeChallenge string
}

func generatePKCECodes() (pkceCodes, error) {
	verifier, err := randomBase64URL(oauthVerifierBytes)
	if err != nil {
		return pkceCodes{}, err
	}

	sum := sha256.Sum256([]byte(verifier))
	return pkceCodes{
		CodeVerifier:  verifier,
		CodeChallenge: base64.RawURLEncoding.EncodeToString(sum[:]),
	}, nil
}

func randomBase64URL(size int) (string, error) {
	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("failed to generate secure random OAuth state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func tokenNeedsRefresh(jwt string, now time.Time) bool {
	expiry, err := extractExpiryFromJWT(jwt)
	if err != nil {
		return true
	}
	return now.Add(tokenRefreshLeeway).After(expiry)
}

func extractExpiryFromJWT(jwt string) (time.Time, error) {
	claims, err := parseTokenClaims(jwt)
	if err != nil {
		return time.Time{}, err
	}
	if claims.Exp == 0 {
		return time.Time{}, errors.New("JWT did not include an exp claim")
	}
	return time.Unix(claims.Exp, 0), nil
}

func extractAccountIDFromJWT(jwt string) (string, error) {
	claims, err := parseTokenClaims(jwt)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(claims.Auth.ChatGPTAccountID), nil
}

func parseTokenClaims(jwt string) (tokenClaims, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) < 2 {
		return tokenClaims{}, errors.New("invalid JWT format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return tokenClaims{}, err
	}

	var claims tokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return tokenClaims{}, err
	}

	return claims, nil
}

func extractErrorMessage(body []byte) string {
	if len(body) == 0 {
		return ""
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return strings.TrimSpace(string(body))
	}

	if errorValue, ok := payload["error"]; ok {
		switch typed := errorValue.(type) {
		case string:
			return strings.TrimSpace(typed)
		case map[string]any:
			if message, ok := typed["message"].(string); ok && strings.TrimSpace(message) != "" {
				return strings.TrimSpace(message)
			}
			if code, ok := typed["code"].(string); ok && strings.TrimSpace(code) != "" {
				return strings.TrimSpace(code)
			}
		}
	}

	if message, ok := payload["message"].(string); ok {
		return strings.TrimSpace(message)
	}
	if detail, ok := payload["detail"].(string); ok {
		return strings.TrimSpace(detail)
	}

	return strings.TrimSpace(string(body))
}

func extractErrorCode(body []byte) string {
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return ""
	}

	if code, ok := payload["code"].(string); ok {
		return strings.TrimSpace(code)
	}

	errorValue, ok := payload["error"]
	if !ok {
		return ""
	}

	switch typed := errorValue.(type) {
	case string:
		return strings.TrimSpace(typed)
	case map[string]any:
		if code, ok := typed["code"].(string); ok {
			return strings.TrimSpace(code)
		}
	}

	return ""
}

func codexClientVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if version := normalizeSemverLikeVersion(info.Main.Version); version != "" {
			return version
		}
	}

	return defaultClientVersion
}

func normalizeSemverLikeVersion(version string) string {
	version = strings.TrimSpace(version)
	version = strings.TrimPrefix(version, "v")
	if version == "" || version == "(devel)" {
		return ""
	}

	end := len(version)
	for i, r := range version {
		if (r < '0' || r > '9') && r != '.' {
			end = i
			break
		}
	}
	version = strings.Trim(version[:end], ".")
	if version == "" {
		return ""
	}

	parts := strings.Split(version, ".")
	if len(parts) < 3 {
		return ""
	}
	if slices.Contains(parts[:3], "") {
		return ""
	}

	return strings.Join(parts[:3], ".")
}

func isUsageLimitMessage(message string) bool {
	lower := strings.ToLower(strings.TrimSpace(message))
	if lower == "" {
		return false
	}

	return strings.Contains(lower, "usage limit") ||
		strings.Contains(lower, "purchase more credits") ||
		strings.Contains(lower, "upgrade to plus") ||
		strings.Contains(lower, "upgrade to pro") ||
		strings.Contains(lower, "plan and billing")
}

func openBrowser(targetURL string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", targetURL)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", targetURL)
	default:
		cmd = exec.Command("xdg-open", targetURL)
	}
	return cmd.Start()
}
