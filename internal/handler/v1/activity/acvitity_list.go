package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

func (hdl *ActivityHandler) ActivityList(ctx *gin.Context) {
	var (
		searchKey         = ctx.Param("searchKey")
		pageNum, pageSize = utils.GetPagination(ctx)
	)

	activitys, count, err := hdl.activitySvc.ActivityList(ctx, pageNum, pageSize, searchKey)
	if err != nil {
		//TODO
	}

	common.Success(ctx, common.ListObj{
		CurrentCount: len(activitys),
		TotalCount:   count,
		TotalPage:    utils.GetPageCount(int(count), pageSize),
		Result: slice.Map(activitys, func(idx int, src domain.Activity) ActivityData {
			return hdl.toActivityData(src)
		}),
	})
}
