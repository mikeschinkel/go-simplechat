package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique"`
	Name  string `json:"name"`
}

type UserCreds struct {
	gorm.Model
	Pwdhash string
	UserID  int
	User    User `gorm:"constraint:OnDelete:CASCADE;"`
}
