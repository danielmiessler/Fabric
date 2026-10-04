package claudecode

import (
	"bufio"
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
	"time"

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

	cmd, _ = command(context.Background(), msgs, &domain.ChatOptions{ImageFile: "/tmp/image.png"})
	if !slices.Contains(cmd.Args, "--add-dir") || !slices.Contains(cmd.Args, "/tmp") {
		t.Errorf("expected --add-dir /tmp in args: %v", cmd.Args)
	}

	// Text parts in MultiContent go to stdin.
	multi := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, MultiContent: []chat.ChatMessagePart{
		{Type: chat.ChatMessagePartTypeText, Text: "from parts"},
	}}}
	cmd, _ = command(context.Background(), multi, &domain.ChatOptions{})
	if s := cmd.Stdin.(*strings.Reader); s.Len() != len("from parts") {
		t.Errorf("stdin length = %d", s.Len())
	}

	multi[0].MultiContent = append(multi[0].MultiContent, chat.ChatMessagePart{
		Type:     chat.ChatMessagePartTypeImageURL,
		ImageURL: &chat.ChatMessageImageURL{URL: "/path/to/image.jpg"},
	})
	cmd, _ = command(context.Background(), multi, &domain.ChatOptions{})
	if !slices.Contains(cmd.Args, "--add-dir") || !slices.Contains(cmd.Args, "/path/to") {
		t.Errorf("expected --add-dir /path/to for local file image: %v", cmd.Args)
	}
}

func TestScanStreamJSONReturnsScanError(t *testing.T) {
	oversizedLine := strings.Repeat("a", 5<<20) // past the 4 MiB buffer
	channel := make(chan domain.StreamUpdate)
	err := scanStreamJSON(strings.NewReader(oversizedLine+"\n"), channel)
	if err == nil {
		t.Fatal("expected an error for a line over the scanner buffer, got nil")
	}
}

func TestSendStreamStopsChildOnScanError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("uses a shell script as the fake claude binary")
	}
	// The fake claude writes a line past the 4 MiB buffer and then writes forever.
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf '%5242880s\\n' x\nwhile :; do echo x; done\n"
	if err := os.WriteFile(filepath.Join(dir, binary), []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)

	channel := make(chan domain.StreamUpdate)
	done := make(chan error, 1)
	go func() {
		done <- NewClient().SendStream(context.Background(),
			[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}, &domain.ChatOptions{}, channel)
	}()
	go func() {
		for range channel {
		}
	}()
	select {
	case err := <-done:
		if !errors.Is(err, bufio.ErrTooLong) {
			t.Errorf("err = %v, want bufio.ErrTooLong", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("SendStream did not return after the scan error")
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
