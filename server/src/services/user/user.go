package user

import (
	userDao "simple-chat-app/server/src/daos/user"
	"simple-chat-app/server/src/models"
)

/**
Fetch all users.
*/
func FetchAll() *[3]models.User {
	return userDao.FetchAll()
}
