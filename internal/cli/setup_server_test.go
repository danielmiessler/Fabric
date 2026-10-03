package cli

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
)

// The --serveOllama path must pass the address, version, and API key
// flags through to ServeOllama unchanged, and must put the yt-dlp
// arguments on the YouTube plugin.
func TestHandleSetupAndServerCommands_ServeOllamaWiring(t *testing.T) {
	var gotAddress, gotVersion, gotKey string
	prev := serveOllama
	serveOllama = func(_ *core.PluginRegistry, address, version, apiKey string) error {
		gotAddress, gotVersion, gotKey = address, version, apiKey
		return nil
	}
	defer func() { serveOllama = prev }()

	registry := &core.PluginRegistry{
		VendorManager: ai.NewVendorsManager(),
		VendorsAll:    ai.NewVendorsManager(),
		YouTube:       youtube.NewYouTube(),
	}
	flags := &Flags{ServeOllama: true, ServeAddress: "127.0.0.1:9999", ServeAPIKey: "secret", YtDlpArgs: "--sub-lang en-orig"}
	handled, err := handleSetupAndServerCommands(flags, registry, "v-test")
	if err != nil {
		t.Fatalf("handleSetupAndServerCommands() error = %v", err)
	}
	if !handled {
		t.Fatal("handleSetupAndServerCommands() handled = false, want true")
	}
	if gotAddress != "127.0.0.1:9999" || gotVersion != "v-test" || gotKey != "secret" {
		t.Fatalf("ServeOllama got (%q, %q, %q), want (127.0.0.1:9999, v-test, secret)",
			gotAddress, gotVersion, gotKey)
	}
	if registry.YouTube.YtDlpArgs != "--sub-lang en-orig" {
		t.Fatalf("registry.YouTube.YtDlpArgs = %q, want %q", registry.YouTube.YtDlpArgs, "--sub-lang en-orig")
	}
}
