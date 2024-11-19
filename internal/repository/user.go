package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
)

type UserRepository interface {
	GetUserByPhone(ctx context.Context, phone string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

type CacheUserRepository struct {
	userDao   dao.UserDao
	userCache cache.UserCache
}

func NewUserRepository(userDao dao.UserDao, userCache cache.UserCache) UserRepository {
	return &CacheUserRepository{
		userDao:   userDao,
		userCache: userCache,
	}
}

func (u *CacheUserRepository) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := u.userDao.FindUserByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return u.toDomainUser(user), nil
}

func (u *CacheUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := u.userDao.InsertUser(ctx, u.toModelUser(user))
	if err != nil {
		return err
	}

	return u.userCache.SetUser(ctx, user)
}

func (u *CacheUserRepository) toModelUser(user domain.User) model.User {
	return model.User{
		UserName: user.UserName,
	}
}

func (u *CacheUserRepository) toDomainUser(user model.User) domain.User {
	return domain.User{
		ID:       user.ID,
		UserName: user.UserName,
	}
}
