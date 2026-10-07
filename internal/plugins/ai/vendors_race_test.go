package ai

import (
	"sync"
	"testing"
)

// TestVendorsManager_GetModelsConcurrent calls GetModels from 8 goroutines
// at the same time. Run it with -race. Without modelsMu, the race detector
// finds a data race on Models.
func TestVendorsManager_GetModelsConcurrent(t *testing.T) {
	manager := NewVendorsManager()
	manager.AddVendors(&stubVendor{name: "OpenAI"}, &stubVendor{name: "Anthropic"})

	const n = 8
	var wg sync.WaitGroup
	results := make([]*VendorsModels, n)
	wg.Add(n)
	for i := range n {
		go func() {
			defer wg.Done()
			results[i], _ = manager.GetModels()
		}()
	}
	wg.Wait()

	// The lazy init runs one time. Thus each call gets the same models.
	for i, got := range results {
		if got == nil || got != results[0] {
			t.Fatalf("call %d got models %p, want %p", i, got, results[0])
		}
	}
}
