package services

import (
	"simple-chat-app/server/src/daos"
	"simple-chat-app/server/src/models"

	"golang.org/x/crypto/bcrypt"
)

/**
Verify user cre
*/
func VerifyAndFetchUser(email string, password string) (*models.User, error) {
	// Search for the user
	user, err := daos.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// Fetch the pwd hash
	pwdHash, err := daos.GetPwdHash(user.ID)
	if err != nil {
		return nil, err
	}
	// Compare the password to the hash
	err = bcrypt.CompareHashAndPassword(pwdHash, []byte(password))
	if err != nil {
		return nil, err
	}
	// Return
	return user, nil
}
