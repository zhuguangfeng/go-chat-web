package common

import (
	"errors"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode" example:"200"`                           // 操作状态码：成功200，失败-1
	Error      string      `json:"error,omitempty" example:"GoChat.User.UnAuthorized"` // 业务状态码，如果操作状态码为200，则为成功业务状态码，如果操作状态码为-1，则为失败业务状态码
	Message    string      `json:"message,omitempty" example:"没有任务操作权限"`               // 业务状态消息，如果操作状态码为200，则为成功业务消息，如果操作状态码为-1，则为失败业务消息
	ReturnObj  interface{} `json:"returnObj,omitempty"`                                // 业务数据
}

type ListObj struct {
	CurrentCount int         `json:"currentCount,omitempty" example:"10"` // 本次数据条数
	TotalCount   int64       `json:"totalCount,omitempty" example:"28"`   // 总数据条数
	TotalPage    int         `json:"totalPage,omitempty" example:"3"`     // 总数据页数
	Result       interface{} `json:"result"`                              // 业务数据
}

func successResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		StatusCode: StatusOk,
		Message:    message,
		ReturnObj:  data,
	})
}

func Success(c *gin.Context, data interface{}) {
	successResponse(c, "成功", data)
}

func SuccessNoData(c *gin.Context) {
	successResponse(c, "成功", "")
}

func SuccessMessage(c *gin.Context, message string) {
	successResponse(c, message, "")
}

func SuccessMessageData(c *gin.Context, message string, data interface{}) {
	successResponse(c, message, data)
}

func errorResponse(ctx *gin.Context, httpStatus int, code, msg string) {
	ctx.AbortWithStatusJSON(httpStatus, Response{
		StatusCode: StatusErr,
		Error:      code,
		Message:    msg,
	})
}

func wrappedErrorResponse(ctx *gin.Context, httpStatus int, code errorx.ErrorCode, msg string, err error) {
	a, b := code.GetCodeMsg()
	if msg != "" {
		b += ": " + msg
	}
	errorResponse(ctx, httpStatus, a, b)
}

func BadRequest(c *gin.Context, code errorx.ErrorCode, err error) {
	wrappedErrorResponse(c, http.StatusOK, code, "", err)
}

func InternalError(c *gin.Context, err error) {
	var bizErr errorx.BizError
	ok := errors.As(err, &bizErr)
	if ok {
		wrappedErrorResponse(c, http.StatusOK, bizErr.ErrorCode(), "", bizErr)
	} else {
		wrappedErrorResponse(c, http.StatusOK, SystemInternalError, "", err)
	}
}

func HttpErrorResp(c *gin.Context, httpStatus int, code errorx.ErrorCode, msg string, err error) {
	wrappedErrorResponse(c, httpStatus, code, msg, err)
}
