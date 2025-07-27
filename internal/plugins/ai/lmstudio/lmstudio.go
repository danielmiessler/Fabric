package lmstudio

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/danielmiessler/fabric/internal/chat"

	"github.com/danielmiessler/fabric/internal/domain"
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
	ret.PluginBase = &plugins.PluginBase{
		Name:            vendorName,
		EnvNamePrefix:   plugins.BuildEnvVariablePrefix(vendorName),
		ConfigureCustom: configureCustom,
	}
	ret.ApiUrl = ret.AddSetupQuestionCustom("API URL", true,
		fmt.Sprintf("Enter your %v URL (as a reminder, it is usually %v')", vendorName, defaultBaseUrl))
	return
}

// Client represents the LM Studio client.
type Client struct {
	*plugins.PluginBase
	ApiUrl     *plugins.SetupQuestion
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
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	models := make([]string, len(result.Data))
	for i, model := range result.Data {
		models[i] = model.ID
	}

	return models, nil
}

func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string) (err error) {
	url := fmt.Sprintf("%s/chat/completions", c.ApiUrl.Value)

	payload := map[string]interface{}{
		"messages": msgs,
		"model":    opts.Model,
		"stream":   true, // Enable streaming
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf("failed to marshal payload: %w", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
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
			err = fmt.Errorf("error reading response: %w", err)
			return
		}

		if len(line) == 0 {
			continue
		}

		if bytes.HasPrefix(line, []byte("data: ")) {
			line = bytes.TrimPrefix(line, []byte("data: "))
		}

		if string(line) == "[DONE]" {
			break
		}

		var result map[string]interface{}
		if err = json.Unmarshal(line, &result); err != nil {
			continue
		}

		var choices []interface{}
		var ok bool
		if choices, ok = result["choices"].([]interface{}); !ok || len(choices) == 0 {
			continue
		}

		var delta map[string]interface{}
		if delta, ok = choices[0].(map[string]interface{})["delta"].(map[string]interface{}); !ok {
			continue
		}

		var content string
		if content, _ = delta["content"].(string); content != "" {
			channel <- content
		}
	}

	return
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (content string, err error) {
	url := fmt.Sprintf("%s/chat/completions", c.ApiUrl.Value)

	payload := map[string]interface{}{
		"messages": msgs,
		"model":    opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf("failed to marshal payload: %w", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf("failed to decode response: %w", err)
		return
	}

	var choices []interface{}
	var ok bool
	if choices, ok = result["choices"].([]interface{}); !ok || len(choices) == 0 {
		err = fmt.Errorf("invalid response format: missing or empty choices")
		return
	}

	var message map[string]interface{}
	if message, ok = choices[0].(map[string]interface{})["message"].(map[string]interface{}); !ok {
		err = fmt.Errorf("invalid response format: missing message in first choice")
		return
	}

	if content, ok = message["content"].(string); !ok {
		err = fmt.Errorf("invalid response format: missing or non-string content in message")
		return
	}

	return
}

func (c *Client) Complete(ctx context.Context, prompt string, opts *domain.ChatOptions) (text string, err error) {
	url := fmt.Sprintf("%s/completions", c.ApiUrl.Value)

	payload := map[string]interface{}{
		"prompt": prompt,
		"model":  opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf("failed to marshal payload: %w", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf("failed to decode response: %w", err)
		return
	}

	var choices []interface{}
	var ok bool
	if choices, ok = result["choices"].([]interface{}); !ok || len(choices) == 0 {
		err = fmt.Errorf("invalid response format: missing or empty choices")
		return
	}

	if text, ok = choices[0].(map[string]interface{})["text"].(string); !ok {
		err = fmt.Errorf("invalid response format: missing or non-string text in first choice")
		return
	}

	return
}

func (c *Client) GetEmbeddings(ctx context.Context, input string, opts *domain.ChatOptions) (embeddings []float64, err error) {
	url := fmt.Sprintf("%s/embeddings", c.ApiUrl.Value)

	payload := map[string]interface{}{
		"input": input,
		"model": opts.Model,
		// Add other options from opts if supported by LM Studio
	}

	var jsonPayload []byte
	if jsonPayload, err = json.Marshal(payload); err != nil {
		err = fmt.Errorf("failed to marshal payload: %w", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonPayload)); err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.HttpClient.Do(req); err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	var result struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		err = fmt.Errorf("failed to decode response: %w", err)
		return
	}

	if len(result.Data) == 0 {
		err = fmt.Errorf("no embeddings returned")
		return
	}

	embeddings = result.Data[0].Embedding
	return
}

func (c *Client) NeedsRawMode(modelName string) bool {
	return false
}
