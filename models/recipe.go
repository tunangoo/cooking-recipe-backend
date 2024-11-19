package models

type Recipe struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:255"`
	Cuisine      string `gorm:"size:50"`
	Instructions string `gorm:"type:text"`
	PhotoURL     string
}
