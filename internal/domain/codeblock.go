package domain

import "strings"

// ExtractFencedCodeBlock returns the content of the first (or, if last is true,
// the last) Markdown fenced code block in text, without the fence lines.
//
// A fence is a line of three or more backticks or tildes, optionally indented
// by up to three spaces and followed by an info string such as a language tag.
// The block is closed by a line of the same character that is at least as long
// as the opening fence. A block left open at the end of text runs to the end.
// The second return value is false when text contains no fenced code block.
func ExtractFencedCodeBlock(text string, last bool) (string, bool) {
	var (
		result string
		found  bool
		inside bool
		fence  string
		body   []string
	)

	for line := range strings.SplitSeq(text, "\n") {
		trimmed := strings.TrimRight(line, "\r")
		if !inside {
			if marker, ok := openingFence(trimmed); ok {
				inside, fence, body = true, marker, nil
			}
			continue
		}
		if isClosingFence(trimmed, fence) {
			inside = false
			result, found = strings.Join(body, "\n"), true
			if !last {
				return result, true
			}
			continue
		}
		body = append(body, trimmed)
	}

	if inside {
		result, found = strings.Join(body, "\n"), true
	}
	return result, found
}

// openingFence reports whether line opens a fenced code block and returns the
// fence marker (the run of backticks or tildes).
func openingFence(line string) (string, bool) {
	marker, rest, ok := fenceMarker(line)
	if !ok {
		return "", false
	}
	// A backtick fence's info string may not contain backticks, so that
	// inline code such as ```foo``` is not mistaken for a fence.
	if marker[0] == '`' && strings.Contains(rest, "`") {
		return "", false
	}
	return marker, true
}

// isClosingFence reports whether line closes a block opened with fence.
func isClosingFence(line, fence string) bool {
	marker, rest, ok := fenceMarker(line)
	return ok && marker[0] == fence[0] && len(marker) >= len(fence) && strings.TrimSpace(rest) == ""
}

// fenceMarker splits a line indented by at most three spaces into its leading
// run of three or more backticks or tildes and the remainder of the line.
func fenceMarker(line string) (marker, rest string, ok bool) {
	content := strings.TrimLeft(line, " ")
	if len(line)-len(content) > 3 || content == "" || (content[0] != '`' && content[0] != '~') {
		return "", "", false
	}
	n := len(content) - len(strings.TrimLeft(content, content[:1]))
	if n < 3 {
		return "", "", false
	}
	return content[:n], content[n:], true
}
