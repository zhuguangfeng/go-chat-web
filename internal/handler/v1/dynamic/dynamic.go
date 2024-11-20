package dynamic

import (
	"github.com/gin-gonic/gin"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	dynamicSvc "github.com/zhuguangfeng/go-chat/internal/service/dynamic"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
)

type DynamicHandler struct {
	dynamicSvc dynamicSvc.DynamicService
}

func NewDynamicHandler(dynamicSvc dynamicSvc.DynamicService) *DynamicHandler {
	return &DynamicHandler{
		dynamicSvc: dynamicSvc,
	}
}

func (hdl *DynamicHandler) RegisterRouter(router *gin.Engine) {
	dynamicG := router.Group("/dynamic")
	{
		dynamicG.POST("/create", ginx.WrapBodyAndClaims[CreateDynamicReq, iJwt.UserClaims](hdl.CreateDynamic))
		dynamicG.POST("/delete", ginx.WrapBodyAndClaims[DeleteDynamicReq, iJwt.UserClaims](hdl.DeleteDynamic))
		dynamicG.POST("/list", ginx.WrapBody[DynamicListReq](hdl.DynamicList))

		dynamicG.GET("/detail", hdl.DetailDynamic)
	}
}
