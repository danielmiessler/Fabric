package converter

import (
	"bytes"

	"github.com/go-shiori/go-readability"
)

// HtmlReadability Convert HTML input into a clean, readable view
// args：
//
//	html (string): full data of web page
//
// return：
//
//	viewContent (string): html main content
//	err (error): parser error
func HtmlReadability(html string) (ret string, err error) {
	buf := bytes.NewBufferString(html)
	var article readability.Article
	if article, err = readability.FromReader(buf, nil); err != nil {
		return
	}
	ret = article.TextContent
	return
}
