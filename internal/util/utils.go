package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// GetAbsolutePath resolves a given path to its absolute form, handling ~, ./, ../, UNC paths, and symlinks.
func GetAbsolutePath(path string) (string, error) {
	if path == "" {
		return "", errors.New(i18n.T("util_error_path_is_empty"))
	}

	// Handle UNC paths on Windows
	if runtime.GOOS == "windows" && strings.HasPrefix(path, `\\`) {
		return path, nil
	}

	// Handle ~ for home directory expansion
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", errors.New(i18n.T("util_error_resolve_home_directory"))
		}
		path = filepath.Join(home, path[1:])
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", errors.New(i18n.T("util_error_get_absolute_path"))
	}

	// Resolve symlinks, but allow non-existent paths
	resolvedPath, err := filepath.EvalSymlinks(absPath)
	if err == nil {
		return resolvedPath, nil
	}
	if os.IsNotExist(err) {
		// Return the absolute path for non-existent paths
		return absPath, nil
	}

	return "", fmt.Errorf(i18n.T("util_error_resolve_symlinks"), err)
}

// Helper function to check if a symlink points to a directory
func IsSymlinkToDir(path string) bool {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		resolvedPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return false
		}

		fileInfo, err = os.Stat(resolvedPath)
		if err != nil {
			return false
		}

		return fileInfo.IsDir()
	}

	return false // Regular directories should not be treated as symlinks
}

// GetDefaultConfigPath returns the default path for the configuration file
// if it exists, otherwise returns an empty string.
func GetDefaultConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf(i18n.T("util_error_determine_home_directory"), err)
	}

	defaultConfigPath := filepath.Join(homeDir, ".config", "fabric", "config.yaml")
	if _, err := os.Stat(defaultConfigPath); err != nil {
		if os.IsNotExist(err) {
			return "", nil // Return no error for non-existent config path
		}
		return "", fmt.Errorf(i18n.T("util_error_accessing_config_path"), err)
	}
	return defaultConfigPath, nil
}
