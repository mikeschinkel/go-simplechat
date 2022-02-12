package routes

import (
	"koa-react-chat-app/server/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string
	Password string
}

/**
Add the auth-router (Group) to the gin-engine.
*/
func setupAuthRouter(router *gin.RouterGroup) {
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
	authService := services.AuthService{}
	jwt, err := authService.VerifyUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Return json
	c.JSON(http.StatusUnauthorized, gin.H{"jwt": jwt})
}

/**
URL: "/api/auth/logout"
*/
func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello, you are signed out"})
}
