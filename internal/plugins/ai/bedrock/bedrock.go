// Package bedrock provides a plugin to use Amazon Bedrock models.
// Supported models are defined in the MODELS variable.
// To add additional models, append them to the MODELS array. Models must support the Converse and ConverseStream operations
// Authentication supports three modes:
//  1. Bearer token: Provide a Bedrock API Key (ABSK token) for simple authentication
//  2. Explicit credentials: Provide AWS Access Key ID and Secret Access Key directly via fabric --setup
//  3. AWS credential provider chain (default fallback): Uses the standard chain similar to the AWS CLI and SDKs
//     https://docs.aws.amazon.com/sdkref/latest/guide/standardized-credentials.html
package bedrock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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
//
// Authentication modes (in priority order):
//  1. Bearer token: BEDROCK_API_KEY (ABSK token) — simplest, like Claude Code
//  2. Explicit credentials: BEDROCK_AWS_ACCESS_KEY_ID + BEDROCK_AWS_SECRET_ACCESS_KEY provided via setup
//  3. AWS credential chain: Standard AWS SDK credential resolution (env vars, profiles, IAM roles, etc.)
type BedrockClient struct {
	*plugins.PluginBase
	runtimeClient      *bedrockruntime.Client
	controlPlaneClient *bedrock.Client

	bedrockRegion    *plugins.SetupQuestion
	bedrockAccessKey *plugins.SetupQuestion
	bedrockSecretKey *plugins.SetupQuestion
	bedrockAPIKey    *plugins.SetupQuestion
}

// bearerTokenTransport is an http.RoundTripper that injects an Authorization
// Bearer header into every outgoing request. Used for ABSK key authentication.
type bearerTokenTransport struct {
	token   string
	wrapped http.RoundTripper
}

func (t *bearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+t.token)
	return t.wrapped.RoundTrip(clone)
}

// String implements fmt.Stringer with token redaction to prevent accidental
// exposure of the ABSK key in logs or debug output.
func (t *bearerTokenTransport) String() string {
	return "bearerTokenTransport{token:REDACTED}"
}

// defaultBedrockModels is a minimal fallback used ONLY when the ListFoundationModels
// and ListInferenceProfiles APIs are not accessible. The primary model listing is always
// fetched programmatically via listModelsFromAPI() which calls the Bedrock control plane.
//
// This fallback is needed because bearer token (ABSK) auth may not have permissions for
// the ListFoundationModels API. In practice, most users will never see this list — it's
// only used when the API call fails AND the user has an API key configured.
var defaultBedrockModels = []string{
	"us.anthropic.claude-sonnet-4-6",
	"us.anthropic.claude-opus-4-6-v1",
	"us.anthropic.claude-haiku-4-5-20251001-v1:0",
	"us.amazon.nova-pro-v1:0",
	"us.meta.llama3-3-70b-instruct-v1:0",
}

// setupModelChoices is shown during interactive setup. Includes both unprefixed
// model IDs (work in any region) and common region-prefixed inference profiles.
var setupModelChoices = []string{
	// Unprefixed (work in any region)
	"anthropic.claude-sonnet-4-6",
	"anthropic.claude-opus-4-6-v1",
	"anthropic.claude-haiku-4-5-20251001-v1:0",
	"amazon.nova-pro-v1:0",
	// US cross-region inference profiles
	"us.anthropic.claude-sonnet-4-6",
	"us.anthropic.claude-opus-4-6-v1",
	// EU cross-region inference profiles
	"eu.anthropic.claude-sonnet-4-6",
	"eu.anthropic.claude-opus-4-6-v1",
	// AP cross-region inference profiles
	"ap.anthropic.claude-sonnet-4-6",
	"ap.anthropic.claude-opus-4-6-v1",
}

// fallbackRegions is used only when the dynamic fetch from botocore fails (e.g., no network).
var fallbackRegions = []string{
	"us-east-1",
	"us-west-2",
	"eu-west-1",
	"eu-west-3",
	"ap-southeast-1",
	"ap-northeast-1",
}

// botocoreEndpointsURL is the public (no-auth) source of truth for which AWS
// regions support Bedrock, maintained by the AWS SDK team.
const botocoreEndpointsURL = "https://raw.githubusercontent.com/boto/botocore/develop/botocore/data/endpoints.json"

// fetchBedrockRegions fetches the list of AWS regions where Bedrock is available
// from the botocore endpoints.json file (public, no authentication required).
// Falls back to the static fallbackRegions list on any error.
func fetchBedrockRegions() []string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(botocoreEndpointsURL)
	if err != nil {
		debuglog.Log("Failed to fetch Bedrock regions from botocore: %v\n", err)
		return fallbackRegions
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		debuglog.Log("Botocore endpoints returned status %d\n", resp.StatusCode)
		return fallbackRegions
	}

	var data struct {
		Partitions []struct {
			Services map[string]struct {
				Endpoints map[string]any `json:"endpoints"`
			} `json:"services"`
		} `json:"partitions"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		debuglog.Log("Failed to parse botocore endpoints.json: %v\n", err)
		return fallbackRegions
	}

	var regions []string
	for _, partition := range data.Partitions {
		if svc, ok := partition.Services["bedrock"]; ok {
			for region := range svc.Endpoints {
				// Skip FIPS and special endpoints (e.g., "bedrock-us-east-1")
				if !strings.HasPrefix(region, "bedrock-") && !strings.Contains(region, "fips") {
					regions = append(regions, region)
				}
			}
		}
	}

	if len(regions) == 0 {
		return fallbackRegions
	}

	sort.Strings(regions)
	return regions
}

// maskSecret redacts a secret value for display, showing only the first 4 and last 4 chars.
func maskSecret(s string) string {
	if len(s) <= 12 {
		return "****"
	}
	return s[:4] + "..." + s[len(s)-4:]
}

// NewClient returns a new Bedrock plugin client.
// Client initialization is deferred to configure() so that explicit credentials
// from the .env file or --setup can be used when available.
func NewClient() (ret *BedrockClient) {
	vendorName := "Bedrock"
	ret = &BedrockClient{}

	ret.PluginBase = plugins.NewVendorPluginBase(vendorName, ret.configure)

	// Settings registered for .env persistence (all optional except region)
	ret.bedrockRegion = ret.PluginBase.AddSetupQuestionWithEnvName(
		"AWS Region", true, i18n.T("bedrock_aws_region_label"))
	ret.bedrockAPIKey = ret.PluginBase.AddSetupQuestionWithEnvName(
		"API Key", false, i18n.T("bedrock_api_key_label"))
	ret.bedrockAccessKey = ret.PluginBase.AddSetupQuestionWithEnvName(
		"AWS Access Key ID", false, i18n.T("bedrock_aws_access_key_label"))
	ret.bedrockSecretKey = ret.PluginBase.AddSetupQuestionWithEnvName(
		"AWS Secret Access Key", false, i18n.T("bedrock_aws_secret_key_label"))

	return
}

// Setup overrides the default plugin Setup to provide a guided auth flow:
// 1. Choose auth method: API Key (ABSK) or AWS Access Key + Secret
// 2. Based on choice, ask only the relevant questions
// 3. Choose region from list or type custom
// 4. Choose model from list or type custom
func (c *BedrockClient) Setup() (err error) {
	fmt.Println()
	fmt.Println(i18n.T("bedrock_setup_header"))
	fmt.Println()
	fmt.Println(i18n.T("bedrock_setup_choose_auth_method"))
	fmt.Println(i18n.T("bedrock_setup_auth_option_apikey"))
	fmt.Println(i18n.T("bedrock_setup_auth_option_accesskey"))
	fmt.Println()

	authChoice := plugins.NewSetupQuestion(i18n.T("bedrock_setup_auth_prompt"))
	if err = authChoice.Ask("Bedrock"); err != nil {
		return
	}

	// Empty input means skip (user pressed enter without typing)
	if authChoice.Value == "" {
		return nil
	}

	switch authChoice.Value {
	case "1":
		// Mask existing API key value before displaying the prompt
		savedKey := c.bedrockAPIKey.Value
		if savedKey != "" {
			c.bedrockAPIKey.Value = maskSecret(savedKey)
		}
		if err = c.bedrockAPIKey.Ask("Bedrock"); err != nil {
			return
		}
		// If user kept the masked value (pressed enter), restore the real key
		if c.bedrockAPIKey.Value == maskSecret(savedKey) {
			c.bedrockAPIKey.Value = savedKey
		}
	case "2":
		// Mask existing credentials before displaying
		savedAccess := c.bedrockAccessKey.Value
		if savedAccess != "" {
			c.bedrockAccessKey.Value = maskSecret(savedAccess)
		}
		if err = c.bedrockAccessKey.Ask("Bedrock"); err != nil {
			return
		}
		if c.bedrockAccessKey.Value == maskSecret(savedAccess) {
			c.bedrockAccessKey.Value = savedAccess
		}

		savedSecret := c.bedrockSecretKey.Value
		if savedSecret != "" {
			c.bedrockSecretKey.Value = maskSecret(savedSecret)
		}
		if err = c.bedrockSecretKey.Ask("Bedrock"); err != nil {
			return
		}
		if c.bedrockSecretKey.Value == maskSecret(savedSecret) {
			c.bedrockSecretKey.Value = savedSecret
		}
	default:
		return fmt.Errorf(i18n.T("bedrock_setup_invalid_auth_selection"), authChoice.Value)
	}

	// Region selection — fetched dynamically from botocore (public, no auth required)
	regions := fetchBedrockRegions()
	fmt.Println()
	fmt.Println(i18n.T("bedrock_setup_choose_region"))
	for i, r := range regions {
		fmt.Printf("    [%d] %s\n", i+1, r)
	}
	fmt.Println(i18n.T("bedrock_setup_region_option_custom"))
	fmt.Println()

	regionChoice := plugins.NewSetupQuestion(i18n.T("bedrock_setup_region_prompt"))
	if err = regionChoice.Ask("Bedrock"); err != nil {
		return
	}

	regionNum := 0
	if _, scanErr := fmt.Sscanf(regionChoice.Value, "%d", &regionNum); scanErr == nil && regionNum >= 1 && regionNum <= len(regions) {
		c.bedrockRegion.Value = regions[regionNum-1]
	} else if regionNum == 0 || regionChoice.Value == "0" {
		customRegion := plugins.NewSetupQuestion(i18n.T("bedrock_setup_region_custom_prompt"))
		if err = customRegion.Ask("Bedrock"); err != nil {
			return
		}
		c.bedrockRegion.Value = customRegion.Value
	} else {
		// They typed a region name directly
		c.bedrockRegion.Value = regionChoice.Value
	}

	// Set the env var so it persists
	if c.bedrockRegion.Value != "" {
		_ = c.bedrockRegion.OnAnswer(c.bedrockRegion.Value)
	}

	// Model selection (shown after auth + region)
	fmt.Println()
	fmt.Println(i18n.T("bedrock_setup_choose_model"))
	for i, m := range setupModelChoices {
		fmt.Printf("    [%d] %s\n", i+1, m)
	}
	fmt.Println(i18n.T("bedrock_setup_model_option_custom"))
	fmt.Println()

	modelChoice := plugins.NewSetupQuestion(i18n.T("bedrock_setup_model_prompt"))
	if err = modelChoice.Ask("Bedrock"); err != nil {
		return
	}

	modelNum := 0
	selectedModel := ""
	if _, scanErr := fmt.Sscanf(modelChoice.Value, "%d", &modelNum); scanErr == nil && modelNum >= 1 && modelNum <= len(setupModelChoices) {
		selectedModel = setupModelChoices[modelNum-1]
	} else if modelNum == 0 || modelChoice.Value == "0" {
		customModel := plugins.NewSetupQuestion(i18n.T("bedrock_setup_model_custom_prompt"))
		if err = customModel.Ask("Bedrock"); err != nil {
			return
		}
		selectedModel = customModel.Value
	} else {
		selectedModel = modelChoice.Value
	}

	if selectedModel != "" {
		fmt.Printf("\n"+i18n.T("bedrock_setup_selected_model")+"\n", selectedModel)
		fmt.Printf(i18n.T("bedrock_setup_use_with")+"\n", selectedModel)
	}

	// Run configure to validate and initialize clients
	if c.ConfigureCustom != nil {
		err = c.ConfigureCustom()
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

// configure initializes the Bedrock clients with the appropriate credentials and region.
//
// Authentication priority:
//  1. If a Bearer token / API key (ABSK) is provided, use it directly via Authorization header.
//     This skips SigV4 signing and is the simplest setup (like Claude Code's BEDROCK_API_KEY).
//  2. If explicit Access Key ID + Secret Access Key are provided (via setup or env vars),
//     use them as static credentials.
//  3. Otherwise, fall back to the standard AWS credential provider chain
//     (env vars like AWS_ACCESS_KEY_ID, AWS profiles, IAM roles, etc.)
func (c *BedrockClient) configure() error {
	if c.bedrockRegion.Value == "" {
		return fmt.Errorf(i18n.T("bedrock_invalid_aws_region"), "(empty)")
	}

	// Validate region format
	if !isValidAWSRegion(c.bedrockRegion.Value) {
		return fmt.Errorf(i18n.T("bedrock_invalid_aws_region"), c.bedrockRegion.Value)
	}

	ctx := context.Background()

	// Build config options
	configOpts := []func(*config.LoadOptions) error{
		config.WithRegion(c.bedrockRegion.Value),
	}

	// Priority 1: Bearer token / API key (ABSK key)
	// We use dummy static credentials (not AnonymousCredentials) to satisfy the
	// AWS SDK's SigV4 auth middleware. AnonymousCredentials causes the SDK to fall
	// through to its bearer token auth path, which panics without a token provider.
	// Our bearerTokenTransport overrides the Authorization header with the real token.
	// When using explicit credentials (bearer token or static keys), temporarily
	// clear AWS_PROFILE to prevent the SDK from trying to load a shared config
	// profile that may not exist. This is a real-world issue: users with
	// AWS_PROFILE set for other tools (terraform, aws-cli) would get
	// "failed to get shared config profile" errors even though they provided
	// credentials directly.
	explicitCreds := c.bedrockAPIKey.Value != "" || (c.bedrockAccessKey.Value != "" && c.bedrockSecretKey.Value != "")
	if explicitCreds {
		if savedProfile, hasProfile := os.LookupEnv("AWS_PROFILE"); hasProfile {
			os.Unsetenv("AWS_PROFILE")
			defer os.Setenv("AWS_PROFILE", savedProfile)
		}
	}

	if c.bedrockAPIKey.Value != "" {
		configOpts = append(configOpts,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider("BEDROCK_BEARER", "BEDROCK_BEARER", ""),
			),
			config.WithHTTPClient(&http.Client{
				Transport: &bearerTokenTransport{
					token:   c.bedrockAPIKey.Value,
					wrapped: http.DefaultTransport,
				},
			}),
		)
	} else if c.bedrockAccessKey.Value != "" && c.bedrockSecretKey.Value != "" {
		// Priority 2: Explicit access key + secret key (static credentials)
		configOpts = append(configOpts, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				c.bedrockAccessKey.Value,
				c.bedrockSecretKey.Value,
				"", // session token (empty for long-term credentials)
			),
		))
	}
	// Priority 3: No explicit credentials → AWS SDK uses the default credential chain
	// (AWS_ACCESS_KEY_ID/AWS_SECRET_ACCESS_KEY env vars, ~/.aws/credentials, IAM roles, etc.)

	cfg, err := config.LoadDefaultConfig(ctx, configOpts...)
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
// When using bearer token auth, the API may not be accessible, so a static
// fallback list of common models is returned instead.
func (c *BedrockClient) ListModels() ([]string, error) {
	models, err := c.listModelsFromAPI()
	if err != nil && c.bedrockAPIKey.Value != "" {
		// Bearer token auth may lack ListFoundationModels permissions;
		// return common models as fallback
		debuglog.Log("Bedrock ListModels API failed (using static fallback): %v\n", err)
		return defaultBedrockModels, nil
	}
	return models, err
}

// listModelsFromAPI queries the Bedrock control plane for available models.
func (c *BedrockClient) listModelsFromAPI() ([]string, error) {
	if c.controlPlaneClient == nil {
		return nil, errors.New(i18n.T("bedrock_client_not_initialized"))
	}
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

	if c.runtimeClient == nil {
		return errors.New(i18n.T("bedrock_client_not_initialized"))
	}

	messages := c.toMessages(msgs)

	// Some models (e.g., Claude on Bedrock) reject requests with both temperature
	// and top_p set simultaneously. Only send temperature as it's the more common parameter.
	var converseInput = bedrockruntime.ConverseStreamInput{
		ModelId:  aws.String(opts.Model),
		Messages: messages,
		InferenceConfig: &types.InferenceConfiguration{
			Temperature: aws.Float32(float32(opts.Temperature)),
		},
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
	if c.runtimeClient == nil {
		return "", errors.New(i18n.T("bedrock_client_not_initialized"))
	}

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
