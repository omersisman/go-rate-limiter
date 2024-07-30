package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Endpoint1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 1"})
}

func Endpoint2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 2"})
}

func Endpoint3(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 3"})
}
