package routes

import (
	"net/http"
	"os"
	authService "simple-chat-app/server/src/services/auth"
	jwtUtil "simple-chat-app/server/src/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
Check the jwt-cookie is present.
*/
func sessionMw(c *gin.Context) {
	// Require login cookie for all /api routes except auth
	urlstr := c.Request.URL.String()
	if strings.HasPrefix(urlstr, "/api/auth") {
		c.Next()
		return
	}
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(os.Getenv("COOKIE_NAME"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		return
	}
	// Pase the string and get the claims
	claims, err := jwtUtil.Parse(jwtstr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		return
	}
	// Set Session Data
	session := authService.UserData{
		ID:    (*claims)["ID"].(uint),
		Email: (*claims)["Email"].(string),
		Name:  (*claims)["Name"].(string),
	}
	c.Set("session", session)
	// Return
	c.Next()
}
