package auth

import (
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"
	jwtUtil "simple-chat-app/server/src/util/jwt"

	"golang.org/x/crypto/bcrypt"
)

/**
Verify user cre
*/
func VerifyUserAndGetToken(
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
	// If password passed create a json web token
	return jwtUtil.Sign(&UserData{user.ID, user.Email, user.Name})
}
