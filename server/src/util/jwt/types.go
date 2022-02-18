package jwt

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	data interface{}
}
