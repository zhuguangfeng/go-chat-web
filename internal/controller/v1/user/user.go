package user

import (
	"github.com/gin-gonic/gin"

	dtoV1 "github.com/zhuguangfeng/go-chat/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/common"
	iJwt "github.com/zhuguangfeng/go-chat/go-chat/internal/controller/v1/jwt"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/go-chat/pkg/ginx"
)

type UserController struct {
	iJwt.JwtHandler
	userSvc user.UserService
}

func NewUserController(jwtHandler iJwt.JwtHandler, userSvc user.UserService) *UserController {
	return &UserController{
		JwtHandler: jwtHandler,
		userSvc:    userSvc}

}

func (u *UserController) RegisterRouter(router *gin.Engine) {
	userG := router.Group(common.GoChatServicePath + "/user")
	{
		userG.POST("/login-pwd", ginx.WrapBody[dtoV1.UserPwdLoginReq](u.LoginPwd))
	}
}
