// Package bedrock provides a plugin to use Amazon Bedrock models.
//
// ListModels reads the model list from the Bedrock control plane. Models must
// support the Converse and ConverseStream operations.
//
// Authentication has three modes, in priority order:
//  1. Bearer token: a Bedrock API key (ABSK token).
//  2. Explicit credentials: an AWS access key ID and secret access key from
//     fabric --setup or the .env file.
//  3. The AWS credential provider chain, the same chain that the AWS CLI and SDKs use:
//     https://docs.aws.amazon.com/sdkref/latest/guide/standardized-credentials.html
package bedrock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

// bearerTokenTransport sets the Authorization header to the ABSK bearer token
// on each request.
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

// defaultBedrockModels is the fallback that ListModels returns when
// listModelsFromAPI fails and an API key is set. Bearer token (ABSK) auth can
// lack permission for the ListFoundationModels API.
var defaultBedrockModels = []string{
	"us.anthropic.claude-sonnet-4-6",
	"us.anthropic.claude-opus-4-6-v1",
	"us.anthropic.claude-haiku-4-5-20251001-v1:0",
	"us.amazon.nova-pro-v1:0",
	"us.meta.llama3-3-70b-instruct-v1:0",
}

// setupModelChoices is the model list that Setup shows. It has unprefixed
// model IDs and us, eu, and ap cross-region inference profiles.
var setupModelChoices = []string{
	"anthropic.claude-sonnet-4-6",
	"anthropic.claude-opus-4-6-v1",
	"anthropic.claude-haiku-4-5-20251001-v1:0",
	"amazon.nova-pro-v1:0",
	"us.anthropic.claude-sonnet-4-6",
	"us.anthropic.claude-opus-4-6-v1",
	"eu.anthropic.claude-sonnet-4-6",
	"eu.anthropic.claude-opus-4-6-v1",
	"ap.anthropic.claude-sonnet-4-6",
	"ap.anthropic.claude-opus-4-6-v1",
}

// fallbackRegions is the list that fetchBedrockRegions returns when the botocore fetch fails.
var fallbackRegions = []string{
	"us-east-1",
	"us-west-2",
	"eu-west-1",
	"eu-west-3",
	"ap-southeast-1",
	"ap-northeast-1",
}

// botocoreEndpointsURL is the public endpoints.json file that lists the AWS
// regions with a Bedrock endpoint. It is a var so that tests can point it at a
// mock server.
var botocoreEndpointsURL = "https://raw.githubusercontent.com/boto/botocore/develop/botocore/data/endpoints.json"

// fetchBedrockRegions reads the Bedrock regions from the botocore endpoints.json
// file, which needs no authentication. It returns fallbackRegions on any error
// or when the file lists no Bedrock region.
func fetchBedrockRegions() []string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(botocoreEndpointsURL)
	if err != nil {
		debuglog.Log(i18n.T("bedrock_fetch_regions_failed")+": %v\n", err)
		return fallbackRegions
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		debuglog.Log(i18n.T("bedrock_fetch_regions_bad_status")+": %d\n", resp.StatusCode)
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
		debuglog.Log(i18n.T("bedrock_fetch_regions_parse_failed")+": %v\n", err)
		return fallbackRegions
	}

	var regions []string
	for _, partition := range data.Partitions {
		if svc, ok := partition.Services["bedrock"]; ok {
			for region := range svc.Endpoints {
				// Skip FIPS endpoints and "bedrock-" entries, which are not region names.
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

// maskSecret returns the first 4 and last 4 characters of a secret, or "****"
// when the secret has 12 characters or fewer.
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

	// Enter with no input skips the Bedrock setup.
	if authChoice.Value == "" {
		return nil
	}

	switch authChoice.Value {
	case "1":
		// Show a saved key masked.
		savedKey := c.bedrockAPIKey.Value
		if savedKey != "" {
			c.bedrockAPIKey.Value = maskSecret(savedKey)
		}
		if err = c.bedrockAPIKey.Ask("Bedrock"); err != nil {
			return
		}
		// Enter keeps the masked value, so restore the saved key.
		if c.bedrockAPIKey.Value == maskSecret(savedKey) {
			c.bedrockAPIKey.Value = savedKey
		}
	case "2":
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
		// A number outside the list. Keep the input as typed.
		c.bedrockRegion.Value = regionChoice.Value
	}

	// OnAnswer sets the env var, which a later Settings.Configure reads back.
	if c.bedrockRegion.Value != "" {
		_ = c.bedrockRegion.OnAnswer(c.bedrockRegion.Value)
	}

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

	// ConfigureCustom is configure. It validates the region and builds the clients.
	if c.ConfigureCustom != nil {
		err = c.ConfigureCustom()
	}
	return
}

// isValidAWSRegion checks only the length of the region name, 5 to 30 characters.
func isValidAWSRegion(region string) bool {
	if len(region) < 5 || len(region) > 30 {
		return false
	}
	return region != ""
}

// configure validates the region and builds the runtime and control plane clients.
//
// Credential priority:
//  1. API key (ABSK): dummy static credentials plus bearerTokenTransport, which
//     replaces the SigV4 Authorization header with the bearer token.
//  2. Access key ID and secret access key as static credentials.
//  3. The default AWS credential provider chain.
func (c *BedrockClient) configure() error {
	if c.bedrockRegion.Value == "" {
		return fmt.Errorf(i18n.T("bedrock_invalid_aws_region"), "(empty)")
	}

	if !isValidAWSRegion(c.bedrockRegion.Value) {
		return fmt.Errorf(i18n.T("bedrock_invalid_aws_region"), c.bedrockRegion.Value)
	}

	ctx := context.Background()

	configOpts := []func(*config.LoadOptions) error{
		config.WithRegion(c.bedrockRegion.Value),
	}

	// An API key needs dummy static credentials, not AnonymousCredentials. With
	// AnonymousCredentials the SDK selects its bearer token auth scheme, which
	// panics without a token provider. The empty shared config and credentials
	// file lists stop AWS_PROFILE from causing "failed to get shared config
	// profile" errors.
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
			config.WithSharedConfigFiles([]string{}),
			config.WithSharedCredentialsFiles([]string{}),
		)
	} else if c.bedrockAccessKey.Value != "" && c.bedrockSecretKey.Value != "" {
		configOpts = append(configOpts,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					c.bedrockAccessKey.Value,
					c.bedrockSecretKey.Value,
					"", // session token (empty for long-term credentials)
				),
			),
			config.WithSharedConfigFiles([]string{}),
			config.WithSharedCredentialsFiles([]string{}),
		)
	}

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
func (c *BedrockClient) ListModels(_ context.Context) ([]string, error) {
	models, err := c.listModelsFromAPI()
	if err != nil && c.bedrockAPIKey.Value != "" {
		debuglog.Log(i18n.T("bedrock_listmodels_fallback")+": %v\n", err)
		return defaultBedrockModels, nil
	}
	return models, err
}

// listModelsFromAPI returns the foundation model IDs and inference profile IDs
// from the Bedrock control plane.
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
func (c *BedrockClient) SendStream(_ context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) (err error) {
	// Close the channel on every exit path, including a panic, so the reader does not block.
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

	// Some models, such as Claude, reject a request that sets both temperature
	// and top_p. Send only temperature.
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
		// ConverseStream event types:
		// https://docs.aws.amazon.com/bedrock/latest/userguide/conversation-inference.html#conversation-inference-call-response-converse-stream
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
			return nil // The deferred func closes the channel.

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

		// Ignored events
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

// toMessages converts chat messages to Bedrock Converse messages. System
// messages get the user role, because they hold the pattern text and the user
// input together. Messages with any other role are dropped.
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
