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
	vendors.AddGroupItems("vendor1", []string{"model1", "model2"}...)
	vendor := vendors.FindGroupsByItemFirst("model1")
	if vendor != "vendor1" {
		t.Fatalf("FindVendorsByModelFirst() = %v, want %v", vendor, "vendor1")
	}
}

func TestFindVendorsByModel(t *testing.T) {
	vendors := NewVendorsModels()
	vendors.AddGroupItems("vendor1", []string{"model1", "model2"}...)
	foundVendors := vendors.FindGroupsByItem("model1")
	if len(foundVendors) != 1 || foundVendors[0] != "vendor1" {
		t.Fatalf("FindVendorsByModel() = %v, want %v", foundVendors, []string{"vendor1"})
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
