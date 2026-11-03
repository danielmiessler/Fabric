package ollama

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/util"
)

const defaultBaseUrl = "http://localhost:11434"

func NewClient() (ret *Client) {
	vendorName := "Ollama"
	ret = &Client{}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, ret.configure)

	ret.ApiUrl = ret.AddSetupQuestionCustom("API URL", true,
		fmt.Sprintf(i18n.T("lmstudio_api_url_question"), vendorName, defaultBaseUrl))
	ret.ApiUrl.Value = defaultBaseUrl
	ret.ApiKey = ret.PluginBase.AddSetupQuestion("API key", false)
	ret.ApiKey.Value = ""
	ret.ApiHttpTimeout = ret.AddSetupQuestionCustom("HTTP Timeout", true,
		i18n.T("ollama_http_timeout_question"))
	ret.ApiHttpTimeout.Value = "20m"

	return
}

type Client struct {
	*plugins.PluginBase
	ApiUrl         *plugins.SetupQuestion
	ApiKey         *plugins.SetupQuestion
	apiUrl         *url.URL
	ApiHttpTimeout *plugins.SetupQuestion
	httpClient     *http.Client
}

type transport_sec struct {
	underlyingTransport http.RoundTripper
	ApiKey              *plugins.SetupQuestion
}

func (t *transport_sec) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.ApiKey.Value != "" {
		req.Header.Add("Authorization", "Bearer "+t.ApiKey.Value)
	}
	return t.underlyingTransport.RoundTrip(req)
}

// imageFetchTimeout is the time limit for one image fetch. It includes
// the dial and the redirects.
const imageFetchTimeout = 30 * time.Second

// maxImageSize is the maximum size of an image that loadImageBytes reads.
const maxImageSize = 20 << 20

// imageDialControl checks the IP address of each connection that
// loadImageBytes makes. A test replaces it to connect to an httptest
// server.
var imageDialControl = util.DenyNonPublicAddress

// IsConfigured returns true only if OLLAMA_API_URL environment variable is explicitly set
func (o *Client) IsConfigured() bool {
	return os.Getenv("OLLAMA_API_URL") != ""
}

func (o *Client) configure() (err error) {
	if o.apiUrl, err = url.Parse(o.ApiUrl.Value); err != nil {
		fmt.Printf("%s\n", fmt.Sprintf(i18n.T("ollama_cannot_parse_url"), o.ApiUrl.Value, err))
		return
	}

	timeout := 20 * time.Minute

	if o.ApiHttpTimeout != nil {
		parsed, err := time.ParseDuration(o.ApiHttpTimeout.Value)
		if err == nil && o.ApiHttpTimeout.Value != "" {
			timeout = parsed
		} else if o.ApiHttpTimeout.Value != "" {
			fmt.Printf("%s\n", fmt.Sprintf(i18n.T("ollama_invalid_http_timeout_using_default"), o.ApiHttpTimeout.Value, err))
		}
	}

	o.httpClient = &http.Client{Timeout: timeout, Transport: &transport_sec{underlyingTransport: http.DefaultTransport, ApiKey: o.ApiKey}}

	return
}

// The types below are the parts of the Ollama API that Fabric uses. Fabric
// does not import github.com/ollama/ollama/api, because govulncheck reports
// all Ollama server advisories against each package of that module.
type ollamaMessage struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Images  [][]byte `json:"images,omitempty"`
}

type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Stream   *bool           `json:"stream,omitempty"`
	Options  map[string]any  `json:"options,omitempty"`
	Think    *bool           `json:"think,omitempty"`
}

type ollamaChatResponse struct {
	Message         ollamaMessage `json:"message"`
	Done            bool          `json:"done"`
	PromptEvalCount int           `json:"prompt_eval_count"`
	EvalCount       int           `json:"eval_count"`
	Error           string        `json:"error"`
}

// maxResponseLineSize is the maximum size of one JSON line from the Ollama server.
const maxResponseLineSize = 8 << 20

// StatusError is an error status from the Ollama server.
type StatusError struct {
	StatusCode int
	Message    string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("ollama: status %d: %s", e.StatusCode, e.Message)
}

// do sends a request to path on the Ollama server. It returns a
// *StatusError when the status is 400 or more.
func (o *Client) do(ctx context.Context, method, path string, data any) (resp *http.Response, err error) {
	var body io.Reader
	if data != nil {
		var b []byte
		if b, err = json.Marshal(data); err != nil {
			return
		}
		body = bytes.NewReader(b)
	}
	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, method, o.apiUrl.JoinPath(path).String(), body); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if resp, err = o.httpClient.Do(req); err != nil {
		return
	}
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 64<<10))
		var errBody struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(msg, &errBody) == nil && errBody.Error != "" {
			msg = []byte(errBody.Error)
		}
		return nil, &StatusError{StatusCode: resp.StatusCode, Message: strings.TrimSpace(string(msg))}
	}
	return
}

// chat sends req to /api/chat and calls fn for each JSON line of the
// response. A response with stream off has one line.
func (o *Client) chat(ctx context.Context, req *ollamaChatRequest, fn func(ollamaChatResponse)) (err error) {
	var resp *http.Response
	if resp, err = o.do(ctx, http.MethodPost, "api/chat", req); err != nil {
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(nil, maxResponseLineSize)
	for scanner.Scan() {
		var line ollamaChatResponse
		if err = json.Unmarshal(scanner.Bytes(), &line); err != nil {
			return
		}
		if line.Error != "" {
			return errors.New(line.Error)
		}
		fn(line)
	}
	return scanner.Err()
}

func (o *Client) ListModels(ctx context.Context) (ret []string, err error) {
	var resp *http.Response
	if resp, err = o.do(ctx, http.MethodGet, "api/tags", nil); err != nil {
		return
	}
	defer resp.Body.Close()

	var list struct {
		Models []struct {
			Model string `json:"model"`
		} `json:"models"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return
	}
	for _, mod := range list.Models {
		ret = append(ret, mod.Model)
	}
	return
}

func (o *Client) SendStream(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) (err error) {
	defer close(channel)

	var req ollamaChatRequest
	if req, err = o.createChatRequest(ctx, msgs, opts); err != nil {
		return
	}

	respFunc := func(resp ollamaChatResponse) {
		channel <- domain.StreamUpdate{
			Type:    domain.StreamTypeContent,
			Content: resp.Message.Content,
		}

		if resp.Done {
			channel <- domain.StreamUpdate{
				Type: domain.StreamTypeUsage,
				Usage: &domain.UsageMetadata{
					InputTokens:  resp.PromptEvalCount,
					OutputTokens: resp.EvalCount,
					TotalTokens:  resp.PromptEvalCount + resp.EvalCount,
				},
			}
		}
	}

	return o.chat(ctx, &req, respFunc)
}

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	bf := false

	var req ollamaChatRequest
	if req, err = o.createChatRequest(ctx, msgs, opts); err != nil {
		return
	}
	req.Stream = &bf

	if err = o.chat(ctx, &req, func(resp ollamaChatResponse) { ret = resp.Message.Content }); err != nil {
		debuglog.Debug(debuglog.Basic, "%s\n", fmt.Sprintf(i18n.T("ollama_chat_request_failed"), err))
	}
	return
}

func (o *Client) createChatRequest(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret ollamaChatRequest, err error) {
	// Some models, for example qwen3-coder and deepseek, return an empty response
	// when the only message is a system message. Send it as a user message instead.
	if len(msgs) == 1 && msgs[0].Role == chat.ChatMessageRoleSystem {
		copy := *msgs[0]
		copy.Role = chat.ChatMessageRoleUser
		msgs = []*chat.ChatCompletionMessage{&copy}
	}

	messages := make([]ollamaMessage, len(msgs))
	for i, message := range msgs {
		if messages[i], err = o.convertMessage(ctx, message); err != nil {
			return
		}
	}

	options := map[string]any{
		"temperature":       opts.Temperature,
		"presence_penalty":  opts.PresencePenalty,
		"frequency_penalty": opts.FrequencyPenalty,
		"top_p":             opts.TopP,
	}

	if opts.ModelContextLength != 0 {
		options["num_ctx"] = opts.ModelContextLength
	}

	if opts.MaxTokens > 0 {
		options["num_predict"] = opts.MaxTokens
	}

	if opts.Seed != 0 {
		options["seed"] = opts.Seed
	}

	ret = ollamaChatRequest{
		Model:    opts.Model,
		Messages: messages,
		Options:  options,
	}

	switch opts.Thinking {
	case domain.ThinkingOff:
		ret.Think = new(false)
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		ret.Think = new(true)
	}

	return
}

func (o *Client) convertMessage(ctx context.Context, message *chat.ChatCompletionMessage) (ret ollamaMessage, err error) {
	ret = ollamaMessage{Role: message.Role, Content: message.Content}

	if len(message.MultiContent) == 0 {
		return
	}

	textParts := make([]string, 0, len(message.MultiContent))
	if strings.TrimSpace(ret.Content) != "" {
		textParts = append(textParts, strings.TrimSpace(ret.Content))
	}

	for _, part := range message.MultiContent {
		switch part.Type {
		case chat.ChatMessagePartTypeText:
			if trimmed := strings.TrimSpace(part.Text); trimmed != "" {
				textParts = append(textParts, trimmed)
			}
		case chat.ChatMessagePartTypeImageURL:
			if part.ImageURL == nil || part.ImageURL.URL == "" {
				continue
			}
			var img []byte
			if img, err = o.loadImageBytes(ctx, part.ImageURL.URL); err != nil {
				return
			}
			ret.Images = append(ret.Images, img)
		}
	}

	ret.Content = strings.Join(textParts, "\n")
	return
}

func (o *Client) loadImageBytes(ctx context.Context, imageURL string) (ret []byte, err error) {
	if strings.HasPrefix(imageURL, "data:") {
		parts := strings.SplitN(imageURL, ",", 2)
		if len(parts) != 2 {
			err = errors.New(i18n.T("ollama_invalid_data_url_format"))
			return
		}
		if ret, err = base64.StdEncoding.DecodeString(parts[1]); err != nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("ollama_failed_decode_data_url"), err))
		}
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil); err != nil {
		return
	}

	// The image URL comes from the session content. Thus use a client that
	// connects only to public addresses. This client does not use
	// transport_sec, thus the Ollama API key does not go to the image host.
	var resp *http.Response
	if resp, err = util.NewPublicHTTPClient(imageFetchTimeout, imageDialControl).Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("ollama_failed_fetch_image"), imageURL, resp.Status))
		return
	}

	if ret, err = io.ReadAll(io.LimitReader(resp.Body, maxImageSize+1)); err != nil {
		return
	}
	if len(ret) > maxImageSize {
		ret = nil
		err = fmt.Errorf(i18n.T("ollama_image_too_large"), imageURL, maxImageSize)
	}
	return
}

func (o *Client) NeedsRawMode(modelName string) bool {
	ollamaSearchStrings := []string{
		"llama3",
		"llama2",
		"mistral",
	}
	for _, searchString := range ollamaSearchStrings {
		if strings.Contains(modelName, searchString) {
			return true
		}
	}
	return false
}
