package review

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
)

func (hdl *ReviewHandler) ReviewDetail(ctx *gin.Context) {
	var (
		uuid = ctx.Query("uuid")
	)

	review, err := hdl.reviewSvc.ReviewDetail(ctx, uuid)
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.Success(ctx, review)
}
