// Package claudecode provides a Fabric vendor plugin that delegates inference
// to the locally-installed Claude Code CLI (`claude`).
package claudecode

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
)

const defaultBinary = "claude"

// streamEvent is a partial representation of the stream-json lines emitted by
// `claude --print --output-format stream-json --include-partial-messages`.
type streamEvent struct {
	Type    string `json:"type"`
	Message *struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"message"`
	Delta *struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"delta"`
	Event *struct {
		Type  string `json:"type"`
		Delta *struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"delta"`
	} `json:"event"`
	Usage *struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

type Client struct {
	*plugins.PluginBase
	BinaryPath *plugins.SetupQuestion
}

func NewClient() (ret *Client) {
	ret = &Client{}
	ret.PluginBase = plugins.NewVendorPluginBase("ClaudeCode", ret.configure)
	ret.BinaryPath = ret.AddSetupQuestion("Binary Path", false)
	ret.BinaryPath.Value = defaultBinary
	return
}

func (c *Client) configure() error {
	return nil
}

// IsConfigured returns true only if the claude binary is found in PATH (or at
// the configured path), meaning the user has Claude Code installed and logged in.
func (c *Client) IsConfigured() bool {
	_, err := exec.LookPath(c.getBinary())
	return err == nil
}

func (c *Client) getBinary() string {
	if c.BinaryPath != nil && c.BinaryPath.Value != "" {
		debuglog.Debug(debuglog.Detailed, "ClaudeCode using configured binary path: %s\n", c.BinaryPath.Value)
		return c.BinaryPath.Value
	}
	debuglog.Debug(debuglog.Detailed, "ClaudeCode using default binary path: %s\n", defaultBinary)
	return defaultBinary
}

func (c *Client) ListModels(_ context.Context) ([]string, error) {
	return []string{
		"claude-opus-4-6",
		"claude-sonnet-4-6",
		"claude-opus-4-5-20251101",
		"claude-opus-4-5",
		"claude-haiku-4-5",
		"claude-haiku-4-5-20251001",
		"claude-sonnet-4-5",
		"claude-sonnet-4-5-20250929",
		"claude-opus-4-1-20250805",
		"claude-sonnet-4-20250514",
		"claude-sonnet-4-0",
		"claude-4-sonnet-20250514",
		"claude-opus-4-0",
		"claude-opus-4-20250514",
		"claude-4-opus-20250514",
	}, nil
}

// extractMessages splits the message list into a system prompt string and a
// formatted user prompt string suitable for piping into the claude CLI.
// If opts.ImageFile is set, the file path is appended to the prompt so Claude's
// Read tool can load it as vision content (requires --add-dir on the directory).
func (c *Client) extractMessages(msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (system string, userPrompt string) {
	var systemParts []string
	var conversationParts []string

	for _, msg := range msgs {
		content := strings.TrimSpace(c.messageText(msg))
		if content == "" {
			continue
		}
		switch msg.Role {
		case chat.ChatMessageRoleSystem:
			systemParts = append(systemParts, content)
		case chat.ChatMessageRoleUser:
			conversationParts = append(conversationParts, content)
		case chat.ChatMessageRoleAssistant:
			conversationParts = append(conversationParts, "Assistant: "+content)
		}
	}

	system = strings.Join(systemParts, "\n\n")
	userPrompt = strings.Join(conversationParts, "\n\n")
	if userPrompt == "" && system != "" {
		// Fabric patterns can inline the user input into a single system message.
		// Claude still needs a prompt argument, so fall back to that composed text.
		userPrompt = system
		system = ""
	}
	if opts.ImageFile != "" {
		userPrompt += "\n\nPlease read and analyze the image at: " + opts.ImageFile
	}
	return
}

// messageText returns the text content of a message, handling both the plain
// Content field and the MultiContent slice (images are skipped since the claude
// CLI does not support image input via stdin).
func (c *Client) messageText(msg *chat.ChatCompletionMessage) string {
	if len(msg.MultiContent) == 0 {
		return msg.Content
	}
	var parts []string
	for _, part := range msg.MultiContent {
		if part.Type == chat.ChatMessagePartTypeText && strings.TrimSpace(part.Text) != "" {
			parts = append(parts, part.Text)
		}
	}
	return strings.Join(parts, "\n")
}

func (c *Client) buildArgs(opts *domain.ChatOptions, system string) []string {
	args := []string{"--print", "--no-session-persistence"}
	if opts.Model != "" {
		args = append(args, "--model", opts.Model)
	}
	if system != "" {
		args = append(args, "--system-prompt", system)
	}
	switch opts.Thinking {
	case domain.ThinkingLow:
		args = append(args, "--effort", "low")
	case domain.ThinkingMedium:
		args = append(args, "--effort", "medium")
	case domain.ThinkingHigh:
		args = append(args, "--effort", "high")
	}
	if opts.ImageFile != "" {
		args = append(args, "--add-dir", filepath.Dir(opts.ImageFile))
	}
	return args
}

func (c *Client) Send(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions) (string, error) {
	system, userPrompt := c.extractMessages(msgs, opts)
	if userPrompt == "" {
		return "", fmt.Errorf("claude: no prompt content after message extraction")
	}

	c.logUnsupportedOptions(opts)

	args := c.buildArgs(opts, system)
	args = append(args, userPrompt)
	binary := c.getBinary()
	debuglog.Debug(debuglog.Detailed, "ClaudeCode Send launching: %s %s\n", binary, truncateArgs(args))
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Env = cleanEnv()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errOut := strings.TrimSpace(stderr.String())
		if errOut == "" {
			errOut = strings.TrimSpace(stdout.String())
		}
		if errOut != "" {
			return "", fmt.Errorf("claude: %w\n%s", err, errOut)
		}
		return "", fmt.Errorf("claude: %w", err)
	}

	return strings.TrimSpace(stdout.String()), nil
}

func (c *Client) SendStream(ctx context.Context, msgs []*chat.ChatCompletionMessage, opts *domain.ChatOptions, channel chan domain.StreamUpdate) error {
	defer close(channel)

	system, userPrompt := c.extractMessages(msgs, opts)
	if userPrompt == "" {
		return fmt.Errorf("claude: no prompt content after message extraction")
	}

	c.logUnsupportedOptions(opts)

	args := c.buildArgs(opts, system)
	args = append(args, "--verbose", "--output-format", "stream-json", "--include-partial-messages")
	args = append(args, userPrompt)

	binary := c.getBinary()
	debuglog.Debug(debuglog.Detailed, "ClaudeCode SendStream launching: %s %s\n", binary, truncateArgs(args))
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Env = cleanEnv()

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("claude: stdout pipe: %w", err)
	}

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("claude: start: %w", err)
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if update, ok := parseStreamEvent(line); ok {
			select {
			case channel <- update:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("claude: scan stdout: %w", err)
	}

	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("claude: %w\n%s", err, stderr.String())
	}

	return nil
}

func (c *Client) NeedsRawMode(_ string) bool {
	return false
}

// parseStreamEvent parses a single stream-json line and returns a StreamUpdate
// if it contains a text delta or usage metadata. Returns false if the line
// should be skipped.
func parseStreamEvent(line string) (domain.StreamUpdate, bool) {
	var event streamEvent
	if err := json.Unmarshal([]byte(line), &event); err != nil {
		return domain.StreamUpdate{}, false
	}

	if event.Delta != nil && event.Delta.Type == "text_delta" && event.Delta.Text != "" {
		return domain.StreamUpdate{Type: domain.StreamTypeContent, Content: event.Delta.Text}, true
	}
	if event.Event != nil && event.Event.Delta != nil && event.Event.Delta.Type == "text_delta" && event.Event.Delta.Text != "" {
		return domain.StreamUpdate{Type: domain.StreamTypeContent, Content: event.Event.Delta.Text}, true
	}

	if event.Type == "message" && event.Message != nil && event.Usage != nil {
		return domain.StreamUpdate{
			Type: domain.StreamTypeUsage,
			Usage: &domain.UsageMetadata{
				InputTokens:  event.Usage.InputTokens,
				OutputTokens: event.Usage.OutputTokens,
				TotalTokens:  event.Usage.InputTokens + event.Usage.OutputTokens,
			},
		}, true
	}

	return domain.StreamUpdate{}, false
}

// cleanEnv returns os.Environ() with CLAUDECODE and all ANTHROPIC_* variables
// removed so the Claude subprocess uses local Claude Code session auth and does
// not detect nested Claude Code execution.
func cleanEnv() []string {
	env := os.Environ()
	filtered := env[:0:0]
	for _, e := range env {
		if strings.HasPrefix(e, "CLAUDECODE=") || strings.HasPrefix(e, "ANTHROPIC_") {
			continue
		}
		filtered = append(filtered, e)
	}
	return filtered
}

// logUnsupportedOptions emits debug warnings for ChatOptions fields that the
// Claude CLI does not support.
func (c *Client) logUnsupportedOptions(opts *domain.ChatOptions) {
	if opts.Temperature > 0 {
		debuglog.Debug(debuglog.Detailed, "ClaudeCode: Temperature option is not supported by the Claude CLI and will be ignored\n")
	}
	if opts.TopP > 0 {
		debuglog.Debug(debuglog.Detailed, "ClaudeCode: TopP option is not supported by the Claude CLI and will be ignored\n")
	}
	if opts.MaxTokens > 0 {
		debuglog.Debug(debuglog.Detailed, "ClaudeCode: MaxTokens option is not supported by the Claude CLI and will be ignored\n")
	}
	if opts.Seed > 0 {
		debuglog.Debug(debuglog.Detailed, "ClaudeCode: Seed option is not supported by the Claude CLI and will be ignored\n")
	}
}

// truncateArgs returns a string representation of args with the last element
// (the prompt) truncated to 200 characters to avoid noisy or sensitive log output.
func truncateArgs(args []string) string {
	if len(args) == 0 {
		return ""
	}
	display := make([]string, len(args))
	copy(display, args)
	last := display[len(display)-1]
	if len(last) > 200 {
		display[len(display)-1] = last[:200] + "..."
	}
	return strings.Join(display, " ")
}
