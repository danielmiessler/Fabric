package core

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
)

type lazyCatalogTestVendor struct {
	*mockVendor
	name   string
	models []string
	calls  atomic.Int32
}

func (v *lazyCatalogTestVendor) GetName() string { return v.name }
func (v *lazyCatalogTestVendor) ListModels(context.Context) ([]string, error) {
	v.calls.Add(1)
	return v.models, nil
}

func TestGetChatterActivatesVendorAfterModelCacheWarmup(t *testing.T) {
	first := &lazyCatalogTestVendor{mockVendor: &mockVendor{}, name: "alpha", models: []string{"alpha-model"}}
	second := &lazyCatalogTestVendor{mockVendor: &mockVendor{}, name: "beta", models: []string{"beta-model"}}
	active, available := ai.NewVendorsManager(), ai.NewVendorsManager()
	active.AddVendors(first)
	available.AddVendors(first, second)
	if _, err := active.GetModels(); err != nil {
		t.Fatal(err)
	}
	registry := &PluginRegistry{
		Db: fsdb.NewDb(t.TempDir()), VendorManager: active, VendorsAll: available,
		Defaults: tools.NeeDefaults(active.GetModels),
	}
	if _, err := registry.GetChatter("beta-model", 0, "beta", true, false); err != nil {
		t.Errorf("healthy newly activated vendor rejected after cache warmup: %v; listCalls=%d", err, second.calls.Load())
	}
	// Explicit refresh is the control; the same configured fake provider works.
	active.Clear()
	active.AddVendors(first, second)
	if _, err := registry.GetChatter("beta-model", 0, "beta", true, false); err != nil {
		t.Fatalf("cold-cache control also failed: %v", err)
	}
}
