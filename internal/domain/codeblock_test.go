package domain

import "testing"

func TestExtractFencedCodeBlock(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		last      bool
		want      string
		wantFound bool
	}{
		{
			name:  "no code block",
			input: "Just some prose.\nNo code here.",
			want:  "",
		},
		{
			name:      "backtick fence with language tag",
			input:     "Here you go:\n```python\ndef f():\n    return 1\n```\nHope that helps.",
			want:      "def f():\n    return 1",
			wantFound: true,
		},
		{
			name:      "tilde fence",
			input:     "~~~\necho hi\n~~~",
			want:      "echo hi",
			wantFound: true,
		},
		{
			name:      "first block by default",
			input:     "```go\nfirst\n```\ntext\n```go\nsecond\n```",
			want:      "first",
			wantFound: true,
		},
		{
			name:      "last block when requested",
			input:     "```go\nfirst\n```\ntext\n```go\nsecond\n```",
			last:      true,
			want:      "second",
			wantFound: true,
		},
		{
			name:      "four backtick fence containing a three backtick fence",
			input:     "````markdown\n```bash\nls\n```\n````",
			want:      "```bash\nls\n```",
			wantFound: true,
		},
		{
			name:      "closing fence may be longer than opening fence",
			input:     "```\ncode\n`````\nafter",
			want:      "code",
			wantFound: true,
		},
		{
			name:      "tilde line does not close backtick fence",
			input:     "```\n~~~\ncode\n```",
			want:      "~~~\ncode",
			wantFound: true,
		},
		{
			name:      "closing fence with info string is content",
			input:     "```\n```js\n```",
			want:      "```js",
			wantFound: true,
		},
		{
			name:      "unterminated final fence runs to end of text",
			input:     "Intro\n```sh\nmake build\nmake test",
			want:      "make build\nmake test",
			wantFound: true,
		},
		{
			name:      "unterminated fence after a closed block is used for last",
			input:     "```\none\n```\n```\ntwo",
			last:      true,
			want:      "two",
			wantFound: true,
		},
		{
			name:  "inline triple backticks are not a fence",
			input: "Use ```code``` inline.",
			want:  "",
		},
		{
			name:  "two backticks are not a fence",
			input: "``\nnot code\n``",
			want:  "",
		},
		{
			name:      "fence indented up to three spaces",
			input:     "   ```\n   indented\n   ```",
			want:      "   indented",
			wantFound: true,
		},
		{
			name:  "fence indented four spaces is not a fence",
			input: "    ```\n    code\n    ```",
			want:  "",
		},
		{
			name:      "CRLF line endings",
			input:     "text\r\n```go\r\nx := 1\r\n```\r\n",
			want:      "x := 1",
			wantFound: true,
		},
		{
			name:      "empty block",
			input:     "```\n```",
			want:      "",
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := ExtractFencedCodeBlock(tt.input, tt.last)
			if got != tt.want || found != tt.wantFound {
				t.Errorf("ExtractFencedCodeBlock(%q, %v) = (%q, %v), want (%q, %v)",
					tt.input, tt.last, got, found, tt.want, tt.wantFound)
			}
		})
	}
}
