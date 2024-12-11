package activity

import (
	"github.com/gin-gonic/gin"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
)

// SignUpActivity 活动报名
func (hdl *ActivityHandler) SignUpActivity(ctx *gin.Context, req dtoV1.SignUpActivityReq, uc iJwt.UserClaims) {
	activity, err := hdl.activitySvc.ActivityDetail(ctx, req.ActivityID)
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	err = hdl.activitySvc.SignUpActivity(ctx, domain.ActivitySignup{
		ActivityID:  req.ActivityID,
		ApplicantID: uc.Uid,
		SponsorID:   activity.Sponsor.ID,
		Status:      common.ReviewStatusPendingReview.Uint(),
	})
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
