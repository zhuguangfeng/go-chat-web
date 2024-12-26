package activity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	activitySvc "github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"strconv"
)

func (hdl *ActivityHandler) ActivityDetail(ctx *gin.Context) {
	var (
		id, _ = strconv.Atoi(ctx.Query("id"))
	)

	activity, err := hdl.activitySvc.ActivityDetail(ctx, int64(id))
	if err != nil {
		if errors.Is(err, activitySvc.ErrActivityNotFound) {
			common.InternalError(ctx, errorx.NewBizError(common.ActivityNotFound))
			return
		}
		hdl.logger.Error("获取活动详情失败",
			logger.Int64("activityId", int64(id)),
			logger.Error(err),
		)
		common.InternalError(ctx, err)
		return
	}

	activity.Sponsor, err = hdl.userSvc.GetUserInfo(ctx, activity.Sponsor.ID)
	if err != nil {
		hdl.logger.Error("获取活动详情用户信息失败",
			logger.Int64("activityId", int64(id)),
			logger.Error(err),
		)
		common.InternalError(ctx, err)
		return
	}

	common.Success(ctx, hdl.toActivityData(activity))

}
