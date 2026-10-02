package core

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools"
)

func TestGetChatterUsesSelectedVendorsModelSpelling(t *testing.T) {
	tests := []struct {
		name, query, vendor, wantVendor, wantModel string
		groups                                     []struct{ vendor, model string }
	}{
		{"explicit vendor after other group", "sHaReD-mOdEl", "bEtA", "Beta", "shared-model",
			[]struct{ vendor, model string }{{"Alpha", "SHARED-MODEL"}, {"Beta", "shared-model"}}},
		{"vendor prefix after other group", "bEtA/sHaReD-mOdEl", "", "Beta", "shared-model",
			[]struct{ vendor, model string }{{"Alpha", "SHARED-MODEL"}, {"Beta", "shared-model"}}},
		{"explicit vendor first", "sHaReD-mOdEl", "Beta", "Beta", "shared-model",
			[]struct{ vendor, model string }{{"Beta", "shared-model"}, {"Alpha", "SHARED-MODEL"}}},
		{"selected uppercase spelling", "sHaReD-mOdEl", "Alpha", "Alpha", "SHARED-MODEL",
			[]struct{ vendor, model string }{{"Beta", "shared-model"}, {"Alpha", "SHARED-MODEL"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := ai.NewVendorsManager()
			manager.AddVendors(
				&testVendor{name: "Alpha", models: []string{"SHARED-MODEL"}},
				&testVendor{name: "Beta", models: []string{"shared-model"}},
			)
			// Seed a completed catalog to test selection independently of discovery order.
			catalog := ai.NewVendorsModels()
			for _, group := range tt.groups {
				catalog.AddGroupItems(group.vendor, group.model)
			}
			manager.Models = catalog
			registry := &PluginRegistry{Db: fsdb.NewDb(t.TempDir()), VendorManager: manager, Defaults: tools.NeeDefaults(manager.GetModels)}
			chatter, err := registry.GetChatter(tt.query, 0, tt.vendor, false, false)
			if err != nil {
				t.Fatal(err)
			}
			if got := chatter.vendor.GetName(); got != tt.wantVendor {
				t.Errorf("vendor = %q, want %q", got, tt.wantVendor)
			}
			if chatter.model != tt.wantModel {
				t.Errorf("model = %q, want selected vendor spelling %q", chatter.model, tt.wantModel)
			}
		})
	}
}
