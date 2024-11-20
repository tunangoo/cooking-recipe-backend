package services

import (
	"cooking-recipe-backend/models"
)

func GetAll() ([]*models.Recipe, error) {
	recipes, err := models.GetAll()
	if err != nil {
		return nil, err
	}
	return recipes, err
}

func GetRecipe(id int) (*models.Recipe, error) {
	recipe, err := models.GetRecipe(id)
	if err != nil {
		return nil, err
	}
	return recipe, err
}
