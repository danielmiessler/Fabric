package restapi

import (
	"errors"
	"log/slog"
	"net/http"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/gin-gonic/gin"
)

// cleanCORSOrigins trims each origin, removes a trailing "/" and removes
// empty entries. A browser sends the Origin header with no "/". It
// refuses "*" when apiKey is empty, because then each web page that
// the user opens can call the server.
func cleanCORSOrigins(origins []string, apiKey string) ([]string, error) {
	var ret []string
	for _, o := range origins {
		if o = strings.TrimSuffix(strings.TrimSpace(o), "/"); o != "" {
			ret = append(ret, o)
		}
	}
	if apiKey == "" && slices.Contains(ret, "*") {
		return nil, errors.New(i18n.T("server_cors_wildcard_requires_api_key"))
	}
	return ret, nil
}

// CORSMiddleware lets browser clients on the listed origins call the
// server. The origin "*" allows each origin: the response copies the
// request Origin, not "*". The middleware never allows the origin
// "null". Register it before APIKeyMiddleware, so a preflight request
// does not get a 401. Serve and ServeOllama do not register it when the
// list is empty.
func CORSMiddleware(origins []string) gin.HandlerFunc {
	wildcard := slices.Contains(origins, "*")
	if wildcard {
		slog.Warn(i18n.T("server_cors_wildcard_warning"))
	}
	return func(c *gin.Context) {
		h := c.Writer.Header()
		h.Add("Vary", "Origin")
		if origin := c.GetHeader("Origin"); origin != "" && origin != "null" && (wildcard || slices.Contains(origins, origin)) {
			h.Set("Access-Control-Allow-Origin", origin)
			h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Content-Type, "+APIKeyHeader)
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
