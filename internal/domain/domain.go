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
	ShowMetadata        bool
	Quiet               bool
	UpdateChan          chan StreamUpdate `json:"-"`
}

// NormalizeInputShape ensures every message array has at least one user-role
// message. Some LLM backends (vLLM, certain Bedrock endpoints) reject
// requests that contain only system-role messages.
//
// When no user-role message exists, the last non-nil message (regardless of role)
// is promoted to user. This keeps the instructional content semantically similar
// while satisfying the API contract.
//
// The function is idempotent: arrays that already contain a user message are
// returned unchanged (returning the original slice). When no user message
// exists, a new slice is returned with the last non-nil message cloned and
// its role changed to user.
func NormalizeInputShape(msgs []*chat.ChatCompletionMessage) []*chat.ChatCompletionMessage {
	if len(msgs) == 0 {
		return msgs
	}

	// Scan for existing user message, skipping nil entries.
	for _, msg := range msgs {
		if msg != nil && msg.Role == chat.ChatMessageRoleUser {
			return msgs
		}
	}

	// Find the last non-nil message to promote.
	var lastIdx int
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i] != nil {
			lastIdx = i
			break
		}
	}
	// If all entries are nil, return original slice (nothing to promote).
	if lastIdx < 0 || msgs[lastIdx] == nil {
		return msgs
	}

	// Build a new slice so the caller's original is not mutated (ret).
	// This prevents side effects when the caller reuses the slice
	// as session history or passes it to multiple consumers.
	ret := make([]*chat.ChatCompletionMessage, len(msgs))
	copy(ret, msgs)

	// Clone the last non-nil message (orig, shallow copy is sufficient since we only change Role).
	orig := ret[lastIdx]
	newMsg := *orig
	newMsg.Role = chat.ChatMessageRoleUser
	ret[lastIdx] = &newMsg
	return ret
}

// NormalizeMessages iterates over messages to enforce the odd-position rule for user
// messages. Empty messages are dropped. When an even position would not contain a user
// message, a synthetic user message with the provided default content is inserted.
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
