package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	restapi "github.com/danielmiessler/fabric/internal/server"
)

// Tests replace serveOllama, because restapi.ServeOllama blocks on a listening socket.
var serveOllama = restapi.ServeOllama

// handleSetupAndServerCommands runs the setup and server commands.
// It returns handled = true when a command ran and the caller must exit.
func handleSetupAndServerCommands(currentFlags *Flags, registry *core.PluginRegistry, version string) (handled bool, err error) {
	if currentFlags.Setup {
		err = registry.Setup()
		return true, err
	}

	// The server handlers do not see the parsed flags. Put the yt-dlp
	// arguments on the plugin so the /youtube/transcript endpoint uses them.
	registry.YouTube.YtDlpArgs = currentFlags.YtDlpArgs

	// The flags are parsed before initializeFabric loads
	// ~/.config/fabric/.env. Read the server variables again, so that
	// file can set them. A flag or a shell variable is already set and wins.
	if currentFlags.ServeAPIKey == "" {
		currentFlags.ServeAPIKey = os.Getenv("FABRIC_API_KEY")
	}
	if v := os.Getenv("FABRIC_CORS_ORIGINS"); len(currentFlags.ServeCORSOrigins) == 0 && v != "" {
		currentFlags.ServeCORSOrigins = strings.Split(v, ",")
	}

	// Show a warning when the API key is on the command line. Other local
	// users can see command-line arguments in the process list.
	if (currentFlags.Serve || currentFlags.ServeOllama) && apiKeyPassedAsArg(os.Args[1:]) {
		fmt.Fprintln(os.Stderr, i18n.T("server_api_key_arg_warning"))
	}

	if currentFlags.Serve {
		registry.ConfigureVendors()
		err = restapi.Serve(registry, currentFlags.ServeAddress, currentFlags.ServeAPIKey, currentFlags.ServeCORSOrigins)
		return true, err
	}

	if currentFlags.ServeOllama {
		registry.ConfigureVendors()
		err = serveOllama(registry, currentFlags.ServeAddress, version, currentFlags.ServeAPIKey, currentFlags.ServeCORSOrigins)
		return true, err
	}

	return false, nil
}

// apiKeyPassedAsArg reports whether the REST API key is on the command line,
// not in the FABRIC_API_KEY environment variable or the .env file.
func apiKeyPassedAsArg(args []string) bool {
	for _, a := range args {
		if a == "--api-key" || strings.HasPrefix(a, "--api-key=") {
			return true
		}
	}
	return false
}
