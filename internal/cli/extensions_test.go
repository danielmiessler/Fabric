package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/template"
	"github.com/stretchr/testify/require"
)

func TestHandleExtensionCommandsNoop(t *testing.T) {
	handled, err := handleExtensionCommands(&Flags{}, nil)
	require.False(t, handled)
	require.NoError(t, err)
}

func TestHandleExtensionCommandsAddListAndRemove(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "echo-extension.yaml")
	scriptPath := filepath.Join(tempDir, "echo-extension.sh")

	require.NoError(t, os.WriteFile(scriptPath, []byte("#!/bin/sh\necho \"$1\"\n"), 0o755))
	require.NoError(t, os.WriteFile(configPath, []byte(`name: echo-extension
executable: `+scriptPath+`
type: executable
timeout: 5s
description: "Echo test extension"
version: "1.0.0"
operations:
  echo:
    cmd_template: "{{executable}} {{1}}"
`), 0o644))

	registry := &core.PluginRegistry{
		TemplateExtensions: template.NewExtensionManager(tempDir),
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handleExtensionCommands(&Flags{AddExtension: configPath}, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "echo-extension")

	stdout, err = captureStdout(func() error {
		handled, runErr := handleExtensionCommands(&Flags{ListExtensions: true}, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "echo-extension")
	require.Contains(t, stdout, scriptPath)

	handled, err := handleExtensionCommands(&Flags{RemoveExtension: "echo-extension"}, registry)
	require.True(t, handled)
	require.NoError(t, err)

	_, processErr := registry.TemplateExtensions.ProcessExtension("echo-extension", "echo", "hello")
	require.Error(t, processErr)
}
