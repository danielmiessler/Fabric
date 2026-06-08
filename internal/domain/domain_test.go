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
		tt := tt // Capture range variable
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

// Regression: ensure NormalizeInputShape does not mutate the caller's slice
// or the message structs it points to (Copilot PR review #2137).
func TestNormalizeInputShapeDoesNotMutateOriginal(t *testing.T) {
	orig := []*chat.ChatCompletionMessage{sys("instruction")}
	got := NormalizeInputShape(orig)
	assert.Equal(t, chat.ChatMessageRoleUser, got[0].Role,
		"normalized message should be user")
	assert.Equal(t, chat.ChatMessageRoleSystem, orig[0].Role,
		"original message should remain system")
}

func TestNormalizeInputShapeWithNilMessages(t *testing.T) {
	// Test nil at beginning, user later - should return original slice (has user)
	in1 := []*chat.ChatCompletionMessage{nil, sys("sys1"), usr("user1")}
	got1 := NormalizeInputShape(in1)
	assert.Equal(t, in1, got1) // Should return original slice (has user)

	// Test nil at end, no user - should promote last non-nil system
	in2 := []*chat.ChatCompletionMessage{sys("sys1"), sys("sys2"), nil}
	got2 := NormalizeInputShape(in2)
	assert.Len(t, got2, 3)
	assert.Equal(t, sys("sys1"), got2[0])
	// Last non-nil (sys2) should be promoted to user
	expectedPromoted := usr("sys2")
	assert.Equal(t, expectedPromoted.Role, got2[1].Role)
	assert.Equal(t, expectedPromoted.Content, got2[1].Content)
	assert.Equal(t, (*chat.ChatCompletionMessage)(nil), got2[2])

	// Test all nil
	in3 := []*chat.ChatCompletionMessage{nil, nil, nil}
	got3 := NormalizeInputShape(in3)
	assert.Equal(t, in3, got3) // Should return original slice

	// Test nil mixed with user (should return early due to user)
	in4 := []*chat.ChatCompletionMessage{nil, usr("user1"), nil}
	got4 := NormalizeInputShape(in4)
	assert.Equal(t, in4, got4) // Should return original slice (has user)
}
