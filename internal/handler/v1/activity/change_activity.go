package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

func (hdl *ActivityHandler) ChangeActivity(ctx *gin.Context, req dtoV1.ChangeActivityReq, uc iJwt.UserClaims) {
	err := hdl.activitySvc.ChangeActivity(ctx, domain.Activity{
		ID: req.ID,
		Sponsor: domain.User{
			ID: uc.Uid,
		},
		Title:           req.Title,
		Desc:            req.Desc,
		Media:           req.Media,
		AgeRestrict:     req.AgeRestrict,
		GenderRestrict:  req.GenderRestrict,
		CostRestrict:    req.CostRestrict,
		Visibility:      req.Visibility,
		MaxPeopleNumber: req.MaxPeopleNumber,
		Address:         req.Address,
		Category:        domain.ActivityCategory(req.Category),
		StartTime:       req.StartTime,
		DeadlineTime:    req.DeadlineTime,
	})
	if err != nil {
		hdl.logger.Error("修改活动信息失败",
			logger.Int64("activityId", req.ID),
			logger.Error(err),
		)
		return
	}

	common.SuccessNoData(ctx)
}
