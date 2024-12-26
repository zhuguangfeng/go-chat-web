package activity

import (
	"errors"
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	activitySvc "github.com/zhuguangfeng/go-chat/internal/service/activity"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

// SignUpActivity 活动报名
func (hdl *ActivityHandler) SignUpActivity(ctx *gin.Context, req dtoV1.SignUpActivityReq, uc iJwt.UserClaims) {
	activity, err := hdl.activitySvc.ActivityDetail(ctx, req.ActivityID)
	if err != nil {
		if errors.Is(err, activitySvc.ErrActivityNotFound) {
			common.InternalError(ctx, errorx.NewBizError(common.ActivityNotFound))
			return
		}
		common.InternalError(ctx, err)
		return
	}

	err = hdl.activitySvc.SignUpActivity(ctx, domain.ActivitySignup{
		Activity: domain.Activity{
			ID: req.ActivityID,
			Sponsor: domain.User{
				ID: activity.Sponsor.ID,
			},
		},
		Applicant: domain.User{
			ID: uc.Uid,
		},
		Status: domain.ActivitySignupStatusPendingReview,
	})
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
