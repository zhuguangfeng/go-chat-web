package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"

	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
)

func (hdl *UserHandler) LoginPwd(ctx *gin.Context, req dtoV1.UserPwdLoginReq) {
	user, err := hdl.userSvc.UserLoginPwd(ctx, req.Phone, req.Password)
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	// 设置登录token 在header头里返回给前端
	err = hdl.SetLoginToken(ctx, user.ID)
	if err != nil {
		common.InternalError(ctx, errorx.NewBizError(common.SystemInternalError).WithError(err))
		return
	}

	common.SuccessNoData(ctx)
}
