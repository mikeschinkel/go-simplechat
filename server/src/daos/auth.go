package daos

import (
	"koa-react-chat-app/server/src/models"
)

type AuthDao struct{}

/**
Fetch a user's login credentials
*/
func (ad *AuthDao) GetUserCreds(userId uint) *models.UserCreds {
	userCreds := models.UserCreds{Pwdhash: "Password@1"}
	return &userCreds
}
