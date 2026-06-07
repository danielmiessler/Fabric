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

func sys(c string) *chat.ChatCompletionMessage {
	return &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleSystem, Content: c}
}

func usr(c string) *chat.ChatCompletionMessage {
	return &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: c}
}

func asst(c string) *chat.ChatCompletionMessage {
	return &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleAssistant, Content: c}
}

func TestNormalizeInputShape(t *testing.T) {
	tests := []struct {
		name string
		in   []*chat.ChatCompletionMessage
		want []*chat.ChatCompletionMessage
	}{
		{
			// Single system message -- the most common failing case
			name: "bare system",
			in:   []*chat.ChatCompletionMessage{sys("instr")},
			want: []*chat.ChatCompletionMessage{usr("instr")},
		},
		{
			// Multiple system messages -- last one promoted to user
			name: "multiple systems",
			in:   []*chat.ChatCompletionMessage{sys("a"), sys("b")},
			want: []*chat.ChatCompletionMessage{sys("a"), usr("b")},
		},
		{
			// Already has user message -- unchanged
			name: "system + user",
			in:   []*chat.ChatCompletionMessage{sys("instr"), usr("input")},
			want: []*chat.ChatCompletionMessage{sys("instr"), usr("input")},
		},
		{
			// Multi-turn session -- unchanged
			name: "multi-turn",
			in:   []*chat.ChatCompletionMessage{sys("instr"), usr("q1"), asst("a1"), usr("q2")},
			want: []*chat.ChatCompletionMessage{sys("instr"), usr("q1"), asst("a1"), usr("q2")},
		},
		{
			// User only -- unchanged
			name: "user only",
			in:   []*chat.ChatCompletionMessage{usr("hello")},
			want: []*chat.ChatCompletionMessage{usr("hello")},
		},
		{
			// Session continuation -- unchanged
			name: "assistant then user",
			in:   []*chat.ChatCompletionMessage{asst("a"), usr("q")},
			want: []*chat.ChatCompletionMessage{asst("a"), usr("q")},
		},
		{
			// Empty -- unchanged
			name: "empty",
			in:   []*chat.ChatCompletionMessage{},
			want: []*chat.ChatCompletionMessage{},
		},
		{
			// System, assistant, system -- last system promoted
			name: "sys-asst-sys",
			in:   []*chat.ChatCompletionMessage{sys("s"), asst("a"), sys("s2")},
			want: []*chat.ChatCompletionMessage{sys("s"), asst("a"), usr("s2")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeInputShape(tt.in)
			assert.Len(t, got, len(tt.want))
			for i := range got {
				assert.Equal(t, tt.want[i].Role, got[i].Role)
				assert.Equal(t, tt.want[i].Content, got[i].Content)
			}
		})
	}
}
