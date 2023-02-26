package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	// Enable CORS for all origins
	router.Use(cors.New(config))

	webPort := os.Getenv("PORT")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hi this is modulity response!")
	})

	router.Run(":" + webPort)
}
