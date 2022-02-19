package auth

import (
	"net/http"
	authService "simple-chat-app/server/src/services/auth"
	envUtil "simple-chat-app/server/src/util/env"
	jwtUtil "simple-chat-app/server/src/util/jwt"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/**
Add the auth-router (Group) to the gin-engine.
*/
func Init(router *gin.RouterGroup) {
	group := router.Group("/auth")
	group.Use(authMiddleware)
	group.PUT("/login", login)
	group.GET("/logout", logout)
	group.GET("/session-data", getSessionData)
}

/**
URL: "/api/auth/login"
*/
func login(c *gin.Context) {
	// Set req data
	var loginReq LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Verify and fetch the user
	user, err := authService.VerifyAndFetchUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Get a jwt string if the user passed authentication
	sessionData := SessionData{user.ID, user.Email, user.Name}
	jwtstr, err := jwtUtil.Sign(&sessionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Set the cookie
	name, exp, path, domain, isSecure := envUtil.GetCookieVals()
	c.SetCookie(name, jwtstr, exp, path, domain, isSecure, true)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": true})
}

/**
- URL: "/api/auth/logout"
- Logout user by setting cookies maxAge = 0 and removing jwtstr
*/
func logout(c *gin.Context) {
	// Set the cookie
	name, _, path, domain, isSecure := envUtil.GetCookieVals()
	c.SetCookie(name, "", 0, path, domain, isSecure, true)
	// Return
	c.JSON(http.StatusOK, gin.H{"success": true})
}

/**
URL: "/api/auth/session"
*/
func getSessionData(c *gin.Context) {
	// Check if the user is not logged in, if not that's okay,
	// there just won't be any session data
	session, exists := c.Get(envUtil.SessionDataKey())
	if !exists {
		c.JSON(http.StatusOK, gin.H{"logged-in": false})
		return
	}
	// Return the data if it's there
	c.JSON(http.StatusOK, gin.H{"data": session})
}
