package DTLProgramming

import (
	"github.com/gin-gonic/gin"
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
}
