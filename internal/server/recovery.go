package restapi

import (
	"io"
	"strings"
)

// secretHeaderPrefixes are the request header lines that the recovery dump
// must not print. gin masks Authorization itself but keeps the others, so
// X-API-Key would be written to the log in clear. The names are in the
// canonical form that net/http uses in a request dump.
var secretHeaderPrefixes = []string{
	"X-Api-Key:",
	"Cookie:",
	"Proxy-Authorization:",
}

// redactRequestDump masks the value of each secret header line in text.
func redactRequestDump(text string) string {
	lines := strings.Split(text, "\r\n")
	for i, line := range lines {
		for _, prefix := range secretHeaderPrefixes {
			if strings.HasPrefix(line, prefix) {
				lines[i] = prefix + " *"
				break
			}
		}
	}
	return strings.Join(lines, "\r\n")
}

// RedactingWriter masks secret request headers in what gin.Recovery writes.
// Recovery prints the whole request dump in one Write call, for both the
// debug-mode panic dump and the broken-pipe dump.
type RedactingWriter struct {
	out io.Writer
}

func (w RedactingWriter) Write(p []byte) (int, error) {
	if _, err := io.WriteString(w.out, redactRequestDump(string(p))); err != nil {
		return 0, err
	}
	return len(p), nil
}
