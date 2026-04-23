package openai

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/pagination"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
	"github.com/openai/openai-go/shared/constant"
)

func NewClient() (ret *Client) {
	return NewClientCompatibleWithResponses("OpenAI", "https://api.openai.com/v1", true, nil)
}

func NewClientCompatible(vendorName string, defaultBaseUrl string, configureCustom func() error) (ret *Client) {
	ret = NewClientCompatibleNoSetupQuestions(vendorName, configureCustom)

	ret.ApiKey = ret.AddSetupQuestion("API Key", true)
	ret.ApiBaseURL = ret.AddSetupQuestion("API Base URL", false)
	ret.ApiBaseURL.Value = defaultBaseUrl

	return
}

func NewClientCompatibleWithResponses(vendorName string, defaultBaseUrl string, implementsResponses bool, configureCustom func() error) (ret *Client) {
	ret = NewClientCompatibleNoSetupQuestions(vendorName, configureCustom)

	ret.ApiKey = ret.AddSetupQuestion("API Key", true)
	ret.ApiBaseURL = ret.AddSetupQuestion("API Base URL", false)
	ret.ApiBaseURL.Value = defaultBaseUrl
	ret.ImplementsResponses = implementsResponses

	return
}

func NewClientCompatibleNoSetupQuestions(vendorName string, configureCustom func() error) (ret *Client) {
	ret = &Client{}

	if configureCustom == nil {
		configureCustom = ret.configure
	}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, configureCustom)

	return
}

type Client struct {
	*plugins.PluginBase
	ApiKey              *plugins.SetupQuestion
	ApiBaseURL          *plugins.SetupQuestion
	ApiClient           *openai.Client
	ImplementsResponses bool // Whether this provider supports the Responses API
	httpClient          *http.Client
	// webSearchToolName, when non-empty, overrides the default
	// "web_search_preview" tool name emitted on the Responses API.
	// Used by OpenAI-compatible providers whose upstream API expects a
	// different tool type string (xAI expects "web_search").
	webSearchToolName string
	// enableXSearch, when true, appends an additional "x_search" tool
	// entry alongside the web search tool when Search is enabled.
	// This is an xAI-specific live search grounding tool.
	enableXSearch bool
}

// SetResponsesAPIEnabled configures whether to use the Responses API
func (o *Client) SetResponsesAPIEnabled(enabled bool) {
	o.ImplementsResponses = enabled
}

// SetWebSearchToolName overrides the default "web_search_preview" tool
// name emitted on the Responses API when Search is enabled. Pass an empty
// string to keep the OpenAI default. Non-OpenAI providers (for example,
// xAI) may require "web_search" instead.
func (o *Client) SetWebSearchToolName(name string) {
	o.webSearchToolName = name
}

// SetEnableXSearch toggles whether an additional xAI "x_search" tool
// entry is appended when Search is enabled. Non-xAI providers should
// leave this false.
func (o *Client) SetEnableXSearch(enabled bool) {
	o.enableXSearch = enabled
}

// checkImageGenerationCompatibility warns if the model doesn't support image generation
func checkImageGenerationCompatibility(model string) {
	if !supportsImageGeneration(model) {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf(i18n.T("openai_warning_model_no_image_generation"),
			model, strings.Join(ImageGenerationSupportedModels, ", ")))
	}
}

func (o *Client) configure() (ret error) {
	opts := []option.RequestOption{option.WithAPIKey(o.ApiKey.Value)}
	if o.ApiBaseURL.Value != "" {
		opts = append(opts, option.WithBaseURL(o.ApiBaseURL.Value))
	}
	client := openai.NewClient(opts...)
	o.ApiClient = &client

	// Initialize HTTP client for direct API calls (reused across requests)
	o.httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	return
}

func (o *Client) ListModels(ctx context.Context) (ret []string, err error) {
	var page *pagination.Page[openai.Model]
	if page, err = o.ApiClient.Models.List(ctx); err == nil {
		for _, mod := range page.Data {
			ret = append(ret, mod.ID)
		}
		// SDK succeeded - return the result even if empty
		return ret, nil
	}

	// SDK returned an error - fall back to direct API fetch.
	// Some providers (e.g., GitHub Models) return non-standard response formats
	// that the SDK fails to parse.
	debuglog.Debug(debuglog.Basic, "SDK Models.List failed for %s: %v, falling back to direct API fetch\n", o.GetName(), err)
	return FetchModelsDirectly(ctx, o.ApiBaseURL.Value, o.ApiKey.Value, o.GetName(), o.httpClient)
}

func (o *Client) SendStream(
	ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) (err error) {
	// Use Responses API for OpenAI, Chat Completions API for other providers
	if o.supportsResponsesAPI() {
		return o.sendStreamResponses(ctx, msgs, opts, channel)
	}
	return o.sendStreamChatCompletions(ctx, msgs, opts, channel)
}

func (o *Client) sendStreamResponses(
	ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate,
) (err error) {
	defer close(channel)

	req := o.buildResponseParams(msgs, opts)
	stream := o.ApiClient.Responses.NewStreaming(ctx, req)
	for stream.Next() {
		event := stream.Current()
		switch event.Type {
		case string(constant.ResponseOutputTextDelta("").Default()):
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: event.AsResponseOutputTextDelta().Delta,
			}
		case string(constant.ResponseOutputTextDone("").Default()):
			// The Responses API sends the full text again in the
			// final "done" event. Since we've already streamed all
			// delta chunks above, sending it would duplicate the
			// output. Ignore it here to prevent doubled results.
			continue
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

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	// Use Responses API for OpenAI, Chat Completions API for other providers
	if o.supportsResponsesAPI() {
		return o.sendResponses(ctx, msgs, opts)
	}
	return o.sendChatCompletions(ctx, msgs, opts)
}

func (o *Client) sendResponses(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	// Warn if model doesn't support image generation when image file is specified
	if opts.ImageFile != "" {
		checkImageGenerationCompatibility(opts.Model)
	}

	// Validate model supports image generation if image file is specified
	if opts.ImageFile != "" && !supportsImageGeneration(opts.Model) {
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_model_no_image_generation"), opts.Model, strings.Join(ImageGenerationSupportedModels, ", ")))
	}

	req := o.buildResponseParams(msgs, opts)

	var resp *responses.Response
	if resp, err = o.ApiClient.Responses.New(ctx, req); err != nil {
		return
	}

	// Extract and save images if requested
	if err = o.extractAndSaveImages(resp, opts); err != nil {
		return
	}

	ret = o.extractText(resp)
	return
}

// supportsResponsesAPI determines if the provider supports the new Responses API
func (o *Client) supportsResponsesAPI() bool {
	return o.ImplementsResponses
}

func (o *Client) NeedsRawMode(modelName string) bool {
	openaiModelsPrefixes := []string{
		"glm",
		"gpt-5",
		"o1",
		"o3",
		"o4",
	}
	openAIModelsNeedingRaw := []string{
		"gpt-4o-mini-search-preview",
		"gpt-4o-mini-search-preview-2025-03-11",
		"gpt-4o-search-preview",
		"gpt-4o-search-preview-2025-03-11",
	}
	for _, prefix := range openaiModelsPrefixes {
		if strings.HasPrefix(modelName, prefix) {
			return true
		}
	}
	return slices.Contains(openAIModelsNeedingRaw, modelName)
}

func parseReasoningEffort(level domain.ThinkingLevel) (shared.ReasoningEffort, bool) {
	switch domain.ThinkingLevel(strings.ToLower(string(level))) {
	case domain.ThinkingLow:
		return shared.ReasoningEffortLow, true
	case domain.ThinkingMedium:
		return shared.ReasoningEffortMedium, true
	case domain.ThinkingHigh:
		return shared.ReasoningEffortHigh, true
	default:
		return "", false
	}
}

func (o *Client) buildResponseParams(
	inputMsgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions,
) (ret responses.ResponseNewParams) {

	items := make([]responses.ResponseInputItemUnionParam, len(inputMsgs))
	for i, msgPtr := range inputMsgs {
		msg := *msgPtr
		if strings.Contains(opts.Model, "deepseek") && len(inputMsgs) == 1 && msg.Role == chat.ChatMessageRoleSystem {
			msg.Role = chat.ChatMessageRoleUser
		}
		items[i] = convertMessage(msg)
	}

	ret = responses.ResponseNewParams{
		Model: shared.ResponsesModel(opts.Model),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: items,
		},
	}

	// Add tools if enabled
	var tools []responses.ToolUnionParam

	// Add web search tool if enabled. The default tool name is OpenAI's
	// "web_search_preview", but providers may override it (for example,
	// xAI's Responses API requires "web_search").
	if opts.Search {
		searchToolName := responses.WebSearchToolType("web_search_preview")
		if o.webSearchToolName != "" {
			searchToolName = responses.WebSearchToolType(o.webSearchToolName)
		}
		webSearchTool := responses.ToolParamOfWebSearchPreview(searchToolName)

		// Add user location if provided. Only attach when the caller
		// asked for it; xAI rejects unexpected location payloads.
		if opts.SearchLocation != "" {
			webSearchTool.OfWebSearchPreview.UserLocation = responses.WebSearchToolUserLocationParam{
				Type:     "approximate",
				Timezone: openai.String(opts.SearchLocation),
			}
		}

		tools = append(tools, webSearchTool)

		// Append xAI's live "x_search" tool when the provider opts in.
		// The xAI Responses API accepts a bare {"type":"x_search"}
		// entry with no other required fields. We reuse the SDK's
		// WebSearchToolParam as a minimal container since every other
		// field is omitzero and will be elided during JSON marshalling.
		if o.enableXSearch {
			xSearchTool := responses.ToolUnionParam{
				OfWebSearchPreview: &responses.WebSearchToolParam{
					Type: responses.WebSearchToolType("x_search"),
				},
			}
			tools = append(tools, xSearchTool)
		}
	}

	// Add image generation tool if needed
	tools = o.addImageGenerationTool(opts, tools)

	if len(tools) > 0 {
		ret.Tools = tools
	}

	if eff, ok := parseReasoningEffort(opts.Thinking); ok {
		ret.Reasoning = shared.ReasoningParam{Effort: eff}
	}

	if !opts.Raw {
		ret.Temperature = openai.Float(opts.Temperature)
		if opts.TopP != 0 {
			ret.TopP = openai.Float(opts.TopP)
		}
		if opts.MaxTokens != 0 {
			ret.MaxOutputTokens = openai.Int(int64(opts.MaxTokens))
		}

		// Add parameters not officially supported by Responses API as extra fields
		extraFields := make(map[string]any)
		if opts.PresencePenalty != 0 {
			extraFields["presence_penalty"] = opts.PresencePenalty
		}
		if opts.FrequencyPenalty != 0 {
			extraFields["frequency_penalty"] = opts.FrequencyPenalty
		}
		if opts.Seed != 0 {
			extraFields["seed"] = opts.Seed
		}
		if len(extraFields) > 0 {
			ret.SetExtraFields(extraFields)
		}
	}
	return
}

// BuildResponseParams exposes the shared Responses API request builder so
// auth-specialized vendors can reuse Fabric's OpenAI-compatible request shape.
func (o *Client) BuildResponseParams(
	inputMsgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions,
) responses.ResponseNewParams {
	return o.buildResponseParams(inputMsgs, opts)
}

func convertMessage(msg chat.ChatCompletionMessage) responses.ResponseInputItemUnionParam {
	result := convertMessageCommon(msg)
	role := responses.EasyInputMessageRole(result.Role)

	if result.HasMultiContent {
		var parts []responses.ResponseInputContentUnionParam
		for _, p := range result.MultiContent {
			switch p.Type {
			case chat.ChatMessagePartTypeText:
				parts = append(parts, responses.ResponseInputContentParamOfInputText(p.Text))
			case chat.ChatMessagePartTypeImageURL:
				part := responses.ResponseInputContentParamOfInputImage(responses.ResponseInputImageDetailAuto)
				if part.OfInputImage != nil {
					part.OfInputImage.ImageURL = openai.String(p.ImageURL.URL)
				}
				parts = append(parts, part)
			}
		}
		contentList := responses.ResponseInputMessageContentListParam(parts)
		return responses.ResponseInputItemParamOfMessage(contentList, role)
	}
	return responses.ResponseInputItemParamOfMessage(result.Content, role)
}

func (o *Client) extractText(resp *responses.Response) (ret string) {
	var textParts []string
	var citations []string
	citationMap := make(map[string]bool) // To avoid duplicate citations

	for _, item := range resp.Output {
		if item.Type == "message" {
			for _, c := range item.Content {
				if c.Type == "output_text" {
					outputText := c.AsOutputText()
					textParts = append(textParts, outputText.Text)

					// Extract citations from annotations
					for _, annotation := range outputText.Annotations {
						if annotation.Type == "url_citation" {
							urlCitation := annotation.AsURLCitation()
							citationKey := urlCitation.URL + "|" + urlCitation.Title
							if !citationMap[citationKey] {
								citationMap[citationKey] = true
								citationText := fmt.Sprintf("- [%s](%s)", urlCitation.Title, urlCitation.URL)
								citations = append(citations, citationText)
							}
						}
					}
				}
			}
			break
		}
	}

	ret = strings.Join(textParts, "")

	// Append citations if any were found
	if len(citations) > 0 {
		ret += "\n\n## Sources\n\n" + strings.Join(citations, "\n")
	}

	return
}

// ExtractText exposes the shared Responses API text extraction logic so other
// vendors can reuse Fabric's response formatting and citation handling.
func (o *Client) ExtractText(resp *responses.Response) string {
	return o.extractText(resp)
}
