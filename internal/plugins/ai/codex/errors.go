package codex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	openaiapi "github.com/openai/openai-go"
)

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
