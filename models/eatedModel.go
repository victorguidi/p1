package models

import (
	"gorm.io/gorm"
)

type Eated struct {
	gorm.Model
	Grade   uint
	UserID  int
	User    User
	FoodsID int
	Foods   Foods
}
