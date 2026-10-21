package plugins

import (
	"fmt"
	"os"
	"strings"
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

func TestRequireSecureBaseURL(t *testing.T) {
	ok := []string{
		"",
		"   ",
		"https://api.example.com/v1",
		"HTTPS://api.example.com",
		"http://localhost:1234/v1",
		"http://127.0.0.1:11434",
		"http://127.1.2.3",
		"http://[::1]:8080",
		"http://app.localhost:3000",
		"http://localhost.:8080",
		"HTTP://LOCALHOST:1",
		"http://[::ffff:127.0.0.1]:1",
	}
	for _, raw := range ok {
		if err := RequireSecureBaseURL(raw); err != nil {
			t.Errorf("RequireSecureBaseURL(%q) = %v, want nil", raw, err)
		}
	}

	bad := []string{
		"http://api.example.com/v1",
		"http://192.168.1.10:8000/v1",
		"http://localhost.evil.example/v1",
		"ftp://localhost/v1",
		"api.example.com/v1",
		"http://",
		"//api.example.com",
		"http://0.0.0.0:1",
		"http://127.0.0.1.nip.io",
		"http://[::1",
	}
	for _, raw := range bad {
		if err := RequireSecureBaseURL(raw); err == nil {
			t.Errorf("RequireSecureBaseURL(%q) = nil, want an error", raw)
		}
	}
}

func TestRequireSecureBaseURLErrorOmitsCredentials(t *testing.T) {
	err := RequireSecureBaseURL("http://user:secret@api.example.com/v1?key=abc")
	if err == nil {
		t.Fatal("want an error for plaintext http to a remote host")
	}
	if !strings.Contains(err.Error(), "http://api.example.com") {
		t.Errorf("error %q does not name the refused origin", err)
	}
	for _, leak := range []string{"user:secret", "key=abc"} {
		if strings.Contains(err.Error(), leak) {
			t.Errorf("error %q leaks %q", err, leak)
		}
	}
}
