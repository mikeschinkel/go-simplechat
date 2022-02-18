package routes

import (
	authRouter "simple-chat-app/server/src/routes/auth"
	userRouter "simple-chat-app/server/src/routes/user"

	"github.com/gin-gonic/gin"
)

/*
Setup API Routes.
*/
func SetupApiRouter(engine *gin.Engine) {
	group := engine.Group("/api")
	group.Use(sessionMw)
	authRouter.Init(group)
	userRouter.Init(group)
}
