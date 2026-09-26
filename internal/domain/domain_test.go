package domain

import (
	"encoding/json"
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

// SessionID is derived internally and must not leak into the REST API surface
// (ChatOptions is embedded in the server's JSON-bound request DTO).
func TestChatOptions_SessionIDExcludedFromJSON(t *testing.T) {
	opts := ChatOptions{Model: "glm-5.1", SessionID: "session-123"}
	data, err := json.Marshal(opts)
	assert.NoError(t, err)
	assert.NotContains(t, string(data), "SessionID")
	assert.NotContains(t, string(data), "session-123")
}
