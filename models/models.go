package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uint      `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConnectDatabase() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB = database

	DB.AutoMigrate(&Recipe{}, &Ingredient{})

}

// Auto-set timestamps before creating a record
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return
}

// Auto-update timestamp before updating a record
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}
