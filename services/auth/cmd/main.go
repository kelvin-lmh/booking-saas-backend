package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "registered"})
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "logged in"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
