package ai

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

type testVendor struct {
	name  string
	calls atomic.Int32
}

func (v *testVendor) GetName() string                       { return v.name }
func (v *testVendor) GetSetupDescription() string           { return "" }
func (v *testVendor) IsConfigured() bool                    { return true }
func (v *testVendor) Configure() error                      { return nil }
func (v *testVendor) Setup() error                          { return nil }
func (v *testVendor) SetupFillEnvFileContent(*bytes.Buffer) {}
func (v *testVendor) ListModels(context.Context) ([]string, error) {
	v.calls.Add(1)
	return []string{v.name + "-model"}, nil
}
func (v *testVendor) SendStream(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error {
	return nil
}
func (v *testVendor) Send(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error) {
	return "", nil
}
func (v *testVendor) NeedsRawMode(string) bool { return false }

// TestAddVendorsInvalidatesCachedModels validates that registering
// a new vendor after warming the cache includes the new vendor's models.
func TestAddVendorsInvalidatesCachedModels(t *testing.T) {
	manager := NewVendorsManager()
	first := &testVendor{name: "alpha"}
	second := &testVendor{name: "beta"}

	manager.AddVendors(first)
	if _, err := manager.GetModels(); err != nil {
		t.Fatal(err)
	}

	manager.AddVendors(second)
	models, err := manager.GetModels()
	if err != nil {
		t.Fatal(err)
	}
	if got := models.FindModelNameCaseInsensitive("beta-model"); got != "beta-model" {
		t.Errorf("new registered vendor missing from warmed cache: got=%q", got)
	}
}

type failingVendor struct {
	testVendor
}

func (v *failingVendor) ListModels(context.Context) ([]string, error) {
	if v.calls.Add(1) == 1 {
		return nil, errors.New("first list fails")
	}
	return []string{v.name + "-model"}, nil
}

// TestGetModelsDoesNotCacheTotalFailure validates that a discovery where all
// vendors fail returns an error and that the next call tries the vendors again.
func TestGetModelsDoesNotCacheTotalFailure(t *testing.T) {
	manager := NewVendorsManager()
	vendor := &failingVendor{testVendor{name: "flaky"}}
	manager.AddVendors(vendor)

	if _, err := manager.GetModels(); err == nil {
		t.Fatal("expected an error when all vendors fail")
	}
	models, err := manager.GetModels()
	if err != nil {
		t.Fatal(err)
	}
	if got := models.FindModelNameCaseInsensitive("flaky-model"); got != "flaky-model" {
		t.Errorf("failed discovery stayed cached: got=%q calls=%d", got, vendor.calls.Load())
	}
}

// TestGetModelsKeepsPartialAndEmptyCatalogs validates that one healthy vendor,
// also with no models, is a successful discovery.
func TestGetModelsKeepsPartialAndEmptyCatalogs(t *testing.T) {
	manager := NewVendorsManager()
	manager.AddVendors(&failingVendor{testVendor{name: "flaky"}}, &emptyVendor{testVendor{name: "empty"}})

	if _, err := manager.GetModels(); err != nil {
		t.Fatalf("partial success returned error: %v", err)
	}
}

type emptyVendor struct {
	testVendor
}

func (v *emptyVendor) ListModels(context.Context) ([]string, error) { return nil, nil }
