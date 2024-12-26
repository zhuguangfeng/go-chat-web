package activity

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
)

func (hdl *ActivityHandler) ReviewSignup(ctx *gin.Context, req dtov1.ReviewSignupReq) {
	err := hdl.activitySvc.ReviewSignup(ctx, domain.ActivitySignup{
		ID:     req.SignupID,
		Status: domain.ActivitySignupStatus(req.Status),
	})
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
