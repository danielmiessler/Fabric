package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/gemini"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/plugins/strategy"
	"github.com/danielmiessler/fabric/internal/tools"
	"github.com/stretchr/testify/require"
)

func TestHandleListingCommandsListPipelines(t *testing.T) {
	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")
	homeDir := filepath.Join(tempDir, "home")
	userDir := filepath.Join(homeDir, ".config", "fabric", "pipelines")

	require.NoError(t, os.MkdirAll(builtInDir, 0o755))
	require.NoError(t, os.MkdirAll(userDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "alpha.yaml"), []byte(minimalPipelineYAML("alpha")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "beta.yaml"), []byte(minimalPipelineYAML("beta")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(userDir, "beta.yaml"), []byte(minimalPipelineYAML("beta")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(userDir, "gamma.yaml"), []byte(minimalPipelineYAML("gamma")), 0o644))

	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", builtInDir)
	t.Setenv("HOME", homeDir)

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns: "0",
			ListPipelines:  true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(stdout), "\n")
	require.Equal(t, []string{
		"alpha\tbuilt-in",
		"beta\tuser (overrides built-in)",
		"gamma\tuser",
	}, lines)
}

func TestHandleListingCommandsListPipelinesShellComplete(t *testing.T) {
	tempDir := t.TempDir()
	builtInDir := filepath.Join(tempDir, "builtins")

	require.NoError(t, os.MkdirAll(builtInDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "zeta.yaml"), []byte(minimalPipelineYAML("zeta")), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(builtInDir, "alpha.yaml"), []byte(minimalPipelineYAML("alpha")), 0o644))

	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", builtInDir)
	t.Setenv("HOME", filepath.Join(tempDir, "home"))

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:      "0",
			ListPipelines:       true,
			ShellCompleteOutput: true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)

	require.Equal(t, []string{"alpha", "zeta"}, strings.Split(strings.TrimSpace(stdout), "\n"))
}

func TestHandleListingCommandsListPipelinesLoaderError(t *testing.T) {
	tempDir := t.TempDir()
	notDirPath := filepath.Join(tempDir, "not-a-dir")
	require.NoError(t, os.WriteFile(notDirPath, []byte("nope"), 0o644))

	t.Setenv("FABRIC_BUILTIN_PIPELINES_DIR", notDirPath)
	t.Setenv("HOME", filepath.Join(tempDir, "home"))

	handled, err := handleListingCommands(&Flags{
		LatestPatterns: "0",
		ListPipelines:  true,
	}, nil, nil)
	require.True(t, handled)
	require.Error(t, err)
	require.Contains(t, err.Error(), "is not a directory")
}

func TestHandleListingCommandsLatestPatterns(t *testing.T) {
	db := fsdb.NewDb(t.TempDir())
	require.NoError(t, os.WriteFile(db.Patterns.UniquePatternsFilePath, []byte("one\ntwo\nthree"), 0o644))

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{LatestPatterns: "2"}, db, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(stdout), "\n")
	require.Equal(t, []string{"three", "two"}, lines)
}

func TestHandleListingCommandsLatestPatternsParseError(t *testing.T) {
	handled, err := handleListingCommands(&Flags{LatestPatterns: "not-a-number"}, nil, nil)
	require.True(t, handled)
	require.Error(t, err)
}

func TestHandleListingCommandsListPatternsGuidanceWhenEmpty(t *testing.T) {
	db := newConfiguredTestDB(t)

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns: "0",
			ListPatterns:   true,
		}, db, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "No patterns found")
	require.Contains(t, stdout, "fabric --setup")
	require.Contains(t, stdout, "fabric -U")
}

func TestHandleListingCommandsListPatternsShellComplete(t *testing.T) {
	db := newConfiguredTestDB(t)
	require.NoError(t, os.MkdirAll(filepath.Join(db.Patterns.Dir, "alpha"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(db.Patterns.Dir, "beta"), 0o755))

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:      "0",
			ListPatterns:        true,
			ShellCompleteOutput: true,
		}, db, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "alpha")
	require.Contains(t, stdout, "beta")
}

func TestHandleListingCommandsListTranscriptionModels(t *testing.T) {
	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:          "0",
			ListTranscriptionModels: true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "Available transcription models:")
	require.Contains(t, stdout, "whisper-1")
	require.Contains(t, stdout, "gpt-4o-mini-transcribe")
	require.Contains(t, stdout, "gpt-4o-transcribe")
}

func TestHandleListingCommandsListTranscriptionModelsShellComplete(t *testing.T) {
	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:          "0",
			ListTranscriptionModels: true,
			ShellCompleteOutput:     true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Equal(t, []string{"whisper-1", "gpt-4o-mini-transcribe", "gpt-4o-transcribe"}, strings.Split(strings.TrimSpace(stdout), "\n"))
}

func TestHandleListingCommandsListAllModels(t *testing.T) {
	models := ai.NewVendorsModels()
	models.AddGroupItems("OpenAI", "gpt-4o", "gpt-4o-mini")
	models.AddGroupItems("Anthropic", "claude-3-7-sonnet")

	manager := ai.NewVendorsManager()
	manager.Models = models

	registry := &core.PluginRegistry{
		VendorManager: manager,
		Defaults: &tools.Defaults{
			PluginBase: &plugins.PluginBase{},
			Vendor:     &plugins.Setting{Value: "OpenAI"},
			Model:      &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "gpt-4o"}},
		},
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns: "0",
			ListAllModels:  true,
		}, nil, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "Available models:")
	require.Contains(t, stdout, "Anthropic|claude-3-7-sonnet")
	require.Contains(t, stdout, "OpenAI|gpt-4o")
	require.Contains(t, stdout, "OpenAI|gpt-4o-mini")
	require.Contains(t, stdout, "*")
}

func TestHandleListingCommandsListAllModelsShellCompleteWithVendorFilter(t *testing.T) {
	models := ai.NewVendorsModels()
	models.AddGroupItems("OpenAI", "gpt-4o", "gpt-4o-mini")
	models.AddGroupItems("Anthropic", "claude-3-7-sonnet")

	manager := ai.NewVendorsManager()
	manager.Models = models

	registry := &core.PluginRegistry{
		VendorManager: manager,
		Defaults: &tools.Defaults{
			PluginBase: &plugins.PluginBase{},
			Vendor:     &plugins.Setting{Value: "OpenAI"},
			Model:      &plugins.SetupQuestion{Setting: &plugins.Setting{Value: "gpt-4o"}},
		},
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:      "0",
			ListAllModels:       true,
			ShellCompleteOutput: true,
			Vendor:              "Anthropic",
		}, nil, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Equal(t, "claude-3-7-sonnet\n", stdout)
}

func TestHandleListingCommandsListVendors(t *testing.T) {
	vendorsAll := ai.NewVendorsManager()
	vendorsAll.AddVendors(
		&cliStubVendor{name: "Anthropic"},
		&cliStubVendor{name: "OpenAI"},
	)

	registry := &core.PluginRegistry{VendorsAll: vendorsAll}

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns: "0",
			ListVendors:    true,
		}, nil, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "Available Vendors:")
	require.Contains(t, stdout, "Anthropic")
	require.Contains(t, stdout, "OpenAI")
}

func TestHandleListingCommandsListAllContexts(t *testing.T) {
	db := newConfiguredTestDB(t)
	require.NoError(t, db.Contexts.Save("project", []byte("ctx")))
	require.NoError(t, db.Contexts.Save("notes", []byte("ctx")))

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:  "0",
			ListAllContexts: true,
		}, db, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "project")
	require.Contains(t, stdout, "notes")
}

func TestHandleListingCommandsListAllSessions(t *testing.T) {
	db := newConfiguredTestDB(t)
	require.NoError(t, db.Sessions.SaveSession(&fsdb.Session{Name: "session-a"}))
	require.NoError(t, db.Sessions.SaveSession(&fsdb.Session{Name: "session-b"}))

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:  "0",
			ListAllSessions: true,
		}, db, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "session-a")
	require.Contains(t, stdout, "session-b")
}

func TestHandleListingCommandsListStrategies(t *testing.T) {
	registry := &core.PluginRegistry{
		Strategies: &strategy.StrategiesManager{
			Strategies: map[string]strategy.Strategy{
				"fast": {Name: "fast", Description: "Low-latency strategy"},
				"deep": {Name: "deep", Description: "High-depth strategy"},
			},
		},
	}

	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns: "0",
			ListStrategies: true,
		}, nil, registry)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "Available Strategies:")
	require.Contains(t, stdout, "fast")
	require.Contains(t, stdout, "deep")
}

func TestHandleListingCommandsListGeminiVoices(t *testing.T) {
	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:   "0",
			ListGeminiVoices: true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "Available Gemini Text-to-Speech voices:")
	require.Contains(t, stdout, "Kore")
}

func TestHandleListingCommandsListGeminiVoicesShellComplete(t *testing.T) {
	stdout, err := captureStdout(func() error {
		handled, runErr := handleListingCommands(&Flags{
			LatestPatterns:      "0",
			ListGeminiVoices:    true,
			ShellCompleteOutput: true,
		}, nil, nil)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, stdout, gemini.GetGeminiVoiceNames()[0])
}

func minimalPipelineYAML(name string) string {
	return "version: 1\n" +
		"name: " + name + "\n" +
		"stages:\n" +
		"  - id: render\n" +
		"    executor: builtin\n" +
		"    builtin:\n" +
		"      name: passthrough\n" +
		"    final_output: true\n" +
		"    primary_output:\n" +
		"      from: stdout\n"
}
