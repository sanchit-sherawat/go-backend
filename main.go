package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sanchit-sherawat/go-backend.git/configs"
	"github.com/sanchit-sherawat/go-backend.git/database"
	"github.com/sanchit-sherawat/go-backend.git/routes"
)

func main() {

	// Load configuration
	configs.LoadConfig()

	// Connect to the database
	configs.ConnectDatabase()

	database.Migrate()

	// Initialize Gin router
	router := gin.Default()

	// Setup CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3002", "*"}, // Change this to match your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(configs.ServerPort)
	// r := gin.Default()

	// routes.SetupRoutes(r)

	// r.Run(":8081") // Start the server
}
