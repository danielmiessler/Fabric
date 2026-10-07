// Package template provides file system operations for the template system.
// Security Note: This plugin provides access to the local filesystem.
// Consider carefully which paths to allow access to in production.
package template

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// MaxFileSize defines the maximum file size that can be read (1MB)
const MaxFileSize = 1 * 1024 * 1024

// fileReadRoot is the folder that the file plugin can read. The default is
// the working folder of the process. Tests change it.
var fileReadRoot = "."

// FilePlugin provides filesystem operations with safety constraints:
// - Each path must be a relative path in fileReadRoot
// - A symbolic link must stay in fileReadRoot
// - Size limits
type FilePlugin struct{}

// safePath rejects a path that contains "..", starts with "~" or is
// absolute. It gives the cleaned path, which is relative to fileReadRoot.
func (p *FilePlugin) safePath(path string) (string, error) {
	debugf(i18n.T("template_file_log_validating_path"), path)

	if strings.Contains(path, "..") {
		return "", errors.New(i18n.T("template_file_error_path_contains_parent_ref"))
	}
	if strings.HasPrefix(path, "~") || filepath.IsAbs(path) {
		return "", errors.New(i18n.T("template_file_error_path_not_confined"))
	}

	cleaned := filepath.Clean(path)
	if cleaned == "." {
		return "", errors.New(i18n.T("template_file_error_path_not_confined"))
	}
	debugf(i18n.T("template_file_log_cleaned_path"), cleaned)
	return cleaned, nil
}

// inRoot checks path with safePath and opens fileReadRoot. All operations
// get to the file through this root. os.Root rejects a path or a symbolic
// link that goes out of the root, and it rejects an absolute symbolic link.
// The caller must close the root.
func (p *FilePlugin) inRoot(path string) (*os.Root, string, error) {
	rel, err := p.safePath(path)
	if err != nil {
		return nil, "", err
	}
	root, err := os.OpenRoot(fileReadRoot)
	if err != nil {
		return nil, "", fmt.Errorf(i18n.T("template_file_error_open_file"), err)
	}
	return root, rel, nil
}

// statInRoot gives the file information for path through inRoot.
func (p *FilePlugin) statInRoot(path string) (os.FileInfo, error) {
	root, rel, err := p.inRoot(path)
	if err != nil {
		return nil, err
	}
	defer root.Close()
	return root.Stat(rel)
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

		root, rel, err := p.inRoot(parts[0])
		if err != nil {
			return "", err
		}
		defer root.Close()

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_invalid_line_count"), parts[1])
		}

		if n < 1 {
			return "", errors.New(i18n.T("template_file_error_line_count_positive"))
		}

		lines, err := p.lastNLines(root, rel, n)
		if err != nil {
			return "", err
		}

		result := strings.Join(lines, "\n")
		debugf(i18n.T("template_file_log_tail_returning_lines"), len(lines))
		return result, nil

	case "read":
		root, rel, err := p.inRoot(value)
		if err != nil {
			return "", err
		}
		defer root.Close()

		info, err := root.Stat(rel)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		if info.Size() > MaxFileSize {
			return "", fmt.Errorf(i18n.T("template_file_error_size_exceeds_limit"),
				info.Size(), MaxFileSize)
		}

		f, err := root.Open(rel)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_read_file"), err)
		}
		defer f.Close()

		// Stat does not give the correct size for all files. For example,
		// /dev/zero has the size 0. Thus read one byte more than the limit,
		// and reject the file if that byte is there.
		content, err := io.ReadAll(io.LimitReader(f, MaxFileSize+1))
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_read_file"), err)
		}

		if len(content) > MaxFileSize {
			return "", fmt.Errorf(i18n.T("template_file_error_size_exceeds_limit"),
				len(content), MaxFileSize)
		}

		debugf(i18n.T("template_file_log_read_bytes"), len(content))
		return string(content), nil

	case "exists":
		root, rel, err := p.inRoot(value)
		if err != nil {
			return "", err
		}
		defer root.Close()

		_, err = root.Stat(rel)
		exists := err == nil
		debugf(i18n.T("template_file_log_exists_for_path"), exists, value)
		return fmt.Sprintf("%t", exists), nil

	case "size":
		info, err := p.statInRoot(value)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		size := info.Size()
		debugf(i18n.T("template_file_log_size_for_path"), size, value)
		return fmt.Sprintf("%d", size), nil

	case "modified":
		info, err := p.statInRoot(value)
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_file_error_stat_file"), err)
		}

		mtime := info.ModTime().Format(time.RFC3339)
		debugf(i18n.T("template_file_log_modified_for_path"), mtime, value)
		return mtime, nil

	default:
		return "", fmt.Errorf(i18n.T("template_file_error_unknown_operation"),
			operation)
	}
}

// lastNLines returns the last n lines from the file rel in root.
func (p *FilePlugin) lastNLines(root *os.Root, rel string, n int) ([]string, error) {
	debugf(i18n.T("template_file_log_reading_last_lines"), n, rel)

	file, err := root.Open(rel)
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

	// n comes from the template. Use it only as a capacity hint up to a
	// small cap, because a large n must not cause a large allocation.
	lines := make([]string, 0, min(n, 4096))
	// Stat does not give the correct size for all files (see read). Thus
	// read one byte more than the limit, and reject the file if that byte
	// is there.
	limited := &io.LimitedReader{R: file, N: MaxFileSize + 1}
	scanner := bufio.NewScanner(limited)

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
	if limited.N == 0 {
		return nil, fmt.Errorf(i18n.T("template_file_error_size_exceeds_limit"),
			MaxFileSize+1, MaxFileSize)
	}

	debugf(i18n.T("template_file_log_read_total_return_last"), lineCount, len(lines))
	return lines, nil
}
