package user

import (
	"github.com/gin-gonic/gin"

	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
)

func (hdl *UserHandler) LoginPwd(ctx *gin.Context, req dtoV1.UserPwdLoginReq) {
	user, err := hdl.userSvc.UserLoginPwd(ctx, req.Phone, req.Password)

	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	err = hdl.SetLoginToken(ctx, user.ID)
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
