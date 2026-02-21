// Package template provides file system operations for the template system.
// Security Note: This plugin provides access to the local filesystem.
// Consider carefully which paths to allow access to in production.
package template

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// MaxFileSize defines the maximum file size that can be read (1MB)
const MaxFileSize = 1 * 1024 * 1024

// FilePlugin provides filesystem operations with safety constraints:
// - No directory traversal
// - Size limits
// - Path sanitization
type FilePlugin struct{}

// safePath validates and normalizes file paths
func (p *FilePlugin) safePath(path string) (string, error) {
	debugf(i18n.T("template_file_log_validating_path"), path)

	// Basic security check - no path traversal
	if strings.Contains(path, "..") {
		return "", errors.New(i18n.T("template_file_error_path_contains_parent_ref"))
	}

	// Expand home directory if needed
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_expand_home_dir"), err)
		}
		path = filepath.Join(home, path[2:])
	}

	// Clean the path
	cleaned := filepath.Clean(path)
	debugf(i18n.T("template_file_log_cleaned_path"), cleaned)
	return cleaned, nil
}

// Apply executes file operations:
//   - read:PATH - Read entire file content
//   - tail:PATH|N - Read last N lines
//   - exists:PATH - Check if file exists
//   - size:PATH - Get file size in bytes
//   - modified:PATH - Get last modified time
func (p *FilePlugin) Apply(operation string, value string) (string, error) {
	debugf(i18n.T("template_file_log_operation_value"), operation, value)

	switch operation {
	case "tail":
		parts := strings.Split(value, "|")
		if len(parts) != 2 {
			return "", errors.New(i18n.T("template_file_error_tail_requires_path_lines"))
		}

		path, err := p.safePath(parts[0])
		if err != nil {
			return "", err
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_invalid_line_count"), parts[1])
		}

		if n < 1 {
			return "", errors.New(i18n.T("template_file_error_line_count_positive"))
		}

		lines, err := p.lastNLines(path, n)
		if err != nil {
			return "", err
		}

		result := strings.Join(lines, "\n")
		debugf(i18n.T("template_file_log_tail_returning_lines"), len(lines))
		return result, nil

	case "read":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		if info.Size() > MaxFileSize {
			return "", fmt.Errorf(i18n.T("template_file_error_size_exceeds_limit"),
				info.Size(), MaxFileSize)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_read_file"), err)
		}

		debugf(i18n.T("template_file_log_read_bytes"), len(content))
		return string(content), nil

	case "exists":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		_, err = os.Stat(path)
		exists := err == nil
		debugf(i18n.T("template_file_log_exists_for_path"), exists, path)
		return fmt.Sprintf("%t", exists), nil

	case "size":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		size := info.Size()
		debugf(i18n.T("template_file_log_size_for_path"), size, path)
		return fmt.Sprintf("%d", size), nil

	case "modified":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		mtime := info.ModTime().Format(time.RFC3339)
		debugf(i18n.T("template_file_log_modified_for_path"), mtime, path)
		return mtime, nil

	default:
		return "", fmt.Errorf(i18n.T("template_file_error_unknown_operation"),
			operation)
	}
}

// lastNLines returns the last n lines from a file
func (p *FilePlugin) lastNLines(path string, n int) ([]string, error) {
	debugf(i18n.T("template_file_log_reading_last_lines"), n, path)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("template_file_error_open_file"), err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf(i18n.T("template_file_error_stat_open_file"), err)
	}

	if info.Size() > MaxFileSize {
		return nil, fmt.Errorf(i18n.T("template_file_error_size_exceeds_limit"),
			info.Size(), MaxFileSize)
	}

	lines := make([]string, 0, n)
	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if len(lines) == n {
			lines = lines[1:]
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(i18n.T("template_file_error_scanner_read"), err)
	}

	debugf(i18n.T("template_file_log_read_total_return_last"), lineCount, len(lines))
	return lines, nil
}
