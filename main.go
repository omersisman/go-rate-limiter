package main

import (
	"github.com/gin-gonic/gin"
	"go-rate-limiter/config"
	"go-rate-limiter/handler"
	"go-rate-limiter/middleware"
	"go-rate-limiter/redis"
)

func main() {
	r := gin.Default()

	config.LoadConfig("./config/config.yaml")
	redis.InitRedis()

	r.GET("/endpoint1", middleware.RateLimitMiddleware("endpoint1"), handler.Endpoint1)
	r.GET("/endpoint2", middleware.RateLimitMiddleware("endpoint2"), handler.Endpoint2)
	r.GET("/endpoint3", middleware.RateLimitMiddleware("endpoint3"), handler.Endpoint3)

	r.Run(":8080")
}
