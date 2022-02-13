package daos_user

import (
	"simple-chat-app/server/src/models"
)

/**
Find a user by email.
*/
func FindByEmail(email string) *models.User {
	user := models.User{Email: email, Name: "steve"}
	return &user
}

/**
Fetch all users.
*/
func FetchAll() *[3]models.User {
	var users [3]models.User
	users[0] = models.User{Email: "john@bar.com", Name: "john"}
	users[1] = models.User{Email: "joe@bar.com", Name: "joe"}
	users[2] = models.User{Email: "jane@bar.com", Name: "jane"}
	return &users
}
