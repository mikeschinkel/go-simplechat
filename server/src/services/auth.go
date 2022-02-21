package services

import (
	"errors"
	"simple-chat-app/server/src/dal"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/util"
	"time"
)

const (
	checkPwdFailed = "password verification failed"
)

/**
Verify user cre
*/
func VerifyAndFetchUser(email string, password string) (*models.User, error) {
	// Search for the user
	user, err := dal.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// Fetch the pwd hash
	pwdHash, err := dal.GetPwdHash(user.ID)
	if err != nil {
		return nil, err
	}
	// Compare the password to the hash. Wait 500 milliseconds if it failed as a security measure.
	passed := util.CheckPwd(pwdHash, password)
	if !passed {
		time.Sleep(time.Millisecond * 500)
		return nil, errors.New(checkPwdFailed)
	}
	// Return
	return user, nil
}
