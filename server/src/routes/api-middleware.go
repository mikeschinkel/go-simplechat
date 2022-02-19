package routes

import (
	"net/http"
	authRoutes "simple-chat-app/server/src/routes/auth"
	envUtil "simple-chat-app/server/src/util/env"
	jwtUtil "simple-chat-app/server/src/util/jwt"

	"github.com/gin-gonic/gin"
)

/**
Check the jwt-cookie is present.
*/
func apiMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(envUtil.CookieName())
	if err != nil || jwtstr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Pase the string and get the claims
	data, err := jwtUtil.Parse(jwtstr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Set Session Data
	session := authRoutes.SessionData{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
	c.Set(envUtil.SessionDataKey(), session)
	// Return
	c.Next()
}
