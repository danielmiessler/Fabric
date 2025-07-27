package openai

import (
	"fmt"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/openai/openai-go/responses"
	"github.com/stretchr/testify/assert"
)

func TestShouldUseImageGeneration(t *testing.T) {
	client := NewClient()

	// Test with image file specified
	opts := &domain.ChatOptions{
		ImageFile: "output.png",
	}
	assert.True(t, client.shouldUseImageGeneration(opts), "Should use image generation when image file is specified")

	// Test without image file
	opts = &domain.ChatOptions{
		ImageFile: "",
	}
	assert.False(t, client.shouldUseImageGeneration(opts), "Should not use image generation when no image file is specified")
}

func TestAddImageGenerationTool(t *testing.T) {
	client := NewClient()

	// Test with image generation enabled
	opts := &domain.ChatOptions{
		ImageFile: "output.png",
	}
	tools := []responses.ToolUnionParam{}
	result := client.addImageGenerationTool(opts, tools)

	assert.Len(t, result, 1, "Should add one image generation tool")
	assert.NotNil(t, result[0].OfImageGeneration, "Should have image generation tool")
	assert.Equal(t, "image_generation", string(result[0].OfImageGeneration.Type))
	assert.Equal(t, "gpt-image-1", result[0].OfImageGeneration.Model)
	assert.Equal(t, "png", result[0].OfImageGeneration.OutputFormat)

	// Test without image generation
	opts = &domain.ChatOptions{
		ImageFile: "",
	}
	tools = []responses.ToolUnionParam{}
	result = client.addImageGenerationTool(opts, tools)

	assert.Len(t, result, 0, "Should not add image generation tool when not needed")
}

func TestBuildResponseParams_WithImageGeneration(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:     "gpt-image-1",
		ImageFile: "output.png",
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Generate an image of a cat"},
	}

	params := client.buildResponseParams(msgs, opts)

	assert.NotNil(t, params.Tools, "Expected tools when image generation is enabled")

	// Should have image generation tool
	hasImageTool := false
	for _, tool := range params.Tools {
		if tool.OfImageGeneration != nil {
			hasImageTool = true
			assert.Equal(t, "image_generation", string(tool.OfImageGeneration.Type))
			assert.Equal(t, "gpt-image-1", tool.OfImageGeneration.Model)
			break
		}
	}
	assert.True(t, hasImageTool, "Should have image generation tool")
}

func TestBuildResponseParams_WithBothSearchAndImage(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{
		Model:          "gpt-image-1",
		Search:         true,
		SearchLocation: "America/Los_Angeles",
		ImageFile:      "output.png",
	}

	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Search for cat images and generate one"},
	}

	params := client.buildResponseParams(msgs, opts)

	assert.NotNil(t, params.Tools, "Expected tools when both search and image generation are enabled")
	assert.Len(t, params.Tools, 2, "Should have both search and image generation tools")

	hasSearchTool := false
	hasImageTool := false

	for _, tool := range params.Tools {
		if tool.OfWebSearchPreview != nil {
			hasSearchTool = true
		}
		if tool.OfImageGeneration != nil {
			hasImageTool = true
		}
	}

	assert.True(t, hasSearchTool, "Should have web search tool")
	assert.True(t, hasImageTool, "Should have image generation tool")
}

func TestGetOutputFormatFromExtension(t *testing.T) {
	tests := []struct {
		name           string
		imagePath      string
		expectedFormat string
	}{
		{
			name:           "PNG extension",
			imagePath:      "/tmp/output.png",
			expectedFormat: "png",
		},
		{
			name:           "WEBP extension",
			imagePath:      "/tmp/output.webp",
			expectedFormat: "webp",
		},
		{
			name:           "JPG extension",
			imagePath:      "/tmp/output.jpg",
			expectedFormat: "jpeg",
		},
		{
			name:           "JPEG extension",
			imagePath:      "/tmp/output.jpeg",
			expectedFormat: "jpeg",
		},
		{
			name:           "Uppercase PNG extension",
			imagePath:      "/tmp/output.PNG",
			expectedFormat: "png",
		},
		{
			name:           "Mixed case JPEG extension",
			imagePath:      "/tmp/output.JpEg",
			expectedFormat: "jpeg",
		},
		{
			name:           "Empty path",
			imagePath:      "",
			expectedFormat: "png",
		},
		{
			name:           "No extension",
			imagePath:      "/tmp/output",
			expectedFormat: "png",
		},
		{
			name:           "Unsupported extension",
			imagePath:      "/tmp/output.gif",
			expectedFormat: "png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getOutputFormatFromExtension(tt.imagePath)
			assert.Equal(t, tt.expectedFormat, result)
		})
	}
}

func TestAddImageGenerationToolWithDynamicFormat(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name           string
		imageFile      string
		expectedFormat string
	}{
		{
			name:           "PNG file",
			imageFile:      "/tmp/output.png",
			expectedFormat: "png",
		},
		{
			name:           "WEBP file",
			imageFile:      "/tmp/output.webp",
			expectedFormat: "webp",
		},
		{
			name:           "JPG file",
			imageFile:      "/tmp/output.jpg",
			expectedFormat: "jpeg",
		},
		{
			name:           "JPEG file",
			imageFile:      "/tmp/output.jpeg",
			expectedFormat: "jpeg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &domain.ChatOptions{
				ImageFile: tt.imageFile,
			}

			tools := client.addImageGenerationTool(opts, []responses.ToolUnionParam{})

			assert.Len(t, tools, 1, "Should have one tool")
			assert.NotNil(t, tools[0].OfImageGeneration, "Should be image generation tool")
			assert.Equal(t, tt.expectedFormat, tools[0].OfImageGeneration.OutputFormat, "Output format should match file extension")
		})
	}
}

func TestSupportsImageGeneration(t *testing.T) {
	tests := []struct {
		name     string
		model    string
		expected bool
	}{
		{
			name:     "gpt-4o supports image generation",
			model:    "gpt-4o",
			expected: true,
		},
		{
			name:     "gpt-4o-mini supports image generation",
			model:    "gpt-4o-mini",
			expected: true,
		},
		{
			name:     "gpt-4.1 supports image generation",
			model:    "gpt-4.1",
			expected: true,
		},
		{
			name:     "gpt-4.1-mini supports image generation",
			model:    "gpt-4.1-mini",
			expected: true,
		},
		{
			name:     "gpt-4.1-nano supports image generation",
			model:    "gpt-4.1-nano",
			expected: true,
		},
		{
			name:     "o3 supports image generation",
			model:    "o3",
			expected: true,
		},
		{
			name:     "o1 does not support image generation",
			model:    "o1",
			expected: false,
		},
		{
			name:     "o1-mini does not support image generation",
			model:    "o1-mini",
			expected: false,
		},
		{
			name:     "o3-mini does not support image generation",
			model:    "o3-mini",
			expected: false,
		},
		{
			name:     "gpt-4 does not support image generation",
			model:    "gpt-4",
			expected: false,
		},
		{
			name:     "gpt-3.5-turbo does not support image generation",
			model:    "gpt-3.5-turbo",
			expected: false,
		},
		{
			name:     "empty model does not support image generation",
			model:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := supportsImageGeneration(tt.model)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestModelValidationLogic(t *testing.T) {
	t.Run("Unsupported model with image file should return validation error", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Model:     "o1-mini",
			ImageFile: "/tmp/output.png",
		}

		// Test the validation logic directly
		if opts.ImageFile != "" && !supportsImageGeneration(opts.Model) {
			err := fmt.Errorf("model '%s' does not support image generation. Supported models: %s", opts.Model, strings.Join(ImageGenerationSupportedModels, ", "))

			assert.Contains(t, err.Error(), "does not support image generation")
			assert.Contains(t, err.Error(), "o1-mini")
			assert.Contains(t, err.Error(), "Supported models:")
		} else {
			t.Error("Expected validation to trigger")
		}
	})

	t.Run("Supported model with image file should not trigger validation", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Model:     "gpt-4o",
			ImageFile: "/tmp/output.png",
		}

		// Test the validation logic directly
		shouldFail := opts.ImageFile != "" && !supportsImageGeneration(opts.Model)
		assert.False(t, shouldFail, "Validation should not trigger for supported model")
	})

	t.Run("Unsupported model without image file should not trigger validation", func(t *testing.T) {
		opts := &domain.ChatOptions{
			Model:     "o1-mini",
			ImageFile: "", // No image file
		}

		// Test the validation logic directly
		shouldFail := opts.ImageFile != "" && !supportsImageGeneration(opts.Model)
		assert.False(t, shouldFail, "Validation should not trigger when no image file is specified")
	})
}

func TestAddImageGenerationToolWithUserParameters(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		opts     *domain.ChatOptions
		expected map[string]interface{}
	}{
		{
			name: "All parameters specified",
			opts: &domain.ChatOptions{
				ImageFile:        "/tmp/test.png",
				ImageSize:        "1536x1024",
				ImageQuality:     "high",
				ImageBackground:  "transparent",
				ImageCompression: 0, // Not applicable for PNG
			},
			expected: map[string]interface{}{
				"size":          "1536x1024",
				"quality":       "high",
				"background":    "transparent",
				"output_format": "png",
			},
		},
		{
			name: "JPEG with compression",
			opts: &domain.ChatOptions{
				ImageFile:        "/tmp/test.jpg",
				ImageSize:        "1024x1024",
				ImageQuality:     "medium",
				ImageBackground:  "opaque",
				ImageCompression: 75,
			},
			expected: map[string]interface{}{
				"size":               "1024x1024",
				"quality":            "medium",
				"background":         "opaque",
				"output_format":      "jpeg",
				"output_compression": int64(75),
			},
		},
		{
			name: "Only some parameters specified",
			opts: &domain.ChatOptions{
				ImageFile:    "/tmp/test.webp",
				ImageQuality: "low",
			},
			expected: map[string]interface{}{
				"quality":       "low",
				"output_format": "webp",
			},
		},
		{
			name: "No parameters specified (defaults)",
			opts: &domain.ChatOptions{
				ImageFile: "/tmp/test.png",
			},
			expected: map[string]interface{}{
				"output_format": "png",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tools := client.addImageGenerationTool(tt.opts, []responses.ToolUnionParam{})

			assert.Len(t, tools, 1)
			assert.NotNil(t, tools[0].OfImageGeneration)

			tool := tools[0].OfImageGeneration

			// Check required fields
			assert.Equal(t, "gpt-image-1", tool.Model)
			assert.Equal(t, tt.expected["output_format"], tool.OutputFormat)

			// Check optional fields
			if expectedSize, ok := tt.expected["size"]; ok {
				assert.Equal(t, expectedSize, tool.Size)
			} else {
				assert.Empty(t, tool.Size, "Size should not be set when not specified")
			}

			if expectedQuality, ok := tt.expected["quality"]; ok {
				assert.Equal(t, expectedQuality, tool.Quality)
			} else {
				assert.Empty(t, tool.Quality, "Quality should not be set when not specified")
			}

			if expectedBackground, ok := tt.expected["background"]; ok {
				assert.Equal(t, expectedBackground, tool.Background)
			} else {
				assert.Empty(t, tool.Background, "Background should not be set when not specified")
			}

			if expectedCompression, ok := tt.expected["output_compression"]; ok {
				assert.Equal(t, expectedCompression, tool.OutputCompression.Value)
			} else {
				assert.Equal(t, int64(0), tool.OutputCompression.Value, "Compression should not be set when not specified")
			}
		})
	}
}
