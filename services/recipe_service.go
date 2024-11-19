package services

import (
	"cooking-recipe-backend/database"
	"cooking-recipe-backend/models"
)

func GetAllRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	if err := database.DB.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}
