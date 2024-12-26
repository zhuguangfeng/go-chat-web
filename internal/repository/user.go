package repository

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

var (
	ErrKeyNotExists = redis.Nil
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByPhone(ctx context.Context, phone string) (domain.User, error)
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]domain.User, error)
}

type CacheUserRepository struct {
	logger    logger.Logger
	userDao   dao.UserDao
	userCache cache.UserCache
}

func NewUserRepository(logger logger.Logger, userDao dao.UserDao, userCache cache.UserCache) UserRepository {
	return &CacheUserRepository{
		logger:    logger,
		userDao:   userDao,
		userCache: userCache,
	}
}

// CreateUser 创建用户
func (repo *CacheUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := repo.userDao.InsertUser(ctx, repo.toModelUser(user))
	if err != nil {
		if errors.Is(err, dao.ErrUserDuplicate) {
			return errorx.NewBizError(common.UserNotFound).WithError(err)
		}
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

// GetUserByPhone 根据手机号获取用户
func (repo *CacheUserRepository) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := repo.userDao.FindUserByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return domain.User{}, errorx.NewBizError(common.UserPhoneNotFound).WithError(err)
		}
		return domain.User{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toDomainUser(user), nil
}

// GetUserByID 根据用户id获取用户信息
func (repo *CacheUserRepository) GetUserByID(ctx context.Context, id int64) (domain.User, error) {
	user, err := repo.userCache.GetUser(ctx, id)
	if err == nil {
		return user, nil
	}

	repo.logger.Error("[CacheUserRepo.GetUserByID]从redis获取用户缓存失败", logger.Int64("id", id), logger.Error(err))

	userM, err := repo.userDao.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return domain.User{}, errorx.NewBizError(common.UserNotFound).WithError(err)
		}
		return domain.User{}, err
	}

	err = repo.userCache.SetUser(ctx, repo.toDomainUser(userM))
	if err != nil {
		repo.logger.Error("[CacheUserRepo.GetUserByID]回写用户信息到redis失败", logger.Int64("id", id), logger.Error(err))
	}
	return repo.toDomainUser(userM), nil

}

// GetUsersByIDs 根据多个用户id查找用户
func (repo *CacheUserRepository) GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]domain.User, error) {
	users, err := repo.userDao.FindUserByIDs(ctx, ids)
	if err != nil {
		return map[int64]domain.User{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
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
		Base: model.Base{
			ID: user.ID,
		},
		Username:      user.Username,
		Password:      user.Password,
		Phone:         user.Phone,
		Age:           user.Age,
		Gender:        user.Gender,
		IsRealName:    user.IsRealName,
		Name:          user.Name,
		IDCard:        user.IDCard,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		Status:        user.Status,
	}
}

func (repo *CacheUserRepository) toDomainUser(user model.User) domain.User {
	return domain.User{
		ID:            user.ID,
		Username:      user.Username,
		Password:      user.Password,
		Phone:         user.Phone,
		Age:           user.Age,
		Gender:        user.Gender,
		IsRealName:    user.IsRealName,
		Name:          user.Name,
		IDCard:        user.IDCard,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		Status:        user.Status,
		CreatedTime:   user.CreatedAt,
		UpdatedTime:   user.UpdatedAt,
	}
}
