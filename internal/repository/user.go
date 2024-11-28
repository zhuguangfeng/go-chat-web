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
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]domain.User, error)
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

func (repo *CacheUserRepository) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := repo.userDao.FindUserByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomainUser(user), nil
}

func (repo *CacheUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := repo.userDao.InsertUser(ctx, repo.toModelUser(user))
	if err != nil {
		return err
	}

	return repo.userCache.SetUser(ctx, user)
}

// GetUserByID 根据用户id获取用户信息
func (repo *CacheUserRepository) GetUserByID(ctx context.Context, id int64) (domain.User, error) {
	user, err := repo.userCache.GetUser(ctx, id)
	if err == nil {
		return user, nil
	}

	//TODO 日志
	userM, err := repo.userDao.FindUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomainUser(userM), nil
}

// GetUsersByIDs 根据多个用户id查找用户
func (repo *CacheUserRepository) GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]domain.User, error) {
	users, err := repo.userDao.FindUserByIDs(ctx, ids)
	if err != nil {
		return map[int64]domain.User{}, err
	}
	var res = make(map[int64]domain.User, len(users))
	if len(users) > 0 {
		for _, user := range users {
			res[user.ID] = repo.toDomainUser(user)
		}
	}
	return res, nil
}

func (repo *CacheUserRepository) toModelUser(user domain.User) model.User {
	return model.User{
		UserName: user.UserName,
	}
}

func (repo *CacheUserRepository) toDomainUser(user model.User) domain.User {
	return domain.User{
		ID:       user.ID,
		UserName: user.UserName,
	}
}
