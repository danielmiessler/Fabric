package restapi

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/gin-gonic/gin"
)

const APIKeyHeader = "X-API-Key"

// requireAPIKeyForBind rejects a non-loopback bind address that has no
// API key. An empty or unspecified host binds each interface, and that
// counts as non-loopback.
func requireAPIKeyForBind(address, apiKey string) error {
	if apiKey != "" {
		return nil
	}
	host := address
	if h, _, err := net.SplitHostPort(address); err == nil {
		host = h
	}
	if host == "localhost" {
		return nil
	}
	if ip := net.ParseIP(host); ip != nil && ip.IsLoopback() {
		return nil
	}
	return fmt.Errorf(i18n.T("server_api_key_required"), address)
}

// APIKeyMiddleware validates API key for protected endpoints.
// Swagger documentation endpoints (/swagger/*) are exempt from authentication
// to allow users to browse and test the API documentation freely.
func APIKeyMiddleware(apiKey string) gin.HandlerFunc {
	// Compare digests, not the raw values. ConstantTimeCompare returns
	// early when the lengths are different, and that shows the length of
	// the configured key.
	expectedKey := sha256.Sum256([]byte(apiKey))
	return func(c *gin.Context) {
		// Skip authentication for Swagger documentation endpoints
		// This allows public access to API docs even when authentication is enabled
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}

		headerApiKey := c.GetHeader(APIKeyHeader)

		if headerApiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			return
		}

		headerKey := sha256.Sum256([]byte(headerApiKey))
		if subtle.ConstantTimeCompare(headerKey[:], expectedKey[:]) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong API Key"})
			return
		}

		c.Next()
	}
}
