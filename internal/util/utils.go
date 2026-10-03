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

	// Return Windows UNC paths unchanged.
	if runtime.GOOS == "windows" && strings.HasPrefix(path, `\\`) {
		return path, nil
	}

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", errors.New(i18n.T("util_error_resolve_home_directory"))
		}
		path = filepath.Join(home, path[1:])
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", errors.New(i18n.T("util_error_get_absolute_path"))
	}

	// Resolve symlinks. A path that does not exist yet is not an error.
	resolvedPath, err := filepath.EvalSymlinks(absPath)
	if err == nil {
		return resolvedPath, nil
	}
	if os.IsNotExist(err) {
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

	return false
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
			return "", nil
		}
		return "", fmt.Errorf(i18n.T("util_error_accessing_config_path"), err)
	}
	return defaultConfigPath, nil
}
