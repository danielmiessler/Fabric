package template

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// ExpandPath expands the ~ to user's home directory and returns absolute path
// It also checks if the path exists
// Returns expanded absolute path or error if:
// - cannot determine user home directory
// - cannot convert to absolute path
// - path doesn't exist
func ExpandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf(i18n.T("template_utils_failed_get_home_dir"), err)
		}
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf(i18n.T("template_utils_failed_get_absolute_path"), err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return "", fmt.Errorf(i18n.T("template_utils_path_not_exist"), err)
	}

	return absPath, nil
}
