package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	redisErrors "main/_lib/redis/errors"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func (rc RedisClient) Get(key string) (string, error) {
	value, err := rc.Client.Get(rc.Ctx, key).Result()
	if err != nil {
		return "", redisErrors.NewGetError(key, err)
	}
	return value, nil
}

func (rc RedisClient) SetEx(key string, value string, expireSec uint) error {
	err := rc.Client.Set(rc.Ctx, key, value, time.Duration(expireSec*1e9)).Err()
	if err != nil {
		return redisErrors.NewSetError(key, value, expireSec, err)
	}
	return nil
}

func (rc RedisClient) Set(key string, value string) error {
	return rc.SetEx(key, value, 0)
}
