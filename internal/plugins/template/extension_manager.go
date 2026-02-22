package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"gopkg.in/yaml.v3"
)

// ExtensionManager handles the high-level operations of the extension system
type ExtensionManager struct {
	registry  *ExtensionRegistry
	executor  *ExtensionExecutor
	configDir string
}

// NewExtensionManager creates a new extension manager instance
func NewExtensionManager(configDir string) *ExtensionManager {
	registry := NewExtensionRegistry(configDir)
	return &ExtensionManager{
		registry:  registry,
		executor:  NewExtensionExecutor(registry),
		configDir: configDir,
	}
}

// ListExtensions handles the listextensions flag action
func (em *ExtensionManager) ListExtensions() error {
	if em.registry == nil || em.registry.registry.Extensions == nil {
		return errors.New(i18n.T("extension_registry_not_initialized"))
	}

	for name, entry := range em.registry.registry.Extensions {
		fmt.Printf(i18n.T("extension_name_label"), name)

		// Try to load extension details
		ext, err := em.registry.GetExtension(name)
		if err != nil {
			fmt.Printf(i18n.T("extension_status_disabled"), err)
			fmt.Printf(i18n.T("extension_config_path_label"), entry.ConfigPath)
			continue
		}

		// Print extension details if verification succeeded
		fmt.Printf("%s", i18n.T("extension_status_enabled"))
		fmt.Printf(i18n.T("extension_executable_label"), ext.Executable)
		fmt.Printf(i18n.T("extension_type_label"), ext.Type)
		fmt.Printf(i18n.T("extension_timeout_label"), ext.Timeout)
		fmt.Printf(i18n.T("extension_description_label"), ext.Description)
		fmt.Printf(i18n.T("extension_version_label"), ext.Version)

		fmt.Printf("%s", i18n.T("extension_operations_label"))
		for opName, opConfig := range ext.Operations {
			fmt.Printf("    %s:\n", opName)
			fmt.Printf(i18n.T("extension_command_template_label"), opConfig.CmdTemplate)
		}

		if fileConfig := ext.GetFileConfig(); fileConfig != nil {
			fmt.Printf("%s", i18n.T("extension_file_config_label"))
			for k, v := range fileConfig {
				fmt.Printf("    %s: %v\n", k, v)
			}
		}
		fmt.Printf("\n")
	}

	return nil
}

// RegisterExtension handles the addextension flag action
func (em *ExtensionManager) RegisterExtension(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_invalid_config_path"), err)
	}

	// Get extension name before registration for status message
	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_read_config"), err)
	}

	var ext ExtensionDefinition
	if err := yaml.Unmarshal(data, &ext); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_parse_config"), err)
	}

	if err := em.registry.Register(absPath); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_register"), err)
	}

	if _, err := time.ParseDuration(ext.Timeout); err != nil {
		return fmt.Errorf(i18n.T("extension_invalid_timeout"), ext.Timeout, err)
	}

	// Print success message with extension details
	fmt.Printf("%s", i18n.T("extension_registered_success"))
	fmt.Printf(i18n.T("extension_name_detail_label"), ext.Name)
	fmt.Printf(i18n.T("extension_executable_label"), ext.Executable)
	fmt.Printf(i18n.T("extension_type_label"), ext.Type)
	fmt.Printf(i18n.T("extension_timeout_label"), ext.Timeout)
	fmt.Printf(i18n.T("extension_description_label"), ext.Description)
	fmt.Printf(i18n.T("extension_version_label"), ext.Version)

	fmt.Printf("%s", i18n.T("extension_operations_label"))
	for opName, opConfig := range ext.Operations {
		fmt.Printf("    %s:\n", opName)
		fmt.Printf(i18n.T("extension_command_template_label"), opConfig.CmdTemplate)
	}

	if fileConfig := ext.GetFileConfig(); fileConfig != nil {
		fmt.Printf("%s", i18n.T("extension_file_config_label"))
		for k, v := range fileConfig {
			fmt.Printf("    %s: %v\n", k, v)
		}
	}

	return nil
}

// RemoveExtension handles the rmextension flag action
func (em *ExtensionManager) RemoveExtension(name string) error {
	if err := em.registry.Remove(name); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_remove"), err)
	}

	return nil
}

// ProcessExtension handles template processing for extension directives
func (em *ExtensionManager) ProcessExtension(name, operation, value string) (string, error) {
	return em.executor.Execute(name, operation, value)
}
