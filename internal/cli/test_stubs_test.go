package cli

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

type cliStubVendor struct {
	name   string
	models []string
}

func (v *cliStubVendor) GetName() string                       { return v.name }
func (v *cliStubVendor) GetSetupDescription() string           { return v.name }
func (v *cliStubVendor) IsConfigured() bool                    { return true }
func (v *cliStubVendor) Configure() error                      { return nil }
func (v *cliStubVendor) Setup() error                          { return nil }
func (v *cliStubVendor) SetupFillEnvFileContent(*bytes.Buffer) {}
func (v *cliStubVendor) ListModels() ([]string, error)         { return v.models, nil }
func (v *cliStubVendor) SendStream([]*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error {
	return nil
}
func (v *cliStubVendor) Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error) {
	return "", nil
}
func (v *cliStubVendor) NeedsRawMode(string) bool { return false }

type cliStubTranscriberVendor struct {
	cliStubVendor
	transcription string
	transcribeErr error
	lastFile      string
	lastModel     string
	lastSplit     bool
	lastNilCtx    bool
}

func (v *cliStubTranscriberVendor) TranscribeFile(ctx context.Context, filePath, model string, split bool) (string, error) {
	v.lastNilCtx = ctx == nil
	v.lastFile = filePath
	v.lastModel = model
	v.lastSplit = split
	return v.transcription, v.transcribeErr
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func installHTTPTransport(t *testing.T, fn roundTripFunc) {
	t.Helper()

	original := http.DefaultTransport
	http.DefaultTransport = fn
	t.Cleanup(func() {
		http.DefaultTransport = original
	})
}
