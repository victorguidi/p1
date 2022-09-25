package models

import "gorm.io/gorm"

type Foods struct {
	gorm.Model
	Name    string
	Country string
	Type    string
}
