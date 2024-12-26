package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/internal/service/review"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type ActivityHandler struct {
	logger      logger.Logger
	activitySvc activity.ActivityService
	userSvc     user.UserService
	reviewSvc   review.ReviewService
}

func NewActivityHandler(logger logger.Logger, activitySvc activity.ActivityService, userSvc user.UserService, reviewSvc review.ReviewService) *ActivityHandler {
	return &ActivityHandler{
		logger:      logger,
		activitySvc: activitySvc,
		userSvc:     userSvc,
	}
}

func (hdl *ActivityHandler) RegisterRouter(router *gin.Engine) {
	activityG := router.Group(common.GoChatServicePath + "/activity")
	{
		activityG.POST("/create", ginx.WrapBodyAndClaims[dtoV1.CreateActivityReq, jwt.UserClaims](hdl.CreateActivity))
		activityG.POST("/cancel", ginx.WrapBodyAndClaims[dtoV1.BaseDeleteReq, jwt.UserClaims](hdl.CancelActivity))
		activityG.POST("/change", ginx.WrapBodyAndClaims[dtoV1.ChangeActivityReq, jwt.UserClaims](hdl.ChangeActivity))
		activityG.GET("/detail", hdl.ActivityDetail)
		activityG.POST("/list", ginx.WrapBody[dtoV1.ActivityListReq](hdl.ActivityList))

		activityG.POST("/signup", ginx.WrapBodyAndClaims[dtoV1.SignUpActivityReq, jwt.UserClaims](hdl.SignUpActivity))
		activityG.POST("/signup-list", ginx.WrapBodyAndClaims[dtoV1.SignUpListReq, jwt.UserClaims](hdl.SignupList))
		activityG.POST("/review-signup", ginx.WrapBody[dtoV1.ReviewSignupReq](hdl.ReviewSignup))
		activityG.POST("/cancel-signup", ginx.WrapBodyAndClaims[dtoV1.CancelSignUpActivityReq, jwt.UserClaims](hdl.CancelSignUp))
	}
}

func (hdl *ActivityHandler) toActivityData(activity domain.Activity) dtoV1.Activity {
	return dtoV1.Activity{
		UserID:          activity.Sponsor.ID,
		Username:        activity.Sponsor.Username,
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
		Category:        activity.Category.Uint(),
		StartTime:       activity.StartTime,
		DeadlineTime:    activity.DeadlineTime,
		Status:          activity.Status.Uint(),
	}
}
