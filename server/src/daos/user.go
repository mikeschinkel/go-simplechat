package daos

import (
	"koa-react-chat-app/server/src/models"
)

type UserDao struct{}

/**
Find a user by email.
*/
func (ud *UserDao) FindByEmail(email string) *models.User {
	user := models.User{Email: email, Name: "steve"}
	return &user
}

/**
Fetch all users.
*/
func (ud *UserDao) FetchAll() *[3]models.User {

	var users [3]models.User
	users[0] = models.User{Email: "john@bar.com", Name: "john"}
	users[1] = models.User{Email: "joe@bar.com", Name: "joe"}
	users[2] = models.User{Email: "jane@bar.com", Name: "jane"}
	return &users
}
