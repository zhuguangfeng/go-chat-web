//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhuguangfeng/go-chat/cmd/server/app"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/activity"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/dynamic"
	iJwt "github.com/zhuguangfeng/go-chat/internal/handler/v1/jwt"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/review"
	"github.com/zhuguangfeng/go-chat/internal/handler/v1/user"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	activitySvc "github.com/zhuguangfeng/go-chat/internal/service/activity"
	dynamicSvc "github.com/zhuguangfeng/go-chat/internal/service/dynamic"
	reviewSvc "github.com/zhuguangfeng/go-chat/internal/service/review"
	userSvc "github.com/zhuguangfeng/go-chat/internal/service/user"
	"github.com/zhuguangfeng/go-chat/ioc"
)

func InitWebServer() *app.App {
	wire.Build(
		ioc.InitLogger,
		ioc.InitMysql,
		ioc.InitRedisCmd,
		ioc.InitWebServer,
		ioc.InitGinMiddleware,
		ioc.InitEsClient,

		cache.NewUserCache,

		dao.NewUserDao,
		dao.NewDynamicDao,
		dao.NewReviewDao,
		dao.NewActivityDao,

		repository.NewUserRepository,
		repository.NewDynamicRepository,
		repository.NewReviewRepository,
		repository.NewActivityRepository,

		userSvc.NewUserService,
		dynamicSvc.NewDynamicService,
		activitySvc.NewActivityService,
		reviewSvc.NewReviewService,

		iJwt.NewJwtHandler,
		user.NewUserController,
		dynamic.NewDynamicHandler,
		activity.NewActivityHandler,
		review.NewReviewHandler,

		wire.Struct(new(app.App), "*"),
	)

	return new(app.App)
}
