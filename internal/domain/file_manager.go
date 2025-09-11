package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// ParseFileChanges extracts and parses file changes from LLM output
// Supports both JSON format (with __CREATE_CODING_FEATURE_FILE_CHANGES__ marker)
// and markdown format (with "## File Changes" section)
func ParseFileChanges(output string) (changeSummary string, changes []FileChange, err error) {
	// First try the JSON format with marker
	fileChangesStart := strings.Index(output, FileChangesMarker)
	if fileChangesStart != -1 {
		return parseJSONFileChanges(output, fileChangesStart)
	}

	// Then try the markdown format
	return parseMarkdownFileChanges(output)
}

// parseJSONFileChanges handles the JSON format with __CREATE_CODING_FEATURE_FILE_CHANGES__ marker
func parseJSONFileChanges(output string, fileChangesStart int) (changeSummary string, changes []FileChange, err error) {
	changeSummary = output[:fileChangesStart] // Everything before the marker

	// Extract the JSON part
	jsonStart := fileChangesStart + len(FileChangesMarker)
	// Find the first [ after the file changes marker
	jsonArrayStart := strings.Index(output[jsonStart:], "[")
	if jsonArrayStart == -1 {
		return output, nil, fmt.Errorf("invalid %s format: no JSON array found", FileChangesMarker)
	}
	jsonStart += jsonArrayStart

	// Find the matching closing bracket for the array with proper bracket counting
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
		return output, nil, fmt.Errorf("invalid %s format: unbalanced brackets", FileChangesMarker)
	}

	// Extract the JSON string and fix escape sequences
	jsonStr := output[jsonStart:jsonEnd]

	// Fix specific invalid escape sequences
	// First try with the common \C issue
	jsonStr = strings.Replace(jsonStr, `\C`, `\\C`, -1)

	// Parse the JSON
	var fileChanges []FileChange
	err = json.Unmarshal([]byte(jsonStr), &fileChanges)
	if err != nil {
		// If still failing, try a more comprehensive fix
		jsonStr = fixInvalidEscapes(jsonStr)
		err = json.Unmarshal([]byte(jsonStr), &fileChanges)
		if err != nil {
			return changeSummary, nil, fmt.Errorf("failed to parse %s JSON: %w", FileChangesMarker, err)
		}
	}

	// Validate file changes
	if err = validateFileChanges(fileChanges); err != nil {
		return changeSummary, nil, err
	}

	return changeSummary, fileChanges, nil
}

// parseMarkdownFileChanges handles markdown format with "## File Changes" section
func parseMarkdownFileChanges(output string) (changeSummary string, changes []FileChange, err error) {
	// Look for "## File Changes" section
	fileChangesStart := strings.Index(output, "## File Changes")
	if fileChangesStart == -1 {
		// No file changes section found
		return output, nil, nil
	}

	// Extract the summary (everything before the file changes section)
	changeSummary = strings.TrimSpace(output[:fileChangesStart])

	// Find the end of the file changes section (next ## header or end of string)
	fileChangesContent := output[fileChangesStart:]
	nextSectionStart := strings.Index(fileChangesContent[1:], "\n## ")
	var fileChangesSection string
	if nextSectionStart == -1 {
		fileChangesSection = fileChangesContent
	} else {
		fileChangesSection = fileChangesContent[:nextSectionStart+1]
	}

	// Parse individual file operations from the markdown format
	lines := strings.Split(fileChangesSection, "\n")
	var currentOperation string
	var currentPath string
	var contentLines []string
	inCodeBlock := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Check for operation headers like "### CREATE: filename" or "### UPDATE: filename"
		if strings.HasPrefix(line, "### CREATE: ") || strings.HasPrefix(line, "### UPDATE: ") {
			// Save previous file change if we have one
			if currentPath != "" && currentOperation != "" {
				content := strings.Join(contentLines, "\n")
				changes = append(changes, FileChange{
					Operation: strings.ToLower(currentOperation),
					Path:      currentPath,
					Content:   content,
				})
			}

			// Parse new operation
			if strings.HasPrefix(line, "### CREATE: ") {
				currentOperation = "create"
				currentPath = strings.TrimSpace(line[12:])
			} else {
				currentOperation = "update"
				currentPath = strings.TrimSpace(line[12:])
			}
			contentLines = []string{}
			inCodeBlock = false
			continue
		}

		// Handle code block markers
		if line == "```" {
			inCodeBlock = !inCodeBlock
			continue
		}

		// Collect content lines inside code blocks
		if inCodeBlock {
			contentLines = append(contentLines, line)
		}
	}

	// Save the last file change
	if currentPath != "" && currentOperation != "" {
		content := strings.Join(contentLines, "\n")
		changes = append(changes, FileChange{
			Operation: strings.ToLower(currentOperation),
			Path:      currentPath,
			Content:   content,
		})
	}

	// Validate file changes
	if err = validateFileChanges(changes); err != nil {
		return changeSummary, nil, err
	}

	return changeSummary, changes, nil
}

// validateFileChanges validates parsed file changes
func validateFileChanges(fileChanges []FileChange) error {
	for i, change := range fileChanges {
		// Validate operation
		if change.Operation != "create" && change.Operation != "update" {
			return fmt.Errorf("invalid operation for file change %d: %s", i, change.Operation)
		}

		// Validate path
		if change.Path == "" {
			return fmt.Errorf("empty path for file change %d", i)
		}

		// Check for suspicious paths (directory traversal)
		if strings.Contains(change.Path, "..") {
			return fmt.Errorf("suspicious path for file change %d: %s", i, change.Path)
		}

		// Check file size
		if len(change.Content) > MaxFileSize {
			return fmt.Errorf("file content too large for file change %d: %d bytes", i, len(change.Content))
		}
	}
	return nil
}

// fixInvalidEscapes replaces invalid escape sequences in JSON strings
func fixInvalidEscapes(jsonStr string) string {
	validEscapes := []byte{'b', 'f', 'n', 'r', 't', '\\', '/', '"', 'u'}

	var result strings.Builder
	inQuotes := false
	i := 0

	for i < len(jsonStr) {
		ch := jsonStr[i]

		// Track whether we're inside a JSON string
		if ch == '"' && (i == 0 || jsonStr[i-1] != '\\') {
			inQuotes = !inQuotes
		}

		// Handle actual control characters inside string literals
		if inQuotes {
			// Convert literal control characters to proper JSON escape sequences
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
				// Handle other control characters
				fmt.Fprintf(&result, "\\u%04x", ch)
				i++
				continue
			}
		}

		// Check for escape sequences only inside strings
		if inQuotes && ch == '\\' && i+1 < len(jsonStr) {
			nextChar := jsonStr[i+1]
			isValid := false

			for _, validEscape := range validEscapes {
				if nextChar == validEscape {
					isValid = true
					break
				}
			}

			if !isValid {
				// Invalid escape sequence - add an extra backslash
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
		// Get the absolute path
		absPath := filepath.Join(projectRoot, change.Path)

		// Create directories if necessary
		dir := filepath.Dir(absPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s for file change %d: %w", dir, i, err)
		}

		// Write the file
		if err := os.WriteFile(absPath, []byte(change.Content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s for file change %d: %w", absPath, i, err)
		}

		fmt.Printf("Applied %s operation to %s\n", change.Operation, change.Path)
	}

	return nil
}
