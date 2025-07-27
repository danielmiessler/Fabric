package cli

import (
	"github.com/danielmiessler/fabric/internal/core"
)

// handleConfigurationCommands handles configuration-related commands
// Returns (handled, error) where handled indicates if a command was processed and should exit
func handleConfigurationCommands(currentFlags *Flags, registry *core.PluginRegistry) (handled bool, err error) {
	if currentFlags.UpdatePatterns {
		if err = registry.PatternsLoader.PopulateDB(); err != nil {
			return true, err
		}
		// Save configuration in case any paths were migrated during pattern loading
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
