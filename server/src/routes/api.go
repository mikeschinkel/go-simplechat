package routes

import (
	authRoutes "simple-chat-app/server/src/routes/auth"
	userRoutes "simple-chat-app/server/src/routes/user"

	"github.com/gin-gonic/gin"
)

/*
Setup API Routes.
*/
func Init(engine *gin.Engine) {
	group := engine.Group("/api")
	authRoutes.Init(group)
	group.Use(apiMiddleware)
	userRoutes.Init(group)
}
