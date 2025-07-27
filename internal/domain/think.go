package domain

import (
	"regexp"
	"sync"
)

// StripThinkBlocks removes any content between the provided start and end tags
// from the input string. Whitespace following the end tag is also removed so
// output resumes at the next non-empty line.
var (
	regexCache = make(map[string]*regexp.Regexp)
	cacheMutex sync.Mutex
)

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
