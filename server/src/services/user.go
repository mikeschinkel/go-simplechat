package services

import (
	"simple-chat-app/server/src/dal"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/util"
)

/**
Fetch all users.
*/
func FetchAllUsers() (*[]models.User, error) {
	return dal.FetchAllUsers()
}

/**
Add a new user object.
*/
func AddUser(email string, name string, password string) error {
	// Save the user
	user, err := dal.AddUser(email, name)
	if err != nil {
		return err
	}
	// Ecrypt password and save it in user_creds table.
	pwdHash, err := util.HashPwd(password)
	if err != nil {
		return err
	}
	err = dal.SaveUserCreds(user.ID, pwdHash)
	if err != nil {
		return err
	}
	return nil
}

/**
Update user's email and name.
*/
func UpdateUser(id uint, email string, name string) error {
	user, err := dal.FindUserById(id)
	if err != nil {
		return err
	}
	dal.UpdateUser(user, email, name)
	return nil
}

/**
Delete one user
*/
func DeleteUser(id uint) error {
	return dal.DeleteUser(id)
}
