package chatfmt

import (
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
)

func FormatMessages(msgs []*chat.ChatCompletionMessage) string {
	var builder strings.Builder

	for _, msg := range msgs {
		builder.WriteString(FormatMessage(msg))
	}

	return builder.String()
}

func FormatMessage(msg *chat.ChatCompletionMessage) string {
	var builder strings.Builder
	header := roleHeader(msg.Role)

	if len(msg.MultiContent) > 0 {
		builder.WriteString(fmt.Sprintf("%s:\n", header))
		for _, part := range msg.MultiContent {
			builder.WriteString(fmt.Sprintf("  - Type: %s\n", part.Type))
			if part.Type == chat.ChatMessagePartTypeImageURL && part.ImageURL != nil {
				builder.WriteString(fmt.Sprintf("    Image URL: %s\n", part.ImageURL.URL))
				continue
			}
			builder.WriteString(fmt.Sprintf("    Text: %s\n", part.Text))
		}
		builder.WriteString("\n")
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%s:\n%s\n\n", header, msg.Content))
	return builder.String()
}

func roleHeader(role string) string {
	switch role {
	case chat.ChatMessageRoleSystem:
		return "System"
	case chat.ChatMessageRoleUser:
		return "User"
	case chat.ChatMessageRoleAssistant:
		return "Assistant"
	case chat.ChatMessageRoleDeveloper:
		return "Developer"
	case chat.ChatMessageRoleTool:
		return "Tool"
	case chat.ChatMessageRoleFunction:
		return "Function"
	default:
		return role
	}
}
