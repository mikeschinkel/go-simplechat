package auth

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	UserData JwtUserData
}

type JwtUserData struct {
	id    uint
	email string
	name  string
}
