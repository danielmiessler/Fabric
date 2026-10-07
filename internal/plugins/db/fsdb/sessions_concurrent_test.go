package fsdb

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
)

// TestSessions_ConcurrentAppendKeepsAllMessages runs many appends to one
// session at the same time. Each goroutine does the same steps as
// core.Chatter.Send: Lock, Get, Append, SaveSession, unlock. All messages
// must stay in the file.
func TestSessions_ConcurrentAppendKeepsAllMessages(t *testing.T) {
	dir := t.TempDir()
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: dir, FileExtension: ".json"},
	}
	if err := sessions.Configure(); err != nil {
		t.Fatalf("configure: %v", err)
	}

	const name = "race"
	const workers = 30

	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := range workers {
		go func(i int) {
			defer wg.Done()
			<-start
			unlock := sessions.Lock(name)
			defer unlock()

			sess, err := sessions.Get(name)
			if err != nil {
				t.Errorf("get: %v", err)
				return
			}
			sess.Append(&chat.ChatCompletionMessage{
				Role:    "user",
				Content: fmt.Sprintf("msg-%d", i),
			})
			if err := sessions.SaveSession(sess); err != nil {
				t.Errorf("save: %v", err)
			}
		}(i)
	}
	close(start)
	wg.Wait()

	final, err := sessions.Get(name)
	if err != nil {
		t.Fatalf("final get: %v", err)
	}
	if len(final.Messages) != workers {
		t.Fatalf("lost updates: got %d messages, want %d", len(final.Messages), workers)
	}
}

// TestSessions_WritesWaitForLock checks that Save, Delete and Rename wait
// for the session lock, and that an invalid name adds no lock.
func TestSessions_WritesWaitForLock(t *testing.T) {
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: t.TempDir(), FileExtension: ".json"},
	}
	if err := sessions.Configure(); err != nil {
		t.Fatalf("configure: %v", err)
	}

	for name, write := range map[string]func() error{
		"save":   func() error { return sessions.Save("s", []byte("[]")) },
		"delete": func() error { return sessions.Delete("s") },
		"rename": func() error { return sessions.Rename("t", "s") },
	} {
		unlock := sessions.Lock("s")
		done := make(chan error, 1)
		go func() { done <- write() }()
		select {
		case <-done:
			t.Errorf("%s did not wait for the lock", name)
		case <-time.After(50 * time.Millisecond):
		}
		unlock()
		<-done
	}

	if err := sessions.Save("../bad", nil); err == nil {
		t.Error("Save accepted an invalid name")
	}
	if _, ok := sessions.locks.Load("../bad"); ok {
		t.Error("Save added a lock for an invalid name")
	}
}
