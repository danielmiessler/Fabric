package perplexity

import (
	"context"
	"encoding/json"
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

func (c *Client) HandleSchema(opts *domain.ChatOptions) (err error) {
	// Perplexity supports JSON Schema structured outputs.
	// We will add the schema to the request options in Send/SendStream.
	return nil
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

	// If schema content is provided, parse it and add the JSON schema response format option
	if opts.SchemaContent != "" {
		var schema map[string]interface{}
		if err := json.Unmarshal([]byte(opts.SchemaContent), &schema); err != nil {
			return "", fmt.Errorf("failed to parse schema content: %w", err)
		}
		requestOptions = append(requestOptions, perplexity.WithJSONSchemaResponseFormat(schema))
	}

	request := perplexity.NewCompletionRequest(requestOptions...)

	// Use SendCompletionRequest method from perplexity-go library
	resp, err := c.client.SendCompletionRequest(request) // Pass request directly
	if err != nil {
		return "", fmt.Errorf("perplexity API request failed: %w", err) // Corrected capitalization
	}

	content := resp.GetLastContent()
	citations := resp.GetSearchResults()

	// Handle structured output with citations embedded in JSON
	if opts.SchemaContent != "" && len(citations) > 0 {
		var contentMap map[string]interface{}
		// Attempt to unmarshal existing content as JSON
		if err := json.Unmarshal([]byte(content), &contentMap); err == nil {
			citationData := make([]map[string]string, len(citations))
			for i, citation := range citations {
				citationData[i] = map[string]string{
					"title": citation.Title,
					"url":   citation.URL,
				}
			}
			contentMap["citations"] = citationData                      // Add citations to the JSON object
			newContent, err := json.MarshalIndent(contentMap, "", "  ") // Use MarshalIndent for readability
			if err == nil {
				content = string(newContent)
			} else {
				// Fallback if marshaling fails, append as plain text
				content += "\n\n# CITATIONS\n\n"
				for i, citation := range citations {
					content += fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
				}
			}
		} else {
			// Fallback if unmarshaling fails (content is not valid JSON), append as plain text
			content += "\n\n# CITATIONS\n\n"
			for i, citation := range citations {
				content += fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
			}
		}
	} else if len(citations) > 0 { // Original logic for non-structured output
		content += "\n\n# CITATIONS\n\n"
		for i, citation := range citations {
			content += fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
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

	// If schema content is provided, parse it and add the JSON schema response format option
	if opts.SchemaContent != "" {
		var schema map[string]interface{}
		if err := json.Unmarshal([]byte(opts.SchemaContent), &schema); err != nil {
			close(channel)
			return fmt.Errorf("failed to parse schema content: %w", err)
		}
		requestOptions = append(requestOptions, perplexity.WithJSONSchemaResponseFormat(schema))
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
		var fullContent string // Accumulate streamed content

		for resp := range responseChan {
			lastResponse = &resp
			if len(resp.Choices) > 0 {
				content := ""
				if resp.Choices[0].Delta.Content != "" {
					content = resp.Choices[0].Delta.Content
				} else if resp.Choices[0].Message.Content != "" {
					content = resp.Choices[0].Message.Content
				}
				if content != "" {
					fullContent += content // Accumulate
					channel <- content
				}
			}
		}

		// After the stream finishes, handle citations
		if lastResponse != nil {
			citations := lastResponse.GetSearchResults()
			if len(citations) > 0 {
				if opts.SchemaContent != "" {
					// Structured output: Embed citations into the JSON object
					var contentMap map[string]interface{}
					if err := json.Unmarshal([]byte(fullContent), &contentMap); err == nil {
						citationData := make([]map[string]string, len(citations))
						for i, citation := range citations {
							citationData[i] = map[string]string{
								"title": citation.Title,
								"url":   citation.URL,
							}
						}
						contentMap["citations"] = citationData
						newContent, err := json.MarshalIndent(contentMap, "", "  ")
						if err == nil {
							// Send the complete JSON object with citations
							channel <- string(newContent)
						} else {
							// Fallback to plain text if JSON marshaling fails
							channel <- "\n\n# CITATIONS\n\n"
							for i, citation := range citations {
								channel <- fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
							}
						}
					} else {
						// Fallback to plain text if unmarshaling fails (fullContent not valid JSON)
						channel <- "\n\n# CITATIONS\n\n"
						for i, citation := range citations {
							channel <- fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
						}
					}
				} else {
					// Non-structured output: Append citations as plain text
					channel <- "\n\n# CITATIONS\n\n"
					for i, citation := range citations {
						channel <- fmt.Sprintf("- [%d] %s %s\n", i+1, citation.Title, citation.URL)
					}
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
