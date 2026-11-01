package core

import (
	"bytes"
	"os"
	"strings"
	"testing"

	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai_compatible"
)

func TestConfigureVendorsReportsConfiguredFailures(t *testing.T) {
	for _, tc := range []struct {
		name       string
		key        string
		baseURL    string
		wantWarn   bool
		wantActive bool
	}{
		{"configured insecure endpoint", "synthetic-key", "http://10.1.2.3:4000/v1", true, false},
		{"unconfigured provider", "", "http://10.1.2.3:4000/v1", false, false},
		{"configured secure endpoint", "synthetic-key", "https://example.com/v1", false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("LITELLM_API_KEY", tc.key)
			t.Setenv("LITELLM_API_BASE_URL", tc.baseURL)
			t.Setenv("FABRIC_ALLOW_INSECURE_LAN", "")
			vendor, _ := openai_compatible.CreateClient("LiteLLM")
			registry := &PluginRegistry{
				VendorsAll:    ai.NewVendorsManager(),
				VendorManager: ai.NewVendorsManager(),
			}
			registry.VendorsAll.AddVendors(vendor)

			var output bytes.Buffer
			oldLevel := debuglog.GetLevel()
			debuglog.SetLevel(debuglog.Off)
			debuglog.SetOutput(&output)
			t.Cleanup(func() {
				debuglog.SetLevel(oldLevel)
				debuglog.SetOutput(os.Stderr)
			})

			registry.ConfigureVendors()

			out := output.String()
			if got := strings.Contains(out, "LiteLLM"); got != tc.wantWarn {
				t.Errorf("warning = %q, want warning %v", out, tc.wantWarn)
			}
			if tc.wantWarn && !strings.Contains(out, tc.baseURL[:strings.LastIndex(tc.baseURL, "/")]) {
				t.Errorf("warning does not explain which endpoint failed: %q", out)
			}
			if strings.Contains(out, "synthetic-key") {
				t.Errorf("warning contains the API key: %q", out)
			}
			if got := registry.VendorManager.FindByName("LiteLLM") != nil; got != tc.wantActive {
				t.Errorf("vendor active = %v, want %v", got, tc.wantActive)
			}
		})
	}
}
