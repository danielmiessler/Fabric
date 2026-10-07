// Package claudecode runs Fabric requests through the Claude Code CLI.
// The CLI uses the local Claude subscription login, so no API key is necessary.
package claudecode

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai/anthropic"
)

const binary = "claude"

type Client struct {
	*plugins.PluginBase
}

func NewClient() *Client {
	return &Client{PluginBase: plugins.NewVendorPluginBase("ClaudeCode", nil)}
}

// IsConfigured is true when the claude binary is on PATH.
func (c *Client) IsConfigured() bool {
	_, err := exec.LookPath(binary)
	return err == nil
}

// ListModels returns the same model names as the Anthropic vendor.
func (c *Client) ListModels(ctx context.Context) ([]string, error) {
	return anthropic.NewClient().ListModels(ctx)
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	cmd, err := command(ctx, msgs, opts)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(cmd.Dir)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fail(err, &stderr)
	}
	return strings.TrimSpace(string(out)), nil
}

func (c *Client) SendStream(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)
	cmd, err := command(ctx, msgs, opts, "--verbose", "--output-format", "stream-json", "--include-partial-messages")
	if err != nil {
		return err
	}
	defer os.RemoveAll(cmd.Dir)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if scanErr := scanStreamJSON(stdout, channel); scanErr != nil {
		// Stop the child first. It can block on a full stdout pipe and Wait would not return.
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
		return fmt.Errorf("claude: reading stream output: %w", scanErr)
	}
	if err := cmd.Wait(); err != nil {
		return fail(err, &stderr)
	}
	return nil
}

// scanStreamJSON reads one stream-json event per line and forwards text deltas
// to channel. It returns scanner.Err() so a truncated line or a broken pipe
// surfaces as an error instead of a silently short answer.
func scanStreamJSON(stdout io.Reader, channel chan domain.StreamUpdate) error {
	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(nil, 4<<20) // the final result line holds the full answer
	for scanner.Scan() {
		if text, ok := textDelta(scanner.Bytes()); ok {
			channel <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: text}
		}
	}
	return scanner.Err()
}

// command builds the claude invocation. System messages go to --system-prompt.
// All other messages go to stdin. If every message is a system message, that
// text becomes the prompt. Image attachments and opts.ImageFile add their
// directories with --add-dir so the Read tool can open them.
func command(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, extra ...string) (*exec.Cmd, error) {
	var system, prompt []string
	for _, m := range msgs {
		content, err := text(m)
		if err != nil {
			return nil, err
		}
		if m.Role == chat.ChatMessageRoleSystem {
			system = append(system, content)
		} else if content != "" {
			prompt = append(prompt, content) // ponytail: assistant turns are unlabelled; label roles if sessions matter
		}
	}
	if len(prompt) == 0 {
		prompt, system = system, nil
	}
	if len(system) == 0 {
		system = []string{"You are a helpful assistant."} // an empty --system-prompt selects the coding-agent persona
	}
	args := []string{"--print", "--no-session-persistence", "--setting-sources", "", "--strict-mcp-config",
		"--system-prompt", strings.Join(system, "\n\n")}
	if opts.Model != "" {
		args = append(args, "--model", opts.Model)
	}
	switch opts.Thinking {
	case domain.ThinkingLow, domain.ThinkingMedium, domain.ThinkingHigh:
		args = append(args, "--effort", string(opts.Thinking))
	}
	// Run the CLI in a new private folder, not in the working folder of
	// Fabric. The CLI read tools can open files in their working folder,
	// for example a .env file. If the folder cannot be made, stop. The
	// caller removes cmd.Dir after the command. Remove the folder before an
	// error return after MkdirTemp, because then the caller does not.
	tmpDir, err := os.MkdirTemp("", "claudecode-*")
	if err != nil {
		return nil, fmt.Errorf("claudecode: cannot make a working folder: %w", err)
	}
	dirs := make(map[string]bool)
	for _, m := range msgs {
		for _, p := range m.MultiContent {
			if p.Type == chat.ChatMessagePartTypeImageURL && p.ImageURL != nil && p.ImageURL.URL != "" {
				url := p.ImageURL.URL
				// The -a flag sends a local image as a data: URL. Write it to a temp file.
				// Use only a data: URL. Do not use a bare local path from a
				// message as an --add-dir folder, because a REST client can
				// write message content. The -a flag sends a local image as a
				// data: URL, and only the CLI sets opts.ImageFile.
				if strings.HasPrefix(url, "data:") {
					tmpFile, err := decodeDataURL(url, tmpDir)
					if err != nil {
						os.RemoveAll(tmpDir)
						return nil, fmt.Errorf("claudecode: image not sent: %w", err)
					}
					dirs[tmpDir] = true
					prompt = append(prompt, "Please analyze the file at: "+filepath.Base(tmpFile))
				}
			}
		}
	}
	if opts.ImageFile != "" {
		imagePath := opts.ImageFile
		if strings.HasPrefix(imagePath, "~/") {
			if u, err := user.Current(); err == nil {
				imagePath = filepath.Join(u.HomeDir, imagePath[2:])
			}
		}
		dirs[filepath.Dir(imagePath)] = true
		prompt = append(prompt, "Please analyze the file at: "+filepath.Base(imagePath))
	}
	for dir := range dirs {
		args = append(args, "--add-dir", dir)
	}
	cmd := exec.CommandContext(ctx, binary, append(args, extra...)...)
	cmd.Stdin = strings.NewReader(strings.Join(prompt, "\n\n"))
	cmd.Dir = tmpDir
	// Drop ANTHROPIC_* so the CLI bills the subscription, not an API key from Fabric's .env.
	cmd.Env = slices.DeleteFunc(os.Environ(), func(kv string) bool { return strings.HasPrefix(kv, "ANTHROPIC_") })
	return cmd, nil
}

// text joins the message content and its text parts. command handles image parts.
func text(m *chat.ChatCompletionMessage) (string, error) {
	parts := []string{m.Content}
	for _, p := range m.MultiContent {
		if p.Type == chat.ChatMessagePartTypeText {
			parts = append(parts, p.Text)
		}
	}
	return strings.TrimSpace(strings.Join(parts, "\n")), nil
}

// decodeDataURL writes a base64 data: URL to a temp file and returns its path.
// Format: data:image/jpeg;base64,<data>
func decodeDataURL(dataURL string, tmpDir string) (string, error) {
	parts := strings.Split(dataURL, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid data URL format")
	}

	header := parts[0]
	ext := ".jpg"
	if strings.Contains(header, "image/png") {
		ext = ".png"
	} else if strings.Contains(header, "image/gif") {
		ext = ".gif"
	} else if strings.Contains(header, "image/webp") {
		ext = ".webp"
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	f, err := os.CreateTemp(tmpDir, "image_*"+ext)
	if err != nil {
		return "", err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return "", err
	}
	return f.Name(), f.Close()
}

// textDelta extracts the text from one stream-json line, if it carries any.
func textDelta(line []byte) (string, bool) {
	var event struct {
		Type  string `json:"type"`
		Event struct {
			Delta struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"delta"`
		} `json:"event"`
	}
	if json.Unmarshal(line, &event) != nil || event.Type != "stream_event" || event.Event.Delta.Type != "text_delta" {
		return "", false
	}
	return event.Event.Delta.Text, true
}

func fail(err error, stderr *bytes.Buffer) error {
	return fmt.Errorf("claude: %w: %s", err, strings.TrimSpace(stderr.String()))
}
