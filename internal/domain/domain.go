package domain

import "github.com/danielmiessler/fabric/internal/chat"

const ChatMessageRoleMeta = "meta"

// Default values for chat options (must match cli/flags.go defaults)
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
}

// NormalizeMessages remove empty messages and ensure messages order user-assist-user
func NormalizeMessages(msgs []*chat.ChatCompletionMessage, defaultUserMessage string) (ret []*chat.ChatCompletionMessage) {
	// Iterate over messages to enforce the odd position rule for user messages
	fullMessageIndex := 0
	for _, message := range msgs {
		if message.Content == "" {
			// Skip empty messages as the anthropic API doesn't accept them
			continue
		}

		// Ensure, that each odd position shall be a user message
		if fullMessageIndex%2 == 0 && message.Role != chat.ChatMessageRoleUser {
			ret = append(ret, &chat.ChatCompletionMessage{Role: chat.ChatMessageRoleUser, Content: defaultUserMessage})
			fullMessageIndex++
		}
		ret = append(ret, message)
		fullMessageIndex++
	}
	return
}
