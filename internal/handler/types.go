package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRouter(server *gin.Engine)
}
