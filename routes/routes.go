package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchit-sherawat/go-backend.git/controllers"
	"github.com/sanchit-sherawat/go-backend.git/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
	}
	// Public routes
	router.POST("/login", controllers.Login)
	router.POST("/adduser", controllers.Register)

	router.GET("/verify-token", controllers.VerifyToken)

	// Protected routes
	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
	}
}
