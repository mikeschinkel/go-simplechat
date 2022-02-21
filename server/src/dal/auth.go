package dal

import (
	"simple-chat-app/server/src/models"
)

/**
Fetch a user's login credentials
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
Create a user credentials table to store confidentials stuff.
*/
func SaveUserCreds(id uint, pwdHash string) error {
	userCreds := models.UserCreds{Pwdhash: pwdHash, UserID: id}
	resp := db.Save(&userCreds)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
