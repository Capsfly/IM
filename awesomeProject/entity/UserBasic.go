package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID      string `gorm:"uniqueIndex;size:15"`
	Name     string
	Password string
}
