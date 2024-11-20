package main

import (
	"cooking-recipe-backend/config"
	"cooking-recipe-backend/models"
	"cooking-recipe-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadEnv()

	// Connect to database
	models.ConnectDatabase()

	// Setup Gin
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
