package auth

import (
	"net/http"
	envUtil "simple-chat-app/server/src/util/env"
	jwtUtil "simple-chat-app/server/src/util/jwt"

	"github.com/gin-gonic/gin"
)

type SessionData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

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
	session := SessionData{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
	c.Set(envUtil.SessionDataKey(), session)
	// Return
	c.Next()
}
