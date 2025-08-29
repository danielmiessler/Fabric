package core

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
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
	name   string
	models []string
}

func (m *testVendor) GetName() string                       { return m.name }
func (m *testVendor) GetSetupDescription() string           { return m.name }
func (m *testVendor) IsConfigured() bool                    { return true }
func (m *testVendor) Configure() error                      { return nil }
func (m *testVendor) Setup() error                          { return nil }
func (m *testVendor) SetupFillEnvFileContent(*bytes.Buffer) {}
func (m *testVendor) ListModels() ([]string, error)         { return m.models, nil }
func (m *testVendor) SendStream([]*chat.ChatCompletionMessage, *domain.ChatOptions, chan string) error {
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

	chatter, err := registry.GetChatter("shared-model", 0, "", "", false, false)
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
