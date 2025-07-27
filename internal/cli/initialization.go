package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

const ConfigDirPerms os.FileMode = 0755
const EnvFilePerms os.FileMode = 0644

// initializeFabric initializes the fabric database and plugin registry
func initializeFabric() (registry *core.PluginRegistry, err error) {
	var homedir string
	if homedir, err = os.UserHomeDir(); err != nil {
		return
	}

	fabricDb := fsdb.NewDb(filepath.Join(homedir, ".config/fabric"))
	if err = fabricDb.Configure(); err != nil {
		return
	}

	if registry, err = core.NewPluginRegistry(fabricDb); err != nil {
		return
	}

	return
}

// ensureEnvFile checks for the default ~/.config/fabric/.env file and creates it
// along with the parent directory if it does not exist.
func ensureEnvFile() (err error) {
	var homedir string
	if homedir, err = os.UserHomeDir(); err != nil {
		return fmt.Errorf("could not determine user home directory: %w", err)
	}
	configDir := filepath.Join(homedir, ".config", "fabric")
	envPath := filepath.Join(configDir, ".env")

	if _, statErr := os.Stat(envPath); statErr != nil {
		if !os.IsNotExist(statErr) {
			return fmt.Errorf("could not stat .env file: %w", statErr)
		}
		if err = os.MkdirAll(configDir, ConfigDirPerms); err != nil {
			return fmt.Errorf("could not create config directory: %w", err)
		}
		if err = os.WriteFile(envPath, []byte{}, EnvFilePerms); err != nil {
			return fmt.Errorf("could not create .env file: %w", err)
		}
	}
	return
}
