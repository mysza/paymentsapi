package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Cache sets max age of cache control header
func Cache(ageInSeconds string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%s", ageInSeconds))
		c.Next()
	}
}
