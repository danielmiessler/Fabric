package openai

import (
	"fmt"
	"strings"
)

const openRouterVendorName = "OpenRouter"

type openRouterProviderRouting struct {
	Order          []string
	AllowFallbacks bool
}

func (o *Client) registerOpenRouterSetupQuestions() {
	if o.GetName() != openRouterVendorName {
		return
	}

	o.openRouterProviderOrder = o.AddSetupQuestionCustom(
		"provider_order",
		false,
		"Enter preferred OpenRouter providers in order (comma-separated, e.g. amazon-bedrock)",
	)
	o.openRouterAllowFallbacks = o.AddSetupQuestionCustomBool(
		"allow_fallbacks",
		false,
		"Allow OpenRouter to fall back to another provider when preferred providers are unavailable",
	)
}

func (o *Client) configureOpenRouterProviderRouting() error {
	o.openRouterProviderRouting = nil
	if o.GetName() != openRouterVendorName || o.openRouterProviderOrder == nil {
		return nil
	}

	order := parseCommaSeparatedValues(o.openRouterProviderOrder.Value)
	if len(order) == 0 {
		return nil
	}

	allowFallbacks := false
	if o.openRouterAllowFallbacks != nil && o.openRouterAllowFallbacks.Value != "" {
		parsed, err := parseFlexibleBool(o.openRouterAllowFallbacks.Value)
		if err != nil {
			return fmt.Errorf("parse OpenRouter allow_fallbacks: %w", err)
		}
		allowFallbacks = parsed
	}

	o.openRouterProviderRouting = &openRouterProviderRouting{
		Order:          order,
		AllowFallbacks: allowFallbacks,
	}
	return nil
}

func (o *Client) buildRequestExtraFields() map[string]any {
	extraFields := map[string]any{}
	if routing := o.openRouterProviderRouting; routing != nil {
		order := make([]string, len(routing.Order))
		copy(order, routing.Order)
		extraFields["provider"] = map[string]any{
			"order":           order,
			"allow_fallbacks": routing.AllowFallbacks,
		}
	}
	return extraFields
}

func parseCommaSeparatedValues(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func parseFlexibleBool(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value %q", value)
	}
}
