package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func WithRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("X-Request-ID", generateRequestID())
		c.Next()
	}
}

func generateRequestID() string {
	return uuid.New().String()
}
