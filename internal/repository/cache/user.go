package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"

	"github.com/zhuguangfeng/go-chat/internal/domain"
	"time"
)

var (
	ErrKeyNotFound = redis.Nil
)

type UserCache interface {
	SetUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, userId int64) (domain.User, error)
}

type RedisUserCache struct {
	redisCli   redis.Cmdable
	keyPrefix  string
	expiration time.Duration
}

func NewUserCache(redisCli redis.Cmdable) UserCache {
	return &RedisUserCache{
		redisCli:   redisCli,
		keyPrefix:  "user:userInfo:userId_",
		expiration: time.Hour * 24 * 7,
	}
}

// SetUser 缓存用户信息
func (cache *RedisUserCache) SetUser(ctx context.Context, user domain.User) error {
	key := fmt.Sprintf("%s%d:", cache.keyPrefix, user.ID)
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return cache.redisCli.Set(ctx, key, val, cache.expiration).Err()
}

// GetUser 获取用户信息
func (cache *RedisUserCache) GetUser(ctx context.Context, userID int64) (domain.User, error) {
	key := fmt.Sprintf("%s%d:", cache.keyPrefix, userID)

	userBytes, err := cache.redisCli.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}

	var res domain.User
	err = json.Unmarshal(userBytes, &res)

	return res, err
}
