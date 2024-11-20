package routes

import (
	"cooking-recipe-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	recipeRoutes := r.Group("/recipe")
	{
		recipeRoutes.GET("/all", controllers.GetAll)
		recipeRoutes.GET("/:id", controllers.GetRecipe)
		recipeRoutes.GET("/search", controllers.SearchRecipes)
	}
}
