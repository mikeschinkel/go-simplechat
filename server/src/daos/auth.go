package daos

import (
	"simple-chat-app/server/src/models"
)

/**
Fetch a user's login credentials
*/
func GetPwdHash(userId uint) ([]byte, error) {
	var userCreds models.UserCreds
	resp := db.Where("user_id = ?", userId).First(&userCreds)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return userCreds.Pwdhash, nil
}

/**
Create a user credentials table to store confidentials stuff.
*/
func SaveUserCreds(id uint, pwdHash []byte) error {
	userCreds := models.UserCreds{Pwdhash: pwdHash, UserID: id}
	resp := db.Save(&userCreds)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
