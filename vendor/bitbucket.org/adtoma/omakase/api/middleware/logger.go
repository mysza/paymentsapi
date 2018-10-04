package middleware

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Logger(appName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		logrus.WithFields(logrus.Fields{
			"reqID":     GetRequestID(c),
			"path":      path,
			"method":    c.Request.Method,
			"clientIP":  c.ClientIP(),
			"app":       appName,
			"userAgent": c.Request.UserAgent(),
			"referer":   c.Request.Referer(),
			"startTime": start,
		}).Info("request_start")
		// call the handler here
		c.Next()
		// log the handle result
		timeTaken := time.Since(start)
		end := time.Now()
		entry := logrus.WithFields(logrus.Fields{
			"reqID":      GetRequestID(c),
			"path":       path,
			"method":     c.Request.Method,
			"clientIP":   c.ClientIP(),
			"app":        appName,
			"userAgent":  c.Request.UserAgent(),
			"referer":    c.Request.Referer(),
			"status":     c.Writer.Status(),
			"statusText": http.StatusText(c.Writer.Status()),
			"latency":    timeTaken,
			"end":        end,
		})
		if len(c.Errors) > 0 {
			// there were errors, log them out
			entry.WithField("errors", c.Errors)
		}
		entry.Info("request_end")
	}
}
