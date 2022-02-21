package routes

import (
	"github.com/gin-gonic/gin"
)

/*
Setup API Routes.
*/
func Init(engine *gin.Engine) {
	group := engine.Group("/api")
	initAuthRoutes(group)
	group.Use(apiMiddleware)
	initUserRoutes(group)
}
