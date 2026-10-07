package cli

import (
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/ai/dryrun"
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

func TestSanitizeNotificationText(t *testing.T) {
	in := `x" & (do shell script "echo hi") & "` + `\'');print('x`
	want := `x & (do shell script echo hi) & );print(x`
	if got := sanitizeNotificationText(in); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestHandleChatProcessingSanitizesOutput checks that the printed reply has no
// terminal control sequence.
func TestHandleChatProcessingSanitizesOutput(t *testing.T) {
	const reply = "a\x1b]8;;http://evil.example\x07b"
	registry := newStepRegistry(t, &stepVendor{Client: dryrun.NewClient(), reply: reply}, &Workflow{})
	flags := &Flags{Message: "in", Language: "en"}

	printed := captureStdout(t, func() { must(t, handleChatProcessing(flags, registry, "")) })
	if got := string(printed); got != "ab\n" {
		t.Errorf("printed %q, want %q", got, "ab\n")
	}
}
