package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	if err := router.Run(":80"); err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
