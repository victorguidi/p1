package models

import (
	"gorm.io/gorm"
)

type Watched struct {
	gorm.Model
	Grade   uint
	UserID  int `sql:"index"`
	User    User
	ShowsID int `sql:"index"`
	Shows   Shows
}
