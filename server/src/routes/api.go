package routes

import (
	"github.com/gin-gonic/gin"
)

/*
Setup API Routes.
*/
func SetupApiRouter(engine *gin.Engine) {
	group := engine.Group("/api")
	setupAuthRouter(group)
	setupUserRouter(group)
}
