package common

import "strings"

const (
	StatusOk  int = 200
	StatusErr int = -1
)

type ErrorCode string

func (c ErrorCode) GetCodeMsg() (string, string) {
	str := string(c)
	index := strings.Index(str, ":")
	return str[:index], str[index+1:]
}

const (
	NoErr               ErrorCode = "200:成功"
	SystemInternalError ErrorCode = "GoChat.System.InternalError:服务内部错误"

	UserInvalidPassword ErrorCode = "GoChat.User.InvalidPassword:密码错误"
	UserNotFound        ErrorCode = "GoChar.User.UserNotFound:账号不存在"
)
