package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

func (hdl *ActivityHandler) SignupList(ctx *gin.Context, req dtoV1.SignUpListReq, uc iJwt.UserClaims) {
	req.UID = uc.Uid
	signups, count, err := hdl.activitySvc.SignupList(ctx, req)
	if err != nil {
		common.InternalError(ctx, errorx.NewBizError(common.SystemInternalError).WithError(err))
		return
	}

	common.Success(ctx, common.ListObj{
		CurrentCount: len(signups),
		TotalCount:   count,
		Result:       signups,
	})

}
