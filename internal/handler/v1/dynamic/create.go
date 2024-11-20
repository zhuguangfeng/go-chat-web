package dynamic

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/pkg/common"
)

type CreateDynamicReq struct {
	Title       string   `json:"title" dc:"标题"`
	Media       []string `json:"media" dc:"资源"`
	DynamicType int64    `json:"dynamicType" dc:"动态类型"`
	Visibility  int64    `json:"visibility" dc:"可见范围"`
}

type CreateDynamicResp struct {
}

func (hdl *DynamicHandler) CreateDynamic(ctx *gin.Context, req CreateDynamicReq, uc jwt.UserClaims) {
	dynamic := domain.Dynamic{
		User: domain.User{
			ID: uc.Uid,
		},
		Title:       req.Title,
		Media:       req.Media,
		Visibility:  req.Visibility,
		DynamicType: req.DynamicType,
		Status:      common.DynamicStatusUnderReview.Int64(),
	}

	err := hdl.dynamicSvc.CreateDynamic(ctx, dynamic)
	if err != nil {

	}

	//TODO 调用审核
}
