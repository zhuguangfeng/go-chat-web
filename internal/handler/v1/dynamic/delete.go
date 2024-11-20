package dynamic

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
)

type DeleteDynamicReq struct {
	Id int64 `json:"id"`
}

func (hdl *DynamicHandler) DeleteDynamic(ctx *gin.Context, req DeleteDynamicReq, uc jwt.UserClaims) {
	err := hdl.dynamicSvc.DeleteDynamic(ctx, req.Id, uc.Uid)
	if err != nil {
		//TODO
	}
}
