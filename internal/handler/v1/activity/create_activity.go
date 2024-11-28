package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
)

func (hdl *ActivityHandler) Create(ctx *gin.Context, req CreateActivityReq) {
	//创建活动
	err := hdl.activitySvc.CreateActivity(ctx, domain.Activity{
		Sponsor: domain.User{
			ID: req.UserID,
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
		//TODO
		return
	}

	//发起活动审批

}
