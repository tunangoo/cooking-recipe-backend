package models

type Ingredient struct {
	BaseModel

	RecipeID uint   `gorm:"not null" json:"recipe_id"`
	Name     string `gorm:"size:255" json:"name"`
	Quantity string `gorm:"size:255" json:"quantity"`
}
