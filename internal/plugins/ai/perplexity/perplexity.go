package perplexity

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	perplexity "github.com/sgaunet/perplexity-go/v2"

	"github.com/danielmiessler/fabric/internal/chat"
)

const (
	providerName = "Perplexity"
)

var models = []string{
	"r1-1776", "sonar", "sonar-pro", "sonar-reasoning", "sonar-reasoning-pro",
}

type Client struct {
	*plugins.PluginBase
	APIKey *plugins.SetupQuestion
	client *perplexity.Client
}

func NewClient() *Client {
	c := &Client{}
	c.PluginBase = &plugins.PluginBase{
		Name:            providerName,
		EnvNamePrefix:   plugins.BuildEnvVariablePrefix(providerName),
		ConfigureCustom: c.Configure, // Assign the Configure method
	}
	c.APIKey = c.AddSetupQuestion("API_KEY", true)
	return c
}

func (c *Client) Configure() error {
	// The PluginBase.Configure() is called by the framework if needed.
	// We only need to handle specific logic for this plugin.
	if c.APIKey.Value == "" {
		// Attempt to get from environment variable if not set by user during setup
		envKey := c.EnvNamePrefix + "API_KEY"
		apiKeyFromEnv := os.Getenv(envKey)
		if apiKeyFromEnv != "" {
			c.APIKey.Value = apiKeyFromEnv
		} else {
			return fmt.Errorf("%s API key not configured. Please set the %s environment variable or run 'fabric --setup %s'", providerName, envKey, providerName)
		}
	}
	c.client = perplexity.NewClient(c.APIKey.Value)
	return nil
}

func (c *Client) ListModels() ([]string, error) {
	// Perplexity API does not have a ListModels endpoint.
	// We return a predefined list.
	return models, nil
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if c.client == nil {
		if err := c.Configure(); err != nil {
			return "", fmt.Errorf("failed to configure Perplexity client: %w", err)
		}
	}

	var perplexityMessages []perplexity.Message
	for _, msg := range msgs {
		perplexityMessages = append(perplexityMessages, perplexity.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	requestOptions := []perplexity.CompletionRequestOption{
		perplexity.WithModel(opts.Model),
		perplexity.WithMessages(perplexityMessages),
	}
	if opts.MaxTokens > 0 {
		requestOptions = append(requestOptions, perplexity.WithMaxTokens(opts.MaxTokens))
	}
	if opts.Temperature > 0 { // Perplexity default is 1.0, only set if user specifies
		requestOptions = append(requestOptions, perplexity.WithTemperature(opts.Temperature))
	}
	if opts.TopP > 0 { // Perplexity default is not specified, typically 1.0
		requestOptions = append(requestOptions, perplexity.WithTopP(opts.TopP))
	}
	if opts.PresencePenalty != 0 {
		// Corrected: Pass float64 directly
		requestOptions = append(requestOptions, perplexity.WithPresencePenalty(opts.PresencePenalty))
	}
	if opts.FrequencyPenalty != 0 {
		// Corrected: Pass float64 directly
		requestOptions = append(requestOptions, perplexity.WithFrequencyPenalty(opts.FrequencyPenalty))
	}

	request := perplexity.NewCompletionRequest(requestOptions...)

	// Corrected: Use SendCompletionRequest method from perplexity-go library
	resp, err := c.client.SendCompletionRequest(request) // Pass request directly
	if err != nil {
		return "", fmt.Errorf("perplexity API request failed: %w", err) // Corrected capitalization
	}

	content := resp.GetLastContent()

	// Append citations if available
	citations := resp.GetCitations()
	if len(citations) > 0 {
		content += "\n\n# CITATIONS\n\n"
		for i, citation := range citations {
			content += fmt.Sprintf("- [%d] %s\n", i+1, citation)
		}
	}

	return content, nil
}

func (c *Client) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string) error {
	if c.client == nil {
		if err := c.Configure(); err != nil {
			close(channel) // Ensure channel is closed on error
			return fmt.Errorf("failed to configure Perplexity client: %w", err)
		}
	}

	var perplexityMessages []perplexity.Message
	for _, msg := range msgs {
		perplexityMessages = append(perplexityMessages, perplexity.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	requestOptions := []perplexity.CompletionRequestOption{
		perplexity.WithModel(opts.Model),
		perplexity.WithMessages(perplexityMessages),
		perplexity.WithStream(true), // Enable streaming
	}

	if opts.MaxTokens > 0 {
		requestOptions = append(requestOptions, perplexity.WithMaxTokens(opts.MaxTokens))
	}
	if opts.Temperature > 0 {
		requestOptions = append(requestOptions, perplexity.WithTemperature(opts.Temperature))
	}
	if opts.TopP > 0 {
		requestOptions = append(requestOptions, perplexity.WithTopP(opts.TopP))
	}
	if opts.PresencePenalty != 0 {
		// Corrected: Pass float64 directly
		requestOptions = append(requestOptions, perplexity.WithPresencePenalty(opts.PresencePenalty))
	}
	if opts.FrequencyPenalty != 0 {
		// Corrected: Pass float64 directly
		requestOptions = append(requestOptions, perplexity.WithFrequencyPenalty(opts.FrequencyPenalty))
	}

	request := perplexity.NewCompletionRequest(requestOptions...)

	responseChan := make(chan perplexity.CompletionResponse)
	var wg sync.WaitGroup // Use sync.WaitGroup
	wg.Add(1)

	go func() {
		err := c.client.SendSSEHTTPRequest(&wg, request, responseChan)
		if err != nil {
			// Log error, can't send to string channel directly.
			// Consider a mechanism to propagate this error if needed.
			debuglog.Log("perplexity streaming error: %v\n", err)
			// If the error occurs during stream setup, the channel might not have been closed by the receiver loop.
			// However, closing it here might cause a panic if the receiver loop also tries to close it.
			// close(channel) // Caution: Uncommenting this may cause panic, as channel is closed in the receiver goroutine.
		}
	}()

	go func() {
		defer close(channel) // Ensure the output channel is closed when this goroutine finishes
		var lastResponse *perplexity.CompletionResponse
		for resp := range responseChan {
			lastResponse = &resp
			if len(resp.Choices) > 0 {
				content := ""
				// Corrected: Check Delta.Content and Message.Content directly for non-emptiness
				// as Delta and Message are structs, not pointers, in perplexity.Choice
				if resp.Choices[0].Delta.Content != "" {
					content = resp.Choices[0].Delta.Content
				} else if resp.Choices[0].Message.Content != "" {
					content = resp.Choices[0].Message.Content
				}
				if content != "" {
					channel <- content
				}
			}
		}

		// Send citations at the end if available
		if lastResponse != nil {
			citations := lastResponse.GetCitations()
			if len(citations) > 0 {
				channel <- "\n\n# CITATIONS\n\n"
				for i, citation := range citations {
					channel <- fmt.Sprintf("- [%d] %s\n", i+1, citation)
				}
			}
		}
	}()

	return nil
}

func (c *Client) NeedsRawMode(modelName string) bool {
	return true
}

// Setup is called by the fabric CLI framework to guide the user through configuration.
func (c *Client) Setup() error {
	return c.PluginBase.Setup()
}

// GetName returns the name of the plugin.
func (c *Client) GetName() string {
	return c.PluginBase.Name
}

// GetEnvNamePrefix returns the environment variable prefix for the plugin.
// Corrected: Receiver name
func (c *Client) GetEnvNamePrefix() string {
	return c.PluginBase.EnvNamePrefix
}

// AddSetupQuestion adds a setup question to the plugin.
// This is a helper method, usually called from NewClient.
func (c *Client) AddSetupQuestion(text string, isSensitive bool) *plugins.SetupQuestion {
	return c.PluginBase.AddSetupQuestion(text, isSensitive)
}

// GetSetupQuestions returns the setup questions for the plugin.
// Corrected: Return the slice of setup questions from PluginBase
func (c *Client) GetSetupQuestions() []*plugins.SetupQuestion {
	return c.PluginBase.SetupQuestions
}
