package models

import (
	"gorm.io/gorm"
)

type UserCreds struct {
	gorm.Model
	Pwdhash string
	UserID  int
	User    User
}

// Also put things like Signup Request, Reset Password Request
