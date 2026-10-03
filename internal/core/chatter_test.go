package core

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// mockVendor implements ai.Vendor.
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
	if m.streamChunks != nil {
		for _, chunk := range m.streamChunks {
			responseChan <- chunk
		}
	}
	// Real vendors close the channel. Send reads it until the close.
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
	}
}

func TestRecordFirstStreamError_ChannelFull(t *testing.T) {
	errChan := make(chan error, 1)
	errChan <- errors.New("first error")
	recordFirstStreamError(errChan, errors.New("second error"))
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
	if len(messages) != 2 {
		t.Fatalf("expected 2 vendor messages, got %d", len(messages))
	}

	systemMessage := messages[0]
	if systemMessage.Role != chat.ChatMessageRoleSystem {
		t.Fatalf("expected first message to be system, got %s", systemMessage.Role)
	}

	// The system message has the strategy, the context, and the pattern.
	// It must not contain the user input.
	expectedSystemMessage := "STRATEGY\nCONTEXT\nPATTERN"
	if systemMessage.Content != expectedSystemMessage {
		t.Fatalf("expected system message %q, got %q", expectedSystemMessage, systemMessage.Content)
	}

	// The user input goes in the user message, one time only.
	userMessage := messages[1]
	if userMessage.Role != chat.ChatMessageRoleUser {
		t.Fatalf("expected second message to be user, got %s", userMessage.Role)
	}
	if userMessage.Content != "user input" {
		t.Fatalf("expected user message %q, got %q", "user input", userMessage.Content)
	}

	if request.Message.Content != "user input" {
		t.Fatalf("expected request user input to remain unchanged, got %q", request.Message.Content)
	}
}

func TestChatter_BuildSession_EndsWithUserMessage(t *testing.T) {
	db := fsdb.NewDb(t.TempDir())
	for name, content := range map[string]string{"plain": "PATTERN", "inline": "PATTERN\n{{input}}"} {
		dir := filepath.Join(db.Patterns.Dir, name)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "system.md"), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name, pattern, input string
		raw                  bool
		want                 []string // "role:content" for each message
	}{
		{"plain pattern", "plain", "IN", false, []string{"system:PATTERN", "user:IN"}},
		{"inline pattern", "inline", "IN", false, []string{"user:PATTERN\nIN"}},
		{"no input", "plain", "", false, []string{"user:PATTERN"}},
		{"raw plain pattern", "plain", "IN", true, []string{"user:PATTERN\nIN"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &domain.ChatRequest{
				PatternName: tt.pattern,
				Message:     &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: tt.input},
			}
			session, err := (&Chatter{db: db}).BuildSession(request, tt.raw)
			if err != nil {
				t.Fatal(err)
			}
			var got []string
			for _, m := range session.GetVendorMessages() {
				got = append(got, m.Role+":"+m.Content)
			}
			if strings.Join(got, "|") != strings.Join(tt.want, "|") {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestChatter_Send_StreamingErrorPropagation(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	expectedError := errors.New("streaming error")
	mockVendor := &mockVendor{
		sendStreamError: expectedError,
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

	session, err := chatter.Send(context.Background(), request, opts)

	if err == nil {
		t.Fatal("Expected error to be returned, but got nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error %q, but got %q", expectedError, err)
	}

	// BuildSession succeeded before the stream failed, so Send returns the session with the error.
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
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

	chunks := []string{"Hello", " ", "world", "!", " This", " is", " a", " test."}
	testChunks := make([]domain.StreamUpdate, len(chunks))
	for i, c := range chunks {
		testChunks[i] = domain.StreamUpdate{Type: domain.StreamTypeContent, Content: c}
	}
	expectedMessage := "Hello world! This is a test."

	mockVendor := &mockVendor{
		sendStreamError: nil,
		streamChunks:    testChunks,
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

	session, err := chatter.Send(context.Background(), request, opts)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if session == nil {
		t.Fatal("Expected session to be returned")
	}

	messages := session.GetVendorMessages()
	if len(messages) != 2 { // user message + assistant response
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}

	assistantMessage := messages[len(messages)-1]
	if assistantMessage.Role != chat.ChatMessageRoleAssistant {
		t.Errorf("Expected assistant role, got %s", assistantMessage.Role)
	}

	if assistantMessage.Content != expectedMessage {
		t.Errorf("Expected aggregated message %q, got %q", expectedMessage, assistantMessage.Content)
	}
}

func TestChatter_Send_StreamingBufferStreamDoesNotPrint(t *testing.T) {
	db := fsdb.NewDb(t.TempDir())

	chunks := []domain.StreamUpdate{
		{Type: domain.StreamTypeContent, Content: "Here:\n```go\n"},
		{Type: domain.StreamTypeContent, Content: "x := 1\n```\n"},
	}
	chatter := &Chatter{
		db:     db,
		Stream: true,
		vendor: &mockVendor{streamChunks: chunks},
		model:  "test-model",
	}
	request := &domain.ChatRequest{
		Message: &chat.ChatCompletionMessage{
			Role:    chat.ChatMessageRoleUser,
			Content: "test message",
		},
	}
	opts := &domain.ChatOptions{
		Model:        "test-model",
		BufferStream: true,
	}

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	session, sendErr := chatter.Send(context.Background(), request, opts)

	w.Close()
	os.Stdout = oldStdout
	printed, _ := io.ReadAll(r)

	if sendErr != nil {
		t.Fatalf("Expected no error, but got: %v", sendErr)
	}
	if len(printed) != 0 {
		t.Errorf("Expected no stdout output while buffering, got %q", printed)
	}
	if got, want := session.GetLastMessage().Content, "Here:\n```go\nx := 1\n```\n"; got != want {
		t.Errorf("Expected full buffered message %q, got %q", want, got)
	}
}

func TestChatter_Send_StreamingMetadataPropagation(t *testing.T) {
	tempDir := t.TempDir()
	db := fsdb.NewDb(tempDir)

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

	mockVendor := &mockVendor{
		sendStreamError: nil,
		streamChunks:    testChunks,
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

	updateChan := make(chan domain.StreamUpdate, 10)

	opts := &domain.ChatOptions{
		Model:      "test-model",
		UpdateChan: updateChan,
		Quiet:      true,
	}

	_, err := chatter.Send(context.Background(), request, opts)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	close(updateChan)

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
