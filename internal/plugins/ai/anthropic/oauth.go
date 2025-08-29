package anthropic

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"

	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/util"
	"golang.org/x/oauth2"
)

// OAuth configuration constants
const (
	oauthClientID    = "9d1c250a-e61b-44d9-88ed-5944d1962f5e"
	oauthAuthURL     = "https://claude.ai/oauth/authorize"
	oauthTokenURL    = "https://console.anthropic.com/v1/oauth/token"
	oauthRedirectURL = "https://console.anthropic.com/oauth/code/callback"
)

// OAuthTransport is a custom HTTP transport that adds OAuth Bearer token and beta header
type OAuthTransport struct {
	client *Client
	base   http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface
func (t *OAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original
	newReq := req.Clone(req.Context())

	// Get current token (may refresh if needed)
	token, err := t.getValidToken(authTokenIdentifier)
	if err != nil {
		return nil, fmt.Errorf("failed to get valid OAuth token: %w", err)
	}

	// Add OAuth Bearer token
	newReq.Header.Set("Authorization", "Bearer "+token)

	// Add the anthropic-beta header for OAuth, preserving existing betas
	existing := newReq.Header.Get("anthropic-beta")
	beta := "oauth-2025-04-20"
	if existing != "" {
		beta = existing + "," + beta
	}
	newReq.Header.Set("anthropic-beta", beta)

	// Set User-Agent to match AI SDK exactly
	newReq.Header.Set("User-Agent", "ai-sdk/anthropic")

	// Remove x-api-key header if present (OAuth doesn't use it)
	newReq.Header.Del("x-api-key")

	return t.base.RoundTrip(newReq)
}

// getValidToken returns a valid access token, refreshing if necessary
func (t *OAuthTransport) getValidToken(tokenIdentifier string) (string, error) {
	storage, err := util.NewOAuthStorage()
	if err != nil {
		return "", fmt.Errorf("failed to create OAuth storage: %w", err)
	}

	// Load stored token
	token, err := storage.LoadToken(tokenIdentifier)
	if err != nil {
		return "", fmt.Errorf("failed to load stored token: %w", err)
	}
	// If no token exists, run OAuth flow
	if token == nil {
		debuglog.Log("No OAuth token found, initiating authentication...\n")
		newAccessToken, err := RunOAuthFlow(tokenIdentifier)
		if err != nil {
			return "", fmt.Errorf("failed to authenticate: %w", err)
		}
		return newAccessToken, nil
	}

	// Check if token needs refresh (5 minute buffer)
	if token.IsExpired(5) {
		debuglog.Log("OAuth token expired, refreshing...\n")
		newAccessToken, err := RefreshToken(tokenIdentifier)
		if err != nil {
			// If refresh fails, try re-authentication
			debuglog.Log("Token refresh failed, re-authenticating...\n")
			newAccessToken, err = RunOAuthFlow(tokenIdentifier)
			if err != nil {
				return "", fmt.Errorf("failed to refresh or re-authenticate: %w", err)
			}
		}

		return newAccessToken, nil
	}

	return token.AccessToken, nil
}

// NewOAuthTransport creates a new OAuth transport for the given client
func NewOAuthTransport(client *Client, base http.RoundTripper) *OAuthTransport {
	return &OAuthTransport{
		client: client,
		base:   base,
	}
}

// generatePKCE generates PKCE code verifier and challenge
func generatePKCE() (verifier, challenge string, err error) {
	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return
	}
	verifier = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(sum[:])
	return
}

// openBrowser attempts to open the given URL in the default browser
func openBrowser(url string) {
	commands := [][]string{{"xdg-open", url}, {"open", url}, {"cmd", "/c", "start", url}}
	for _, cmd := range commands {
		if exec.Command(cmd[0], cmd[1:]...).Start() == nil {
			return
		}
	}
}

// RunOAuthFlow executes the complete OAuth authorization flow
func RunOAuthFlow(tokenIdentifier string) (token string, err error) {
	// First check if we have an existing token that can be refreshed
	storage, err := util.NewOAuthStorage()
	if err == nil {
		existingToken, err := storage.LoadToken(tokenIdentifier)
		if err == nil && existingToken != nil {
			// If token exists but is expired, try refreshing first
			if existingToken.IsExpired(5) {
				debuglog.Log("Found expired OAuth token, attempting refresh...\n")
				refreshedToken, refreshErr := RefreshToken(tokenIdentifier)
				if refreshErr == nil {
					debuglog.Log("Token refresh successful\n")
					return refreshedToken, nil
				}
				debuglog.Log("Token refresh failed (%v), proceeding with full OAuth flow...\n", refreshErr)
			} else {
				// Token exists and is still valid
				return existingToken.AccessToken, nil
			}
		}
	}

	verifier, challenge, err := generatePKCE()
	if err != nil {
		return
	}

	cfg := oauth2.Config{
		ClientID:    oauthClientID,
		Endpoint:    oauth2.Endpoint{AuthURL: oauthAuthURL, TokenURL: oauthTokenURL},
		RedirectURL: oauthRedirectURL,
		Scopes:      []string{"org:create_api_key", "user:profile", "user:inference"},
	}

	authURL := cfg.AuthCodeURL(verifier,
		oauth2.SetAuthURLParam("code_challenge", challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code", "true"),
		oauth2.SetAuthURLParam("state", verifier),
	)

	debuglog.Log("Open the following URL in your browser. Fabric would like to authorize:\n")
	debuglog.Log("%s\n", authURL)
	openBrowser(authURL)
	debuglog.Log("Paste the authorization code here: ")
	var code string
	fmt.Scanln(&code)
	parts := strings.SplitN(code, "#", 2)
	state := verifier
	if len(parts) == 2 {
		state = parts[1]
	}

	// Manual token exchange to match opencode implementation
	tokenReq := map[string]string{
		"code":          parts[0],
		"state":         state,
		"grant_type":    "authorization_code",
		"client_id":     oauthClientID,
		"redirect_uri":  oauthRedirectURL,
		"code_verifier": verifier,
	}

	token, err = exchangeToken(tokenIdentifier, tokenReq)
	return
}

// exchangeToken exchanges authorization code for access token
func exchangeToken(tokenIdentifier string, params map[string]string) (token string, err error) {
	reqBody, err := json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := http.Post(oauthTokenURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("token exchange failed: %s - %s", resp.Status, string(body))
		return
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	// Save the complete token information
	storage, err := util.NewOAuthStorage()
	if err != nil {
		return result.AccessToken, fmt.Errorf("failed to create OAuth storage: %w", err)
	}

	oauthToken := &util.OAuthToken{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresAt:    time.Now().Unix() + int64(result.ExpiresIn),
		TokenType:    result.TokenType,
		Scope:        result.Scope,
	}

	if err = storage.SaveToken(tokenIdentifier, oauthToken); err != nil {
		return result.AccessToken, fmt.Errorf("failed to save OAuth token: %w", err)
	}

	token = result.AccessToken
	return
}

// RefreshToken refreshes an expired OAuth token using the refresh token
func RefreshToken(tokenIdentifier string) (string, error) {
	storage, err := util.NewOAuthStorage()
	if err != nil {
		return "", fmt.Errorf("failed to create OAuth storage: %w", err)
	}

	// Load existing token
	token, err := storage.LoadToken(tokenIdentifier)
	if err != nil {
		return "", fmt.Errorf("failed to load stored token: %w", err)
	}
	if token == nil || token.RefreshToken == "" {
		return "", fmt.Errorf("no refresh token available")
	}

	// Prepare refresh request
	refreshReq := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     oauthClientID,
	}

	reqBody, err := json.Marshal(refreshReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal refresh request: %w", err)
	}

	// Make refresh request
	resp, err := http.Post(oauthTokenURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("refresh request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token refresh failed: %s - %s", resp.Status, string(body))
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse refresh response: %w", err)
	}

	// Update stored token
	newToken := &util.OAuthToken{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresAt:    time.Now().Unix() + int64(result.ExpiresIn),
		TokenType:    result.TokenType,
		Scope:        result.Scope,
	}

	// Use existing refresh token if new one not provided
	if newToken.RefreshToken == "" {
		newToken.RefreshToken = token.RefreshToken
	}

	if err = storage.SaveToken(tokenIdentifier, newToken); err != nil {
		return "", fmt.Errorf("failed to save refreshed token: %w", err)
	}

	return result.AccessToken, nil
}
