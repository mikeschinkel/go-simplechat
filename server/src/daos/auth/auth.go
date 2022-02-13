package auth

import (
	"simple-chat-app/server/src/daos"
	"simple-chat-app/server/src/models"
)

/**
Fetch a user's login credentials
*/
func GetUserCreds(userId uint) *models.UserCreds {
	userCreds := models.UserCreds{Pwdhash: "Password@1"}
	return &userCreds
}

/**
Create a user credentials table to store confidentials stuff.
*/
func SaveUserCreds(id uint, pwdHash string) error {
	db := daos.GetDbConn()
	userCreds := models.UserCreds{Pwdhash: pwdHash}
	resp := db.Save(userCreds)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
