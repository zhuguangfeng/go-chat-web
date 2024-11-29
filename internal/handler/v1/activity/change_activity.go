package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

func (hdl *ActivityHandler) ChangeActivity(ctx *gin.Context, req ChangeActivityReq, uc iJwt.UserClaims) {
	err := hdl.activitySvc.ChangeActivity(ctx, domain.Activity{
		Title:           req.Title,
		Desc:            req.Desc,
		Media:           req.Media,
		AgeRestrict:     req.AgeRestrict,
		GenderRestrict:  req.GenderRestrict,
		CostRestrict:    req.CostRestrict,
		Visibility:      req.Visibility,
		MaxPeopleNumber: req.MaxPeopleNumber,
		Address:         req.Address,
		Category:        req.Category,
		StartTime:       req.StartTime,
		DeadlineTime:    req.DeadlineTime,
	})
	if err != nil {
		hdl.logger.Error("[activity.hdl.change]修改活动信息失败",
			logger.Int64("activityId", req.ID),
			logger.Error(err),
		)
		return
	}

	common.SuccessNoData(ctx)
}
