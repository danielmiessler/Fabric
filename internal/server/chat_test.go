package restapi

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
	"github.com/gin-gonic/gin"
)

func TestBuildPromptChatRequest_PreservesStrategyAndUserInput(t *testing.T) {
	prompt := PromptRequest{
		UserInput:    "user input",
		Vendor:       "TestVendor",
		Model:        "test-model",
		ContextName:  "ctx",
		PatternName:  "pattern",
		StrategyName: "strategy",
		SessionName:  "session",
		Variables: map[string]string{
			"topic": "pipelines",
		},
	}

	request := buildPromptChatRequest(prompt, "en")

	if request.Message == nil {
		t.Fatal("expected request message to be set")
	}
	if request.Message.Content != "user input" {
		t.Fatalf("expected user input to stay unchanged, got %q", request.Message.Content)
	}
	if request.StrategyName != "strategy" {
		t.Fatalf("expected strategy name to be preserved, got %q", request.StrategyName)
	}
	if request.PatternName != "pattern" {
		t.Fatalf("expected pattern name to be preserved, got %q", request.PatternName)
	}
	if request.ContextName != "ctx" {
		t.Fatalf("expected context name to be preserved, got %q", request.ContextName)
	}
	if request.SessionName != "session" {
		t.Fatalf("expected session name to be preserved, got %q", request.SessionName)
	}
	if request.Language != "en" {
		t.Fatalf("expected language to be preserved, got %q", request.Language)
	}
	if got := request.PatternVariables["topic"]; got != "pipelines" {
		t.Fatalf("expected variables to be preserved, got %q", got)
	}
}

type serverTestVendor struct {
	name   string
	models []string
}

func (m *serverTestVendor) GetName() string                              { return m.name }
func (m *serverTestVendor) GetSetupDescription() string                  { return m.name }
func (m *serverTestVendor) IsConfigured() bool                           { return true }
func (m *serverTestVendor) Configure() error                             { return nil }
func (m *serverTestVendor) Setup() error                                 { return nil }
func (m *serverTestVendor) SetupFillEnvFileContent(*bytes.Buffer)        {}
func (m *serverTestVendor) ListModels(context.Context) ([]string, error) { return m.models, nil }
func (m *serverTestVendor) SendStream(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error {
	return nil
}
func (m *serverTestVendor) Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error) {
	return "", nil
}
func (m *serverTestVendor) NeedsRawMode(string) bool { return false }

type closeNotifierRecorder struct {
	*httptest.ResponseRecorder
	closeCh chan bool
}

func (r *closeNotifierRecorder) CloseNotify() <-chan bool {
	return r.closeCh
}

func TestHandleChat_InvalidStrategyEmitsErrorAndSkipsComplete(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	gin.SetMode(gin.TestMode)

	db := fsdb.NewDb(t.TempDir())
	vendor := &serverTestVendor{name: "TestVendor", models: []string{"test-model"}}
	vm := ai.NewVendorsManager()
	vm.AddVendors(vendor)

	registry := &core.PluginRegistry{
		Db:            db,
		VendorManager: vm,
		Defaults: &tools.Defaults{
			PluginBase:         &plugins.PluginBase{},
			Vendor:             &plugins.Setting{Value: "TestVendor"},
			Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "test-model"}},
			ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
		},
	}

	router := gin.New()
	NewChatHandler(router, registry, db)

	body := `{"prompts":[{"userInput":"hello","vendor":"TestVendor","model":"test-model","strategyName":"missing-strategy"}]}`
	req := httptest.NewRequest(http.MethodPost, "/chat", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := &closeNotifierRecorder{
		ResponseRecorder: httptest.NewRecorder(),
		closeCh:          make(chan bool, 1),
	}

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %q", resp.Code, resp.Body.String())
	}

	responseBody := resp.Body.String()
	if !strings.Contains(responseBody, `"type":"error"`) {
		t.Fatalf("expected SSE error event for missing strategy, got body %q", responseBody)
	}
	if strings.Contains(responseBody, `"type":"complete"`) {
		t.Fatalf("expected no complete SSE event for missing strategy, got body %q", responseBody)
	}
}
