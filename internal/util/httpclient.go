package util

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strings"
	"syscall"
	"time"
	"unicode"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// maxRedirects is the maximum number of redirects that a client from
// NewPublicHTTPClient follows.
const maxRedirects = 5

// nonPublicPrefixes are ranges that the netip methods do not find. They are
// not public, or they can send a connection to an IPv4 address in an IPv6
// address.
var nonPublicPrefixes = []netip.Prefix{
	// Shared address space (RFC 6598). Carrier-grade NAT and some VPN products
	// use it, and the Alibaba Cloud metadata service is at 100.100.100.200.
	netip.MustParsePrefix("100.64.0.0/10"),
	netip.MustParsePrefix("0.0.0.0/8"),     // "this network"; a connect to 0.0.0.0 goes to the local host
	netip.MustParsePrefix("192.0.0.0/24"),  // IETF protocol assignments
	netip.MustParsePrefix("198.18.0.0/15"), // benchmark tests
	netip.MustParsePrefix("64:ff9b::/96"),  // NAT64
	netip.MustParsePrefix("2002::/16"),     // 6to4
	netip.MustParsePrefix("::/96"),         // IPv4-compatible IPv6 (::a.b.c.d)
	netip.MustParsePrefix("240.0.0.0/4"),   // reserved, and the broadcast address
	netip.MustParsePrefix("2001::/32"),     // Teredo
	netip.MustParsePrefix("fec0::/10"),     // site-local (deprecated)
}

// nonPublicAddressError is the error of DenyNonPublicAddress. It has no
// address, thus the client of a server does not get the IP address of a
// host on the internal network.
type nonPublicAddressError struct{}

func (nonPublicAddressError) Error() string { return i18n.T("util_error_non_public_address") }

// DenyNonPublicAddress is a Control function for net.Dialer. The dialer
// calls it after DNS resolution, with the IP address and the port of the
// connection. It returns an error if the address is not a public IP
// address: loopback, private, unique local, link-local, multicast,
// unspecified, or in nonPublicPrefixes.
func DenyNonPublicAddress(_, address string, _ syscall.RawConn) error {
	addrPort, err := netip.ParseAddrPort(address)
	// Unmap changes an IPv4-mapped IPv6 address to IPv4, because
	// IsUnspecified does not do this. IsPrivate includes the unique local
	// IPv6 addresses. IsMulticast includes the link-local and the
	// interface-local multicast addresses.
	ip := addrPort.Addr().Unmap()
	if err != nil || ip.IsLoopback() || ip.IsPrivate() ||
		ip.IsLinkLocalUnicast() || ip.IsMulticast() || ip.IsUnspecified() {
		return nonPublicAddressError{}
	}
	for _, prefix := range nonPublicPrefixes {
		if prefix.Contains(ip) {
			return nonPublicAddressError{}
		}
	}
	return nil
}

// NewPublicHTTPClient returns an HTTP client for a URL that a template or a
// request supplies. control is the Control function of the dialer. If
// control is nil, the client uses DenyNonPublicAddress. The dialer calls
// control for each connection, thus also for each redirect. A test gives a
// different function to connect to an httptest server on a loopback
// address. When DenyNonPublicAddress refuses an address, the error of the
// client does not include the IP address.
//
// The client stops a request after timeout. It follows a maximum of
// maxRedirects redirects, and it refuses a redirect to a scheme that is
// not http or https. The transport refuses these schemes on the first
// request. The client does not use a proxy, because control must see the
// address of the server. Each request uses a new connection, thus a client
// for one request keeps no idle connection.
func NewPublicHTTPClient(timeout time.Duration, control func(network, address string, c syscall.RawConn) error) *http.Client {
	if control == nil {
		control = DenyNonPublicAddress
	}
	dialer := &net.Dialer{Control: control}
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				conn, err := dialer.DialContext(ctx, network, address)
				if errors.Is(err, nonPublicAddressError{}) {
					return nil, nonPublicAddressError{}
				}
				return conn, err
			},
			DisableKeepAlives: true,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= maxRedirects {
				return fmt.Errorf(i18n.T("util_error_too_many_redirects"), maxRedirects)
			}
			if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
				return fmt.Errorf(i18n.T("util_error_unsupported_url_scheme"), req.URL.Scheme)
			}
			return nil
		},
	}
}

// IsLoopbackHost tells if host is "localhost", a subdomain of "localhost" or
// a loopback IP address. The host has no port and no IPv6 brackets.
func IsLoopbackHost(host string) bool {
	host = strings.ToLower(strings.TrimSuffix(host, "."))
	if host == "localhost" || strings.HasSuffix(host, ".localhost") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// privateLANSuffixes are DNS suffixes that only ever name hosts on a private
// network: mDNS (.local), tailscale MagicDNS (.ts.net), and the common
// private-use suffixes. A name under one of these cannot resolve to a public
// host, so plaintext http to it carries the same risk profile as a LAN IP.
var privateLANSuffixes = []string{".local", ".ts.net", ".lan", ".internal", ".home.arpa"}

// cgnat is 100.64.0.0/10, which tailscale uses. net.IP.IsPrivate does not
// include it.
var cgnat = netip.MustParsePrefix("100.64.0.0/10")

// isPrivateLANHost tells if host is a private / LAN destination: a loopback,
// RFC1918, CGNAT, link-local or IPv6 ULA address, or a hostname under a
// private-use DNS suffix. It does not resolve names. The host must be
// lowercase, with no port, no IPv6 brackets and no trailing dot.
func isPrivateLANHost(host string) bool {
	if addr, err := netip.ParseAddr(host); err == nil {
		addr = addr.Unmap()
		return addr.IsLoopback() || addr.IsPrivate() || addr.IsLinkLocalUnicast() || cgnat.Contains(addr)
	}
	if IsLoopbackHost(host) {
		return true
	}
	for _, suffix := range privateLANSuffixes {
		if strings.HasSuffix(host, suffix) {
			return true
		}
	}
	return false
}

// HostAllowedForInsecureHTTP tells if an operator-supplied allowlist (the
// FABRIC_ALLOW_INSECURE_LAN value) permits plaintext http to host. The spec
// is a list of entries with comma or whitespace separators. An empty spec
// matches nothing. The host has no port and no IPv6 brackets. Entry forms:
//
//	192.168.20.4      exact IP
//	192.168.20.0/24   CIDR (IPv4 or IPv6)
//	b450.ts.net       exact hostname
//	*.ts.net          hostname suffix (also ".ts.net")
//	lan               any private / LAN / tailnet host
func HostAllowedForInsecureHTTP(host, spec string) bool {
	host = strings.ToLower(strings.TrimSuffix(host, "."))
	if host == "" {
		return false
	}
	// A zone (fe80::1%eth0) names the interface, not the address.
	ipHost, _, _ := strings.Cut(host, "%")
	hostIP := net.ParseIP(ipHost)
	entries := strings.FieldsFunc(strings.ToLower(spec), func(r rune) bool {
		return r == ',' || unicode.IsSpace(r)
	})
	for _, entry := range entries {
		entry = strings.TrimSuffix(strings.Trim(entry, "[]"), ".")
		if entry == "lan" {
			if isPrivateLANHost(host) {
				return true
			}
			continue
		}
		if _, ipnet, err := net.ParseCIDR(entry); err == nil {
			if hostIP != nil && ipnet.Contains(hostIP) {
				return true
			}
			continue
		}
		if entry = strings.TrimPrefix(entry, "*"); strings.HasPrefix(entry, ".") {
			if host == entry[1:] || strings.HasSuffix(host, entry) {
				return true
			}
			continue
		}
		if ip := net.ParseIP(entry); ip != nil {
			if hostIP != nil && ip.Equal(hostIP) {
				return true
			}
			continue
		}
		if host == entry {
			return true
		}
	}
	return false
}
