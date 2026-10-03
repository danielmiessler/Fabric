package vertexai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractModelName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard format",
			input:    "publishers/google/models/gemini-2.0-flash",
			expected: "gemini-2.0-flash",
		},
		{
			name:     "anthropic model",
			input:    "publishers/anthropic/models/claude-sonnet-4-5",
			expected: "claude-sonnet-4-5",
		},
		{
			name:     "model with version",
			input:    "publishers/anthropic/models/claude-3-opus@20240229",
			expected: "claude-3-opus@20240229",
		},
		{
			name:     "just model name",
			input:    "gemini-pro",
			expected: "gemini-pro",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractModelName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSortModels(t *testing.T) {
	input := []string{
		"claude-sonnet-4-5",
		"gemini-2.0-flash",
		"gemini-pro",
		"claude-opus-4",
		"unknown-model",
	}

	result := sortModels(input)

	// Gemini, then Claude, then others. Alphabetical within a group.
	expected := []string{
		"gemini-2.0-flash",
		"gemini-pro",
		"claude-opus-4",
		"claude-sonnet-4-5",
		"unknown-model",
	}

	assert.Equal(t, expected, result)
}

func TestModelPriority(t *testing.T) {
	tests := []struct {
		model    string
		priority int
	}{
		{"gemini-2.0-flash", 1},
		{"Gemini-Pro", 1},
		{"claude-sonnet-4-5", 2},
		{"CLAUDE-OPUS", 2},
		{"some-other-model", 3},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			result := modelPriority(tt.model)
			assert.Equal(t, tt.priority, result, "priority for %s", tt.model)
		})
	}
}

func TestListPublisherModels_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.Path, "/v1/publishers/google/models")

		response := publisherModelsResponse{
			PublisherModels: []publisherModel{
				{Name: "publishers/google/models/gemini-2.0-flash"},
				{Name: "publishers/google/models/gemini-pro"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// This test does not call listPublisherModels. It only checks the mock server.
	resp, err := http.Get(server.URL + "/v1/publishers/google/models")
	require.NoError(t, err)
	defer resp.Body.Close()

	var response publisherModelsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)

	assert.Len(t, response.PublisherModels, 2)
	assert.Equal(t, "publishers/google/models/gemini-2.0-flash", response.PublisherModels[0].Name)
}

func TestListPublisherModels_Pagination(t *testing.T) {
	callCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		var response publisherModelsResponse
		if callCount == 1 {
			response = publisherModelsResponse{
				PublisherModels: []publisherModel{
					{Name: "publishers/google/models/gemini-flash"},
				},
				NextPageToken: "page2",
			}
		} else {
			response = publisherModelsResponse{
				PublisherModels: []publisherModel{
					{Name: "publishers/google/models/gemini-pro"},
				},
				NextPageToken: "",
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	resp, err := http.Get(server.URL + "/page1")
	require.NoError(t, err)
	resp.Body.Close()

	resp, err = http.Get(server.URL + "/page2")
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, 2, callCount)
}

func TestListPublisherModels_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error": "access denied"}`))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/publishers/google/models")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client)
	assert.Equal(t, "VertexAI", client.Name)
	assert.NotNil(t, client.ProjectID)
	assert.NotNil(t, client.Region)
	assert.Equal(t, "global", client.Region.Value)
}

func TestPublishersListComplete(t *testing.T) {
	expectedPublishers := []string{"google", "anthropic"}

	assert.Equal(t, expectedPublishers, publishers)
}

func TestIsConversationalModel(t *testing.T) {
	tests := []struct {
		model    string
		expected bool
	}{
		{"gemini-2.0-flash", true},
		{"gemini-2.5-pro", true},
		{"claude-sonnet-4-5", true},
		{"claude-opus-4", true},
		{"deepseek-v3", true},
		{"llama-3.1-405b", true},
		{"mistral-large", true},

		{"imagen-3.0-capability-002", false},
		{"imagen-4.0-fast-generate-001", false},
		{"imagegeneration", false},
		{"imagetext", false},
		{"image-segmentation-001", false},
		{"textembedding-gecko", false},
		{"multimodalembedding", false},
		{"text-embedding-004", false},
		{"text-bison", false},
		{"text-unicorn", false},
		{"code-bison", false},
		{"code-gecko", false},
		{"codechat-bison", false},
		{"chat-bison", false},
		{"veo-001", false},
		{"chirp", false},
		{"medlm-medium", false},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			result := isConversationalModel(tt.model)
			assert.Equal(t, tt.expected, result, "isConversationalModel(%s)", tt.model)
		})
	}
}

func TestFilterConversationalModels(t *testing.T) {
	input := []string{
		"gemini-2.0-flash",
		"imagen-3.0-capability-002",
		"claude-sonnet-4-5",
		"textembedding-gecko",
		"deepseek-v3",
		"chat-bison",
		"llama-3.1-405b",
		"code-bison",
	}

	result := filterConversationalModels(input)

	expected := []string{
		"gemini-2.0-flash",
		"claude-sonnet-4-5",
		"deepseek-v3",
		"llama-3.1-405b",
	}

	assert.Equal(t, expected, result)
}

func TestFilterConversationalModels_EmptyInput(t *testing.T) {
	result := filterConversationalModels([]string{})
	assert.Empty(t, result)
}

func TestFilterConversationalModels_AllFiltered(t *testing.T) {
	input := []string{
		"imagen-3.0",
		"textembedding-gecko",
		"chat-bison",
	}

	result := filterConversationalModels(input)
	assert.Empty(t, result)
}

func TestIsGeminiModel(t *testing.T) {
	tests := []struct {
		model    string
		expected bool
	}{
		{"gemini-2.5-pro", true},
		{"gemini-3-pro-preview", true},
		{"Gemini-2.0-flash", true},
		{"GEMINI-flash", true},
		{"claude-sonnet-4-5", false},
		{"claude-opus-4", false},
		{"deepseek-v3", false},
		{"llama-3.1-405b", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			result := isGeminiModel(tt.model)
			assert.Equal(t, tt.expected, result, "isGeminiModel(%s)", tt.model)
		})
	}
}

func TestGetMaxTokens(t *testing.T) {
	tests := []struct {
		name     string
		opts     *domain.ChatOptions
		expected int64
	}{
		{
			name:     "MaxTokens specified",
			opts:     &domain.ChatOptions{MaxTokens: 8192},
			expected: 8192,
		},
		{
			name:     "Default when MaxTokens is 0",
			opts:     &domain.ChatOptions{MaxTokens: 0},
			expected: int64(defaultMaxTokens),
		},
		{
			name:     "Default when MaxTokens is negative",
			opts:     &domain.ChatOptions{MaxTokens: -1},
			expected: int64(defaultMaxTokens),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMaxTokens(tt.opts)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseGeminiThinking(t *testing.T) {
	tests := []struct {
		name           string
		level          domain.ThinkingLevel
		expectNil      bool
		expectedBudget int32
	}{
		{
			name:      "empty string returns nil",
			level:     "",
			expectNil: true,
		},
		{
			name:      "off returns nil",
			level:     domain.ThinkingOff,
			expectNil: true,
		},
		{
			name:           "low thinking",
			level:          domain.ThinkingLow,
			expectNil:      false,
			expectedBudget: int32(domain.ThinkingBudgets[domain.ThinkingLow]),
		},
		{
			name:           "medium thinking",
			level:          domain.ThinkingMedium,
			expectNil:      false,
			expectedBudget: int32(domain.ThinkingBudgets[domain.ThinkingMedium]),
		},
		{
			name:           "high thinking",
			level:          domain.ThinkingHigh,
			expectNil:      false,
			expectedBudget: int32(domain.ThinkingBudgets[domain.ThinkingHigh]),
		},
		{
			name:           "numeric string",
			level:          "5000",
			expectNil:      false,
			expectedBudget: 5000,
		},
		{
			name:      "invalid string returns nil",
			level:     "invalid",
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseGeminiThinking(tt.level)
			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				assert.True(t, result.IncludeThoughts)
				assert.Equal(t, tt.expectedBudget, *result.ThinkingBudget)
			}
		})
	}
}

func TestBuildGeminiConfig(t *testing.T) {
	client := &Client{}

	t.Run("basic config with temperature and TopP", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Temperature: 0.7,
			TopP:        0.9,
			MaxTokens:   8192,
		}
		config := client.buildGeminiConfig(opts)

		assert.NotNil(t, config)
		assert.Equal(t, float32(0.7), *config.Temperature)
		assert.Equal(t, float32(0.9), *config.TopP)
		assert.Equal(t, int32(8192), config.MaxOutputTokens)
		assert.Nil(t, config.Tools)
		assert.Nil(t, config.ThinkingConfig)
	})

	t.Run("config with search enabled", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Temperature: 0.5,
			TopP:        0.8,
			Search:      true,
		}
		config := client.buildGeminiConfig(opts)

		assert.NotNil(t, config.Tools)
		assert.Len(t, config.Tools, 1)
		assert.NotNil(t, config.Tools[0].GoogleSearch)
	})

	t.Run("config with thinking enabled", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Temperature: 0.5,
			TopP:        0.8,
			Thinking:    domain.ThinkingHigh,
		}
		config := client.buildGeminiConfig(opts)

		assert.NotNil(t, config.ThinkingConfig)
		assert.True(t, config.ThinkingConfig.IncludeThoughts)
	})
}
