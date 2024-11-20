package dynamic

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (hdl *DynamicHandler) DetailDynamic(ctx *gin.Context) {
	var (
		id, _ = strconv.Atoi(ctx.Param("id"))
	)
	dynamic, err := hdl.dynamicSvc.DetailDynamic(ctx, int64(id))
	if err != nil {
		//TODO
	}

	ctx.JSON(200, dynamic)
}
