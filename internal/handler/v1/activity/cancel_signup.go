package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
)

func (hdl *ActivityHandler) CancelSignUp(ctx *gin.Context, req dtoV1.CancelSignUpActivityReq, uc iJwt.UserClaims) {
	hdl.activitySvc.CancelSignup(ctx, domain.ActivitySignup{
		ID: req.ActivityID,
	})
}
