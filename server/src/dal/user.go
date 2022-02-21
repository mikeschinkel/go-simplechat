package dal

import (
	"simple-chat-app/server/src/models"
)

/**
Find a user by their id.
*/
func FindUserById(id uint) (*models.User, error) {
	user := models.User{}
	resp := db.First(&user, id)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

/**
Find a user by email.
*/
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	resp := db.Where("email = ?", email).First(&user)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

/**
Fetch all users.
*/
func FetchAllUsers() (*[]models.User, error) {
	var users []models.User
	resp := db.Find(&users)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &users, nil
}

/**
Add a new user.
*/
func AddUser(email string, name string) (*models.User, error) {
	newUser := models.User{Email: email, Name: name}
	resp := db.Save(&newUser)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &newUser, nil
}

/**
Update user's email and name.
*/
func UpdateUser(user *models.User, email string, name string) {
	db.Model(user).Updates(models.User{Email: email, Name: name})
}

/**
Delete one user.
*/
func DeleteUser(id uint) error {
	resp := db.Unscoped().Where("id = ?", id).Delete(&models.User{})
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

/**
Fetch a user's hashed password
*/
func GetPwdHash(userId uint) (string, error) {
	var userCreds models.UserCreds
	resp := db.Where("user_id = ?", userId).First(&userCreds)
	if resp.Error != nil {
		return "", resp.Error
	}
	return userCreds.Pwdhash, nil
}

/**
Create a user credentials row to store confidentials stuff.
*/
func SaveUserCreds(id uint, pwdHash string) error {
	userCreds := models.UserCreds{Pwdhash: pwdHash, UserID: id}
	resp := db.Save(&userCreds)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
