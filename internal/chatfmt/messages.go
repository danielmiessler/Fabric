package chatfmt

import (
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
)

// FormatMessages writes each message as a role header line and its content.
// --dry-run and --print-prompt use this format.
func FormatMessages(msgs []*chat.ChatCompletionMessage) string {
	var builder strings.Builder

	for _, msg := range msgs {
		header := roleHeader(msg.Role)
		if len(msg.MultiContent) == 0 {
			fmt.Fprintf(&builder, "%s:\n%s\n\n", header, msg.Content)
			continue
		}
		fmt.Fprintf(&builder, "%s:\n", header)
		for _, part := range msg.MultiContent {
			fmt.Fprintf(&builder, "  - Type: %s\n", part.Type)
			switch {
			case part.Type != chat.ChatMessagePartTypeImageURL:
				fmt.Fprintf(&builder, "    Text: %s\n", part.Text)
			case part.ImageURL == nil:
				builder.WriteString("    Image URL: <missing>\n")
			default:
				fmt.Fprintf(&builder, "    Image URL: %s\n", part.ImageURL.URL)
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// roleHeader changes the first letter of the role to uppercase, for example "user" to "User".
func roleHeader(role string) string {
	if role == "" {
		return role
	}
	return strings.ToUpper(role[:1]) + role[1:]
}
