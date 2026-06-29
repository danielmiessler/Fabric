package openai

import (
	"strings"
	"testing"
)

func TestParseSSEAndConcat_JSONDeltas(t *testing.T) {
	s := `data: {"id":"chatcmpl-1","choices":[{"delta":{"content":"A fox"}}]}
data: {"id":"chatcmpl-1","choices":[{"delta":{"content":" swiftly"}}]}
data: {"id":"chatcmpl-1","choices":[{"delta":{"content":" leaps"}}]}
data: {"id":"chatcmpl-1","choices":[{"delta":{},"finish_reason":"stop"}]}
data: [DONE]
`
	r := strings.NewReader(s)
	got, err := parseSSEAndConcat(r)
	if err != nil {
		t.Fatalf("parseSSEAndConcat returned error: %v", err)
	}
	want := "A fox swiftly leaps"
	if got != want {
		t.Fatalf("unexpected result\nwant: %q\ngot:  %q", want, got)
	}
}

func TestParseSSEAndConcat_PlainText(t *testing.T) {
	s := `data: Hello
data: world
data: [DONE]
`
	r := strings.NewReader(s)
	got, err := parseSSEAndConcat(r)
	if err != nil {
		t.Fatalf("parseSSEAndConcat returned error: %v", err)
	}
	// parseSSEAndConcat trims data lines, so expected concatenation has no extra spaces
	want := "Helloworld"
	if got != want {
		t.Fatalf("unexpected result\nwant: %q\ngot:  %q", want, got)
	}
}
