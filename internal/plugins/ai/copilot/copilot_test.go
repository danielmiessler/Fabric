package copilot

import (
	"strings"
	"testing"

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
