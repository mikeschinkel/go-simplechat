package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique"`
	Name  string `json:"name"`
}
