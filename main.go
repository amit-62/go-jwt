package main

import (
	"os"
	"log"

	routes "github.com/amit/go-jwt/routes"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})

	router.Run(":" + port)

}
