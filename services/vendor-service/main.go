package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "vendor-service",
			"status":  "running",
		})
	})

	r.GET("/vendors", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"vendors": []string{
				"Wedding Photographer",
				"Cake Designer",
				"DJ Service",
			},
		})
	})

	r.Run(":8082")
}