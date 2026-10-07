package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

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

		// A control character in a path makes a bad file name, and it can
		// change the terminal when Fabric shows the path.
		if !filepath.IsLocal(change.Path) || strings.ContainsFunc(change.Path, unicode.IsControl) {
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

// ApplyFileChanges applies the parsed file changes to the file system.
//
// The change list comes from model output. Thus all writes go through os.Root,
// which refuses a path that goes out of projectRoot through a symlink.
// filepath.IsLocal only examines the path text and cannot find this.
// ApplyFileChanges also refuses a path through a symlink and a path into .git,
// so that a change cannot write a git hook.
func ApplyFileChanges(projectRoot string, changes []FileChange) error {
	root, err := os.OpenRoot(projectRoot)
	if err != nil {
		return fmt.Errorf(i18n.T("file_manager_failed_open_root"), projectRoot, err)
	}
	defer root.Close()

	// Examine all paths before the first write, so that a bad path does not
	// leave a part of the change set on disk.
	for i, change := range changes {
		if !filepath.IsLocal(change.Path) || unsafePath(root, filepath.Clean(change.Path)) {
			return fmt.Errorf(i18n.T("file_manager_suspicious_path"), i, change.Path)
		}
	}

	for i, change := range changes {
		clean := filepath.Clean(change.Path)
		if dir := filepath.Dir(clean); dir != "." {
			if err := root.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf(i18n.T("file_manager_failed_create_directory"), dir, i, err)
			}
		}

		if err := root.WriteFile(clean, []byte(change.Content), 0644); err != nil {
			return fmt.Errorf(i18n.T("file_manager_failed_write_file"), clean, i, err)
		}

		fmt.Printf(i18n.T("file_manager_applied_operation")+"\n", change.Operation, change.Path)
	}

	return nil
}

// unsafePath reports whether a component of the clean, relative path is .git
// (any case) or a symlink that is in the root.
func unsafePath(root *os.Root, clean string) bool {
	parts := strings.Split(clean, string(filepath.Separator))
	for i, part := range parts {
		if strings.EqualFold(part, ".git") {
			return true
		}
		if info, err := root.Lstat(filepath.Join(parts[:i+1]...)); err == nil && info.Mode()&os.ModeSymlink != 0 {
			return true
		}
	}
	return false
}
