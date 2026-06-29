package openai

// This file contains helper methods for the Chat Completions API.
// These methods are used as fallbacks for OpenAI-compatible providers
// that don't support the newer Responses API (e.g., Groq, Mistral, etc.).

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

// sendChatCompletions sends a request using the Chat Completions API
// If the SDK fails (for example when the server returns SSE even for
// non-stream responses), fall back to a direct HTTP request that can
// parse text/event-stream (SSE) and concatenate 'data:' fields.
func (o *Client) sendChatCompletions(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	req := o.buildChatCompletionParams(msgs, opts)

	var resp *openai.ChatCompletion
	if resp, err = o.ApiClient.Chat.Completions.New(ctx, req); err == nil {
		if len(resp.Choices) > 0 {
			ret = resp.Choices[0].Message.Content
		}
		return
	}

	// SDK failed - attempt direct HTTP fallback that handles SSE
	return o.sendChatCompletionsDirect(ctx, msgs, opts)
}

// sendChatCompletionsDirect performs a direct HTTP POST to the chat/completions
// endpoint and handles both application/json and text/event-stream responses.
// It builds the request from the provided messages and options instead of
// relying on SDK param types.
func (o *Client) sendChatCompletionsDirect(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	// Build JSON body
	payload := make(map[string]any)
	payload["model"] = opts.Model

	// Build messages array
	var messages []map[string]any
	for _, m := range msgs {
		entry := map[string]any{"role": m.Role, "content": m.Content}
		messages = append(messages, entry)
	}
	payload["messages"] = messages

	if !opts.Raw {
		payload["temperature"] = opts.Temperature
		if opts.TopP != 0 {
			payload["top_p"] = opts.TopP
		}
		if opts.MaxTokens != 0 {
			payload["max_tokens"] = opts.MaxTokens
		}
	}

	body, jerr := json.Marshal(payload)
	if jerr != nil {
		return "", jerr
	}

	// Ensure base URL ends without trailing slash
	base := strings.TrimRight(o.ApiBaseURL.Value, "/")
	url := base + "/chat/completions"

	req, rerr := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if rerr != nil {
		return "", rerr
	}
	req.Header.Set("Content-Type", "application/json")
	if o.ApiKey != nil && o.ApiKey.Value != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.ApiKey.Value))
	}

	resp, derr := o.httpClient.Do(req)
	if derr != nil {
		return "", derr
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if strings.Contains(ct, "application/json") {
		var parsed struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
			return "", err
		}
		if len(parsed.Choices) > 0 {
			return parsed.Choices[0].Message.Content, nil
		}
		return "", nil
	}

	// Handle text/event-stream (SSE) by scanning data: lines and concatenating
	if strings.Contains(ct, "text/event-stream") || strings.Contains(ct, "event-stream") {
		return parseSSEAndConcat(resp.Body)
	}

	// Unknown content-type: attempt to read body as text
	b, rerr := io.ReadAll(resp.Body)
	if rerr != nil {
		return "", rerr
	}
	return string(b), nil
}

// parseSSEAndConcat reads an SSE stream and concatenates any 'data:' JSON
// payloads. It handles both JSON objects containing choices deltas and plain text.
func parseSSEAndConcat(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	var parts []string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// SSE data lines start with 'data:'
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			if data == "[DONE]" || data == "[done]" {
				break
			}
			// Try to parse JSON
			var obj any
			if err := json.Unmarshal([]byte(data), &obj); err == nil {
				// Traverse object to find any 'choices' -> delta -> content or message->content
				if m, ok := obj.(map[string]any); ok {
					if choices, ok := m["choices"].([]any); ok {
						for _, ch := range choices {
							if chm, ok := ch.(map[string]any); ok {
								// delta.content
								if delta, ok := chm["delta"].(map[string]any); ok {
									if c, ok := delta["content"].(string); ok {
										parts = append(parts, c)
									}
								}
								// message.content (final messages)
								if msg, ok := chm["message"].(map[string]any); ok {
									if content, ok := msg["content"].(string); ok {
										parts = append(parts, content)
									}
								}
							}
						}
					}
				}
				continue
			}
			// Not JSON — treat data as raw text
			parts = append(parts, data)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(parts, ""), nil
}

// sendStreamChatCompletions sends a streaming request using the Chat Completions API
func (o *Client) sendStreamChatCompletions(
	ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) (err error) {
	defer close(channel)

	req := o.buildChatCompletionParams(msgs, opts)
	// Set StreamOptions only for streaming requests (required to get usage stats)
	req.StreamOptions = openai.ChatCompletionStreamOptionsParam{
		IncludeUsage: openai.Bool(true),
	}
	stream := o.ApiClient.Chat.Completions.NewStreaming(ctx, req)
	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: chunk.Choices[0].Delta.Content,
			}
		}

		if chunk.Usage.TotalTokens > 0 {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(chunk.Usage.PromptTokens),
					OutputTokens: int(chunk.Usage.CompletionTokens),
					TotalTokens:  int(chunk.Usage.TotalTokens),
				},
			}
		}
	}
	if stream.Err() == nil {
		channel <- domain.StreamUpdate{
			Type:    domain.StreamTypeContent,
			Content: "\n",
		}
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
