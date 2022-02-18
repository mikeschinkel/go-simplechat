package jwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	tokenValFailedErr   = "token validation failed"
	extractingClaimsErr = "extracting claims failed"
)

var (
	secret = []byte(os.Getenv("JWT_SECRET"))
)

type JwtClaims struct {
	jwt.StandardClaims
	data interface{}
}

/**
Get a jwt string with the data encoded.
*/
func Sign(data interface{}) (string, error) {
	// If passed, get the cookie expiration time, use the same exp for jwt and cookie
	expSeconds, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		return "", err
	}
	// If passed, create a *jwt.Token with the claims
	claims := JwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expSeconds)).Unix(),
			Issuer:    "simple-chat-app/server",
		},
		data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	// Return
	return tokenString, err
}

/**
Parse a jwt string and return the data
*/
func Parse(jwtstr string) (*jwt.MapClaims, error) {
	// Parse the the token, Don't forget to validate the alg is what you expect:
	token, err := jwt.Parse(jwtstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return nil, errors.New(tokenValFailedErr)
	}
	// Check valid, extract data
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(extractingClaimsErr)
	}
	// Return
	return &claims, nil
}
