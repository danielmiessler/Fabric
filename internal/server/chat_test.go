package restapi

import "testing"

func TestBuildPromptChatRequest_PreservesStrategyAndUserInput(t *testing.T) {
	prompt := PromptRequest{
		UserInput:    "user input",
		Vendor:       "TestVendor",
		Model:        "test-model",
		ContextName:  "ctx",
		PatternName:  "pattern",
		StrategyName: "strategy",
		SessionName:  "session",
		Variables: map[string]string{
			"topic": "pipelines",
		},
	}

	request := buildPromptChatRequest(prompt, "en")

	if request.Message == nil {
		t.Fatal("expected request message to be set")
	}
	if request.Message.Content != "user input" {
		t.Fatalf("expected user input to stay unchanged, got %q", request.Message.Content)
	}
	if request.StrategyName != "strategy" {
		t.Fatalf("expected strategy name to be preserved, got %q", request.StrategyName)
	}
	if request.PatternName != "pattern" {
		t.Fatalf("expected pattern name to be preserved, got %q", request.PatternName)
	}
	if request.ContextName != "ctx" {
		t.Fatalf("expected context name to be preserved, got %q", request.ContextName)
	}
	if request.SessionName != "session" {
		t.Fatalf("expected session name to be preserved, got %q", request.SessionName)
	}
	if request.Language != "en" {
		t.Fatalf("expected language to be preserved, got %q", request.Language)
	}
	if got := request.PatternVariables["topic"]; got != "pipelines" {
		t.Fatalf("expected variables to be preserved, got %q", got)
	}
}
