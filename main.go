package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	router.Use(staticServePage("/", "./public"))

	api := router.Group("/api")

	api.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}

func staticServePage(relativePath string, localFolder string) gin.HandlerFunc {
	return static.Serve(relativePath, static.LocalFile(localFolder, true))
}
