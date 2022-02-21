package routes

import (
	"net/http"
	"simple-chat-app/server/src/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

type AddUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

/**
Setup the User route.
*/
func initUserRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")
	group.GET("/", fetchAll)
	group.POST("/", addOne)
	group.DELETE("/:id", deleteOne)
}

/**
Fetch all users.
*/
func fetchAll(c *gin.Context) {
	users, err := services.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

/**
Add a new user.
*/
func addOne(c *gin.Context) {
	// Extract user from json
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = services.AddUser(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

/**
Update one user.
*/
// func updateOne(c *gin.Context) {
// 	// Extract user from json
// 	var req UpdateUserReq
// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
// 		return
// 	}
// 	// Query db
// 	err = userService.UpdateOne(req.Idreq.Email, req.Name, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"status": "success"})
// }

/**
Delete one user.
*/
func deleteOne(c *gin.Context) {
	// Convert query string to unint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Delete the user
	err = services.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
