package cli

import (
	"github.com/danielmiessler/fabric/internal/core"
)

// handleExtensionCommands handles extension-related commands
// Returns (handled, error) where handled indicates if a command was processed and should exit
func handleExtensionCommands(currentFlags *Flags, registry *core.PluginRegistry) (handled bool, err error) {
	if currentFlags.ListExtensions {
		err = registry.TemplateExtensions.ListExtensions()
		return true, err
	}

	if currentFlags.AddExtension != "" {
		err = registry.TemplateExtensions.RegisterExtension(currentFlags.AddExtension)
		return true, err
	}

	if currentFlags.RemoveExtension != "" {
		err = registry.TemplateExtensions.RemoveExtension(currentFlags.RemoveExtension)
		return true, err
	}

	return false, nil
}
