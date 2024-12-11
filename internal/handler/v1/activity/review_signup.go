package activity

import (
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
)

func (hdl *ActivityHandler) ReviewSignup(ctx *gin.Context, req dtov1.ReviewActivityReq) {
	err := hdl.activitySvc.ReviewSignup(ctx, domain.ActivitySignup{
		ID:     req.SignupID,
		Status: req.Status,
	})
}
