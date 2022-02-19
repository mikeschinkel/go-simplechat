/**
JWT Documentation: https://pkg.go.dev/github.com/golang-jwt/jwt
Examples: https://github.com/dgrijalva/jwt-go/blob/master/example_test.go
*/

package jwt

import (
	"errors"
	"fmt"
	"time"

	envUtil "simple-chat-app/server/src/util/env"

	"github.com/golang-jwt/jwt"
)

const (
	tokenValFailedErr   = "token validation failed"
	extractingClaimsErr = "extracting claims failed"
	signMethodErr       = "unexpected signing method: %v"
)

type JwtClaims struct {
	jwt.StandardClaims
	Data interface{} `json:"data"`
}

/**
Get a jwt string with the data encoded.
*/
func Sign(data interface{}) (string, error) {
	// If passed, create a *jwt.Token with the claims
	exp := envUtil.JwtExp()
	claims := JwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(exp)).Unix(),
			Issuer:    "simple-chat-app/server",
		},
		data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	tokenStr, err := token.SignedString(envUtil.JwtSecret())
	if err != nil {
		return "", err
	}
	// Return
	return tokenStr, err
}

/**
Parse a jwt string and return the data.
*/
func Parse(jwtstr string) (*map[string]interface{}, error) {
	// Parse the the token, Don't forget to validate the alg is what you expect.
	token, err := jwt.Parse(jwtstr, parseHelper)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New(tokenValFailedErr)
	}
	// Check valid, extract data
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(extractingClaimsErr)
	}
	data := claims["data"].(map[string]interface{})
	// Return
	return &data, nil
}

/**
Provide the secret and algorithm to the jwt.Parse() method above.
*/
func parseHelper(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf(signMethodErr, token.Header["alg"])
	}
	return envUtil.JwtSecret(), nil
}
