package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

/**
Check the jwt-cookie is present.
*/
func sessionMw(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(os.Getenv("COOKIE_NAME"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		return
	}
	// Parse the the token
	token, err := jwt.Parse(jwtstr, sessionMwHelper)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		return
	}
	// Check valid, extract data
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	// pick up here
	c.Set("Session", claims)
	// Next
	c.Next()
}

/**
Check the algo and pass the jwt secret.
*/
func sessionMwHelper(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	// jwt secret is a []byte containing your secret, e.g. []byte("my_secret_key")
	secret := []byte(os.Getenv("JWT_SECRET"))
	return secret, nil
}
