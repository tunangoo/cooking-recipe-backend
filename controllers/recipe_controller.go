package controllers

import (
	"cooking-recipe-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecipes(c *gin.Context) {
	recipes, err := services.GetAllRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}
