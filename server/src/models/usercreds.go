package models

import (
	"gorm.io/gorm"
)

type UserCreds struct {
	gorm.Model
	Pwdhash []byte `gorm:"size:255;not null"`
	UserID  uint
}
