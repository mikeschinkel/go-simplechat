package auth

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	*jwt.StandardClaims
	TokenType string
	JwtUserParams
}

type JwtUserParams struct {
	id    uint
	email string
	name  string
}
