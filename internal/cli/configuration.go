package cli

import (
	"github.com/danielmiessler/fabric/internal/core"
)

// handleConfigurationCommands runs the pattern update and default model commands.
// It returns handled = true when a command ran and the caller must exit.
func handleConfigurationCommands(currentFlags *Flags, registry *core.PluginRegistry) (handled bool, err error) {
	if currentFlags.UpdatePatterns {
		if err = registry.PatternsLoader.PopulateDB(); err != nil {
			return true, err
		}
		// PopulateDB can change the patterns folder path. Save the env file so the new path persists.
		err = registry.SaveEnvFile()
		return true, err
	}

	if currentFlags.ChangeDefaultModel {
		if err = registry.Defaults.Setup(); err != nil {
			return true, err
		}
		err = registry.SaveEnvFile()
		return true, err
	}

	return false, nil
}
