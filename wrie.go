//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhuguangfeng/go-chat/go-chat/cmd/server/app"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/controller/v1/user"
	iJwt "github.com/zhuguangfeng/go-chat/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/go-chat/internal/repository/dao"
	userSvc "github.com/zhuguangfeng/go-chat/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/go-chat/ioc"
)

func InitWebServer() *app.App {
	wire.Build(
		ioc.InitMysql,
		ioc.InitRedisCmd,
		ioc.InitWebServer,

		dao.NewUserDao,
		cache.NewUserCache,

		repository.NewUserRepository,

		userSvc.NewUserService,

		iJwt.NewJwtHandler,
		user.NewUserController,

		wire.Struct(new(app.App), "*"),
	)

	return new(app.App)
}
