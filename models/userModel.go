package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Watched  []Watched `gorm:"foreignkey:UserID"`
	Eated    []Eated   `gorm:"foreignkey:UserID"`
}
