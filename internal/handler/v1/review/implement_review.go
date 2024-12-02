package review

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
)

func (hdl *ReviewHandler) ImplementReview(ctx *gin.Context, req ImplementReviewReq, uc iJwt.UserClaims) {
	err := hdl.reviewSvc.ImplementReview(ctx, domain.Review{
		UUID:    req.UUID,
		Opinion: req.Opinion,
		Reviewer: domain.User{
			ID: uc.Uid,
		},
		Status: req.Status,
	})
	if err != nil {
		common.InternalError(ctx, err)
		return
	}

	common.SuccessNoData(ctx)
}
