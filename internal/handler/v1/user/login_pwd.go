package user

import (
	"fmt"
	"github.com/gin-gonic/gin"

	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/pkg/common"
)

func (u *UserHandler) LoginPwd(ctx *gin.Context, req dtoV1.UserPwdLoginReq) {
	user, errCode, err := u.userSvc.UserLoginPwd(ctx, req.Phone, req.Password)

	if err != nil {

		if errCode == common.UserNotFound {
			common.BadRequest(ctx, errCode, err)
			return
		}
		common.InternalError(ctx, errCode, err)
		return
	}

	fmt.Println(user.ID)
	err = u.SetLoginToken(ctx, user.ID)
	if err != nil {
		common.InternalError(ctx, common.SystemInternalError, err)
		return
	}

	common.SuccessNoData(ctx)
}
