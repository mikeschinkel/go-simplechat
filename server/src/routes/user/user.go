package user

import (
	"net/http"
	userService "simple-chat-app/server/src/services/user"

	"github.com/gin-gonic/gin"
)

/**
Setup the User route.
*/
func Init(router *gin.RouterGroup) {
	group := router.Group("/users")
	group.GET("/", fetchAll)
	group.POST("/", addOne)
}

/**
Fetch all users.
*/
func fetchAll(c *gin.Context) {
	users := userService.FetchAll()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

/**
Add a new user.
*/
func addOne(c *gin.Context) {
	// Extra user from json
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = userService.AddOne(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
