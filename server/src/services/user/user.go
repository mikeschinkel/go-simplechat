package user

import (
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"
	"simple-chat-app/server/src/models"

	"golang.org/x/crypto/bcrypt"
)

/**
Fetch all users.
*/
func FetchAll() *[]models.User {
	return userDao.FetchAll()
}

/**
Add a new user object.
*/
func AddOne(email string, name string, password string) error {
	// Save the user
	user, err := userDao.AddOne(email, name)
	if err != nil {
		return err
	}
	// Ecrypt password and save it in user_creds table.
	// Note bcrypt using byte[] not strings.
	pwdArr := []byte(password)
	hashedPassword, errr := bcrypt.GenerateFromPassword(pwdArr, bcrypt.DefaultCost)
	if errr != nil {
		return err
	}
	pwdHash := string(hashedPassword)
	err = authDao.SaveUserCreds(user.ID, pwdHash)
	if err != nil {
		return err
	}
	return nil
}
