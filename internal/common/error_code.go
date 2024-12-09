package common

import (
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

const (
	StatusOk  int = 200
	StatusErr int = -1
)

const (
	NoErr               errorx.ErrorCode = "200:成功"
	SystemInternalError errorx.ErrorCode = "GoChat.System.InternalError:服务内部错误"
	InvalidParam        errorx.ErrorCode = "GoChat.System.InvalidParam:请求参数有误"

	UserInvalidPassword errorx.ErrorCode = "GoChat.User.InvalidPassword:密码错误"
	UserPhoneNotFound   errorx.ErrorCode = "GoChat.User.UserPhoneNotFound:该手机号码不存在"
	UserNotFound        errorx.ErrorCode = "GoChat.User.UserNotFound:用户不存在"
	UserIsExists        errorx.ErrorCode = "GoChat.User.UserIsExists:用户已存在"

	ActivityNotFound  errorx.ErrorCode = "GoChat.Activity.ActivityNotFound:活动不存在"
	ActivityNotChange errorx.ErrorCode = "GoChat.Activity.NotChange:该活动暂时不能修改"
	ActivityNotCancel errorx.ErrorCode = "GoChat.Activity.NotCancel:该活动暂时不能取消"

	ReviewNotReview errorx.ErrorCode = "GoChat.Review.NotReview:已经审核过的不能再次审核"
)
