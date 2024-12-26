package activity

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

func (hdl *ActivityHandler) CancelActivity(ctx *gin.Context, req dtov1.BaseDeleteReq, uc iJwt.UserClaims) {
	err := hdl.activitySvc.CancelActivity(ctx, req.ID)
	if err != nil {
		hdl.logger.Error("[activity.hdl.cancel]取消活动失败",
			logger.Int64("activityId", req.ID),
			logger.Error(err),
		)
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
