package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"time"

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

func (hdl *UserHandler) RegisterRouter(router *gin.Engine) {
	userG := router.Group(common.GoChatServicePath + "/user")
	{
		userG.POST("/login-pwd", ginx.WrapBody[dtoV1.UserPwdLoginReq](hdl.LoginPwd))
		userG.GET("user-info", hdl.UserInfo)
	}
}

func (hdl *UserHandler) toUser(user domain.User) dtoV1.User {
	return dtoV1.User{
		ID:            user.ID,
		Username:      user.UserName,
		Phone:         user.Phone,
		Age:           user.Age,
		Gender:        user.Gender,
		IsRealName:    user.IsRealName,
		LoginIp:       user.LastLoginIp,
		LastLoginTime: time.Unix(int64(user.LastLoginTime), 0),
		Status:        user.Status,
		CreatedTime:   time.Unix(int64(user.CreatedTime), 0),
		UpdatedTime:   time.Unix(int64(user.CreatedTime), 0),
	}
}
