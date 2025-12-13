package ai

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewVendorsModels(t *testing.T) {
	vendors := NewVendorsModels()
	if vendors == nil {
		t.Fatalf("NewVendorsModels() returned nil")
	}
	if len(vendors.GroupsItems) != 0 {
		t.Fatalf("NewVendorsModels() returned non-empty VendorsModels map")
	}
}

func TestFindVendorsByModelFirst(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("Vendor1", []string{"Model1", "model2"}...)
	vendor := vendors.FindGroupsByItemFirst("model1")
	if vendor != "Vendor1" {
		t.Fatalf("FindVendorsByModelFirst() = %v, want %v", vendor, "Vendor1")
	}
}

func TestFindVendorsByModel(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("Vendor1", []string{"Model1", "model2"}...)
	foundVendors := vendors.FindGroupsByItem("MODEL1")
	if len(foundVendors) != 1 || foundVendors[0] != "Vendor1" {
		t.Fatalf("FindVendorsByModel() = %v, want %v", foundVendors, []string{"Vendor1"})
	}
}

func TestPrintWithVendorMarksDefault(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("vendor1", []string{"model1"}...)
	vendors.AddGroupItems("vendor2", []string{"model2"}...)

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	vendors.PrintWithVendor(false, "vendor2", "model2")

	w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "      *\t[2]\tvendor2|model2") {
		t.Fatalf("default model not marked: %s", out)
	}
}

func TestFilterByVendorCaseInsensitive(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("vendor1", []string{"model1"}...)
	vendors.AddGroupItems("vendor2", []string{"model2"}...)

	filtered := vendors.FilterByVendor("VENDOR2")

	if len(filtered.GroupsItems) != 1 {
		t.Fatalf("expected 1 vendor group, got %d", len(filtered.GroupsItems))
	}

	if filtered.GroupsItems[0].Group != "vendor2" {
		t.Fatalf("expected vendor2, got %s", filtered.GroupsItems[0].Group)
	}

	if len(filtered.GroupsItems[0].Items) != 1 || filtered.GroupsItems[0].Items[0] != "model2" {
		t.Fatalf("unexpected models for vendor2: %v", filtered.GroupsItems[0].Items)
	}
}

func TestFindModelNameCaseInsensitive(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("OpenAI", []string{"gpt-4o", "gpt-5"}...)
	vendors.AddGroupItems("Anthropic", []string{"claude-3-opus"}...)

	tests := []struct {
		name          string
		query         string
		expectedModel string
	}{
		{"exact match lowercase", "gpt-4o", "gpt-4o"},
		{"uppercase query", "GPT-4O", "gpt-4o"},
		{"mixed case query", "GpT-5", "gpt-5"},
		{"exact match with hyphens", "claude-3-opus", "claude-3-opus"},
		{"uppercase with hyphens", "CLAUDE-3-OPUS", "claude-3-opus"},
		{"non-existent model", "gpt-999", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := vendors.FindModelNameCaseInsensitive(tt.query)
			if result != tt.expectedModel {
				t.Errorf("FindModelNameCaseInsensitive(%q) = %q, want %q", tt.query, result, tt.expectedModel)
			}
		})
	}
}
