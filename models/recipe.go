package models

type Recipe struct {
	BaseModel

	Name         string       `gorm:"size:255" json:"name"`
	Cuisine      string       `gorm:"size:50" json:"cuisine"`
	Instructions string       `gorm:"type:text" json:"instructions"`
	PhotoURL     string       `json:"photo_url"`
	Ingredients  []Ingredient `gorm:"foreignkey:RecipeID" json:"ingredients"`
}

func GetRecipe(id int) (*Recipe, error) {
	var recipe *Recipe
	err := DB.Preload("Ingredients").First(&recipe, id).Error
	return recipe, err
}

func GetAll() ([]*Recipe, error) {
	var recipes []*Recipe
	err := DB.Preload("Ingredients").Find(&recipes).Error
	return recipes, err
}

func Search(query string) ([]*Recipe, error) {
	var recipes []*Recipe
	err := DB.Preload("Ingredients").
		Joins("JOIN ingredients ON ingredients.recipe_id = recipes.id").
		Where("recipes.name LIKE ? OR recipes.cuisine LIKE ? OR ingredients.name LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Distinct().
		Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}
