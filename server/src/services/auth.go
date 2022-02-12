package services

import (
	"errors"
	"koa-react-chat-app/server/src/daos"
)

type AuthService struct{}

/**
Verify user cre
*/
func (as *AuthService) VerifyUser(
	email string,
	password string,
) (string, error) {
	userDao := daos.UserDao{}
	authDao := daos.AuthDao{}
	// Search for the user
	user := userDao.FindByEmail(email)
	userCreds := authDao.GetUserCreds(user.ID)
	// Check the password
	if password != userCreds.Pwdhash {
		return "", errors.New("password verification failed")
	}
	jwt := "ima json web token"
	return jwt, nil
}
