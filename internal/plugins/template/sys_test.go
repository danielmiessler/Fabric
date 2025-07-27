package template

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestSysPlugin(t *testing.T) {
	plugin := &SysPlugin{}

	// Set up test environment variable
	const testEnvVar = "FABRIC_TEST_VAR"
	const testEnvValue = "test_value"
	os.Setenv(testEnvVar, testEnvValue)
	defer os.Unsetenv(testEnvVar)

	tests := []struct {
		name      string
		operation string
		value     string
		validate  func(string) error
		wantErr   bool
	}{
		{
			name:      "hostname returns valid name",
			operation: "hostname",
			validate: func(got string) error {
				if got == "" {
					return fmt.Errorf("hostname is empty")
				}
				return nil
			},
		},
		{
			name:      "user returns current user",
			operation: "user",
			validate: func(got string) error {
				if got == "" {
					return fmt.Errorf("username is empty")
				}
				return nil
			},
		},
		{
			name:      "os returns valid OS",
			operation: "os",
			validate: func(got string) error {
				if got != runtime.GOOS {
					return fmt.Errorf("expected OS %s, got %s", runtime.GOOS, got)
				}
				return nil
			},
		},
		{
			name:      "arch returns valid architecture",
			operation: "arch",
			validate: func(got string) error {
				if got != runtime.GOARCH {
					return fmt.Errorf("expected arch %s, got %s", runtime.GOARCH, got)
				}
				return nil
			},
		},
		{
			name:      "env returns environment variable",
			operation: "env",
			value:     testEnvVar,
			validate: func(got string) error {
				if got != testEnvValue {
					return fmt.Errorf("expected env var %s, got %s", testEnvValue, got)
				}
				return nil
			},
		},
		{
			name:      "pwd returns valid directory",
			operation: "pwd",
			validate: func(got string) error {
				if !filepath.IsAbs(got) {
					return fmt.Errorf("expected absolute path, got %s", got)
				}
				return nil
			},
		},
		{
			name:      "home returns valid home directory",
			operation: "home",
			validate: func(got string) error {
				if !filepath.IsAbs(got) {
					return fmt.Errorf("expected absolute path, got %s", got)
				}
				if !strings.Contains(got, "home") && !strings.Contains(got, "Users") && got != "/root" {
					return fmt.Errorf("path %s doesn't look like a home directory", got)
				}
				return nil
			},
		},
		// Error cases
		{
			name:      "unknown operation",
			operation: "invalid",
			wantErr:   true,
		},
		{
			name:      "env without variable",
			operation: "env",
			wantErr:   true,
		},
		{
			name:      "env with non-existent variable",
			operation: "env",
			value:     "NONEXISTENT_VAR_123456",
			validate: func(got string) error {
				if got != "" {
					return fmt.Errorf("expected empty string for non-existent env var, got %s", got)
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := plugin.Apply(tt.operation, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SysPlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.validate != nil {
				if err := tt.validate(got); err != nil {
					t.Errorf("SysPlugin.Apply() validation failed: %v", err)
				}
			}
		})
	}
}
