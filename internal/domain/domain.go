package domain

import "github.com/danielmiessler/fabric/internal/chat"

const ChatMessageRoleMeta = "meta"

// Default chat options. They must match the defaults in internal/cli/flags.go.
const (
	DefaultTemperature      = 0.7
	DefaultTopP             = 0.9
	DefaultPresencePenalty  = 0.0
	DefaultFrequencyPenalty = 0.0
)

type ChatRequest struct {
	ContextName           string
	SessionName           string
	PatternName           string
	PatternVariables      map[string]string
	Message               *chat.ChatCompletionMessage
	Language              string
	Meta                  string
	InputHasVars          bool
	NoVariableReplacement bool
	StrategyName          string
}

type ChatOptions struct {
	Model               string
	Temperature         float64
	TopP                float64
	PresencePenalty     float64
	FrequencyPenalty    float64
	Raw                 bool
	Seed                int
	Thinking            ThinkingLevel
	ModelContextLength  int
	MaxTokens           int
	Search              bool
	SearchLocation      string
	ImageFile           string
	ImageSize           string
	ImageQuality        string
	ImageCompression    int
	ImageBackground     string
	SuppressThink       bool
	ThinkStartTag       string
	ThinkEndTag         string
	AudioOutput         bool
	AudioFormat         string
	Voice               string
	Notification        bool
	NotificationCommand string
	ShowMetadata        bool
	Quiet               bool
	// SessionID is a stable identifier for the current conversation. Providers
	// that support session-based routing (e.g. OpenCode) forward it as a header.
	// It is derived internally and must not be part of the REST API surface.
	SessionID    string            `json:"-"`
	BufferStream bool              // collect streamed content without printing it
	UpdateChan   chan StreamUpdate `json:"-"`
}

// NormalizeMessages remove empty messages and ensure messages order user-assist-user
func NormalizeMessages(msgs []*chat.ChatCompletionMessage, defaultUserMessage string) (ret []*chat.ChatCompletionMessage) {
	fullMessageIndex := 0
	for _, message := range msgs {
		if message.Content == "" {
			// The Anthropic API rejects empty messages.
			continue
		}

		// The message at each even index must come from the user.
		// Insert a default user message when it does not.
		if fullMessageIndex%2 == 0 && message.Role != chat.ChatMessageRoleUser {
			ret = append(ret, &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: defaultUserMessage})
			fullMessageIndex++
		}
		ret = append(ret, message)
		fullMessageIndex++
	}
	return
}
