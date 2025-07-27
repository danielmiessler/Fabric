package domain

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeMessages(t *testing.T) {
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
		{Role: chat.ChatMessageRoleAssistant, Content: "Hi there!"},
		{Role: chat.ChatMessageRoleUser, Content: ""},
		{Role: chat.ChatMessageRoleUser, Content: ""},
		{Role: chat.ChatMessageRoleUser, Content: "How are you?"},
	}

	expected := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
		{Role: chat.ChatMessageRoleAssistant, Content: "Hi there!"},
		{Role: chat.ChatMessageRoleUser, Content: "How are you?"},
	}

	actual := NormalizeMessages(msgs, "default")
	assert.Equal(t, expected, actual)
}
