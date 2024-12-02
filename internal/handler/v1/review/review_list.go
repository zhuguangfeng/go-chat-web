package review

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"strconv"
)

func (hdl *ReviewHandler) ReviewList(ctx *gin.Context) {
	var (
		pageNum, _  = strconv.Atoi(ctx.Query("pageNum"))
		pageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
		biz         = ctx.Query("biz")
		status, _   = strconv.Atoi(ctx.Query("status"))
	)
	reviews, _, err := hdl.reviewSvc.ReviewList(ctx, pageNum, pageSize, biz, uint(status))
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.Success(ctx, common.ListObj{
		Result: reviews,
	})
}
