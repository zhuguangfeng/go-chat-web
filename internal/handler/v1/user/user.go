package user

import (
	"github.com/gin-gonic/gin"

	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/pkg/ginx"
)

type UserHandler struct {
	iJwt.JwtHandler
	userSvc user.UserService
}

func NewUserController(jwtHandler iJwt.JwtHandler, userSvc user.UserService) *UserHandler {
	return &UserHandler{
		JwtHandler: jwtHandler,
		userSvc:    userSvc,
	}
}

func (u *UserHandler) RegisterRouter(router *gin.Engine) {
	userG := router.Group(common.GoChatServicePath + "/user")
	{
		userG.POST("/login-pwd", ginx.WrapBody[dtoV1.UserPwdLoginReq](u.LoginPwd))
	}
}
