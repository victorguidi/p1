package models

import "gorm.io/gorm"

type ShowWatched struct {
	gorm.Model
	User User
	Show Shows
	rate float32
}
