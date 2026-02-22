// Package template provides system information operations for the template system.
package template

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// SysPlugin provides access to system-level information.
// Security Note: This plugin provides access to system information and
// environment variables. Be cautious with exposed variables in templates.
type SysPlugin struct{}

// Apply executes system operations with the following options:
//   - hostname: System hostname
//   - user: Current username
//   - os: Operating system (linux, darwin, windows)
//   - arch: System architecture (amd64, arm64, etc)
//   - env:VALUE: Environment variable lookup
//   - pwd: Current working directory
//   - home: User's home directory
func (p *SysPlugin) Apply(operation string, value string) (string, error) {
	debugf("Sys: operation=%q value=%q", operation, value)

	switch operation {
	case "hostname":
		hostname, err := os.Hostname()
		if err != nil {
			debugf("Sys: hostname error: %v", err)
			return "", fmt.Errorf(i18n.T("template_sys_error_hostname"), err)
		}
		debugf("Sys: hostname=%q", hostname)
		return hostname, nil

	case "user":
		currentUser, err := user.Current()
		if err != nil {
			debugf("Sys: user error: %v", err)
			return "", fmt.Errorf(i18n.T("template_sys_error_user"), err)
		}
		debugf("Sys: user=%q", currentUser.Username)
		return currentUser.Username, nil

	case "os":
		result := runtime.GOOS
		debugf("Sys: os=%q", result)
		return result, nil

	case "arch":
		result := runtime.GOARCH
		debugf("Sys: arch=%q", result)
		return result, nil

	case "env":
		if value == "" {
			debugf("Sys: env error: missing variable name")
			return "", errors.New(i18n.T("template_sys_error_env_requires_var"))
		}
		result := os.Getenv(value)
		debugf("Sys: env %q=%q", value, result)
		return result, nil

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			debugf("Sys: pwd error: %v", err)
			return "", fmt.Errorf(i18n.T("template_sys_error_pwd"), err)
		}
		debugf("Sys: pwd=%q", dir)
		return dir, nil

	case "home":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			debugf("Sys: home error: %v", err)
			return "", fmt.Errorf(i18n.T("template_sys_error_home"), err)
		}
		debugf("Sys: home=%q", homeDir)
		return homeDir, nil

	default:
		debugf("Sys: unknown operation %q", operation)
		return "", fmt.Errorf(i18n.T("template_sys_error_unknown_operation"), operation)
	}
}
