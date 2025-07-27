package chat

import (
	"encoding/json"
	"errors"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleFunction  = "function"
	ChatMessageRoleTool      = "tool"
	ChatMessageRoleDeveloper = "developer"
)

var ErrContentFieldsMisused = errors.New("can't use both Content and MultiContent properties simultaneously")

type ChatMessagePartType string

const (
	ChatMessagePartTypeText     ChatMessagePartType = "text"
	ChatMessagePartTypeImageURL ChatMessagePartType = "image_url"
)

type ChatMessageImageURL struct {
	URL string `json:"url,omitempty"`
}

type ChatMessagePart struct {
	Type     ChatMessagePartType  `json:"type,omitempty"`
	Text     string               `json:"text,omitempty"`
	ImageURL *ChatMessageImageURL `json:"image_url,omitempty"`
}

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type ToolType string

const (
	ToolTypeFunction ToolType = "function"
)

type ToolCall struct {
	Index    *int         `json:"index,omitempty"`
	ID       string       `json:"id,omitempty"`
	Type     ToolType     `json:"type"`
	Function FunctionCall `json:"function"`
}

type ChatCompletionMessage struct {
	Role             string            `json:"role"`
	Content          string            `json:"content,omitempty"`
	Refusal          string            `json:"refusal,omitempty"`
	MultiContent     []ChatMessagePart `json:"-"`
	Name             string            `json:"name,omitempty"`
	ReasoningContent string            `json:"reasoning_content,omitempty"`
	FunctionCall     *FunctionCall     `json:"function_call,omitempty"`
	ToolCalls        []ToolCall        `json:"tool_calls,omitempty"`
	ToolCallID       string            `json:"tool_call_id,omitempty"`
}

func (m ChatCompletionMessage) MarshalJSON() ([]byte, error) {
	if m.Content != "" && m.MultiContent != nil {
		return nil, ErrContentFieldsMisused
	}
	if len(m.MultiContent) > 0 {
		msg := struct {
			Role             string            `json:"role"`
			Content          string            `json:"-"`
			Refusal          string            `json:"refusal,omitempty"`
			MultiContent     []ChatMessagePart `json:"content,omitempty"`
			Name             string            `json:"name,omitempty"`
			ReasoningContent string            `json:"reasoning_content,omitempty"`
			FunctionCall     *FunctionCall     `json:"function_call,omitempty"`
			ToolCalls        []ToolCall        `json:"tool_calls,omitempty"`
			ToolCallID       string            `json:"tool_call_id,omitempty"`
		}(m)
		return json.Marshal(msg)
	}

	msg := struct {
		Role             string            `json:"role"`
		Content          string            `json:"content,omitempty"`
		Refusal          string            `json:"refusal,omitempty"`
		MultiContent     []ChatMessagePart `json:"-"`
		Name             string            `json:"name,omitempty"`
		ReasoningContent string            `json:"reasoning_content,omitempty"`
		FunctionCall     *FunctionCall     `json:"function_call,omitempty"`
		ToolCalls        []ToolCall        `json:"tool_calls,omitempty"`
		ToolCallID       string            `json:"tool_call_id,omitempty"`
	}(m)
	return json.Marshal(msg)
}

func (m *ChatCompletionMessage) UnmarshalJSON(bs []byte) error {
	msg := struct {
		Role             string `json:"role"`
		Content          string `json:"content"`
		Refusal          string `json:"refusal,omitempty"`
		MultiContent     []ChatMessagePart
		Name             string        `json:"name,omitempty"`
		ReasoningContent string        `json:"reasoning_content,omitempty"`
		FunctionCall     *FunctionCall `json:"function_call,omitempty"`
		ToolCalls        []ToolCall    `json:"tool_calls,omitempty"`
		ToolCallID       string        `json:"tool_call_id,omitempty"`
	}{}

	if err := json.Unmarshal(bs, &msg); err == nil {
		*m = ChatCompletionMessage(msg)
		return nil
	}
	multiMsg := struct {
		Role             string `json:"role"`
		Content          string
		Refusal          string            `json:"refusal,omitempty"`
		MultiContent     []ChatMessagePart `json:"content"`
		Name             string            `json:"name,omitempty"`
		ReasoningContent string            `json:"reasoning_content,omitempty"`
		FunctionCall     *FunctionCall     `json:"function_call,omitempty"`
		ToolCalls        []ToolCall        `json:"tool_calls,omitempty"`
		ToolCallID       string            `json:"tool_call_id,omitempty"`
	}{}
	if err := json.Unmarshal(bs, &multiMsg); err != nil {
		return err
	}
	*m = ChatCompletionMessage(multiMsg)
	return nil
}
