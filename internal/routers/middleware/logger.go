package middleware

import (
	"ammount-in-words/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// DefaultStructuredLogger logs a gin HTTP request in JSON format. Uses the
// default logger from rs/zerolog.
func DefaultStructuredLogger() gin.HandlerFunc {
	logger := logger.CreateLogger()
	return StructuredLogger(logger)
}

// StructuredLogger logs a gin HTTP request in JSON format. Allows to set the
// logger for testing purposes.
func StructuredLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypeAny).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Log using the params
		logg := logger.Sugar()
		defer logg.Sync()

		params := []interface{}{
			"client_id", param.ClientIP,
			"method", param.Method,
			"status_code", param.StatusCode,
			"body_size", param.BodySize,
			"path", param.Path,
			"latency", param.Latency.String(),
			"message", param.ErrorMessage}

		if c.Writer.Status() >= 500 {
			logg.Errorw("", params...)
		} else {
			logg.Infow("", params...)
		}
	}
}
