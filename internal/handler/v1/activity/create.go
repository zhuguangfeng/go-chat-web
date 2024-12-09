package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

func (hdl *ActivityHandler) CreateActivity(ctx *gin.Context, req CreateActivityReq, uc iJwt.UserClaims) {
	//创建活动
	err := hdl.activitySvc.CreateActivity(ctx, domain.Activity{
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
		Category:        req.Category,
		StartTime:       req.StartTime,
		DeadlineTime:    req.DeadlineTime,
		Status:          common.ActivityStatusPendingReview.Uint(),
	})
	if err != nil {
		hdl.logger.Error("[ActivityHdl.CreateActivity]创建活动失败", logger.Error(err))
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
