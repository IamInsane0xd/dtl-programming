package DTLProgramming

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"DTLProgramming/dbquery"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Unable to find .env file in go project root, using system env values.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	err := dbquery.ConnectToDb(dbUser, dbPass, dbAddr, dbName)

	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	router := gin.Default()

	api := router.Group("/api")

	apiAuth := api.Group("/auth")
}
