package ollama

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	ollamaapi "github.com/ollama/ollama/api"
)

const defaultBaseUrl = "http://localhost:11434"

func NewClient() (ret *Client) {
	vendorName := "Ollama"
	ret = &Client{}

	ret.PluginBase = &plugins.PluginBase{
		Name:            vendorName,
		EnvNamePrefix:   plugins.BuildEnvVariablePrefix(vendorName),
		ConfigureCustom: ret.configure,
	}

	ret.ApiUrl = ret.AddSetupQuestionCustom("API URL", true,
		"Enter your Ollama URL (as a reminder, it is usually http://localhost:11434')")
	ret.ApiUrl.Value = defaultBaseUrl
	ret.ApiKey = ret.PluginBase.AddSetupQuestion("API key", false)
	ret.ApiKey.Value = ""
	ret.ApiHttpTimeout = ret.AddSetupQuestionCustom("HTTP Timeout", true,
		"Specify HTTP timeout duration for Ollama requests (e.g. 30s, 5m, 1h)")
	ret.ApiHttpTimeout.Value = "20m"

	return
}

type Client struct {
	*plugins.PluginBase
	ApiUrl         *plugins.SetupQuestion
	ApiKey         *plugins.SetupQuestion
	apiUrl         *url.URL
	client         *ollamaapi.Client
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

// IsConfigured returns true only if OLLAMA_API_URL environment variable is explicitly set
func (o *Client) IsConfigured() bool {
	return os.Getenv("OLLAMA_API_URL") != ""
}

func (o *Client) configure() (err error) {
	if o.apiUrl, err = url.Parse(o.ApiUrl.Value); err != nil {
		fmt.Printf("cannot parse URL: %s: %v\n", o.ApiUrl.Value, err)
		return
	}

	timeout := 20 * time.Minute // Default timeout

	if o.ApiHttpTimeout != nil {
		parsed, err := time.ParseDuration(o.ApiHttpTimeout.Value)
		if err == nil && o.ApiHttpTimeout.Value != "" {
			timeout = parsed
		} else if o.ApiHttpTimeout.Value != "" {
			fmt.Printf("Invalid HTTP timeout format (%q), using default (20m): %v\n", o.ApiHttpTimeout.Value, err)
		}
	}

	o.httpClient = &http.Client{Timeout: timeout, Transport: &transport_sec{underlyingTransport: http.DefaultTransport, ApiKey: o.ApiKey}}
	o.client = ollamaapi.NewClient(o.apiUrl, o.httpClient)

	return
}

func (o *Client) ListModels() (ret []string, err error) {
	ctx := context.Background()

	var listResp *ollamaapi.ListResponse
	if listResp, err = o.client.List(ctx); err != nil {
		return
	}

	for _, mod := range listResp.Models {
		ret = append(ret, mod.Model)
	}
	return
}

func (o *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string) (err error) {
	ctx := context.Background()

	var req ollamaapi.ChatRequest
	if req, err = o.createChatRequest(ctx, msgs, opts); err != nil {
		return
	}

	respFunc := func(resp ollamaapi.ChatResponse) (streamErr error) {
		channel <- resp.Message.Content
		return
	}

	if err = o.client.Chat(ctx, &req, respFunc); err != nil {
		return
	}

	close(channel)
	return
}

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	bf := false

	var req ollamaapi.ChatRequest
	if req, err = o.createChatRequest(ctx, msgs, opts); err != nil {
		return
	}
	req.Stream = &bf

	respFunc := func(resp ollamaapi.ChatResponse) (streamErr error) {
		ret = resp.Message.Content
		return
	}

	if err = o.client.Chat(ctx, &req, respFunc); err != nil {
		debuglog.Debug(debuglog.Basic, "Ollama chat request failed: %v\n", err)
	}
	return
}

func (o *Client) createChatRequest(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret ollamaapi.ChatRequest, err error) {
	messages := make([]ollamaapi.Message, len(msgs))
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

	ret = ollamaapi.ChatRequest{
		Model:    opts.Model,
		Messages: messages,
		Options:  options,
	}
	return
}

func (o *Client) convertMessage(ctx context.Context, message *chat.ChatCompletionMessage) (ret ollamaapi.Message, err error) {
	ret = ollamaapi.Message{Role: message.Role, Content: message.Content}

	if len(message.MultiContent) == 0 {
		return
	}

	// Pre-allocate with capacity hint
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
			// Nil guard
			if part.ImageURL == nil || part.ImageURL.URL == "" {
				continue
			}
			var img []byte
			if img, err = o.loadImageBytes(ctx, part.ImageURL.URL); err != nil {
				return
			}
			ret.Images = append(ret.Images, ollamaapi.ImageData(img))
		}
	}

	ret.Content = strings.Join(textParts, "\n")
	return
}

func (o *Client) loadImageBytes(ctx context.Context, imageURL string) (ret []byte, err error) {
	// Handle data URLs (base64 encoded)
	if strings.HasPrefix(imageURL, "data:") {
		parts := strings.SplitN(imageURL, ",", 2)
		if len(parts) != 2 {
			err = fmt.Errorf("invalid data URL format")
			return
		}
		if ret, err = base64.StdEncoding.DecodeString(parts[1]); err != nil {
			err = fmt.Errorf("failed to decode data URL: %w", err)
		}
		return
	}

	// Handle HTTP URLs with context
	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil); err != nil {
		return
	}

	var resp *http.Response
	if resp, err = o.httpClient.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("failed to fetch image %s: %s", imageURL, resp.Status)
		return
	}

	ret, err = io.ReadAll(resp.Body)
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
