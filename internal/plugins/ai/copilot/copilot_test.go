package copilot

import (
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
)

// TestParseSSEStreamLargeFrame reproduces the buffer-size failure in
// parseSSEStream. Each Copilot SSE frame carries the cumulative message-so-far
// on a single "data: {json}" line. Once the accumulated response exceeds
// bufio.MaxScanTokenSize (64 KiB) the default scanner aborts with
// bufio.ErrTooLong, parseSSEStream returns copilot_error_reading_stream, and
// the response is silently truncated.
//
// parseSSEStream is unexported; this in-package test calls it directly on a
// zero-value *Client (it only reaches c.extractResponseText, which reads no
// Client state).
func TestParseSSEStreamLargeFrame(t *testing.T) {
	// Build a response text larger than the default 64 KiB scanner token.
	largeText := strings.Repeat("A", 100*1024)

	// One SSE frame: a conversationResponse whose single assistant message
	// carries the whole cumulative text on one line.
	frame := `data: {"messages":[{"@odata.type":"#microsoft.graph.copilotConversationResponseMessage","text":"` +
		largeText + `"}]}` + "\n"

	reader := strings.NewReader(frame)
	channel := make(chan domain.StreamUpdate, 64)

	c := &Client{}

	errCh := make(chan error, 1)
	go func() {
		errCh <- c.parseSSEStream(reader, channel)
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

	// parseSSEStream appends a trailing newline after the stream ends.
	want := largeText + "\n"
	if got.String() != want {
		t.Fatalf("content truncated: got %d bytes, want %d bytes", got.Len(), len(want))
	}
}
