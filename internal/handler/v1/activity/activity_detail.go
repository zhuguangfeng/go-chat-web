package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"strconv"
)

func (hdl *ActivityHandler) ActivityDetail(ctx *gin.Context) {
	var (
		id, _ = strconv.Atoi(ctx.Param("id"))
	)

	activity, err := hdl.activitySvc.ActivityDetail(ctx, int64(id))
	if err != nil {
		//TODO
		return
	}

	activity.Sponsor, err = hdl.userSvc.UserDetail(ctx, activity.Sponsor.ID)
	if err != nil {
		//TODO
		return
	}

	common.Success(ctx, hdl.toActivityData(activity))

}
