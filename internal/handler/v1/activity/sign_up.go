package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
)

// SignUpActivity 活动报名
func (hdl *ActivityHandler) SignUpActivity(ctx *gin.Context, req dtoV1.SignUpActivityReq, uc iJwt.UserClaims) {
	err := hdl.activitySvc.SignUpActivity(ctx, req.ActivityID, uc.Uid)
	if err != nil {
		common.InternalError(ctx, err)
		return
	}
	common.SuccessNoData(ctx)
}
