package exolab

import (
	"fmt"
	"os"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
)

func TestMain(m *testing.M) {
	if _, err := i18n.Init("en"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to init i18n: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestConfigureRequiresHTTPSForRemoteBaseURL(t *testing.T) {
	for baseURL, wantErr := range map[string]bool{
		"http://localhost:52415/v1": false,
		"http://api.example.com/v1": true,
	} {
		client := NewClient()
		client.ApiBaseURL.Value = baseURL
		if err := client.configure(); (err != nil) != wantErr {
			t.Errorf("configure() with %s: error = %v, want error %t", baseURL, err, wantErr)
		}
	}
}
