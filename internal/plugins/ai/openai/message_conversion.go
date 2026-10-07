package openai

import "github.com/danielmiessler/fabric/internal/chat"

// MessageConversionResult holds the common conversion result
type MessageConversionResult struct {
	Role            string
	Content         string
	MultiContent    []chat.ChatMessagePart
	HasMultiContent bool
}

func convertMessageCommon(msg chat.ChatCompletionMessage) MessageConversionResult {
	// A session file can have an image part with no image_url. Skip it,
	// because the converters read ImageURL.URL.
	var parts []chat.ChatMessagePart
	for _, p := range msg.MultiContent {
		if p.Type != chat.ChatMessagePartTypeImageURL || p.ImageURL != nil {
			parts = append(parts, p)
		}
	}
	return MessageConversionResult{
		Role:            msg.Role,
		Content:         msg.Content,
		MultiContent:    parts,
		HasMultiContent: len(parts) > 0,
	}
}
