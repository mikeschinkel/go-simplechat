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
}

/**
Fetch all users.
*/
func fetchAll(c *gin.Context) {
	users := userService.FetchAll()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
