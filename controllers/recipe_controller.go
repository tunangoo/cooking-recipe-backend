package controllers

import (
	"cooking-recipe-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	recipes, err := services.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	recipe, err := services.GetRecipe(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func SearchRecipes(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	recipes, err := services.SearchRecipes(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)
}
