package restapi

import (
	"errors"
	"testing"
)

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

func TestUnreportedSendError(t *testing.T) {
	patternErr := errors.New("could not get pattern write_essay: missing required variable: author_name")

	tests := []struct {
		name        string
		sendErr     error
		sawError    bool
		wantReport  bool
		wantContent string
	}{
		{
			// A pattern that needs a variable the request does not give fails
			// before Send starts to stream, so no error reached the client yet.
			// This is the only chance to tell the client about it.
			name:        "error before the stream reaches the client",
			sendErr:     patternErr,
			sawError:    false,
			wantReport:  true,
			wantContent: "Error: " + patternErr.Error(),
		},
		{
			// Send gives a stream error to the update channel and also returns
			// it. The client has it already, so a second report would show the
			// same failure twice.
			name:       "error during the stream is not sent again",
			sendErr:    errors.New("vendor stream failed"),
			sawError:   true,
			wantReport: false,
		},
		{
			name:       "no error and no report",
			sendErr:    nil,
			sawError:   false,
			wantReport: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sendErrChan := make(chan error, 1)
			if test.sendErr != nil {
				sendErrChan <- test.sendErr
			}

			response, ok := unreportedSendError(sendErrChan, test.sawError)

			if ok != test.wantReport {
				t.Fatalf("want report=%v, got %v", test.wantReport, ok)
			}
			if !test.wantReport {
				return
			}
			if response.Type != "error" {
				t.Errorf("want type %q, got %q", "error", response.Type)
			}
			if response.Content != test.wantContent {
				t.Errorf("want content %q, got %q", test.wantContent, response.Content)
			}
		})
	}
}

// An empty channel must not block the response, which the client waits for.
func TestUnreportedSendErrorWithEmptyChannel(t *testing.T) {
	if _, ok := unreportedSendError(make(chan error, 1), false); ok {
		t.Error("want no report from an empty channel")
	}
}
