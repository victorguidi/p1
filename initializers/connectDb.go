package initializers

import (
	"os"

	"github.com/victorguidi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Yo, database is screeeewd")
	}

	DB = db

	DB.AutoMigrate(&models.User{})

}
