// utils.go in template package for now
package template

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ExpandPath expands the ~ to user's home directory and returns absolute path
// It also checks if the path exists
// Returns expanded absolute path or error if:
// - cannot determine user home directory
// - cannot convert to absolute path
// - path doesn't exist
func ExpandPath(path string) (string, error) {
	// If path starts with ~
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		// Replace ~/ with actual home directory
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Check if path exists
	if _, err := os.Stat(absPath); err != nil {
		return "", fmt.Errorf("path does not exist: %w", err)
	}

	return absPath, nil
}
