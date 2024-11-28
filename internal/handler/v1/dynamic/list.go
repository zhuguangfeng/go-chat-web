package dynamic

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DynamicListReq struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	SearchKey string `json:"searchKey"`
}

func (hdl *DynamicHandler) DynamicList(ctx *gin.Context, req DynamicListReq) {

	dynamic, count, err := hdl.dynamicSvc.ListDynamic(ctx, req.PageNum, req.PageSize, req.SearchKey)
	if err != nil {
		//TODO
	}
	fmt.Println(count)

	ctx.JSON(200, dynamic)
}
