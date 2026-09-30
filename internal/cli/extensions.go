package cli

import (
	"github.com/danielmiessler/fabric/internal/core"
)

// handleExtensionCommands runs the extension list, add, and remove commands.
// It returns handled = true when a command ran and the caller must exit.
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
