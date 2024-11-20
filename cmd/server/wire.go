//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhuguangfeng/go-chat/cmd/server/app"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/dynamic"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	dynamicSvc "github.com/zhuguangfeng/go-chat/internal/service/dynamic"
	userSvc "github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/ioc"
)

func InitWebServer() *app.App {
	wire.Build(
		ioc.InitMysql,
		ioc.InitRedisCmd,
		ioc.InitWebServer,

		dao.NewUserDao,
		cache.NewUserCache,
		dao.NewDynamicDao,

		repository.NewUserRepository,
		repository.NewDynamicRepository,

		userSvc.NewUserService,
		dynamicSvc.NewDynamicService,

		iJwt.NewJwtHandler,
		user.NewUserController,
		dynamic.NewDynamicHandler,

		wire.Struct(new(app.App), "*"),
	)

	return new(app.App)
}
