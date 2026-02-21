package template

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"

	"gopkg.in/yaml.v3"
)

// ExtensionDefinition represents a single extension configuration
type ExtensionDefinition struct {
	// Global properties
	Name        string   `yaml:"name"`
	Executable  string   `yaml:"executable"`
	Type        string   `yaml:"type"`
	Timeout     string   `yaml:"timeout"`
	Description string   `yaml:"description"`
	Version     string   `yaml:"version"`
	Env         []string `yaml:"env"`

	// Operation-specific commands
	Operations map[string]OperationConfig `yaml:"operations"`

	// Additional config
	Config map[string]any `yaml:"config"`
}

type OperationConfig struct {
	CmdTemplate string `yaml:"cmd_template"`
}

// RegistryEntry represents a registered extension
type RegistryEntry struct {
	ConfigPath     string `yaml:"config_path"`
	ConfigHash     string `yaml:"config_hash"`
	ExecutableHash string `yaml:"executable_hash"`
}

type ExtensionRegistry struct {
	configDir string
	registry  struct {
		Extensions map[string]*RegistryEntry `yaml:"extensions"`
	}
}

// Helper methods for Config access
func (e *ExtensionDefinition) GetOutputMethod() string {
	if output, ok := e.Config["output"].(map[string]any); ok {
		if method, ok := output["method"].(string); ok {
			return method
		}
	}
	return "stdout" // default to stdout if not specified
}

func (e *ExtensionDefinition) GetFileConfig() map[string]any {
	if output, ok := e.Config["output"].(map[string]any); ok {
		if fileConfig, ok := output["file_config"].(map[string]any); ok {
			return fileConfig
		}
	}
	return nil
}

func (e *ExtensionDefinition) IsCleanupEnabled() bool {
	if fc := e.GetFileConfig(); fc != nil {
		if cleanup, ok := fc["cleanup"].(bool); ok {
			return cleanup
		}
	}
	return false // default to no cleanup
}

func NewExtensionRegistry(configDir string) *ExtensionRegistry {
	r := &ExtensionRegistry{
		configDir: configDir,
	}
	r.registry.Extensions = make(map[string]*RegistryEntry)

	r.ensureConfigDir()

	if err := r.loadRegistry(); err != nil {
		debuglog.Log(i18n.T("extension_warning_load_registry"), err)
	}

	return r
}

func (r *ExtensionRegistry) ensureConfigDir() error {
	extDir := filepath.Join(r.configDir, "extensions")
	return os.MkdirAll(extDir, 0755)
}

// Update the Register method in extension_registry.go

func (r *ExtensionRegistry) Register(configPath string) error {
	// Read and parse the extension definition to verify it
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_read_config"), err)
	}

	var ext ExtensionDefinition
	if err := yaml.Unmarshal(data, &ext); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_parse_config"), err)
	}

	// Validate extension name
	if ext.Name == "" {
		return errors.New(i18n.T("extension_name_empty"))
	}

	if strings.Contains(ext.Name, " ") {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_name_contains_spaces"), ext.Name))
	}

	// Verify executable exists
	if _, err := os.Stat(ext.Executable); err != nil {
		return fmt.Errorf(i18n.T("extension_executable_not_found"), err)
	}

	// Get absolute path to config
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_get_absolute_path"), err)
	}

	// Calculate hashes
	configHash := ComputeStringHash(string(data))
	executableHash, err := ComputeHash(ext.Executable)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_hash_executable"), err)
	}

	// Validate full extension definition (ensures operations and cmd_template present)
	if err := r.validateExtensionDefinition(&ext); err != nil {
		return fmt.Errorf(i18n.T("extension_invalid_definition"), err)
	}

	// Store entry
	r.registry.Extensions[ext.Name] = &RegistryEntry{
		ConfigPath:     absPath,
		ConfigHash:     configHash,
		ExecutableHash: executableHash,
	}

	return r.saveRegistry()
}

func (r *ExtensionRegistry) validateExtensionDefinition(ext *ExtensionDefinition) error {
	// Validate required fields
	if ext.Name == "" {
		return errors.New(i18n.T("extension_name_required"))
	}
	if ext.Executable == "" {
		return errors.New(i18n.T("extension_executable_required"))
	}
	if ext.Type == "" {
		return errors.New(i18n.T("extension_type_required"))
	}

	// Validate timeout format
	if ext.Timeout != "" {
		if _, err := time.ParseDuration(ext.Timeout); err != nil {
			return fmt.Errorf(i18n.T("extension_invalid_timeout_format"), err)
		}
	}

	// Validate operations
	if len(ext.Operations) == 0 {
		return errors.New(i18n.T("extension_operation_required"))
	}
	for name, op := range ext.Operations {
		if op.CmdTemplate == "" {
			return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_cmd_template_required"), name))
		}
	}

	return nil
}

func (r *ExtensionRegistry) Remove(name string) error {
	if _, exists := r.registry.Extensions[name]; !exists {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_not_found"), name))
	}

	delete(r.registry.Extensions, name)

	return r.saveRegistry()
}

func (r *ExtensionRegistry) Verify(name string) error {
	// Get the registry entry
	entry, exists := r.registry.Extensions[name]
	if !exists {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_not_found"), name))
	}

	// Load and parse the config file
	data, err := os.ReadFile(entry.ConfigPath)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_read_config"), err)
	}

	// Verify config hash
	currentConfigHash := ComputeStringHash(string(data))
	if currentConfigHash != entry.ConfigHash {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_config_hash_mismatch"), name))
	}

	// Parse to get executable path
	var ext ExtensionDefinition
	if err := yaml.Unmarshal(data, &ext); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_parse_config"), err)
	}

	// Verify executable hash
	currentExecutableHash, err := ComputeHash(ext.Executable)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_verify_executable"), err)
	}

	if currentExecutableHash != entry.ExecutableHash {
		return fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_executable_hash_mismatch"), name))
	}

	return nil
}

func (r *ExtensionRegistry) GetExtension(name string) (*ExtensionDefinition, error) {
	entry, exists := r.registry.Extensions[name]
	if !exists {
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_not_found"), name))
	}

	// Read current config file
	data, err := os.ReadFile(entry.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("extension_failed_read_config"), err)
	}

	// Verify config hash
	currentHash := ComputeStringHash(string(data))
	if currentHash != entry.ConfigHash {
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_config_hash_mismatch"), name))
	}

	// Parse config
	var ext ExtensionDefinition
	if err := yaml.Unmarshal(data, &ext); err != nil {
		return nil, fmt.Errorf(i18n.T("extension_failed_parse_config"), err)
	}

	// Verify executable hash
	currentExecHash, err := ComputeHash(ext.Executable)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("extension_failed_verify_executable"), err)
	}

	if currentExecHash != entry.ExecutableHash {
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("extension_executable_hash_mismatch"), name))
	}

	return &ext, nil
}

func (r *ExtensionRegistry) ListExtensions() ([]*ExtensionDefinition, error) {
	var exts []*ExtensionDefinition

	for name := range r.registry.Extensions {
		ext, err := r.GetExtension(name)
		if err != nil {
			// Instead of failing, we'll return nil for this extension
			// The manager will handle displaying the error
			exts = append(exts, nil)
			continue
		}
		exts = append(exts, ext)
	}

	return exts, nil
}

func (r *ExtensionRegistry) calculateFileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func (r *ExtensionRegistry) saveRegistry() error {
	data, err := yaml.Marshal(r.registry)
	if err != nil {
		return fmt.Errorf(i18n.T("extension_failed_marshal_registry"), err)
	}

	registryPath := filepath.Join(r.configDir, "extensions", "extensions.yaml")
	return os.WriteFile(registryPath, data, 0644)
}

func (r *ExtensionRegistry) loadRegistry() error {
	registryPath := filepath.Join(r.configDir, "extensions", "extensions.yaml")
	data, err := os.ReadFile(registryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // New registry
		}
		return fmt.Errorf(i18n.T("extension_failed_read_registry"), err)
	}

	// Need to unmarshal the data into our registry
	if err := yaml.Unmarshal(data, &r.registry); err != nil {
		return fmt.Errorf(i18n.T("extension_failed_parse_registry"), err)
	}

	return nil
}
