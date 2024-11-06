package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRouter(server *gin.Engine)
}
