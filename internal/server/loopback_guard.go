package restapi

import (
	"log/slog"
	"mime"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// hostOnly removes the port from a Host value or from an address. It
// accepts an IPv6 literal such as "[::1]:8080". If the value has no port,
// hostOnly removes only the brackets of an IPv6 literal.
func hostOnly(hostport string) string {
	if h, _, err := net.SplitHostPort(hostport); err == nil {
		return h
	}
	return strings.Trim(hostport, "[]")
}

// isLoopbackHost tells if host is "localhost" or a loopback IP address.
func isLoopbackHost(host string) bool {
	if host == "localhost" {
		return true
	}
	if ip := net.ParseIP(host); ip != nil {
		return ip.IsLoopback()
	}
	return false
}

// originHost gives the host name of an Origin header value, with no port.
func originHost(origin string) (string, bool) {
	u, err := url.Parse(origin)
	if err != nil || u.Host == "" {
		return "", false
	}
	return hostOnly(u.Host), true
}

// LoopbackSecurityMiddleware examines the requests to a server that has no
// API key. Serve and ServeOllama use it only when there is no key, because
// then the server does no authentication. corsOrigins is the list from
// cleanCORSOrigins.
//
// The middleware does two checks:
//
//   - The Host header must be "localhost", a loopback IP address or the host
//     of a CORS origin. The host name of a web page can resolve to 127.0.0.1.
//     The browser then sends that name in the Host header, and the
//     middleware rejects the request.
//   - A POST, PUT, DELETE or PATCH request must not have an Origin header
//     from a different origin. A browser sends some of these requests to a
//     different origin with no CORS preflight, for example a text/plain
//     POST. A page on a different loopback port is a different origin. The
//     middleware accepts a request with no Origin header, for example from
//     curl or the CLI.
//
// A server with an API key does not use this middleware. A browser must do
// a CORS preflight before it sends the X-API-Key header to a different
// origin.
func LoopbackSecurityMiddleware(corsOrigins []string) gin.HandlerFunc {
	allowedHosts := map[string]bool{}
	allowedOrigins := map[string]bool{}
	for _, o := range corsOrigins {
		allowedOrigins[o] = true
		if h, ok := originHost(o); ok {
			allowedHosts[h] = true
		}
	}

	return func(c *gin.Context) {
		// A preflight request and the Swagger documents do not change data.
		// The CORS and Swagger handlers answer them.
		if c.Request.Method == http.MethodOptions ||
			strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}

		// The Host must be a loopback host or the host of a CORS origin.
		host := hostOnly(c.Request.Host)
		if !isLoopbackHost(host) && !allowedHosts[host] {
			slog.Warn("request blocked: unexpected Host", "client", c.ClientIP(), "host", c.Request.Host, "path", c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: unexpected Host header"})
			return
		}

		// A request that can change data must not come from a different
		// origin. The middleware accepts the same origin, a CORS origin, the
		// web UI dev server and a request with no Origin header (curl, the
		// CLI).
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
			if origin := c.GetHeader("Origin"); origin != "" {
				if !originAllowed(origin, c.Request.Host, allowedOrigins, c.ContentType()) {
					slog.Warn("request blocked: cross-site Origin", "client", c.ClientIP(), "origin", origin, "path", c.Request.URL.Path)
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: cross-site request blocked"})
					return
				}
			}
		}

		c.Next()
	}
}

// webUIOrigins are the origins of the web UI (web/): the dev server on port
// 5173 and the preview server on port 4173. The Vite proxy sends the Origin
// of the browser page to the server. Each project that uses Vite has the
// same default ports. Thus a page from a different
// project can have this origin, and originAllowed accepts it only for a JSON
// request (see originAllowed).
var webUIOrigins = map[string]bool{
	"http://localhost:5173": true,
	"http://127.0.0.1:5173": true,
	"http://localhost:4173": true,
	"http://127.0.0.1:4173": true,
}

// originAllowed tells if a request with this Origin can change data on a
// server with no API key. requestHost is the Host header of the request.
// originAllowed accepts a CORS origin, and an origin with the same host and
// port as requestHost. It accepts the web UI dev server only when contentType
// is application/json. The web UI sends JSON. A browser does a CORS preflight
// before it sends JSON to a different origin, and this server does not allow
// the web UI origin in a preflight. originAllowed rejects the origin "null".
func originAllowed(origin, requestHost string, allowedOrigins map[string]bool, contentType string) bool {
	if origin == "null" {
		return false
	}
	if allowedOrigins[origin] || (webUIOrigins[origin] && contentType == "application/json") {
		return true
	}
	u, err := url.Parse(origin)
	return err == nil && u.Host != "" && strings.EqualFold(u.Host, requestHost)
}

// requireJSON answers 415 when the Content-Type of a request is not
// application/json. Put it in front of each handler that reads a JSON body.
// gin reads a JSON body with any Content-Type, and a browser sends a
// text/plain POST to a different site with no CORS preflight. The save
// routes, such as POST /patterns/:name, take a raw body and do not use it.
func requireJSON(c *gin.Context) {
	if mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type")); mediaType != "application/json" {
		c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
	}
}
