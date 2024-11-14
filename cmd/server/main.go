package main

import (
	"ignbmd/blogging-platform-api/config"
	"ignbmd/blogging-platform-api/internal/app/controllers"
	"ignbmd/blogging-platform-api/internal/app/repositories"
	"ignbmd/blogging-platform-api/internal/app/services"
	"ignbmd/blogging-platform-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Connect to MongoDB
	config.ConnectDB()

	// Set up Gin Router
	router := gin.Default()

	// Initialize repositories, services, and controllers
	postRepo := repositories.NewPostRepository()
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	// Load routes
	routes.RegisterPostRoutes(router, postController)

	// Start the server on specified port
	port := config.GetPort()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
