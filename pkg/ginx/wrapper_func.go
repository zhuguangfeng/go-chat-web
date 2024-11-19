package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func WrapBody[Req any](bizFn func(ctx *gin.Context, req Req)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			fmt.Println("")
		}
		bizFn(ctx, req)
	}
}
