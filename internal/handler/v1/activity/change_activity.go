package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/domain"
)

func (hdl *ActivityHandler) ChangeActivity(ctx *gin.Context, req ChangeActivityReq) {
	err := hdl.activitySvc.ChangeActivity(ctx, domain.Activity{
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
	})
	if err != nil {
		//TODO
		return
	}

}
