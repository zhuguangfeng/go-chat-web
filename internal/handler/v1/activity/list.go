package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

func (hdl *ActivityHandler) ActivityList(ctx *gin.Context, req dtoV1.ActivityListReq) {

	activitys, count, err := hdl.activitySvc.ActivityList(ctx, req)
	if err != nil {
		hdl.logger.Error("[activity.hdl.list]获取活动列表失败",
			logger.Int("pageNum", req.PageNum),
			logger.Int("pageSize", req.PageSize),
			logger.String("searchKey", req.SearchKey),
			logger.Error(err),
		)
		common.InternalError(ctx, err)
		return
	}

	common.Success(ctx, common.ListObj{
		CurrentCount: len(activitys),
		TotalCount:   count,
		TotalPage:    utils.GetPageCount(int(count), req.PageSize),
		Result: slice.Map(activitys, func(idx int, src domain.Activity) dtoV1.Activity {
			return hdl.toActivityData(src)
		}),
	})
}
