package dynamic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"strconv"
)

type DynamicListReq struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Query    string `json:"query"`
	UserID   int64  `json:"userId"`
}

func (hdl *DynamicHandler) DynamicList(ctx *gin.Context, req DynamicListReq) {

	var conditions = make([]mysqlx.Condition, 0)

	if req.Query != "" {
		conditions = append(conditions, mysqlx.Condition{
			Key:   "title",
			Where: "like",
			Val:   "%" + req.Query + "%",
		})
	}

	if req.UserID > 0 {
		conditions = append(conditions, mysqlx.Condition{
			Key:   "uid",
			Where: "=",
			Val:   strconv.FormatInt(req.UserID, 10),
		})
	}

	dynamic, count, err := hdl.dynamicSvc.ListDynamic(ctx, req.PageNum, req.PageSize, conditions)
	if err != nil {
		//TODO
	}
	fmt.Println(count)

	ctx.JSON(200, dynamic)
}
