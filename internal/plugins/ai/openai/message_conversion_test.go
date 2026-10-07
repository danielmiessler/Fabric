package openai

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
)

// TestConvertSkipsImagePartWithNoURL checks that an image part with no
// image_url does not stop the converters with a nil pointer.
func TestConvertSkipsImagePartWithNoURL(t *testing.T) {
	msg := chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "hi"},
			{Type: chat.ChatMessagePartTypeImageURL},
		},
	}
	if got := convertMessageCommon(msg).MultiContent; len(got) != 1 || got[0].Text != "hi" {
		t.Fatalf("MultiContent = %+v, want only the text part", got)
	}
	_ = (&Client{}).convertChatMessage(msg)
	_ = convertMessage(msg)
}
