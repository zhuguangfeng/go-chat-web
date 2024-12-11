package review

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/review"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
)

type ReviewHandler struct {
	reviewSvc review.ReviewService
}

func NewReviewHandler(reviewSvc review.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewSvc: reviewSvc,
	}

}

func (hdl *ReviewHandler) RegisterRouter(router *gin.Engine) {
	reviewG := router.Group(common.GoChatServicePath + "/review")
	{
		reviewG.POST("/create-activity", ginx.WrapBodyAndClaims[ImplementReviewReq, iJwt.UserClaims](hdl.ReviewCreateActivity))
		reviewG.GET("/detail", hdl.ReviewDetail)
		reviewG.GET("/list", hdl.ReviewList)

	}
}

type ImplementReviewReq struct {
	UUID    string `json:"uuid" dc:"uuid"`
	Status  uint   `json:"status" dc:"审批状态"`
	Opinion string `json:"opinion" dc:"审批意见"`
}
