package domain

import "regexp"

var (
	// terminalControl matches one escape sequence or one control character:
	// CSI, OSC (to BEL or ST), DCS/SOS/PM/APC (to ST), other ESC sequences,
	// and C0, DEL and C1 characters other than tab, newline and carriage return.
	terminalControl = regexp.MustCompile(`\x1b\[[0-?]*[ -/]*[@-~]|\x1b\][^\x07\x1b]*(?:\x07|\x1b\\)?|\x1b[P^_X][^\x1b]*(?:\x1b\\)?|\x1b[ -/]*[0-~]?|[\x00-\x08\x0b\x0c\x0e-\x1f\x7f\x{80}-\x{9f}]`)
	// sgrSequence matches a complete SGR (color and style) sequence.
	sgrSequence = regexp.MustCompile(`^\x1b\[[0-9;:]*m$`)
	// openEscape matches an escape sequence at the end of the text that is
	// not complete: a lone ESC, or a CSI, OSC or DCS/SOS/PM/APC sequence
	// with no final byte.
	openEscape = regexp.MustCompile(`\x1b(?:\[[0-?]*[ -/]*|\][^\x07\x1b]*|[P^_X][^\x1b]*)?$`)
)

// SplitOpenEscape divides streamed text before an escape sequence at its end
// that is not complete. Print SanitizeTerminalOutput(done), and put rest in
// front of the next chunk. Then a sequence that two chunks divide is removed
// or kept as one sequence, and its bytes do not show as text.
func SplitOpenEscape(s string) (done, rest string) {
	if loc := openEscape.FindStringIndex(s); loc != nil {
		return s[:loc[0]], s[loc[0]:]
	}
	return s, ""
}

// SanitizeTerminalOutput removes escape sequences and control characters from
// text before Fabric writes it to the terminal. Model output can contain
// sequences that move the cursor, clear the screen, set links (OSC 8) or write
// to the clipboard (OSC 52). The function keeps SGR color sequences, tab,
// newline and carriage return. Only the displayed copy goes through this
// function. The session, clipboard and output file keep the full text.
func SanitizeTerminalOutput(s string) string {
	return terminalControl.ReplaceAllStringFunc(s, func(m string) string {
		if sgrSequence.MatchString(m) {
			return m
		}
		return ""
	})
}
