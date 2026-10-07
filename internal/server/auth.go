package restapi

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/gin-gonic/gin"
)

const APIKeyHeader = "X-API-Key"

// minAPIKeyLength is the key length below which the server warns. The
// server has no limit on wrong keys, so a short key is easy to guess.
const minAPIKeyLength = 16

// requireAPIKeyForBind rejects a non-loopback bind address that has no
// API key. An empty or unspecified host binds each interface, and that
// counts as non-loopback. It warns when the key is short.
func requireAPIKeyForBind(address, apiKey string) error {
	if apiKey != "" && len(apiKey) < minAPIKeyLength {
		slog.Warn("API key is short: use a random key of 16 or more characters", "length", len(apiKey))
	}
	if apiKey != "" || isLoopbackHost(hostOnly(address)) {
		return nil
	}
	return fmt.Errorf(i18n.T("server_api_key_required"), address)
}

// APIKeyMiddleware validates API key for protected endpoints.
// Swagger documentation endpoints (/swagger/*) are exempt from authentication
// to allow users to browse and test the API documentation freely.
func APIKeyMiddleware(apiKey string) gin.HandlerFunc {
	// Hash the configured key so the comparison avoids leaking its length through timing.
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
			slog.Warn("API key missing", "client", c.ClientIP(), "path", c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			return
		}

		headerKey := sha256.Sum256([]byte(headerApiKey))
		if subtle.ConstantTimeCompare(headerKey[:], expectedKey[:]) != 1 {
			// An operator can find repeated guesses in this log.
			slog.Warn("API key wrong", "client", c.ClientIP(), "path", c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong API Key"})
			return
		}

		c.Next()
	}
}
