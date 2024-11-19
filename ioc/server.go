package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
)

func InitWebServer(userHandler *user.UserController) *gin.Engine {
	server := gin.Default()

	userController.RegisterRouter(server)

	return server
}
