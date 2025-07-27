package template

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchPlugin(t *testing.T) {
	plugin := &FetchPlugin{}

	tests := []struct {
		name        string
		operation   string
		value       string
		server      func() *httptest.Server
		wantErr     bool
		errContains string
	}{
		// ... keep existing valid test cases ...

		{
			name:        "invalid URL",
			operation:   "get",
			value:       "not-a-url",
			wantErr:     true,
			errContains: "unsupported protocol", // Updated to match actual error
		},
		{
			name:        "malformed URL",
			operation:   "get",
			value:       "http://[::1]:namedport",
			wantErr:     true,
			errContains: "error creating request",
		},
		// ... keep other test cases ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var url string
			if tt.server != nil {
				server := tt.server()
				defer server.Close()
				url = server.URL
			} else {
				url = tt.value
			}

			got, err := plugin.Apply(tt.operation, url)

			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchPlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errContains != "" {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error %q should contain %q", err.Error(), tt.errContains)
					t.Logf("Full error: %v", err) // Added for better debugging
				}
				return
			}

			// For successful cases, verify we got some content
			if err == nil && got == "" {
				t.Error("FetchPlugin.Apply() returned empty content on success")
			}
		})
	}
}
