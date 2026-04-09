package codex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	plugins "github.com/danielmiessler/fabric/internal/plugins"
)

var errReplayBodyUnavailable = errors.New(i18n.T("codex_replay_body_unavailable"))

type authTransport struct {
	client  *Client
	wrapped http.RoundTripper
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
		return oauthTokens{}, fmt.Errorf(i18n.T("codex_refresh_login_failed"), err)
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
		return oauthTokens{}, fmt.Errorf(i18n.T("codex_decode_refresh_response_failed"), err)
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
