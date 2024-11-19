package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"time"
)

type UserCache interface {
	SetUser(ctx context.Context, user domain.User) error
}

type RedisUserCache struct {
	redisCli   redis.Cmdable
	keyPrefix  string
	expiration time.Duration
}

func NewUserCache(redisCli redis.Cmdable) UserCache {
	return &RedisUserCache{
		redisCli:   redisCli,
		keyPrefix:  "user:userId_",
		expiration: time.Hour * 24 * 7,
	}
}

// SetUser 缓存用户信息
func (u *RedisUserCache) SetUser(ctx context.Context, user domain.User) error {
	key := fmt.Sprintf("%s%d:", u.keyPrefix, user.ID)

	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return u.redisCli.Set(ctx, key, val, u.expiration).Err()
}
