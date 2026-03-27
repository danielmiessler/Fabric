package codex

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"runtime/debug"
	"slices"
	"strings"
	"time"
)

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
		return time.Time{}, errors.New("jwt did not include an exp claim")
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
		return tokenClaims{}, errors.New("invalid jwt format")
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
