// Package bedrock provides a plugin to use Amazon Bedrock models.
// Supported models are defined in the MODELS variable.
// To add additional models, append them to the MODELS array. Models must support the Converse and ConverseStream operations
// Authentication uses the  AWS credential provider chain, similar.to the AWS CLI and SDKs
// https://docs.aws.amazon.com/sdkref/latest/guide/standardized-credentials.html
package bedrock

import (
	"context"
	"fmt"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"

	"github.com/danielmiessler/fabric/internal/chat"
)

const (
	userAgentKey   = "aiosc"
	userAgentValue = "fabric"
)

// Ensure BedrockClient implements the ai.Vendor interface
var _ ai.Vendor = (*BedrockClient)(nil)

// BedrockClient is a plugin to add support for Amazon Bedrock.
// It implements the plugins.Plugin interface and provides methods
// for interacting with AWS Bedrock's Converse and ConverseStream APIs.
type BedrockClient struct {
	*plugins.PluginBase
	runtimeClient      *bedrockruntime.Client
	controlPlaneClient *bedrock.Client

	bedrockRegion *plugins.SetupQuestion
}

// NewClient returns a new Bedrock plugin client
func NewClient() (ret *BedrockClient) {
	vendorName := "Bedrock"
	ret = &BedrockClient{}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		// Create a minimal client that will fail gracefully during configuration
		ret.PluginBase = &plugins.PluginBase{
			Name:          vendorName,
			EnvNamePrefix: plugins.BuildEnvVariablePrefix(vendorName),
			ConfigureCustom: func() error {
				return fmt.Errorf("unable to load AWS Config: %w", err)
			},
		}
		ret.bedrockRegion = ret.PluginBase.AddSetupQuestion("AWS Region", true)
		return
	}

	cfg.APIOptions = append(cfg.APIOptions, middleware.AddUserAgentKeyValue(userAgentKey, userAgentValue))

	runtimeClient := bedrockruntime.NewFromConfig(cfg)
	controlPlaneClient := bedrock.NewFromConfig(cfg)

	ret.PluginBase = &plugins.PluginBase{
		Name:            vendorName,
		EnvNamePrefix:   plugins.BuildEnvVariablePrefix(vendorName),
		ConfigureCustom: ret.configure,
	}

	ret.runtimeClient = runtimeClient
	ret.controlPlaneClient = controlPlaneClient

	ret.bedrockRegion = ret.PluginBase.AddSetupQuestion("AWS Region", true)

	if cfg.Region != "" {
		ret.bedrockRegion.Value = cfg.Region
	}

	return
}

// isValidAWSRegion validates AWS region format
func isValidAWSRegion(region string) bool {
	// Simple validation - AWS regions are typically 2-3 parts separated by hyphens
	// Examples: us-east-1, eu-west-1, ap-southeast-2
	if len(region) < 5 || len(region) > 30 {
		return false
	}
	// Basic pattern check for AWS region format
	return region != ""
}

// configure initializes the Bedrock clients with the specified AWS region.
// If no region is specified, the default region from AWS config is used.
func (c *BedrockClient) configure() error {
	if c.bedrockRegion.Value == "" {
		return nil // Use default region from AWS config
	}

	// Validate region format
	if !isValidAWSRegion(c.bedrockRegion.Value) {
		return fmt.Errorf("invalid AWS region: %s", c.bedrockRegion.Value)
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(c.bedrockRegion.Value))
	if err != nil {
		return fmt.Errorf("unable to load AWS Config with region %s: %w", c.bedrockRegion.Value, err)
	}

	cfg.APIOptions = append(cfg.APIOptions, middleware.AddUserAgentKeyValue(userAgentKey, userAgentValue))

	c.runtimeClient = bedrockruntime.NewFromConfig(cfg)
	c.controlPlaneClient = bedrock.NewFromConfig(cfg)

	return nil
}

// ListModels retrieves all available foundation models and inference profiles
// from AWS Bedrock that can be used with this plugin.
func (c *BedrockClient) ListModels() ([]string, error) {
	models := []string{}
	ctx := context.Background()

	foundationModels, err := c.controlPlaneClient.ListFoundationModels(ctx, &bedrock.ListFoundationModelsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list foundation models: %w", err)
	}

	for _, model := range foundationModels.ModelSummaries {
		models = append(models, *model.ModelId)
	}

	inferenceProfilesPaginator := bedrock.NewListInferenceProfilesPaginator(c.controlPlaneClient, &bedrock.ListInferenceProfilesInput{})

	for inferenceProfilesPaginator.HasMorePages() {
		inferenceProfiles, err := inferenceProfilesPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list inference profiles: %w", err)
		}

		for _, profile := range inferenceProfiles.InferenceProfileSummaries {
			models = append(models, *profile.InferenceProfileId)
		}
	}

	return models, nil
}

// SendStream sends the messages to the Bedrock ConverseStream API
func (c *BedrockClient) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan string) (err error) {
	// Ensure channel is closed on all exit paths to prevent goroutine leaks
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in SendStream: %v", r)
		}
		close(channel)
	}()

	messages := c.toMessages(msgs)

	var converseInput = bedrockruntime.ConverseStreamInput{
		ModelId:  aws.String(opts.Model),
		Messages: messages,
		InferenceConfig: &types.InferenceConfiguration{
			Temperature: aws.Float32(float32(opts.Temperature)),
			TopP:        aws.Float32(float32(opts.TopP))},
	}

	response, err := c.runtimeClient.ConverseStream(context.Background(), &converseInput)
	if err != nil {
		return fmt.Errorf("bedrock conversestream failed for model %s: %w", opts.Model, err)
	}

	for event := range response.GetStream().Events() {
		// Possible ConverseStream event types
		// https://docs.aws.amazon.com/bedrock/latest/userguide/conversation-inference-call.html#conversation-inference-call-response-converse-stream
		switch v := event.(type) {

		case *types.ConverseStreamOutputMemberContentBlockDelta:
			text, ok := v.Value.Delta.(*types.ContentBlockDeltaMemberText)
			if ok {
				channel <- text.Value
			}

		case *types.ConverseStreamOutputMemberMessageStop:
			channel <- "\n"
			return nil // Let defer handle the close

		// Unused Events
		case *types.ConverseStreamOutputMemberMessageStart,
			*types.ConverseStreamOutputMemberContentBlockStart,
			*types.ConverseStreamOutputMemberContentBlockStop,
			*types.ConverseStreamOutputMemberMetadata:

		default:
			return fmt.Errorf("unknown stream event type: %T", v)
		}
	}

	return nil
}

// Send sends the messages the Bedrock Converse API
func (c *BedrockClient) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (ret string, err error) {

	messages := c.toMessages(msgs)

	var converseInput = bedrockruntime.ConverseInput{
		ModelId:  aws.String(opts.Model),
		Messages: messages,
	}
	response, err := c.runtimeClient.Converse(ctx, &converseInput)
	if err != nil {
		return "", fmt.Errorf("bedrock converse failed for model %s: %w", opts.Model, err)
	}

	responseText, ok := response.Output.(*types.ConverseOutputMemberMessage)
	if !ok {
		return "", fmt.Errorf("unexpected response type: %T", response.Output)
	}

	if len(responseText.Value.Content) == 0 {
		return "", fmt.Errorf("empty response content")
	}

	responseContentBlock := responseText.Value.Content[0]
	text, ok := responseContentBlock.(*types.ContentBlockMemberText)
	if !ok {
		return "", fmt.Errorf("unexpected content block type: %T", responseContentBlock)
	}

	return text.Value, nil
}

// NeedsRawMode indicates whether the model requires raw mode processing.
// Bedrock models do not require raw mode.
func (c *BedrockClient) NeedsRawMode(modelName string) bool {
	return false
}

// toMessages converts the array of input messages from the ChatCompletionMessageType to the
// Bedrock Converse Message type.
// The system role messages are mapped to the user role as they contain a mix of system messages,
// pattern content and user input.
func (c *BedrockClient) toMessages(inputMessages []*chat.ChatCompletionMessage) (messages []types.Message) {
	for _, msg := range inputMessages {
		roles := map[string]types.ConversationRole{
			chat.ChatMessageRoleUser:      types.ConversationRoleUser,
			chat.ChatMessageRoleAssistant: types.ConversationRoleAssistant,
			chat.ChatMessageRoleSystem:    types.ConversationRoleUser,
		}

		role, ok := roles[msg.Role]
		if !ok {
			continue
		}

		message := types.Message{
			Role:    role,
			Content: []types.ContentBlock{&types.ContentBlockMemberText{Value: msg.Content}},
		}
		messages = append(messages, message)

	}

	return
}
