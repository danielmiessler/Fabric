package domain

// ThinkingLevel represents reasoning/thinking levels supported across providers.
type ThinkingLevel string

const (
	ThinkingOff    ThinkingLevel = "off"
	ThinkingLow    ThinkingLevel = "low"
	ThinkingMedium ThinkingLevel = "medium"
	ThinkingHigh   ThinkingLevel = "high"
)

// Token budgets for each ThinkingLevel. Providers send them as the thinking
// budget of a request. Adjust them when model capabilities change.
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
