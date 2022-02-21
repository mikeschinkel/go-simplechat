package util

import "golang.org/x/crypto/bcrypt"

/**
Generate a hash from a password.
*/
func HashPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

/**
Check password against hash.
*/
func CheckPwd(pwdHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password))
	return err == nil
}
