package domain

import (
	"regexp"
	"sync"
)

var (
	regexCache = make(map[string]*regexp.Regexp)
	cacheMutex sync.Mutex
)

// StripThinkBlocks removes each start tag, end tag, and the text between them
// from input. It also removes the whitespace after the end tag, so the output
// continues at the next non-blank text.
func StripThinkBlocks(input, startTag, endTag string) string {
	if startTag == "" || endTag == "" {
		return input
	}

	cacheKey := startTag + "|" + endTag
	cacheMutex.Lock()
	re, exists := regexCache[cacheKey]
	if !exists {
		pattern := "(?s)" + regexp.QuoteMeta(startTag) + ".*?" + regexp.QuoteMeta(endTag) + "\\s*"
		re = regexp.MustCompile(pattern)
		regexCache[cacheKey] = re
	}
	cacheMutex.Unlock()

	return re.ReplaceAllString(input, "")
}
