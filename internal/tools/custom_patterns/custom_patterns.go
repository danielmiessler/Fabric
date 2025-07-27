package custom_patterns

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins"
)

func NewCustomPatterns() (ret *CustomPatterns) {
	label := "Custom Patterns"
	ret = &CustomPatterns{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             label,
		SetupDescription: "Custom Patterns - Set directory for your custom patterns (optional)",
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  ret.configure,
	}

	ret.CustomPatternsDir = ret.AddSetupQuestionCustom("Directory", false,
		"Enter the path to your custom patterns directory (leave empty to skip)")

	return
}

type CustomPatterns struct {
	*plugins.PluginBase
	CustomPatternsDir *plugins.SetupQuestion
}

func (o *CustomPatterns) configure() error {
	if o.CustomPatternsDir.Value != "" {
		// Expand home directory if needed
		if strings.HasPrefix(o.CustomPatternsDir.Value, "~/") {
			if homeDir, err := os.UserHomeDir(); err == nil {
				o.CustomPatternsDir.Value = filepath.Join(homeDir, o.CustomPatternsDir.Value[2:])
			}
		}

		// Convert to absolute path
		if absPath, err := filepath.Abs(o.CustomPatternsDir.Value); err == nil {
			o.CustomPatternsDir.Value = absPath
		}

		// Check if directory exists, create only if it doesn't
		if _, err := os.Stat(o.CustomPatternsDir.Value); os.IsNotExist(err) {
			if err := os.MkdirAll(o.CustomPatternsDir.Value, 0755); err != nil {
				// Log the error but don't clear the value - let it persist in env file
				fmt.Printf("Warning: Could not create custom patterns directory %s: %v\n", o.CustomPatternsDir.Value, err)
			}
		}
	}

	return nil
}

// IsConfigured returns true if a custom patterns directory has been set
func (o *CustomPatterns) IsConfigured() bool {
	// First configure to load values from environment variables
	o.Configure()
	// Check if the plugin has been configured with a directory
	return o.CustomPatternsDir.Value != ""
}
