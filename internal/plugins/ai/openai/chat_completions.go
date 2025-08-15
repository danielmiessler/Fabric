package openai

// This file contains helper methods for the Chat Completions API.
// These methods are used as fallbacks for OpenAI-compatible providers
// that don't support the newer Responses API (e.g., Groq, Mistral, etc.).

import (
	"context"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

// sendChatCompletions sends a request using the Chat Completions API
func (o *Client) sendChatCompletions(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	req := o.buildChatCompletionParams(msgs, opts)

	var resp *openai.ChatCompletion
	if resp, err = o.ApiClient.Chat.Completions.New(ctx, req); err != nil {
		return
	}
	if len(resp.Choices) > 0 {
		ret = resp.Choices[0].Message.Content
	}
	return
}

// sendStreamChatCompletions sends a streaming request using the Chat Completions API
func (o *Client) sendStreamChatCompletions(
	msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string,
) (err error) {
	defer close(channel)

	req := o.buildChatCompletionParams(msgs, opts)
	stream := o.ApiClient.Chat.Completions.NewStreaming(context.Background(), req)
	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			channel <- chunk.Choices[0].Delta.Content
		}
	}
	if stream.Err() == nil {
		channel <- "\n"
	}
	return stream.Err()
}

// buildChatCompletionParams builds parameters for the Chat Completions API
func (o *Client) buildChatCompletionParams(
	inputMsgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions,
) (ret openai.ChatCompletionNewParams) {

	messages := make([]openai.ChatCompletionMessageParamUnion, len(inputMsgs))
	for i, msgPtr := range inputMsgs {
		msg := *msgPtr
		if strings.Contains(opts.Model, "deepseek") && len(inputMsgs) == 1 && msg.Role == chat.ChatMessageRoleSystem {
			msg.Role = chat.ChatMessageRoleUser
		}
		messages[i] = o.convertChatMessage(msg)
	}

	ret = openai.ChatCompletionNewParams{
		Model:    shared.ChatModel(opts.Model),
		Messages: messages,
	}

	if !opts.Raw {
		ret.Temperature = openai.Float(opts.Temperature)
		if opts.TopP != 0 {
			ret.TopP = openai.Float(opts.TopP)
		}
		if opts.MaxTokens != 0 {
			ret.MaxTokens = openai.Int(int64(opts.MaxTokens))
		}
		if opts.PresencePenalty != 0 {
			ret.PresencePenalty = openai.Float(opts.PresencePenalty)
		}
		if opts.FrequencyPenalty != 0 {
			ret.FrequencyPenalty = openai.Float(opts.FrequencyPenalty)
		}
		if opts.Seed != 0 {
			ret.Seed = openai.Int(int64(opts.Seed))
		}
	}
	if eff, ok := parseReasoningEffort(opts.Thinking); ok {
		ret.ReasoningEffort = eff
	}
	return
}

// convertChatMessage converts fabric chat message to OpenAI chat completion message
func (o *Client) convertChatMessage(msg chat.ChatCompletionMessage) openai.ChatCompletionMessageParamUnion {
	result := convertMessageCommon(msg)

	switch result.Role {
	case chat.ChatMessageRoleSystem:
		return openai.SystemMessage(result.Content)
	case chat.ChatMessageRoleUser:
		// Handle multi-content messages (text + images)
		if result.HasMultiContent {
			var parts []openai.ChatCompletionContentPartUnionParam
			for _, p := range result.MultiContent {
				switch p.Type {
				case chat.ChatMessagePartTypeText:
					parts = append(parts, openai.TextContentPart(p.Text))
				case chat.ChatMessagePartTypeImageURL:
					parts = append(parts, openai.ImageContentPart(openai.ChatCompletionContentPartImageImageURLParam{URL: p.ImageURL.URL}))
				}
			}
			return openai.UserMessage(parts)
		}
		return openai.UserMessage(result.Content)
	case chat.ChatMessageRoleAssistant:
		return openai.AssistantMessage(result.Content)
	default:
		return openai.UserMessage(result.Content)
	}
}
