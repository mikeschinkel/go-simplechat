package auth

import (
	"net/http"
	"simple-chat-app/server/src/shared/globals"
	envUtil "simple-chat-app/server/src/util/env"
	jwtUtil "simple-chat-app/server/src/util/jwt"

	"github.com/gin-gonic/gin"
)

/**
Check the jwt-cookie is present.
*/
func authMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(envUtil.CookieName())
	if jwtstr == "" || err != nil {
		c.Next()
		return
	}
	// Pase the string and get the claims
	data, err := jwtUtil.Parse(jwtstr)
	if data == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// Set Session Data
	session := GetSessionData(data)
	c.Set(globals.SessionDataKey(), session)
	// Return
	c.Next()
}

/**
The API middleware needs this too
*/
func GetSessionData(data *map[string]interface{}) *SessionData {
	return &SessionData{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
}
