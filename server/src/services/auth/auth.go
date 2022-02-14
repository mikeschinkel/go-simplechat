package auth

import (
	"os"
	authDao "simple-chat-app/server/src/daos/auth"
	userDao "simple-chat-app/server/src/daos/user"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	pwdVerificationFailedMsg = "password verification failed"
)

/**
Verify user cre
*/
func VerifyUserAndGetToken(
	email string,
	password string,
) (string, error) {
	// Search for the user
	user, err := userDao.FindByEmail(email)
	if err != nil {
		return "", err
	}
	// Fetch the pwd hash
	pwdHash, err := authDao.GetPwdHash(user.ID)
	if err != nil {
		return "", err
	}
	// Compare the password to the hash
	err = bcrypt.CompareHashAndPassword(pwdHash, []byte(password))
	if err != nil {
		return "", err
	}
	// If passed, get the cookie expiration time, use the same exp for jwt and cookie
	expSeconds, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		return "", err
	}
	// If passed, create a *jwt.Token with the claims
	claims := JwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expSeconds)).Unix(),
			Issuer:    "simple-chat-app/server",
		},
		"level1",
		JwtUserParams{
			id:    user.ID,
			email: user.Email,
			name:  user.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	tokenSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	// Generate a jsonwebtoken if passed
	return tokenString, nil
}
