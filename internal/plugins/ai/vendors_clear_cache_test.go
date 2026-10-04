package ai

import (
	"sync"
	"testing"
	"time"
)

func TestVendorsManagerClearReleasesReadersFromObsoleteDiscovery(t *testing.T) {
	manager := NewVendorsManager()
	old := &blockedCatalogTestVendor{stubVendor: &stubVendor{name: "removed"}, entered: make(chan struct{}), release: make(chan struct{})}
	healthy := &catalogTestVendor{stubVendor: &stubVendor{name: "replacement"}, models: []string{"replacement-model"}}
	manager.AddVendors(old)
	var once sync.Once
	release := func() { once.Do(func() { close(old.release) }) }
	defer release()
	type result struct {
		models *VendorsModels
		err    error
	}
	original, cleared, replacement := make(chan result, 1), make(chan result, 1), make(chan result, 1)
	go func() { m, err := manager.GetModels(); original <- result{m, err} }()
	select {
	case <-old.entered:
	case <-time.After(time.Second):
		t.Fatal("old discovery did not start")
	}
	manager.mu.Lock()
	oldLoading := manager.modelsLoading
	manager.mu.Unlock()
	manager.Clear()
	select {
	case <-oldLoading.done:
	default:
		t.Error("invalidation did not notify readers already waiting on the old discovery")
	}
	go func() { m, err := manager.GetModels(); cleared <- result{m, err} }()
	select {
	case r := <-cleared:
		if r.err == nil {
			t.Error("cleared manager must report that no vendors are configured")
		}
	case <-time.After(time.Second):
		t.Error("cleared manager waited on the removed provider")
		release()
		select {
		case <-cleared:
		case <-time.After(time.Second):
			t.Fatal("cleared reader did not exit during cleanup")
		}
	}
	manager.AddVendors(healthy)
	go func() { m, err := manager.GetModels(); replacement <- result{m, err} }()
	for _, ch := range []chan result{replacement} {
		select {
		case r := <-ch:
			if r.err != nil || r.models.FindModelNameCaseInsensitive("replacement-model") != "replacement-model" {
				t.Errorf("replacement catalog unavailable: err=%v", r.err)
			}
		case <-time.After(time.Second):
			t.Error("new membership waited on the removed provider")
			release() // Cleanup the bad control so no goroutine is left behind.
			select {
			case <-ch:
			case <-time.After(time.Second):
				t.Fatal("reader did not exit during cleanup")
			}
		}
	}
	release()
	select {
	case <-original:
	case <-time.After(time.Second):
		t.Fatal("old caller did not exit after release")
	}
	models, err := manager.GetModels()
	if err != nil || models.FindModelNameCaseInsensitive("replacement-model") != "replacement-model" || models.FindModelNameCaseInsensitive("slow-model") != "" {
		t.Fatalf("obsolete completion replaced the current catalog: err=%v", err)
	}
}

func TestVendorsManagerOldCompletionPreservesPendingReplacementSlot(t *testing.T) {
	manager := NewVendorsManager()
	old := &blockedCatalogTestVendor{stubVendor: &stubVendor{name: "old"}, entered: make(chan struct{}), release: make(chan struct{})}
	next := &blockedCatalogTestVendor{stubVendor: &stubVendor{name: "new"}, entered: make(chan struct{}), release: make(chan struct{})}
	var oldOnce, nextOnce sync.Once
	releaseOld := func() { oldOnce.Do(func() { close(old.release) }) }
	releaseNext := func() { nextOnce.Do(func() { close(next.release) }) }
	defer releaseOld()
	defer releaseNext()
	manager.AddVendors(old)
	doneOld, doneNext := make(chan error, 1), make(chan error, 1)
	go func() { _, err := manager.GetModels(); doneOld <- err }()
	select {
	case <-old.entered:
	case <-time.After(time.Second):
		t.Fatal("old lookup did not start")
	}
	manager.Clear()
	manager.AddVendors(next)
	go func() { _, err := manager.GetModels(); doneNext <- err }()
	select {
	case <-next.entered:
	case <-time.After(time.Second):
		t.Fatal("replacement lookup did not start")
	}
	manager.mu.Lock()
	pending := manager.modelsLoading
	manager.mu.Unlock()
	releaseOld()
	select {
	case err := <-doneOld:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("old lookup did not finish")
	}
	manager.mu.Lock()
	if manager.modelsLoading != pending || manager.Models != nil {
		t.Error("old completion altered the new loading slot or cache")
	}
	manager.mu.Unlock()
	releaseNext()
	select {
	case err := <-doneNext:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("replacement lookup did not finish")
	}
	if next.calls.Load() != 1 {
		t.Error("replacement discovery was not coalesced")
	}
}
