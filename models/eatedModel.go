package models

import "gorm.io/gorm"

type FoodEated struct {
	gorm.Model
	User User
	Food Foods
	rate float32
}
