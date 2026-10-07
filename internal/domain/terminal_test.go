package domain

import "testing"

func TestSanitizeTerminalOutput(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"plain text and whitespace stay", "a\tb\nc\rd", "a\tb\nc\rd"},
		{"SGR color stays", "\x1b[1;31mred\x1b[0m", "\x1b[1;31mred\x1b[0m"},
		{"OSC 52 clipboard write is removed", "a\x1b]52;c;aGVsbG8=\x07b", "ab"},
		{"OSC 8 link with ST is removed", "\x1b]8;;http://x\x1b\\link\x1b]8;;\x1b\\", "link"},
		{"DCS is removed", "a\x1bPq#0;2;0;0;0\x1b\\b", "ab"},
		{"cursor and screen control is removed", "a\x1b[2J\x1b[1;1Hb\x1b[?25l", "ab"},
		{"other ESC sequences are removed", "a\x1b7\x1b(Bb\x1bc", "ab"},
		{"C0, DEL and C1 are removed", "a\x07\x00\x7f\u0090\u009bb", "ab"},
		{"unterminated OSC is removed", "a\x1b]52;c;aGVsbG8=", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeTerminalOutput(tt.in); got != tt.want {
				t.Errorf("SanitizeTerminalOutput(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestSplitOpenEscape checks that a sequence that two stream chunks divide
// does not show its bytes as text.
func TestSplitOpenEscape(t *testing.T) {
	cases := []struct {
		chunks []string
		want   string
	}{
		{[]string{"a\x1b", "[31mb"}, "a\x1b[31mb"},
		{[]string{"a\x1b[3", "1mb"}, "a\x1b[31mb"},
		{[]string{"a\x1b]8;;http://ev", "il.example\x07b"}, "ab"},
		{[]string{"a\x1b[31mb", "c"}, "a\x1b[31mbc"},
		{[]string{"a\x1b["}, "a"},
	}
	for _, tc := range cases {
		got, held := "", ""
		for _, c := range tc.chunks {
			var done string
			done, held = SplitOpenEscape(held + c)
			got += SanitizeTerminalOutput(done)
		}
		got += SanitizeTerminalOutput(held)
		if got != tc.want {
			t.Errorf("chunks %q: got %q, want %q", tc.chunks, got, tc.want)
		}
	}
}
