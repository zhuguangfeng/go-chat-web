package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"net/http"
)

func WrapBody[Req any](bizFn func(ctx *gin.Context, req Req)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			common.BadRequest(ctx, common.InvalidParam, err)
			return
		}
		bizFn(ctx, req)
	}
}

func WrapBodyAndClaims[Req any, Claims jwt.Claims](bizFn func(ctx *gin.Context, req Req, uc Claims)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			common.BadRequest(ctx, common.InvalidParam, err)
			return
		}

		val, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		bizFn(ctx, req, uc)

	}

}
