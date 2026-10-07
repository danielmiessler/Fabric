package cli

import (
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
)

// The --serveOllama path must pass the address, version, API key, and
// CORS origin flags through to ServeOllama unchanged, and must put the yt-dlp
// arguments on the YouTube plugin.
func TestHandleSetupAndServerCommands_ServeOllamaWiring(t *testing.T) {
	var gotAddress, gotVersion, gotKey string
	var gotOrigins []string
	prev := serveOllama
	serveOllama = func(_ *core.PluginRegistry, address, version, apiKey string, corsOrigins []string) error {
		gotAddress, gotVersion, gotKey, gotOrigins = address, version, apiKey, corsOrigins
		return nil
	}
	defer func() { serveOllama = prev }()

	registry := &core.PluginRegistry{
		VendorManager: ai.NewVendorsManager(),
		VendorsAll:    ai.NewVendorsManager(),
		YouTube:       youtube.NewYouTube(),
	}
	flags := &Flags{ServeOllama: true, ServeAddress: "127.0.0.1:9999", ServeAPIKey: "secret", ServeCORSOrigins: []string{"http://localhost:1420"}, YtDlpArgs: "--sub-lang en-orig"}
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
	if len(gotOrigins) != 1 || gotOrigins[0] != "http://localhost:1420" {
		t.Fatalf("ServeOllama got corsOrigins %v, want [http://localhost:1420]", gotOrigins)
	}
	if registry.YouTube.YtDlpArgs != "--sub-lang en-orig" {
		t.Fatalf("registry.YouTube.YtDlpArgs = %q, want %q", registry.YouTube.YtDlpArgs, "--sub-lang en-orig")
	}
}

// The flags are parsed before the .env file loads. Values from that file
// reach the server when the command line did not set them.
func TestHandleSetupAndServerCommands_ServeOllamaEnvFile(t *testing.T) {
	t.Setenv("FABRIC_API_KEY", "from-env")
	t.Setenv("FABRIC_CORS_ORIGINS", "http://a.example,http://b.example")
	var gotKey string
	var gotOrigins []string
	prev := serveOllama
	serveOllama = func(_ *core.PluginRegistry, _, _, apiKey string, corsOrigins []string) error {
		gotKey, gotOrigins = apiKey, corsOrigins
		return nil
	}
	defer func() { serveOllama = prev }()

	registry := &core.PluginRegistry{
		VendorManager: ai.NewVendorsManager(),
		VendorsAll:    ai.NewVendorsManager(),
		YouTube:       youtube.NewYouTube(),
	}
	if _, err := handleSetupAndServerCommands(&Flags{ServeOllama: true}, registry, "v"); err != nil {
		t.Fatal(err)
	}
	if gotKey != "from-env" || len(gotOrigins) != 2 || gotOrigins[1] != "http://b.example" {
		t.Fatalf("got key %q, origins %q", gotKey, gotOrigins)
	}

	flags := &Flags{ServeOllama: true, ServeAPIKey: "flag", ServeCORSOrigins: []string{"http://c.example"}}
	if _, err := handleSetupAndServerCommands(flags, registry, "v"); err != nil {
		t.Fatal(err)
	}
	if gotKey != "flag" || len(gotOrigins) != 1 {
		t.Fatalf("flags must win: got key %q, origins %q", gotKey, gotOrigins)
	}
}

// TestAPIKeyPassedAsArg makes sure that only an API key on the command line
// gives the warning.
func TestAPIKeyPassedAsArg(t *testing.T) {
	cases := []struct {
		args []string
		want bool
	}{
		{[]string{"--serve", "--api-key", "secret"}, true},
		{[]string{"--serve", "--api-key=secret"}, true},
		{[]string{"--serve"}, false},
		{[]string{"--serve", "--address", "127.0.0.1:8080"}, false},
	}
	for _, c := range cases {
		if got := apiKeyPassedAsArg(c.args); got != c.want {
			t.Errorf("apiKeyPassedAsArg(%q) = %v, want %v", c.args, got, c.want)
		}
	}
}
