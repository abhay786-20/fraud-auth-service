package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
)

// Logger returns a gin middleware that logs HTTP requests.
func Logger(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Log after request is processed
		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Info(method + " " + path + " " + clientIP + " " +
			string(rune(status)) + " " + latency.String())
	}
}
