package custom_patterns

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

func NewCustomPatterns() (ret *CustomPatterns) {
	label := "Custom Patterns"
	ret = &CustomPatterns{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             i18n.T("custom_patterns_label"),
		SetupDescription: i18n.T("custom_patterns_setup_description") + " " + i18n.T("optional_marker"),
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  ret.configure,
	}

	ret.CustomPatternsDir = ret.AddSetupQuestionWithEnvName("Directory", false,
		i18n.T("custom_patterns_directory_question"))

	return
}

type CustomPatterns struct {
	*plugins.PluginBase
	CustomPatternsDir *plugins.SetupQuestion
}

func (o *CustomPatterns) configure() error {
	if o.CustomPatternsDir.Value != "" {
		if strings.HasPrefix(o.CustomPatternsDir.Value, "~/") {
			if homeDir, err := os.UserHomeDir(); err == nil {
				o.CustomPatternsDir.Value = filepath.Join(homeDir, o.CustomPatternsDir.Value[2:])
			}
		}

		if absPath, err := filepath.Abs(o.CustomPatternsDir.Value); err == nil {
			o.CustomPatternsDir.Value = absPath
		}

		if _, err := os.Stat(o.CustomPatternsDir.Value); os.IsNotExist(err) {
			if err := os.MkdirAll(o.CustomPatternsDir.Value, 0755); err != nil {
				// Keep the value when the create fails, so setup still writes it to the env file.
				fmt.Printf(i18n.T("custom_patterns_warning_create_directory"), o.CustomPatternsDir.Value, err)
			}
		}
	}

	return nil
}

// IsConfigured returns true if a custom patterns directory has been set
func (o *CustomPatterns) IsConfigured() bool {
	// Configure first, so the check sees the value from the environment.
	o.Configure()
	return o.CustomPatternsDir.Value != ""
}
