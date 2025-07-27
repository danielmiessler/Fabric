// Package template provides file system operations for the template system.
// Security Note: This plugin provides access to the local filesystem.
// Consider carefully which paths to allow access to in production.
package template

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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
	debugf("File: validating path %q", path)

	// Basic security check - no path traversal
	if strings.Contains(path, "..") {
		return "", fmt.Errorf("file: path cannot contain '..'")
	}

	// Expand home directory if needed
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("file: could not expand home directory: %v", err)
		}
		path = filepath.Join(home, path[2:])
	}

	// Clean the path
	cleaned := filepath.Clean(path)
	debugf("File: cleaned path %q", cleaned)
	return cleaned, nil
}

// Apply executes file operations:
//   - read:PATH - Read entire file content
//   - tail:PATH|N - Read last N lines
//   - exists:PATH - Check if file exists
//   - size:PATH - Get file size in bytes
//   - modified:PATH - Get last modified time
func (p *FilePlugin) Apply(operation string, value string) (string, error) {
	debugf("File: operation=%q value=%q", operation, value)

	switch operation {
	case "tail":
		parts := strings.Split(value, "|")
		if len(parts) != 2 {
			return "", fmt.Errorf("file: tail requires format path|lines")
		}

		path, err := p.safePath(parts[0])
		if err != nil {
			return "", err
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("file: invalid line count %q", parts[1])
		}

		if n < 1 {
			return "", fmt.Errorf("file: line count must be positive")
		}

		lines, err := p.lastNLines(path, n)
		if err != nil {
			return "", err
		}

		result := strings.Join(lines, "\n")
		debugf("File: tail returning %d lines", len(lines))
		return result, nil

	case "read":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf("file: could not stat file: %v", err)
		}

		if info.Size() > MaxFileSize {
			return "", fmt.Errorf("file: size %d exceeds limit of %d bytes",
				info.Size(), MaxFileSize)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("file: could not read: %v", err)
		}

		debugf("File: read %d bytes", len(content))
		return string(content), nil

	case "exists":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		_, err = os.Stat(path)
		exists := err == nil
		debugf("File: exists=%v for path %q", exists, path)
		return fmt.Sprintf("%t", exists), nil

	case "size":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf("file: could not stat file: %v", err)
		}

		size := info.Size()
		debugf("File: size=%d for path %q", size, path)
		return fmt.Sprintf("%d", size), nil

	case "modified":
		path, err := p.safePath(value)
		if err != nil {
			return "", err
		}

		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf("file: could not stat file: %v", err)
		}

		mtime := info.ModTime().Format(time.RFC3339)
		debugf("File: modified=%q for path %q", mtime, path)
		return mtime, nil

	default:
		return "", fmt.Errorf("file: unknown operation %q (supported: read, tail, exists, size, modified)",
			operation)
	}
}

// lastNLines returns the last n lines from a file
func (p *FilePlugin) lastNLines(path string, n int) ([]string, error) {
	debugf("File: reading last %d lines from %q", n, path)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("file: could not open: %v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("file: could not stat: %v", err)
	}

	if info.Size() > MaxFileSize {
		return nil, fmt.Errorf("file: size %d exceeds limit of %d bytes",
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
		return nil, fmt.Errorf("file: error reading: %v", err)
	}

	debugf("File: read %d lines total, returning last %d", lineCount, len(lines))
	return lines, nil
}
