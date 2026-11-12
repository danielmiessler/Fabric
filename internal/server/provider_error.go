package restapi

import (
	"errors"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/ai/ollama"
	"github.com/openai/openai-go/v3"
	"google.golang.org/genai"
)

// publicError returns the text of err that the server sends to a client or
// writes to its log. The text of a provider error can repeat private
// content, for example the pattern, the context or the session. Thus, for a
// provider error, publicError gives only the HTTP status and the error code.
// When the server starts with --debug, publicError gives the full text.
func publicError(err error) string {
	if _, ok := errors.AsType[*domain.ProviderError](err); !ok || debuglog.GetLevel() >= debuglog.Basic {
		return err.Error()
	}
	status, code := providerStatus(err)
	switch {
	case status != 0 && code != "":
		return fmt.Sprintf("provider error: status %d, code %s", status, code)
	case status != 0:
		return fmt.Sprintf("provider error: status %d", status)
	case code != "":
		return fmt.Sprintf("provider error: code %s", code)
	}
	return "provider error"
}

// providerStatus finds the HTTP status and the error code of a provider
// error. It returns zero values when it does not know the error type.
func providerStatus(err error) (status int, code string) {
	if e, ok := errors.AsType[*openai.Error](err); ok {
		return e.StatusCode, e.Code
	}
	if e, ok := errors.AsType[*anthropic.Error](err); ok {
		return e.StatusCode, ""
	}
	if e, ok := errors.AsType[genai.APIError](err); ok {
		return e.Code, e.Status
	}
	if e, ok := errors.AsType[*ollama.StatusError](err); ok {
		return e.StatusCode, ""
	}
	// The AWS SDK errors, which Bedrock returns, have these methods.
	if e, ok := errors.AsType[interface {
		error
		HTTPStatusCode() int
	}](err); ok {
		status = e.HTTPStatusCode()
	}
	if e, ok := errors.AsType[interface {
		error
		ErrorCode() string
	}](err); ok {
		code = e.ErrorCode()
	}
	return
}
