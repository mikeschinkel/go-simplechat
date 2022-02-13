package user

import (
	userDao "simple-chat-app/server/src/daos/user"
	"simple-chat-app/server/src/models"
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
func AddOne(newUser *models.User) error {
	return userDao.AddOne(newUser)
}
