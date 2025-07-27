package openai

import "github.com/danielmiessler/fabric/internal/chat"

// MessageConversionResult holds the common conversion result
type MessageConversionResult struct {
	Role            string
	Content         string
	MultiContent    []chat.ChatMessagePart
	HasMultiContent bool
}

// convertMessageCommon extracts common conversion logic
func convertMessageCommon(msg chat.ChatCompletionMessage) MessageConversionResult {
	return MessageConversionResult{
		Role:            msg.Role,
		Content:         msg.Content,
		MultiContent:    msg.MultiContent,
		HasMultiContent: len(msg.MultiContent) > 0,
	}
}
