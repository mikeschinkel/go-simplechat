package auth

import (
	"errors"
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"
)

const (
	pwdVerificationFailedMsg = "password verification failed"
)

/**
Verify user cre
*/
func VerifyUser(
	email string,
	password string,
) (string, error) {
	// Search for the user
	user, err := userDao.FindByEmail(email)
	if err != nil {
		return "", err
	}
	userCreds := authDao.GetUserCreds(user.ID)
	// Check the password
	if password != userCreds.Pwdhash {
		return "", errors.New(pwdVerificationFailedMsg)
	}
	jwt := "ima json web token"
	return jwt, nil
}
