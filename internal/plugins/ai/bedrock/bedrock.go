// Package bedrock provides a plugin to use Amazon Bedrock models.
// Supported models are defined in the MODELS variable.
// To add additional models, append them to the MODELS array. Models must support the Converse and ConverseStream operations
// Authentication uses the  AWS credential provider chain, similar.to the AWS CLI and SDKs
// https://docs.aws.amazon.com/sdkref/latest/guide/standardized-credentials.html
package bedrock

import (
	"context"
	"errors"
	"fmt"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
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
		ret.PluginBase = plugins.NewVendorPluginBase(vendorName, func() error {
			return fmt.Errorf(i18n.T("bedrock_unable_load_aws_config"), err)
		})
		ret.bedrockRegion = ret.PluginBase.AddSetupQuestionWithEnvName("AWS Region", true, i18n.T("bedrock_aws_region_label"))
		return
	}

	cfg.APIOptions = append(cfg.APIOptions, middleware.AddUserAgentKeyValue(userAgentKey, userAgentValue))

	runtimeClient := bedrockruntime.NewFromConfig(cfg)
	controlPlaneClient := bedrock.NewFromConfig(cfg)

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, ret.configure)

	ret.runtimeClient = runtimeClient
	ret.controlPlaneClient = controlPlaneClient

	ret.bedrockRegion = ret.PluginBase.AddSetupQuestionWithEnvName("AWS Region", true, i18n.T("bedrock_aws_region_label"))

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
		return fmt.Errorf(i18n.T("bedrock_invalid_aws_region"), c.bedrockRegion.Value)
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(c.bedrockRegion.Value))
	if err != nil {
		return fmt.Errorf(i18n.T("bedrock_unable_load_aws_config_with_region"), c.bedrockRegion.Value, err)
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
		return nil, fmt.Errorf(i18n.T("bedrock_failed_list_foundation_models"), err)
	}

	for _, model := range foundationModels.ModelSummaries {
		models = append(models, *model.ModelId)
	}

	inferenceProfilesPaginator := bedrock.NewListInferenceProfilesPaginator(c.controlPlaneClient, &bedrock.ListInferenceProfilesInput{})

	for inferenceProfilesPaginator.HasMorePages() {
		inferenceProfiles, err := inferenceProfilesPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("bedrock_failed_list_inference_profiles"), err)
		}

		for _, profile := range inferenceProfiles.InferenceProfileSummaries {
			models = append(models, *profile.InferenceProfileId)
		}
	}

	return models, nil
}

// SendStream sends the messages to the Bedrock ConverseStream API
func (c *BedrockClient) SendStream(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) (err error) {
	// Ensure channel is closed on all exit paths to prevent goroutine leaks
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(i18n.T("bedrock_panic_sendstream"), r)
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
		return fmt.Errorf(i18n.T("bedrock_conversestream_failed"), opts.Model, err)
	}

	for event := range response.GetStream().Events() {
		// Possible ConverseStream event types
		// https://docs.aws.amazon.com/bedrock/latest/userguide/conversation-inference-call.html#conversation-inference-call-response-converse-stream
		switch v := event.(type) {

		case *types.ConverseStreamOutputMemberContentBlockDelta:
			text, ok := v.Value.Delta.(*types.ContentBlockDeltaMemberText)
			if ok {
				channel <- domain.StreamUpdate{
					Type:    domain.StreamTypeContent,
					Content: text.Value,
				}
			}

		case *types.ConverseStreamOutputMemberMessageStop:
			channel <- domain.StreamUpdate{
				Type:    domain.StreamTypeContent,
				Content: "\n",
			}
			return nil // Let defer handle the close

		case *types.ConverseStreamOutputMemberMetadata:
			if v.Value.Usage != nil {
				channel <- domain.StreamUpdate{
					Type: domain.StreamTypeUsage,
					Usage: &domain.UsageMetadata{
						InputTokens:  int(*v.Value.Usage.InputTokens),
						OutputTokens: int(*v.Value.Usage.OutputTokens),
						TotalTokens:  int(*v.Value.Usage.TotalTokens),
					},
				}
			}

		// Unused Events
		case *types.ConverseStreamOutputMemberMessageStart,
			*types.ConverseStreamOutputMemberContentBlockStart,
			*types.ConverseStreamOutputMemberContentBlockStop:

		default:
			return fmt.Errorf(i18n.T("bedrock_unknown_stream_event_type"), v)
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
		return "", fmt.Errorf(i18n.T("bedrock_converse_failed"), opts.Model, err)
	}

	responseText, ok := response.Output.(*types.ConverseOutputMemberMessage)
	if !ok {
		return "", fmt.Errorf(i18n.T("bedrock_unexpected_response_type"), response.Output)
	}

	if len(responseText.Value.Content) == 0 {
		return "", errors.New(i18n.T("bedrock_empty_response_content"))
	}

	responseContentBlock := responseText.Value.Content[0]
	text, ok := responseContentBlock.(*types.ContentBlockMemberText)
	if !ok {
		return "", fmt.Errorf(i18n.T("bedrock_unexpected_content_block_type"), responseContentBlock)
	}

	return text.Value, nil
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
