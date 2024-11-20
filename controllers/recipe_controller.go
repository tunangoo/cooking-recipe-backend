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
	}

	recipe, err := services.GetRecipe(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
	}
	c.JSON(http.StatusOK, recipe)
}
