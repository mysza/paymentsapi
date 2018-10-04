package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// RequestIDKey is a key used to set request id in the context.
const RequestIDKey = "reqID"

var (
	prefix string
	reqID  uint64
)

func init() {
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = "localhost"
	}
	var buf = make([]byte, 16)
	rand.Read(buf)
	uid := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(buf)
	uid = strings.NewReplacer("+", "", "/", "").Replace(uid)
	prefix = fmt.Sprintf("%s/%s", hostname, uid)
}

// RequestID is a middleware that injects a request id into the context of
// each request. A request ID is in a form of "host.example.com/random-00001",
// where "random" is a random base64 string unique to this process, followed by
// atomically incremented request counter.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := atomic.AddUint64(&reqID, 1)
		c.Set(RequestIDKey, fmt.Sprintf("%s-%010d", prefix, requestID))
		c.Next()
	}
}

// GetRequestID is a helper method retrieving the generated request id from the context.
func GetRequestID(c *gin.Context) string {
	if c == nil {
		return ""
	}
	reqID, exists := c.Get(RequestIDKey)
	if exists {
		return reqID.(string)
	}
	return ""
}
