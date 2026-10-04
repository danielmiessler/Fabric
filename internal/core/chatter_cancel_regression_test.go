package core

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// A finite provider that delivers one update, then an optional pending tail.
// This isolates forwarding and drain cleanup from provider cancellation.
type completedStreamVendor struct {
	*mockVendor
	sent     chan struct{}
	finished chan struct{}
	chunks   int
	output   chan domain.StreamUpdate
}

func (v *completedStreamVendor) SendStream(_ context.Context, _ []*chat.ChatCompletionMessage, _ *domain.ChatOptions, output chan domain.StreamUpdate) error {
	defer close(output)
	if v.finished != nil {
		defer close(v.finished)
	}
	v.output = output
	chunks := v.chunks
	if chunks == 0 {
		chunks = 1
	}
	for i := 0; i < chunks; i++ {
		output <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: "chunk"}
		if i == 0 {
			close(v.sent)
		}
	}
	return nil
}

func TestChatterSendCancellationDrainsPendingVendorUpdates(t *testing.T) {
	vendor := &completedStreamVendor{
		mockVendor: &mockVendor{}, sent: make(chan struct{}),
		finished: make(chan struct{}), chunks: 3,
	}
	chatter := &Chatter{db: fsdb.NewDb(t.TempDir()), Stream: true, vendor: vendor, model: "test-model"}
	request := &domain.ChatRequest{Message: &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: "test"}}
	updates := make(chan domain.StreamUpdate)
	opts := &domain.ChatOptions{Model: "test-model", UpdateChan: updates, Quiet: true}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		_, err := chatter.Send(ctx, request, opts)
		done <- err
	}()
	select {
	case <-vendor.sent:
	case <-time.After(2 * time.Second):
		t.Fatal("vendor did not deliver the first update")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context cancellation, got %v", err)
		}
	case <-time.After(250 * time.Millisecond):
		t.Error("Send did not complete after cancellation")
		// Cleanup for the original blocking implementation.
	cleanup:
		for {
			select {
			case <-updates:
			case <-done:
				break cleanup
			case <-time.After(time.Second):
				t.Fatal("could not release Send during cleanup")
			}
		}
	}
	select {
	case <-vendor.finished:
	case <-time.After(time.Second):
		// A naive early return from Send strands this provider send.
		// Drain the captured upstream channel to clean up that bad control.
		for range vendor.output {
		}
		t.Fatal("vendor was stranded; cancellation must drain its pending updates")
	}
}

func TestChatterSendCanceledAfterUpdateConsumerExit(t *testing.T) {
	vendor := &completedStreamVendor{mockVendor: &mockVendor{}, sent: make(chan struct{})}
	chatter := &Chatter{db: fsdb.NewDb(t.TempDir()), Stream: true, vendor: vendor, model: "test-model"}
	request := &domain.ChatRequest{Message: &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: "test"}}
	updates := make(chan domain.StreamUpdate) // HTTP consumer has exited: nobody drains this.
	opts := &domain.ChatOptions{Model: "test-model", UpdateChan: updates, Quiet: true}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		_, err := chatter.Send(ctx, request, opts)
		done <- err
	}()
	select {
	case <-vendor.sent:
	case err := <-done:
		t.Fatalf("Send returned before receiving the update: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("vendor did not deliver the update")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected context cancellation, got %v", err)
		}
	case <-time.After(250 * time.Millisecond):
		t.Error("Chatter.Send stayed blocked on UpdateChan after request cancellation and consumer exit")
		// Release the original implementation's send so the regression leaves no leaked goroutine.
		select {
		case <-updates:
		case <-time.After(time.Second):
			t.Fatal("could not release the blocked update")
		}
		select {
		case err := <-done:
			t.Logf("Send completed only after the consumer resumed; returned error: %v", err)
		case <-time.After(time.Second):
			t.Fatal("Send did not exit after the consumer resumed")
		}
	}
}
