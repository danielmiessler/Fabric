package anthropic

import (
	"context"
	"fmt"
	neturl "net/url"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
)

const defaultBaseUrl = "https://api.anthropic.com/"

const webSearchToolName = "web_search"
const webSearchToolType = "web_search_20250305"
const sourcesHeader = "## Sources"

// These models reject non-default sampling parameters.
// Requests to them do not include temperature or top_p.
var samplingParamsDisallowedPrefixes = []string{
	"claude-opus-4-7",
	"claude-opus-4-8",
	"claude-opus-5",
	"claude-sonnet-5",
	"claude-fable-5",
}

func modelDisallowsSamplingParams(model string) bool {
	return slices.ContainsFunc(samplingParamsDisallowedPrefixes, func(prefix string) bool {
		return strings.HasPrefix(model, prefix)
	})
}

// These models reject the legacy `thinking.type=enabled` and `budget_tokens`
// shape. Requests to them use `thinking.type=adaptive` with
// `output_config.effort`. For the legacy shape, the API returns this error:
//
//	"thinking.type.enabled" is not supported for this model.
//	Use "thinking.type.adaptive" and "output_config.effort" to control
//	thinking behavior.
var adaptiveThinkingPrefixes = []string{
	"claude-opus-5",
	"claude-sonnet-5",
	"claude-fable-5",
}

func modelUsesAdaptiveThinking(model string) bool {
	return slices.ContainsFunc(adaptiveThinkingPrefixes, func(prefix string) bool {
		return strings.HasPrefix(model, prefix)
	})
}

// effortForBudget changes a numeric --thinking value into an effort level for
// adaptive models, which have no budget_tokens. It returns the lowest level
// with a budget equal to or more than tokens. Values above the high budget get
// high. For example, `--thinking=2048` and `--thinking=medium` both give medium.
func effortForBudget(tokens int64) anthropic.OutputConfigEffort {
	switch {
	case tokens <= domain.TokenBudgetLow:
		return anthropic.OutputConfigEffortLow
	case tokens <= domain.TokenBudgetMedium:
		return anthropic.OutputConfigEffortMedium
	default:
		return anthropic.OutputConfigEffortHigh
	}
}

func NewClient() (ret *Client) {
	vendorName := "Anthropic"
	ret = &Client{}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, ret.configure)

	ret.ApiBaseURL = ret.AddSetupQuestion("API Base URL", false)
	ret.ApiBaseURL.Value = defaultBaseUrl
	ret.ApiKey = ret.PluginBase.AddSetupQuestion("API key", false)

	ret.maxTokens = 4096
	ret.defaultRequiredUserMessage = "Hi"
	ret.models = []string{
		string(anthropic.ModelClaudeOpus5_5),
		string(anthropic.ModelClaudeFable5),
		string(anthropic.ModelClaudeSonnet5),
		string(anthropic.ModelClaudeOpus5),
		string(anthropic.ModelClaudeOpus4_8),
		string(anthropic.ModelClaudeOpus4_7),
		string(anthropic.ModelClaudeSonnet4_6),
		string(anthropic.ModelClaudeOpus4_6),
		string(anthropic.ModelClaudeOpus4_5_20251101),
		string(anthropic.ModelClaudeOpus4_5),
		string(anthropic.ModelClaudeHaiku4_5),
		string(anthropic.ModelClaudeHaiku4_5_20251001),
		string(anthropic.ModelClaudeSonnet4_5),
		string(anthropic.ModelClaudeSonnet4_5_20250929),
	}

	// context1M is the beta header for the 1M-token context window. Models with
	// a 1M window use it by default, so the header is not necessary. The code
	// still sends the header as a precaution. If a request with the header gets
	// an error, Send and SendStream send the request again without the header.
	// Only models with a 1M window belong in this map.
	//
	// Window sizes by model:
	// https://platform.claude.com/docs/en/build-with-claude/context-windows#context-window-sizes-by-model
	// Sonnet 4.5, Opus 4.5, Opus 4.1, and Haiku 4.5 have a 200K window, so they
	// are not in the map.
	const context1M = "context-1m-2025-08-07"
	ret.modelBetas = map[string][]string{
		string(anthropic.ModelClaudeOpus5_5): {context1M},
		string(anthropic.ModelClaudeFable5):  {context1M},
		string(anthropic.ModelClaudeOpus5):   {context1M},
		string(anthropic.ModelClaudeSonnet5): {context1M},

		string(anthropic.ModelClaudeOpus4_8): {context1M},
		string(anthropic.ModelClaudeOpus4_7): {context1M},
		string(anthropic.ModelClaudeOpus4_6): {context1M},

		string(anthropic.ModelClaudeSonnet4_6): {context1M},
	}

	return
}

// IsConfigured returns true if the API key is configured
func (an *Client) IsConfigured() bool {
	if an.ApiKey.Value != "" {
		return true
	}

	return false
}

type Client struct {
	*plugins.PluginBase
	ApiBaseURL *plugins.SetupQuestion
	ApiKey     *plugins.SetupQuestion

	maxTokens                  int
	defaultRequiredUserMessage string
	models                     []string
	modelBetas                 map[string][]string

	client anthropic.Client
}

func (an *Client) Setup() (err error) {
	if err = an.PluginBase.Ask(an.Name); err != nil {
		return
	}

	err = an.configure()
	return
}

func (an *Client) configure() (err error) {
	opts := []option.RequestOption{}

	if an.ApiBaseURL.Value != "" {
		opts = append(opts, option.WithBaseURL(an.ApiBaseURL.Value))
	}

	opts = append(opts, option.WithAPIKey(an.ApiKey.Value))

	an.client = anthropic.NewClient(opts...)
	return
}

func (an *Client) ListModels(context.Context) (ret []string, err error) {
	return an.models, nil
}

// parseThinking changes a thinking level into the thinking shape for model.
// It returns a non-empty effort only for adaptive models. The caller must then
// set params.OutputConfig.Effort, because an adaptive request has no budget
// and gets its level from effort.
func parseThinking(level domain.ThinkingLevel, model string) (
	thinking anthropic.ThinkingConfigParamUnion, effort anthropic.OutputConfigEffort, ok bool) {

	lower := strings.ToLower(string(level))
	adaptive := modelUsesAdaptiveThinking(model)
	adaptiveThinking := anthropic.ThinkingConfigParamUnion{OfAdaptive: &anthropic.ThinkingConfigAdaptiveParam{}}

	switch domain.ThinkingLevel(lower) {
	case domain.ThinkingOff:
		disabled := anthropic.NewThinkingConfigDisabledParam()
		return anthropic.ThinkingConfigParamUnion{OfDisabled: &disabled}, "", true
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		if adaptive {
			// The level names are also the effort names.
			return adaptiveThinking, anthropic.OutputConfigEffort(lower), true
		}
		return anthropic.ThinkingConfigParamOfEnabled(domain.ThinkingBudgets[domain.ThinkingLevel(lower)]), "", true
	default:
		if tokens, err := strconv.ParseInt(lower, 10, 64); err == nil && tokens >= 1 && tokens <= 10000 {
			if adaptive {
				return adaptiveThinking, effortForBudget(tokens), true
			}
			return anthropic.ThinkingConfigParamOfEnabled(tokens), "", true
		}
	}
	return anthropic.ThinkingConfigParamUnion{}, "", false
}

func (an *Client) SendStream(
	ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) (err error) {
	messages := an.toMessages(msgs)
	if len(messages) == 0 {
		close(channel)
		// No messages remain after normalization. This is not an error.
		return
	}

	params := an.buildMessageParams(messages, opts)
	betas := an.modelBetas[opts.Model]
	var reqOpts []option.RequestOption
	if len(betas) > 0 {
		reqOpts = append(reqOpts, option.WithHeader("anthropic-beta", strings.Join(betas, ",")))
	}
	stream := an.client.Messages.NewStreaming(ctx, params, reqOpts...)
	if stream.Err() != nil && len(betas) > 0 {
		debuglog.Debug(debuglog.Basic, "Anthropic beta feature %s failed: %v\n", strings.Join(betas, ","), stream.Err())
		stream = an.client.Messages.NewStreaming(ctx, params)
	}

	for stream.Next() {
		event := stream.Current()

		if event.Delta.Text != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.Delta.Text,
			}
		}

		if event.Message.Usage.InputTokens != 0 || event.Message.Usage.OutputTokens != 0 {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(event.Message.Usage.InputTokens),
					OutputTokens: int(event.Message.Usage.OutputTokens),
					TotalTokens:  int(event.Message.Usage.InputTokens + event.Message.Usage.OutputTokens),
				},
			}
		} else if event.Usage.InputTokens != 0 || event.Usage.OutputTokens != 0 {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  int(event.Usage.InputTokens),
					OutputTokens: int(event.Usage.OutputTokens),
					TotalTokens:  int(event.Usage.InputTokens + event.Usage.OutputTokens),
				},
			}
		}
	}

	if stream.Err() != nil {
		fmt.Fprintf(os.Stderr, i18n.T("anthropic_stream_error"), stream.Err())
	}
	close(channel)
	return
}

func (an *Client) buildMessageParams(msgs []anthropic.MessageParam, opts *domain.ChatOptions) (
	params anthropic.MessageNewParams) {

	maxTokens := an.maxTokens
	if opts.MaxTokens > 0 {
		maxTokens = opts.MaxTokens
	}

	params = anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: int64(maxTokens),
		Messages:  msgs,
	}

	if modelDisallowsSamplingParams(opts.Model) {
		// Send no sampling parameters to these models.
	} else if opts.TopP != domain.DefaultTopP {
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		// Send temperature also at its default value. The Fabric default is 0.7,
		// and the API default is 1.0.
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	if opts.Search {
		webTool := anthropic.WebSearchTool20250305Param{
			Name:         webSearchToolName,
			Type:         webSearchToolType,
			CacheControl: anthropic.NewCacheControlEphemeralParam(),
		}

		if opts.SearchLocation != "" {
			webTool.UserLocation.Type = "approximate"
			webTool.UserLocation.Timezone = anthropic.Opt(opts.SearchLocation)
		}

		params.Tools = []anthropic.ToolUnionParam{
			{OfWebSearchTool20250305: &webTool},
		}
	}

	if t, effort, ok := parseThinking(opts.Thinking, opts.Model); ok {
		params.Thinking = t
		params.OutputConfig.Effort = effort
	}

	return
}

func (an *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (
	ret string, err error) {

	messages := an.toMessages(msgs)
	if len(messages) == 0 {
		// No messages remain after normalization. This is not an error.
		return
	}

	var message *anthropic.Message
	params := an.buildMessageParams(messages, opts)
	betas := an.modelBetas[opts.Model]
	var reqOpts []option.RequestOption
	if len(betas) > 0 {
		reqOpts = append(reqOpts, option.WithHeader("anthropic-beta", strings.Join(betas, ",")))
	}
	if message, err = an.client.Messages.New(ctx, params, reqOpts...); err != nil {
		if len(betas) > 0 {
			debuglog.Debug(debuglog.Basic, "Anthropic beta feature %s failed: %v\n", strings.Join(betas, ","), err)
			if message, err = an.client.Messages.New(ctx, params); err != nil {
				return
			}
		} else {
			return
		}
	}

	var textParts []string
	var citations []string
	citationMap := make(map[string]bool) // To prevent duplicate citations

	for _, block := range message.Content {
		if block.Type == "text" && block.Text != "" {
			textParts = append(textParts, block.Text)

			for _, citation := range block.Citations {
				if citation.Type == "web_search_result_location" {
					citationKey := citation.URL + "|" + citation.Title
					if !citationMap[citationKey] {
						citationMap[citationKey] = true
						citationText := fmt.Sprintf("- [%s](%s)", citation.Title, citation.URL)
						if citation.CitedText != "" {
							citationText += fmt.Sprintf(" - \"%s\"", citation.CitedText)
						}
						citations = append(citations, citationText)
					}
				}
			}
		}
	}

	var resultBuilder strings.Builder
	resultBuilder.WriteString(strings.Join(textParts, ""))

	if len(citations) > 0 {
		resultBuilder.WriteString("\n\n")
		resultBuilder.WriteString(sourcesHeader)
		resultBuilder.WriteString("\n\n")
		resultBuilder.WriteString(strings.Join(citations, "\n"))
	}
	ret = resultBuilder.String()

	return
}

func (an *Client) toMessages(msgs []*chat.ChatCompletionMessage) (ret []anthropic.MessageParam) {
	// Normalization rules:
	// - The code puts system content at the start of the next user message.
	//   It does this one time only and ignores later system messages.
	// - Two messages in a row must not have the same role.
	// - The code ignores empty messages.

	var anthropicMessages []anthropic.MessageParam
	var systemContent string

	isFirstUserMessage := true
	lastRoleWasUser := false

	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" && len(msg.MultiContent) == 0 {
			continue
		}

		switch msg.Role {
		case chat.ChatMessageRoleSystem:
			systemText := messageTextFromParts(msg)
			if systemText == "" {
				continue
			}
			if systemContent != "" {
				systemContent += "\n" + systemText
			} else {
				systemContent = systemText
			}
		case chat.ChatMessageRoleUser:
			blocks := contentBlocksFromMessage(msg)
			if len(blocks) == 0 {
				continue
			}
			if isFirstUserMessage && systemContent != "" {
				blocks = prependSystemContentToBlocks(systemContent, blocks)
				isFirstUserMessage = false
			}
			if lastRoleWasUser {
				// Add a short assistant message between two user messages.
				// The chatter.go flow does not usually send this sequence.
				anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock("Okay.")))
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(blocks...))
			lastRoleWasUser = true
		case chat.ChatMessageRoleAssistant:
			// The system content is not in a message yet. Put it in a new user
			// message before this assistant message.
			if isFirstUserMessage && systemContent != "" {
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
				lastRoleWasUser = true
				isFirstUserMessage = false
			} else if !lastRoleWasUser && len(anthropicMessages) > 0 {
				// Add a short user message between two assistant messages.
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(an.defaultRequiredUserMessage)))
				lastRoleWasUser = true
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(msg.Content)))
			lastRoleWasUser = false
		default:
			continue
		}
	}

	if len(anthropicMessages) == 0 && systemContent != "" {
		anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
	}

	return anthropicMessages
}

// messageTextFromParts returns the text in Content and in the text parts of
// MultiContent, with a newline between each.
func messageTextFromParts(msg *chat.ChatCompletionMessage) string {
	textParts := []string{}
	if strings.TrimSpace(msg.Content) != "" {
		textParts = append(textParts, msg.Content)
	}
	for _, part := range msg.MultiContent {
		if part.Type == chat.ChatMessagePartTypeText && strings.TrimSpace(part.Text) != "" {
			textParts = append(textParts, part.Text)
		}
	}
	return strings.Join(textParts, "\n")
}

// contentBlocksFromMessage makes Anthropic content blocks from the text, image,
// and PDF content of a chat message. Images and PDFs can be data URLs or remote URLs.
func contentBlocksFromMessage(msg *chat.ChatCompletionMessage) []anthropic.ContentBlockParamUnion {
	var blocks []anthropic.ContentBlockParamUnion
	if strings.TrimSpace(msg.Content) != "" {
		blocks = append(blocks, anthropic.NewTextBlock(msg.Content))
	}
	for _, part := range msg.MultiContent {
		switch part.Type {
		case chat.ChatMessagePartTypeText:
			if strings.TrimSpace(part.Text) != "" {
				blocks = append(blocks, anthropic.NewTextBlock(part.Text))
			}
		case chat.ChatMessagePartTypeImageURL:
			if part.ImageURL == nil || strings.TrimSpace(part.ImageURL.URL) == "" {
				continue
			}
			if block, ok := contentBlockFromAttachmentURL(part.ImageURL.URL); ok {
				blocks = append(blocks, block)
			}
		}
	}
	return blocks
}

// prependSystemContentToBlocks puts systemContent at the start of blocks. If the first
// block is text, it adds systemContent and a blank line to the start of that text.
// If not, it adds a new text block at the start.
func prependSystemContentToBlocks(systemContent string, blocks []anthropic.ContentBlockParamUnion) []anthropic.ContentBlockParamUnion {
	if len(blocks) == 0 {
		return []anthropic.ContentBlockParamUnion{anthropic.NewTextBlock(systemContent)}
	}
	if blocks[0].OfText != nil {
		blocks[0].OfText.Text = systemContent + "\n\n" + blocks[0].OfText.Text
		return blocks
	}
	return append([]anthropic.ContentBlockParamUnion{anthropic.NewTextBlock(systemContent)}, blocks...)
}

// contentBlockFromAttachmentURL makes an Anthropic content block from an attachment URL.
// A data URL gives an image block or a PDF block, from its MIME type and base64 data.
// A remote URL gives a PDF document block if its path has a .pdf extension, and an
// image block if not. It returns false for a data URL that it cannot parse or that
// has an unsupported MIME type.
func contentBlockFromAttachmentURL(url string) (anthropic.ContentBlockParamUnion, bool) {
	if strings.HasPrefix(url, "data:") {
		mimeType, data, ok := parseDataURL(url)
		if !ok {
			debuglog.Debug(debuglog.Basic, "contentBlockFromAttachmentURL: failed to parse data URL")
			return anthropic.ContentBlockParamUnion{}, false
		}
		if strings.EqualFold(mimeType, "application/pdf") {
			return anthropic.NewDocumentBlock(anthropic.Base64PDFSourceParam{Data: data}), true
		}
		if normalized := normalizeImageMimeType(mimeType); normalized != "" {
			return anthropic.NewImageBlockBase64(normalized, data), true
		}
		debuglog.Debug(debuglog.Basic, "contentBlockFromAttachmentURL: unsupported MIME type %s", mimeType)
		return anthropic.ContentBlockParamUnion{}, false
	}
	if isPDFURL(url) {
		return anthropic.NewDocumentBlock(anthropic.URLPDFSourceParam{URL: url}), true
	}
	return anthropic.NewImageBlock(anthropic.URLImageSourceParam{URL: url}), true
}

// parseDataURL returns the MIME type and the base64 data of an RFC 2397 data URL.
// It returns ok=false for a data URL that is not base64-encoded.
func parseDataURL(value string) (mimeType string, data string, ok bool) {
	if !strings.HasPrefix(value, "data:") {
		return "", "", false
	}
	withoutPrefix := strings.TrimPrefix(value, "data:")
	parts := strings.SplitN(withoutPrefix, ",", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	meta := strings.TrimSpace(parts[0])
	data = strings.TrimSpace(parts[1])
	if data == "" {
		return "", "", false
	}
	metaParts := strings.Split(meta, ";")
	mimeType = strings.TrimSpace(metaParts[0])
	if mimeType == "" {
		return "", "", false
	}
	hasBase64 := false
	for _, part := range metaParts[1:] {
		if strings.EqualFold(strings.TrimSpace(part), "base64") {
			hasBase64 = true
			break
		}
	}
	if !hasBase64 {
		debuglog.Debug(debuglog.Basic, "parseDataURL: data URL without base64 encoding is not supported")
		return "", "", false
	}
	return mimeType, data, true
}

// normalizeImageMimeType returns the standard form of an image MIME type, or an empty
// string if the Anthropic API does not accept the type. The API accepts image/jpeg,
// image/png, image/gif, and image/webp.
// See: https://platform.claude.com/docs/en/build-with-claude/vision
func normalizeImageMimeType(mimeType string) string {
	switch strings.ToLower(strings.TrimSpace(mimeType)) {
	case "image/jpg", "image/jpeg":
		return "image/jpeg"
	case "image/png":
		return "image/png"
	case "image/gif":
		return "image/gif"
	case "image/webp":
		return "image/webp"
	default:
		return ""
	}
}

// isPDFURL returns true if the URL path has a .pdf extension. It does not send a
// request to read the Content-Type header, so it does not find a PDF at a path
// with no extension, such as /documents/12345.
func isPDFURL(url string) bool {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return false
	}
	return strings.EqualFold(path.Ext(parsedURL.Path), ".pdf")
}
