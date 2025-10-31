package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware es un middleware personalizado para logging
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Procesar request
		c.Next()

		// Calcular latencia
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		// Log
		log.Printf("[%s] %s %s - Status: %d - Latency: %v",
			method,
			path,
			c.ClientIP(),
			statusCode,
			latency,
		)
	}
}
