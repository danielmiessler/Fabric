package copilot

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Each Copilot SSE frame holds all the text so far on one "data:" line. A
// frame larger than the 64 KiB default scanner limit must not stop the stream.
func TestParseSSEStreamLargeFrame(t *testing.T) {
	largeText := strings.Repeat("A", 100*1024)
	frame := `data: {"messages":[{"@odata.type":"#microsoft.graph.copilotConversationResponseMessage","text":"` +
		largeText + `"}]}` + "\n"

	channel := make(chan domain.StreamUpdate, 64)
	errCh := make(chan error, 1)
	go func() {
		// parseSSEStream reads no Client state, so a zero Client is enough.
		errCh <- (&Client{}).parseSSEStream(strings.NewReader(frame), channel)
		close(channel)
	}()

	var got strings.Builder
	for update := range channel {
		if update.Type == domain.StreamTypeContent {
			got.WriteString(update.Content)
		}
	}

	if err := <-errCh; err != nil {
		t.Fatalf("parseSSEStream returned error on a >64KiB frame: %v", err)
	}
	// parseSSEStream adds a newline after the stream ends.
	if want := largeText + "\n"; got.String() != want {
		t.Fatalf("content truncated: got %d bytes, want %d bytes", got.Len(), len(want))
	}
}

// Cancelling the caller's context must abort an in-flight Copilot stream
// instead of letting it run to server completion.
func TestSendStreamHonorsContextCancellation(t *testing.T) {
	const totalFrames = 40

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/conversations") {
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"id":"conv-1"}`))
			return
		}
		for range totalFrames {
			if r.Context().Err() != nil {
				return
			}
			_, _ = w.Write([]byte(`data: {"messages":[{"@odata.type":"#microsoft.graph.copilotConversationResponseMessage","text":"chunk"}]}` + "\n\n"))
			w.(http.Flusher).Flush()
			time.Sleep(25 * time.Millisecond)
		}
	}))
	t.Cleanup(server.Close)

	client := NewClient()
	client.ApiBaseURL.Value = server.URL
	client.httpClient = server.Client()

	channel := make(chan domain.StreamUpdate)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	errCh := make(chan error, 1)
	go func() {
		errCh <- client.SendStream(ctx, []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hello"}}, &domain.ChatOptions{}, channel)
	}()

	// Consume the first update, then cancel mid-stream.
	if _, ok := <-channel; !ok {
		t.Fatal("expected at least one stream update before cancellation")
	}
	cancel()

	// Drain the remaining updates so SendStream can return.
	for range channel {
	}

	if err := <-errCh; !errors.Is(err, context.Canceled) {
		t.Fatalf("SendStream error = %v, want context.Canceled", err)
	}
}
