package restapi

import (
	"encoding/json"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai"
)

func TestBuildVendorsMapGivesEmptyArrayForVendorWithoutModels(t *testing.T) {
	vendorsModels := ai.NewVendorsModels()
	vendorsModels.AddGroupItems("Anthropic", "claude-opus-5", "claude-sonnet-5")
	// A variadic call with no items makes a nil slice, which is what a vendor in
	// the configuration that serves no models produces.
	vendorsModels.AddGroupItems("Ollama")

	vendors := buildVendorsMap(vendorsModels)

	if got := vendors["Ollama"]; got == nil {
		t.Error("want an empty slice for a vendor with no models, got nil")
	}
	if got := len(vendors["Ollama"]); got != 0 {
		t.Errorf("want 0 models for Ollama, got %d", got)
	}
	if got := len(vendors["Anthropic"]); got != 2 {
		t.Errorf("want 2 models for Anthropic, got %d", got)
	}

	// The client reads vendors[name] as an array. Confirm the JSON holds an
	// array and not null, which is what stopped the web UI from listing models.
	encoded, err := json.Marshal(vendors)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded map[string][]string
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded["Ollama"] == nil {
		t.Errorf("want an array for Ollama in JSON, got null: %s", encoded)
	}
}

func TestBuildVendorsMapWithNoVendors(t *testing.T) {
	if got := len(buildVendorsMap(ai.NewVendorsModels())); got != 0 {
		t.Errorf("want an empty map, got %d entries", got)
	}
}
