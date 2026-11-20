package core

import (
	"errors"
	"testing"
)

func TestGetChatterRecoversPrefixedVendor(t *testing.T) {
	cases := []struct {
		name         string
		query        string
		vendor       string
		warm         bool
		reconfigure  bool
		keepError    bool
		wantVendor   string
		wantModel    string
		wantInactive bool
	}{
		{"explicit recovered control", "bEtA-mOdEl", "bEtA", false, false, false, "Beta", "beta-model", false},
		{"active prefix control", "bEtA/bEtA-mOdEl", "", false, true, false, "Beta", "beta-model", false},
		{"recovered prefix cold", "bEtA/bEtA-mOdEl", "", false, false, false, "Beta", "beta-model", false},
		{"recovered prefix warm", "bEtA/bEtA-mOdEl", "", true, false, false, "Beta", "beta-model", false},
		{"prefix configure error", "bEtA/beta-model", "", false, false, true, "", "", true},
		{"known vendor opaque ID control", "Beta/opaque-model", "", false, false, false, "Alpha", "Beta/opaque-model", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			failed := errors.New("temporary vendor configuration failure")
			alpha := &testVendor{name: "Alpha", models: []string{"alpha-model", "Beta/opaque-model"}}
			beta := &testVendor{name: "Beta", models: []string{"beta-model"}, configureErr: failed}
			registry := newInactiveVendorRegistry(t, "Alpha", "alpha-model", alpha)
			registry.VendorsAll.AddVendors(beta)
			registry.ConfigureVendors()
			if registry.VendorManager.FindByName("Beta") != nil {
				t.Fatal("failed vendor was unexpectedly activated")
			}
			if !tc.keepError {
				beta.configureErr = nil
			}
			if tc.reconfigure {
				registry.ConfigureVendors()
			}
			if tc.warm {
				if _, err := registry.VendorManager.GetModels(); err != nil {
					t.Fatal(err)
				}
			}

			chatter, err := registry.GetChatter(tc.query, 0, tc.vendor, false, false)
			if tc.keepError {
				if !errors.Is(err, failed) {
					t.Fatalf("error = %v, want configuration failure", err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if got := chatter.vendor.GetName(); got != tc.wantVendor {
					t.Errorf("vendor = %q, want %q", got, tc.wantVendor)
				}
				if chatter.model != tc.wantModel {
					t.Errorf("model = %q, want %q", chatter.model, tc.wantModel)
				}
			}
			if tc.wantInactive && registry.VendorManager.FindByName("Beta") != nil {
				t.Fatal("unselected or failed vendor was unexpectedly activated")
			}
		})
	}
}
