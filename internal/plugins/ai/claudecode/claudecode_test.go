package claudecode

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

func TestNewClient_DefaultInitialization(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Fatal("expected non-nil client")
	}
	if c.BinaryPath == nil {
		t.Fatal("expected BinaryPath to be initialized")
	}
	if c.BinaryPath.Value != defaultBinary {
		t.Errorf("expected BinaryPath %q, got %q", defaultBinary, c.BinaryPath.Value)
	}
	if c.GetName() != "ClaudeCode" {
		t.Errorf("expected name ClaudeCode, got %q", c.GetName())
	}
}

func TestListModels(t *testing.T) {
	c := NewClient()
	models, err := c.ListModels(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(models) == 0 {
		t.Fatal("expected at least one model")
	}
	for _, m := range models {
		if !strings.HasPrefix(m, "claude-") {
			t.Errorf("expected model name to start with 'claude-', got %q", m)
		}
	}
}

func TestNeedsRawMode(t *testing.T) {
	c := NewClient()
	if c.NeedsRawMode("claude-sonnet-4-6") {
		t.Error("expected NeedsRawMode to return false")
	}
}

func TestGetBinary_Default(t *testing.T) {
	c := NewClient()
	if c.getBinary() != defaultBinary {
		t.Errorf("expected %q, got %q", defaultBinary, c.getBinary())
	}
}

func TestGetBinary_Custom(t *testing.T) {
	c := NewClient()
	c.BinaryPath.Value = "/usr/local/bin/claude"
	if c.getBinary() != "/usr/local/bin/claude" {
		t.Errorf("expected custom path, got %q", c.getBinary())
	}
}

func TestGetBinary_EmptyFallsBackToDefault(t *testing.T) {
	c := NewClient()
	c.BinaryPath.Value = ""
	if c.getBinary() != defaultBinary {
		t.Errorf("expected fallback to %q, got %q", defaultBinary, c.getBinary())
	}
}

func TestIsConfigured_KnownBinary(t *testing.T) {
	c := NewClient()
	c.BinaryPath.Value = "ls" // always exists on unix
	if !c.IsConfigured() {
		t.Error("expected IsConfigured to return true for 'ls'")
	}
}

func TestIsConfigured_MissingBinary(t *testing.T) {
	c := NewClient()
	c.BinaryPath.Value = "definitely_not_a_real_binary_xyz"
	if c.IsConfigured() {
		t.Error("expected IsConfigured to return false for non-existent binary")
	}
}

func TestMessageText_PlainContent(t *testing.T) {
	c := NewClient()
	msg := &chat.ChatCompletionMessage{
		Role:    chat.ChatMessageRoleUser,
		Content: "hello world",
	}
	if got := c.messageText(msg); got != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", got)
	}
}

func TestMessageText_MultiContentTextOnly(t *testing.T) {
	c := NewClient()
	msg := &chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "first"},
			{Type: chat.ChatMessagePartTypeText, Text: "second"},
		},
	}
	got := c.messageText(msg)
	if got != "first\nsecond" {
		t.Errorf("expected %q, got %q", "first\nsecond", got)
	}
}

func TestMessageText_MultiContentSkipsImages(t *testing.T) {
	c := NewClient()
	msg := &chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "describe this"},
			{Type: chat.ChatMessagePartTypeImageURL, ImageURL: &chat.ChatMessageImageURL{URL: "data:image/png;base64,abc"}},
		},
	}
	got := c.messageText(msg)
	if got != "describe this" {
		t.Errorf("expected only text part, got %q", got)
	}
}

func TestMessageText_MultiContentBlankPartsSkipped(t *testing.T) {
	c := NewClient()
	msg := &chat.ChatCompletionMessage{
		Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{
			{Type: chat.ChatMessagePartTypeText, Text: "   "},
			{Type: chat.ChatMessagePartTypeText, Text: "real content"},
		},
	}
	got := c.messageText(msg)
	if got != "real content" {
		t.Errorf("expected %q, got %q", "real content", got)
	}
}

func TestExtractMessages_SystemAndUser(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "you are helpful"},
		{Role: chat.ChatMessageRoleUser, Content: "tell me a joke"},
	}
	system, userPrompt := c.extractMessages(msgs, &domain.ChatOptions{})
	if system != "you are helpful" {
		t.Errorf("expected system %q, got %q", "you are helpful", system)
	}
	if userPrompt != "tell me a joke" {
		t.Errorf("expected userPrompt %q, got %q", "tell me a joke", userPrompt)
	}
}

func TestExtractMessages_AssistantPrefixed(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
		{Role: chat.ChatMessageRoleAssistant, Content: "hi there"},
		{Role: chat.ChatMessageRoleUser, Content: "how are you"},
	}
	_, userPrompt := c.extractMessages(msgs, &domain.ChatOptions{})
	if !strings.Contains(userPrompt, "Assistant: hi there") {
		t.Errorf("expected assistant prefix in prompt, got %q", userPrompt)
	}
}

func TestExtractMessages_EmptyMessagesSkipped(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: ""},
		{Role: chat.ChatMessageRoleUser, Content: "  "},
	}
	system, userPrompt := c.extractMessages(msgs, &domain.ChatOptions{})
	if system != "" || userPrompt != "" {
		t.Errorf("expected empty results for blank messages, got system=%q userPrompt=%q", system, userPrompt)
	}
}

func TestExtractMessages_ImageFileAppendsPath(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "what is this?"},
	}
	opts := &domain.ChatOptions{ImageFile: "/tmp/photo.jpg"}
	_, userPrompt := c.extractMessages(msgs, opts)
	if !strings.Contains(userPrompt, "/tmp/photo.jpg") {
		t.Errorf("expected image path in prompt, got %q", userPrompt)
	}
}

func TestExtractMessages_MultipleSystemsJoined(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "be concise"},
		{Role: chat.ChatMessageRoleSystem, Content: "use markdown"},
		{Role: chat.ChatMessageRoleUser, Content: "hi"},
	}
	system, _ := c.extractMessages(msgs, &domain.ChatOptions{})
	if !strings.Contains(system, "be concise") || !strings.Contains(system, "use markdown") {
		t.Errorf("expected both system messages joined, got %q", system)
	}
}

func TestExtractMessages_SystemOnlyFallsBackToUserPrompt(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "pattern instructions with embedded input"},
	}
	system, userPrompt := c.extractMessages(msgs, &domain.ChatOptions{})
	if system != "" {
		t.Fatalf("expected system prompt to be cleared when used as fallback prompt, got %q", system)
	}
	if userPrompt != "pattern instructions with embedded input" {
		t.Fatalf("expected fallback user prompt, got %q", userPrompt)
	}
}

func TestBuildArgs_BaseFlags(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{}, "")
	if !slices.Contains(args, "--print") {
		t.Error("expected --print flag")
	}
	if !slices.Contains(args, "--no-session-persistence") {
		t.Error("expected --no-session-persistence flag")
	}
}

func TestBuildArgs_WithModel(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{Model: "claude-sonnet-4-6"}, "")
	if !containsPair(args, "--model", "claude-sonnet-4-6") {
		t.Errorf("expected --model claude-sonnet-4-6 in args: %v", args)
	}
}

func TestBuildArgs_WithSystemPrompt(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{}, "be helpful")
	if !containsPair(args, "--system-prompt", "be helpful") {
		t.Errorf("expected --system-prompt in args: %v", args)
	}
}

func TestBuildArgs_NoSystemPromptWhenEmpty(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{}, "")
	if slices.Contains(args, "--system-prompt") {
		t.Error("expected no --system-prompt when system is empty")
	}
}

func TestBuildArgs_ThinkingLevels(t *testing.T) {
	c := NewClient()
	cases := []struct {
		level    domain.ThinkingLevel
		expected string
	}{
		{domain.ThinkingLow, "low"},
		{domain.ThinkingMedium, "medium"},
		{domain.ThinkingHigh, "high"},
	}
	for _, tc := range cases {
		args := c.buildArgs(&domain.ChatOptions{Thinking: tc.level}, "")
		if !containsPair(args, "--effort", tc.expected) {
			t.Errorf("thinking %v: expected --effort %s in args: %v", tc.level, tc.expected, args)
		}
	}
}

func TestBuildArgs_ThinkingOffOmitsEffort(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{Thinking: domain.ThinkingOff}, "")
	if slices.Contains(args, "--effort") {
		t.Errorf("expected no --effort for ThinkingOff, got: %v", args)
	}
}

func TestBuildArgs_ImageFileAddsDir(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{ImageFile: "/tmp/images/photo.jpg"}, "")
	if !containsPair(args, "--add-dir", "/tmp/images") {
		t.Errorf("expected --add-dir /tmp/images in args: %v", args)
	}
}

func TestBuildArgs_NoImageFileNoAddDir(t *testing.T) {
	c := NewClient()
	args := c.buildArgs(&domain.ChatOptions{}, "")
	if slices.Contains(args, "--add-dir") {
		t.Errorf("expected no --add-dir when ImageFile is empty, got: %v", args)
	}
}

func TestSend_EmptyPromptReturnsError(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "  "},
	}
	_, err := c.Send(context.Background(), msgs, &domain.ChatOptions{})
	if err == nil {
		t.Fatal("expected error for empty prompt")
	}
	if !strings.Contains(err.Error(), "no prompt content") {
		t.Fatalf("expected 'no prompt content' error, got %q", err.Error())
	}
}

func TestSendStream_EmptyPromptReturnsError(t *testing.T) {
	c := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleUser, Content: "  "},
	}
	ch := make(chan domain.StreamUpdate, 1)
	err := c.SendStream(context.Background(), msgs, &domain.ChatOptions{}, ch)
	if err == nil {
		t.Fatal("expected error for empty prompt")
	}
	if !strings.Contains(err.Error(), "no prompt content") {
		t.Fatalf("expected 'no prompt content' error, got %q", err.Error())
	}
}

func TestSend_ExecutesBinaryAndReturnsTrimmedOutput(t *testing.T) {
	t.Setenv("CLAUDECODE", "nested-session")

	script := writeFakeClaudeScript(t, `
if [ -n "${CLAUDECODE:-}" ]; then
  echo "CLAUDECODE should be unset" >&2
  exit 1
fi
printf "  reply from fake claude  \n"
`)

	c := NewClient()
	c.BinaryPath.Value = script

	msgs := []*chat.ChatCompletionMessage{{
		Role:    chat.ChatMessageRoleUser,
		Content: "hello",
	}}

	got, err := c.Send(t.Context(), msgs, &domain.ChatOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "reply from fake claude" {
		t.Fatalf("expected trimmed output, got %q", got)
	}
}

func TestSendStream_ParsesTextDeltas(t *testing.T) {
	script := writeFakeClaudeScript(t, `
echo '{"delta":{"type":"text_delta","text":"Hello"}}'
echo '{"type":"message"}'
echo '{"delta":{"type":"text_delta","text":" world"}}'
`)

	c := NewClient()
	c.BinaryPath.Value = script

	msgs := []*chat.ChatCompletionMessage{{
		Role:    chat.ChatMessageRoleUser,
		Content: "hello",
	}}

	ch := make(chan domain.StreamUpdate, 8)
	err := c.SendStream(context.Background(), msgs, &domain.ChatOptions{}, ch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parts []string
	for u := range ch {
		if u.Type == domain.StreamTypeContent {
			parts = append(parts, u.Content)
		}
	}

	got := strings.Join(parts, "")
	if got != "Hello world" {
		t.Fatalf("expected streamed content %q, got %q", "Hello world", got)
	}
}

func TestSendStream_ParsesNestedStreamEvents(t *testing.T) {
	script := writeFakeClaudeScript(t, `
echo '{"type":"stream_event","event":{"type":"content_block_delta","delta":{"type":"text_delta","text":"streaming"}}}'
echo '{"type":"stream_event","event":{"type":"content_block_delta","delta":{"type":"text_delta","text":"-ok"}}}'
`)

	c := NewClient()
	c.BinaryPath.Value = script

	msgs := []*chat.ChatCompletionMessage{{
		Role:    chat.ChatMessageRoleUser,
		Content: "hello",
	}}

	ch := make(chan domain.StreamUpdate, 8)
	err := c.SendStream(context.Background(), msgs, &domain.ChatOptions{}, ch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parts []string
	for u := range ch {
		if u.Type == domain.StreamTypeContent {
			parts = append(parts, u.Content)
		}
	}

	got := strings.Join(parts, "")
	if got != "streaming-ok" {
		t.Fatalf("expected streamed content %q, got %q", "streaming-ok", got)
	}
}

func TestSendStream_CommandFailureIncludesStderr(t *testing.T) {
	script := writeFakeClaudeScript(t, `
echo 'boom from fake claude' >&2
exit 7
`)

	c := NewClient()
	c.BinaryPath.Value = script

	msgs := []*chat.ChatCompletionMessage{{
		Role:    chat.ChatMessageRoleUser,
		Content: "hello",
	}}

	ch := make(chan domain.StreamUpdate, 1)
	err := c.SendStream(context.Background(), msgs, &domain.ChatOptions{}, ch)
	if err == nil {
		t.Fatal("expected error")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "boom from fake claude") {
		t.Fatalf("expected stderr in error, got %q", errStr)
	}
}

func TestCleanEnv_RemovesClaudeCodeAndAnthropicVars(t *testing.T) {
	t.Setenv("CLAUDECODE", "nested-session")
	t.Setenv("ANTHROPIC_API_KEY", "sk-test")
	t.Setenv("ANTHROPIC_BASE_URL", "https://example.com")
	t.Setenv("ANTHROPIC_AUTH_TOKEN", "tok-test")
	t.Setenv("FABRIC_TEST_KEEP", "1")

	env := cleanEnv()
	for _, kv := range env {
		if strings.HasPrefix(kv, "CLAUDECODE=") {
			t.Fatalf("CLAUDECODE should have been removed, got %q", kv)
		}
		if strings.HasPrefix(kv, "ANTHROPIC_") {
			t.Fatalf("ANTHROPIC_* vars should have been removed, got %q", kv)
		}
	}
	if !slices.Contains(env, "FABRIC_TEST_KEEP=1") {
		t.Fatalf("expected FABRIC_TEST_KEEP to remain in env")
	}
}

func writeFakeClaudeScript(t *testing.T, body string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "fake-claude.sh")
	content := fmt.Sprintf("#!/bin/sh\nset -eu\n%s\n", strings.TrimSpace(body))
	if err := os.WriteFile(path, []byte(content), 0o700); err != nil {
		t.Fatalf("write fake script: %v", err)
	}
	if err := os.Chmod(path, 0o700); err != nil {
		t.Fatalf("chmod fake script: %v", err)
	}
	return path
}

// containsPair reports whether flag followed immediately by value appears in the slice.
func containsPair(slice []string, flag, value string) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == flag && slice[i+1] == value {
			return true
		}
	}
	return false
}
