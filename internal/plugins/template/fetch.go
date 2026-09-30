// Package template provides URL fetching operations for the template system.
// Security Note: This plugin makes outbound HTTP requests. Use with caution
// and consider implementing URL allowlists in production.
package template

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/danielmiessler/fabric/internal/i18n"
)

const (
	// MaxContentSize limits response size to 1MB to prevent memory issues
	MaxContentSize = 1024 * 1024

	// UserAgent identifies the client in HTTP requests
	UserAgent = "Fabric-Fetch/1.0"
)

// FetchPlugin provides HTTP fetching capabilities with safety constraints:
// - Only text content types allowed
// - Size limited to MaxContentSize
// - UTF-8 validation
// - Null byte checking
type FetchPlugin struct{}

// Apply executes fetch operations:
//   - get:URL: Fetches content from URL, returns text content
func (p *FetchPlugin) Apply(operation string, value string) (string, error) {
	debugf("Fetch: operation=%q value=%q", operation, value)

	switch operation {
	case "get":
		return p.fetch(value)
	default:
		return "", fmt.Errorf(i18n.T("fetch_unknown_operation"), operation)
	}
}

// isTextContent checks if the content type is text-based
func (p *FetchPlugin) isTextContent(contentType string) bool {
	debugf("Fetch: checking content type %q", contentType)

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		debugf("Fetch: error parsing media type: %v", err)
		return false
	}

	isText := strings.HasPrefix(mediaType, "text/") ||
		mediaType == "application/json" ||
		mediaType == "application/xml" ||
		mediaType == "application/yaml" ||
		mediaType == "application/x-yaml" ||
		strings.HasSuffix(mediaType, "+json") ||
		strings.HasSuffix(mediaType, "+xml") ||
		strings.HasSuffix(mediaType, "+yaml")

	debugf("Fetch: content type %q is text: %v", mediaType, isText)
	return isText
}

// validateTextContent returns an error if content is not valid UTF-8 or contains a null byte.
func (p *FetchPlugin) validateTextContent(content []byte) error {
	debugf("Fetch: validating content length=%d bytes", len(content))

	if !utf8.Valid(content) {
		return errors.New(i18n.T("fetch_content_not_utf8"))
	}

	if bytes.Contains(content, []byte{0}) {
		return errors.New(i18n.T("fetch_content_null_bytes"))
	}

	debugf("Fetch: content validation successful")
	return nil
}

// fetch downloads urlStr and applies the size, content-type, and UTF-8 checks.
func (p *FetchPlugin) fetch(urlStr string) (string, error) {
	debugf("Fetch: requesting URL %q", urlStr)

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", fmt.Errorf(i18n.T("fetch_error_create_request"), err)
	}
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf(i18n.T("fetch_error_fetching_url"), err)
	}
	defer resp.Body.Close()

	debugf("Fetch: got response status=%q", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(i18n.T("fetch_http_error"), resp.StatusCode, resp.Status)
	}

	if contentLength := resp.ContentLength; contentLength > MaxContentSize {
		return "", fmt.Errorf(i18n.T("fetch_content_too_large"), contentLength, MaxContentSize)
	}

	contentType := resp.Header.Get("Content-Type")
	debugf("Fetch: content-type=%q", contentType)
	if !p.isTextContent(contentType) {
		return "", fmt.Errorf(i18n.T("fetch_unsupported_content_type"), contentType)
	}

	debugf("Fetch: reading response body")
	limitReader := io.LimitReader(resp.Body, MaxContentSize+1)
	content, err := io.ReadAll(limitReader)
	if err != nil {
		return "", fmt.Errorf(i18n.T("fetch_error_reading_response"), err)
	}

	if len(content) > MaxContentSize {
		return "", fmt.Errorf(i18n.T("fetch_content_exceeds_limit"), MaxContentSize)
	}

	if err := p.validateTextContent(content); err != nil {
		return "", err
	}

	debugf("Fetch: operation completed successfully, read %d bytes", len(content))
	return string(content), nil
}
