package anthropic

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// An error event in the middle of a stream must come back as the error of
// SendStream, as it does for the other vendors, so that the command fails
// instead of printing a cut-off answer and exiting with status 0.
func TestSendStream_ReturnsMidStreamError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		events := []string{
			`event: message_start` + "\n" + `data: {"type":"message_start","message":{"id":"msg_1","type":"message","role":"assistant","model":"claude-sonnet-5-5","content":[],"stop_reason":null,"stop_sequence":null,"usage":{"input_tokens":5,"output_tokens":1}}}`,
			`event: content_block_start` + "\n" + `data: {"type":"content_block_start","index":0,"content_block":{"type":"text","text":""}}`,
			`event: content_block_delta` + "\n" + `data: {"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":"Hel"}}`,
			`event: error` + "\n" + `data: {"type":"error","error":{"type":"overloaded_error","message":"Overloaded"}}`,
		}
		fmt.Fprint(w, strings.Join(events, "\n\n")+"\n\n")
	}))
	defer server.Close()

	client := &Client{
		maxTokens: 4096,
		client: anthropic.NewClient(
			option.WithBaseURL(server.URL),
			option.WithAPIKey("test-key"),
			option.WithMaxRetries(0),
		),
	}

	channel := make(chan domain.StreamUpdate, 16)
	err := client.SendStream(
		context.Background(),
		[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}},
		&domain.ChatOptions{Model: "claude-sonnet-5-5", Temperature: domain.DefaultTemperature, TopP: domain.DefaultTopP},
		channel,
	)

	if err == nil {
		t.Fatal("expected the stream error to be returned, got nil")
	}
	if !strings.Contains(err.Error(), "Overloaded") {
		t.Errorf("expected the error to name the cause, got %v", err)
	}

	var got string
	for update := range channel {
		got += update.Content
	}
	if got != "Hel" {
		t.Errorf("expected the text streamed before the error, got %q", got)
	}
}
