package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHtmlReadability(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "Empty HTML",
			html:     "",
			expected: "",
		},
		{
			name:     "HTML with text",
			html:     "<p>Hello World</p>",
			expected: "Hello World",
		},
		{
			name:     "HTML with nested tags",
			html:     "<div><p>Hello</p><p>World</p></div>",
			expected: "HelloWorld",
		},
		{
			name:     "HTML missing tags",
			html:     "<div><p>Hello</p><p>World</div>",
			expected: "HelloWorld",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := HtmlReadability(tc.html)

			// 验证结果
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
