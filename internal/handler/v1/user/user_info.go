package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"strconv"
)

func (hdl *UserHandler) UserInfo(ctx *gin.Context) {
	var (
		id, _ = strconv.Atoi(ctx.Param("id"))
	)
	user, err := hdl.userSvc.GetUserInfo(ctx, int64(id))
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.Success(ctx, hdl.toUser(user))

}
