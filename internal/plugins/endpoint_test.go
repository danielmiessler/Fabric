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

func TestRequireSecureBaseURLAllowInsecureLAN(t *testing.T) {
	// Off by default: a LAN http URL must still be rejected.
	if err := RequireSecureBaseURL("http://192.168.20.4:11440/v1"); err == nil {
		t.Error("LAN http must be rejected when FABRIC_ALLOW_INSECURE_LAN is unset")
	}

	// An allowlist only relaxes the hosts it names.
	t.Setenv("FABRIC_ALLOW_INSECURE_LAN", "192.168.20.0/24, *.ts.net")
	scopedOK := []string{
		"http://192.168.20.4:11440/v1",           // in 192.168.20.0/24
		"http://192.168.20.250/v1",               // in 192.168.20.0/24
		"http://b450.mink-spica.ts.net:11440/v1", // *.ts.net
	}
	for _, raw := range scopedOK {
		if err := RequireSecureBaseURL(raw); err != nil {
			t.Errorf("allowlist %q: RequireSecureBaseURL(%q) = %v, want nil", "192.168.20.0/24,*.ts.net", raw, err)
		}
	}
	scopedBad := []string{
		"http://192.168.21.4/v1", // different /24, not listed
		"http://10.0.0.5/v1",     // not listed
		"http://api.openai.com/v1",
		"http://8.8.8.8/v1",
	}
	for _, raw := range scopedBad {
		if err := RequireSecureBaseURL(raw); err == nil {
			t.Errorf("allowlist %q: RequireSecureBaseURL(%q) = nil, want an error", "192.168.20.0/24,*.ts.net", raw)
		}
	}

	// CIDR and exact-IP entries.
	t.Setenv("FABRIC_ALLOW_INSECURE_LAN", "100.64.0.0/10,10.1.2.3")
	for _, raw := range []string{"http://100.84.110.2:11440/v1", "http://10.1.2.3/v1"} {
		if err := RequireSecureBaseURL(raw); err != nil {
			t.Errorf("cidr/ip allowlist: RequireSecureBaseURL(%q) = %v, want nil", raw, err)
		}
	}
	if err := RequireSecureBaseURL("http://10.1.2.4/v1"); err == nil {
		t.Error("cidr/ip allowlist: 10.1.2.4 is not listed, want an error")
	}

	// The "lan" keyword relaxes any private/LAN/tailnet host, but never public.
	t.Setenv("FABRIC_ALLOW_INSECURE_LAN", "lan")
	for _, raw := range []string{"http://192.168.1.10/v1", "http://nas.local:8080", "http://box.internal/v1"} {
		if err := RequireSecureBaseURL(raw); err != nil {
			t.Errorf("lan keyword: RequireSecureBaseURL(%q) = %v, want nil", raw, err)
		}
	}
	if err := RequireSecureBaseURL("http://api.openai.com/v1"); err == nil {
		t.Error("lan keyword must still reject a public host")
	}

	// IPv6 zones, bracketed entries and trailing dots, as "lan" already allows.
	t.Setenv("FABRIC_ALLOW_INSECURE_LAN", "fe80::1, fd00:1::/32, [fd00::5], nas.lan.")
	for _, raw := range []string{
		"http://[fe80::1%25eth0]:11434",   // exact IP, zoned host
		"http://[fd00:1::7%25eth0]:11434", // CIDR, zoned host
		"http://[fd00::5]:11434",          // bracketed entry
		"http://nas.lan:11434",            // trailing-dot entry
	} {
		if err := RequireSecureBaseURL(raw); err != nil {
			t.Errorf("normalized allowlist: RequireSecureBaseURL(%q) = %v, want nil", raw, err)
		}
	}
	for _, raw := range []string{
		"http://[fe80::2%25eth0]:11434",
		"http://[fd00:2::7%25eth0]:11434",
		"http://[2001:db8::1%25eth0]:11434",
		"http://[fd00::6]:11434",
		"http://nas2.lan:11434",
	} {
		if err := RequireSecureBaseURL(raw); err == nil {
			t.Errorf("normalized allowlist: RequireSecureBaseURL(%q) = nil, want an error", raw)
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
