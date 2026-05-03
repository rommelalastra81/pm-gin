package main

import (
	"pm-gin/config"
	"pm-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()

	// Auto-migrate the Users table
	//config.DB.AutoMigrate(&models.Users{})

	// Initialize router
	router := gin.Default()

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register auth routes
	routes.RegisterAuthRoutes(router)
	routes.UserRoutes(router)

	router.Run() // listens on 0.0.0.0:8080 by default
}
