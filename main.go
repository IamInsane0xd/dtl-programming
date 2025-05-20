package DTLProgramming

import (
	"github.com/gin-gonic/gin"
	"log"
)

//goland:noinspection GoUnusedFunction
func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
