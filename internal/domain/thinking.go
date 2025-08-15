package domain

// ThinkingLevel represents reasoning/thinking levels supported across providers.
type ThinkingLevel string

const (
	ThinkingOff    ThinkingLevel = "off"
	ThinkingLow    ThinkingLevel = "low"
	ThinkingMedium ThinkingLevel = "medium"
	ThinkingHigh   ThinkingLevel = "high"
)

// ThinkingBudgets defines standardized token budgets for reasoning-enabled models.
// The map assigns a maximum token count to each ThinkingLevel, representing the
// amount of context or computation that can be used for reasoning at that level.
// These values (e.g., 1024 for low, 2048 for medium, 4096 for high) are used to
// Token budget constants for each ThinkingLevel.
// These values are chosen to align with typical context window sizes for LLMs at different reasoning levels.
// Adjust these if model capabilities change.
const (
	// TokenBudgetLow is suitable for basic reasoning or smaller models (e.g., 1k context window).
	TokenBudgetLow int64 = 1024
	// TokenBudgetMedium is suitable for intermediate reasoning or mid-sized models (e.g., 2k context window).
	TokenBudgetMedium int64 = 2048
	// TokenBudgetHigh is suitable for advanced reasoning or large models (e.g., 4k context window).
	TokenBudgetHigh int64 = 4096
)

// ThinkingBudgets defines standardized token budgets for reasoning-enabled models.
var ThinkingBudgets = map[ThinkingLevel]int64{
	ThinkingLow:    TokenBudgetLow,
	ThinkingMedium: TokenBudgetMedium,
	ThinkingHigh:   TokenBudgetHigh,
}
