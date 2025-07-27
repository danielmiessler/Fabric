package ollama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	ollamaapi "github.com/ollama/ollama/api"
	"github.com/samber/lo"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
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

	o.client = ollamaapi.NewClient(o.apiUrl, &http.Client{Timeout: timeout, Transport: &transport_sec{underlyingTransport: http.DefaultTransport, ApiKey: o.ApiKey}})

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
	req := o.createChatRequest(msgs, opts)

	respFunc := func(resp ollamaapi.ChatResponse) (streamErr error) {
		channel <- resp.Message.Content
		return
	}

	ctx := context.Background()

	if err = o.client.Chat(ctx, &req, respFunc); err != nil {
		return
	}

	close(channel)
	return
}

func (o *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {
	bf := false

	req := o.createChatRequest(msgs, opts)
	req.Stream = &bf

	respFunc := func(resp ollamaapi.ChatResponse) (streamErr error) {
		ret = resp.Message.Content
		return
	}

	if err = o.client.Chat(ctx, &req, respFunc); err != nil {
		fmt.Printf("FRED --> %s\n", err)
	}
	return
}

func (o *Client) createChatRequest(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret ollamaapi.ChatRequest) {
	messages := lo.Map(msgs, func(message *chat.ChatCompletionMessage, _ int) (ret ollamaapi.Message) {
		return ollamaapi.Message{Role: message.Role, Content: message.Content}
	})

	options := map[string]interface{}{
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

func (o *Client) NeedsRawMode(modelName string) bool {
	ollamaPrefixes := []string{
		"llama3",
		"llama2",
		"mistral",
	}
	for _, prefix := range ollamaPrefixes {
		if strings.HasPrefix(modelName, prefix) {
			return true
		}
	}
	return false
}
