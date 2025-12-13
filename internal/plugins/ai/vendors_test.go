package ai

import (
	"bytes"
	"context"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

type stubVendor struct {
	name string
}

func (v *stubVendor) GetName() string                       { return v.name }
func (v *stubVendor) GetSetupDescription() string           { return "" }
func (v *stubVendor) IsConfigured() bool                    { return true }
func (v *stubVendor) Configure() error                      { return nil }
func (v *stubVendor) Setup() error                          { return nil }
func (v *stubVendor) SetupFillEnvFileContent(*bytes.Buffer) {}
func (v *stubVendor) ListModels() ([]string, error)         { return nil, nil }
func (v *stubVendor) SendStream([]*chat.ChatCompletionMessage, *domain.ChatOptions, chan string) error {
	return nil
}
func (v *stubVendor) Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error) {
	return "", nil
}
func (v *stubVendor) NeedsRawMode(string) bool { return false }

func TestVendorsManagerFindByNameCaseInsensitive(t *testing.T) {
	manager := NewVendorsManager()
	vendor := &stubVendor{name: "OpenAI"}

	manager.AddVendors(vendor)

	if got := manager.FindByName("openai"); got != vendor {
		t.Fatalf("FindByName lowercase = %v, want %v", got, vendor)
	}

	if got := manager.FindByName("OPENAI"); got != vendor {
		t.Fatalf("FindByName uppercase = %v, want %v", got, vendor)
	}

	if got := manager.FindByName("OpenAI"); got != vendor {
		t.Fatalf("FindByName mixed case = %v, want %v", got, vendor)
	}
}

func TestVendorsManagerSetupVendorToCaseInsensitive(t *testing.T) {
	manager := NewVendorsManager()
	vendor := &stubVendor{name: "OpenAI"}

	configured := map[string]Vendor{}
	manager.setupVendorTo(vendor, configured)

	// Verify vendor is stored with lowercase key
	if _, ok := configured["openai"]; !ok {
		t.Fatalf("setupVendorTo should store vendor using lowercase key")
	}

	// Verify original case key is not used
	if _, ok := configured["OpenAI"]; ok {
		t.Fatalf("setupVendorTo should not store vendor using original case key")
	}
}
