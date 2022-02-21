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

type UpdateUserReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

/**
Setup the User route.
*/
func initUserRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")
	group.GET("/", fetchAllUsers)
	group.POST("/", addUser)
	group.PUT("/", updateUser)
	group.DELETE("/:id", deleteUser)
}

/**
Fetch all users.
*/
func fetchAllUsers(c *gin.Context) {
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
func addUser(c *gin.Context) {
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
Update user's email and name.
*/
func updateUser(c *gin.Context) {
	// Extract user from json
	var req UpdateUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = services.UpdateUser(req.ID, req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

/**
Delete one user.
*/
func deleteUser(c *gin.Context) {
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
