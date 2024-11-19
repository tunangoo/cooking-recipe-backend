package main

import (
	"cooking-recipe-backend/config"
	"cooking-recipe-backend/database"
	"cooking-recipe-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadEnv()

	// Connect to database
	database.ConnectDatabase()

	// Setup Gin
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
