package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const APIKeyHeader = "X-API-Key"

func APIKeyMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerApiKey := c.GetHeader(APIKeyHeader)

		if headerApiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			return
		}

		if headerApiKey != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong API Key"})
			return
		}

		c.Next()
	}
}
