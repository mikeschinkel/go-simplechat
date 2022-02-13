package auth

import (
	"simple-chat-app/server/src/models"
)

/**
Fetch a user's login credentials
*/
func GetUserCreds(userId uint) *models.UserCreds {
	userCreds := models.UserCreds{Pwdhash: "Password@1"}
	return &userCreds
}
