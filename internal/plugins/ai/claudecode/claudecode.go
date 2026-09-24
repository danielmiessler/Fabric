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
	defer cleanupTmpDir()
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
	defer cleanupTmpDir()
	cmd, err := command(ctx, msgs, opts, "--verbose", "--output-format", "stream-json", "--include-partial-messages")
	if err != nil {
		return err
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(nil, 4<<20) // the final result line holds the full answer
	for scanner.Scan() {
		if text, ok := textDelta(scanner.Bytes()); ok {
			channel <- domain.StreamUpdate{Type: domain.StreamTypeContent, Content: text}
		}
	}
	if err := cmd.Wait(); err != nil {
		return fail(err, &stderr)
	}
	return nil
}

// command builds the claude invocation. System messages go to --system-prompt,
// all other messages go to stdin. A pattern that inlines {{input}} produces only
// a system message, so that text becomes the prompt instead. When opts.ImageFile
// is set, the directory is exposed via --add-dir so the Read tool can access it.
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
	// Collect directories from image attachments and explicit ImageFile
	dirs := make(map[string]bool)
	tmpDir := filepath.Join(os.TempDir(), "claudecode")
	for _, m := range msgs {
		for _, p := range m.MultiContent {
			if p.Type == chat.ChatMessagePartTypeImageURL && p.ImageURL != nil && p.ImageURL.URL != "" {
				url := p.ImageURL.URL
				// Base64 images from -a are written to temp files
				if strings.HasPrefix(url, "data:") {
					tmpFile, err := decodeDataURL(url, tmpDir)
					if err == nil {
						dirs[tmpDir] = true
						prompt = append(prompt, "Please analyze the file at: "+filepath.Base(tmpFile))
					}
				} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
					// Local file paths become references
					dir := filepath.Dir(url)
					if dir != "" && dir != "." {
						dirs[dir] = true
					}
					prompt = append(prompt, "Please analyze the file at: "+filepath.Base(url))
				}
			}
		}
	}
	if opts.ImageFile != "" {
		imagePath := opts.ImageFile
		// Expand ~ to home directory
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
	// Drop ANTHROPIC_* so the CLI bills the subscription, not an API key from Fabric's .env.
	cmd.Env = slices.DeleteFunc(os.Environ(), func(kv string) bool { return strings.HasPrefix(kv, "ANTHROPIC_") })
	return cmd, nil
}

// text returns the message text. Image attachments are handled separately via decodeDataURL.
func text(m *chat.ChatCompletionMessage) (string, error) {
	parts := []string{m.Content}
	for _, p := range m.MultiContent {
		if p.Type == chat.ChatMessagePartTypeText {
			parts = append(parts, p.Text)
		}
	}
	return strings.TrimSpace(strings.Join(parts, "\n")), nil
}

// cleanupTmpDir removes the temp directory used for decoded images.
func cleanupTmpDir() {
	tmpDir := filepath.Join(os.TempDir(), "claudecode")
	_ = os.RemoveAll(tmpDir)
}

// decodeDataURL writes a base64 data: URL to a temp file and returns its path.
// Format: data:image/jpeg;base64,<data>
func decodeDataURL(dataURL string, tmpDir string) (string, error) {
	parts := strings.Split(dataURL, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid data URL format")
	}

	// Extract MIME type to guess file extension
	header := parts[0]
	ext := ".jpg"
	if strings.Contains(header, "image/png") {
		ext = ".png"
	} else if strings.Contains(header, "image/gif") {
		ext = ".gif"
	} else if strings.Contains(header, "image/webp") {
		ext = ".webp"
	}

	// Decode base64
	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	// Create temp dir if needed
	if err := os.MkdirAll(tmpDir, 0700); err != nil {
		return "", err
	}

	// Write to temp file
	tmpFile := filepath.Join(tmpDir, fmt.Sprintf("image_%d%s", os.Getpid(), ext))
	if err := os.WriteFile(tmpFile, data, 0600); err != nil {
		return "", err
	}

	return tmpFile, nil
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
