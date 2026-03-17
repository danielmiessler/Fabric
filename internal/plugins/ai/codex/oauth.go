package codex

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

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

type pkceCodes struct {
	CodeVerifier  string
	CodeChallenge string
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

	if !oauthStatesMatch(expectedState, r.URL.Query().Get("state")) {
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

func oauthStatesMatch(expected string, actual string) bool {
	if len(expected) != len(actual) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
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
