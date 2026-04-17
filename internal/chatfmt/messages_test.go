package chatfmt

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
)

func TestFormatMessages(t *testing.T) {
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "Follow the system prompt"},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
	}

	got := FormatMessages(msgs)
	want := "System:\nFollow the system prompt\n\nUser:\nHello\n\n"

	if got != want {
		t.Fatalf("FormatMessages() = %q, want %q", got, want)
	}
}

func TestFormatMessagesMultiContent(t *testing.T) {
	msgs := []*chat.ChatCompletionMessage{
		{
			Role: chat.ChatMessageRoleUser,
			MultiContent: []chat.ChatMessagePart{
				{Type: chat.ChatMessagePartTypeText, Text: "Describe this image"},
				{
					Type: chat.ChatMessagePartTypeImageURL,
					ImageURL: &chat.ChatMessageImageURL{
						URL: "https://example.com/image.png",
					},
				},
			},
		},
	}

	got := FormatMessages(msgs)
	want := "User:\n" +
		"  - Type: text\n" +
		"    Text: Describe this image\n" +
		"  - Type: image_url\n" +
		"    Image URL: https://example.com/image.png\n\n"

	if got != want {
		t.Fatalf("FormatMessages() = %q, want %q", got, want)
	}
}
