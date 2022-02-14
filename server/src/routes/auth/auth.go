package auth

import (
	"net/http"
	"os"
	authService "simple-chat-app/server/src/services/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
Add the auth-router (Group) to the gin-engine.
*/
func Init(router *gin.RouterGroup) {
	group := router.Group("/auth")
	group.PUT("/login", login)
	group.GET("/logout", logout)
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
	jwtstr, err := authService.VerifyUserAndGetToken(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Setup environment variables (convert string to various types)
	maxAge, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	isSecure, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Store the jwt in a cookie if passed
	c.SetCookie(os.Getenv("COOKIE_NAME"), jwtstr, maxAge, os.Getenv("COOKIE_PATH"),
		os.Getenv("COOKIE_DOMAIN"), isSecure, true)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": jwtstr})
}

/**
URL: "/api/auth/logout"
*/
func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello, you are signed out"})
}
