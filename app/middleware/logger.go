package middleware

import (
	"awesomeProject/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start) // 延迟
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		logger.Logger.Info(
			"HTTP request",
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("statusCode", statusCode),
			zap.Duration("latency", latency),
		)
	}
}
