package main

import (
	"net/http"
	"io"
	"log"

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

	r.GET("/check-auth", func(c *gin.Context) {

		resp, err := http.Get("http://auth-service:8081/health")
		if err != nil {
			log.Println("error calling auth-service:", err)
			c.JSON(500, gin.H{"error": "auth-service unreachable"})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		c.JSON(200, gin.H{
			"message": "response from auth-service",
			"data":    string(body),
		})
	})

	r.Run(":8082")
}