package models

import (
	"gorm.io/gorm"
)

type Watched struct {
	gorm.Model
	UserID  int
	User    User
	ShowsID int
	Shows   Shows
	grade   float32 `sql:"type:decimal(10,2);"`
}
