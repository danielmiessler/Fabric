package claudecode

import (
	"bufio"
	"context"
	"errors"
	"os"
	"os/exec"
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
	cmd := newCommand(t, msgs, &domain.ChatOptions{Model: "sonnet", Thinking: domain.ThinkingHigh}, "--verbose")
	want := []string{"--system-prompt", "be terse", "--model", "sonnet", "--effort", "high", "--verbose"}
	if !slices.Equal(cmd.Args[len(cmd.Args)-len(want):], want) {
		t.Errorf("args = %q", cmd.Args)
	}
	if slices.ContainsFunc(cmd.Env, func(kv string) bool { return kv == "ANTHROPIC_API_KEY=secret" }) {
		t.Error("ANTHROPIC_API_KEY leaked into the child environment")
	}

	// A system-only message list becomes the prompt.
	cmd = newCommand(t, msgs[:1], &domain.ChatOptions{})
	if i := slices.Index(cmd.Args, "--system-prompt"); cmd.Args[i+1] != "You are a helpful assistant." {
		t.Errorf("system prompt = %q, want the default", cmd.Args[i+1])
	}

	cmd = newCommand(t, msgs, &domain.ChatOptions{ImageFile: "/tmp/image.png"})
	if !slices.Contains(cmd.Args, "--add-dir") || !slices.Contains(cmd.Args, "/tmp") {
		t.Errorf("expected --add-dir /tmp in args: %v", cmd.Args)
	}

	// Text parts in MultiContent go to stdin.
	multi := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, MultiContent: []chat.ChatMessagePart{
		{Type: chat.ChatMessagePartTypeText, Text: "from parts"},
	}}}
	cmd = newCommand(t, multi, &domain.ChatOptions{})
	if s := cmd.Stdin.(*strings.Reader); s.Len() != len("from parts") {
		t.Errorf("stdin length = %d", s.Len())
	}

	// A bare local path in message content must not become an --add-dir folder.
	multi[0].MultiContent = append(multi[0].MultiContent, chat.ChatMessagePart{
		Type:     chat.ChatMessagePartTypeImageURL,
		ImageURL: &chat.ChatMessageImageURL{URL: "/some/private/dir/image.jpg"},
	})
	cmd = newCommand(t, multi, &domain.ChatOptions{})
	if slices.Contains(cmd.Args, "/some/private/dir") {
		t.Errorf("a message image path must not add a folder: %v", cmd.Args)
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

// newCommand calls command and removes its working folder after the test.
func newCommand(t *testing.T, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, extra ...string) *exec.Cmd {
	t.Helper()
	cmd, err := command(context.Background(), msgs, opts, extra...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(cmd.Dir) })
	return cmd
}

// The CLI runs in a new private folder, not in the Fabric working folder.
// Decoded images go into that folder, each in its own file.
func TestCommandUsesNewWorkingFolder(t *testing.T) {
	img := chat.ChatMessagePart{Type: chat.ChatMessagePartTypeImageURL,
		ImageURL: &chat.ChatMessageImageURL{URL: "data:image/png;base64,aGk="}}
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{img, img}}}
	cmd := newCommand(t, msgs, &domain.ChatOptions{})
	wd, _ := os.Getwd()
	if cmd.Dir == "" || cmd.Dir == wd {
		t.Fatalf("cmd.Dir = %q, want a new folder", cmd.Dir)
	}
	if filepath.Dir(cmd.Dir) != filepath.Clean(os.TempDir()) || !strings.HasPrefix(filepath.Base(cmd.Dir), "claudecode-") {
		t.Errorf("cmd.Dir = %q, want a MkdirTemp folder in %q", cmd.Dir, os.TempDir())
	}
	images, _ := filepath.Glob(filepath.Join(cmd.Dir, "image_*.png"))
	if len(images) != 2 {
		t.Errorf("images in cmd.Dir = %v, want 2", images)
	}
	if cmd2 := newCommand(t, msgs, &domain.ChatOptions{}); cmd2.Dir == cmd.Dir {
		t.Errorf("two calls used the same folder %q", cmd.Dir)
	}
}

// A bad data: URL gives an error, so that a REST client knows that the
// image was not sent.
func TestCommandRejectsBadImage(t *testing.T) {
	img := chat.ChatMessagePart{Type: chat.ChatMessagePartTypeImageURL,
		ImageURL: &chat.ChatMessageImageURL{URL: "data:image/png;base64"}}
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{img}}}
	if cmd, err := command(context.Background(), msgs, &domain.ChatOptions{}); err == nil {
		os.RemoveAll(cmd.Dir)
		t.Fatal("command() accepted a bad data URL")
	}
}
