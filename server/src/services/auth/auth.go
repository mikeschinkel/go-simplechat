package auth

import (
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"

	"golang.org/x/crypto/bcrypt"
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
	// Fetch the pwd hash
	pwdHash, err := authDao.GetPwdHash(user.ID)
	if err != nil {
		return "", err
	}
	// Compare the password to the hash
	err = bcrypt.CompareHashAndPassword(pwdHash, []byte(password))
	if err != nil {
		return "", err
	}
	// Generate a jsonwebtoken if passed
	jwt := "ima json web token"
	return jwt, nil
}
