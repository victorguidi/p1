package models

import "gorm.io/gorm"

type Shows struct {
	gorm.Model
	Name  string
	Type  string
	Class string
}
