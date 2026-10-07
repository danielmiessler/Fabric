package template

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"syscall"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/util"
)

// newTextServer starts a test server on a loopback address. The server
// answers each request with body as text/plain.
func newTextServer(body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(body))
	}))
}

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
		{
			name:        "invalid URL",
			operation:   "get",
			value:       "not-a-url",
			wantErr:     true,
			errContains: "unsupported protocol",
		},
		{
			name:        "file URL",
			operation:   "get",
			value:       "file:///tmp/notes.txt",
			wantErr:     true,
			errContains: "unsupported protocol",
		},
		{
			name:        "malformed URL",
			operation:   "get",
			value:       "http://[::1]:namedport",
			wantErr:     true,
			errContains: "error creating request",
		},
		{
			name:        "loopback address",
			operation:   "get",
			server:      func() *httptest.Server { return newTextServer("loopback content") },
			wantErr:     true,
			errContains: i18n.T("util_error_non_public_address"),
		},
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

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchPlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.errContains != "" {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error %q should contain %q", err.Error(), tt.errContains)
					t.Logf("Full error: %v", err)
				}
				return
			}

			if err == nil && got == "" {
				t.Error("FetchPlugin.Apply() returned empty content on success")
			}
		})
	}
}

// allowTestServer lets fetch connect to the loopback address of server.
func allowTestServer(t *testing.T, server *httptest.Server) {
	old := fetchDialControl
	fetchDialControl = func(network, address string, c syscall.RawConn) error {
		if address == server.Listener.Addr().String() {
			return nil
		}
		return util.DenyNonPublicAddress(network, address, c)
	}
	t.Cleanup(func() { fetchDialControl = old })
}

// TestFetchWithTestServerAllowed checks that fetch reads the response when
// the address check accepts the address of the server.
func TestFetchWithTestServerAllowed(t *testing.T) {
	server := newTextServer("ok")
	defer server.Close()
	allowTestServer(t, server)

	got, err := (&FetchPlugin{}).Apply("get", server.URL)
	if err != nil || got != "ok" {
		t.Fatalf("Apply() = %q, %v; want %q", got, err, "ok")
	}
}

// TestFetchContentValidation checks that fetch rejects each response that
// is not small, valid text with status 200.
func TestFetchContentValidation(t *testing.T) {
	big := strings.Repeat("a", MaxContentSize+1)
	tests := []struct {
		name    string
		handler http.HandlerFunc
		wantKey string // i18n key of the error, "" for success
	}{
		{"status 404", func(w http.ResponseWriter, _ *http.Request) {
			http.NotFound(w, nil)
		}, "fetch_http_error"},
		{"Content-Length too large", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte(big))
		}, "fetch_content_too_large"},
		{"streamed body too large", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.(http.Flusher).Flush() // no Content-Length
			_, _ = w.Write([]byte(big))
		}, "fetch_content_exceeds_limit"},
		{"binary content type", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "image/png")
			_, _ = w.Write([]byte("png"))
		}, "fetch_unsupported_content_type"},
		{"no content type", func(w http.ResponseWriter, _ *http.Request) {
			w.Header()["Content-Type"] = nil
			_, _ = w.Write([]byte("text"))
		}, "fetch_unsupported_content_type"},
		{"not UTF-8", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte{0xff, 0xfe})
		}, "fetch_content_not_utf8"},
		{"null byte", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte("a\x00b"))
		}, "fetch_content_null_bytes"},
		{"body shorter than Content-Length", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", "100")
			_, _ = w.Write([]byte("short"))
		}, "fetch_error_reading_response"},
		{"JSON", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write([]byte(`{"a":1}`))
		}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handler)
			defer server.Close()
			allowTestServer(t, server)

			_, err := (&FetchPlugin{}).Apply("get", server.URL)
			if tt.wantKey == "" {
				if err != nil {
					t.Fatalf("Apply() error = %v, want nil", err)
				}
				return
			}
			// The part of the message before the first verb identifies it.
			want, _, _ := strings.Cut(i18n.T(tt.wantKey), "%")
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Fatalf("Apply() error = %v, want %q", err, want)
			}
		})
	}
}
