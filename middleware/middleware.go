package middleware

import (
	"github.com/gin-gonic/gin"
	"go-rate-limiter/config"
	"go-rate-limiter/rate_limiter"
	"net/http"
	"time"
)

func RateLimitMiddleware(endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rateLimit, exists := config.GetRateLimit(endpoint)
		if !exists {
			c.Next()
			return
		}

		clientIP := c.GetHeader("X-Forwarded-For")
		key := endpoint + ":" + clientIP

		if !rate_limiter.RateLimit(key, rateLimit.Limit, time.Duration(rateLimit.Duration)*time.Second) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
