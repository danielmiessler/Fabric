package util

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"syscall"
	"time"

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
