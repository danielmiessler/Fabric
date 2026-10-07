package util

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
)

func TestDenyNonPublicAddress(t *testing.T) {
	refused := []string{
		"127.0.0.1:80", "[::1]:80", // loopback
		"10.0.0.1:80", "172.16.0.1:80", "192.168.1.1:80", // private
		"[fd00::1]:80",                                              // unique local
		"100.64.0.1:80", "100.100.100.200:80", "100.127.255.255:80", // shared
		"[::ffff:100.100.100.200]:80",
		"169.254.169.254:80", "[fe80::1%en0]:80", // link-local unicast
		"224.0.0.1:80", "[ff02::1]:80", // link-local multicast
		"[ff01::1]:80",                 // interface-local multicast
		"239.1.2.3:80", "[ff0e::1]:80", // multicast
		"0.0.0.0:80", "[::]:80", // unspecified
		"[::ffff:127.0.0.1]:80", "[::ffff:0.0.0.0]:80", // IPv4-mapped IPv6
		"0.1.2.3:80", "192.0.0.1:80", "198.18.0.1:80", "198.19.255.255:80", // special use
		"[64:ff9b::a9fe:a9fe]:80", "[2002:a9fe:a9fe::1]:80", "[::169.254.169.254]:80", // IPv4 in IPv6
		"240.0.0.1:80", "255.255.255.255:80", // reserved and broadcast
		"[2001::1]:80", "[fec0::1]:80", // Teredo and site-local
		"example.com:80", "8.8.8.8", // not an IP address and a port
	}
	for _, address := range refused {
		if DenyNonPublicAddress("tcp", address, nil) == nil {
			t.Errorf("%s: no error, want an error", address)
		}
	}

	for _, address := range []string{
		"8.8.8.8:443", "[2606:4700:4700::1111]:443", "[::ffff:8.8.8.8]:443",
		"100.63.255.255:443", "100.128.0.1:443", // next to the shared range
		"223.255.255.255:443", "[2001:4860:4860::8888]:443", // next to the reserved and Teredo ranges
		"1.0.0.1:443", "192.0.1.1:443", "198.17.255.255:443", "198.20.0.1:443", // next to the special-use ranges
	} {
		if err := DenyNonPublicAddress("tcp", address, nil); err != nil {
			t.Errorf("%s: %v, want no error", address, err)
		}
	}
}

// TestPublicHTTPClientRedirects checks that the client refuses a redirect to
// a private address and a redirect to a scheme that is not http or https.
// It also checks that the client stops after maxRedirects redirects.
func TestPublicHTTPClientRedirects(t *testing.T) {
	tests := map[string]struct{ location, want string }{
		"/private": {"http://10.0.0.1/", i18n.T("util_error_non_public_address")},
		"/ftp":     {"ftp://example.com/", fmt.Sprintf(i18n.T("util_error_unsupported_url_scheme"), "ftp")},
		"/loop":    {"/loop", fmt.Sprintf(i18n.T("util_error_too_many_redirects"), maxRedirects)},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, tests[r.URL.Path].location, http.StatusFound)
	}))
	defer server.Close()

	// Accept the loopback address of the test server. Check all other
	// addresses.
	control := func(network, address string, c syscall.RawConn) error {
		if address == server.Listener.Addr().String() {
			return nil
		}
		return DenyNonPublicAddress(network, address, c)
	}
	client := NewPublicHTTPClient(5*time.Second, control)

	for path, tt := range tests {
		resp, err := client.Get(server.URL + path)
		if err == nil {
			resp.Body.Close()
		}
		if err == nil || !strings.Contains(err.Error(), tt.want) {
			t.Errorf("%s: got error %v, want an error with %q", path, err, tt.want)
		}
	}
}

// TestPublicHTTPClientTimeout checks that the client stops a request to a
// slow server after the timeout.
func TestPublicHTTPClientTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Wait until the client closes the connection. If the client does
		// not stop the request, send the response after 5 seconds.
		select {
		case <-r.Context().Done():
		case <-time.After(5 * time.Second):
		}
	}))
	defer server.Close()

	resp, err := NewPublicHTTPClient(100*time.Millisecond, allowAnyAddress).Get(server.URL)
	if err == nil {
		resp.Body.Close()
	}
	var urlErr *url.Error
	if !errors.As(err, &urlErr) || !urlErr.Timeout() {
		t.Fatalf("got error %v, want a timeout error", err)
	}
}

func allowAnyAddress(string, string, syscall.RawConn) error { return nil }

// TestPublicHTTPClientNilControl checks that a nil control refuses a
// loopback address, and that the error does not show the IP address.
func TestPublicHTTPClientNilControl(t *testing.T) {
	var called bool
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { called = true }))
	defer server.Close()

	_, port, _ := strings.Cut(server.Listener.Addr().String(), ":")
	resp, err := NewPublicHTTPClient(5*time.Second, nil).Get("http://localhost:" + port)
	if err == nil {
		resp.Body.Close()
	}
	if !errors.Is(err, nonPublicAddressError{}) || called {
		t.Fatalf("got error %v, called %v; want the non-public address error and no request", err, called)
	}
	if strings.Contains(err.Error(), "127.0.0.1") || strings.Contains(err.Error(), "::1") {
		t.Fatalf("error %q shows the IP address", err)
	}
}
