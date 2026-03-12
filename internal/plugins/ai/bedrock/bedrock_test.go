package bedrock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient_CreatesAllSetupQuestions(t *testing.T) {
	client := NewClient()

	require.NotNil(t, client)
	require.NotNil(t, client.PluginBase)
	assert.Equal(t, "Bedrock", client.GetName())

	// Verify all 4 setup questions exist
	require.NotNil(t, client.bedrockRegion, "bedrockRegion setup question should exist")
	require.NotNil(t, client.bedrockAPIKey, "bedrockAPIKey setup question should exist")
	require.NotNil(t, client.bedrockAccessKey, "bedrockAccessKey setup question should exist")
	require.NotNil(t, client.bedrockSecretKey, "bedrockSecretKey setup question should exist")

	// Region is required, others are optional
	assert.True(t, client.bedrockRegion.Required, "region should be required")
	assert.False(t, client.bedrockAPIKey.Required, "API key should be optional")
	assert.False(t, client.bedrockAccessKey.Required, "access key should be optional")
	assert.False(t, client.bedrockSecretKey.Required, "secret key should be optional")
}

func TestNewClient_SetupQuestionOrder(t *testing.T) {
	client := NewClient()

	// Verify question order: Region → API Key → Access Key → Secret Key
	// (API Key should come before Access/Secret for best UX since it's simplest)
	require.Len(t, client.SetupQuestions, 4)
	assert.Contains(t, client.SetupQuestions[0].EnvVariable, "AWS_REGION")
	assert.Contains(t, client.SetupQuestions[1].EnvVariable, "API_KEY")
	assert.Contains(t, client.SetupQuestions[2].EnvVariable, "AWS_ACCESS_KEY_ID")
	assert.Contains(t, client.SetupQuestions[3].EnvVariable, "AWS_SECRET_ACCESS_KEY")
}

func TestNewClient_DeferredInit(t *testing.T) {
	client := NewClient()

	// Clients should be nil before configure() — deferred initialization
	assert.Nil(t, client.runtimeClient, "runtimeClient should be nil before configure()")
	assert.Nil(t, client.controlPlaneClient, "controlPlaneClient should be nil before configure()")
}

func TestConfigure_EmptyRegion_ReturnsError(t *testing.T) {
	client := NewClient()
	client.bedrockRegion.Value = ""

	err := client.configure()
	assert.Error(t, err, "configure() should return error when region is empty")
}

func TestConfigure_InvalidRegion_ReturnsError(t *testing.T) {
	client := NewClient()
	client.bedrockRegion.Value = "bad"

	err := client.configure()
	assert.Error(t, err, "configure() should return error for invalid region")
}

func TestConfigure_ValidRegion_BearerToken(t *testing.T) {
	t.Setenv("AWS_PROFILE", "")
	client := NewClient()
	client.bedrockRegion.Value = "us-east-1"
	client.bedrockAPIKey.Value = "test-absk-token"

	err := client.configure()
	assert.NoError(t, err, "configure() should succeed with valid region + API key")

	// Clients should be initialized after configure()
	assert.NotNil(t, client.runtimeClient, "runtimeClient should be initialized")
	assert.NotNil(t, client.controlPlaneClient, "controlPlaneClient should be initialized")
}

func TestConfigure_ValidRegion_StaticCredentials(t *testing.T) {
	t.Setenv("AWS_PROFILE", "")
	client := NewClient()
	client.bedrockRegion.Value = "us-west-2"
	client.bedrockAccessKey.Value = "AKIAIOSFODNN7EXAMPLE"
	client.bedrockSecretKey.Value = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

	err := client.configure()
	assert.NoError(t, err, "configure() should succeed with valid region + access key + secret key")

	assert.NotNil(t, client.runtimeClient, "runtimeClient should be initialized")
	assert.NotNil(t, client.controlPlaneClient, "controlPlaneClient should be initialized")
}

func TestConfigure_ValidRegion_DefaultChain(t *testing.T) {
	t.Setenv("AWS_PROFILE", "")
	client := NewClient()
	client.bedrockRegion.Value = "eu-west-1"
	// No API key, no access key — should fall back to default credential chain

	err := client.configure()
	assert.NoError(t, err, "configure() should succeed with valid region and default credential chain")

	assert.NotNil(t, client.runtimeClient, "runtimeClient should be initialized")
	assert.NotNil(t, client.controlPlaneClient, "controlPlaneClient should be initialized")
}

func TestConfigure_BearerTokenPriority(t *testing.T) {
	t.Setenv("AWS_PROFILE", "")
	// If both API key and access key are provided, API key (bearer) should win
	client := NewClient()
	client.bedrockRegion.Value = "us-east-1"
	client.bedrockAPIKey.Value = "test-absk-token"
	client.bedrockAccessKey.Value = "AKIAIOSFODNN7EXAMPLE"
	client.bedrockSecretKey.Value = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

	err := client.configure()
	assert.NoError(t, err, "configure() should succeed when both auth methods are provided")

	// We can't easily inspect which credential provider was used, but at least
	// verify it initialized successfully (bearer token takes priority)
	assert.NotNil(t, client.runtimeClient)
}

func TestIsValidAWSRegion(t *testing.T) {
	tests := []struct {
		name     string
		region   string
		expected bool
	}{
		{"valid us-east-1", "us-east-1", true},
		{"valid eu-west-1", "eu-west-1", true},
		{"valid ap-southeast-2", "ap-southeast-2", true},
		{"too short", "us", false},
		{"too short 2", "bad", false},
		{"empty", "", false},
		{"just long enough", "us-ea", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isValidAWSRegion(tt.region))
		})
	}
}

func TestBearerTokenTransport_InjectsHeader(t *testing.T) {
	token := "test-absk-token-12345"

	// Create a transport that records the request
	var capturedReq *http.Request
	mockTransport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		capturedReq = req
		return &http.Response{StatusCode: 200}, nil
	})

	transport := &bearerTokenTransport{
		token:   token,
		wrapped: mockTransport,
	}

	req, _ := http.NewRequest("POST", "https://bedrock.us-east-1.amazonaws.com/model/invoke", nil)
	req.Header.Set("X-Original", "preserved")

	_, err := transport.RoundTrip(req)
	require.NoError(t, err)
	require.NotNil(t, capturedReq)

	// Verify Authorization header is set
	assert.Equal(t, "Bearer "+token, capturedReq.Header.Get("Authorization"))

	// Verify original request is NOT modified (clone is used)
	assert.Empty(t, req.Header.Get("Authorization"), "original request should not be modified")

	// Verify other headers are preserved in clone
	assert.Equal(t, "preserved", capturedReq.Header.Get("X-Original"))
}

func TestBearerTokenTransport_StringRedactsToken(t *testing.T) {
	transport := &bearerTokenTransport{
		token:   "super-secret-absk-key",
		wrapped: http.DefaultTransport,
	}

	str := transport.String()
	assert.Contains(t, str, "REDACTED")
	assert.NotContains(t, str, "super-secret-absk-key", "token should not appear in String() output")
}

func TestDefaultBedrockModels_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, defaultBedrockModels, "default models list should not be empty")
	for _, model := range defaultBedrockModels {
		assert.NotEmpty(t, model, "each model ID should be non-empty")
	}
}

func TestListModels_NilClient_WithApiKey_ReturnsFallback(t *testing.T) {
	client := NewClient()
	client.bedrockAPIKey.Value = "test-absk-token"
	// Don't call configure() — clients are nil

	models, err := client.ListModels()
	assert.NoError(t, err, "ListModels should not error when falling back to static list")
	assert.Equal(t, defaultBedrockModels, models, "should return default models as fallback")
}

func TestListModels_NilClient_NoApiKey_ReturnsError(t *testing.T) {
	client := NewClient()
	// Don't call configure() and no API key — should propagate error

	_, err := client.ListModels()
	assert.Error(t, err, "ListModels should error when client is nil and no API key for fallback")
}

func TestSendStream_NilClient_ReturnsError(t *testing.T) {
	client := NewClient()
	// Don't call configure() — runtimeClient is nil

	ch := make(chan domain.StreamUpdate, 10)
	opts := &domain.ChatOptions{Model: "test-model", Temperature: 0.7, TopP: 0.9}

	err := client.SendStream(nil, opts, ch)
	assert.Error(t, err, "SendStream should return error when client is nil")
	assert.Contains(t, err.Error(), i18n.T("bedrock_client_not_initialized"))
}

func TestSend_NilClient_ReturnsError(t *testing.T) {
	client := NewClient()
	// Don't call configure() — runtimeClient is nil

	opts := &domain.ChatOptions{Model: "test-model"}
	_, err := client.Send(context.Background(), nil, opts)
	assert.Error(t, err, "Send should return error when client is nil")
	assert.Contains(t, err.Error(), i18n.T("bedrock_client_not_initialized"))
}

func TestMaskSecret(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"long token", "ABSKabcdefghijklmnopqrstuvwxyz1234", "ABSK...1234"},
		{"short string", "short", "****"},
		{"empty", "", "****"},
		{"exactly 12", "123456789012", "****"},
		{"13 chars", "1234567890123", "1234...0123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maskSecret(tt.input)
			assert.Equal(t, tt.expected, result)
			if tt.input != "" && len(tt.input) > 12 {
				assert.NotContains(t, result, tt.input, "full secret should not appear in masked output")
			}
		})
	}
}

func TestSetupModelChoices_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, setupModelChoices, "setup model choices should not be empty")

	// Should contain both unprefixed and region-prefixed models
	hasUnprefixed := false
	hasUS := false
	hasEU := false
	hasAP := false
	for _, m := range setupModelChoices {
		if m == "anthropic.claude-sonnet-4-6" {
			hasUnprefixed = true
		}
		if m == "us.anthropic.claude-sonnet-4-6" {
			hasUS = true
		}
		if m == "eu.anthropic.claude-sonnet-4-6" {
			hasEU = true
		}
		if m == "ap.anthropic.claude-sonnet-4-6" {
			hasAP = true
		}
	}
	assert.True(t, hasUnprefixed, "should have unprefixed models")
	assert.True(t, hasUS, "should have US-prefixed models")
	assert.True(t, hasEU, "should have EU-prefixed models")
	assert.True(t, hasAP, "should have AP-prefixed models")
}

func TestToMessages(t *testing.T) {
	client := NewClient()

	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "You are helpful"},
		{Role: chat.ChatMessageRoleUser, Content: "Hello"},
		{Role: chat.ChatMessageRoleAssistant, Content: "Hi there"},
	}

	result := client.toMessages(msgs)
	require.Len(t, result, 3)

	// System maps to User in Bedrock
	assert.Equal(t, types.ConversationRoleUser, result[0].Role)
	assert.Equal(t, types.ConversationRoleUser, result[1].Role)
	assert.Equal(t, types.ConversationRoleAssistant, result[2].Role)
}

func TestToMessages_SkipsUnknownRoles(t *testing.T) {
	client := NewClient()

	msgs := []*chat.ChatCompletionMessage{
		{Role: "unknown_role", Content: "skip me"},
		{Role: chat.ChatMessageRoleUser, Content: "keep me"},
	}

	result := client.toMessages(msgs)
	require.Len(t, result, 1, "should skip unknown roles")
	assert.Equal(t, types.ConversationRoleUser, result[0].Role)
}

func TestToMessages_Empty(t *testing.T) {
	client := NewClient()
	result := client.toMessages(nil)
	assert.Empty(t, result)
}

// --- fetchBedrockRegions mock HTTP tests ---

// withMockEndpointsURL temporarily overrides the botocore endpoints URL for testing.
// NOTE: Not safe with t.Parallel() — tests using this helper must run sequentially.
func withMockEndpointsURL(url string, fn func()) {
	orig := botocoreEndpointsURL
	botocoreEndpointsURL = url
	defer func() { botocoreEndpointsURL = orig }()
	fn()
}

func TestFetchBedrockRegions_ValidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"partitions":[{"services":{"bedrock":{"endpoints":{"us-east-1":{},"eu-west-1":{},"ap-southeast-1":{},"bedrock-us-east-1":{"hostname":"x"}}}}}]}`)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		assert.Contains(t, regions, "us-east-1")
		assert.Contains(t, regions, "eu-west-1")
		assert.Contains(t, regions, "ap-southeast-1")
		// bedrock- prefixed should be filtered
		for _, r := range regions {
			assert.False(t, len(r) > 8 && r[:8] == "bedrock-", "should filter bedrock- prefix: %s", r)
		}
	})
}

func TestFetchBedrockRegions_HTTPError_ReturnsFallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		assert.Equal(t, fallbackRegions, regions)
	})
}

func TestFetchBedrockRegions_InvalidJSON_ReturnsFallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{not valid json`)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		assert.Equal(t, fallbackRegions, regions)
	})
}

func TestFetchBedrockRegions_NoBedrock_ReturnsFallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"partitions":[{"services":{"s3":{"endpoints":{"us-east-1":{}}}}}]}`)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		assert.Equal(t, fallbackRegions, regions)
	})
}

func TestFetchBedrockRegions_Unreachable_ReturnsFallback(t *testing.T) {
	withMockEndpointsURL("http://127.0.0.1:1", func() {
		regions := fetchBedrockRegions()
		assert.Equal(t, fallbackRegions, regions)
	})
}

func TestFetchBedrockRegions_EmptyEndpoints_ReturnsFallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"partitions":[{"services":{"bedrock":{"endpoints":{}}}}]}`)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		assert.Equal(t, fallbackRegions, regions)
	})
}

func TestFetchBedrockRegions_ResultsSorted(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"partitions":[{"services":{"bedrock":{"endpoints":{"us-west-2":{},"ap-northeast-1":{},"eu-central-1":{}}}}}]}`)
	}))
	defer server.Close()

	withMockEndpointsURL(server.URL, func() {
		regions := fetchBedrockRegions()
		require.Len(t, regions, 3)
		assert.Equal(t, "ap-northeast-1", regions[0])
		assert.Equal(t, "eu-central-1", regions[1])
		assert.Equal(t, "us-west-2", regions[2])
	})
}

func TestFallbackRegions_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, fallbackRegions)
	for _, r := range fallbackRegions {
		assert.True(t, isValidAWSRegion(r), "fallback region %q should be valid", r)
	}
}

// roundTripFunc is a helper to create http.RoundTripper from a function
type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
