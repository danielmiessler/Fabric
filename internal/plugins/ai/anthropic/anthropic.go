package anthropic

import (
	"context"
	"fmt"
	neturl "net/url"
	"os"
	"path"
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
		// The following are the current supported models
		string(anthropic.ModelClaudeSonnet4_6),
		string(anthropic.ModelClaudeOpus4_6),
		string(anthropic.ModelClaudeOpus4_5_20251101),
		string(anthropic.ModelClaudeOpus4_5),
		string(anthropic.ModelClaudeHaiku4_5),
		string(anthropic.ModelClaudeHaiku4_5_20251001),
		string(anthropic.ModelClaudeSonnet4_20250514),
		string(anthropic.ModelClaudeSonnet4_0),
		string(anthropic.ModelClaude4Sonnet20250514),
		string(anthropic.ModelClaudeSonnet4_5),
		string(anthropic.ModelClaudeSonnet4_5_20250929),
		string(anthropic.ModelClaudeOpus4_0),
		string(anthropic.ModelClaudeOpus4_20250514),
		string(anthropic.ModelClaude4Opus20250514),
		string(anthropic.ModelClaudeOpus4_1_20250805),
	}

	ret.modelBetas = map[string][]string{
		// See https://platform.claude.com/docs/en/build-with-claude/context-windows#1-m-token-context-window
		// Claude Opus 4.6, Sonnet 4.6, Sonnet 4.5, and Sonnet 4 support a 1-million token context window.

		// This list can change over time as Anthropic updates their models and beta features, so we maintain it separately from the main model list
		// for easier updates.

		// Claude Sonnet 4 variants (1M context support)
		string(anthropic.ModelClaudeSonnet4_20250514): {"context-1m-2025-08-07"},
		string(anthropic.ModelClaudeSonnet4_0):        {"context-1m-2025-08-07"},
		string(anthropic.ModelClaude4Sonnet20250514):  {"context-1m-2025-08-07"},

		// Claude Sonnet 4.5 variants (1M context support)
		string(anthropic.ModelClaudeSonnet4_5):          {"context-1m-2025-08-07"},
		string(anthropic.ModelClaudeSonnet4_5_20250929): {"context-1m-2025-08-07"},

		// Claude Sonnet 4.6 (1M context support)
		string(anthropic.ModelClaudeSonnet4_6): {"context-1m-2025-08-07"},

		// Claude Opus 4.5 and 4.6 variants (1M context support)
		string(anthropic.ModelClaudeOpus4_5):          {"context-1m-2025-08-07"},
		string(anthropic.ModelClaudeOpus4_6):          {"context-1m-2025-08-07"},
		string(anthropic.ModelClaudeOpus4_5_20251101): {"context-1m-2025-08-07"},
	}

	return
}

// IsConfigured returns true if the API key is configured
func (an *Client) IsConfigured() bool {
	// Check if API key is configured
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

func (an *Client) ListModels() (ret []string, err error) {
	return an.models, nil
}

func parseThinking(level domain.ThinkingLevel) (anthropic.ThinkingConfigParamUnion, bool) {
	lower := strings.ToLower(string(level))
	switch domain.ThinkingLevel(lower) {
	case domain.ThinkingOff:
		disabled := anthropic.NewThinkingConfigDisabledParam()
		return anthropic.ThinkingConfigParamUnion{OfDisabled: &disabled}, true
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		if budget, ok := domain.ThinkingBudgets[domain.ThinkingLevel(lower)]; ok {
			return anthropic.ThinkingConfigParamOfEnabled(budget), true
		}
	default:
		if tokens, err := strconv.ParseInt(lower, 10, 64); err == nil {
			if tokens >= 1 && tokens <= 10000 {
				return anthropic.ThinkingConfigParamOfEnabled(tokens), true
			}
		}
	}
	return anthropic.ThinkingConfigParamUnion{}, false
}

func (an *Client) SendStream(
	msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) (err error) {
	messages := an.toMessages(msgs)
	if len(messages) == 0 {
		close(channel)
		// No messages to send after normalization, consider this a non-error condition for streaming.
		return
	}

	ctx := context.Background()

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

		// Handle Content
		if event.Delta.Text != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.Delta.Text,
			}
		}

		// Handle Usage
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

	params = anthropic.MessageNewParams{
		Model:     anthropic.Model(opts.Model),
		MaxTokens: int64(an.maxTokens),
		Messages:  msgs,
	}

	// Only set one of Temperature or TopP as some models don't allow both
	// Always set temperature to ensure consistent behavior (Anthropic default is 1.0, Fabric default is 0.7)
	if opts.TopP != domain.DefaultTopP {
		// User explicitly set TopP, so use that instead of temperature
		params.TopP = anthropic.Opt(opts.TopP)
	} else {
		// Use temperature (always set to ensure Fabric's default of 0.7, not Anthropic's 1.0)
		params.Temperature = anthropic.Opt(opts.Temperature)
	}

	if opts.Search {
		// Build the web-search tool definition:
		webTool := anthropic.WebSearchTool20250305Param{
			Name:         webSearchToolName,
			Type:         webSearchToolType,
			CacheControl: anthropic.NewCacheControlEphemeralParam(),
		}

		if opts.SearchLocation != "" {
			webTool.UserLocation.Type = "approximate"
			webTool.UserLocation.Timezone = anthropic.Opt(opts.SearchLocation)
		}

		// Wrap it in the union:
		params.Tools = []anthropic.ToolUnionParam{
			{OfWebSearchTool20250305: &webTool},
		}
	}

	if t, ok := parseThinking(opts.Thinking); ok {
		params.Thinking = t
	}

	return
}

func (an *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (
	ret string, err error) {

	messages := an.toMessages(msgs)
	if len(messages) == 0 {
		// No messages to send after normalization, return empty string and no error.
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
	citationMap := make(map[string]bool) // To avoid duplicate citations

	for _, block := range message.Content {
		if block.Type == "text" && block.Text != "" {
			textParts = append(textParts, block.Text)

			// Extract citations from this text block
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

	// Append citations if any were found
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
	// Custom normalization for Anthropic:
	// - System messages become the first part of the first user message.
	// - Messages must alternate user/assistant.
	// - Skip empty messages.

	var anthropicMessages []anthropic.MessageParam
	var systemContent string

	// Note: Claude Code spoofing is now handled in buildMessageParams

	isFirstUserMessage := true
	lastRoleWasUser := false

	for _, msg := range msgs {
		if strings.TrimSpace(msg.Content) == "" && len(msg.MultiContent) == 0 {
			continue // Skip empty messages
		}

		switch msg.Role {
		case chat.ChatMessageRoleSystem:
			// Accumulate system content. It will be prepended to the first user message.
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
				isFirstUserMessage = false // System content now consumed
			}
			if lastRoleWasUser {
				// Enforce alternation: add a minimal assistant message if two user messages are consecutive.
				// This shouldn't happen with current chatter.go logic but is a safeguard.
				anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock("Okay.")))
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(blocks...))
			lastRoleWasUser = true
		case chat.ChatMessageRoleAssistant:
			// If the first message is an assistant message, and we have system content,
			// prepend a user message with the system content.
			if isFirstUserMessage && systemContent != "" {
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
				lastRoleWasUser = true
				isFirstUserMessage = false // System content now consumed
			} else if !lastRoleWasUser && len(anthropicMessages) > 0 {
				// Enforce alternation: add a minimal user message if two assistant messages are consecutive
				// or if an assistant message is first without prior system prompt handling.
				anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(an.defaultRequiredUserMessage)))
				lastRoleWasUser = true
			}
			anthropicMessages = append(anthropicMessages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(msg.Content)))
			lastRoleWasUser = false
		default:
			// Other roles (like 'meta') are ignored for Anthropic's message structure.
			continue
		}
	}

	// If only system content was provided, create a user message with it.
	if len(anthropicMessages) == 0 && systemContent != "" {
		anthropicMessages = append(anthropicMessages, anthropic.NewUserMessage(anthropic.NewTextBlock(systemContent)))
	}

	return anthropicMessages
}

// messageTextFromParts extracts and concatenates all text content from a message,
// combining both the Content field and any text parts in MultiContent.
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

// contentBlocksFromMessage converts a chat message into Anthropic content blocks,
// handling text content, image URLs (both data URLs and remote URLs), and PDF attachments.
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

// prependSystemContentToBlocks prepends system content to content blocks. If the first
// block is text, it merges the system content with it; otherwise, it prepends a new text block.
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

// contentBlockFromAttachmentURL converts an attachment URL into an Anthropic content block.
// For data URLs, it parses the MIME type and base64 data to create image or PDF blocks.
// For remote URLs, it creates URL-based image blocks, or PDF document blocks if the URL ends in .pdf.
// Returns the content block and true on success, or an empty block and false if unsupported.
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

// parseDataURL parses an RFC 2397 data URL, extracting the MIME type and base64-encoded data.
// Only base64-encoded data URLs are supported; URL-encoded data URLs will return ok=false.
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

// normalizeImageMimeType validates and normalizes image MIME types to those supported
// by the Anthropic API. Supported formats: image/jpeg, image/png, image/gif, image/webp.
// See: https://docs.anthropic.com/en/docs/build-with-claude/vision
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

// isPDFURL checks if a URL appears to point to a PDF based on its path extension.
// NOTE: This only checks the URL path extension (.pdf) and will not detect PDFs served
// from extension-less endpoints (e.g., /documents/12345) or based on Content-Type headers.
// This is an intentional limitation; callers should not assume this guarantees the
// remote resource is actually a PDF.
func isPDFURL(url string) bool {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return false
	}
	return strings.EqualFold(path.Ext(parsedURL.Path), ".pdf")
}
