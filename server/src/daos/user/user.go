package user

import (
	"fmt"
	"simple-chat-app/server/src/daos"
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
func FetchAll() *[]models.User {
	db := daos.GetDbConn()
	// var users [3]models.User
	// users[0] = models.User{Email: "john@bar.com", Name: "john"}
	// users[1] = models.User{Email: "joe@bar.com", Name: "joe"}
	// users[2] = models.User{Email: "jane@bar.com", Name: "jane"}
	var users []models.User
	result := db.Find(&models.User{})
	fmt.Println(result.Error)

	return &users
}
