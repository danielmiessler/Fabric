package core

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// mockVendor implements the ai.Vendor interface for testing
type mockVendor struct {
	sendStreamError error
	streamChunks    []string
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

func (m *mockVendor) ListModels() ([]string, error) {
	return []string{"test-model"}, nil
}

func (m *mockVendor) SendStream(messages []*chat.ChatCompletionMessage, opts *domain.ChatOptions, responseChan chan string) error {
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

func (m *mockVendor) GetProviderName() string {
	return "mock"
}

// Enhanced mock vendor for schema testing
type mockSchemaVendor struct {
	*mockVendor
	supportsSchema bool
}

func (m *mockSchemaVendor) GetProviderName() string {
	if m.supportsSchema {
		return "anthropic" // Return a provider that supports schema
	}
	return "mock"
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

	session, err := chatter.Send(request, opts)
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
	session, err := chatter.Send(request, opts)

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

func TestChatter_Send_StreamingSuccessfulAggregation(t *testing.T) {
	// Create a temporary database for testing
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create test chunks that should be aggregated
	testChunks := []string{"Hello", " ", "world", "!", " This", " is", " a", " test."}
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
	session, err := chatter.Send(request, opts)

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

// Test schema manager mock for testing schema-related functionality
type mockSchemaManager struct {
	transformError error
	parseError     error
	validateError  error
	parsedResponse string
}

func (m *mockSchemaManager) HandleSchemaTransformation(provider string, opts *domain.ChatOptions) error {
	return m.transformError
}

func (m *mockSchemaManager) HandleSchemaTransformationWithContext(provider string, opts *domain.ChatOptions, context map[string]interface{}) error {
	return m.transformError
}

func (m *mockSchemaManager) HandleResponseParsing(provider string, response string, opts *domain.ChatOptions) (string, error) {
	if m.parseError != nil {
		return "", m.parseError
	}
	if m.parsedResponse != "" {
		return m.parsedResponse, nil
	}
	return response, nil
}

func (m *mockSchemaManager) ValidateOutput(output string, schemaContent string, provider string) error {
	return m.validateError
}

func TestChatter_Send_EmptyRequest(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	mockVendor := &mockVendor{}
	chatter := &Chatter{
		db:     db,
		vendor: mockVendor,
		model:  "test-model",
	}

	// Test with request that has empty content message
	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "", // Empty content
		},
	}
	opts := &domain.ChatOptions{Model: "test-model"}

	// This should succeed because BuildSession creates the space message
	session, err := chatter.Send(request, opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Test response content should be the test response from mock
	last := session.GetLastMessage()
	if last.Content != "test response" {
		t.Errorf("Expected 'test response', got %q", last.Content)
	}
}

func TestChatter_Send_EmptyResponse(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	mockVendor := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return "", nil // Empty response
		},
	}

	chatter := &Chatter{
		db:     db,
		vendor: mockVendor,
		model:  "test-model",
	}

	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test",
		},
	}

	opts := &domain.ChatOptions{Model: "test-model"}

	session, err := chatter.Send(request, opts)
	if err == nil {
		t.Fatal("Expected error for empty response, got nil")
	}
	if !strings.Contains(err.Error(), "empty response") {
		t.Errorf("Expected 'empty response' error, got: %v", err)
	}
	if session != nil {
		t.Fatal("Expected nil session for empty response")
	}
}

func TestChatter_Send_RawMode(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	mockVendor := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return "raw response", nil
		},
	}

	chatter := &Chatter{
		db:     db,
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
		Model: "test-model",
		Raw:   true,
	}

	session, err := chatter.Send(request, opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// In raw mode, all messages should be in user role
	messages := session.GetVendorMessages()
	if len(messages) != 2 { // user + assistant
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}
}

func TestChatter_Send_DryRun(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	mockVendor := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return "<think>hidden</think> visible content", nil
		},
	}

	chatter := &Chatter{
		db:     db,
		vendor: mockVendor,
		model:  "test-model",
		DryRun: true,
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

	session, err := chatter.Send(request, opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// In dry run mode with SuppressThink=false, think blocks should be preserved
	last := session.GetLastMessage()
	// Since DryRun is true, think blocks should NOT be stripped
	expected := "<think>hidden</think> visible content"
	if last.Content != expected {
		t.Errorf("Expected %q, got %q", expected, last.Content)
	}
}

func TestChatter_BuildSession_EmptyRequest(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	chatter := &Chatter{db: db}

	// Test with completely empty request - BuildSession creates a default message
	request := &domain.ChatRequest{}

	session, err := chatter.BuildSession(request, false)
	// BuildSession should succeed - it creates a message with space content
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// The session should have a user message with space content
	messages := session.GetVendorMessages()
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(messages))
	}
	if messages[0].Content != " " {
		t.Errorf("Expected space content, got %q", messages[0].Content)
	}
}

func TestChatter_BuildSession_WithContext(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create a test context
	contextsDir := filepath.Join(tempDir, "contexts")
	err := os.MkdirAll(contextsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create contexts directory: %v", err)
	}

	// Create context without .md extension since fsdb looks for files without extension
	contextPath := filepath.Join(contextsDir, "test-context")
	contextContent := "Test context content"
	err = os.WriteFile(contextPath, []byte(contextContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write context file: %v", err)
	}

	chatter := &Chatter{db: db}

	request := &domain.ChatRequest{
		ContextName: "test-context",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	session, err := chatter.BuildSession(request, false)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Check that context was loaded into system message
	messages := session.GetVendorMessages()
	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages (system + user), got %d", len(messages))
	}

	systemMsg := messages[0]
	if systemMsg.Role != chat.ChatMessageRoleSystem {
		t.Errorf("Expected system role, got %s", systemMsg.Role)
	}
	if !strings.Contains(systemMsg.Content, contextContent) {
		t.Errorf("System message should contain context content")
	}
}

func TestChatter_BuildSession_WithPattern(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create a test pattern
	patternsDir := filepath.Join(tempDir, "patterns", "test-pattern")
	err := os.MkdirAll(patternsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create patterns directory: %v", err)
	}

	patternPath := filepath.Join(patternsDir, "system.md")
	patternContent := "# Test Pattern\n\nTest pattern content"
	err = os.WriteFile(patternPath, []byte(patternContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write pattern file: %v", err)
	}

	chatter := &Chatter{db: db}

	request := &domain.ChatRequest{
		PatternName: "test-pattern",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test input",
		},
	}

	session, err := chatter.BuildSession(request, false)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Check that pattern was loaded into system message
	messages := session.GetVendorMessages()
	if len(messages) != 1 { // Only system message, user input was consumed by pattern
		t.Fatalf("Expected 1 message (system), got %d", len(messages))
	}

	systemMsg := messages[0]
	if systemMsg.Role != chat.ChatMessageRoleSystem {
		t.Errorf("Expected system role, got %s", systemMsg.Role)
	}
	if !strings.Contains(systemMsg.Content, "Test pattern content") {
		t.Errorf("System message should contain pattern content")
	}
}

func TestChatter_BuildSession_WithMeta(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	chatter := &Chatter{db: db}

	request := &domain.ChatRequest{
		Meta: "meta information",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	session, err := chatter.BuildSession(request, false)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Check that meta was added
	messages := session.GetVendorMessages()
	// Meta messages don't appear in vendor messages, only user messages do
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message (user), got %d", len(messages))
	}

	// Meta messages are handled internally, check user message exists
	userMsg := messages[0]
	if userMsg.Role != chat.ChatMessageRoleUser {
		t.Errorf("Expected user role, got %s", userMsg.Role)
	}
	if userMsg.Content != "test message" {
		t.Errorf("Expected user content, got %s", userMsg.Content)
	}
}

func TestChatter_BuildSession_WithLanguage(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	chatter := &Chatter{db: db}

	request := &domain.ChatRequest{
		Language: "Spanish",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}

	session, err := chatter.BuildSession(request, false)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Check that language instruction was added to system message
	messages := session.GetVendorMessages()
	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages (system + user), got %d", len(messages))
	}

	systemMsg := messages[0]
	if systemMsg.Role != chat.ChatMessageRoleSystem {
		t.Errorf("Expected system role, got %s", systemMsg.Role)
	}
	if !strings.Contains(strings.ToLower(systemMsg.Content), "spanish") {
		t.Errorf("System message should contain language instruction")
	}
}

func TestChatter_BuildSession_RawMode(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	chatter := &Chatter{db: db}

	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
			MultiContent: []chat.ChatMessagePart{
				{Type: chat.ChatMessagePartTypeText, Text: "original text"},
				{Type: "image", Text: "image data"},
			},
		},
	}

	session, err := chatter.BuildSession(request, true) // raw mode = true
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// In raw mode, everything should be combined into user message
	messages := session.GetVendorMessages()
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message in raw mode, got %d", len(messages))
	}

	userMsg := messages[0]
	if userMsg.Role != chat.ChatMessageRoleUser {
		t.Errorf("Expected user role, got %s", userMsg.Role)
	}

	// Should preserve MultiContent with non-text parts
	if len(userMsg.MultiContent) != 2 {
		t.Errorf("Expected 2 MultiContent parts, got %d", len(userMsg.MultiContent))
	}
}

func TestChatter_Send_WithSchemaContent(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Use a mock vendor that supports schema
	baseMock := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return `{"name": "test", "value": 123}`, nil
		},
	}
	mockVendor := &mockSchemaVendor{
		mockVendor:     baseMock,
		supportsSchema: true,
	}

	chatter := &Chatter{
		db:     db,
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
		Model: "test-model",
	}

	session, err := chatter.Send(request, opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Response should be the JSON response
	last := session.GetLastMessage()
	if !strings.Contains(last.Content, "test") || !strings.Contains(last.Content, "123") {
		t.Errorf("Expected JSON response, got: %s", last.Content)
	}
}

func TestChatter_Send_CreateCodingFeaturePattern(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	// Create the pattern directory and file
	patternDir := filepath.Join(tempDir, "patterns", "create_coding_feature")
	err := os.MkdirAll(patternDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create pattern directory: %v", err)
	}
	patternFile := filepath.Join(patternDir, "system.md")
	err = os.WriteFile(patternFile, []byte("# Create Coding Feature\nPattern content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create pattern file: %v", err)
	}

	// Mock response that includes file changes
	response := "## Summary\nCreated a new feature\n\n## File Changes\n\n### CREATE: test.txt\n```\ntest content\n```\n\n### UPDATE: existing.txt\n```\nupdated content\n```\n\n## End"

	mockVendor := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return response, nil
		},
	}

	chatter := &Chatter{
		db:     db,
		vendor: mockVendor,
		model:  "test-model",
	}

	request := &domain.ChatRequest{
		PatternName: "create_coding_feature",
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "Create a test feature",
		},
	}

	opts := &domain.ChatOptions{Model: "test-model"}

	// Change to temp directory so file operations don't affect real files
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	session, err := chatter.Send(request, opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if session == nil {
		t.Fatal("Expected session")
	}

	// Response should be just the summary, file changes should be applied
	last := session.GetLastMessage()
	if !strings.Contains(last.Content, "Created a new feature") {
		t.Errorf("Expected summary in response, got: %s", last.Content)
	}
	if strings.Contains(last.Content, "File Changes") {
		t.Errorf("File changes should be removed from response, got: %s", last.Content)
	}
}

func TestChatter_Send_NonStreamingError(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	expectedError := errors.New("send error")
	mockVendor := &mockVendor{
		sendFunc: func(ctx context.Context, msgs []*chat.ChatCompletionMessage, o *domain.ChatOptions) (string, error) {
			return "", expectedError
		},
	}

	chatter := &Chatter{
		db:     db,
		vendor: mockVendor,
		model:  "test-model",
		Stream: false,
	}

	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test",
		},
	}

	opts := &domain.ChatOptions{Model: "test-model"}

	session, err := chatter.Send(request, opts)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected %v, got %v", expectedError, err)
	}
	// Session is still created even if Send fails, because BuildSession succeeded
	if session == nil {
		t.Error("Expected session to be returned (BuildSession succeeded)")
	}
}
