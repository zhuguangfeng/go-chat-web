package review

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"time"
)

func (hdl *ReviewHandler) ReviewCreateActivity(ctx *gin.Context, req ImplementReviewReq, uc iJwt.UserClaims) {
	err := hdl.reviewSvc.ReviewCreateActivity(ctx, domain.Review{
		UUID:    req.UUID,
		Opinion: req.Opinion,
		Reviewer: domain.User{
			ID: uc.Uid,
		},
		Status:     domain.ReviewStatus(req.Status),
		ReviewTime: uint(time.Now().Unix()),
	})
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
