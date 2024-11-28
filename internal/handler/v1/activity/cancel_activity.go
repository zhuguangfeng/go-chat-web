package activity

import (
	"github.com/gin-gonic/gin"
)

func (hdl *ActivityHandler) CancelActivity(ctx *gin.Context, req BaseReq) {
	err := hdl.activitySvc.CancelActivity(ctx, req.ID)
	if err != nil {
		//TODO
		return
	}
}
