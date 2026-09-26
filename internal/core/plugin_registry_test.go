package core

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/codex"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
	"github.com/joho/godotenv"
)

func TestSaveEnvFile(t *testing.T) {
	db := fsdb.NewDb(os.TempDir())
	registry, err := NewPluginRegistry(db)
	if err != nil {
		t.Fatalf("NewPluginRegistry() error = %v", err)
	}

	err = registry.SaveEnvFile()
	if err != nil {
		t.Fatalf("SaveEnvFile() error = %v", err)
	}
}

// testVendor implements ai.Vendor for testing purposes
type testVendor struct {
	name         string
	models       []string
	configureErr error
}

func (m *testVendor) GetName() string                              { return m.name }
func (m *testVendor) GetSetupDescription() string                  { return m.name }
func (m *testVendor) IsConfigured() bool                           { return true }
func (m *testVendor) Configure() error                             { return m.configureErr }
func (m *testVendor) Setup() error                                 { return nil }
func (m *testVendor) SetupFillEnvFileContent(*bytes.Buffer)        {}
func (m *testVendor) ListModels(context.Context) ([]string, error) { return m.models, nil }
func (m *testVendor) SendStream(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error {
	return nil
}
func (m *testVendor) Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error) {
	return "", nil
}
func (m *testVendor) NeedsRawMode(string) bool { return false }

func TestGetChatter_WarnsOnAmbiguousModel(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	vendorA := &testVendor{name: "VendorA", models: []string{"shared-model"}}
	vendorB := &testVendor{name: "VendorB", models: []string{"shared-model"}}

	vm := ai.NewVendorsManager()
	vm.AddVendors(vendorA, vendorB)

	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: "VendorA"},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "shared-model"}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}

	registry := &PluginRegistry{Db: db, VendorManager: vm, Defaults: defaults}

	r, w, _ := os.Pipe()
	oldStderr := os.Stderr
	os.Stderr = w
	// Redirect log output to our pipe to capture unconditional log messages
	debuglog.SetOutput(w)
	defer func() {
		os.Stderr = oldStderr
		debuglog.SetOutput(oldStderr)
	}()

	chatter, err := registry.GetChatter("shared-model", 0, "", false, false)
	w.Close()
	warning, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("GetChatter() error = %v", err)
	}
	// Verify that one of the valid vendors was selected (don't care which one due to map iteration randomness)
	vendorName := chatter.vendor.GetName()
	if vendorName != "VendorA" && vendorName != "VendorB" {
		t.Fatalf("expected vendor VendorA or VendorB, got %s", vendorName)
	}
	if !strings.Contains(string(warning), "multiple vendors provide model shared-model") {
		t.Fatalf("expected warning about multiple vendors, got %q", string(warning))
	}
}

func TestGetChatter_AllowsExplicitCodexManualModel(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	codexVendor := &testVendor{name: "Codex", models: []string{"gpt-5.4"}}

	vm := ai.NewVendorsManager()
	vm.AddVendors(codexVendor)

	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: "Codex"},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "gpt-5.4"}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}

	registry := &PluginRegistry{Db: db, VendorManager: vm, Defaults: defaults}

	chatter, err := registry.GetChatter("gpt-5.1-codex", 0, "Codex", false, false)
	if err != nil {
		t.Fatalf("GetChatter() error = %v", err)
	}
	if chatter.vendor.GetName() != "Codex" {
		t.Fatalf("expected Codex vendor, got %s", chatter.vendor.GetName())
	}
	if chatter.model != "gpt-5.1-codex" {
		t.Fatalf("expected manual Codex model to pass through, got %s", chatter.model)
	}
}

func TestGetChatter_RejectsExplicitCodexModelFromOtherVendor(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	codexVendor := &testVendor{name: "Codex", models: []string{"gpt-5.4"}}
	anthropicVendor := &testVendor{name: "Anthropic", models: []string{"claude-3.7-sonnet"}}

	vm := ai.NewVendorsManager()
	vm.AddVendors(codexVendor, anthropicVendor)

	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: "Codex"},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "gpt-5.4"}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}

	registry := &PluginRegistry{Db: db, VendorManager: vm, Defaults: defaults}

	if _, err := registry.GetChatter("claude-3.7-sonnet", 0, "Codex", false, false); err == nil {
		t.Fatal("expected GetChatter() to reject models that only belong to another vendor")
	}
}

func TestGetChatter_ParsesVendorModelPrefix(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	ollamaVendor := &testVendor{name: "Ollama", models: []string{"some-namespace/model-name"}}

	vm := ai.NewVendorsManager()
	vm.AddVendors(ollamaVendor)

	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: "Ollama"},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "some-namespace/model-name"}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}

	registry := &PluginRegistry{Db: db, VendorManager: vm, Defaults: defaults}

	chatter, err := registry.GetChatter("ollama/some-namespace/model-name", 0, "", false, false)
	if err != nil {
		t.Fatalf("GetChatter() error = %v", err)
	}
	if chatter.vendor.GetName() != "Ollama" {
		t.Fatalf("expected Ollama vendor, got %s", chatter.vendor.GetName())
	}
	if chatter.model != "some-namespace/model-name" {
		t.Fatalf("expected model 'some-namespace/model-name', got %s", chatter.model)
	}
}

func TestGetChatter_VendorPrefixIgnoredWhenNotAVendor(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	vendorA := &testVendor{name: "VendorA", models: []string{"notavendor/model"}}

	vm := ai.NewVendorsManager()
	vm.AddVendors(vendorA)

	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: "VendorA"},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "notavendor/model"}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}

	registry := &PluginRegistry{Db: db, VendorManager: vm, Defaults: defaults}

	chatter, err := registry.GetChatter("notavendor/model", 0, "", false, false)
	if err != nil {
		t.Fatalf("GetChatter() error = %v", err)
	}
	if chatter.vendor.GetName() != "VendorA" {
		t.Fatalf("expected VendorA vendor, got %s", chatter.vendor.GetName())
	}
	if chatter.model != "notavendor/model" {
		t.Fatalf("expected model 'notavendor/model', got %s", chatter.model)
	}
}

func TestGetChatter_ReportsConfigureError(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	for _, tt := range []struct {
		name, model, vendor string
	}{
		{"empty model", "", ""},
		{"model", "gpt-5.6-sol", ""},
		{"named vendor", "gpt-5.6-sol", "Codex"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			configureErr := errors.New(i18n.T("codex_login_refresh_failed"))
			registry := newInactiveVendorRegistry(t, "Codex", "gpt-5.6-sol", &testVendor{
				name: "Codex", models: []string{"gpt-5.6-sol"}, configureErr: configureErr,
			})
			_, err := registry.GetChatter(tt.model, 0, tt.vendor, false, false)
			if !errors.Is(err, configureErr) {
				t.Fatalf("GetChatter() error = %v, want %v", err, configureErr)
			}
		})
	}
}

func TestGetChatter_ActivatesInactiveVendor(t *testing.T) {
	for _, tt := range []struct {
		name, model, vendor string
	}{
		{"named", "gpt-5.6-sol", "Codex"},
		{"default for model", "gpt-5.6-sol", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			codexVendor := &testVendor{name: "Codex", models: []string{"gpt-5.6-sol"}}
			registry := newInactiveVendorRegistry(t, "Codex", "gpt-5.6-sol", codexVendor)
			chatter, err := registry.GetChatter(tt.model, 0, tt.vendor, false, false)
			if err != nil {
				t.Fatalf("GetChatter() error = %v", err)
			}
			if chatter.vendor.GetName() != "Codex" {
				t.Fatalf("vendor = %s, want Codex", chatter.vendor.GetName())
			}
			if chatter.model != "gpt-5.6-sol" {
				t.Fatalf("model = %s, want gpt-5.6-sol", chatter.model)
			}
		})
	}
}

func TestNewPluginRegistry_CodexTokenPersist(t *testing.T) {
	dir := t.TempDir()
	db := fsdb.NewDb(dir)
	if err := db.SaveEnv("KEEP=old\n"); err != nil {
		t.Fatalf("SaveEnv() error = %v", err)
	}

	registry, err := NewPluginRegistry(db)
	if err != nil {
		t.Fatalf("NewPluginRegistry() error = %v", err)
	}

	vendor := registry.VendorsAll.FindByName("Codex")
	client, ok := vendor.(*codex.Client)
	if !ok || client == nil {
		t.Fatal("expected Codex client in VendorsAll")
	}
	if client.TokenPersist == nil {
		t.Fatal("TokenPersist is nil")
	}

	client.AccessToken.Value = "access-live"
	client.RefreshToken.Value = "refresh-live"
	client.AccountID.Value = "acct-live"
	if err := client.TokenPersist(); err != nil {
		t.Fatalf("TokenPersist() error = %v", err)
	}

	parsed, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("godotenv.Read() error = %v", err)
	}
	if parsed["KEEP"] != "old" {
		t.Fatalf("KEEP = %q, want old", parsed["KEEP"])
	}
	if parsed[client.AccessToken.EnvVariable] != "access-live" {
		t.Fatalf("access = %q, want access-live", parsed[client.AccessToken.EnvVariable])
	}
	if parsed[client.RefreshToken.EnvVariable] != "refresh-live" {
		t.Fatalf("refresh = %q, want refresh-live", parsed[client.RefreshToken.EnvVariable])
	}
	if parsed[client.AccountID.EnvVariable] != "acct-live" {
		t.Fatalf("account = %q, want acct-live", parsed[client.AccountID.EnvVariable])
	}
}

func newInactiveVendorRegistry(t *testing.T, defaultVendor, defaultModel string, vendor *testVendor) *PluginRegistry {
	t.Helper()
	db := fsdb.NewDb(t.TempDir())
	all := ai.NewVendorsManager()
	all.AddVendors(vendor)
	defaults := &tools.Defaults{
		PluginBase:         &plugins.PluginBase{},
		Vendor:             &plugins.Setting{Value: defaultVendor},
		Model:              &plugins.SetupQuestion{Setting: &plugins.Setting{Value: defaultModel}},
		ModelContextLength: &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "0"}},
	}
	return &PluginRegistry{
		Db:            db,
		VendorManager: ai.NewVendorsManager(),
		VendorsAll:    all,
		Defaults:      defaults,
	}
}
