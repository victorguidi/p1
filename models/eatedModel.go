package models

import (
	"gorm.io/gorm"
)

type Eated struct {
	gorm.Model
	UserID  int
	User    User
	FoodsID int
	Foods   Foods
	grade   float32 `sql:"type:decimal(10,2);"`
}
