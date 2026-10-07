package restapi

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaxRequestBodyBytes is the maximum size of a request body that the
// server reads. It is large enough for chat requests with images and
// attached documents.
const MaxRequestBodyBytes = 50 << 20 // 50 MiB

// MaxBodyBytesMiddleware stops a request body at limit bytes. When
// Content-Length is more than limit, it answers 413 before the handler
// runs. The middleware reads a body with no length (chunked) before the
// handler runs, and answers 413 if the body is more than limit. The handler
// then reads a second copy of a chunked body, thus a chunked request can use
// two times limit. Without this read, a handler gets the MaxBytesError and
// answers 400, not 413. limit must be more than 0.
func MaxBodyBytesMiddleware(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > limit {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{"error": "request body too large"})
			return
		}
		if c.Request.Body != nil {
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
			if c.Request.ContentLength < 0 {
				body, err := io.ReadAll(c.Request.Body)
				if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
					c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{"error": "request body too large"})
					return
				}
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot read request body"})
					return
				}
				c.Request.Body = io.NopCloser(bytes.NewReader(body))
			}
		}
		c.Next()
	}
}
