package services

import (
	"simple-chat-app/server/src/daos"
	"simple-chat-app/server/src/models"

	"golang.org/x/crypto/bcrypt"
)

/**
Fetch all users.
*/
func FetchAllUsers() (*[]models.User, error) {
	return daos.FetchAllUsers()
}

/**
Add a new user object.
*/
func AddUser(email string, name string, password string) error {
	// Save the user
	user, err := daos.AddUser(email, name)
	if err != nil {
		return err
	}
	// Ecrypt password and save it in user_creds table. Note bcrypt using byte[] not strings.
	pwdHash, errr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errr != nil {
		return err
	}
	err = daos.SaveUserCreds(user.ID, pwdHash)
	if err != nil {
		return err
	}
	return nil
}

/**
Update user's email and name.
*/
func UpdateUser(id uint, email string, name string) error {
	user, err := daos.FindUserById(id)
	if err != nil {
		return err
	}
	daos.UpdateUser(user, email, name)
	return nil
}

/**
Delete one user
*/
func DeleteUser(id uint) error {
	return daos.DeleteUser(id)
}
