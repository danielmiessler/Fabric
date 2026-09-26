package claudecode

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

func TestCommand(t *testing.T) {
	t.Setenv("ANTHROPIC_API_KEY", "secret")
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "be terse"},
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
	}
	cmd, err := command(context.Background(), msgs, &domain.ChatOptions{Model: "sonnet", Thinking: domain.ThinkingHigh}, "--verbose")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"--system-prompt", "be terse", "--model", "sonnet", "--effort", "high", "--verbose"}
	// (--tools "" removed to allow Read tool for file access)
	if !slices.Equal(cmd.Args[len(cmd.Args)-len(want):], want) {
		t.Errorf("args = %q", cmd.Args)
	}
	if slices.ContainsFunc(cmd.Env, func(kv string) bool { return kv == "ANTHROPIC_API_KEY=secret" }) {
		t.Error("ANTHROPIC_API_KEY leaked into the child environment")
	}

	// A system-only message list becomes the prompt.
	cmd, _ = command(context.Background(), msgs[:1], &domain.ChatOptions{})
	if i := slices.Index(cmd.Args, "--system-prompt"); cmd.Args[i+1] != "You are a helpful assistant." {
		t.Errorf("system prompt = %q, want the default", cmd.Args[i+1])
	}

	// Image files are exposed via --add-dir.
	cmd, _ = command(context.Background(), msgs, &domain.ChatOptions{ImageFile: "/tmp/image.png"})
	if !slices.Contains(cmd.Args, "--add-dir") || !slices.Contains(cmd.Args, "/tmp") {
		t.Errorf("expected --add-dir /tmp in args: %v", cmd.Args)
	}

	// Text parts in MultiContent are included; image parts are skipped in text().
	multi := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, MultiContent: []chat.ChatMessagePart{
		{Type: chat.ChatMessagePartTypeText, Text: "from parts"},
	}}}
	cmd, _ = command(context.Background(), multi, &domain.ChatOptions{})
	if s := cmd.Stdin.(*strings.Reader); s.Len() != len("from parts") {
		t.Errorf("stdin length = %d", s.Len())
	}

	// Local file paths in image_url attachments get --add-dir.
	multi[0].MultiContent = append(multi[0].MultiContent, chat.ChatMessagePart{
		Type:     chat.ChatMessagePartTypeImageURL,
		ImageURL: &chat.ChatMessageImageURL{URL: "/path/to/image.jpg"},
	})
	cmd, _ = command(context.Background(), multi, &domain.ChatOptions{})
	if !slices.Contains(cmd.Args, "--add-dir") || !slices.Contains(cmd.Args, "/path/to") {
		t.Errorf("expected --add-dir /path/to for local file image: %v", cmd.Args)
	}
}

func TestTextDelta(t *testing.T) {
	text, ok := textDelta([]byte(`{"type":"stream_event","event":{"type":"content_block_delta","delta":{"type":"text_delta","text":"Hey"}}}`))
	if !ok || text != "Hey" {
		t.Errorf("got %q, %v", text, ok)
	}
	for _, line := range []string{
		`{"type":"stream_event","event":{"type":"content_block_delta","delta":{"type":"thinking_delta","thinking":"x"}}}`,
		`{"type":"assistant","message":{"content":[{"type":"text","text":"Hey"}]}}`,
		`not json`,
	} {
		if _, ok := textDelta([]byte(line)); ok {
			t.Errorf("unexpected text from %s", line)
		}
	}
}
