package auth

import (
	"net/http"
	"os"
	authService "simple-chat-app/server/src/services/auth"
	jwtUtil "simple-chat-app/server/src/util/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	sessionDataKey = "session"
)

/**
Add the auth-router (Group) to the gin-engine.
*/
func Init(router *gin.RouterGroup) {
	group := router.Group("/auth")
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
	// Verify the user and get a jwt if they passed.
	user, err := authService.VerifyUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	jwtstr, err := jwtUtil.Sign(&UserData{user.ID, user.Email, user.Name})
	// Get the time to expire in seconds
	maxAge, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Store the jwt in a cookie if passed
	name, path, domain, isSecure, err := getCookieVals()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	c.SetCookie(name, jwtstr, maxAge, path, domain, isSecure, true)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": true})
}

/**
URL: "/api/auth/logout"

Logout user by setting cookies maxAge = 0 and removing jwtstr
*/
func logout(c *gin.Context) {
	// Erase cookie data
	name, path, domain, isSecure, err := getCookieVals()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"logout": err.Error()})
		return
	}
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
	session, exists := c.Get(sessionDataKey)
	if !exists {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	// Return the data if it's there
	c.JSON(http.StatusOK, gin.H{"data": session})
}

/***************************************************************************************
                                      Shared
***************************************************************************************/

func getCookieVals() (string, string, string, bool, error) {
	name := os.Getenv("COOKIE_DOMAIN")
	path := os.Getenv("COOKIE_PATH")
	domain := os.Getenv("COOKIE_DOMAIN")
	isSecure, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	return name, path, domain, isSecure, err
}
