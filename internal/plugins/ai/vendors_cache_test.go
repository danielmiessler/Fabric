package ai

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type catalogTestVendor struct {
	*stubVendor
	models []string
	err    error
	calls  atomic.Int32
}

func (v *catalogTestVendor) ListModels(context.Context) ([]string, error) {
	v.calls.Add(1)
	return v.models, v.err
}

func TestVendorsManagerDoesNotMutateBorrowedCatalogSlice(t *testing.T) {
	manager := NewVendorsManager()
	vendor := &catalogTestVendor{stubVendor: &stubVendor{name: "borrowed"}, models: []string{"z-model", "a-model"}}
	manager.AddVendors(vendor)
	if _, err := manager.GetModels(); err != nil {
		t.Fatal(err)
	}
	if vendor.models[0] != "z-model" || vendor.models[1] != "a-model" {
		t.Fatal("manager sorted the provider-owned slice in place")
	}
}

func TestVendorsManagerGetModelsKeepsHealthyCatalogAfterOtherVendorFailure(t *testing.T) {
	manager := NewVendorsManager()
	good := &catalogTestVendor{stubVendor: &stubVendor{name: "alpha"}, models: []string{"alpha-model"}}
	bad := &catalogTestVendor{stubVendor: &stubVendor{name: "beta"}, err: errors.New("synthetic model-list failure")}
	manager.AddVendors(good, bad)
	models, err := manager.GetModels()
	if err != nil || models.FindModelNameCaseInsensitive("alpha-model") != "alpha-model" {
		t.Fatalf("healthy vendor models lost: err=%v", err)
	}
}

func TestVendorsManagerAddVendorsInvalidatesCachedModels(t *testing.T) {
	manager := NewVendorsManager()
	first := &catalogTestVendor{stubVendor: &stubVendor{name: "alpha"}, models: []string{"alpha-model"}}
	second := &catalogTestVendor{stubVendor: &stubVendor{name: "beta"}, models: []string{"beta-model"}}
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
		t.Errorf("new registered vendor missing from warmed cache: model=%q listCalls=%d", got, second.calls.Load())
	}
	// A cold read is the control: the vendor itself is healthy and has the model.
	manager.Clear()
	manager.AddVendors(first, second)
	models, err = manager.GetModels()
	if err != nil || models.FindModelNameCaseInsensitive("beta-model") != "beta-model" {
		t.Fatalf("cold-read control did not find new model: err=%v", err)
	}
}

type blockedCatalogTestVendor struct {
	*stubVendor
	entered chan struct{}
	release chan struct{}
	once    sync.Once
	calls   atomic.Int32
}

func (v *blockedCatalogTestVendor) ListModels(context.Context) ([]string, error) {
	v.calls.Add(1)
	v.once.Do(func() { close(v.entered) })
	<-v.release
	return []string{"slow-model"}, nil
}

func TestVendorsManagerGetModelsWaitsForPendingCatalog(t *testing.T) {
	manager := NewVendorsManager()
	vendor := &blockedCatalogTestVendor{stubVendor: &stubVendor{name: "slow"}, entered: make(chan struct{}), release: make(chan struct{})}
	manager.AddVendors(vendor)
	var releaseOnce sync.Once
	unblock := func() { releaseOnce.Do(func() { close(vendor.release) }) }
	defer unblock()
	type result struct {
		models *VendorsModels
		err    error
	}
	first, second := make(chan result, 1), make(chan result, 1)
	go func() { m, err := manager.GetModels(); first <- result{m, err} }()
	select {
	case <-vendor.entered:
	case <-time.After(time.Second):
		t.Fatal("model discovery did not start")
	}
	go func() { m, err := manager.GetModels(); second <- result{m, err} }()
	secondReturned := false
	select {
	case r := <-second:
		secondReturned = true
		if r.err == nil && (r.models == nil || r.models.FindModelNameCaseInsensitive("slow-model") == "") {
			t.Error("concurrent GetModels returned an incomplete catalog with nil error while discovery was pending")
		}
	case <-time.After(100 * time.Millisecond):
		t.Log("reader waited for the pending discovery")
	}
	unblock()
	for i, ch := range []chan result{first, second} {
		if i == 1 && secondReturned {
			continue
		}
		select {
		case r := <-ch:
			if r.err != nil || r.models.FindModelNameCaseInsensitive("slow-model") != "slow-model" {
				t.Errorf("completed catalog missing: err=%v", r.err)
			}
		case <-time.After(time.Second):
			t.Fatal("discovery did not finish after release")
		}
	}
	// The provider and later cached read work; only the in-flight reader was wrong.
	m, err := manager.GetModels()
	if err != nil || m.FindModelNameCaseInsensitive("slow-model") != "slow-model" {
		t.Fatal("completed cache control failed")
	}
	if got := vendor.calls.Load(); got != 1 {
		t.Errorf("concurrent discovery was not coalesced: calls=%d", got)
	}
}

func TestVendorsManagerRegistrationDuringDiscoveryDoesNotCacheOldGeneration(t *testing.T) {
	manager := NewVendorsManager()
	slow := &blockedCatalogTestVendor{stubVendor: &stubVendor{name: "slow"}, entered: make(chan struct{}), release: make(chan struct{})}
	added := &catalogTestVendor{stubVendor: &stubVendor{name: "beta"}, models: []string{"beta-model"}}
	manager.AddVendors(slow)
	var once sync.Once
	unblock := func() { once.Do(func() { close(slow.release) }) }
	defer unblock()
	done := make(chan error, 1)
	go func() { _, err := manager.GetModels(); done <- err }()
	select {
	case <-slow.entered:
	case <-time.After(time.Second):
		t.Fatal("discovery did not start")
	}
	registered := make(chan struct{})
	go func() { manager.AddVendors(added); close(registered) }()
	select {
	case <-registered:
	case <-time.After(time.Second):
		t.Error("registration blocked behind model provider callback")
		unblock()
		select {
		case <-registered:
		case <-time.After(time.Second):
			t.Fatal("registration did not finish")
		}
	}
	unblock()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("discovery did not finish")
	}
	m, err := manager.GetModels()
	if err != nil || m.FindModelNameCaseInsensitive("beta-model") != "beta-model" {
		t.Errorf("in-flight old generation poisoned new membership cache: err=%v calls=%d", err, added.calls.Load())
	}
	before := added.calls.Load()
	if _, err := manager.GetModels(); err != nil {
		t.Fatal(err)
	}
	if got := added.calls.Load(); got != before {
		t.Errorf("completed model cache was not reused: before=%d after=%d", before, got)
	}
}
