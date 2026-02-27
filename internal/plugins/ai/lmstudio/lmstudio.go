package lmstudio

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

// NewClient creates a new LM Studio client with default configuration.
func NewClient() (ret *Client) {
	return NewClientCompatible("LM Studio", "http://localhost:1234/v1", nil)
}

// NewClientCompatible creates a new LM Studio client with custom configuration.
func NewClientCompatible(vendorName string, defaultBaseUrl string, configureCustom func() error) (ret *Client) {
	ret = &Client{}

	if configureCustom == nil {
		configureCustom = ret.configure
	}
	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, configureCustom)
	ret.ApiUrl = ret.AddSetupQuestionCustom("API URL", true,
		fmt.Sprintf(i18n.T("lmstudio_api_url_question"), vendorName, defaultBaseUrl))
	ret.ApiKey = ret.AddSetupQuestion("API key", false)
	return
}

// Client represents the LM Studio client.
type Client struct {
	*plugins.PluginBase
	ApiUrl     *plugins.SetupQuestion
	ApiKey     *plugins.SetupQuestion
	HttpClient *http.Client
}

// configure sets up the HTTP client.
func (c *Client) configure() error {
	c.HttpClient = &http.Client{}
	return nil
}

// ListModels returns a list of available models.
func (c *Client) ListModels() ([]string, error) {
	url := fmt.Sprintf("%s/models", c.ApiUrl.Value)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("lmstudio_failed_create_request"), err)
	}
	c.addAuthorizationHeader(req)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("lmstudio_failed_send_request"), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(i18n.T("lmstudio_unexpected_status_code"), resp.StatusCode)
	}

	var result struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf(i18n.T("lmstudio_failed_decode_response"), err)
	}

	models := make([]string, len(result.Data))
	for i, model := range result.Data {
		models[i] = model.ID
	}

	return models, nil
}

func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) (err error) {
	url := fmt.Sprintf("%s/chat/completions", c.ApiUrl.Value)

	payload := map[string]any{
		"messages": msgs,
		"model":    opts.Model,
		"stream":   true, // Enable streaming
		"stream_options": map[string]any{
			"include_usage": true,
		},
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_marshal_payload"), err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_create_request"), err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthorizationHeader(req)

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_send_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf(i18n.T("lmstudio_unexpected_status_code"), resp.StatusCode)
		return
	}

	defer close(channel)

	reader := bufio.NewReader(resp.Body)
	for {
		var line []byte
		if line, err = reader.ReadBytes('\n'); err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			err = fmt.Errorf(i18n.T("lmstudio_error_reading_response"), err)
			return
		}

		if len(line) == 0 {
			continue
		}

		if after, ok := bytes.CutPrefix(line, []byte("data: ")); ok {
			line = after
		}

		if string(bytes.TrimSpace(line)) == "[DONE]" {
			break
		}

		var result map[string]any
		if err = json.Unmarshal(line, &result); err != nil {
			continue
		}

		// Handle Usage
		if usage, ok := result["usage"].(map[string]any); ok {
			var metadata domain.UsageMetadata
			if val, ok := usage["prompt_tokens"].(float64); ok {
				metadata.InputTokens = int(val)
			}
			if val, ok := usage["completion_tokens"].(float64); ok {
				metadata.OutputTokens = int(val)
			}
			if val, ok := usage["total_tokens"].(float64); ok {
				metadata.TotalTokens = int(val)
			}
			channel <- domain.StreamUpdate{
				Type:  domain.StreamTypeUsage,
				Usage: &metadata,
			}
		}

		var choices []any
		var ok bool
		if choices, ok = result["choices"].([]any); !ok || len(choices) == 0 {
			continue
		}

		var delta map[string]any
		if delta, ok = choices[0].(map[string]any)["delta"].(map[string]any); !ok {
			continue
		}

		var content string
		if content, _ = delta["content"].(string); content != "" {
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: content,
			}
		}
	}

	return
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (content string, err error) {
	url := fmt.Sprintf("%s/chat/completions", c.ApiUrl.Value)

	payload := map[string]any{
		"messages": msgs,
		"model":    opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_marshal_payload"), err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_create_request"), err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthorizationHeader(req)

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_send_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf(i18n.T("lmstudio_unexpected_status_code"), resp.StatusCode)
		return
	}

	var result map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_decode_response"), err)
		return
	}

	var choices []any
	var ok bool
	if choices, ok = result["choices"].([]any); !ok || len(choices) == 0 {
		err = errors.New(i18n.T("lmstudio_invalid_response_missing_choices"))
		return
	}

	var message map[string]any
	if message, ok = choices[0].(map[string]any)["message"].(map[string]any); !ok {
		err = errors.New(i18n.T("lmstudio_invalid_response_missing_message"))
		return
	}

	if content, ok = message["content"].(string); !ok {
		err = errors.New(i18n.T("lmstudio_invalid_response_missing_content"))
		return
	}

	return
}

func (c *Client) Complete(ctx context.Context, prompt string, opts *domain.ChatOptions) (text string, err error) {
	url := fmt.Sprintf("%s/completions", c.ApiUrl.Value)

	payload := map[string]any{
		"prompt": prompt,
		"model":  opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_marshal_payload"), err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_create_request"), err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthorizationHeader(req)

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_send_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf(i18n.T("lmstudio_unexpected_status_code"), resp.StatusCode)
		return
	}

	var result map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_decode_response"), err)
		return
	}

	var choices []any
	var ok bool
	if choices, ok = result["choices"].([]any); !ok || len(choices) == 0 {
		err = errors.New(i18n.T("lmstudio_invalid_response_missing_choices"))
		return
	}

	if text, ok = choices[0].(map[string]any)["text"].(string); !ok {
		err = errors.New(i18n.T("lmstudio_invalid_response_missing_text"))
		return
	}

	return
}

func (c *Client) GetEmbeddings(ctx context.Context, input string, opts *domain.ChatOptions) (embeddings []float64, err error) {
	url := fmt.Sprintf("%s/embeddings", c.ApiUrl.Value)

	payload := map[string]any{
		"input": input,
		"model": opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_marshal_payload"), err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_create_request"), err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	c.addAuthorizationHeader(req)

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_send_request"), err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf(i18n.T("lmstudio_unexpected_status_code"), resp.StatusCode)
		return
	}

	var result struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf(i18n.T("lmstudio_failed_decode_response"), err)
		return
	}

	if len(result.Data) == 0 {
		err = errors.New(i18n.T("lmstudio_no_embeddings_returned"))
		return
	}

	embeddings = result.Data[0].Embedding
	return
}

func (c *Client) addAuthorizationHeader(req *http.Request) {
	if c.ApiKey == nil {
		return
	}
	apiKey := strings.TrimSpace(c.ApiKey.Value)
	if apiKey == "" {
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
}
