package routes

import (
	"koa-react-chat-app/server/src/daos"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
Setup the User route.
*/
func setupUserRouter(router *gin.RouterGroup) {
	group := router.Group("/users")
	group.GET("/", fetchAll)
}

/**
Fetch all users.
*/
func fetchAll(c *gin.Context) {
	userDao := daos.UserDao{}
	users := userDao.FetchAll()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
