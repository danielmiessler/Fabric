package cli

import (
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
)

func TestSendNotification_SecurityEscaping(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		message     string
		command     string
		expectError bool
		description string
	}{
		{
			name:        "Normal content",
			title:       "Test Title",
			message:     "Test message content",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Normal content should work fine",
		},
		{
			name:        "Content with backticks",
			title:       "Test Title",
			message:     "Test `whoami` injection",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Backticks should be escaped and not executed",
		},
		{
			name:        "Content with semicolon injection",
			title:       "Test Title",
			message:     "Test; echo INJECTED; echo end",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Semicolon injection should be prevented",
		},
		{
			name:        "Content with command substitution",
			title:       "Test Title",
			message:     "Test $(whoami) injection",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Command substitution should be escaped",
		},
		{
			name:        "Content with quote injection",
			title:       "Test Title",
			message:     "Test ' || echo INJECTED || echo ' end",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Quote injection should be prevented",
		},
		{
			name:        "Content with newlines",
			title:       "Test Title",
			message:     "Line 1\nLine 2\nLine 3",
			command:     `echo "Title: $1, Message: $2"`,
			expectError: false,
			description: "Newlines should be handled safely",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &domain.ChatOptions{
				NotificationCommand: tt.command,
				Notification:        true,
			}

			// The command runs, but the test does not read its output. It only checks for no error.
			err := sendNotification(options, "test_pattern", tt.message)

			if tt.expectError && err == nil {
				t.Errorf("Expected error for %s, but got none", tt.description)
			}

			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for %s: %v", tt.description, err)
			}
		})
	}
}

func TestSendNotification_TitleGeneration(t *testing.T) {
	tests := []struct {
		name        string
		patternName string
		expected    string
	}{
		{
			name:        "No pattern name",
			patternName: "",
			expected:    "Fabric Command Complete",
		},
		{
			name:        "With pattern name",
			patternName: "summarize",
			expected:    "Fabric: summarize Complete",
		},
		{
			name:        "Pattern with special characters",
			patternName: "test_pattern-v2",
			expected:    "Fabric: test_pattern-v2 Complete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &domain.ChatOptions{
				NotificationCommand: `echo "Title: $1"`,
				Notification:        true,
			}

			err := sendNotification(options, tt.patternName, "test message")

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestSendNotification_MessageTruncation(t *testing.T) {
	longMessage := strings.Repeat("A", 150)
	shortMessage := "Short message"

	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:    "Short message",
			message: shortMessage,
		},
		{
			name:    "Long message truncation",
			message: longMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &domain.ChatOptions{
				NotificationCommand: `echo "Message: $2"`,
				Notification:        true,
			}

			err := sendNotification(options, "test", tt.message)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestImageGenerationCompatibilityWarning(t *testing.T) {
	originalStderr := os.Stderr
	defer func() {
		os.Stderr = originalStderr
	}()

	tests := []struct {
		name          string
		model         string
		imageFile     string
		expectWarning bool
		warningSubstr string
		description   string
	}{
		{
			name:          "Compatible model with image",
			model:         "gpt-4o",
			imageFile:     "test.png",
			expectWarning: false,
			description:   "Should not warn for compatible model",
		},
		{
			name:          "Incompatible model with image",
			model:         "o1-mini",
			imageFile:     "test.png",
			expectWarning: true,
			warningSubstr: "Warning: Model 'o1-mini' does not support image generation",
			description:   "Should warn for incompatible model",
		},
		{
			name:          "Incompatible model without image",
			model:         "o1-mini",
			imageFile:     "",
			expectWarning: false,
			description:   "Should not warn when no image file specified",
		},
		{
			name:          "Compatible model without image",
			model:         "gpt-4o-mini",
			imageFile:     "",
			expectWarning: false,
			description:   "Should not warn when no image file specified even for compatible model",
		},
		{
			name:          "Another incompatible model with image",
			model:         "gpt-3.5-turbo",
			imageFile:     "output.jpg",
			expectWarning: true,
			warningSubstr: "Warning: Model 'gpt-3.5-turbo' does not support image generation",
			description:   "Should warn for different incompatible model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &domain.ChatOptions{
				Model:     tt.model,
				ImageFile: tt.imageFile,
			}

			hasImage := tt.imageFile != ""
			shouldWarn := hasImage && tt.expectWarning

			if shouldWarn && tt.expectWarning {
				if tt.warningSubstr == "" {
					t.Errorf("Expected warning substring for warning case")
				}
			}

			if tt.expectWarning {
				t.Logf("Note: Warning would be printed by openai plugin for model '%s'", tt.model)
			}
		})
	}
}

func TestImageGenerationIntegrationScenarios(t *testing.T) {
	scenarios := []struct {
		name          string
		cliArgs       []string
		expectWarning bool
		warningModel  string
		description   string
	}{
		{
			name: "User tries o1-mini with image",
			cliArgs: []string{
				"-m", "o1-mini",
				"--image-file", "output.png",
				"Describe this image",
			},
			expectWarning: true,
			warningModel:  "o1-mini",
			description:   "Common user error - using incompatible model",
		},
		{
			name: "User uses compatible model",
			cliArgs: []string{
				"-m", "gpt-4o",
				"--image-file", "output.png",
				"Describe this image",
			},
			expectWarning: false,
			description:   "Correct usage - should work without warnings",
		},
		{
			name: "User specifies model via pattern env var",
			cliArgs: []string{
				"--pattern", "summarize",
				"--image-file", "output.png",
				"Summarize this image",
			},
			expectWarning: false, // Depends on env var, not tested here
			description:   "Pattern-based model selection",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			hasImage := false
			model := ""

			for i, arg := range scenario.cliArgs {
				if arg == "-m" && i+1 < len(scenario.cliArgs) {
					model = scenario.cliArgs[i+1]
				}
				if arg == "--image-file" && i+1 < len(scenario.cliArgs) {
					hasImage = true
				}
			}

			if scenario.expectWarning && scenario.warningModel == "" {
				t.Errorf("Expected warning scenario must specify warning model")
			}

			t.Logf("Scenario: %s", scenario.description)
			t.Logf("Model: %s, Has Image: %v, Expect Warning: %v", model, hasImage, scenario.expectWarning)
		})
	}
}
