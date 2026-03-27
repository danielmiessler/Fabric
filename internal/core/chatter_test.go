package core

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// mockVendor implements the ai.Vendor interface for testing
type mockVendor struct {
	sendStreamError error
	streamChunks    []domain.StreamUpdate
	sendFunc        func(context.Context, []*chat.ChatCompletionMessage, *domain.ChatOptions) (string, error)
}

func (m *mockVendor) GetName() string {
	return "mock"
}

func (m *mockVendor) GetSetupDescription() string {
	return "mock vendor"
}

func (m *mockVendor) IsConfigured() bool {
	return true
}

func (m *mockVendor) Configure() error {
	return nil
}

func (m *mockVendor) Setup() error {
	return nil
}

func (m *mockVendor) SetupFillEnvFileContent(*bytes.Buffer) {
}

func (m *mockVendor) ListModels(context.Context) ([]string, error) {
	return []string{"test-model"}, nil
}

func (m *mockVendor) SendStream(_ context.Context, messages []*chat.ChatCompletionMessage, opts *domain.ChatOptions, responseChan chan domain.StreamUpdate) error {
	// Send chunks if provided (for successful streaming test)
	if m.streamChunks != nil {
		for _, chunk := range m.streamChunks {
			responseChan <- chunk
		}
	}
	// Close the channel like real vendors do
	close(responseChan)
	return m.sendStreamError
}

func (m *mockVendor) Send(ctx context.Context, messages []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	if m.sendFunc != nil {
		return m.sendFunc(ctx, messages, opts)
	}
	return "test response", nil
}

func (m *mockVendor) NeedsRawMode(modelName string) bool {
	return false
}

func TestJoinPromptSections(t *testing.T) {
	tests := []struct {
		name     string
		parts    []string
		expected string
	}{
		{
			name:     "multiple non-empty sections",
			parts:    []string{"STRATEGY", "CONTEXT", "PATTERN"},
			expected: "STRATEGY\nCONTEXT\nPATTERN",
		},
		{
			name:     "single section",
			parts:    []string{"only one"},
			expected: "only one",
		},
		{
			name:     "empty strings filtered out",
			parts:    []string{"first", "", "third"},
			expected: "first\nthird",
		},
		{
			name:     "whitespace-only strings filtered out",
			parts:    []string{"first", "   ", "\t\n", "last"},
			expected: "first\nlast",
		},
		{
			name:     "all empty returns empty",
			parts:    []string{"", "  ", "\n"},
			expected: "",
		},
		{
			name:     "no parts returns empty",
			parts:    []string{},
			expected: "",
		},
		{
			name:     "surrounding whitespace trimmed",
			parts:    []string{"  hello  ", "\nworld\n"},
			expected: "hello\nworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := joinPromptSections(tt.parts...)
			if got != tt.expected {
				t.Errorf("joinPromptSections(%v) = %q, want %q", tt.parts, got, tt.expected)
			}
		})
	}
}

func TestRecordFirstStreamError_NilError(t *testing.T) {
	errChan := make(chan error, 1)
	recordFirstStreamError(errChan, nil)
	select {
	case err := <-errChan:
		t.Fatalf("expected no error in channel, got %v", err)
	default:
		// Good — nil error should not be sent
	}
}

func TestRecordFirstStreamError_ChannelFull(t *testing.T) {
	errChan := make(chan error, 1)
	// Fill the channel with the first error
	errChan <- errors.New("first error")
	// Second error should be discarded (default branch + debug log)
	recordFirstStreamError(errChan, errors.New("second error"))
	// Channel should still contain the first error
	err := <-errChan
	if err.Error() != "first error" {
		t.Errorf("expected first error, got %q", err.Error())
	}
}

func TestChatter_Send_SuppressThink(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	mockVendor := &mockVendor{}

	chatter := &Chatter{
		db:     db,
		Stream: false,
		vendor: mockVendor,
		model:  "test-model",
	}

	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test",
		},
	}

	opts := &domain.ChatOptions{
		Model:         "test-model",
		SuppressThink: true,
		ThinkStartTag: "<think>",
		ThinkEndTag:   "</think>",
	}

	// custom send function returning a message with think tags
	mockVendor.sendFunc = func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
		return "<think>hidden</think> visible", nil
	}

	session, err := chatter.Send(context.Background(), request, opts)
	if err != nil {
		t.Fatalf("Send returned error: %v", err)
	}
	if session == nil {
		t.Fatal("expected session")
	}
	last := session.GetLastMessage()
	if last.Content != "visible" {
		t.Errorf("expected filtered content 'visible', got %q", last.Content)
	}
}

func TestChatter_BuildSession_SeparatesSystemSections(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	if err := os.MkdirAll(filepath.Join(db.Patterns.Dir, "test-pattern"), 0o755); err != nil {
		t.Fatalf("failed to create pattern directory: %v", err)
	}
	if err := os.MkdirAll(db.Contexts.Dir, 0o755); err != nil {
		t.Fatalf("failed to create context directory: %v", err)
	}

	patternPath := filepath.Join(db.Patterns.Dir, "test-pattern", "system.md")
	if err := os.WriteFile(patternPath, []byte("PATTERN"), 0o644); err != nil {
		t.Fatalf("failed to write pattern: %v", err)
	}

	contextPath := filepath.Join(db.Contexts.Dir, "test-context")
	if err := os.WriteFile(contextPath, []byte("CONTEXT"), 0o644); err != nil {
		t.Fatalf("failed to write context: %v", err)
	}

	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")
	if err := os.MkdirAll(strategyDir, 0o755); err != nil {
		t.Fatalf("failed to create strategy directory: %v", err)
	}

	strategyPath := filepath.Join(strategyDir, "test-strategy.json")
	if err := os.WriteFile(strategyPath, []byte(`{"prompt":"STRATEGY"}`), 0o644); err != nil {
		t.Fatalf("failed to write strategy: %v", err)
	}

	chatter := &Chatter{db: db}
	request := &domain.ChatRequest{
		ContextName:  "test-context",
		PatternName:  "test-pattern",
		StrategyName: "test-strategy",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "user input",
		},
	}

	session, err := chatter.BuildSession(request, false)
	if err != nil {
		t.Fatalf("BuildSession returned error: %v", err)
	}

	messages := session.GetVendorMessages()
	if len(messages) != 1 {
		t.Fatalf("expected 1 vendor message, got %d", len(messages))
	}

	systemMessage := messages[0]
	if systemMessage.Role != chat.ChatMessageRoleSystem {
		t.Fatalf("expected first message to be system, got %s", systemMessage.Role)
	}

	expectedSystemMessage := "STRATEGY\nCONTEXT\nPATTERN\nuser input"
	if systemMessage.Content != expectedSystemMessage {
		t.Fatalf("expected system message %q, got %q", expectedSystemMessage, systemMessage.Content)
	}

	if request.Message.Content != "user input" {
		t.Fatalf("expected request user input to remain unchanged, got %q", request.Message.Content)
	}
}

func TestChatter_Send_StreamingErrorPropagation(t *testing.T) {
	// Create a temporary database for testing
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create a mock vendor that will return an error from SendStream
	expectedError := errors.New("streaming error")
	mockVendor := &mockVendor{
		sendStreamError: expectedError,
	}

	// Create chatter with streaming enabled
	chatter := &Chatter{
		db:     db,
		Stream: true, // Enable streaming to trigger SendStream path
		vendor: mockVendor,
		model:  "test-model",
	}

	// Create a test request
	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	// Create test options
	opts := &domain.ChatOptions{
		Model: "test-model",
	}

	// Call Send and expect it to return the streaming error
	session, err := chatter.Send(context.Background(), request, opts)

	// Verify that the error from SendStream is propagated
	if err == nil {
		t.Fatal("Expected error to be returned, but got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error %q, but got %q", expectedError, err)
	}

	// Session should still be returned (it was built successfully before the streaming error)
	if session == nil {
		t.Error("Expected session to be returned even when streaming error occurs")
	}
}

func TestChatter_Send_StreamingErrorUpdateAndReturnDoesNotDeadlock(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	expectedError := errors.New("streaming error")
	mockVendor := &mockVendor{
		sendStreamError: expectedError,
		streamChunks: []domain.StreamUpdate{
			{
				Type:    domain.StreamTypeError,
				Content: "stream update error",
			},
		},
	}

	chatter := &Chatter{
		db:     db,
		Stream: true,
		vendor: mockVendor,
		model:  "test-model",
	}

	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	opts := &domain.ChatOptions{
		Model: "test-model",
	}

	type sendResult struct {
		session *fsdb.Session
		err     error
	}

	done := make(chan sendResult, 1)
	go func() {
		session, err := chatter.Send(context.Background(), request, opts)
		done <- sendResult{session: session, err: err}
	}()

	select {
	case result := <-done:
		if result.err == nil {
			t.Fatal("expected streaming error, got nil")
		}
		if result.session == nil {
			t.Fatal("expected session to be returned")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Send deadlocked when stream emitted an error update and returned an error")
	}
}

func TestChatter_Send_StreamingSuccessfulAggregation(t *testing.T) {
	// Create a temporary database for testing
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create test chunks that should be aggregated
	chunks := []string{"Hello", " ", "world", "!", " This", " is", " a", " test."}
	testChunks := make([]domain.StreamUpdate, len(chunks))
	for i, c := range chunks {
		testChunks[i] = domain.StreamUpdate{Type: domain.StreamTypeContent, Content: c}
	}
	expectedMessage := "Hello world! This is a test."

	// Create a mock vendor that will send chunks successfully
	mockVendor := &mockVendor{
		sendStreamError: nil, // No error for successful streaming
		streamChunks:    testChunks,
	}

	// Create chatter with streaming enabled
	chatter := &Chatter{
		db:     db,
		Stream: true, // Enable streaming to trigger SendStream path
		vendor: mockVendor,
		model:  "test-model",
	}

	// Create a test request
	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	// Create test options
	opts := &domain.ChatOptions{
		Model: "test-model",
	}

	// Call Send and expect successful aggregation
	session, err := chatter.Send(context.Background(), request, opts)

	// Verify no error occurred
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Verify session was returned
	if session == nil {
		t.Fatal("Expected session to be returned")
	}

	// Verify the message was aggregated correctly
	messages := session.GetVendorMessages()
	if len(messages) != 2 { // user message + assistant response
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}

	// Check the assistant's response (last message)
	assistantMessage := messages[len(messages)-1]
	if assistantMessage.Role != chat.ChatMessageRoleAssistant {
		t.Errorf("Expected assistant role, got %s", assistantMessage.Role)
	}

	if assistantMessage.Content != expectedMessage {
		t.Errorf("Expected aggregated message %q, got %q", expectedMessage, assistantMessage.Content)
	}
}

func TestChatter_Send_StreamingMetadataPropagation(t *testing.T) {
	// Create a temporary database for testing
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create test chunks: one content, one usage metadata
	testChunks := []domain.StreamUpdate{
		{
			Type:    domain.StreamTypeContent,
			Content: "Test content",
		},
		{
			Type: domain.StreamTypeUsage,
			Usage: &domain.UsageMetadata{
				InputTokens:  10,
				OutputTokens: 5,
				TotalTokens:  15,
			},
		},
	}

	// Create a mock vendor
	mockVendor := &mockVendor{
		sendStreamError: nil,
		streamChunks:    testChunks,
	}

	// Create chatter with streaming enabled
	chatter := &Chatter{
		db:     db,
		Stream: true,
		vendor: mockVendor,
		model:  "test-model",
	}

	// Create a test request
	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	// Create an update channel to capture stream events
	updateChan := make(chan domain.StreamUpdate, 10)

	// Create test options with UpdateChan
	opts := &domain.ChatOptions{
		Model:      "test-model",
		UpdateChan: updateChan,
		Quiet:      true, // Suppress stdout/stderr
	}

	// Call Send
	_, err := chatter.Send(context.Background(), request, opts)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	close(updateChan)

	// Verify we received the metadata event
	var usageReceived bool
	for update := range updateChan {
		if update.Type == domain.StreamTypeUsage {
			usageReceived = true
			if update.Usage == nil {
				t.Error("Expected usage metadata to be non-nil")
			} else {
				if update.Usage.TotalTokens != 15 {
					t.Errorf("Expected 15 total tokens, got %d", update.Usage.TotalTokens)
				}
			}
		}
	}

	if !usageReceived {
		t.Error("Expected to receive a usage metadata update, but didn't")
	}
}
