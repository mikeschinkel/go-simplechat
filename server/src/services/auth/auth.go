package auth

import (
	"errors"
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"
)

/**
Verify user cre
*/
func VerifyUser(
	email string,
	password string,
) (string, error) {
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
