package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/controller/v1/user"
)

func InitWebServer(userController *user.UserController) *gin.Engine {
	server := gin.Default()

	userController.RegisterRouter(server)
	return server
}
