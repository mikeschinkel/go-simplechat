package simplechat

import "github.com/gin-gonic/gin"

// addRoutes sets up the API URL routes
// NOTE I put this here to be part of the reusable server, but
// it could easily be moved to app/main.go and only be part of
// the app.
func addRoutes(engine *gin.Engine) {

	engine.GET("/", func(c *gin.Context) {
		c.String(200, serverStartMsg)
	})

	api := engine.Group("/api")

	// et up the routes for users
	group := api.Group("/users")
	group.GET("/", handleGetAllUsersRequest)
	group.POST("/", handleUserAddRequest)
	group.PUT("/", handleUserUpdateRequest)
	group.DELETE("/:id", handleUserDeleteRequest)

	group.Use(APIMiddleware)

	// Add the auth-router (Group) to the gin-engine.
	group = api.Group("/auth")
	group.PUT("/handleLoginRequest", handleLoginRequest)
	group.GET("/handleLogoutRequest", handleLogoutRequest)

	group.Use(AuthMiddleware)

	group.GET("/session-data", handleGetSessionRequest)

}
