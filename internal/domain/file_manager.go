package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// FileChangesMarker identifies the start of a file changes section in output
const FileChangesMarker = "__CREATE_CODING_FEATURE_FILE_CHANGES__"

const (
	// MaxFileSize is the maximum size of a file that can be created (10MB)
	MaxFileSize = 10 * 1024 * 1024
)

// FileChange represents a single file change operation to be performed
type FileChange struct {
	Operation string `json:"operation"` // "create" or "update"
	Path      string `json:"path"`      // Relative path from project root
	Content   string `json:"content"`   // New file content
}

// ParseFileChanges extracts and parses the file change marker section from LLM output
func ParseFileChanges(output string) (changeSummary string, changes []FileChange, err error) {
	fileChangesStart := strings.Index(output, FileChangesMarker)
	if fileChangesStart == -1 {
		return output, nil, nil
	}
	changeSummary = output[:fileChangesStart]

	jsonStart := fileChangesStart + len(FileChangesMarker)
	jsonArrayStart := strings.Index(output[jsonStart:], "[")
	if jsonArrayStart == -1 {
		return output, nil, fmt.Errorf(i18n.T("file_manager_invalid_format_no_json_array"), FileChangesMarker)
	}
	jsonStart += jsonArrayStart

	bracketCount := 0
	jsonEnd := jsonStart
	for i := jsonStart; i < len(output); i++ {
		if output[i] == '[' {
			bracketCount++
		} else if output[i] == ']' {
			bracketCount--
			if bracketCount == 0 {
				jsonEnd = i + 1
				break
			}
		}
	}

	if bracketCount != 0 {
		return output, nil, fmt.Errorf(i18n.T("file_manager_invalid_format_unbalanced_brackets"), FileChangesMarker)
	}

	jsonStr := output[jsonStart:jsonEnd]

	jsonStr = strings.Replace(jsonStr, `\C`, `\\C`, -1)

	var fileChanges []FileChange
	err = json.Unmarshal([]byte(jsonStr), &fileChanges)
	if err != nil {
		jsonStr = fixInvalidEscapes(jsonStr)
		err = json.Unmarshal([]byte(jsonStr), &fileChanges)
		if err != nil {
			return changeSummary, nil, fmt.Errorf(i18n.T("file_manager_failed_parse_json"), FileChangesMarker, err)
		}
	}

	for i, change := range fileChanges {
		if change.Operation != "create" && change.Operation != "update" {
			return changeSummary, nil, fmt.Errorf(i18n.T("file_manager_invalid_operation"), i, change.Operation)
		}

		if change.Path == "" {
			return changeSummary, nil, fmt.Errorf(i18n.T("file_manager_empty_path"), i)
		}

		if !filepath.IsLocal(change.Path) {
			return changeSummary, nil, fmt.Errorf(i18n.T("file_manager_suspicious_path"), i, change.Path)
		}

		if len(change.Content) > MaxFileSize {
			return changeSummary, nil, fmt.Errorf(i18n.T("file_manager_file_content_too_large"), i, len(change.Content))
		}
	}

	return changeSummary, fileChanges, nil
}

// fixInvalidEscapes escapes raw control characters inside JSON strings and
// doubles the backslash of an unknown escape sequence.
func fixInvalidEscapes(jsonStr string) string {
	validEscapes := []byte{'b', 'f', 'n', 'r', 't', '\\', '/', '"', 'u'}

	var result strings.Builder
	inQuotes := false
	i := 0

	for i < len(jsonStr) {
		ch := jsonStr[i]

		if ch == '"' && (i == 0 || jsonStr[i-1] != '\\') {
			inQuotes = !inQuotes
		}

		if inQuotes {
			// JSON does not accept a raw control character in a string.
			if ch == '\n' {
				result.WriteString("\\n")
				i++
				continue
			} else if ch == '\r' {
				result.WriteString("\\r")
				i++
				continue
			} else if ch == '\t' {
				result.WriteString("\\t")
				i++
				continue
			} else if ch < 32 {
				fmt.Fprintf(&result, "\\u%04x", ch)
				i++
				continue
			}
		}

		if inQuotes && ch == '\\' && i+1 < len(jsonStr) {
			nextChar := jsonStr[i+1]
			isValid := slices.Contains(validEscapes, nextChar)

			if !isValid {
				// Double the backslash so the sequence decodes as a literal backslash.
				result.WriteByte('\\')
				result.WriteByte('\\')
				i++
				continue
			}
		}

		result.WriteByte(ch)
		i++
	}

	return result.String()
}

// ApplyFileChanges applies the parsed file changes to the file system
func ApplyFileChanges(projectRoot string, changes []FileChange) error {
	for i, change := range changes {
		if !filepath.IsLocal(change.Path) {
			return fmt.Errorf(i18n.T("file_manager_suspicious_path"), i, change.Path)
		}

		absPath := filepath.Join(projectRoot, change.Path)

		dir := filepath.Dir(absPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf(i18n.T("file_manager_failed_create_directory"), dir, i, err)
		}

		if err := os.WriteFile(absPath, []byte(change.Content), 0644); err != nil {
			return fmt.Errorf(i18n.T("file_manager_failed_write_file"), absPath, i, err)
		}

		fmt.Printf(i18n.T("file_manager_applied_operation")+"\n", change.Operation, change.Path)
	}

	return nil
}
