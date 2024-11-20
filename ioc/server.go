package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/dynamic"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
)

func InitWebServer(userHandler *user.UserHandler, dynamicHandler *dynamic.DynamicHandler) *gin.Engine {
	server := gin.Default()

	userHandler.RegisterRouter(server)
	dynamicHandler.RegisterRouter(server)

	return server
}
