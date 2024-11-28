package activity

import (
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
)

type ActivityHandler struct {
	activitySvc activity.ActivityService
	userSvc     user.UserService
}

func NewActivityHandler(activitySvc activity.ActivityService) *ActivityHandler {
	return &ActivityHandler{
		activitySvc: activitySvc,
	}
}

type BaseReq struct {
	ID int64 `json:"id" dc:"活动id"`
}

type CreateActivityReq struct {
	UserID          int64    `json:"user" dc:"活动发起人"`
	Title           string   `json:"title" dc:"活动标题"`
	Desc            string   `json:"desc" dc:"活动描述"`
	Media           []string `json:"media" dc:"资源 视频或图片"`
	AgeRestrict     uint     `json:"ageRestrict" dc:"最大年龄"`
	GenderRestrict  uint     `json:"genderRestrict" dc:"性别限制"`
	CostRestrict    uint     `json:"costRestrict" dc:"费用限制"`
	Visibility      uint     `json:"visibility" dc:"可见度"`
	MaxPeopleNumber int64    `json:"maxPeopleNumber" dc:"最大报名人数"`
	Address         string   `json:"address" dc:"活动地址"`
	Category        int64    `json:"category" dc:"活动分类"`
	StartTime       uint     `json:"startTime" dc:"活动开始时间"`
	DeadlineTime    uint     `json:"deadlineTime" dc:"活动截止时间"`
}

type ChangeActivityReq struct {
	ID int64 `json:"id" dc:"活动ID"`
	CreateActivityReq
}

type ActivityData struct {
	UserID          int64    `json:"user_id,omitempty"`
	UserName        string   `json:"user_name,omitempty"`
	Avatar          string   `json:"avatar,omitempty"`
	Title           string   `json:"title,omitempty" dc:"活动标题"`
	Desc            string   `json:"desc,omitempty" dc:"活动描述"`
	Media           []string `json:"media,omitempty" dc:"资源 视频或图片"`
	AgeRestrict     uint     `json:"ageRestrict,omitempty" dc:"最大年龄"`
	GenderRestrict  uint     `json:"genderRestrict,omitempty" dc:"性别限制"`
	CostRestrict    uint     `json:"costRestrict,omitempty" dc:"费用限制"`
	Visibility      uint     `json:"visibility,omitempty" dc:"可见度"`
	MaxPeopleNumber int64    `json:"maxPeopleNumber,omitempty" dc:"最大报名人数"`
	Address         string   `json:"address,omitempty" dc:"活动地址"`
	Category        int64    `json:"category,omitempty" dc:"活动分类"`
	StartTime       uint     `json:"startTime,omitempty" dc:"活动开始时间"`
	DeadlineTime    uint     `json:"deadlineTime,omitempty" dc:"活动截止时间"`
	Status          uint     `json:"status,omitempty" dc:"活动截止时间"`
}

func (hdl *ActivityHandler) toActivityData(activity domain.Activity) ActivityData {
	return ActivityData{
		UserID:          activity.Sponsor.ID,
		UserName:        activity.Sponsor.UserName,
		Avatar:          "",
		Title:           activity.Title,
		Desc:            activity.Desc,
		Media:           activity.Media,
		AgeRestrict:     activity.AgeRestrict,
		GenderRestrict:  activity.GenderRestrict,
		CostRestrict:    activity.CostRestrict,
		Visibility:      activity.Visibility,
		MaxPeopleNumber: activity.MaxPeopleNumber,
		Address:         activity.Address,
		Category:        activity.Category,
		StartTime:       activity.StartTime,
		DeadlineTime:    activity.DeadlineTime,
		Status:          activity.Status,
	}
}
